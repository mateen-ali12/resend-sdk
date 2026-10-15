

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { ResendSDK, BaseFeature, config, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('SuppressionBatchEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when RESEND_TEST_LIVE=TRUE.
  afterEach(liveDelay('RESEND_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ResendSDK.test()
    const ent = testsdk.SuppressionBatch()
    assert(null != ent)
  })




  test('basic', async (t) => {

    const live = 'TRUE' === process.env.RESEND_TEST_LIVE
    for (const op of []) {
      if (!live && maybeSkipControl(t, 'entityOp', 'suppression_batch.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{},"name":"suppression_batch","op":{"create":{"input":"data","name":"create","points":[{"a":true,"bf":["emails"],"co":{"id":"POST /suppressions/batch/add","source":"openapi3","version":2},"g":{},"k":"http","m":"POST","o":"/suppressions/batch/add","q":{"$action":"add"},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"suppressions"},{"lit":"batch"},{"lit":"add"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"bf":["emails","ids"],"co":{"id":"POST /suppressions/batch/remove","source":"openapi3","version":2},"g":{},"k":"http","m":"POST","o":"/suppressions/batch/remove","q":{"$action":"remove"},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"suppressions"},{"lit":"batch"},{"lit":"remove"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"suppression_batch","name__orig":"suppression_batch","Name":"SuppressionBatch","name_":"suppression_batch","name-":"suppression-batch","NAME":"SUPPRESSION_BATCH","index$":32}, {"active":true,"entity":"suppression_batch","key$":"BasicSuppressionBatchFlow","kind":"basic","name":"BasicSuppressionBatchFlow","param":{},"step":[{"a":false,"d":{},"i":{"ref":"suppression_batch_ref01"},"m":{},"o":"create","s":[],"v":[],"unreachable":true}]}, 'SuppressionBatch', {"POST /suppressions/batch/add":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","required":["emails"],"properties":{"emails":{"type":"array","minItems":1,"maxItems":100,"items":{"type":"string"},"description":"Email addresses to suppress.","example":["steve.wozniak@gmail.com"]}},"x-ref":"#/components/schemas/BatchAddSuppressionsOptions"}}}},"parameters":[]},"POST /suppressions/batch/remove":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","description":"Provide either `emails` or `ids`, but not both.","properties":{"emails":{"type":"array","minItems":1,"maxItems":100,"items":{"type":"string"},"description":"Email addresses to remove from the suppression list.","example":["steve.wozniak@gmail.com"]},"ids":{"type":"array","minItems":1,"maxItems":100,"items":{"type":"string"},"description":"Suppression IDs to remove from the suppression list.","example":["e169aa45-1ecf-4183-9955-b1499d5701d3"]}},"x-ref":"#/components/schemas/BatchRemoveSuppressionsOptions"}}}},"parameters":[]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let suppression_batch_ref01_data = Object.values(setup.data.existing.suppression_batch)[0] as any

  })
})



// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const LIVE_STRICT = true

function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/suppression_batch/SuppressionBatchTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = ResendSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['suppression_batch01','suppression_batch02','suppression_batch03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'RESEND_TEST_SUPPRESSION_BATCH_ENTID': idmap,
    'RESEND_TEST_LIVE': 'FALSE',
    'RESEND_TEST_EXPLAIN': 'FALSE',
    'RESEND_APIKEY': '',
  })

  idmap = env['RESEND_TEST_SUPPRESSION_BATCH_ENTID']

  const live = 'TRUE' === env.RESEND_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['RESEND_TEST_SUPPRESSION_BATCH_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new ResendSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
        apikey: env.RESEND_APIKEY,
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.RESEND_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
