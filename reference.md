# Reference
<details><summary><code>client.Send(request) -> *v3.SendMessageResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use the send API to send a message to one or more recipients.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Send(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**message:** `*v3.Message` — Defines the message to be delivered
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Audiences
<details><summary><code>client.Audiences.Get(AudienceId) -> *v3.Audience</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the specified audience by id.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Audiences.Get(
        context.TODO(),
        "audience_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**audienceId:** `string` — A unique identifier representing the audience_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.Update(AudienceId, request) -> *v3.AudienceUpdateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates audience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Audiences.Update(
        context.TODO(),
        "audience_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**audienceId:** `string` — A unique identifier representing the audience id
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the audience
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A description of the audience
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*v3.Filter` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.Delete(AudienceId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the specified audience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Audiences.Delete(
        context.TODO(),
        "audience_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**audienceId:** `string` — A unique identifier representing the audience id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.ListMembers(AudienceId) -> *v3.AudienceMemberListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of members of an audience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Audiences.ListMembers(
        context.TODO(),
        "audience_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**audienceId:** `string` — A unique identifier representing the audience id
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — A unique identifier that allows for fetching the next set of members
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.ListAudiences() -> *v3.AudienceListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the audiences associated with the authorization token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Audiences.ListAudiences(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — A unique identifier that allows for fetching the next set of audiences
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AuditEvents
<details><summary><code>client.AuditEvents.List() -> *v3.ListAuditEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetch the list of audit events
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.AuditEvents.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — A unique identifier that allows for fetching the next set of audit events.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AuditEvents.Get(AuditEventId) -> *v3.AuditEvent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetch a specific audit event by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.AuditEvents.Get(
        context.TODO(),
        "audit-event-id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**auditEventId:** `string` — A unique identifier associated with the audit event you wish to retrieve
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AuthTokens
<details><summary><code>client.AuthTokens.IssueToken(request) -> *v3.IssueTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a new access token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.AuthTokens.IssueToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**scope:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expiresIn:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Automations
<details><summary><code>client.Automations.InvokeAutomationTemplate(TemplateId, request) -> *v3.AutomationInvokeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Invoke an automation run from an automation template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Automations.InvokeAutomationTemplate(
        context.TODO(),
        "templateId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateId:** `string` — A unique identifier representing the automation template to be invoked. This could be the Automation Template ID or the Automation Template Alias.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v3.AutomationInvokeParams` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.InvokeAdHocAutomation(request) -> *v3.AutomationInvokeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Invoke an ad hoc automation run. This endpoint accepts a JSON payload with a series of automation steps. For information about what steps are available, checkout the ad hoc automation guide [here](https://www.courier.com/docs/automations/steps/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Automations.InvokeAdHocAutomation(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*v3.AutomationAdHocInvokeParams` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Brands
<details><summary><code>client.Brands.Create(request) -> *v3.Brand</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Brands.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*v3.BrandParameters` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Brands.Get(BrandId) -> *v3.Brand</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetch a specific brand by brand ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Brands.Get(
        context.TODO(),
        "brand_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**brandId:** `string` — A unique identifier associated with the brand you wish to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Brands.List() -> *v3.BrandsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the list of brands.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Brands.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — A unique identifier that allows for fetching the next set of brands.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Brands.Delete(BrandId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a brand by brand ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Brands.Delete(
        context.TODO(),
        "brand_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**brandId:** `string` — A unique identifier associated with the brand you wish to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Brands.Replace(BrandId, request) -> *v3.Brand</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replace an existing brand with the supplied values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Brands.Replace(
        context.TODO(),
        "brand_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**brandId:** `string` — A unique identifier associated with the brand you wish to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the brand.
    
</dd>
</dl>

<dl>
<dd>

**settings:** `*v3.BrandSettings` 
    
</dd>
</dl>

<dl>
<dd>

**snippets:** `*v3.BrandSnippets` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Bulk
<details><summary><code>client.Bulk.CreateJob(request) -> *v3.BulkCreateJobResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Bulk.CreateJob(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**message:** `*v3.InboundBulkMessage` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bulk.IngestUsers(JobId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Ingest user data into a Bulk Job
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Bulk.IngestUsers(
        context.TODO(),
        "job_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**jobId:** `string` — A unique identifier representing the bulk job
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v3.BulkIngestUsersParams` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bulk.RunJob(JobId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Run a bulk job
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Bulk.RunJob(
        context.TODO(),
        "job_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**jobId:** `string` — A unique identifier representing the bulk job
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bulk.GetJob(JobId) -> *v3.BulkGetJobResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a bulk job
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Bulk.GetJob(
        context.TODO(),
        "job_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**jobId:** `string` — A unique identifier representing the bulk job
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bulk.GetUsers(JobId) -> *v3.BulkGetJobUsersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get Bulk Job Users
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Bulk.GetUsers(
        context.TODO(),
        "job_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**jobId:** `string` — A unique identifier representing the bulk job
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — A unique identifier that allows for fetching the next set of users added to the bulk job
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Inbound
<details><summary><code>client.Inbound.Track(request) -> *v3.TrackAcceptedResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Inbound.Track(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*v3.InboundTrackEvent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Lists
<details><summary><code>client.Lists.List() -> *v3.ListGetAllResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all of the lists, with the ability to filter based on a pattern.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Lists.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — A unique identifier that allows for fetching the next page of lists.
    
</dd>
</dl>

<dl>
<dd>

**pattern:** `*string` — "A pattern used to filter the list items returned. Pattern types supported: exact match on `list_id` or a pattern of one or more pattern parts. you may replace a part with either: `*` to match all parts in that position, or `**` to signify a wildcard `endsWith` pattern match."
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.Get(ListId) -> *v3.List</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list based on the list ID provided.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Lists.Get(
        context.TODO(),
        "list_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listId:** `string` — A unique identifier representing the list you wish to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.Update(ListId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create or replace an existing list with the supplied values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Lists.Update(
        context.TODO(),
        "list_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listId:** `string` — A unique identifier representing the list you wish to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v3.ListPutParams` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.Delete(ListId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a list by list ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Lists.Delete(
        context.TODO(),
        "list_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listId:** `string` — A unique identifier representing the list you wish to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.Restore(ListId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Restore a previously deleted list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Lists.Restore(
        context.TODO(),
        "list_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listId:** `string` — A unique identifier representing the list you wish to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetSubscribers(ListId) -> *v3.ListGetSubscriptionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the list's subscriptions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Lists.GetSubscribers(
        context.TODO(),
        "list_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listId:** `string` — A unique identifier representing the list you wish to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — A unique identifier that allows for fetching the next set of list subscriptions
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.UpdateSubscribers(ListId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Subscribes the users to the list, overwriting existing subscriptions. If the list does not exist, it will be automatically created.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Lists.UpdateSubscribers(
        context.TODO(),
        "list_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listId:** `string` — A unique identifier representing the list you wish to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**recipients:** `[]*v3.PutSubscriptionsRecipient` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.AddSubscribers(ListId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Subscribes additional users to the list, without modifying existing subscriptions. If the list does not exist, it will be automatically created.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Lists.AddSubscribers(
        context.TODO(),
        "list_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listId:** `string` — A unique identifier representing the list you wish to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**recipients:** `[]*v3.PutSubscriptionsRecipient` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.Subscribe(ListId, UserId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Subscribe a user to an existing list (note: if the List does not exist, it will be automatically created).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Lists.Subscribe(
        context.TODO(),
        "list_id",
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listId:** `string` — A unique identifier representing the list you wish to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**userId:** `string` — A unique identifier representing the recipient associated with the list
    
</dd>
</dl>

<dl>
<dd>

**preferences:** `*v3.RecipientPreferences` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.Unsubscribe(ListId, UserId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a subscription to a list by list ID and user ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Lists.Unsubscribe(
        context.TODO(),
        "list_id",
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listId:** `string` — A unique identifier representing the list you wish to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**userId:** `string` — A unique identifier representing the recipient associated with the list
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Messages
<details><summary><code>client.Messages.List() -> *v3.ListMessagesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetch the statuses of messages you've previously sent.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Messages.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**archived:** `*bool` — A boolean value that indicates whether archived messages should be included in the response.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — A unique identifier that allows for fetching the next set of messages.
    
</dd>
</dl>

<dl>
<dd>

**event:** `*string` — A unique identifier representing the event that was used to send the event.
    
</dd>
</dl>

<dl>
<dd>

**list:** `*string` — A unique identifier representing the list the message was sent to.
    
</dd>
</dl>

<dl>
<dd>

**messageId:** `*string` — A unique identifier representing the message_id returned from either /send or /send/list.
    
</dd>
</dl>

<dl>
<dd>

**notification:** `*string` — A unique identifier representing the notification that was used to send the event.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — The key assocated to the provider you want to filter on. E.g., sendgrid, inbox, twilio, slack, msteams, etc. Allows multiple values to be set in query parameters.
    
</dd>
</dl>

<dl>
<dd>

**recipient:** `*string` — A unique identifier representing the recipient associated with the requested profile.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` — An indicator of the current status of the message. Allows multiple values to be set in query parameters.
    
</dd>
</dl>

<dl>
<dd>

**tag:** `*string` — A tag placed in the metadata.tags during a notification send. Allows multiple values to be set in query parameters.
    
</dd>
</dl>

<dl>
<dd>

**tags:** `*string` — A comma delimited list of 'tags'. Messages will be returned if they match any of the tags passed in.
    
</dd>
</dl>

<dl>
<dd>

**tenantId:** `*string` — Messages sent with the context of a Tenant
    
</dd>
</dl>

<dl>
<dd>

**enqueuedAfter:** `*string` — The enqueued datetime of a message to filter out messages received before.
    
</dd>
</dl>

<dl>
<dd>

**traceId:** `*string` — The unique identifier used to trace the requests
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Messages.Get(MessageId) -> *v3.MessageDetailsExtended</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetch the status of a message you've previously sent.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Messages.Get(
        context.TODO(),
        "message_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**messageId:** `string` — A unique identifier associated with the message you wish to retrieve (results from a send).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Messages.Cancel(MessageId) -> *v3.MessageDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancel a message that is currently in the process of being delivered. A well-formatted API call to the cancel message API will return either `200` status code for a successful cancellation or `409` status code for an unsuccessful cancellation. Both cases will include the actual message record in the response body (see details below).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Messages.Cancel(
        context.TODO(),
        "message_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**messageId:** `string` — A unique identifier representing the message ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Messages.GetHistory(MessageId) -> *v3.MessageHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetch the array of events of a message you've previously sent.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Messages.GetHistory(
        context.TODO(),
        "message_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**messageId:** `string` — A unique identifier representing the message ID
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — A supported Message History type that will filter the events returned.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Messages.GetContent(MessageId) -> *v3.RenderOutputResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Messages.GetContent(
        context.TODO(),
        "message_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**messageId:** `string` — A unique identifier associated with the message you wish to retrieve (results from a send).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Messages.Archive(RequestId) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Messages.Archive(
        context.TODO(),
        "request_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**requestId:** `string` — A unique identifier representing the request ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Notifications
<details><summary><code>client.Notifications.List() -> *v3.NotificationListResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Notifications.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**notes:** `*bool` — Retrieve the notes from the Notification template settings.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.GetContent(Id) -> *v3.NotificationGetContentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Notifications.GetContent(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.GetDraftContent(Id) -> *v3.NotificationGetContentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Notifications.GetDraftContent(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.GetSubmissionChecks(Id, SubmissionId) -> *v3.SubmissionChecksGetResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Notifications.GetSubmissionChecks(
        context.TODO(),
        "id",
        "submissionId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**submissionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.ReplaceSubmissionChecks(Id, SubmissionId, request) -> *v3.SubmissionChecksReplaceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Notifications.ReplaceSubmissionChecks(
        context.TODO(),
        "id",
        "submissionId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**submissionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**checks:** `[]*v3.BaseCheck` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.CancelSubmission(Id, SubmissionId) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Notifications.CancelSubmission(
        context.TODO(),
        "id",
        "submissionId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**submissionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Profiles
<details><summary><code>client.Profiles.Get(UserId) -> *v3.ProfileGetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the specified user profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Profiles.Get(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier representing the user associated with the requested profile.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Profiles.Create(UserId, request) -> *v3.MergeProfileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Merge the supplied values with an existing profile or create a new profile if one doesn't already exist.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Profiles.Create(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier representing the user associated with the requested profile.
    
</dd>
</dl>

<dl>
<dd>

**profile:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Profiles.Replace(UserId, request) -> *v3.ReplaceProfileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

When using `PUT`, be sure to include all the key-value pairs required by the recipient's profile. 
Any key-value pairs that exist in the profile but fail to be included in the `PUT` request will be 
removed from the profile. Remember, a `PUT` update is a full replacement of the data. For partial updates, 
use the [Patch](https://www.courier.com/docs/reference/profiles/patch/) request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Profiles.Replace(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier representing the user associated with the requested profile.
    
</dd>
</dl>

<dl>
<dd>

**profile:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Profiles.MergeProfile(UserId, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Profiles.MergeProfile(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier representing the user associated with the requested profile.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v3.ProfileUpdateRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Profiles.Delete(UserId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the specified user profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Profiles.Delete(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier representing the user associated with the requested profile.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Profiles.GetListSubscriptions(UserId) -> *v3.GetListSubscriptionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the subscribed lists for a specified user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Profiles.GetListSubscriptions(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier representing the user associated with the requested profile.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — A unique identifier that allows for fetching the next set of message statuses.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Profiles.SubscribeToLists(UserId, request) -> *v3.SubscribeToListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Subscribes the given user to one or more lists. If the list does not exist, it will be created.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Profiles.SubscribeToLists(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier representing the user associated with the requested profile.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v3.SubscribeToListsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Profiles.DeleteListSubscription(UserId) -> *v3.DeleteListSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes all list subscriptions for given user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Profiles.DeleteListSubscription(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier representing the user associated with the requested profile.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Templates
<details><summary><code>client.Templates.List() -> *v3.ListTemplatesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of notification templates
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Notifications.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — A unique identifier that allows for fetching the next set of templates
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tenants
<details><summary><code>client.Tenants.CreateOrReplace(TenantId, request) -> *v3.Tenant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.CreateOrReplace(
        context.TODO(),
        "tenant_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantId:** `string` — A unique identifier representing the tenant to be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Name of the tenant.
    
</dd>
</dl>

<dl>
<dd>

**parentTenantId:** `*string` — Tenant's parent id (if any).
    
</dd>
</dl>

<dl>
<dd>

**defaultPreferences:** `*v3.DefaultPreferences` — Defines the preferences used for the tenant when the user hasn't specified their own.
    
</dd>
</dl>

<dl>
<dd>

**properties:** `map[string]any` — Arbitrary properties accessible to a template.
    
</dd>
</dl>

<dl>
<dd>

**userProfile:** `map[string]any` — A user profile object merged with user profile on send.
    
</dd>
</dl>

<dl>
<dd>

**brandId:** `*string` — Brand to be used for the account when one is not specified by the send call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Get(TenantId) -> *v3.Tenant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Get(
        context.TODO(),
        "tenant_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantId:** `string` — A unique identifier representing the tenant to be returned.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.List() -> *v3.TenantListResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**parentTenantId:** `*string` — Filter the list of tenants by parent_id
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The number of tenants to return 
(defaults to 20, maximum value of 100)
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Continue the pagination with the next cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Delete(TenantId) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Delete(
        context.TODO(),
        "tenant_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantId:** `string` — Id of the tenant to be deleted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.GetUsersByTenant(TenantId) -> *v3.ListUsersForTenantResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.GetUsersByTenant(
        context.TODO(),
        "tenant_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantId:** `string` — Id of the tenant for user membership.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The number of accounts to return 
(defaults to 20, maximum value of 100)
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Continue the pagination with the next cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.CreateOrReplaceDefaultPreferencesForTopic(TenantId, TopicId, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.CreateOrReplaceDefaultPreferencesForTopic(
        context.TODO(),
        "tenantABC",
        "HB529N49MD4D5PMX9WR5P4JH78NA",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantId:** `string` — Id of the tenant to update the default preferences for.
    
</dd>
</dl>

<dl>
<dd>

**topicId:** `string` — Id fo the susbcription topic you want to have a default preference for.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*v3.SubscriptionTopicNew` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.RemoveDefaultPreferencesForTopic(TenantId, TopicId) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.RemoveDefaultPreferencesForTopic(
        context.TODO(),
        "tenant_id",
        "topic_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantId:** `string` — Id of the tenant to update the default preferences for.
    
</dd>
</dl>

<dl>
<dd>

**topicId:** `string` — Id fo the susbcription topic you want to have a default preference for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.GetTemplateByTenant(TenantId, TemplateId) -> *v3.GetTemplateByTenantResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.GetTemplateByTenant(
        context.TODO(),
        "tenant_id",
        "template_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantId:** `string` — Id of the tenant for which to retrieve the template.
    
</dd>
</dl>

<dl>
<dd>

**templateId:** `string` — Id of the template to be retrieved.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.GetTemplateListByTenant(TenantId) -> *v3.ListTemplatesByTenantResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.GetTemplateListByTenant(
        context.TODO(),
        "tenant_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantId:** `string` — Id of the tenant for which to retrieve the templates.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The number of templates to return (defaults to 20, maximum value of 100)
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Continue the pagination with the next cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Translations
<details><summary><code>client.Translations.Get(Domain, Locale) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get translations by locale
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Translations.Get(
        context.TODO(),
        "domain",
        "locale",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domain:** `string` — The domain you want to retrieve translations for. Only `default` is supported at the moment
    
</dd>
</dl>

<dl>
<dd>

**locale:** `string` — The locale you want to retrieve the translations for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Translations.Update(Domain, Locale, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a translation
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Translations.Update(
        context.TODO(),
        "domain",
        "locale",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domain:** `string` — The domain you want to retrieve translations for. Only `default` is supported at the moment
    
</dd>
</dl>

<dl>
<dd>

**locale:** `string` — The locale you want to retrieve the translations for
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — .po file translation content
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Preferences
<details><summary><code>client.Users.Preferences.List(UserId) -> *users.UserPreferencesListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetch all user preferences.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Preferences.List(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier associated with the user whose preferences you wish to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**tenantId:** `*string` — Query the preferences of a user for this specific tenant context.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Preferences.Get(UserId, TopicId) -> *users.UserPreferencesGetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetch user preferences for a specific subscription topic.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Preferences.Get(
        context.TODO(),
        "user_id",
        "topic_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier associated with the user whose preferences you wish to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**topicId:** `string` — A unique identifier associated with a subscription topic.
    
</dd>
</dl>

<dl>
<dd>

**tenantId:** `*string` — Query the preferences of a user for this specific tenant context.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Preferences.Update(UserId, TopicId, request) -> *users.UserPreferencesUpdateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update or Create user preferences for a specific subscription topic.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Preferences.Update(
        context.TODO(),
        "abc-123",
        "74Q4QGFBEX481DP6JRPMV751H4XT",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — A unique identifier associated with the user whose preferences you wish to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**topicId:** `string` — A unique identifier associated with a subscription topic.
    
</dd>
</dl>

<dl>
<dd>

**tenantId:** `*string` — Update the preferences of a user for this specific tenant context.
    
</dd>
</dl>

<dl>
<dd>

**topic:** `*users.TopicPreferenceUpdate` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Tenants
<details><summary><code>client.Users.Tenants.AddMultple(UserId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint is used to add a user to
multiple tenants in one call.
A custom profile can also be supplied for each tenant. 
This profile will be merged with the user's main 
profile when sending to the user with that tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tenants.AddMultple(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — The user's ID. This can be any uniquely identifiable string.
    
</dd>
</dl>

<dl>
<dd>

**tenants:** `[]*v3.UserTenantAssociation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Tenants.Add(UserId, TenantId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint is used to add a single tenant.

A custom profile can also be supplied with the tenant. 
This profile will be merged with the user's main profile 
when sending to the user with that tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tenants.Add(
        context.TODO(),
        "user_id",
        "tenant_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — Id of the user to be added to the supplied tenant.
    
</dd>
</dl>

<dl>
<dd>

**tenantId:** `string` — Id of the tenant the user should be added to.
    
</dd>
</dl>

<dl>
<dd>

**profile:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Tenants.RemoveAll(UserId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes a user from any tenants they may have been associated with.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tenants.RemoveAll(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — Id of the user to be removed from the supplied tenant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Tenants.Remove(UserId, TenantId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes a user from the supplied tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tenants.Remove(
        context.TODO(),
        "user_id",
        "tenant_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — Id of the user to be removed from the supplied tenant.
    
</dd>
</dl>

<dl>
<dd>

**tenantId:** `string` — Id of the tenant the user should be removed from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Tenants.List(UserId) -> *users.ListTenantsForUserResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of user tenant associations.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tenants.List(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — Id of the user to retrieve all associated tenants for.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The number of accounts to return 
(defaults to 20, maximum value of 100)
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Continue the pagination with the next cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Tokens
<details><summary><code>client.Users.Tokens.AddMultiple(UserId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds multiple tokens to a user and overwrites matching existing tokens.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tokens.AddMultiple(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — The user's ID. This can be any uniquely identifiable string.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Tokens.Add(UserId, Token, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds a single token to a user and overwrites a matching existing token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tokens.Add(
        context.TODO(),
        "user_id",
        "token",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — The user's ID. This can be any uniquely identifiable string.
    
</dd>
</dl>

<dl>
<dd>

**token:** `string` — The full token string.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*users.UserToken` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Tokens.Update(UserId, Token, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Apply a JSON Patch (RFC 6902) to the specified token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tokens.Update(
        context.TODO(),
        "user_id",
        "token",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — The user's ID. This can be any uniquely identifiable string.
    
</dd>
</dl>

<dl>
<dd>

**token:** `string` — The full token string.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*users.PatchUserTokenOpts` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Tokens.Get(UserId, Token) -> *users.GetUserTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get single token available for a `:token`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tokens.Get(
        context.TODO(),
        "user_id",
        "token",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — The user's ID. This can be any uniquely identifiable string.
    
</dd>
</dl>

<dl>
<dd>

**token:** `string` — The full token string.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Tokens.List(UserId) -> users.GetAllTokensResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets all tokens available for a :user_id
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tokens.List(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — The user's ID. This can be any uniquely identifiable string.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Tokens.Delete(UserId, Token) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Tokens.Delete(
        context.TODO(),
        "user_id",
        "token",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — The user's ID. This can be any uniquely identifiable string.
    
</dd>
</dl>

<dl>
<dd>

**token:** `string` — The full token string.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
