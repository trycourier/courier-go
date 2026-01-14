// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/trycourier/courier-go/v4/internal/apierror"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type AudienceFilterParam = shared.AudienceFilterParam

// Send to users only if they are member of the account
//
// This is an alias to an internal type.
type AudienceFilterOperator = shared.AudienceFilterOperator

// Equals "MEMBER_OF"
const AudienceFilterOperatorMemberOf = shared.AudienceFilterOperatorMemberOf

// This is an alias to an internal type.
type AudienceFilterPath = shared.AudienceFilterPath

// Equals "account_id"
const AudienceFilterPathAccountID = shared.AudienceFilterPathAccountID

// Filter configuration for audience membership containing an array of filter rules
//
// This is an alias to an internal type.
type AudienceFilterConfig = shared.AudienceFilterConfig

// Filter configuration for audience membership containing an array of filter rules
//
// This is an alias to an internal type.
type AudienceFilterConfigParam = shared.AudienceFilterConfigParam

// Send to all users in an audience
//
// This is an alias to an internal type.
type AudienceRecipientParam = shared.AudienceRecipientParam

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

// A filter rule that can be either a single condition (with path/value) or a
// nested group (with filters array). Use comparison operators (EQ, GT, etc.) for
// single conditions, and logical operators (AND, OR) for nested groups.
//
// This is an alias to an internal type.
type FilterConfig = shared.FilterConfig

// A filter rule that can be either a single condition (with path/value) or a
// nested group (with filters array). Use comparison operators (EQ, GT, etc.) for
// single conditions, and logical operators (AND, OR) for nested groups.
//
// This is an alias to an internal type.
type FilterConfigParam = shared.FilterConfigParam

// This is an alias to an internal type.
type ListFilterParam = shared.ListFilterParam

// Send to users only if they are member of the account
//
// This is an alias to an internal type.
type ListFilterOperator = shared.ListFilterOperator

// Equals "MEMBER_OF"
const ListFilterOperatorMemberOf = shared.ListFilterOperatorMemberOf

// This is an alias to an internal type.
type ListFilterPath = shared.ListFilterPath

// Equals "account_id"
const ListFilterPathAccountID = shared.ListFilterPathAccountID

// Send to users in lists matching a pattern
//
// This is an alias to an internal type.
type ListPatternRecipientParam = shared.ListPatternRecipientParam

// Send to all users in a specific list
//
// This is an alias to an internal type.
type ListRecipientParam = shared.ListRecipientParam

// This is an alias to an internal type.
type MessageContext = shared.MessageContext

// This is an alias to an internal type.
type MessageContextParam = shared.MessageContextParam

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
type MsTeamsUnionParam = shared.MsTeamsUnionParam

// Send via Microsoft Teams
//
// This is an alias to an internal type.
type MsTeamsRecipientParam = shared.MsTeamsRecipientParam

// This is an alias to an internal type.
type NotificationPreferenceDetails = shared.NotificationPreferenceDetails

// This is an alias to an internal type.
type NotificationPreferenceDetailsParam = shared.NotificationPreferenceDetailsParam

// This is an alias to an internal type.
type PagerdutyParam = shared.PagerdutyParam

// Send via PagerDuty
//
// This is an alias to an internal type.
type PagerdutyRecipientParam = shared.PagerdutyRecipientParam

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
type RecipientPreferences = shared.RecipientPreferences

// This is an alias to an internal type.
type RecipientPreferencesParam = shared.RecipientPreferencesParam

// This is an alias to an internal type.
type Rule = shared.Rule

// This is an alias to an internal type.
type RuleParam = shared.RuleParam

// This is an alias to an internal type.
type SendToMsTeamsChannelIDParam = shared.SendToMsTeamsChannelIDParam

// This is an alias to an internal type.
type SendToMsTeamsChannelNameParam = shared.SendToMsTeamsChannelNameParam

// This is an alias to an internal type.
type SendToMsTeamsConversationIDParam = shared.SendToMsTeamsConversationIDParam

// This is an alias to an internal type.
type SendToMsTeamsEmailParam = shared.SendToMsTeamsEmailParam

// This is an alias to an internal type.
type SendToMsTeamsUserIDParam = shared.SendToMsTeamsUserIDParam

// This is an alias to an internal type.
type SendToSlackChannelParam = shared.SendToSlackChannelParam

// This is an alias to an internal type.
type SendToSlackEmailParam = shared.SendToSlackEmailParam

// This is an alias to an internal type.
type SendToSlackUserIDParam = shared.SendToSlackUserIDParam

// This is an alias to an internal type.
type SlackUnionParam = shared.SlackUnionParam

// Send via Slack (channel, email, or user_id)
//
// This is an alias to an internal type.
type SlackRecipientParam = shared.SlackRecipientParam

// This is an alias to an internal type.
type UserRecipient = shared.UserRecipient

// This is an alias to an internal type.
type UserRecipientPreferences = shared.UserRecipientPreferences

// This is an alias to an internal type.
type UserRecipientParam = shared.UserRecipientParam

// This is an alias to an internal type.
type UserRecipientPreferencesParam = shared.UserRecipientPreferencesParam

// This is an alias to an internal type.
type UtmParam = shared.UtmParam

// This is an alias to an internal type.
type WebhookAuthMode = shared.WebhookAuthMode

// Equals "none"
const WebhookAuthModeNone = shared.WebhookAuthModeNone

// Equals "basic"
const WebhookAuthModeBasic = shared.WebhookAuthModeBasic

// Equals "bearer"
const WebhookAuthModeBearer = shared.WebhookAuthModeBearer

// This is an alias to an internal type.
type WebhookAuthenticationParam = shared.WebhookAuthenticationParam

// This is an alias to an internal type.
type WebhookMethod = shared.WebhookMethod

// Equals "POST"
const WebhookMethodPost = shared.WebhookMethodPost

// Equals "PUT"
const WebhookMethodPut = shared.WebhookMethodPut

// This is an alias to an internal type.
type WebhookProfileParam = shared.WebhookProfileParam

// This is an alias to an internal type.
type WebhookProfileType = shared.WebhookProfileType

// Equals "limited"
const WebhookProfileTypeLimited = shared.WebhookProfileTypeLimited

// Equals "expanded"
const WebhookProfileTypeExpanded = shared.WebhookProfileTypeExpanded

// Send via webhook
//
// This is an alias to an internal type.
type WebhookRecipientParam = shared.WebhookRecipientParam
