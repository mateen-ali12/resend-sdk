export interface ApiKey {
    created_at?: string;
    domain_id?: string;
    id?: string;
    last_used_at?: string | null;
    name: string;
    object?: string;
    permission?: string;
}
export interface ApiKeyListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface ApiKeyCreateData {
    created_at?: string;
    domain_id?: string;
    id?: string;
    last_used_at?: string | null;
    name: string;
    object?: string;
    permission?: string;
}
export interface ApiKeyUpdateData {
    id: string;
    created_at?: string;
    domain_id?: string;
    last_used_at?: string | null;
    name?: string;
    object?: string;
    permission?: string;
}
export interface ApiKeyRemoveMatch {
    id: string;
}
export interface Audience {
    created_at?: string;
    id?: string;
    name?: string;
    object?: string;
}
export interface AudienceLoadMatch {
    id: string;
}
export interface AudienceListMatch {
    created_at?: string;
    id?: string;
    name?: string;
    object?: string;
}
export interface AudienceCreateData {
    created_at?: string;
    id?: string;
    name?: string;
    object?: string;
}
export interface AudienceRemoveMatch {
    id: string;
}
export interface Automation {
    connections?: any[];
    created_at?: string;
    id?: string;
    name?: string;
    object?: string;
    status?: string;
    steps?: any[];
    updated_at?: string;
}
export interface AutomationLoadMatch {
    id: string;
}
export interface AutomationListMatch {
    after?: string;
    before?: string;
    limit?: number;
    status?: string;
}
export interface AutomationCreateData {
    connections?: any[];
    created_at?: string;
    id?: string;
    name?: string;
    object?: string;
    status?: string;
    steps?: any[];
    updated_at?: string;
    $action?: string;
    [action: string]: any;
}
export interface AutomationUpdateData {
    id: string;
    connections?: any[];
    created_at?: string;
    name?: string;
    object?: string;
    status?: string;
    steps?: any[];
    updated_at?: string;
}
export interface AutomationRemoveMatch {
    id: string;
}
export interface AutomationRun {
    completed_at?: string | null;
    created_at?: string;
    id?: string;
    object?: string;
    started_at?: string | null;
    status?: string;
    steps?: any[];
}
export interface AutomationRunLoadMatch {
    automation_id: string;
    id: string;
}
export interface AutomationRunListMatch {
    automation_id: string;
    after?: string;
    before?: string;
    limit?: number;
    status?: string;
}
export interface Broadcast {
    audience_id?: string | null;
    created_at?: string;
    from?: string;
    html?: string | null;
    id?: string;
    name?: string;
    object?: string;
    preview_text?: string;
    reply_to?: any[];
    scheduled_at?: string;
    segment_id?: string | null;
    send?: boolean;
    sent_at?: string;
    status?: string;
    subject?: string;
    text?: string | null;
    topic_id?: string | null;
}
export interface BroadcastLoadMatch {
    id: string;
}
export interface BroadcastListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface BroadcastCreateData {
    audience_id?: string | null;
    created_at?: string;
    from?: string;
    html?: string | null;
    id?: string;
    name?: string;
    object?: string;
    preview_text?: string;
    reply_to?: any[];
    scheduled_at?: string;
    segment_id?: string | null;
    send?: boolean;
    sent_at?: string;
    status?: string;
    subject?: string;
    text?: string | null;
    topic_id?: string | null;
    $action?: string;
    [action: string]: any;
}
export interface BroadcastUpdateData {
    id: string;
    audience_id?: string | null;
    created_at?: string;
    from?: string;
    html?: string | null;
    name?: string;
    object?: string;
    preview_text?: string;
    reply_to?: any[];
    scheduled_at?: string;
    segment_id?: string | null;
    send?: boolean;
    sent_at?: string;
    status?: string;
    subject?: string;
    text?: string | null;
    topic_id?: string | null;
}
export interface BroadcastRemoveMatch {
    id: string;
}
export interface BroadcastClickedLink {
    clicks?: number;
    id?: string;
    unique_clicks?: number;
    url?: string;
}
export interface BroadcastClickedLinkListMatch {
    broadcast_id: string;
    after?: string;
    before?: string;
    limit?: number;
}
export interface BroadcastRecipient {
    bounce_type?: string;
    clicked_links?: any[];
    contact_id?: string | null;
    count?: number;
    email?: string;
    id?: string;
}
export interface BroadcastRecipientListMatch {
    broadcast_id: string;
    after?: string;
    before?: string;
    bounce_type?: string;
    email?: string;
    limit?: number;
    type: string;
}
export interface Contact {
    audience_id?: string;
    created_at?: string;
    email?: string;
    first_name?: string | null;
    id?: string;
    last_name?: string | null;
    object?: string;
    properties?: Record<string, any>;
    segments?: any[];
    topics?: any[];
    unsubscribed?: boolean;
}
export interface ContactLoadMatch {
    id: string;
}
export interface ContactListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface ContactCreateData {
    audience_id?: string;
    created_at?: string;
    email?: string;
    first_name?: string | null;
    id?: string;
    last_name?: string | null;
    object?: string;
    properties?: Record<string, any>;
    segments?: any[];
    topics?: any[];
    unsubscribed?: boolean;
}
export interface ContactUpdateData {
    id: string;
    audience_id?: string;
    created_at?: string;
    email?: string;
    first_name?: string | null;
    last_name?: string | null;
    object?: string;
    properties?: Record<string, any>;
    segments?: any[];
    topics?: any[];
    unsubscribed?: boolean;
}
export interface ContactRemoveMatch {
    id: string;
}
export interface ContactImport {
    completed_at?: string | null;
    counts?: Record<string, any>;
    created_at?: string;
    id?: string;
    object?: string;
    status?: string;
}
export interface ContactImportLoadMatch {
    id: string;
}
export interface ContactImportListMatch {
    after?: string;
    before?: string;
    limit?: number;
    status?: string;
}
export interface ContactImportCreateData {
    completed_at?: string | null;
    counts?: Record<string, any>;
    created_at?: string;
    id?: string;
    object?: string;
    status?: string;
}
export interface ContactProperty {
    created_at?: string;
    fallback_value?: any;
    id?: string;
    key?: string;
    object?: string;
    type?: string;
}
export interface ContactPropertyLoadMatch {
    id: string;
}
export interface ContactPropertyListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface ContactPropertyCreateData {
    created_at?: string;
    fallback_value?: any;
    id?: string;
    key?: string;
    object?: string;
    type?: string;
}
export interface ContactPropertyUpdateData {
    id: string;
    created_at?: string;
    fallback_value?: any;
    key?: string;
    object?: string;
    type?: string;
}
export interface ContactPropertyRemoveMatch {
    id: string;
}
export interface ContactSegment {
    contact_id?: string;
    created_at?: string;
    id?: string;
    name?: string;
    object?: string;
    segment_id?: string;
}
export interface ContactSegmentListMatch {
    contact_id: string;
    after?: string;
    before?: string;
    limit?: number;
}
export interface ContactSegmentCreateData {
    contact_id: string;
    id: string;
    created_at?: string;
    name?: string;
    object?: string;
    segment_id?: string;
}
export interface ContactSegmentRemoveMatch {
    contact_id: string;
    id: string;
}
export interface ContactTopic {
    contact_id?: string;
    description?: string;
    id?: string;
    name?: string;
    object?: string;
    subscription?: string;
    topics?: any[];
}
export interface ContactTopicListMatch {
    contact_id: string;
    after?: string;
    before?: string;
    limit?: number;
}
export interface ContactTopicUpdateData {
    contact_id: string;
    description?: string;
    id?: string;
    name?: string;
    object?: string;
    subscription?: string;
    topics?: any[];
}
export interface Domain {
    capabilities?: Record<string, any>;
    click_tracking?: boolean;
    created_at?: string;
    custom_return_path?: string;
    id?: string;
    name?: string;
    object?: string;
    open_tracking?: boolean;
    records?: any[];
    region?: string;
    status?: string;
    tls?: string;
    tracking_subdomain?: string;
}
export interface DomainLoadMatch {
    id: string;
}
export interface DomainListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface DomainCreateData {
    capabilities?: Record<string, any>;
    click_tracking?: boolean;
    created_at?: string;
    custom_return_path?: string;
    id?: string;
    name?: string;
    object?: string;
    open_tracking?: boolean;
    records?: any[];
    region?: string;
    status?: string;
    tls?: string;
    tracking_subdomain?: string;
    $action?: string;
    [action: string]: any;
}
export interface DomainUpdateData {
    id: string;
    capabilities?: Record<string, any>;
    click_tracking?: boolean;
    created_at?: string;
    custom_return_path?: string;
    name?: string;
    object?: string;
    open_tracking?: boolean;
    records?: any[];
    region?: string;
    status?: string;
    tls?: string;
    tracking_subdomain?: string;
}
export interface DomainRemoveMatch {
    id: string;
}
export interface DomainClaim {
    blocked_reason?: string | null;
    click_tracking?: boolean;
    created_at?: string;
    custom_return_path?: string;
    domain_id?: string | null;
    expires_at?: string;
    failure_reason?: string | null;
    id?: string;
    name?: string;
    object?: string;
    open_tracking?: boolean;
    record?: Record<string, any>;
    region?: string | null;
    status?: string;
    tracking_subdomain?: string;
}
export interface DomainClaimLoadMatch {
    id: string;
}
export interface DomainClaimCreateData {
    blocked_reason?: string | null;
    click_tracking?: boolean;
    created_at?: string;
    custom_return_path?: string;
    domain_id?: string | null;
    expires_at?: string;
    failure_reason?: string | null;
    id?: string;
    name?: string;
    object?: string;
    open_tracking?: boolean;
    record?: Record<string, any>;
    region?: string | null;
    status?: string;
    tracking_subdomain?: string;
}
export interface Email {
    attachments?: any[];
    bcc?: any[];
    cc?: any[];
    created_at?: string;
    from?: string;
    headers?: Record<string, any>;
    html?: string;
    id?: string;
    last_event?: string;
    message_id?: string;
    object?: string;
    reply_to?: any[];
    scheduled_at?: string;
    subject?: string;
    tags?: any[];
    template?: any;
    text?: string;
    to?: any[];
    topic_id?: string;
}
export interface EmailLoadMatch {
    id: string;
}
export interface EmailListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface EmailCreateData {
    attachments?: any[];
    bcc?: any[];
    cc?: any[];
    created_at?: string;
    from?: string;
    headers?: Record<string, any>;
    html?: string;
    id?: string;
    last_event?: string;
    message_id?: string;
    object?: string;
    reply_to?: any[];
    scheduled_at?: string;
    subject?: string;
    tags?: any[];
    template?: any;
    text?: string;
    to?: any[];
    topic_id?: string;
    $action?: string;
    [action: string]: any;
}
export interface EmailUpdateData {
    id: string;
    attachments?: any[];
    bcc?: any[];
    cc?: any[];
    created_at?: string;
    from?: string;
    headers?: Record<string, any>;
    html?: string;
    last_event?: string;
    message_id?: string;
    object?: string;
    reply_to?: any[];
    scheduled_at?: string;
    subject?: string;
    tags?: any[];
    template?: any;
    text?: string;
    to?: any[];
    topic_id?: string;
}
export interface EmailAttachment {
    content_disposition?: string | null;
    content_id?: string;
    content_type?: string;
    download_url?: string;
    expires_at?: string;
    filename?: string | null;
    id?: string;
    object?: string;
    size?: number;
}
export interface EmailAttachmentLoadMatch {
    email_id: string;
    id: string;
}
export interface EmailAttachmentListMatch {
    email_id: string;
    after?: string;
    before?: string;
    limit?: number;
}
export interface EmailBatch {
    data?: any[];
    send_emails?: any[];
}
export interface EmailBatchCreateData {
    data?: any[];
    send_emails?: any[];
}
export interface EmailsMetric {
    broadcast_id?: string;
    broadcast_name?: string;
    domain_id?: string;
    domain_name?: string;
    email_id?: string;
    period?: string;
}
export interface EmailsMetricListMatch {
    broadcast_id?: any[];
    dimension?: any[];
    domain_id?: any[];
    email_id?: any[];
    end_date?: string;
    granularity?: string;
    metric?: any[];
    start_date?: string;
    timezone?: string;
}
export interface Event {
    created_at?: string;
    id?: string;
    name?: string;
    object?: string;
    schema?: Record<string, any> | null;
    updated_at?: string | null;
}
export interface EventLoadMatch {
    id: string;
}
export interface EventListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface EventCreateData {
    created_at?: string;
    id?: string;
    name?: string;
    object?: string;
    schema?: Record<string, any> | null;
    updated_at?: string | null;
    $action?: string;
    [action: string]: any;
}
export interface EventUpdateData {
    id: string;
    created_at?: string;
    name?: string;
    object?: string;
    schema?: Record<string, any> | null;
    updated_at?: string | null;
}
export interface EventRemoveMatch {
    id: string;
}
export interface Inbox {
    created_at: string;
    domain_id: string;
    drafts: number;
    email_address: string;
    forwarding?: boolean;
    from_name: string | null;
    id: string;
    last_received: string | null;
    name: string;
    object: string;
    receiving_address: string | null;
    unread: number;
}
export interface InboxLoadMatch {
    id: string;
}
export interface InboxListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface InboxCreateData {
    created_at: string;
    domain_id: string;
    drafts: number;
    email_address: string;
    forwarding?: boolean;
    from_name: string | null;
    id: string;
    last_received: string | null;
    name: string;
    object: string;
    receiving_address: string | null;
    unread: number;
}
export interface InboxUpdateData {
    id: string;
    created_at?: string;
    domain_id?: string;
    drafts?: number;
    email_address?: string;
    forwarding?: boolean;
    from_name?: string | null;
    last_received?: string | null;
    name?: string;
    object?: string;
    receiving_address?: string | null;
    unread?: number;
}
export interface InboxRemoveMatch {
    id: string;
}
export interface InboxAgent {
    enabled_actions: any[];
    id: string;
    instructions: string | null;
    object: string;
    tone: string | null;
}
export interface InboxAgentLoadMatch {
    inbox_id: string;
}
export interface InboxAgentUpdateData {
    inbox_id: string;
    enabled_actions?: any[];
    id?: string;
    instructions?: string | null;
    object?: string;
    tone?: string | null;
}
export interface InboxDraft {
    attachments: any[];
    bcc: any[];
    cc: any[];
    created_at: string;
    email_id: string | null;
    html: string | null;
    id: string;
    object: string;
    reply_to_email_id: string | null;
    snippet: string | null;
    subject: string | null;
    text: string | null;
    thread_id: string | null;
    to: any[] | null;
    type: string;
    updated_at: string;
}
export interface InboxDraftLoadMatch {
    id: string;
    inbox_id: string;
}
export interface InboxDraftListMatch {
    inbox_id: string;
    after?: string;
    before?: string;
    limit?: number;
}
export interface InboxDraftCreateData {
    inbox_id: string;
    attachments: any[];
    bcc: any[];
    cc: any[];
    created_at: string;
    email_id: string | null;
    html: string | null;
    id: string;
    object: string;
    reply_to_email_id: string | null;
    snippet: string | null;
    subject: string | null;
    text: string | null;
    thread_id: string | null;
    to: any[] | null;
    type: string;
    updated_at: string;
    $action?: string;
    [action: string]: any;
}
export interface InboxDraftUpdateData {
    id: string;
    inbox_id: string;
    attachments?: any[];
    bcc?: any[];
    cc?: any[];
    created_at?: string;
    email_id?: string | null;
    html?: string | null;
    object?: string;
    reply_to_email_id?: string | null;
    snippet?: string | null;
    subject?: string | null;
    text?: string | null;
    thread_id?: string | null;
    to?: any[] | null;
    type?: string;
    updated_at?: string;
}
export interface InboxDraftRemoveMatch {
    id: string;
    inbox_id: string;
}
export interface InboxLabel {
    color?: string;
    created_at: string;
    id: string;
    name: string;
    object: string;
}
export interface InboxLabelListMatch {
    id: string;
}
export interface InboxLabelCreateData {
    id: string;
    color?: string;
    created_at: string;
    name: string;
    object: string;
}
export interface InboxLabelUpdateData {
    id: string;
    inbox_id: string;
    color?: string;
    created_at?: string;
    name?: string;
    object?: string;
}
export interface InboxLabelRemoveMatch {
    id: string;
    inbox_id: string;
}
export interface InboxThread {
    add_label_ids?: any[];
    bcc: any[];
    cc: any[];
    email_count: number;
    folder: string;
    from: string | null;
    has_attachment: boolean;
    has_draft: boolean;
    highlights: Record<string, any>;
    id: string;
    labels: any[];
    matched_email_id: string | null;
    object: string;
    read: boolean;
    received_at: string;
    remove_label_ids?: any[];
    subject: string | null;
    to: any[];
}
export interface InboxThreadLoadMatch {
    id: string;
    inbox_id: string;
}
export interface InboxThreadListMatch {
    id: string;
    after?: string;
    before?: string;
    folder?: any[];
    label?: any[];
    limit?: number;
    read?: boolean;
}
export interface InboxThreadUpdateData {
    id: string;
    inbox_id: string;
    add_label_ids?: any[];
    bcc?: any[];
    cc?: any[];
    email_count?: number;
    folder?: string;
    from?: string | null;
    has_attachment?: boolean;
    has_draft?: boolean;
    highlights?: Record<string, any>;
    labels?: any[];
    matched_email_id?: string | null;
    object?: string;
    read?: boolean;
    received_at?: string;
    remove_label_ids?: any[];
    subject?: string | null;
    to?: any[];
}
export interface InboxThreadRemoveMatch {
    id: string;
    inbox_id: string;
}
export interface InboxThreadEmail {
    attachments: any[];
    bcc: any[];
    cc: any[];
    direction: string;
    from: string;
    html: string | null;
    id: string;
    message_id: string | null;
    read: boolean;
    received_at: string;
    reply_to: any[];
    subject: string | null;
    text: string | null;
    to: any[];
}
export interface InboxThreadEmailLoadMatch {
    id: string;
    inbox_id: string;
    thread_id: string;
}
export interface InboxThreadEmailListMatch {
    inbox_id: string;
    thread_id: string;
    after?: string;
    before?: string;
    limit?: number;
}
export interface InboxThreadEmailCreateData {
    id: string;
    inbox_id: string;
    thread_id: string;
    attachments: any[];
    bcc: any[];
    cc: any[];
    direction: string;
    from: string;
    html: string | null;
    message_id: string | null;
    read: boolean;
    received_at: string;
    reply_to: any[];
    subject: string | null;
    text: string | null;
    to: any[];
    $action?: string;
    [action: string]: any;
}
export interface Log {
    created_at?: string;
    endpoint?: string;
    id?: string;
    method?: string;
    object?: string;
    request_body?: Record<string, any> | null;
    response_body?: Record<string, any> | null;
    response_status?: number;
    user_agent?: string | null;
}
export interface LogLoadMatch {
    id: string;
}
export interface LogListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface OauthGrant {
    client?: Record<string, any>;
    client_id?: string;
    created_at?: string;
    id?: string;
    revoked_at?: string | null;
    revoked_reason?: string | null;
    scopes?: any[];
}
export interface OauthGrantListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface OauthGrantRemoveMatch {
    id: string;
}
export interface ReceivedEmail {
    attachments?: any[];
    bcc?: any[] | null;
    cc?: any[] | null;
    created_at?: string;
    from?: string;
    headers?: Record<string, any> | null;
    html?: string | null;
    id?: string;
    message_id?: string;
    object?: string;
    received_for?: any[];
    reply_to?: any[] | null;
    subject?: string;
    text?: string | null;
    to?: any[];
}
export interface ReceivedEmailLoadMatch {
    email_id: string;
}
export interface ReceivedEmailListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface ReceivedEmailAttachment {
    content_disposition?: string | null;
    content_id?: string;
    content_type?: string;
    download_url?: string;
    expires_at?: string;
    filename?: string | null;
    id?: string;
    object?: string;
    size?: number;
}
export interface ReceivedEmailAttachmentLoadMatch {
    email_id: string;
    id: string;
}
export interface ReceivedEmailAttachmentListMatch {
    email_id: string;
    after?: string;
    before?: string;
    limit?: number;
}
export interface Segment {
    audience_id?: string;
    created_at?: string;
    filter?: Record<string, any>;
    id?: string;
    name?: string;
    object?: string;
}
export interface SegmentLoadMatch {
    id: string;
}
export interface SegmentListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface SegmentCreateData {
    audience_id?: string;
    created_at?: string;
    filter?: Record<string, any>;
    id?: string;
    name?: string;
    object?: string;
}
export interface SegmentUpdateData {
    id: string;
    audience_id?: string;
    created_at?: string;
    filter?: Record<string, any>;
    name?: string;
    object?: string;
}
export interface SegmentRemoveMatch {
    id: string;
}
export interface SegmentContact {
    created_at?: string;
    email?: string;
    first_name?: string | null;
    id?: string;
    last_name?: string | null;
    unsubscribed?: boolean;
}
export interface SegmentContactListMatch {
    segment_id: string;
    after?: string;
    before?: string;
    limit?: number;
}
export interface Suppression {
    created_at?: string;
    email?: string;
    id?: string;
    object?: string;
    origin?: string;
    source_id?: string | null;
}
export interface SuppressionLoadMatch {
    id: string;
}
export interface SuppressionListMatch {
    after?: string;
    before?: string;
    limit?: number;
    origin?: string;
}
export interface SuppressionCreateData {
    created_at?: string;
    email?: string;
    id?: string;
    object?: string;
    origin?: string;
    source_id?: string | null;
}
export interface SuppressionRemoveMatch {
    id: string;
}
export interface SuppressionBatch {
}
export interface SuppressionBatchCreateData {
    $action?: string;
    [action: string]: any;
}
export interface Template {
    alias?: string;
    created_at?: string;
    current_version_id?: string;
    from?: string;
    has_unpublished_versions?: boolean;
    html?: string;
    id?: string;
    name?: string;
    object?: string;
    published_at?: string | null;
    reply_to?: any[] | null;
    status?: string;
    subject?: string;
    text?: string;
    updated_at?: string;
    variables?: any[];
}
export interface TemplateLoadMatch {
    id: string;
}
export interface TemplateListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface TemplateCreateData {
    alias?: string;
    created_at?: string;
    current_version_id?: string;
    from?: string;
    has_unpublished_versions?: boolean;
    html?: string;
    id?: string;
    name?: string;
    object?: string;
    published_at?: string | null;
    reply_to?: any[] | null;
    status?: string;
    subject?: string;
    text?: string;
    updated_at?: string;
    variables?: any[];
    $action?: string;
    [action: string]: any;
}
export interface TemplateUpdateData {
    id: string;
    alias?: string;
    created_at?: string;
    current_version_id?: string;
    from?: string;
    has_unpublished_versions?: boolean;
    html?: string;
    name?: string;
    object?: string;
    published_at?: string | null;
    reply_to?: any[] | null;
    status?: string;
    subject?: string;
    text?: string;
    updated_at?: string;
    variables?: any[];
}
export interface TemplateRemoveMatch {
    id: string;
}
export interface Topic {
    created_at?: string;
    default_subscription?: string;
    description?: string;
    id?: string;
    name?: string;
    object?: string;
    visibility?: string;
}
export interface TopicLoadMatch {
    id: string;
}
export interface TopicListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface TopicCreateData {
    created_at?: string;
    default_subscription?: string;
    description?: string;
    id?: string;
    name?: string;
    object?: string;
    visibility?: string;
}
export interface TopicUpdateData {
    id: string;
    created_at?: string;
    default_subscription?: string;
    description?: string;
    name?: string;
    object?: string;
    visibility?: string;
}
export interface TopicRemoveMatch {
    id: string;
}
export interface Usage {
    ai_credits?: Record<string, any>;
    automation_runs?: Record<string, any>;
    broadcasts?: Record<string, any>;
    contacts?: Record<string, any>;
    domains?: Record<string, any>;
    emails?: Record<string, any>;
    object?: string;
    rate_limit?: Record<string, any>;
    segments?: Record<string, any>;
}
export interface UsageLoadMatch {
    ai_credits?: Record<string, any>;
    automation_runs?: Record<string, any>;
    broadcasts?: Record<string, any>;
    contacts?: Record<string, any>;
    domains?: Record<string, any>;
    emails?: Record<string, any>;
    object?: string;
    rate_limit?: Record<string, any>;
    segments?: Record<string, any>;
}
export interface Webhook {
    created_at?: string;
    endpoint?: string;
    events?: any[] | null;
    id?: string;
    object?: string;
    signing_secret?: string;
    status?: string;
}
export interface WebhookLoadMatch {
    id: string;
}
export interface WebhookListMatch {
    after?: string;
    before?: string;
    limit?: number;
}
export interface WebhookCreateData {
    created_at?: string;
    endpoint?: string;
    events?: any[] | null;
    id?: string;
    object?: string;
    signing_secret?: string;
    status?: string;
    $action?: string;
    [action: string]: any;
}
export interface WebhookUpdateData {
    id: string;
    created_at?: string;
    endpoint?: string;
    events?: any[] | null;
    object?: string;
    signing_secret?: string;
    status?: string;
}
export interface WebhookRemoveMatch {
    id: string;
}
export interface WebhookEvent {
    created_at?: string;
    id?: string;
    next_attempt_at?: string | null;
    object?: string;
    payload?: Record<string, any>;
    status?: string;
    type?: string;
}
export interface WebhookEventLoadMatch {
    id: string;
    webhook_id: string;
}
export interface WebhookEventListMatch {
    webhook_id: string;
    after?: string;
    limit?: number;
}
export interface WebhookEventCreateData {
    event_id: string;
    webhook_id: string;
    created_at?: string;
    id?: string;
    next_attempt_at?: string | null;
    object?: string;
    payload?: Record<string, any>;
    status?: string;
    type?: string;
    $action?: string;
    [action: string]: any;
}
export interface WebhookEventAttempt {
    http_status_code?: number;
    id?: string;
    response?: string;
    sent_at?: string;
}
export interface WebhookEventAttemptListMatch {
    event_id: string;
    webhook_id: string;
    after?: string;
    limit?: number;
}
