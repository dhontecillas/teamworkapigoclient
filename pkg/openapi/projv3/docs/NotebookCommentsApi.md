# \NotebookCommentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3NotebooksIdCommentsJson**](NotebookCommentsApi.md#GETProjectsApiV3NotebooksIdCommentsJson) | **Get** /projects/api/v3/notebooks/:id/comments.json | Get a list of comments for a notebook



## GETProjectsApiV3NotebooksIdCommentsJson

> CommentCommentsResponse GETProjectsApiV3NotebooksIdCommentsJson(ctx).Execute()

Get a list of comments for a notebook

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
    resp, r, err := api_client.NotebookCommentsApi.GETProjectsApiV3NotebooksIdCommentsJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebookCommentsApi.GETProjectsApiV3NotebooksIdCommentsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksIdCommentsJson`: CommentCommentsResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebookCommentsApi.GETProjectsApiV3NotebooksIdCommentsJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksIdCommentsJsonRequest struct via the builder pattern


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

