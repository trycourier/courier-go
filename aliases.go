// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/trycourier/courier-go/internal/apierror"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type ChannelPreference = shared.ChannelPreference

// This is an alias to an internal type.
type ChannelPreferenceParam = shared.ChannelPreferenceParam

// This is an alias to an internal type.
type ElementalBaseNode = shared.ElementalBaseNode

// This is an alias to an internal type.
type ElementalBaseNodeParam = shared.ElementalBaseNodeParam

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// This is an alias to an internal type.
type ElementalContentSugar = shared.ElementalContentSugar

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// This is an alias to an internal type.
type ElementalContentSugarParam = shared.ElementalContentSugarParam

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
type Rule = shared.Rule

// This is an alias to an internal type.
type RuleParam = shared.RuleParam
