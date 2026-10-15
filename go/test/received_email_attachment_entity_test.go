package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/mateen-ali12/resend-sdk/go"
	"github.com/mateen-ali12/resend-sdk/go/core"

	vs "github.com/mateen-ali12/resend-sdk/go/utility/struct"
)

// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const received_email_attachmentEntityLiveStrict = true


func TestReceivedEmailAttachmentEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.ReceivedEmailAttachment(nil)
		if ent == nil {
			t.Fatal("expected non-nil ReceivedEmailAttachmentEntity")
		}
	})

	t.Run("validate", func(t *testing.T) {
		if !fhHasFeature("validate") {
			t.Skip("feature not present in this SDK: validate")
		}
		client := sdk.TestSDK(nil, map[string]any{
			"feature": map[string]any{"validate": map[string]any{"active": true}},
		})
		_, err := client.ReceivedEmailAttachment(nil).List(map[string]any{"email_id": 1}, nil)
		if sdkerr, ok := err.(*core.ResendError); !ok || "validate_failed" != sdkerr.Code {
			t.Fatalf("expected validate_failed, got %v", err)
		}
	})

	t.Run("basic", func(tt *testing.T) {
		var t testing.TB = tt
		setup := received_email_attachmentBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "received_email_attachment." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		if setup.live {
			for _, _liveKey := range []string{"email01"} {
				if setup.syntheticOnly || setup.idmap[_liveKey] == nil {
					liveMiss(t, received_email_attachmentEntityLiveStrict, "Live entity test blocked: needs %s via RESEND_TEST_RECEIVED_EMAIL_ATTACHMENT_ENTID", _liveKey)
				}
			}
		}
		client := setup.client
		_ = client
		if setup.live {
			liveExisting(t, received_email_attachmentEntityLiveStrict, setup.data, "received_email_attachment", func() (any, error) {
				return client.ReceivedEmailAttachment(nil).List(map[string]any{"email_id": setup.idmap["email01"]}, nil)
			})
		}

		// Bootstrap entity data from existing test data (no create step in flow).
		receivedEmailAttachmentRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath(setup.data, "existing.received_email_attachment")))
		var receivedEmailAttachmentRef01Data map[string]any
		if len(receivedEmailAttachmentRef01DataRaw) > 0 {
			receivedEmailAttachmentRef01Data = core.ToMapAny(receivedEmailAttachmentRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = receivedEmailAttachmentRef01Data

		// LIST
		receivedEmailAttachmentRef01Ent := client.ReceivedEmailAttachment(nil)
		receivedEmailAttachmentRef01Match := map[string]any{
			"email_id": setup.idmap["email01"],
		}

		receivedEmailAttachmentRef01ListResult, err := receivedEmailAttachmentRef01Ent.List(receivedEmailAttachmentRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, receivedEmailAttachmentRef01ListOk := receivedEmailAttachmentRef01ListResult.([]any)
		if !receivedEmailAttachmentRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", receivedEmailAttachmentRef01ListResult)
		}

		// LOAD
		receivedEmailAttachmentRef01MatchDt0 := map[string]any{
			"id": receivedEmailAttachmentRef01Data["id"],
		}
		if v := setup.idmap["email_id"]; v != nil {
			receivedEmailAttachmentRef01MatchDt0["email_id"] = v
		} else {
			receivedEmailAttachmentRef01MatchDt0["email_id"] = setup.idmap["email01"]
		}
		receivedEmailAttachmentRef01DataDt0Loaded, err := receivedEmailAttachmentRef01Ent.Load(receivedEmailAttachmentRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		receivedEmailAttachmentRef01DataDt0LoadResult := core.ToMapAny(entityData(receivedEmailAttachmentRef01DataDt0Loaded))
		if receivedEmailAttachmentRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if receivedEmailAttachmentRef01DataDt0LoadResult["id"] != receivedEmailAttachmentRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func received_email_attachmentBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "received_email_attachment", "ReceivedEmailAttachmentTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read received_email_attachment test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse received_email_attachment test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"received_email_attachment01", "received_email_attachment02", "received_email_attachment03", "email01"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Whether *_ENTID supplied the idmap, read before envOverride consumes it:
	// without it, the ids a live flow binds are the fixture's synthetic ones.
	entidEnvRaw := os.Getenv("RESEND_TEST_RECEIVED_EMAIL_ATTACHMENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"RESEND_TEST_RECEIVED_EMAIL_ATTACHMENT_ENTID": idmap,
		"RESEND_TEST_LIVE":      "FALSE",
		"RESEND_TEST_EXPLAIN":   "FALSE",
		"RESEND_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["RESEND_TEST_RECEIVED_EMAIL_ATTACHMENT_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["RESEND_TEST_LIVE"] == "TRUE" {
		// An empty map, not a nil one: Merge returns nil when its last entry
		// is nil, and BasicSetup is normally called with no extras - so a
		// bare nil silently discarded the apikey and server values below.
		extraOpts := extra
		if extraOpts == nil {
			extraOpts = map[string]any{}
		}

		mergedOpts := vs.Merge([]any{
			// liveClientOptions() FIRST, so the generated fields below win:
			// sdk-test-control.json's test.client.options adds to the live
			// client, it does not redirect it.
			liveClientOptions(),
			map[string]any{
				"apikey": env["RESEND_APIKEY"],
			},
			extraOpts,
		})
		client = sdk.NewResendSDK(core.ToMapAny(mergedOpts))
	}

	live := env["RESEND_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["RESEND_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
