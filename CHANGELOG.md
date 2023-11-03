# Change Log

All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased][unreleased]

## [v3.1.0] - 2023-11-03
- Add cancel message endpoint #60
- bumps version #61

## [v3.0.0] - 2023-08-09

- Accounts -> Tenants


## [v2.11.0] - 2023-07-31

- Add pagination attributes for accounts API

## [v2.10.0] - 2023-07-24

- Add accounts API

## [v2.9.0] - 2023-06-30

- Add missing type to message struct

## [v2.8.0] - 2023-06-22

- Support for Cancel Message API

## [v2.7.0] - 2022-07-26

- Support for Audit Events API

## [v2.5.0] - 2022-01-31

### Added

- Support for enhanced Send API features by @tk26

## [v2.4.0] - 2022-01-26

### Added

- Support for Bulk API by @tk26

## [v2.3.0] - 2021-08-23

### Added

- Support for Notifications API by @tk26
  - `GET /notifications` with `courier.GetNotifications` method
  - `GET /notifications/{notification_id}/content` with `courier.GetNotificationContent` method
  - `GET /notifications/{notification_id}/draft/content` with `courier.GetNotificationDraftContent` method
  - `POST /notifications/{notification_id}/variations` with `courier.PostNotificationVariations` method
  - `POST /notifications/{notification_id}/draft/variations` with `courier.PostNotificationDraftVariations` method
  - `GET /notifications/{notification_id}/{submission_id}/checks` with `courier.GetNotificationSubmissionChecks` method
  - `PUT /notifications/{notification_id}/{submission_id}/checks` with `courier.PutNotificationSubmissionChecks` method
  - `DELETE /notifications/{notification_id}/{submission_id}/checks` with `courier.CancelNotificationSubmission` method
  - Usage Examples in v2/examples/notifications

## [v2.2.1] - 2021-05-11

- Fixes [Issue 26](https://github.com/trycourier/courier-go/issues/26)

## [v2.2.0] - 2021-05-07

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
  - `PUT /lists/{list_id}/subscriptions/{recipient_id}` with `courier.ListSubscribe` method
  - `DELETE /lists/{list_id}/subscriptions/{recipient_id}` with `courier.ListUnsubscribe` method
  - Usage Examples in v2/examples/lists

- Support for [Brands API](https://docs.courier.com/reference/brands-api) by @tk26
  - `GET /brands` with `courier.GetBrands` method
  - `GET /brands/{brand_id}` with `courier.GetBrand` method
  - `POST /brands/{brand_id}` with `courier.PostBrand` method
  - `PUT /brands/{brand_id}` with `courier.PutBrand` method
  - `DELETE /brands/{brand_id}` with `courier.DeleteBrand` method
  - Usage Examples in v2/examples/brands

## [v2.1.0] - 2021-04-09

### Added

- Support for [Automation API](https://docs.courier.com/reference/automation-api) by @tk26
  - `POST /automations/invoke` with `courier.InvokeAutomation` method
  - `POST /automations/{template_id}/invoke` with `courier.InvokeAutomationTemplate` method
  - Usage Examples in v2/examples/automations
- Changed default base URL to `https://api.courier.com`

## [v2.0.2] - 2020-09-22

- Adds brand and override support to SendBody struct

## [v2.0.1] - 2020-09-17

- Fixes issues where a body was being set on GET requests

## [v2.0.0] - 2020-06-24

- `courier.CourierClient` has been renamed to `courier.CreateClient`
- `courier.CreateClient` now takes a `nil` second parameter if you wish to use our default API URL, not `""`
- Our underlying API communication functions are now exposed via `client.API`
- renamed `SendRequest` to `SendBody` and modified the struct
- `Send` now takes one `body interface{}` argument rather than separate `profile` and `data` arguments
- added `SendMap`
- `GetProfile` now requires a `context` argument
- `GetProfile` now returns a `map[string]json.RawMessage{}` instead of hydrating a struct reference
- removed `MergeProfile`; use `client.API.SendRequestWithBytes(context.Background(), "PUT", "/profiles/"+recipientID, profile)` instead
- removed `UpdateProfile`; use `client.API.SendRequestWithBytes(context.Background(), "PUT", "/profiles/"+recipientID, profile)` instead
