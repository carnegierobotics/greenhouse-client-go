# greenhouse-client-go

[![Go Reference](https://pkg.go.dev/badge/github.com/carnegierobotics/greenhouse-client-go.svg)](https://pkg.go.dev/github.com/carnegierobotics/greenhouse-client-go)

This library provides a Go client for the Greenhouse [Harvest](https://developers.greenhouse.io/harvest.html#introduction) and [Job Board](https://developers.greenhouse.io/job-board.html)\* APIs.

\* The Job Board API has not yet been added to this library. 

## Installation
```
# Go Modules
require github.com/carnegierobotics/greenhouse-client-go/greenhouse
```

## Usage
The functions included in this package are more or less 1:1 with the API specification. Below are some snippets demonstrating how to use this library to interact with the Greenhouse APIs. 
### Initializing the client
This library uses [Resty](https://github.com/go-resty/resty) to build the client. Additionally, all functions are [context](https://pkg.go.dev/context)-aware, so be sure to include one in your calls.
```go
import (
  "github.com/carnegierobotics/greenhouse-client-go/greenhouse"
)
client := greenhouse.Client{
  BaseUrl:      "https://boards-api.greenhouse.io", // Harvest API URL
  Token:        "abc123",                           // Harvest API token
  OnBehalfOf:   "12345",                            // On-Behalf-Of user ID
  RetryCount:   5,                                  // Number of retries per failed API call 
  RetryWait:    5,                                  // Minimum time to wait between retries
  RetryMaxWait: 30,                                 // Maximum time to wait between retries
}
client.BuildResty()
```
### Initialize a context
```go
import (
  "context"
)
ctx := context.TODO()
```
### Examples
Below are some examples of how this package can be used. Note that not all functions are shown here; these are merely examples to get you started.
#### Get a candidate's activity feed
```go
feed, err := greenhouse.GetActivityFeed(&client, ctx, candidateId)
```
#### Get all applications
```go
list, err := greenhouse.GetAllApplications(&client, ctx)
```
#### Add an application to a candidate
```go
var applicationObj greenhouse.Application
applicationId, err := greenhouse.AddApplicationToCandidate(&client, ctx, candidateId, &applicationObj)
```
#### Update an application
```go
var applicationObj greenhouse.Application
err := greenhouse.UpdateApplication(&client, ctx, applicationId, &applicationObj)
```
#### Advance an application
```go
err := greenhouse.AdvanceApplication(&client, ctx, applicationId, fromStageId)
```
#### Move an application to a different job
```go
err := greenhouse.MoveApplicationDifferentJob(&client, ctx, applicationId, newJobId, newStageId)
```
#### Move an application between stages in a job
```go
err := greenhouse.MoveApplicationSameJob(&client, ctx, applicationId, fromStageId, toStageId)
```
#### Add an attachment to an application
```go
var attachmentObj greenhouse.Attachment
err := greenhouse.AddAttachmentToApplication(&client, ctx, applicationId, &attachmentObj)
```
#### Hire an application
```go
var hireObj greenhouse.ApplicationHire
err := greenhouse.HireApplication(&client, ctx, applicationId, &hireObj)
```
#### Reject an application
```go
var rejectObj greenhouse.ApplicationReject
err := greenhouse.RejectApplication(&client, ctx, applicationId, &rejectObj)
```
#### Update a rejection reason
```go
err := greenhouse.UpdateRejectionReason(&client, ctx, applicationId, newReasonId)
```
#### Unreject an application
```go
err := greenhouse.UnrejectApplication(&client, ctx, applicationId)
```

## Links
[Greenhouse Harvest API](https://developers.greenhouse.io/harvest.html#introduction)
[Greenhouse Job Board API](https://developers.greenhouse.io/job-board.html)
