# Typed models for the Resend SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class ApiKeyRequired(TypedDict):
    name: str


class ApiKey(ApiKeyRequired, total=False):
    created_at: str
    domain_id: str
    id: str
    last_used_at: str | None
    object: str
    permission: str


class ApiKeyListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class ApiKeyCreateDataRequired(TypedDict):
    name: str


class ApiKeyCreateData(ApiKeyCreateDataRequired, total=False):
    created_at: str
    domain_id: str
    id: str
    last_used_at: str | None
    object: str
    permission: str


class ApiKeyUpdateDataRequired(TypedDict):
    id: str


class ApiKeyUpdateData(ApiKeyUpdateDataRequired, total=False):
    created_at: str
    domain_id: str
    last_used_at: str | None
    name: str
    object: str
    permission: str


class ApiKeyRemoveMatch(TypedDict):
    id: str


class Audience(TypedDict, total=False):
    created_at: str
    id: str
    name: str
    object: str


class AudienceLoadMatch(TypedDict):
    id: str


class AudienceListMatch(TypedDict, total=False):
    created_at: str
    id: str
    name: str
    object: str


class AudienceCreateData(TypedDict, total=False):
    created_at: str
    id: str
    name: str
    object: str


class AudienceRemoveMatch(TypedDict):
    id: str


class Automation(TypedDict, total=False):
    connections: list
    created_at: str
    id: str
    name: str
    object: str
    status: str
    steps: list
    updated_at: str


class AutomationLoadMatch(TypedDict):
    id: str


class AutomationListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int
    status: str


class AutomationCreateData(TypedDict, total=False):
    connections: list
    created_at: str
    id: str
    name: str
    object: str
    status: str
    steps: list
    updated_at: str


class AutomationUpdateDataRequired(TypedDict):
    id: str


class AutomationUpdateData(AutomationUpdateDataRequired, total=False):
    connections: list
    created_at: str
    name: str
    object: str
    status: str
    steps: list
    updated_at: str


class AutomationRemoveMatch(TypedDict):
    id: str


class AutomationRun(TypedDict, total=False):
    completed_at: str | None
    created_at: str
    id: str
    object: str
    started_at: str | None
    status: str
    steps: list


class AutomationRunLoadMatch(TypedDict):
    automation_id: str
    id: str


class AutomationRunListMatchRequired(TypedDict):
    automation_id: str


class AutomationRunListMatch(AutomationRunListMatchRequired, total=False):
    after: str
    before: str
    limit: int
    status: str


class Broadcast(TypedDict, total=False):
    audience_id: str | None
    created_at: str
    html: str | None
    id: str
    name: str
    object: str
    preview_text: str
    reply_to: list
    scheduled_at: str
    segment_id: str | None
    send: bool
    sent_at: str
    status: str
    subject: str
    text: str | None
    topic_id: str | None


class BroadcastLoadMatch(TypedDict):
    id: str


class BroadcastListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class BroadcastCreateData(TypedDict, total=False):
    audience_id: str | None
    created_at: str
    html: str | None
    id: str
    name: str
    object: str
    preview_text: str
    reply_to: list
    scheduled_at: str
    segment_id: str | None
    send: bool
    sent_at: str
    status: str
    subject: str
    text: str | None
    topic_id: str | None


class BroadcastUpdateDataRequired(TypedDict):
    id: str


class BroadcastUpdateData(BroadcastUpdateDataRequired, total=False):
    audience_id: str | None
    created_at: str
    html: str | None
    name: str
    object: str
    preview_text: str
    reply_to: list
    scheduled_at: str
    segment_id: str | None
    send: bool
    sent_at: str
    status: str
    subject: str
    text: str | None
    topic_id: str | None


class BroadcastRemoveMatch(TypedDict):
    id: str


class BroadcastClickedLink(TypedDict, total=False):
    clicks: int
    id: str
    unique_clicks: int
    url: str


class BroadcastClickedLinkListMatchRequired(TypedDict):
    broadcast_id: str


class BroadcastClickedLinkListMatch(BroadcastClickedLinkListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class BroadcastRecipient(TypedDict, total=False):
    bounce_type: str
    clicked_links: list
    contact_id: str | None
    count: int
    email: str
    id: str


class BroadcastRecipientListMatchRequired(TypedDict):
    broadcast_id: str
    type: str


class BroadcastRecipientListMatch(BroadcastRecipientListMatchRequired, total=False):
    after: str
    before: str
    bounce_type: str
    email: str
    limit: int


class Contact(TypedDict, total=False):
    audience_id: str
    created_at: str
    email: str
    first_name: str | None
    id: str
    last_name: str | None
    object: str
    properties: dict
    segments: list
    topics: list
    unsubscribed: bool


class ContactLoadMatch(TypedDict):
    id: str


class ContactListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class ContactCreateData(TypedDict, total=False):
    audience_id: str
    created_at: str
    email: str
    first_name: str | None
    id: str
    last_name: str | None
    object: str
    properties: dict
    segments: list
    topics: list
    unsubscribed: bool


class ContactUpdateDataRequired(TypedDict):
    id: str


class ContactUpdateData(ContactUpdateDataRequired, total=False):
    audience_id: str
    created_at: str
    email: str
    first_name: str | None
    last_name: str | None
    object: str
    properties: dict
    segments: list
    topics: list
    unsubscribed: bool


class ContactRemoveMatch(TypedDict):
    id: str


class ContactImport(TypedDict, total=False):
    completed_at: str | None
    counts: dict
    created_at: str
    id: str
    object: str
    status: str


class ContactImportLoadMatch(TypedDict):
    id: str


class ContactImportListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int
    status: str


class ContactImportCreateData(TypedDict, total=False):
    completed_at: str | None
    counts: dict
    created_at: str
    id: str
    object: str
    status: str


class ContactProperty(TypedDict, total=False):
    created_at: str
    fallback_value: Any
    id: str
    key: str
    object: str
    type: str


class ContactPropertyLoadMatch(TypedDict):
    id: str


class ContactPropertyListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class ContactPropertyCreateData(TypedDict, total=False):
    created_at: str
    fallback_value: Any
    id: str
    key: str
    object: str
    type: str


class ContactPropertyUpdateDataRequired(TypedDict):
    id: str


class ContactPropertyUpdateData(ContactPropertyUpdateDataRequired, total=False):
    created_at: str
    fallback_value: Any
    key: str
    object: str
    type: str


class ContactPropertyRemoveMatch(TypedDict):
    id: str


class ContactSegment(TypedDict, total=False):
    contact_id: str
    created_at: str
    id: str
    name: str
    object: str
    segment_id: str


class ContactSegmentListMatchRequired(TypedDict):
    contact_id: str


class ContactSegmentListMatch(ContactSegmentListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class ContactSegmentCreateDataRequired(TypedDict):
    contact_id: str
    id: str


class ContactSegmentCreateData(ContactSegmentCreateDataRequired, total=False):
    created_at: str
    name: str
    object: str
    segment_id: str


class ContactSegmentRemoveMatch(TypedDict):
    contact_id: str
    id: str


class ContactTopic(TypedDict, total=False):
    contact_id: str
    description: str
    id: str
    name: str
    object: str
    subscription: str
    topics: list


class ContactTopicListMatchRequired(TypedDict):
    contact_id: str


class ContactTopicListMatch(ContactTopicListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class ContactTopicUpdateDataRequired(TypedDict):
    contact_id: str


class ContactTopicUpdateData(ContactTopicUpdateDataRequired, total=False):
    description: str
    id: str
    name: str
    object: str
    subscription: str
    topics: list


class Domain(TypedDict, total=False):
    capabilities: dict
    click_tracking: bool
    created_at: str
    custom_return_path: str
    id: str
    name: str
    object: str
    open_tracking: bool
    records: list
    region: str
    status: str
    tls: str
    tracking_subdomain: str


class DomainLoadMatch(TypedDict):
    id: str


class DomainListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class DomainCreateData(TypedDict, total=False):
    capabilities: dict
    click_tracking: bool
    created_at: str
    custom_return_path: str
    id: str
    name: str
    object: str
    open_tracking: bool
    records: list
    region: str
    status: str
    tls: str
    tracking_subdomain: str


class DomainUpdateDataRequired(TypedDict):
    id: str


class DomainUpdateData(DomainUpdateDataRequired, total=False):
    capabilities: dict
    click_tracking: bool
    created_at: str
    custom_return_path: str
    name: str
    object: str
    open_tracking: bool
    records: list
    region: str
    status: str
    tls: str
    tracking_subdomain: str


class DomainRemoveMatch(TypedDict):
    id: str


class DomainClaim(TypedDict, total=False):
    blocked_reason: str | None
    click_tracking: bool
    created_at: str
    custom_return_path: str
    domain_id: str | None
    expires_at: str
    failure_reason: str | None
    id: str
    name: str
    object: str
    open_tracking: bool
    record: dict
    region: str | None
    status: str
    tracking_subdomain: str


class DomainClaimLoadMatch(TypedDict):
    id: str


class DomainClaimCreateData(TypedDict, total=False):
    blocked_reason: str | None
    click_tracking: bool
    created_at: str
    custom_return_path: str
    domain_id: str | None
    expires_at: str
    failure_reason: str | None
    id: str
    name: str
    object: str
    open_tracking: bool
    record: dict
    region: str | None
    status: str
    tracking_subdomain: str


class Email(TypedDict, total=False):
    attachments: list
    bcc: list
    cc: list
    created_at: str
    headers: dict
    html: str
    id: str
    last_event: str
    message_id: str
    object: str
    reply_to: list
    scheduled_at: str
    subject: str
    tags: list
    template: Any
    text: str
    to: list
    topic_id: str


class EmailLoadMatch(TypedDict):
    id: str


class EmailListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class EmailCreateData(TypedDict, total=False):
    attachments: list
    bcc: list
    cc: list
    created_at: str
    headers: dict
    html: str
    id: str
    last_event: str
    message_id: str
    object: str
    reply_to: list
    scheduled_at: str
    subject: str
    tags: list
    template: Any
    text: str
    to: list
    topic_id: str


class EmailUpdateDataRequired(TypedDict):
    id: str


class EmailUpdateData(EmailUpdateDataRequired, total=False):
    attachments: list
    bcc: list
    cc: list
    created_at: str
    headers: dict
    html: str
    last_event: str
    message_id: str
    object: str
    reply_to: list
    scheduled_at: str
    subject: str
    tags: list
    template: Any
    text: str
    to: list
    topic_id: str


class EmailAttachment(TypedDict, total=False):
    content_disposition: str | None
    content_id: str
    content_type: str
    download_url: str
    expires_at: str
    filename: str | None
    id: str
    object: str
    size: int


class EmailAttachmentLoadMatch(TypedDict):
    email_id: str
    id: str


class EmailAttachmentListMatchRequired(TypedDict):
    email_id: str


class EmailAttachmentListMatch(EmailAttachmentListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class EmailBatch(TypedDict, total=False):
    data: list
    send_emails: list


class EmailBatchCreateData(TypedDict, total=False):
    data: list
    send_emails: list


class EmailsMetric(TypedDict, total=False):
    broadcast_id: str
    broadcast_name: str
    domain_id: str
    domain_name: str
    email_id: str
    period: str


class EmailsMetricListMatch(TypedDict, total=False):
    broadcast_id: list
    dimension: list
    domain_id: list
    email_id: list
    end_date: str
    granularity: str
    metric: list
    start_date: str
    timezone: str


class Event(TypedDict, total=False):
    created_at: str
    id: str
    name: str
    object: str
    schema: dict | None
    updated_at: str | None


class EventLoadMatch(TypedDict):
    id: str


class EventListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class EventCreateData(TypedDict, total=False):
    created_at: str
    id: str
    name: str
    object: str
    schema: dict | None
    updated_at: str | None


class EventUpdateDataRequired(TypedDict):
    id: str


class EventUpdateData(EventUpdateDataRequired, total=False):
    created_at: str
    name: str
    object: str
    schema: dict | None
    updated_at: str | None


class EventRemoveMatch(TypedDict):
    id: str


class InboxRequired(TypedDict):
    created_at: str
    domain_id: str
    drafts: int
    email_address: str
    from_name: str | None
    id: str
    last_received: str | None
    name: str
    object: str
    receiving_address: str | None
    unread: int


class Inbox(InboxRequired, total=False):
    forwarding: bool


class InboxLoadMatch(TypedDict):
    id: str


class InboxListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class InboxCreateDataRequired(TypedDict):
    created_at: str
    domain_id: str
    drafts: int
    email_address: str
    from_name: str | None
    id: str
    last_received: str | None
    name: str
    object: str
    receiving_address: str | None
    unread: int


class InboxCreateData(InboxCreateDataRequired, total=False):
    forwarding: bool


class InboxUpdateDataRequired(TypedDict):
    id: str


class InboxUpdateData(InboxUpdateDataRequired, total=False):
    created_at: str
    domain_id: str
    drafts: int
    email_address: str
    forwarding: bool
    from_name: str | None
    last_received: str | None
    name: str
    object: str
    receiving_address: str | None
    unread: int


class InboxRemoveMatch(TypedDict):
    id: str


class InboxAgent(TypedDict):
    enabled_actions: list
    id: str
    instructions: str | None
    object: str
    tone: str | None


class InboxAgentLoadMatch(TypedDict):
    inbox_id: str


class InboxAgentUpdateDataRequired(TypedDict):
    inbox_id: str


class InboxAgentUpdateData(InboxAgentUpdateDataRequired, total=False):
    enabled_actions: list
    id: str
    instructions: str | None
    object: str
    tone: str | None


class InboxDraft(TypedDict):
    attachments: list
    bcc: list
    cc: list
    created_at: str
    email_id: str | None
    html: str | None
    id: str
    object: str
    reply_to_email_id: str | None
    snippet: str | None
    subject: str | None
    text: str | None
    thread_id: str | None
    to: list | None
    type: str
    updated_at: str


class InboxDraftLoadMatch(TypedDict):
    id: str
    inbox_id: str


class InboxDraftListMatchRequired(TypedDict):
    inbox_id: str


class InboxDraftListMatch(InboxDraftListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class InboxDraftCreateData(TypedDict):
    inbox_id: str
    attachments: list
    bcc: list
    cc: list
    created_at: str
    email_id: str | None
    html: str | None
    id: str
    object: str
    reply_to_email_id: str | None
    snippet: str | None
    subject: str | None
    text: str | None
    thread_id: str | None
    to: list | None
    type: str
    updated_at: str


class InboxDraftUpdateDataRequired(TypedDict):
    id: str
    inbox_id: str


class InboxDraftUpdateData(InboxDraftUpdateDataRequired, total=False):
    attachments: list
    bcc: list
    cc: list
    created_at: str
    email_id: str | None
    html: str | None
    object: str
    reply_to_email_id: str | None
    snippet: str | None
    subject: str | None
    text: str | None
    thread_id: str | None
    to: list | None
    type: str
    updated_at: str


class InboxDraftRemoveMatch(TypedDict):
    id: str
    inbox_id: str


class InboxLabelRequired(TypedDict):
    created_at: str
    id: str
    name: str
    object: str


class InboxLabel(InboxLabelRequired, total=False):
    color: str


class InboxLabelListMatch(TypedDict):
    id: str


class InboxLabelCreateDataRequired(TypedDict):
    id: str
    created_at: str
    name: str
    object: str


class InboxLabelCreateData(InboxLabelCreateDataRequired, total=False):
    color: str


class InboxLabelUpdateDataRequired(TypedDict):
    id: str
    inbox_id: str


class InboxLabelUpdateData(InboxLabelUpdateDataRequired, total=False):
    color: str
    created_at: str
    name: str
    object: str


class InboxLabelRemoveMatch(TypedDict):
    id: str
    inbox_id: str


class InboxThreadRequired(TypedDict):
    bcc: list
    cc: list
    email_count: int
    folder: str
    has_attachment: bool
    has_draft: bool
    highlights: dict
    id: str
    labels: list
    matched_email_id: str | None
    object: str
    read: bool
    received_at: str
    subject: str | None
    to: list


class InboxThread(InboxThreadRequired, total=False):
    add_label_ids: list
    remove_label_ids: list


class InboxThreadLoadMatch(TypedDict):
    id: str
    inbox_id: str


class InboxThreadListMatchRequired(TypedDict):
    id: str


class InboxThreadListMatch(InboxThreadListMatchRequired, total=False):
    after: str
    before: str
    folder: list
    label: list
    limit: int
    read: bool


class InboxThreadUpdateDataRequired(TypedDict):
    id: str
    inbox_id: str


class InboxThreadUpdateData(InboxThreadUpdateDataRequired, total=False):
    add_label_ids: list
    bcc: list
    cc: list
    email_count: int
    folder: str
    has_attachment: bool
    has_draft: bool
    highlights: dict
    labels: list
    matched_email_id: str | None
    object: str
    read: bool
    received_at: str
    remove_label_ids: list
    subject: str | None
    to: list


class InboxThreadRemoveMatch(TypedDict):
    id: str
    inbox_id: str


class InboxThreadEmail(TypedDict):
    attachments: list
    bcc: list
    cc: list
    direction: str
    html: str | None
    id: str
    message_id: str | None
    read: bool
    received_at: str
    reply_to: list
    subject: str | None
    text: str | None
    to: list


class InboxThreadEmailLoadMatch(TypedDict):
    id: str
    inbox_id: str
    thread_id: str


class InboxThreadEmailListMatchRequired(TypedDict):
    inbox_id: str
    thread_id: str


class InboxThreadEmailListMatch(InboxThreadEmailListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class InboxThreadEmailCreateData(TypedDict):
    id: str
    inbox_id: str
    thread_id: str
    attachments: list
    bcc: list
    cc: list
    direction: str
    html: str | None
    message_id: str | None
    read: bool
    received_at: str
    reply_to: list
    subject: str | None
    text: str | None
    to: list


class Log(TypedDict, total=False):
    created_at: str
    endpoint: str
    id: str
    method: str
    object: str
    request_body: dict | None
    response_body: dict | None
    response_status: int
    user_agent: str | None


class LogLoadMatch(TypedDict):
    id: str


class LogListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class OauthGrant(TypedDict, total=False):
    client: dict
    client_id: str
    created_at: str
    id: str
    revoked_at: str | None
    revoked_reason: str | None
    scopes: list


class OauthGrantListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class OauthGrantRemoveMatch(TypedDict):
    id: str


class ReceivedEmail(TypedDict, total=False):
    attachments: list
    bcc: list | None
    cc: list | None
    created_at: str
    headers: dict | None
    html: str | None
    id: str
    message_id: str
    object: str
    received_for: list
    reply_to: list | None
    subject: str
    text: str | None
    to: list


class ReceivedEmailLoadMatch(TypedDict):
    email_id: str


class ReceivedEmailListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class ReceivedEmailAttachment(TypedDict, total=False):
    content_disposition: str | None
    content_id: str
    content_type: str
    download_url: str
    expires_at: str
    filename: str | None
    id: str
    object: str
    size: int


class ReceivedEmailAttachmentLoadMatch(TypedDict):
    email_id: str
    id: str


class ReceivedEmailAttachmentListMatchRequired(TypedDict):
    email_id: str


class ReceivedEmailAttachmentListMatch(ReceivedEmailAttachmentListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class Segment(TypedDict, total=False):
    audience_id: str
    created_at: str
    filter: dict
    id: str
    name: str
    object: str


class SegmentLoadMatch(TypedDict):
    id: str


class SegmentListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class SegmentCreateData(TypedDict, total=False):
    audience_id: str
    created_at: str
    filter: dict
    id: str
    name: str
    object: str


class SegmentUpdateDataRequired(TypedDict):
    id: str


class SegmentUpdateData(SegmentUpdateDataRequired, total=False):
    audience_id: str
    created_at: str
    filter: dict
    name: str
    object: str


class SegmentRemoveMatch(TypedDict):
    id: str


class SegmentContact(TypedDict, total=False):
    created_at: str
    email: str
    first_name: str | None
    id: str
    last_name: str | None
    unsubscribed: bool


class SegmentContactListMatchRequired(TypedDict):
    segment_id: str


class SegmentContactListMatch(SegmentContactListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class Suppression(TypedDict, total=False):
    created_at: str
    email: str
    id: str
    object: str
    origin: str
    source_id: str | None


class SuppressionLoadMatch(TypedDict):
    id: str


class SuppressionListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int
    origin: str


class SuppressionCreateData(TypedDict, total=False):
    created_at: str
    email: str
    id: str
    object: str
    origin: str
    source_id: str | None


class SuppressionRemoveMatch(TypedDict):
    id: str


class SuppressionBatch(TypedDict):
    pass


class SuppressionBatchCreateData(TypedDict):
    pass


class Template(TypedDict, total=False):
    alias: str
    created_at: str
    current_version_id: str
    has_unpublished_versions: bool
    html: str
    id: str
    name: str
    object: str
    published_at: str | None
    reply_to: list | None
    status: str
    subject: str
    text: str
    updated_at: str
    variables: list


class TemplateLoadMatch(TypedDict):
    id: str


class TemplateListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class TemplateCreateData(TypedDict, total=False):
    alias: str
    created_at: str
    current_version_id: str
    has_unpublished_versions: bool
    html: str
    id: str
    name: str
    object: str
    published_at: str | None
    reply_to: list | None
    status: str
    subject: str
    text: str
    updated_at: str
    variables: list


class TemplateUpdateDataRequired(TypedDict):
    id: str


class TemplateUpdateData(TemplateUpdateDataRequired, total=False):
    alias: str
    created_at: str
    current_version_id: str
    has_unpublished_versions: bool
    html: str
    name: str
    object: str
    published_at: str | None
    reply_to: list | None
    status: str
    subject: str
    text: str
    updated_at: str
    variables: list


class TemplateRemoveMatch(TypedDict):
    id: str


class Topic(TypedDict, total=False):
    created_at: str
    default_subscription: str
    description: str
    id: str
    name: str
    object: str
    visibility: str


class TopicLoadMatch(TypedDict):
    id: str


class TopicListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class TopicCreateData(TypedDict, total=False):
    created_at: str
    default_subscription: str
    description: str
    id: str
    name: str
    object: str
    visibility: str


class TopicUpdateDataRequired(TypedDict):
    id: str


class TopicUpdateData(TopicUpdateDataRequired, total=False):
    created_at: str
    default_subscription: str
    description: str
    name: str
    object: str
    visibility: str


class TopicRemoveMatch(TypedDict):
    id: str


class Usage(TypedDict, total=False):
    ai_credits: dict
    automation_runs: dict
    broadcasts: dict
    contacts: dict
    domains: dict
    emails: dict
    object: str
    rate_limit: dict
    segments: dict


class UsageLoadMatch(TypedDict, total=False):
    ai_credits: dict
    automation_runs: dict
    broadcasts: dict
    contacts: dict
    domains: dict
    emails: dict
    object: str
    rate_limit: dict
    segments: dict


class Webhook(TypedDict, total=False):
    created_at: str
    endpoint: str
    events: list | None
    id: str
    object: str
    signing_secret: str
    status: str


class WebhookLoadMatch(TypedDict):
    id: str


class WebhookListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class WebhookCreateData(TypedDict, total=False):
    created_at: str
    endpoint: str
    events: list | None
    id: str
    object: str
    signing_secret: str
    status: str


class WebhookUpdateDataRequired(TypedDict):
    id: str


class WebhookUpdateData(WebhookUpdateDataRequired, total=False):
    created_at: str
    endpoint: str
    events: list | None
    object: str
    signing_secret: str
    status: str


class WebhookRemoveMatch(TypedDict):
    id: str


class WebhookEvent(TypedDict, total=False):
    created_at: str
    id: str
    next_attempt_at: str | None
    object: str
    payload: dict
    status: str
    type: str


class WebhookEventLoadMatch(TypedDict):
    id: str
    webhook_id: str


class WebhookEventListMatchRequired(TypedDict):
    webhook_id: str


class WebhookEventListMatch(WebhookEventListMatchRequired, total=False):
    after: str
    limit: int


class WebhookEventCreateDataRequired(TypedDict):
    event_id: str
    webhook_id: str


class WebhookEventCreateData(WebhookEventCreateDataRequired, total=False):
    created_at: str
    id: str
    next_attempt_at: str | None
    object: str
    payload: dict
    status: str
    type: str


class WebhookEventAttempt(TypedDict, total=False):
    http_status_code: int
    id: str
    response: str
    sent_at: str


class WebhookEventAttemptListMatchRequired(TypedDict):
    event_id: str
    webhook_id: str


class WebhookEventAttemptListMatch(WebhookEventAttemptListMatchRequired, total=False):
    after: str
    limit: int
