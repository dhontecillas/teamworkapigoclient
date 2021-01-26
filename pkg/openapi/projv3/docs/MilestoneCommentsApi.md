# \MilestoneCommentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3MilestonesIdCommentsJson**](MilestoneCommentsApi.md#GETProjectsApiV3MilestonesIdCommentsJson) | **Get** /projects/api/v3/milestones/:id/comments.json | Get a list of comments for a milestone



## GETProjectsApiV3MilestonesIdCommentsJson

> CommentCommentsResponse GETProjectsApiV3MilestonesIdCommentsJson(ctx).Execute()

Get a list of comments for a milestone

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestoneCommentsApi.GETProjectsApiV3MilestonesIdCommentsJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestoneCommentsApi.GETProjectsApiV3MilestonesIdCommentsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3MilestonesIdCommentsJson`: CommentCommentsResponse
    fmt.Fprintf(os.Stdout, "Response from `MilestoneCommentsApi.GETProjectsApiV3MilestonesIdCommentsJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MilestonesIdCommentsJsonRequest struct via the builder pattern


### Return type

[**CommentCommentsResponse**](comment.CommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

