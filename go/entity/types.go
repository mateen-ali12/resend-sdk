// Typed models for the Resend SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/mateen-ali12/resend-sdk/go/core"
)

// ApiKey is the typed data model for the api_key entity.
type ApiKey struct {
}

// ApiKeyListMatch is the typed request payload for ApiKey.ListTyped.
type ApiKeyListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ApiKeyCreateData is the typed request payload for ApiKey.CreateTyped.
type ApiKeyCreateData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	DomainId *string `json:"domain_id,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUsedAt *any `json:"last_used_at,omitempty"`
	Name string `json:"name"`
	Object *string `json:"object,omitempty"`
	Permission *string `json:"permission,omitempty"`
}

// ApiKeyUpdateData is the typed request payload for ApiKey.UpdateTyped.
type ApiKeyUpdateData struct {
	Id string `json:"id"`
	CreatedAt *string `json:"created_at,omitempty"`
	DomainId *string `json:"domain_id,omitempty"`
	LastUsedAt *any `json:"last_used_at,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	Permission *string `json:"permission,omitempty"`
}

// ApiKeyRemoveMatch is the typed request payload for ApiKey.RemoveTyped.
type ApiKeyRemoveMatch struct {
	Id string `json:"id"`
}

// Audience is the typed data model for the audience entity.
type Audience struct {
}

// AudienceLoadMatch is the typed request payload for Audience.LoadTyped.
type AudienceLoadMatch struct {
	Id string `json:"id"`
}

// AudienceListMatch is the typed request payload for Audience.ListTyped.
type AudienceListMatch struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
}

// AudienceCreateData is the typed request payload for Audience.CreateTyped.
type AudienceCreateData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
}

// AudienceRemoveMatch is the typed request payload for Audience.RemoveTyped.
type AudienceRemoveMatch struct {
	Id string `json:"id"`
}

// Automation is the typed data model for the automation entity.
type Automation struct {
}

// AutomationLoadMatch is the typed request payload for Automation.LoadTyped.
type AutomationLoadMatch struct {
	Id string `json:"id"`
}

// AutomationListMatch is the typed request payload for Automation.ListTyped.
type AutomationListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Status *string `json:"status,omitempty"`
}

// AutomationCreateData is the typed request payload for Automation.CreateTyped.
type AutomationCreateData struct {
	Connections *[]any `json:"connections,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	Status *string `json:"status,omitempty"`
	Steps *[]any `json:"steps,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// AutomationUpdateData is the typed request payload for Automation.UpdateTyped.
type AutomationUpdateData struct {
	Id string `json:"id"`
	Connections *[]any `json:"connections,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	Status *string `json:"status,omitempty"`
	Steps *[]any `json:"steps,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// AutomationRemoveMatch is the typed request payload for Automation.RemoveTyped.
type AutomationRemoveMatch struct {
	Id string `json:"id"`
}

// AutomationRun is the typed data model for the automation_run entity.
type AutomationRun struct {
}

// AutomationRunLoadMatch is the typed request payload for AutomationRun.LoadTyped.
type AutomationRunLoadMatch struct {
	AutomationId string `json:"automation_id"`
	Id string `json:"id"`
}

// AutomationRunListMatch is the typed request payload for AutomationRun.ListTyped.
type AutomationRunListMatch struct {
	AutomationId string `json:"automation_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Broadcast is the typed data model for the broadcast entity.
type Broadcast struct {
}

// BroadcastLoadMatch is the typed request payload for Broadcast.LoadTyped.
type BroadcastLoadMatch struct {
	Id string `json:"id"`
}

// BroadcastListMatch is the typed request payload for Broadcast.ListTyped.
type BroadcastListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// BroadcastCreateData is the typed request payload for Broadcast.CreateTyped.
type BroadcastCreateData struct {
	AudienceId *any `json:"audience_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	From *string `json:"from,omitempty"`
	Html *any `json:"html,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	PreviewText *string `json:"preview_text,omitempty"`
	ReplyTo *[]any `json:"reply_to,omitempty"`
	ScheduledAt *string `json:"scheduled_at,omitempty"`
	SegmentId *any `json:"segment_id,omitempty"`
	Send *bool `json:"send,omitempty"`
	SentAt *string `json:"sent_at,omitempty"`
	Status *string `json:"status,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Text *any `json:"text,omitempty"`
	TopicId *any `json:"topic_id,omitempty"`
}

// BroadcastUpdateData is the typed request payload for Broadcast.UpdateTyped.
type BroadcastUpdateData struct {
	Id string `json:"id"`
	AudienceId *any `json:"audience_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	From *string `json:"from,omitempty"`
	Html *any `json:"html,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	PreviewText *string `json:"preview_text,omitempty"`
	ReplyTo *[]any `json:"reply_to,omitempty"`
	ScheduledAt *string `json:"scheduled_at,omitempty"`
	SegmentId *any `json:"segment_id,omitempty"`
	Send *bool `json:"send,omitempty"`
	SentAt *string `json:"sent_at,omitempty"`
	Status *string `json:"status,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Text *any `json:"text,omitempty"`
	TopicId *any `json:"topic_id,omitempty"`
}

// BroadcastRemoveMatch is the typed request payload for Broadcast.RemoveTyped.
type BroadcastRemoveMatch struct {
	Id string `json:"id"`
}

// BroadcastClickedLink is the typed data model for the broadcast_clicked_link entity.
type BroadcastClickedLink struct {
}

// BroadcastClickedLinkListMatch is the typed request payload for BroadcastClickedLink.ListTyped.
type BroadcastClickedLinkListMatch struct {
	BroadcastId string `json:"broadcast_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// BroadcastRecipient is the typed data model for the broadcast_recipient entity.
type BroadcastRecipient struct {
}

// BroadcastRecipientListMatch is the typed request payload for BroadcastRecipient.ListTyped.
type BroadcastRecipientListMatch struct {
	BroadcastId string `json:"broadcast_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	BounceType *string `json:"bounce_type,omitempty"`
	Email *string `json:"email,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Type string `json:"type"`
}

// Contact is the typed data model for the contact entity.
type Contact struct {
}

// ContactLoadMatch is the typed request payload for Contact.LoadTyped.
type ContactLoadMatch struct {
	Id string `json:"id"`
}

// ContactListMatch is the typed request payload for Contact.ListTyped.
type ContactListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ContactCreateData is the typed request payload for Contact.CreateTyped.
type ContactCreateData struct {
	AudienceId *string `json:"audience_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *any `json:"first_name,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *any `json:"last_name,omitempty"`
	Object *string `json:"object,omitempty"`
	Properties *map[string]any `json:"properties,omitempty"`
	Segments *[]any `json:"segments,omitempty"`
	Topics *[]any `json:"topics,omitempty"`
	Unsubscribed *bool `json:"unsubscribed,omitempty"`
}

// ContactUpdateData is the typed request payload for Contact.UpdateTyped.
type ContactUpdateData struct {
	Id string `json:"id"`
	AudienceId *string `json:"audience_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *any `json:"first_name,omitempty"`
	LastName *any `json:"last_name,omitempty"`
	Object *string `json:"object,omitempty"`
	Properties *map[string]any `json:"properties,omitempty"`
	Segments *[]any `json:"segments,omitempty"`
	Topics *[]any `json:"topics,omitempty"`
	Unsubscribed *bool `json:"unsubscribed,omitempty"`
}

// ContactRemoveMatch is the typed request payload for Contact.RemoveTyped.
type ContactRemoveMatch struct {
	Id string `json:"id"`
}

// ContactImport is the typed data model for the contact_import entity.
type ContactImport struct {
}

// ContactImportLoadMatch is the typed request payload for ContactImport.LoadTyped.
type ContactImportLoadMatch struct {
	Id string `json:"id"`
}

// ContactImportListMatch is the typed request payload for ContactImport.ListTyped.
type ContactImportListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Status *string `json:"status,omitempty"`
}

// ContactImportCreateData is the typed request payload for ContactImport.CreateTyped.
type ContactImportCreateData struct {
	CompletedAt *any `json:"completed_at,omitempty"`
	Counts *map[string]any `json:"counts,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	Status *string `json:"status,omitempty"`
}

// ContactProperty is the typed data model for the contact_property entity.
type ContactProperty struct {
}

// ContactPropertyLoadMatch is the typed request payload for ContactProperty.LoadTyped.
type ContactPropertyLoadMatch struct {
	Id string `json:"id"`
}

// ContactPropertyListMatch is the typed request payload for ContactProperty.ListTyped.
type ContactPropertyListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ContactPropertyCreateData is the typed request payload for ContactProperty.CreateTyped.
type ContactPropertyCreateData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	FallbackValue *any `json:"fallback_value,omitempty"`
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
	Object *string `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ContactPropertyUpdateData is the typed request payload for ContactProperty.UpdateTyped.
type ContactPropertyUpdateData struct {
	Id string `json:"id"`
	CreatedAt *string `json:"created_at,omitempty"`
	FallbackValue *any `json:"fallback_value,omitempty"`
	Key *string `json:"key,omitempty"`
	Object *string `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ContactPropertyRemoveMatch is the typed request payload for ContactProperty.RemoveTyped.
type ContactPropertyRemoveMatch struct {
	Id string `json:"id"`
}

// ContactSegment is the typed data model for the contact_segment entity.
type ContactSegment struct {
}

// ContactSegmentListMatch is the typed request payload for ContactSegment.ListTyped.
type ContactSegmentListMatch struct {
	ContactId string `json:"contact_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ContactSegmentCreateData is the typed request payload for ContactSegment.CreateTyped.
type ContactSegmentCreateData struct {
	ContactId string `json:"contact_id"`
	Id string `json:"id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	SegmentId *string `json:"segment_id,omitempty"`
}

// ContactSegmentRemoveMatch is the typed request payload for ContactSegment.RemoveTyped.
type ContactSegmentRemoveMatch struct {
	ContactId string `json:"contact_id"`
	Id string `json:"id"`
}

// ContactTopic is the typed data model for the contact_topic entity.
type ContactTopic struct {
}

// ContactTopicListMatch is the typed request payload for ContactTopic.ListTyped.
type ContactTopicListMatch struct {
	ContactId string `json:"contact_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ContactTopicUpdateData is the typed request payload for ContactTopic.UpdateTyped.
type ContactTopicUpdateData struct {
	ContactId string `json:"contact_id"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	Subscription *string `json:"subscription,omitempty"`
	Topics *[]any `json:"topics,omitempty"`
}

// Domain is the typed data model for the domain entity.
type Domain struct {
}

// DomainLoadMatch is the typed request payload for Domain.LoadTyped.
type DomainLoadMatch struct {
	Id string `json:"id"`
}

// DomainListMatch is the typed request payload for Domain.ListTyped.
type DomainListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// DomainCreateData is the typed request payload for Domain.CreateTyped.
type DomainCreateData struct {
	Capabilities *map[string]any `json:"capabilities,omitempty"`
	ClickTracking *bool `json:"click_tracking,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomReturnPath *string `json:"custom_return_path,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	OpenTracking *bool `json:"open_tracking,omitempty"`
	Records *[]any `json:"records,omitempty"`
	Region *string `json:"region,omitempty"`
	Status *string `json:"status,omitempty"`
	Tls *string `json:"tls,omitempty"`
	TrackingSubdomain *string `json:"tracking_subdomain,omitempty"`
}

// DomainUpdateData is the typed request payload for Domain.UpdateTyped.
type DomainUpdateData struct {
	Id string `json:"id"`
	Capabilities *map[string]any `json:"capabilities,omitempty"`
	ClickTracking *bool `json:"click_tracking,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomReturnPath *string `json:"custom_return_path,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	OpenTracking *bool `json:"open_tracking,omitempty"`
	Records *[]any `json:"records,omitempty"`
	Region *string `json:"region,omitempty"`
	Status *string `json:"status,omitempty"`
	Tls *string `json:"tls,omitempty"`
	TrackingSubdomain *string `json:"tracking_subdomain,omitempty"`
}

// DomainRemoveMatch is the typed request payload for Domain.RemoveTyped.
type DomainRemoveMatch struct {
	Id string `json:"id"`
}

// DomainClaim is the typed data model for the domain_claim entity.
type DomainClaim struct {
}

// DomainClaimLoadMatch is the typed request payload for DomainClaim.LoadTyped.
type DomainClaimLoadMatch struct {
	Id string `json:"id"`
}

// DomainClaimCreateData is the typed request payload for DomainClaim.CreateTyped.
type DomainClaimCreateData struct {
	BlockedReason *any `json:"blocked_reason,omitempty"`
	ClickTracking *bool `json:"click_tracking,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomReturnPath *string `json:"custom_return_path,omitempty"`
	DomainId *any `json:"domain_id,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	FailureReason *any `json:"failure_reason,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	OpenTracking *bool `json:"open_tracking,omitempty"`
	Record *map[string]any `json:"record,omitempty"`
	Region *any `json:"region,omitempty"`
	Status *string `json:"status,omitempty"`
	TrackingSubdomain *string `json:"tracking_subdomain,omitempty"`
}

// Email is the typed data model for the email entity.
type Email struct {
}

// EmailLoadMatch is the typed request payload for Email.LoadTyped.
type EmailLoadMatch struct {
	Id string `json:"id"`
}

// EmailListMatch is the typed request payload for Email.ListTyped.
type EmailListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// EmailCreateData is the typed request payload for Email.CreateTyped.
type EmailCreateData struct {
	Attachments *[]any `json:"attachments,omitempty"`
	Bcc *[]any `json:"bcc,omitempty"`
	Cc *[]any `json:"cc,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	From *string `json:"from,omitempty"`
	Headers *map[string]any `json:"headers,omitempty"`
	Html *string `json:"html,omitempty"`
	Id *string `json:"id,omitempty"`
	LastEvent *string `json:"last_event,omitempty"`
	MessageId *string `json:"message_id,omitempty"`
	Object *string `json:"object,omitempty"`
	ReplyTo *[]any `json:"reply_to,omitempty"`
	ScheduledAt *string `json:"scheduled_at,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Template *any `json:"template,omitempty"`
	Text *string `json:"text,omitempty"`
	To *[]any `json:"to,omitempty"`
	TopicId *string `json:"topic_id,omitempty"`
}

// EmailUpdateData is the typed request payload for Email.UpdateTyped.
type EmailUpdateData struct {
	Id string `json:"id"`
	Attachments *[]any `json:"attachments,omitempty"`
	Bcc *[]any `json:"bcc,omitempty"`
	Cc *[]any `json:"cc,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	From *string `json:"from,omitempty"`
	Headers *map[string]any `json:"headers,omitempty"`
	Html *string `json:"html,omitempty"`
	LastEvent *string `json:"last_event,omitempty"`
	MessageId *string `json:"message_id,omitempty"`
	Object *string `json:"object,omitempty"`
	ReplyTo *[]any `json:"reply_to,omitempty"`
	ScheduledAt *string `json:"scheduled_at,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Template *any `json:"template,omitempty"`
	Text *string `json:"text,omitempty"`
	To *[]any `json:"to,omitempty"`
	TopicId *string `json:"topic_id,omitempty"`
}

// EmailAttachment is the typed data model for the email_attachment entity.
type EmailAttachment struct {
}

// EmailAttachmentLoadMatch is the typed request payload for EmailAttachment.LoadTyped.
type EmailAttachmentLoadMatch struct {
	EmailId string `json:"email_id"`
	Id string `json:"id"`
}

// EmailAttachmentListMatch is the typed request payload for EmailAttachment.ListTyped.
type EmailAttachmentListMatch struct {
	EmailId string `json:"email_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// EmailBatch is the typed data model for the email_batch entity.
type EmailBatch struct {
}

// EmailBatchCreateData is the typed request payload for EmailBatch.CreateTyped.
type EmailBatchCreateData struct {
	Data *[]any `json:"data,omitempty"`
	SendEmails *[]any `json:"send_emails,omitempty"`
}

// EmailsMetric is the typed data model for the emails_metric entity.
type EmailsMetric struct {
}

// EmailsMetricListMatch is the typed request payload for EmailsMetric.ListTyped.
type EmailsMetricListMatch struct {
	BroadcastId *[]any `json:"broadcast_id,omitempty"`
	Dimension *[]any `json:"dimension,omitempty"`
	DomainId *[]any `json:"domain_id,omitempty"`
	EmailId *[]any `json:"email_id,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
	Granularity *string `json:"granularity,omitempty"`
	Metric *[]any `json:"metric,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// Event is the typed data model for the event entity.
type Event struct {
}

// EventLoadMatch is the typed request payload for Event.LoadTyped.
type EventLoadMatch struct {
	Id string `json:"id"`
}

// EventListMatch is the typed request payload for Event.ListTyped.
type EventListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// EventCreateData is the typed request payload for Event.CreateTyped.
type EventCreateData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	Schema *any `json:"schema,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// EventUpdateData is the typed request payload for Event.UpdateTyped.
type EventUpdateData struct {
	Id string `json:"id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	Schema *any `json:"schema,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// EventRemoveMatch is the typed request payload for Event.RemoveTyped.
type EventRemoveMatch struct {
	Id string `json:"id"`
}

// Inbox is the typed data model for the inbox entity.
type Inbox struct {
}

// InboxLoadMatch is the typed request payload for Inbox.LoadTyped.
type InboxLoadMatch struct {
	Id string `json:"id"`
}

// InboxListMatch is the typed request payload for Inbox.ListTyped.
type InboxListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// InboxCreateData is the typed request payload for Inbox.CreateTyped.
type InboxCreateData struct {
	CreatedAt string `json:"created_at"`
	DomainId string `json:"domain_id"`
	Drafts int `json:"drafts"`
	EmailAddress string `json:"email_address"`
	Forwarding *bool `json:"forwarding,omitempty"`
	FromName any `json:"from_name"`
	Id string `json:"id"`
	LastReceived any `json:"last_received"`
	Name string `json:"name"`
	Object string `json:"object"`
	ReceivingAddress any `json:"receiving_address"`
	Unread int `json:"unread"`
}

// InboxUpdateData is the typed request payload for Inbox.UpdateTyped.
type InboxUpdateData struct {
	Id string `json:"id"`
	CreatedAt *string `json:"created_at,omitempty"`
	DomainId *string `json:"domain_id,omitempty"`
	Drafts *int `json:"drafts,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	Forwarding *bool `json:"forwarding,omitempty"`
	FromName *any `json:"from_name,omitempty"`
	LastReceived *any `json:"last_received,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	ReceivingAddress *any `json:"receiving_address,omitempty"`
	Unread *int `json:"unread,omitempty"`
}

// InboxRemoveMatch is the typed request payload for Inbox.RemoveTyped.
type InboxRemoveMatch struct {
	Id string `json:"id"`
}

// InboxAgent is the typed data model for the inbox_agent entity.
type InboxAgent struct {
}

// InboxAgentLoadMatch is the typed request payload for InboxAgent.LoadTyped.
type InboxAgentLoadMatch struct {
	InboxId string `json:"inbox_id"`
}

// InboxAgentUpdateData is the typed request payload for InboxAgent.UpdateTyped.
type InboxAgentUpdateData struct {
	InboxId string `json:"inbox_id"`
	EnabledActions *[]any `json:"enabled_actions,omitempty"`
	Id *string `json:"id,omitempty"`
	Instructions *any `json:"instructions,omitempty"`
	Object *string `json:"object,omitempty"`
	Tone *any `json:"tone,omitempty"`
}

// InboxDraft is the typed data model for the inbox_draft entity.
type InboxDraft struct {
}

// InboxDraftLoadMatch is the typed request payload for InboxDraft.LoadTyped.
type InboxDraftLoadMatch struct {
	Id string `json:"id"`
	InboxId string `json:"inbox_id"`
}

// InboxDraftListMatch is the typed request payload for InboxDraft.ListTyped.
type InboxDraftListMatch struct {
	InboxId string `json:"inbox_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// InboxDraftCreateData is the typed request payload for InboxDraft.CreateTyped.
type InboxDraftCreateData struct {
	InboxId string `json:"inbox_id"`
	Attachments []any `json:"attachments"`
	Bcc []any `json:"bcc"`
	Cc []any `json:"cc"`
	CreatedAt string `json:"created_at"`
	EmailId any `json:"email_id"`
	Html any `json:"html"`
	Id string `json:"id"`
	Object string `json:"object"`
	ReplyToEmailId any `json:"reply_to_email_id"`
	Snippet any `json:"snippet"`
	Subject any `json:"subject"`
	Text any `json:"text"`
	ThreadId any `json:"thread_id"`
	To any `json:"to"`
	Type string `json:"type"`
	UpdatedAt string `json:"updated_at"`
}

// InboxDraftUpdateData is the typed request payload for InboxDraft.UpdateTyped.
type InboxDraftUpdateData struct {
	Id string `json:"id"`
	InboxId string `json:"inbox_id"`
	Attachments *[]any `json:"attachments,omitempty"`
	Bcc *[]any `json:"bcc,omitempty"`
	Cc *[]any `json:"cc,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	EmailId *any `json:"email_id,omitempty"`
	Html *any `json:"html,omitempty"`
	Object *string `json:"object,omitempty"`
	ReplyToEmailId *any `json:"reply_to_email_id,omitempty"`
	Snippet *any `json:"snippet,omitempty"`
	Subject *any `json:"subject,omitempty"`
	Text *any `json:"text,omitempty"`
	ThreadId *any `json:"thread_id,omitempty"`
	To *any `json:"to,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// InboxDraftRemoveMatch is the typed request payload for InboxDraft.RemoveTyped.
type InboxDraftRemoveMatch struct {
	Id string `json:"id"`
	InboxId string `json:"inbox_id"`
}

// InboxLabel is the typed data model for the inbox_label entity.
type InboxLabel struct {
}

// InboxLabelListMatch is the typed request payload for InboxLabel.ListTyped.
type InboxLabelListMatch struct {
	Id string `json:"id"`
}

// InboxLabelCreateData is the typed request payload for InboxLabel.CreateTyped.
type InboxLabelCreateData struct {
	Id string `json:"id"`
	Color *string `json:"color,omitempty"`
	CreatedAt string `json:"created_at"`
	Name string `json:"name"`
	Object string `json:"object"`
}

// InboxLabelUpdateData is the typed request payload for InboxLabel.UpdateTyped.
type InboxLabelUpdateData struct {
	Id string `json:"id"`
	InboxId string `json:"inbox_id"`
	Color *string `json:"color,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
}

// InboxLabelRemoveMatch is the typed request payload for InboxLabel.RemoveTyped.
type InboxLabelRemoveMatch struct {
	Id string `json:"id"`
	InboxId string `json:"inbox_id"`
}

// InboxThread is the typed data model for the inbox_thread entity.
type InboxThread struct {
}

// InboxThreadLoadMatch is the typed request payload for InboxThread.LoadTyped.
type InboxThreadLoadMatch struct {
	Id string `json:"id"`
	InboxId string `json:"inbox_id"`
}

// InboxThreadListMatch is the typed request payload for InboxThread.ListTyped.
type InboxThreadListMatch struct {
	Id string `json:"id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Folder *[]any `json:"folder,omitempty"`
	Label *[]any `json:"label,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Read *bool `json:"read,omitempty"`
}

// InboxThreadUpdateData is the typed request payload for InboxThread.UpdateTyped.
type InboxThreadUpdateData struct {
	Id string `json:"id"`
	InboxId string `json:"inbox_id"`
	AddLabelIds *[]any `json:"add_label_ids,omitempty"`
	Bcc *[]any `json:"bcc,omitempty"`
	Cc *[]any `json:"cc,omitempty"`
	EmailCount *int `json:"email_count,omitempty"`
	Folder *string `json:"folder,omitempty"`
	From *any `json:"from,omitempty"`
	HasAttachment *bool `json:"has_attachment,omitempty"`
	HasDraft *bool `json:"has_draft,omitempty"`
	Highlights *map[string]any `json:"highlights,omitempty"`
	Labels *[]any `json:"labels,omitempty"`
	MatchedEmailId *any `json:"matched_email_id,omitempty"`
	Object *string `json:"object,omitempty"`
	Read *bool `json:"read,omitempty"`
	ReceivedAt *string `json:"received_at,omitempty"`
	RemoveLabelIds *[]any `json:"remove_label_ids,omitempty"`
	Subject *any `json:"subject,omitempty"`
	To *[]any `json:"to,omitempty"`
}

// InboxThreadRemoveMatch is the typed request payload for InboxThread.RemoveTyped.
type InboxThreadRemoveMatch struct {
	Id string `json:"id"`
	InboxId string `json:"inbox_id"`
}

// InboxThreadEmail is the typed data model for the inbox_thread_email entity.
type InboxThreadEmail struct {
}

// InboxThreadEmailLoadMatch is the typed request payload for InboxThreadEmail.LoadTyped.
type InboxThreadEmailLoadMatch struct {
	Id string `json:"id"`
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
}

// InboxThreadEmailListMatch is the typed request payload for InboxThreadEmail.ListTyped.
type InboxThreadEmailListMatch struct {
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// InboxThreadEmailCreateData is the typed request payload for InboxThreadEmail.CreateTyped.
type InboxThreadEmailCreateData struct {
	Id string `json:"id"`
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
	Attachments []any `json:"attachments"`
	Bcc []any `json:"bcc"`
	Cc []any `json:"cc"`
	Direction string `json:"direction"`
	From string `json:"from"`
	Html any `json:"html"`
	MessageId any `json:"message_id"`
	Read bool `json:"read"`
	ReceivedAt string `json:"received_at"`
	ReplyTo []any `json:"reply_to"`
	Subject any `json:"subject"`
	Text any `json:"text"`
	To []any `json:"to"`
}

// Log is the typed data model for the log entity.
type Log struct {
}

// LogLoadMatch is the typed request payload for Log.LoadTyped.
type LogLoadMatch struct {
	Id string `json:"id"`
}

// LogListMatch is the typed request payload for Log.ListTyped.
type LogListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// OauthGrant is the typed data model for the oauth_grant entity.
type OauthGrant struct {
}

// OauthGrantListMatch is the typed request payload for OauthGrant.ListTyped.
type OauthGrantListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// OauthGrantRemoveMatch is the typed request payload for OauthGrant.RemoveTyped.
type OauthGrantRemoveMatch struct {
	Id string `json:"id"`
}

// ReceivedEmail is the typed data model for the received_email entity.
type ReceivedEmail struct {
}

// ReceivedEmailLoadMatch is the typed request payload for ReceivedEmail.LoadTyped.
type ReceivedEmailLoadMatch struct {
	EmailId string `json:"email_id"`
}

// ReceivedEmailListMatch is the typed request payload for ReceivedEmail.ListTyped.
type ReceivedEmailListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ReceivedEmailAttachment is the typed data model for the received_email_attachment entity.
type ReceivedEmailAttachment struct {
}

// ReceivedEmailAttachmentLoadMatch is the typed request payload for ReceivedEmailAttachment.LoadTyped.
type ReceivedEmailAttachmentLoadMatch struct {
	EmailId string `json:"email_id"`
	Id string `json:"id"`
}

// ReceivedEmailAttachmentListMatch is the typed request payload for ReceivedEmailAttachment.ListTyped.
type ReceivedEmailAttachmentListMatch struct {
	EmailId string `json:"email_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// Segment is the typed data model for the segment entity.
type Segment struct {
}

// SegmentLoadMatch is the typed request payload for Segment.LoadTyped.
type SegmentLoadMatch struct {
	Id string `json:"id"`
}

// SegmentListMatch is the typed request payload for Segment.ListTyped.
type SegmentListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// SegmentCreateData is the typed request payload for Segment.CreateTyped.
type SegmentCreateData struct {
	AudienceId *string `json:"audience_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Filter *map[string]any `json:"filter,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
}

// SegmentUpdateData is the typed request payload for Segment.UpdateTyped.
type SegmentUpdateData struct {
	Id string `json:"id"`
	AudienceId *string `json:"audience_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Filter *map[string]any `json:"filter,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
}

// SegmentRemoveMatch is the typed request payload for Segment.RemoveTyped.
type SegmentRemoveMatch struct {
	Id string `json:"id"`
}

// SegmentContact is the typed data model for the segment_contact entity.
type SegmentContact struct {
}

// SegmentContactListMatch is the typed request payload for SegmentContact.ListTyped.
type SegmentContactListMatch struct {
	SegmentId string `json:"segment_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// Suppression is the typed data model for the suppression entity.
type Suppression struct {
}

// SuppressionLoadMatch is the typed request payload for Suppression.LoadTyped.
type SuppressionLoadMatch struct {
	Id string `json:"id"`
}

// SuppressionListMatch is the typed request payload for Suppression.ListTyped.
type SuppressionListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Origin *string `json:"origin,omitempty"`
}

// SuppressionCreateData is the typed request payload for Suppression.CreateTyped.
type SuppressionCreateData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	Origin *string `json:"origin,omitempty"`
	SourceId *any `json:"source_id,omitempty"`
}

// SuppressionRemoveMatch is the typed request payload for Suppression.RemoveTyped.
type SuppressionRemoveMatch struct {
	Id string `json:"id"`
}

// SuppressionBatch is the typed data model for the suppression_batch entity.
type SuppressionBatch struct {
}

// SuppressionBatchCreateData is the typed request payload for SuppressionBatch.CreateTyped.
type SuppressionBatchCreateData struct {
}

// Template is the typed data model for the template entity.
type Template struct {
}

// TemplateLoadMatch is the typed request payload for Template.LoadTyped.
type TemplateLoadMatch struct {
	Id string `json:"id"`
}

// TemplateListMatch is the typed request payload for Template.ListTyped.
type TemplateListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// TemplateCreateData is the typed request payload for Template.CreateTyped.
type TemplateCreateData struct {
	Alias *string `json:"alias,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CurrentVersionId *string `json:"current_version_id,omitempty"`
	From *string `json:"from,omitempty"`
	HasUnpublishedVersions *bool `json:"has_unpublished_versions,omitempty"`
	Html *string `json:"html,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	PublishedAt *any `json:"published_at,omitempty"`
	ReplyTo *any `json:"reply_to,omitempty"`
	Status *string `json:"status,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Text *string `json:"text,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Variables *[]any `json:"variables,omitempty"`
}

// TemplateUpdateData is the typed request payload for Template.UpdateTyped.
type TemplateUpdateData struct {
	Id string `json:"id"`
	Alias *string `json:"alias,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CurrentVersionId *string `json:"current_version_id,omitempty"`
	From *string `json:"from,omitempty"`
	HasUnpublishedVersions *bool `json:"has_unpublished_versions,omitempty"`
	Html *string `json:"html,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	PublishedAt *any `json:"published_at,omitempty"`
	ReplyTo *any `json:"reply_to,omitempty"`
	Status *string `json:"status,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Text *string `json:"text,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Variables *[]any `json:"variables,omitempty"`
}

// TemplateRemoveMatch is the typed request payload for Template.RemoveTyped.
type TemplateRemoveMatch struct {
	Id string `json:"id"`
}

// Topic is the typed data model for the topic entity.
type Topic struct {
}

// TopicLoadMatch is the typed request payload for Topic.LoadTyped.
type TopicLoadMatch struct {
	Id string `json:"id"`
}

// TopicListMatch is the typed request payload for Topic.ListTyped.
type TopicListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// TopicCreateData is the typed request payload for Topic.CreateTyped.
type TopicCreateData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	DefaultSubscription *string `json:"default_subscription,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

// TopicUpdateData is the typed request payload for Topic.UpdateTyped.
type TopicUpdateData struct {
	Id string `json:"id"`
	CreatedAt *string `json:"created_at,omitempty"`
	DefaultSubscription *string `json:"default_subscription,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

// TopicRemoveMatch is the typed request payload for Topic.RemoveTyped.
type TopicRemoveMatch struct {
	Id string `json:"id"`
}

// Usage is the typed data model for the usage entity.
type Usage struct {
}

// UsageLoadMatch is the typed request payload for Usage.LoadTyped.
type UsageLoadMatch struct {
	AiCredits *map[string]any `json:"ai_credits,omitempty"`
	AutomationRuns *map[string]any `json:"automation_runs,omitempty"`
	Broadcasts *map[string]any `json:"broadcasts,omitempty"`
	Contacts *map[string]any `json:"contacts,omitempty"`
	Domains *map[string]any `json:"domains,omitempty"`
	Emails *map[string]any `json:"emails,omitempty"`
	Object *string `json:"object,omitempty"`
	RateLimit *map[string]any `json:"rate_limit,omitempty"`
	Segments *map[string]any `json:"segments,omitempty"`
}

// Webhook is the typed data model for the webhook entity.
type Webhook struct {
}

// WebhookLoadMatch is the typed request payload for Webhook.LoadTyped.
type WebhookLoadMatch struct {
	Id string `json:"id"`
}

// WebhookListMatch is the typed request payload for Webhook.ListTyped.
type WebhookListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// WebhookCreateData is the typed request payload for Webhook.CreateTyped.
type WebhookCreateData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Events *any `json:"events,omitempty"`
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	SigningSecret *string `json:"signing_secret,omitempty"`
	Status *string `json:"status,omitempty"`
}

// WebhookUpdateData is the typed request payload for Webhook.UpdateTyped.
type WebhookUpdateData struct {
	Id string `json:"id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Events *any `json:"events,omitempty"`
	Object *string `json:"object,omitempty"`
	SigningSecret *string `json:"signing_secret,omitempty"`
	Status *string `json:"status,omitempty"`
}

// WebhookRemoveMatch is the typed request payload for Webhook.RemoveTyped.
type WebhookRemoveMatch struct {
	Id string `json:"id"`
}

// WebhookEvent is the typed data model for the webhook_event entity.
type WebhookEvent struct {
}

// WebhookEventLoadMatch is the typed request payload for WebhookEvent.LoadTyped.
type WebhookEventLoadMatch struct {
	Id string `json:"id"`
	WebhookId string `json:"webhook_id"`
}

// WebhookEventListMatch is the typed request payload for WebhookEvent.ListTyped.
type WebhookEventListMatch struct {
	WebhookId string `json:"webhook_id"`
	After *string `json:"after,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// WebhookEventCreateData is the typed request payload for WebhookEvent.CreateTyped.
type WebhookEventCreateData struct {
	EventId string `json:"event_id"`
	WebhookId string `json:"webhook_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	NextAttemptAt *any `json:"next_attempt_at,omitempty"`
	Object *string `json:"object,omitempty"`
	Payload *map[string]any `json:"payload,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
}

// WebhookEventAttempt is the typed data model for the webhook_event_attempt entity.
type WebhookEventAttempt struct {
}

// WebhookEventAttemptListMatch is the typed request payload for WebhookEventAttempt.ListTyped.
type WebhookEventAttemptListMatch struct {
	EventId string `json:"event_id"`
	WebhookId string `json:"webhook_id"`
	After *string `json:"after,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
