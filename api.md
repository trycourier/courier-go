# Shared Params Types

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared#ChannelPreferenceParam">ChannelPreferenceParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared#PreferenceParam">PreferenceParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared#RuleParam">RuleParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared#ChannelPreference">ChannelPreference</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared#Preference">Preference</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go/shared#Rule">Rule</a>

# Send

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ContentUnionParam">ContentUnionParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageContextParam">MessageContextParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#RecipientParam">RecipientParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UtmParam">UtmParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ContentUnion">ContentUnion</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageContext">MessageContext</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#SendSendMessageResponse">SendSendMessageResponse</a>

Methods:

- <code title="post /send">client.Send.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#SendService.SendMessage">SendMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#SendSendMessageParams">SendSendMessageParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#SendSendMessageResponse">SendSendMessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Audiences

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#FilterParam">FilterParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Audience">Audience</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Filter">Filter</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Paging">Paging</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceUpdateResponse">AudienceUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceListResponse">AudienceListResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceListMembersResponse">AudienceListMembersResponse</a>

Methods:

- <code title="get /audiences/{audience_id}">client.Audiences.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, audienceID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Audience">Audience</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /audiences/{audience_id}">client.Audiences.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, audienceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceUpdateParams">AudienceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceUpdateResponse">AudienceUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /audiences">client.Audiences.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceListParams">AudienceListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceListResponse">AudienceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /audiences/{audience_id}">client.Audiences.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, audienceID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /audiences/{audience_id}/members">client.Audiences.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceService.ListMembers">ListMembers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, audienceID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceListMembersParams">AudienceListMembersParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AudienceListMembersResponse">AudienceListMembersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AuditEvents

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuditEvent">AuditEvent</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuditEventListResponse">AuditEventListResponse</a>

Methods:

- <code title="get /audit-events/{audit-event-id}">client.AuditEvents.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuditEventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, auditEventID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuditEvent">AuditEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /audit-events">client.AuditEvents.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuditEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuditEventListParams">AuditEventListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuditEventListResponse">AuditEventListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuthIssueTokenResponse">AuthIssueTokenResponse</a>

Methods:

- <code title="post /auth/issue-token">client.Auth.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuthService.IssueToken">IssueToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuthIssueTokenParams">AuthIssueTokenParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AuthIssueTokenResponse">AuthIssueTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Automations

## Invoke

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AutomationInvokeResponse">AutomationInvokeResponse</a>

Methods:

- <code title="post /automations/invoke">client.Automations.Invoke.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AutomationInvokeService.InvokeAdHoc">InvokeAdHoc</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AutomationInvokeInvokeAdHocParams">AutomationInvokeInvokeAdHocParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AutomationInvokeResponse">AutomationInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /automations/{templateId}/invoke">client.Automations.Invoke.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AutomationInvokeService.InvokeByTemplate">InvokeByTemplate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, templateID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AutomationInvokeInvokeByTemplateParams">AutomationInvokeInvokeByTemplateParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#AutomationInvokeResponse">AutomationInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Brands

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandColorsParam">BrandColorsParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandSettingsParam">BrandSettingsParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandSettingsEmailParam">BrandSettingsEmailParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandSettingsInAppParam">BrandSettingsInAppParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandSnippetParam">BrandSnippetParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandSnippetsParam">BrandSnippetsParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandTemplateParam">BrandTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#EmailFooterParam">EmailFooterParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#EmailHeadParam">EmailHeadParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#EmailHeaderParam">EmailHeaderParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#IconsParam">IconsParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#LogoParam">LogoParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#WidgetBackgroundParam">WidgetBackgroundParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Brand">Brand</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandColors">BrandColors</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandSettings">BrandSettings</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandSettingsEmail">BrandSettingsEmail</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandSettingsInApp">BrandSettingsInApp</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandSnippet">BrandSnippet</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandSnippets">BrandSnippets</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandTemplate">BrandTemplate</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#EmailFooter">EmailFooter</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#EmailHead">EmailHead</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#EmailHeader">EmailHeader</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Icons">Icons</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Logo">Logo</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#WidgetBackground">WidgetBackground</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandListResponse">BrandListResponse</a>

Methods:

- <code title="post /brands">client.Brands.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandNewParams">BrandNewParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Brand">Brand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /brands/{brand_id}">client.Brands.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, brandID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Brand">Brand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /brands/{brand_id}">client.Brands.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, brandID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandUpdateParams">BrandUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Brand">Brand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /brands">client.Brands.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandListParams">BrandListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandListResponse">BrandListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /brands/{brand_id}">client.Brands.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BrandService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, brandID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Bulk

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#InboundBulkMessageUnionParam">InboundBulkMessageUnionParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#InboundBulkMessageUserParam">InboundBulkMessageUserParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserRecipientParam">UserRecipientParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#InboundBulkMessageUnion">InboundBulkMessageUnion</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#InboundBulkMessageUser">InboundBulkMessageUser</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserRecipient">UserRecipient</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkNewJobResponse">BulkNewJobResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkListUsersResponse">BulkListUsersResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkGetJobResponse">BulkGetJobResponse</a>

Methods:

- <code title="post /bulk/{job_id}">client.Bulk.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkService.AddUsers">AddUsers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkAddUsersParams">BulkAddUsersParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /bulk">client.Bulk.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkService.NewJob">NewJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkNewJobParams">BulkNewJobParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkNewJobResponse">BulkNewJobResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bulk/{job_id}/users">client.Bulk.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkService.ListUsers">ListUsers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkListUsersParams">BulkListUsersParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkListUsersResponse">BulkListUsersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bulk/{job_id}">client.Bulk.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkService.GetJob">GetJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkGetJobResponse">BulkGetJobResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /bulk/{job_id}/run">client.Bulk.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BulkService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Inbound

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#InboundTrackEventResponse">InboundTrackEventResponse</a>

Methods:

- <code title="post /inbound/courier">client.Inbound.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#InboundService.TrackEvent">TrackEvent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#InboundTrackEventParams">InboundTrackEventParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#InboundTrackEventResponse">InboundTrackEventResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Lists

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#List">List</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListListResponse">ListListResponse</a>

Methods:

- <code title="get /lists/{list_id}">client.Lists.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, listID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#List">List</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /lists/{list_id}">client.Lists.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, listID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListUpdateParams">ListUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /lists">client.Lists.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListListParams">ListListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListListResponse">ListListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /lists/{list_id}">client.Lists.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, listID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /lists/{list_id}/restore">client.Lists.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListService.Restore">Restore</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, listID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListRestoreParams">ListRestoreParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Subscriptions

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationPreferenceDetailsParam">NotificationPreferenceDetailsParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#PutSubscriptionsRecipientParam">PutSubscriptionsRecipientParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#RecipientPreferencesParam">RecipientPreferencesParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationPreferenceDetails">NotificationPreferenceDetails</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#RecipientPreferences">RecipientPreferences</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionListResponse">ListSubscriptionListResponse</a>

Methods:

- <code title="get /lists/{list_id}/subscriptions">client.Lists.Subscriptions.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, listID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionListParams">ListSubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionListResponse">ListSubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /lists/{list_id}/subscriptions">client.Lists.Subscriptions.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, listID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionAddParams">ListSubscriptionAddParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /lists/{list_id}/subscriptions">client.Lists.Subscriptions.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionService.Subscribe">Subscribe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, listID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionSubscribeParams">ListSubscriptionSubscribeParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /lists/{list_id}/subscriptions/{user_id}">client.Lists.Subscriptions.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionService.SubscribeUser">SubscribeUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionSubscribeUserParams">ListSubscriptionSubscribeUserParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /lists/{list_id}/subscriptions/{user_id}">client.Lists.Subscriptions.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionService.UnsubscribeUser">UnsubscribeUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ListSubscriptionUnsubscribeUserParams">ListSubscriptionUnsubscribeUserParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageDetails">MessageDetails</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageGetResponse">MessageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageListResponse">MessageListResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageContentResponse">MessageContentResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageHistoryResponse">MessageHistoryResponse</a>

Methods:

- <code title="get /messages/{message_id}">client.Messages.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageGetResponse">MessageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /messages">client.Messages.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageListParams">MessageListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageListResponse">MessageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /messages/{message_id}/cancel">client.Messages.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageDetails">MessageDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /messages/{message_id}/output">client.Messages.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageService.Content">Content</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageContentResponse">MessageContentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /messages/{message_id}/history">client.Messages.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageService.History">History</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageHistoryParams">MessageHistoryParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageHistoryResponse">MessageHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Requests

Methods:

- <code title="put /requests/{request_id}/archive">client.Requests.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#RequestService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Notifications

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageRoutingParam">MessageRoutingParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageRoutingChannelUnionParam">MessageRoutingChannelUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageRouting">MessageRouting</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#MessageRoutingChannelUnion">MessageRoutingChannelUnion</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationGetContent">NotificationGetContent</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationListResponse">NotificationListResponse</a>

Methods:

- <code title="get /notifications">client.Notifications.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationListParams">NotificationListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationListResponse">NotificationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /notifications/{id}/content">client.Notifications.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationService.GetContent">GetContent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationGetContent">NotificationGetContent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Draft

Methods:

- <code title="get /notifications/{id}/draft/content">client.Notifications.Draft.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationDraftService.GetContent">GetContent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationGetContent">NotificationGetContent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Checks

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BaseCheckParam">BaseCheckParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BaseCheck">BaseCheck</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Check">Check</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationCheckUpdateResponse">NotificationCheckUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationCheckListResponse">NotificationCheckListResponse</a>

Methods:

- <code title="put /notifications/{id}/{submissionId}/checks">client.Notifications.Checks.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationCheckService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, submissionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationCheckUpdateParams">NotificationCheckUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationCheckUpdateResponse">NotificationCheckUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /notifications/{id}/{submissionId}/checks">client.Notifications.Checks.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationCheckService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, submissionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationCheckListParams">NotificationCheckListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationCheckListResponse">NotificationCheckListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /notifications/{id}/{submissionId}/checks">client.Notifications.Checks.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationCheckService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, submissionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#NotificationCheckDeleteParams">NotificationCheckDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Profiles

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileNewResponse">ProfileNewResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileGetResponse">ProfileGetResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileReplaceResponse">ProfileReplaceResponse</a>

Methods:

- <code title="post /profiles/{user_id}">client.Profiles.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileNewParams">ProfileNewParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileNewResponse">ProfileNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /profiles/{user_id}">client.Profiles.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileGetResponse">ProfileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /profiles/{user_id}">client.Profiles.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileUpdateParams">ProfileUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /profiles/{user_id}">client.Profiles.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /profiles/{user_id}">client.Profiles.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileService.Replace">Replace</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileReplaceParams">ProfileReplaceParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileReplaceResponse">ProfileReplaceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Lists

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListGetResponse">ProfileListGetResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListDeleteResponse">ProfileListDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListSubscribeResponse">ProfileListSubscribeResponse</a>

Methods:

- <code title="get /profiles/{user_id}/lists">client.Profiles.Lists.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListGetParams">ProfileListGetParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListGetResponse">ProfileListGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /profiles/{user_id}/lists">client.Profiles.Lists.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListDeleteResponse">ProfileListDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /profiles/{user_id}/lists">client.Profiles.Lists.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListService.Subscribe">Subscribe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListSubscribeParams">ProfileListSubscribeParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ProfileListSubscribeResponse">ProfileListSubscribeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tenants

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#DefaultPreferencesParam">DefaultPreferencesParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#DefaultPreferences">DefaultPreferences</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Tenant">Tenant</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantListResponse">TenantListResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantListUsersResponse">TenantListUsersResponse</a>

Methods:

- <code title="get /tenants/{tenant_id}">client.Tenants.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Tenant">Tenant</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /tenants/{tenant_id}">client.Tenants.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantUpdateParams">TenantUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Tenant">Tenant</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /tenants">client.Tenants.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantListParams">TenantListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantListResponse">TenantListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /tenants/{tenant_id}">client.Tenants.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /tenants/{tenant_id}/users">client.Tenants.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantService.ListUsers">ListUsers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantListUsersParams">TenantListUsersParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantListUsersResponse">TenantListUsersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DefaultPreferences

### Items

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ChannelClassification">ChannelClassification</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#SubscriptionTopicNewParam">SubscriptionTopicNewParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ChannelClassification">ChannelClassification</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#SubscriptionTopicNew">SubscriptionTopicNew</a>

Methods:

- <code title="put /tenants/{tenant_id}/default_preferences/items/{topic_id}">client.Tenants.DefaultPreferences.Items.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantDefaultPreferenceItemService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, topicID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantDefaultPreferenceItemUpdateParams">TenantDefaultPreferenceItemUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /tenants/{tenant_id}/default_preferences/items/{topic_id}">client.Tenants.DefaultPreferences.Items.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantDefaultPreferenceItemService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, topicID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantDefaultPreferenceItemDeleteParams">TenantDefaultPreferenceItemDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Templates

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Alignment">Alignment</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ElementalBaseNodeParam">ElementalBaseNodeParam</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ElementalContentParam">ElementalContentParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#Alignment">Alignment</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BaseTemplateTenantAssociation">BaseTemplateTenantAssociation</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ElementalBaseNode">ElementalBaseNode</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#ElementalContent">ElementalContent</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantTemplateListResponse">TenantTemplateListResponse</a>

Methods:

- <code title="get /tenants/{tenant_id}/templates/{template_id}">client.Tenants.Templates.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantTemplateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, templateID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantTemplateGetParams">TenantTemplateGetParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#BaseTemplateTenantAssociation">BaseTemplateTenantAssociation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /tenants/{tenant_id}/templates">client.Tenants.Templates.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantTemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantTemplateListParams">TenantTemplateListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantTemplateListResponse">TenantTemplateListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Translations

Methods:

- <code title="get /translations/{domain}/{locale}">client.Translations.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TranslationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, locale <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TranslationGetParams">TranslationGetParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /translations/{domain}/{locale}">client.Translations.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TranslationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, locale <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TranslationUpdateParams">TranslationUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Users

## Preferences

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#PreferenceStatus">PreferenceStatus</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#PreferenceStatus">PreferenceStatus</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TopicPreference">TopicPreference</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceGetResponse">UserPreferenceGetResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceGetTopicResponse">UserPreferenceGetTopicResponse</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceUpdateOrNewTopicResponse">UserPreferenceUpdateOrNewTopicResponse</a>

Methods:

- <code title="get /users/{user_id}/preferences">client.Users.Preferences.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceGetParams">UserPreferenceGetParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceGetResponse">UserPreferenceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users/{user_id}/preferences/{topic_id}">client.Users.Preferences.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceService.GetTopic">GetTopic</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, topicID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceGetTopicParams">UserPreferenceGetTopicParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceGetTopicResponse">UserPreferenceGetTopicResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/{user_id}/preferences/{topic_id}">client.Users.Preferences.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceService.UpdateOrNewTopic">UpdateOrNewTopic</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, topicID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceUpdateOrNewTopicParams">UserPreferenceUpdateOrNewTopicParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserPreferenceUpdateOrNewTopicResponse">UserPreferenceUpdateOrNewTopicResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tenants

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantAssociationParam">TenantAssociationParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#TenantAssociation">TenantAssociation</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantListResponse">UserTenantListResponse</a>

Methods:

- <code title="get /users/{user_id}/tenants">client.Users.Tenants.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantListParams">UserTenantListParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantListResponse">UserTenantListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/{user_id}/tenants">client.Users.Tenants.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantService.AddMultiple">AddMultiple</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantAddMultipleParams">UserTenantAddMultipleParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /users/{user_id}/tenants/{tenant_id}">client.Users.Tenants.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantService.AddSingle">AddSingle</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantAddSingleParams">UserTenantAddSingleParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /users/{user_id}/tenants">client.Users.Tenants.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantService.RemoveAll">RemoveAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /users/{user_id}/tenants/{tenant_id}">client.Users.Tenants.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantService.RemoveSingle">RemoveSingle</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTenantRemoveSingleParams">UserTenantRemoveSingleParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Tokens

Params Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenParam">UserTokenParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserToken">UserToken</a>
- <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenGetResponse">UserTokenGetResponse</a>

Methods:

- <code title="get /users/{user_id}/tokens/{token}">client.Users.Tokens.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenGetParams">UserTokenGetParams</a>) (<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenGetResponse">UserTokenGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /users/{user_id}/tokens/{token}">client.Users.Tokens.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenUpdateParams">UserTokenUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /users/{user_id}/tokens">client.Users.Tokens.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserToken">UserToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /users/{user_id}/tokens/{token}">client.Users.Tokens.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenDeleteParams">UserTokenDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /users/{user_id}/tokens">client.Users.Tokens.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenService.AddMultiple">AddMultiple</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /users/{user_id}/tokens/{token}">client.Users.Tokens.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenService.AddSingle">AddSingle</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/trycourier/courier-go">courier</a>.<a href="https://pkg.go.dev/github.com/trycourier/courier-go#UserTokenAddSingleParams">UserTokenAddSingleParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
