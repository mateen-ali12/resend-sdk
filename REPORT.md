# Mini task 1: Resend SDK with the Voxgig SDK generator

**Author:** Mateen Ali · **Date:** 15 October 2026
**Repository:** https://github.com/mateen-ali12/resend-sdk (MIT)
**API:** [Resend](https://resend.com), an email API with a free tier and an official MIT-licensed OpenAPI 3.1 spec ([resend/resend-openapi](https://github.com/resend/resend-openapi)). It isn't in the voxgig-sdk catalogue (checked against all 804 repositories).

## Result

| | First generate (heuristic only) | After `guide.aontu` corrections |
|---|---|---|
| Entities | 81 | 39 |
| TypeScript tests | 664 pass, **3 fail** | 575 pass, 0 fail (45 skipped: live-only) |
| Go tests | **1 fail** | all pass |
| Python tests | not run | 413 pass, 0 fail (124 skipped: live-only) |
| `voxgig-sdkgen doctor` | n/a | passes (after removing 56 orphan files, see 5) |

The repository has TypeScript, Python and Go SDKs, offline test suites, generated READMEs and reference docs. I also added `examples/live-smoke.js`. It runs the TypeScript SDK against the real API: it sends a single email and a batch, reads the account, and runs a create, update, load, list and remove cycle on a Topic it creates for the purpose.

**Live check** (against api.resend.com, sending to Resend's `delivered@resend.dev` test address):

```
PASS  Email.create (send) 01a13f1b-b824-7b54-8105-b81386154684
PASS  EmailBatch.create (batch send) {"data":[{"id":"01a13f1b-bb4d-…"},{"id":"01a13f1b-bb4d-…"}]}
SKIPPED Email.load - ResendSDK: load: request: 401: Unauthorized This API key is restricted to only send emails
…
2 passed, 0 failed, 5 skipped
```

The key on my account is a "Sending access" key, so the read and Topic steps were refused by Resend (401) rather than failing in the SDK. With a "Full access" key the same script runs all 10 steps.

## How I worked, and how I used AI

I used Claude Code (Claude Opus) as the operator and directed it step by step: pick an API, scaffold, generate, test, correct the model, re-test. I checked its work in four ways:

1. **Diff review.** I read every line of the guide corrections and checked each path against the spec.
2. **The generator's own gates.** I ran all three test suites, which include the README-example tests, plus `doctor`. I only accepted changes after they went green.
3. **A wire check.** I called the API with an invalid key through both the SDK and `curl`. Both got the same `400` from Resend, which confirms the base URL and auth header.
4. **Live calls** with a real API key (above). The script creates its own test data and always removes it.

The SDK is generated, and I didn't edit any target code by hand. Every change is in `.sdk/model/guide/guide.aontu` or `.sdk/model/project.aontu`.

Time: about 30 minutes of human work, within the 30-minute box. Machine time was small: scaffolding took 19 s and each generate run about 8 s.

## Observations, most important first

### 1. The entity heuristic splits resources on per-operation response schema names

Resend names most DELETE and PATCH responses per operation (`RemoveContactResponseSuccess`, `UpdateTopicResponseSuccess`, and so on). apidef turned each of those into a separate entity, so the first generate produced 81 entities for about 20 real resources. `Contact` got only `GET /contacts/{id}`, while the PATCH and DELETE on the same path became `UpdateContactResponseSuccess` and `RemoveContactResponseSuccess`. The collection move made it worse: list and create on `/templates`, `/topics`, `/segments` and `/suppressions` landed on the `remove_*` entities, and `/webhooks` on `update_webhook`. The first README showcased `client.AddContactToSegmentResponseSuccess().create(...)`.

The guide fixed all of this, and the override design worked first time: defaults written as `*`, and `active: false` to keep a record of what was switched off. But it took 56 entity switch-offs and 53 path assignments. That's the bulk of the work, and a newcomer wouldn't know it was needed.

Suggestions:
- When an item path already has an entity from its `GET`, attach that path's `PATCH`/`DELETE` to the same entity, whatever the response schema is called.
- Strip affixes like `(Create|Update|Remove|List|Get)…(ResponseSuccess|Response)` before naming.
- apidef already prints `emr=0.583`. Below a threshold, warn and print the likely guide corrections, ready to paste.

### 2. A freshly generated project fails its own tests

Before any edits, 3 TypeScript tests and 1 Go test failed:
- A README snippet calls `Email().load()` without `email_id`.
- The `RemoveSuppressionResponseSuccess` remove test expects the removed record to be gone from a list that belongs to a different entity.

The documented guarantee is that green tests mean the docs are correct, so a red first run is the worst first impression. It is a symptom of observation 1, but it should degrade to a skip or a warning, not a failure.

### 3. Python drops `from`, the one field you need to send an email

Generation logs `entity-types-skip-field py: field "from" ... has no legal class-syntax TypedDict key form` 19 times (14 after the corrections). `from` is a Python keyword, so it's left out of the typed model. The functional form, `TypedDict('EmailCreateData', {'from': str, ...})`, accepts keyword keys, so the field could stay typed.

### 4. Request types are looser than the spec

- `EmailCreateData` makes `from`, `to` and `subject` optional, although the spec marks them `required`.
- It includes response-only fields (`id`, `created_at`, `last_event`).
- `to` is typed `any[]`, but the spec says `oneOf: [string, string[]]`. So `to: 'a@b.com'`, the most common call, fails type-checking.
- The README's Email create example is `client.Email().create({})`. That passes in test mode but would fail against the real API.

### 5. `doctor` misreports orphan files

When entities are switched off, apidef deletes their entity files but leaves 56 flow files behind. `doctor` then exits 1 with: ".sdk has drifted from the scaffold: 0 forked, 0 edited, 0 stale, 0 missing. Run `voxgig-sdkgen target add <lang>` to resync". The counts say nothing is wrong, and the suggested fix doesn't remove orphans. I deleted them by hand, and `doctor` then passed. Suggestion: apidef cleans up flow files with their entity, and `doctor` names "orphans" in its summary with a `--fix` option.

### 6. API error messages are hidden

An invalid key gives `err.message = "ResendSDK: list: request: 400: Bad Request"`. Resend's actual message, "API key is invalid", is only on `err.result.body.message`. A sending-only key gives "401: Unauthorized" in the same way, and the useful part, "This API key is restricted to only send emails", is again only in the body. Appending the body's message (or `error`/`detail`) to `err.message` would make most failures self-explanatory.

Two related naming points came up in the live run:
- The SDK and its live tests read `RESEND_APIKEY`, but Resend's own docs use `RESEND_API_KEY`, which is what I had in my `.env`. I had to map one to the other. Accepting the provider's conventional name too, or naming the expected variable in the error, would save that step.
- `POST /emails/batch` takes a bare JSON array. The SDK sends it from an invented field, `EmailBatch().create({ send_emails: [...] })`. That works, but I could only find the name in `.sdk/model/entity/email_batch.aontu`.

### 7. Defaults assume publishing into the voxgig-sdk organisation

Out of the box:
- the LICENSE says "Copyright (c) 2026 Voxgig"
- install links and the Go module path point to `github.com/voxgig-sdk/resend-sdk`
- `.sdk/package.json` lists the author as "Resend", which wrongly suggests affiliation

Fixing it is easy (`main: kit: repo`, `publisher` and `contributor` in `project.aontu`), but I only found those keys by reading `node_modules/@voxgig/sdkgen/model/sdkgen.aontu`. Suggestion: `create-sdkgen --repo owner/name --publisher "Name"`, or ask interactively.

### 8. Documentation is hard to find from the starting point

voxgig.com/sdk shows `npm create @voxgig/sdkgen`. The fix I needed for observation 1 is in `apidef/docs/reference/guide.md`, three links away. The generated `guide.aontu` is two include lines with no hint of what can go there. Suggestion: scaffold `guide.aontu` with commented examples (switch off, add an op, rename a param, allowlist) and a link to the reference.

Two smaller points:
- The Node 24 requirement is only in AGENTS.md. On Node 22 everything still worked, with `EBADENGINE` warnings.
- 9 `flow-step-unreachable` warnings remain after correction. They say what's wrong but not how to fix it.

### 9. The README tutorial picks entities alphabetically

The tutorial's "first API call" uses `ApiKey` and `AutomationRun`. For an email API the obvious first call is `Email().create()`. On the uncorrected model, the tutorial also skipped step 2. A `primary entity` setting, or ranking entities by CRUD completeness, would make the first page much stronger.

## What worked well

- **Speed.** 19 s to scaffold and about 8 s per regenerate, so the edit loop is effectively instant.
- **Corrections in the model.** The model-first design means every correction lives in two small files that survive regeneration. I never touched generated code.
- **Offline tests by default.** Hundreds of tests in three languages, with no server, key or network.
- **Generated extras.** The "unofficial, not affiliated" disclaimer, MIT LICENSE, SECURITY.md and AGENTS.md all come for free, and the agent docs are good enough that an AI assistant can drive the toolchain from them.

## Where I stopped

I didn't add the `go-cli` and `go-mcp` targets or the PHP, Ruby and Lua ports, to stay within the time box. I also didn't check every inbox endpoint (a beta Resend feature) against the live API, only the email, domain and API-key paths in the smoke test.
