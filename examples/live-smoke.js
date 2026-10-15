// Live smoke test against the real Resend API, using the generated TS SDK.
//
//   cd ts && npm install && npm run build && cd ..
//   RESEND_APIKEY=re_... RESEND_TO=you@example.com node examples/live-smoke.js
//
// RESEND_TO must be the email address of your Resend account: without a
// verified domain, Resend only delivers from onboarding@resend.dev to it.

const { ResendSDK } = require('../ts')

async function main() {
  const apikey = process.env.RESEND_APIKEY
  const to = process.env.RESEND_TO
  if (!apikey || !to) {
    throw new Error('Set RESEND_APIKEY and RESEND_TO')
  }

  const client = new ResendSDK({ apikey })

  const domains = await client.Domain().list()
  console.log('domains:', domains.map((d) => d.data().name))

  const keys = await client.ApiKey().list()
  console.log('api keys:', keys.map((k) => k.data().name))

  const sent = await client.Email().create({
    from: 'onboarding@resend.dev',
    to: [to],
    subject: 'Hello from the Voxgig-generated Resend SDK',
    html: '<p>Sent with <strong>@voxgig-sdk/resend-sdk</strong>.</p>',
  })
  const id = sent.data().id
  console.log('sent email id:', id)

  const loaded = await client.Email().load({ id })
  console.log('loaded email:', loaded.data().subject, '/', loaded.data().last_event)

  const emails = await client.Email().list({ limit: 3 })
  console.log('recent emails:', emails.length)
}

main().catch((err) => {
  // The API's own message is on err.result.body, not in err.message.
  console.error('FAILED:', err.message, err.result?.body?.message ?? '')
  process.exit(1)
})
