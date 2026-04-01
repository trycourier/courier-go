# Courier Go SDK

Courier is a notifications API for sending messages across email, SMS, push, in-app inbox, Slack, and WhatsApp from a single API call.

## Setup

```go
import (
    "github.com/trycourier/courier-go/v4"
    "github.com/trycourier/courier-go/v4/option"
    "github.com/trycourier/courier-go/v4/shared"
)

client := courier.NewClient() // reads COURIER_API_KEY from env
```

## Core pattern

```go
response, err := client.Send.Message(ctx, courier.SendMessageParams{
    Message: courier.SendMessageParamsMessage{
        To: courier.SendMessageParamsMessageToUnion{
            OfUserRecipient: &shared.UserRecipientParam{
                UserID: courier.String("user_123"),
            },
        },
        Template: courier.String("TEMPLATE_ID"),
        Data:     map[string]any{"order_id": "456"},
    },
})
fmt.Println(response.RequestID)
```

## Key rules

- Use `routing.method: "single"` (fallback chain) unless the user explicitly asks for parallel delivery (`"all"`).
- Use `client.Profiles.Create()` for partial profile updates (it merges). Use `client.Profiles.Replace()` only when fully replacing all profile data.
- Test and production use different API keys from the same workspace. Always confirm which environment before sending.
- For transactional sends (OTP, orders, billing), pass an `Idempotency-Key` header via `option.WithHeader()` to prevent duplicates.
- Bulk sends are a 3-step flow: `client.Bulk.Create()` → `client.Bulk.AddUsers()` → `client.Bulk.Run()`.
- `RequestID` from a single-recipient send doubles as the `message_id`. For multi-recipient sends, each recipient gets a unique `message_id`.
- Go uses typed structs for all params; use union wrappers like `SendMessageParamsMessageToUnion` for polymorphic fields.

## Concepts

- `Template` — notification template ID from the Courier dashboard
- `Routing.Method` — `"single"` = try channels in order until one succeeds; `"all"` = send on every channel simultaneously
- `TenantID` — multi-tenant context; affects brand and preference defaults for the message
- `ListID` — send to all subscribers of a named list
- `UserRecipientParam` — registered user whose profile has channel addresses
- `EmailRecipientParam` / `PhoneNumberRecipientParam` — ad-hoc recipient (no stored profile needed)

## More context

- Full docs index: https://www.courier.com/docs/llms.txt
- API reference: https://www.courier.com/docs/reference/get-started
- MCP server: https://mcp.courier.com
- Courier Skills (Cursor / Claude Code): https://github.com/trycourier/courier-skills
