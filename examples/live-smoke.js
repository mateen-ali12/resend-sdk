// Live smoke test against the real Resend API, using the generated TS SDK.
//
//   cd ts && npm install && npm run build && cd ..
//   RESEND_APIKEY=re_... node examples/live-smoke.js
//
// It sends one email, reads the account, and runs a full create/load/update/
// list/remove cycle on a Topic it creates for the purpose (and always removes).
// With a "Sending access" key only the send runs; the rest report SKIPPED.
//
// RESEND_TO defaults to delivered@resend.dev, Resend's test address, which
// simulates a delivery without a verified domain and without emailing anyone.

const { ResendSDK } = require('../ts')

const results = []

async function step(name, fn) {
  try {
    const out = await fn()
    results.push([name, 'PASS'])
    console.log('PASS ', name, out === undefined ? '' : out)
    return out
  } catch (err) {
    // The API's own message is on err.result.body, not in err.message.
    const msg = err.result?.body?.message ?? ''
    const status = msg.includes('restricted') ? 'SKIPPED' : 'FAIL'
    results.push([name, status])
    console.log(status.padEnd(5), name, '-', err.message, msg)
  }
}

async function main() {
  const apikey = process.env.RESEND_APIKEY || process.env.RESEND_API_KEY
  const to = process.env.RESEND_TO || 'delivered@resend.dev'
  if (!apikey) {
    throw new Error('Set RESEND_APIKEY')
  }

  const client = new ResendSDK({ apikey })

  const emailId = await step('Email.create (send)', async () => {
    const sent = await client.Email().create({
      from: 'onboarding@resend.dev',
      to: [to],
      subject: 'Hello from the Voxgig-generated Resend SDK',
      html: '<p>Sent with <strong>@voxgig-sdk/resend-sdk</strong>.</p>',
    })
    return sent.data().id
  })

  await step('EmailBatch.create (batch send)', async () => {
    const batch = await client.EmailBatch().create({
      send_emails: [1, 2].map((n) => ({
        from: 'onboarding@resend.dev',
        to: [to],
        subject: 'Batch email ' + n + ' from the Voxgig-generated Resend SDK',
        text: 'Batch email ' + n,
      })),
    })
    return JSON.stringify(batch.data())
  })

  if (emailId) {
    await step('Email.load', async () => {
      const loaded = await client.Email().load({ id: emailId })
      return loaded.data().subject + ' / ' + loaded.data().last_event
    })
  }

  await step('Email.list', async () => (await client.Email().list({ limit: 3 })).length)
  await step('Domain.list', async () => (await client.Domain().list()).map((d) => d.data().name))
  await step('ApiKey.list', async () => (await client.ApiKey().list()).length)

  const topicId = await step('Topic.create', async () => {
    const topic = await client.Topic().create({
      name: 'voxgig-sdk-smoke-' + Date.now(),
      default_subscription: 'opt_out',
    })
    return topic.data().id
  })

  if (topicId) {
    await step('Topic.update', async () => {
      await client.Topic().update({ id: topicId, description: 'updated by smoke test' })
    })
    await step('Topic.load', async () => (await client.Topic().load({ id: topicId })).data().description)
    await step('Topic.list', async () =>
      (await client.Topic().list()).some((t) => t.data().id === topicId))
    await step('Topic.remove', async () => {
      await client.Topic().remove({ id: topicId })
    })
  }

  const failed = results.filter(([, s]) => s === 'FAIL').length
  const skipped = results.filter(([, s]) => s === 'SKIPPED').length
  console.log(`\n${results.length - failed - skipped} passed, ${failed} failed, ${skipped} skipped`)
  process.exit(failed ? 1 : 0)
}

main().catch((err) => {
  console.error('FAILED:', err.message)
  process.exit(1)
})
