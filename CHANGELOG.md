# Change Log

All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased][unreleased]
### Added
- Support for [Lists API](https://docs.courier.com/reference/lists-api) by @tk26
  - `GET /lists` with `courier.GetLists` method
  - `GET /list/{list_id}` with `courier.GetList` method
  - `PUT /list/{list_id}` with `courier.PutList` method
  - `DELETE /list/{list_id}` with `courier.DeleteList` method
  - `PUT /list/{list_id}/restore` with `courier.RestoreList` method
  - `GET /list/{list_id}/subscriptions` with `courier.GetListSubscriptions` method
  - `PUT /list/{list_id}/subscriptions` with `courier.PutListSubscriptions` method
  - `POST /list/{list_id}/subscriptions` with `courier.PostListSubscriptions` method
  - `PUT /lists/{list_id}/subscriptions/{recipient_id}` with `courier.PutListSubscriptionsRecipient` method
  - `DELETE /lists/{list_id}/subscriptions/{recipient_id}` with `courier.DeleteListSubscriptionsRecipient` method
  - Usage Examples in v2/examples/lists

## [v2.1.0] - 2021-04-09
### Added
- Support for [Automation API](https://docs.courier.com/reference/automation-api) by @tk26
  - `POST /automations/invoke` with `courier.InvokeAutomation` method
  - `POST /automations/{template_id}/invoke` with `courier.InvokeAutomationTemplate` method
  - Usage Examples in v2/examples/automations
- Changed default base URL to `https://api.courier.com`

## [v2.0.2] - 2020-09-22
* Adds brand and override support to SendBody struct

## [v2.0.1] - 2020-09-17
* Fixes issues where a body was being set on GET requests

## [v2.0.0] - 2020-06-24

* `courier.CourierClient` has been renamed to `courier.CreateClient`
* `courier.CreateClient` now takes a `nil` second parameter if you wish to use our default API URL, not `""`
* Our underlying API communication functions are now exposed via `client.API`
* renamed `SendRequest` to `SendBody` and modified the struct
* `Send` now takes one `body interface{}` argument rather than separate `profile` and `data` arguments
* added `SendMap`
* `GetProfile` now requires a `context` argument 
* `GetProfile` now returns a `map[string]json.RawMessage{}` instead of hydrating a struct reference
* removed `MergeProfile`; use `client.API.SendRequestWithBytes(context.Background(), "PUT", "/profiles/"+recipientID, profile)` instead
* removed `UpdateProfile`; use `client.API.SendRequestWithBytes(context.Background(), "PUT", "/profiles/"+recipientID, profile)` instead
