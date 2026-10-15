package voxgigresendsdk

import (
	"github.com/mateen-ali12/resend-sdk/go/core"
	"github.com/mateen-ali12/resend-sdk/go/entity"
	"github.com/mateen-ali12/resend-sdk/go/feature"
	_ "github.com/mateen-ali12/resend-sdk/go/utility"
)

// Type aliases preserve external API.
type ResendSDK = core.ResendSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type ResendEntity = core.ResendEntity
type StreamItem = core.StreamItem
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type ResendError = core.ResendError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewApiKeyEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewApiKeyEntity(client, entopts)
	}
	core.NewAudienceEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewAudienceEntity(client, entopts)
	}
	core.NewAutomationEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewAutomationEntity(client, entopts)
	}
	core.NewAutomationRunEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewAutomationRunEntity(client, entopts)
	}
	core.NewBroadcastEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewBroadcastEntity(client, entopts)
	}
	core.NewBroadcastClickedLinkEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewBroadcastClickedLinkEntity(client, entopts)
	}
	core.NewBroadcastRecipientEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewBroadcastRecipientEntity(client, entopts)
	}
	core.NewContactEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactEntity(client, entopts)
	}
	core.NewContactImportEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactImportEntity(client, entopts)
	}
	core.NewContactPropertyEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactPropertyEntity(client, entopts)
	}
	core.NewContactSegmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactSegmentEntity(client, entopts)
	}
	core.NewContactTopicEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactTopicEntity(client, entopts)
	}
	core.NewDomainEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewDomainEntity(client, entopts)
	}
	core.NewDomainClaimEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewDomainClaimEntity(client, entopts)
	}
	core.NewEmailEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEmailEntity(client, entopts)
	}
	core.NewEmailAttachmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEmailAttachmentEntity(client, entopts)
	}
	core.NewEmailBatchEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEmailBatchEntity(client, entopts)
	}
	core.NewEmailsMetricEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEmailsMetricEntity(client, entopts)
	}
	core.NewEventEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEventEntity(client, entopts)
	}
	core.NewInboxEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxEntity(client, entopts)
	}
	core.NewInboxAgentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxAgentEntity(client, entopts)
	}
	core.NewInboxDraftEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxDraftEntity(client, entopts)
	}
	core.NewInboxLabelEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxLabelEntity(client, entopts)
	}
	core.NewInboxThreadEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxThreadEntity(client, entopts)
	}
	core.NewInboxThreadEmailEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxThreadEmailEntity(client, entopts)
	}
	core.NewLogEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewLogEntity(client, entopts)
	}
	core.NewOauthGrantEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewOauthGrantEntity(client, entopts)
	}
	core.NewReceivedEmailEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewReceivedEmailEntity(client, entopts)
	}
	core.NewReceivedEmailAttachmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewReceivedEmailAttachmentEntity(client, entopts)
	}
	core.NewSegmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSegmentEntity(client, entopts)
	}
	core.NewSegmentContactEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSegmentContactEntity(client, entopts)
	}
	core.NewSuppressionEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSuppressionEntity(client, entopts)
	}
	core.NewSuppressionBatchEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSuppressionBatchEntity(client, entopts)
	}
	core.NewTemplateEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewTemplateEntity(client, entopts)
	}
	core.NewTopicEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewTopicEntity(client, entopts)
	}
	core.NewUsageEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUsageEntity(client, entopts)
	}
	core.NewWebhookEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewWebhookEntity(client, entopts)
	}
	core.NewWebhookEventEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewWebhookEventEntity(client, entopts)
	}
	core.NewWebhookEventAttemptEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewWebhookEventAttemptEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewResendSDK = core.NewResendSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewResendSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *ResendSDK  { return NewResendSDK(nil) }
func Test() *ResendSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
