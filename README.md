# greenhouse-client-go
This library provides a Go client for the Greenhouse [Harvest](https://developers.greenhouse.io/harvest.html#introduction) and [Job Board](https://developers.greenhouse.io/job-board.html)\* APIs.
\* The Job Board API has not yet been added to this library. 

## Installation
```
# Go Modules
require github.com/carnegierobotics/greenhouse-client-go/greenhouse
```
## Usage
Below are some snippets demonstrating how to use this library to interact with the Greenhouse APIs. The methods included here are more or less 1:1 with the API specification.
### Initializing the client
This library uses [Resty](https://github.com/go-resty/resty) to build the client. Additionally, all functions are [context](https://pkg.go.dev/context)-aware, so be sure to include one in your calls.
```
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
```
import (
  "context"
)
ctx := context.TODO()
```
### Get a candidate's activity feed
```
feed, err := greenhouse.GetActivityFeed(&client, ctx, candidateId)
```
### Get all applications
```
list, err := greenhouse.GetAllApplications(&client, ctx)
```
### Get a single application by ID
```
application, err := greenhouse.GetApplication(&client, ctx, applicationId)
```
### Delete an application
```
err := greenhouse.DeleteApplication(&client, ctx, applicationId)
```
### Add an application to a candidate
```
var applicationObj greenhouse.Application
applicationId, err := greenhouse.AddApplicationToCandidate(&client, ctx, candidateId, &applicationObj)
```
### Update an application
```
var applicationObj greenhouse.Application
err := greenhouse.UpdateApplication(&client, ctx, applicationId, &applicationObj)
```
### Advance an application
```
err := greenhouse.AdvanceApplication(&client, ctx, applicationId, fromStageId)
```
### Move an application to a different job
```
err := greenhouse.MoveApplicationDifferentJob(&client, ctx, applicationId, newJobId, newStageId)
```
### Move an application between stages in a job
```
err := greenhouse.MoveApplicationSameJob(&client, ctx, applicationId, fromStageId, toStageId)
```
### Add an attachment to an application
```
var attachmentObj greenhouse.Attachment
err := greenhouse.AddAttachmentToApplication(&client, ctx, applicationId, &attachmentObj)
```
### Hire an application
```
var hireObj greenhouse.ApplicationHire
err := greenhouse.HireApplication(&client, ctx, applicationId, &hireObj)
```
### Reject an application
```
var rejectObj greenhouse.ApplicationReject
err := greenhouse.RejectApplication(&client, ctx, applicationId, &rejectObj)
```
### Update a rejection reason
```
err := greenhouse.UpdateRejectionReason(&client, ctx, applicationId, newReasonId)
```
### Unreject an application
```
err := greenhouse.UnrejectApplication(&client, ctx, applicationId)
```
### Get all candidates
```
list, err := greenhouse.GetAllCandidates(&client, ctx)
```
### Get a single candidate by ID
```
candidate, err := greenhouse.GetCandidate(&client, ctx, candidateId)
```
### Get a single candidate by Name
```
candidate, err := greenhouse.GetCandidateByName(&client, ctx, firstName, lastName)
```
