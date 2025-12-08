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
type ProfilePreferences = shared.ProfilePreferences

// This is an alias to an internal type.
type ProfilePreferencesParam = shared.ProfilePreferencesParam

// This is an alias to an internal type.
type RecipientParam = shared.RecipientParam

// This is an alias to an internal type.
type RecipientPreferences = shared.RecipientPreferences

// This is an alias to an internal type.
type RecipientPreferencesParam = shared.RecipientPreferencesParam

// This is an alias to an internal type.
type Rule = shared.Rule

// This is an alias to an internal type.
type RuleParam = shared.RuleParam

// This is an alias to an internal type.
type UserRecipient = shared.UserRecipient

// This is an alias to an internal type.
type UserRecipientParam = shared.UserRecipientParam

// This is an alias to an internal type.
type UtmParam = shared.UtmParam
