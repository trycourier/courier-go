// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/trycourier/courier-go/v3/internal/apierror"
	"github.com/trycourier/courier-go/v3/packages/param"
	"github.com/trycourier/courier-go/v3/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type Audience = shared.Audience

// This is an alias to an internal type.
type AuditEvent = shared.AuditEvent

// This is an alias to an internal type.
type AuditEventActor = shared.AuditEventActor

// This is an alias to an internal type.
type AutomationInvokeResponse = shared.AutomationInvokeResponse

// This is an alias to an internal type.
type BaseCheck = shared.BaseCheck

// This is an alias to an internal type.
type BaseCheckStatus = shared.BaseCheckStatus

// Equals "RESOLVED"
const BaseCheckStatusResolved = shared.BaseCheckStatusResolved

// Equals "FAILED"
const BaseCheckStatusFailed = shared.BaseCheckStatusFailed

// Equals "PENDING"
const BaseCheckStatusPending = shared.BaseCheckStatusPending

// This is an alias to an internal type.
type BaseCheckType = shared.BaseCheckType

// Equals "custom"
const BaseCheckTypeCustom = shared.BaseCheckTypeCustom

// This is an alias to an internal type.
type BaseCheckParam = shared.BaseCheckParam

// This is an alias to an internal type.
type BaseTemplateTenantAssociation = shared.BaseTemplateTenantAssociation

// This is an alias to an internal type.
type Brand = shared.Brand

// This is an alias to an internal type.
type BrandColors = shared.BrandColors

// This is an alias to an internal type.
type BrandColorsParam = shared.BrandColorsParam

// This is an alias to an internal type.
type BrandSettings = shared.BrandSettings

// This is an alias to an internal type.
type BrandSettingsParam = shared.BrandSettingsParam

// This is an alias to an internal type.
type BrandSettingsEmail = shared.BrandSettingsEmail

// This is an alias to an internal type.
type BrandSettingsEmailTemplateOverride = shared.BrandSettingsEmailTemplateOverride

// This is an alias to an internal type.
type BrandSettingsEmailParam = shared.BrandSettingsEmailParam

// This is an alias to an internal type.
type BrandSettingsEmailTemplateOverrideParam = shared.BrandSettingsEmailTemplateOverrideParam

// This is an alias to an internal type.
type BrandSettingsInApp = shared.BrandSettingsInApp

// This is an alias to an internal type.
type BrandSettingsInAppPlacement = shared.BrandSettingsInAppPlacement

// Equals "top"
const BrandSettingsInAppPlacementTop = shared.BrandSettingsInAppPlacementTop

// Equals "bottom"
const BrandSettingsInAppPlacementBottom = shared.BrandSettingsInAppPlacementBottom

// Equals "left"
const BrandSettingsInAppPlacementLeft = shared.BrandSettingsInAppPlacementLeft

// Equals "right"
const BrandSettingsInAppPlacementRight = shared.BrandSettingsInAppPlacementRight

// This is an alias to an internal type.
type BrandSettingsInAppParam = shared.BrandSettingsInAppParam

// This is an alias to an internal type.
type BrandSnippet = shared.BrandSnippet

// This is an alias to an internal type.
type BrandSnippetParam = shared.BrandSnippetParam

// This is an alias to an internal type.
type BrandSnippets = shared.BrandSnippets

// This is an alias to an internal type.
type BrandSnippetsParam = shared.BrandSnippetsParam

// This is an alias to an internal type.
type BrandTemplate = shared.BrandTemplate

// This is an alias to an internal type.
type BrandTemplateParam = shared.BrandTemplateParam

// This is an alias to an internal type.
type ChannelClassification = shared.ChannelClassification

// Equals "direct_message"
const ChannelClassificationDirectMessage = shared.ChannelClassificationDirectMessage

// Equals "email"
const ChannelClassificationEmail = shared.ChannelClassificationEmail

// Equals "push"
const ChannelClassificationPush = shared.ChannelClassificationPush

// Equals "sms"
const ChannelClassificationSMS = shared.ChannelClassificationSMS

// Equals "webhook"
const ChannelClassificationWebhook = shared.ChannelClassificationWebhook

// Equals "inbox"
const ChannelClassificationInbox = shared.ChannelClassificationInbox

// This is an alias to an internal type.
type ChannelPreference = shared.ChannelPreference

// This is an alias to an internal type.
type ChannelPreferenceParam = shared.ChannelPreferenceParam

// This is an alias to an internal type.
type Check = shared.Check

// This is an alias to an internal type.
type DefaultPreferences = shared.DefaultPreferences

// This is an alias to an internal type.
type DefaultPreferencesItem = shared.DefaultPreferencesItem

// This is an alias to an internal type.
type DefaultPreferencesParam = shared.DefaultPreferencesParam

// This is an alias to an internal type.
type DefaultPreferencesItemParam = shared.DefaultPreferencesItemParam

// This is an alias to an internal type.
type ElementalActionNodeWithType = shared.ElementalActionNodeWithType

// This is an alias to an internal type.
type ElementalActionNodeWithTypeParam = shared.ElementalActionNodeWithTypeParam

// This is an alias to an internal type.
type ElementalBaseNode = shared.ElementalBaseNode

// This is an alias to an internal type.
type ElementalBaseNodeParam = shared.ElementalBaseNodeParam

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
//
// This is an alias to an internal type.
type ElementalChannelNode = shared.ElementalChannelNode

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
//
// This is an alias to an internal type.
type ElementalChannelNodeParam = shared.ElementalChannelNodeParam

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
//
// This is an alias to an internal type.
type ElementalChannelNodeWithType = shared.ElementalChannelNodeWithType

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
//
// This is an alias to an internal type.
type ElementalChannelNodeWithTypeParam = shared.ElementalChannelNodeWithTypeParam

// This is an alias to an internal type.
type ElementalContent = shared.ElementalContent

// This is an alias to an internal type.
type ElementalContentParam = shared.ElementalContentParam

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// This is an alias to an internal type.
type ElementalContentSugar = shared.ElementalContentSugar

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// This is an alias to an internal type.
type ElementalContentSugarParam = shared.ElementalContentSugarParam

// This is an alias to an internal type.
type ElementalDividerNodeWithType = shared.ElementalDividerNodeWithType

// This is an alias to an internal type.
type ElementalDividerNodeWithTypeParam = shared.ElementalDividerNodeWithTypeParam

// This is an alias to an internal type.
type ElementalImageNodeWithType = shared.ElementalImageNodeWithType

// This is an alias to an internal type.
type ElementalImageNodeWithTypeParam = shared.ElementalImageNodeWithTypeParam

// This is an alias to an internal type.
type ElementalMetaNodeWithType = shared.ElementalMetaNodeWithType

// This is an alias to an internal type.
type ElementalMetaNodeWithTypeParam = shared.ElementalMetaNodeWithTypeParam

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
//
// This is an alias to an internal type.
type ElementalNodeUnion = shared.ElementalNodeUnion

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
//
// This is an alias to an internal type.
type ElementalNodeUnionParam = shared.ElementalNodeUnionParam

// This is an alias to an internal type.
type ElementalQuoteNodeWithType = shared.ElementalQuoteNodeWithType

// This is an alias to an internal type.
type ElementalQuoteNodeWithTypeParam = shared.ElementalQuoteNodeWithTypeParam

// This is an alias to an internal type.
type ElementalTextNodeWithType = shared.ElementalTextNodeWithType

// This is an alias to an internal type.
type ElementalTextNodeWithTypeParam = shared.ElementalTextNodeWithTypeParam

// This is an alias to an internal type.
type EmailFooter = shared.EmailFooter

// This is an alias to an internal type.
type EmailFooterParam = shared.EmailFooterParam

// This is an alias to an internal type.
type EmailHead = shared.EmailHead

// This is an alias to an internal type.
type EmailHeadParam = shared.EmailHeadParam

// This is an alias to an internal type.
type EmailHeader = shared.EmailHeader

// This is an alias to an internal type.
type EmailHeaderParam = shared.EmailHeaderParam

// This is an alias to an internal type.
type Filter = shared.Filter

// The operator to use for filtering
//
// This is an alias to an internal type.
type FilterOperator = shared.FilterOperator

// Equals "ENDS_WITH"
const FilterOperatorEndsWith = shared.FilterOperatorEndsWith

// Equals "EQ"
const FilterOperatorEq = shared.FilterOperatorEq

// Equals "EXISTS"
const FilterOperatorExists = shared.FilterOperatorExists

// Equals "GT"
const FilterOperatorGt = shared.FilterOperatorGt

// Equals "GTE"
const FilterOperatorGte = shared.FilterOperatorGte

// Equals "INCLUDES"
const FilterOperatorIncludes = shared.FilterOperatorIncludes

// Equals "IS_AFTER"
const FilterOperatorIsAfter = shared.FilterOperatorIsAfter

// Equals "IS_BEFORE"
const FilterOperatorIsBefore = shared.FilterOperatorIsBefore

// Equals "LT"
const FilterOperatorLt = shared.FilterOperatorLt

// Equals "LTE"
const FilterOperatorLte = shared.FilterOperatorLte

// Equals "NEQ"
const FilterOperatorNeq = shared.FilterOperatorNeq

// Equals "OMIT"
const FilterOperatorOmit = shared.FilterOperatorOmit

// Equals "STARTS_WITH"
const FilterOperatorStartsWith = shared.FilterOperatorStartsWith

// Equals "AND"
const FilterOperatorAnd = shared.FilterOperatorAnd

// Equals "OR"
const FilterOperatorOr = shared.FilterOperatorOr

// This is an alias to an internal type.
type FilterParam = shared.FilterParam

// This is an alias to an internal type.
type Icons = shared.Icons

// This is an alias to an internal type.
type IconsParam = shared.IconsParam

// This is an alias to an internal type.
type InboundBulkMessageUnion = shared.InboundBulkMessageUnion

// This is an alias to an internal type.
type InboundBulkMessageInboundBulkTemplateMessage = shared.InboundBulkMessageInboundBulkTemplateMessage

// This is an alias to an internal type.
type InboundBulkMessageInboundBulkContentMessage = shared.InboundBulkMessageInboundBulkContentMessage

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// This is an alias to an internal type.
type InboundBulkMessageInboundBulkContentMessageContentUnion = shared.InboundBulkMessageInboundBulkContentMessageContentUnion

// This is an alias to an internal type.
type InboundBulkMessageUnionParam = shared.InboundBulkMessageUnionParam

// This is an alias to an internal type.
type InboundBulkMessageInboundBulkTemplateMessageParam = shared.InboundBulkMessageInboundBulkTemplateMessageParam

// This is an alias to an internal type.
type InboundBulkMessageInboundBulkContentMessageParam = shared.InboundBulkMessageInboundBulkContentMessageParam

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// This is an alias to an internal type.
type InboundBulkMessageInboundBulkContentMessageContentUnionParam = shared.InboundBulkMessageInboundBulkContentMessageContentUnionParam

// This is an alias to an internal type.
type InboundBulkMessageUser = shared.InboundBulkMessageUser

// This is an alias to an internal type.
type InboundBulkMessageUserParam = shared.InboundBulkMessageUserParam

// This is an alias to an internal type.
type Logo = shared.Logo

// This is an alias to an internal type.
type LogoParam = shared.LogoParam

// This is an alias to an internal type.
type MessageContext = shared.MessageContext

// This is an alias to an internal type.
type MessageContextParam = shared.MessageContextParam

// This is an alias to an internal type.
type MessageDetails = shared.MessageDetails

// The current status of the message.
//
// This is an alias to an internal type.
type MessageDetailsStatus = shared.MessageDetailsStatus

// Equals "CANCELED"
const MessageDetailsStatusCanceled = shared.MessageDetailsStatusCanceled

// Equals "CLICKED"
const MessageDetailsStatusClicked = shared.MessageDetailsStatusClicked

// Equals "DELAYED"
const MessageDetailsStatusDelayed = shared.MessageDetailsStatusDelayed

// Equals "DELIVERED"
const MessageDetailsStatusDelivered = shared.MessageDetailsStatusDelivered

// Equals "DIGESTED"
const MessageDetailsStatusDigested = shared.MessageDetailsStatusDigested

// Equals "ENQUEUED"
const MessageDetailsStatusEnqueued = shared.MessageDetailsStatusEnqueued

// Equals "FILTERED"
const MessageDetailsStatusFiltered = shared.MessageDetailsStatusFiltered

// Equals "OPENED"
const MessageDetailsStatusOpened = shared.MessageDetailsStatusOpened

// Equals "ROUTED"
const MessageDetailsStatusRouted = shared.MessageDetailsStatusRouted

// Equals "SENT"
const MessageDetailsStatusSent = shared.MessageDetailsStatusSent

// Equals "SIMULATED"
const MessageDetailsStatusSimulated = shared.MessageDetailsStatusSimulated

// Equals "THROTTLED"
const MessageDetailsStatusThrottled = shared.MessageDetailsStatusThrottled

// Equals "UNDELIVERABLE"
const MessageDetailsStatusUndeliverable = shared.MessageDetailsStatusUndeliverable

// Equals "UNMAPPED"
const MessageDetailsStatusUnmapped = shared.MessageDetailsStatusUnmapped

// Equals "UNROUTABLE"
const MessageDetailsStatusUnroutable = shared.MessageDetailsStatusUnroutable

// The reason for the current status of the message.
//
// This is an alias to an internal type.
type MessageDetailsReason = shared.MessageDetailsReason

// Equals "BOUNCED"
const MessageDetailsReasonBounced = shared.MessageDetailsReasonBounced

// Equals "FAILED"
const MessageDetailsReasonFailed = shared.MessageDetailsReasonFailed

// Equals "FILTERED"
const MessageDetailsReasonFiltered = shared.MessageDetailsReasonFiltered

// Equals "NO_CHANNELS"
const MessageDetailsReasonNoChannels = shared.MessageDetailsReasonNoChannels

// Equals "NO_PROVIDERS"
const MessageDetailsReasonNoProviders = shared.MessageDetailsReasonNoProviders

// Equals "OPT_IN_REQUIRED"
const MessageDetailsReasonOptInRequired = shared.MessageDetailsReasonOptInRequired

// Equals "PROVIDER_ERROR"
const MessageDetailsReasonProviderError = shared.MessageDetailsReasonProviderError

// Equals "UNPUBLISHED"
const MessageDetailsReasonUnpublished = shared.MessageDetailsReasonUnpublished

// Equals "UNSUBSCRIBED"
const MessageDetailsReasonUnsubscribed = shared.MessageDetailsReasonUnsubscribed

// This is an alias to an internal type.
type MessageRouting = shared.MessageRouting

// This is an alias to an internal type.
type MessageRoutingMethod = shared.MessageRoutingMethod

// Equals "all"
const MessageRoutingMethodAll = shared.MessageRoutingMethodAll

// Equals "single"
const MessageRoutingMethodSingle = shared.MessageRoutingMethodSingle

// This is an alias to an internal type.
type MessageRoutingParam = shared.MessageRoutingParam

// This is an alias to an internal type.
type MessageRoutingChannelUnion = shared.MessageRoutingChannelUnion

// This is an alias to an internal type.
type MessageRoutingChannelUnionParam = shared.MessageRoutingChannelUnionParam

// This is an alias to an internal type.
type NotificationGetContent = shared.NotificationGetContent

// This is an alias to an internal type.
type NotificationGetContentBlock = shared.NotificationGetContentBlock

// This is an alias to an internal type.
type NotificationGetContentBlockContentUnion = shared.NotificationGetContentBlockContentUnion

// This is an alias to an internal type.
type NotificationGetContentBlockContentNotificationContentHierarchy = shared.NotificationGetContentBlockContentNotificationContentHierarchy

// This is an alias to an internal type.
type NotificationGetContentBlockLocaleUnion = shared.NotificationGetContentBlockLocaleUnion

// This is an alias to an internal type.
type NotificationGetContentBlockLocaleNotificationContentHierarchy = shared.NotificationGetContentBlockLocaleNotificationContentHierarchy

// This is an alias to an internal type.
type NotificationGetContentChannel = shared.NotificationGetContentChannel

// This is an alias to an internal type.
type NotificationGetContentChannelContent = shared.NotificationGetContentChannelContent

// This is an alias to an internal type.
type NotificationGetContentChannelLocale = shared.NotificationGetContentChannelLocale

// This is an alias to an internal type.
type NotificationPreferenceDetails = shared.NotificationPreferenceDetails

// This is an alias to an internal type.
type NotificationPreferenceDetailsParam = shared.NotificationPreferenceDetailsParam

// This is an alias to an internal type.
type Paging = shared.Paging

// This is an alias to an internal type.
type Preference = shared.Preference

// This is an alias to an internal type.
type PreferenceSource = shared.PreferenceSource

// Equals "subscription"
const PreferenceSourceSubscription = shared.PreferenceSourceSubscription

// Equals "list"
const PreferenceSourceList = shared.PreferenceSourceList

// Equals "recipient"
const PreferenceSourceRecipient = shared.PreferenceSourceRecipient

// This is an alias to an internal type.
type PreferenceParam = shared.PreferenceParam

// This is an alias to an internal type.
type PreferenceStatus = shared.PreferenceStatus

// Equals "OPTED_IN"
const PreferenceStatusOptedIn = shared.PreferenceStatusOptedIn

// Equals "OPTED_OUT"
const PreferenceStatusOptedOut = shared.PreferenceStatusOptedOut

// Equals "REQUIRED"
const PreferenceStatusRequired = shared.PreferenceStatusRequired

// This is an alias to an internal type.
type PutSubscriptionsRecipientParam = shared.PutSubscriptionsRecipientParam

// This is an alias to an internal type.
type RecipientParam = shared.RecipientParam

// This is an alias to an internal type.
type RecipientPreferencesParam = shared.RecipientPreferencesParam

// This is an alias to an internal type.
type RecipientPreferences = shared.RecipientPreferences

// This is an alias to an internal type.
type Rule = shared.Rule

// This is an alias to an internal type.
type RuleParam = shared.RuleParam

// This is an alias to an internal type.
type SubscribeToListsRequestItemParam = shared.SubscribeToListsRequestItemParam

// This is an alias to an internal type.
type SubscriptionList = shared.SubscriptionList

// This is an alias to an internal type.
type SubscriptionTopicNew = shared.SubscriptionTopicNew

// This is an alias to an internal type.
type SubscriptionTopicNewStatus = shared.SubscriptionTopicNewStatus

// Equals "OPTED_OUT"
const SubscriptionTopicNewStatusOptedOut = shared.SubscriptionTopicNewStatusOptedOut

// Equals "OPTED_IN"
const SubscriptionTopicNewStatusOptedIn = shared.SubscriptionTopicNewStatusOptedIn

// Equals "REQUIRED"
const SubscriptionTopicNewStatusRequired = shared.SubscriptionTopicNewStatusRequired

// This is an alias to an internal type.
type SubscriptionTopicNewParam = shared.SubscriptionTopicNewParam

// This is an alias to an internal type.
type Tenant = shared.Tenant

// This is an alias to an internal type.
type TenantAssociation = shared.TenantAssociation

// This is an alias to an internal type.
type TenantAssociationType = shared.TenantAssociationType

// Equals "user"
const TenantAssociationTypeUser = shared.TenantAssociationTypeUser

// This is an alias to an internal type.
type TenantAssociationParam = shared.TenantAssociationParam

// This is an alias to an internal type.
type TopicPreference = shared.TopicPreference

// This is an alias to an internal type.
type UserRecipient = shared.UserRecipient

// This is an alias to an internal type.
type UserRecipientPreferences = shared.UserRecipientPreferences

// This is an alias to an internal type.
type UserRecipientParam = shared.UserRecipientParam

// This is an alias to an internal type.
type UserRecipientPreferencesParam = shared.UserRecipientPreferencesParam

// This is an alias to an internal type.
type UserToken = shared.UserToken

// This is an alias to an internal type.
type UserTokenProviderKey = shared.UserTokenProviderKey

// Equals "firebase-fcm"
const UserTokenProviderKeyFirebaseFcm = shared.UserTokenProviderKeyFirebaseFcm

// Equals "apn"
const UserTokenProviderKeyApn = shared.UserTokenProviderKeyApn

// Equals "expo"
const UserTokenProviderKeyExpo = shared.UserTokenProviderKeyExpo

// Equals "onesignal"
const UserTokenProviderKeyOnesignal = shared.UserTokenProviderKeyOnesignal

// Information about the device the token is associated with.
//
// This is an alias to an internal type.
type UserTokenDevice = shared.UserTokenDevice

// ISO 8601 formatted date the token expires. Defaults to 2 months. Set to false to
// disable expiration.
//
// This is an alias to an internal type.
type UserTokenExpiryDateUnion = shared.UserTokenExpiryDateUnion

// Information about the device the token is associated with.
//
// This is an alias to an internal type.
type UserTokenTracking = shared.UserTokenTracking

// This is an alias to an internal type.
type UserTokenParam = shared.UserTokenParam

// Information about the device the token is associated with.
//
// This is an alias to an internal type.
type UserTokenDeviceParam = shared.UserTokenDeviceParam

// ISO 8601 formatted date the token expires. Defaults to 2 months. Set to false to
// disable expiration.
//
// This is an alias to an internal type.
type UserTokenExpiryDateUnionParam = shared.UserTokenExpiryDateUnionParam

// Information about the device the token is associated with.
//
// This is an alias to an internal type.
type UserTokenTrackingParam = shared.UserTokenTrackingParam

// This is an alias to an internal type.
type UtmParam = shared.UtmParam

// This is an alias to an internal type.
type WidgetBackground = shared.WidgetBackground

// This is an alias to an internal type.
type WidgetBackgroundParam = shared.WidgetBackgroundParam
