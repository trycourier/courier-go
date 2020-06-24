# Change Log

All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [v2.0.0] - 2020-06-24

* `courier.CourierClient` has been renamed to `courier.CreateClient`
* `courier.CreateClient` now takes a `nil` second parameter if you wish to use our default API URL, not `""`
* Our underlying API communication functions are now exposed via `client.API`
* renamed `SendRequest` to `SendBody` and modified the struct
* `Send` now takes one `body interface{}` argument rather than separate `profile` and `data` arguments
* added `SendMap`
* `GetProfile` now requires a `context` argument 
* `GetProfile` now returns a `map[string]json.RawMessage{}` instead of hydrating a struct reference
* renamed `MergeProfile` to `MergeProfileBytes`
* `MergeProfileBytes` now requires a `context` argument
* renamed `UpdateProfile` to `UpdateProfileBytes`
* `UpdateProfileBytes` now requires a `context` argument
