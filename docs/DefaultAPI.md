# \DefaultAPI

All URIs are relative to *https://api.samsara.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAddress**](DefaultAPI.md#CreateAddress) | **Post** /addresses | Create an address
[**CreateContact**](DefaultAPI.md#CreateContact) | **Post** /contacts | Create a contact
[**CreateDriver**](DefaultAPI.md#CreateDriver) | **Post** /fleet/drivers | Create a driver
[**CreateTag**](DefaultAPI.md#CreateTag) | **Post** /tags | Create a tag
[**CreateUser**](DefaultAPI.md#CreateUser) | **Post** /users | Create a user
[**DeleteAddress**](DefaultAPI.md#DeleteAddress) | **Delete** /addresses/{id} | Delete an address
[**DeleteContact**](DefaultAPI.md#DeleteContact) | **Delete** /contacts/{id} | Delete a contact
[**DeleteTag**](DefaultAPI.md#DeleteTag) | **Delete** /tags/{id} | Delete a tag
[**DeleteUser**](DefaultAPI.md#DeleteUser) | **Delete** /users/{id} | Delete a user
[**GetAddress**](DefaultAPI.md#GetAddress) | **Get** /addresses/{id} | Retrieve an address
[**GetContact**](DefaultAPI.md#GetContact) | **Get** /contacts/{id} | Retrieve a contact
[**GetDriver**](DefaultAPI.md#GetDriver) | **Get** /fleet/drivers/{id} | Retrieve a driver
[**GetTag**](DefaultAPI.md#GetTag) | **Get** /tags/{id} | Retrieve a tag
[**GetUser**](DefaultAPI.md#GetUser) | **Get** /users/{id} | Retrieve a user
[**GetVehicle**](DefaultAPI.md#GetVehicle) | **Get** /fleet/vehicles/{id} | Retrieve a vehicle
[**GetVehicleLocations**](DefaultAPI.md#GetVehicleLocations) | **Get** /fleet/vehicles/locations | Get most recent vehicle locations
[**GetVehicleLocationsFeed**](DefaultAPI.md#GetVehicleLocationsFeed) | **Get** /fleet/vehicles/locations/feed | Follow a feed of vehicle locations
[**GetVehicleLocationsHistory**](DefaultAPI.md#GetVehicleLocationsHistory) | **Get** /fleet/vehicles/locations/history | Get historical vehicle locations
[**GetVehicleStats**](DefaultAPI.md#GetVehicleStats) | **Get** /fleet/vehicles/stats | List most recent vehicle stats
[**GetVehicleStatsFeed**](DefaultAPI.md#GetVehicleStatsFeed) | **Get** /fleet/vehicles/stats/feed | Follow a feed of vehicle stats
[**GetVehicleStatsHistory**](DefaultAPI.md#GetVehicleStatsHistory) | **Get** /fleet/vehicles/stats/history | Get historical vehicle stats
[**ListAddresses**](DefaultAPI.md#ListAddresses) | **Get** /addresses | List all addresses
[**ListContacts**](DefaultAPI.md#ListContacts) | **Get** /contacts | List all contacts
[**ListDrivers**](DefaultAPI.md#ListDrivers) | **Get** /fleet/drivers | List all drivers
[**ListTags**](DefaultAPI.md#ListTags) | **Get** /tags | List all tags
[**ListUserRoles**](DefaultAPI.md#ListUserRoles) | **Get** /user-roles | List all user roles
[**ListUsers**](DefaultAPI.md#ListUsers) | **Get** /users | List all users
[**ListVehicles**](DefaultAPI.md#ListVehicles) | **Get** /fleet/vehicles | List all vehicles
[**ReplaceTag**](DefaultAPI.md#ReplaceTag) | **Put** /tags/{id} | Update a tag
[**UpdateAddress**](DefaultAPI.md#UpdateAddress) | **Patch** /addresses/{id} | Update an address
[**UpdateContact**](DefaultAPI.md#UpdateContact) | **Patch** /contacts/{id} | Update a contact
[**UpdateDriver**](DefaultAPI.md#UpdateDriver) | **Patch** /fleet/drivers/{id} | Update a driver
[**UpdateUser**](DefaultAPI.md#UpdateUser) | **Patch** /users/{id} | Update a user
[**UpdateVehicle**](DefaultAPI.md#UpdateVehicle) | **Patch** /fleet/vehicles/{id} | Update a vehicle



## CreateAddress

> AddressResponse CreateAddress(ctx).Address(address).Execute()

Create an address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	address := *openapiclient.NewCreateAddressRequest("350 Rhode Island St, San Francisco, CA", *openapiclient.NewAddressGeofence(), "Samsara HQ") // CreateAddressRequest | The address to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateAddress(context.Background()).Address(address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAddress`: AddressResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | [**CreateAddressRequest**](CreateAddressRequest.md) | The address to create. | 

### Return type

[**AddressResponse**](AddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContact

> ContactResponse CreateContact(ctx).Contact(contact).Execute()

Create a contact



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	contact := *openapiclient.NewCreateContactRequest() // CreateContactRequest | The contact create parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateContact(context.Background()).Contact(contact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContact`: ContactResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contact** | [**CreateContactRequest**](CreateContactRequest.md) | The contact create parameters. | 

### Return type

[**ContactResponse**](ContactResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDriver

> DriverResponse CreateDriver(ctx).Driver(driver).Execute()

Create a driver



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	driver := *openapiclient.NewCreateDriverRequest("Susan Jones", "aSecurePassword1234", "SusanJones") // CreateDriverRequest | The driver to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateDriver(context.Background()).Driver(driver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateDriver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDriver`: DriverResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateDriver`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDriverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **driver** | [**CreateDriverRequest**](CreateDriverRequest.md) | The driver to create. | 

### Return type

[**DriverResponse**](DriverResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTag

> TagResponse CreateTag(ctx).Tag(tag).Execute()

Create a tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tag := *openapiclient.NewCreateTagRequest("California") // CreateTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTag(context.Background()).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTag`: TagResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**CreateTagRequest**](CreateTagRequest.md) |  | 

### Return type

[**TagResponse**](TagResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> UserResponse CreateUser(ctx).User(user).Execute()

Create a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	user := *openapiclient.NewCreateUserRequest(openapiclient.UserAuthType("default"), "user@company.com", "Bob Smith", []openapiclient.UserRoleAssignmentRequest{*openapiclient.NewUserRoleAssignmentRequest()}) // CreateUserRequest | The user to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateUser(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**CreateUserRequest**](CreateUserRequest.md) | The user to create. | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddress

> DeleteAddress(ctx, id).Execute()

Delete an address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | ID of the Address. This can either be the Samsara-provided ID or an external ID. External IDs are customer-specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: `key:value`. For example, `crmId:abc123`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteAddress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Address. This can either be the Samsara-provided ID or an external ID. External IDs are customer-specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: &#x60;key:value&#x60;. For example, &#x60;crmId:abc123&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContact

> DeleteContact(ctx, id).Execute()

Delete a contact



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique identifier for the contact.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteContact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the contact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteTag(ctx, id).Execute()

Delete a tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique identifier for the tag.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteTag(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, id).Execute()

Delete a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique identifier for the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddress

> AddressResponse GetAddress(ctx, id).Execute()

Retrieve an address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | ID of the Address. This can either be the Samsara-provided ID or an external ID. External IDs are customer-specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: `key:value`. For example, `crmId:abc123`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAddress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddress`: AddressResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Address. This can either be the Samsara-provided ID or an external ID. External IDs are customer-specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: &#x60;key:value&#x60;. For example, &#x60;crmId:abc123&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressResponse**](AddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContact

> ContactResponse GetContact(ctx, id).Execute()

Retrieve a contact



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique identifier for the contact.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetContact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContact`: ContactResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the contact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContactResponse**](ContactResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDriver

> DriverResponse GetDriver(ctx, id).Execute()

Retrieve a driver



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | ID of the driver. This can either be the Samsara-specified ID, or an external ID. External IDs are customer specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: `key:value`. For example, `payrollId:ABFS18600`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDriver(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDriver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriver`: DriverResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDriver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the driver. This can either be the Samsara-specified ID, or an external ID. External IDs are customer specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: &#x60;key:value&#x60;. For example, &#x60;payrollId:ABFS18600&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriverResponse**](DriverResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTag

> TagResponse GetTag(ctx, id).Execute()

Retrieve a tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique identifier for the tag.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTag(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTag`: TagResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagResponse**](TagResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserResponse GetUser(ctx, id).Execute()

Retrieve a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique identifier for the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicle

> VehicleResponse GetVehicle(ctx, id).Execute()

Retrieve a vehicle



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | ID of the vehicle. This can either be the Samsara-specified ID, or an external ID. External IDs are customer specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: `key:value`. For example, `maintenanceId:250020`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetVehicle(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVehicle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicle`: VehicleResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVehicle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the vehicle. This can either be the Samsara-specified ID, or an external ID. External IDs are customer specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: &#x60;key:value&#x60;. For example, &#x60;maintenanceId:250020&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VehicleResponse**](VehicleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicleLocations

> VehicleLocationsResponse GetVehicleLocations(ctx).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()

Get most recent vehicle locations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)
	tagIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of tag IDs. Example: `tagIds=1234,5678` (optional)
	vehicleIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of vehicle IDs. Example: `vehicleIds=1234,5678` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetVehicleLocations(context.Background()).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVehicleLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicleLocations`: VehicleLocationsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVehicleLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVehicleLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 
 **tagIds** | **[]string** | A filter on the data based on this comma-separated list of tag IDs. Example: &#x60;tagIds&#x3D;1234,5678&#x60; | 
 **vehicleIds** | **[]string** | A filter on the data based on this comma-separated list of vehicle IDs. Example: &#x60;vehicleIds&#x3D;1234,5678&#x60; | 

### Return type

[**VehicleLocationsResponse**](VehicleLocationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicleLocationsFeed

> VehicleLocationsListResponse GetVehicleLocationsFeed(ctx).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()

Follow a feed of vehicle locations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)
	tagIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of tag IDs. Example: `tagIds=1234,5678` (optional)
	vehicleIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of vehicle IDs. Example: `vehicleIds=1234,5678` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetVehicleLocationsFeed(context.Background()).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVehicleLocationsFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicleLocationsFeed`: VehicleLocationsListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVehicleLocationsFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVehicleLocationsFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 
 **tagIds** | **[]string** | A filter on the data based on this comma-separated list of tag IDs. Example: &#x60;tagIds&#x3D;1234,5678&#x60; | 
 **vehicleIds** | **[]string** | A filter on the data based on this comma-separated list of vehicle IDs. Example: &#x60;vehicleIds&#x3D;1234,5678&#x60; | 

### Return type

[**VehicleLocationsListResponse**](VehicleLocationsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicleLocationsHistory

> VehicleLocationsListResponse GetVehicleLocationsHistory(ctx).StartTime(startTime).EndTime(endTime).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()

Get historical vehicle locations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	startTime := time.Now() // time.Time | A start time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00).
	endTime := time.Now() // time.Time | An end time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00).
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)
	tagIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of tag IDs. Example: `tagIds=1234,5678` (optional)
	vehicleIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of vehicle IDs. Example: `vehicleIds=1234,5678` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetVehicleLocationsHistory(context.Background()).StartTime(startTime).EndTime(endTime).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVehicleLocationsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicleLocationsHistory`: VehicleLocationsListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVehicleLocationsHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVehicleLocationsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **time.Time** | A start time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00). | 
 **endTime** | **time.Time** | An end time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00). | 
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 
 **tagIds** | **[]string** | A filter on the data based on this comma-separated list of tag IDs. Example: &#x60;tagIds&#x3D;1234,5678&#x60; | 
 **vehicleIds** | **[]string** | A filter on the data based on this comma-separated list of vehicle IDs. Example: &#x60;vehicleIds&#x3D;1234,5678&#x60; | 

### Return type

[**VehicleLocationsListResponse**](VehicleLocationsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicleStats

> VehicleStatsResponse GetVehicleStats(ctx).Types(types).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()

List most recent vehicle stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	types := []string{"obdOdometerMeters"} // []string | The stat type you want this endpoint to return information on. Currently only one stat type is accepted per request.  - `engineStates`: The state of the engine (`Off`, `On`, `Idle`). - `fuelPercents`: The engine fuel level in percentage points (e.g. `99`, `50`, etc). - `obdOdometerMeters`: The odometer reading according to on-board diagnostics. If Samsara does not have diagnostic coverage for a particular vehicle, the value for this stat type will be omitted. In these cases, we recommend using `gpsOdometerMeters`. - `gpsOdometerMeters`: The odometer reading according to GPS calculations. This calculation is based off GPS distance traveled and a manual odometer offset for a given vehicle, specified by the user in Samsara's dashboard UI or through the `odometerMeters` field in the [PATCH /fleet/vehicles/{id}](#operation/updateVehicleById) endpoint. `gpsOdometerMeters` is equal to the manual offset plus the GPS distance traveled since the offset was set. The value for this stat type will be omitted if a manual offset is not provided for a given vehicle. *A manual offset can only be provided when we do not have diagnostic coverage for a particular vehicle.* - `obdEngineSeconds`: The cumulative number of seconds the engine has run according to on-board diagnostics. - `gpsDistanceMeters`: The distance the vehicle has traveled since the gateway was installed based on GPS calculations.
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)
	tagIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of tag IDs. Example: `tagIds=1234,5678` (optional)
	vehicleIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of vehicle IDs. Example: `vehicleIds=1234,5678` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetVehicleStats(context.Background()).Types(types).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVehicleStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicleStats`: VehicleStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVehicleStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVehicleStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **types** | **[]string** | The stat type you want this endpoint to return information on. Currently only one stat type is accepted per request.  - &#x60;engineStates&#x60;: The state of the engine (&#x60;Off&#x60;, &#x60;On&#x60;, &#x60;Idle&#x60;). - &#x60;fuelPercents&#x60;: The engine fuel level in percentage points (e.g. &#x60;99&#x60;, &#x60;50&#x60;, etc). - &#x60;obdOdometerMeters&#x60;: The odometer reading according to on-board diagnostics. If Samsara does not have diagnostic coverage for a particular vehicle, the value for this stat type will be omitted. In these cases, we recommend using &#x60;gpsOdometerMeters&#x60;. - &#x60;gpsOdometerMeters&#x60;: The odometer reading according to GPS calculations. This calculation is based off GPS distance traveled and a manual odometer offset for a given vehicle, specified by the user in Samsara&#39;s dashboard UI or through the &#x60;odometerMeters&#x60; field in the [PATCH /fleet/vehicles/{id}](#operation/updateVehicleById) endpoint. &#x60;gpsOdometerMeters&#x60; is equal to the manual offset plus the GPS distance traveled since the offset was set. The value for this stat type will be omitted if a manual offset is not provided for a given vehicle. *A manual offset can only be provided when we do not have diagnostic coverage for a particular vehicle.* - &#x60;obdEngineSeconds&#x60;: The cumulative number of seconds the engine has run according to on-board diagnostics. - &#x60;gpsDistanceMeters&#x60;: The distance the vehicle has traveled since the gateway was installed based on GPS calculations. | 
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 
 **tagIds** | **[]string** | A filter on the data based on this comma-separated list of tag IDs. Example: &#x60;tagIds&#x3D;1234,5678&#x60; | 
 **vehicleIds** | **[]string** | A filter on the data based on this comma-separated list of vehicle IDs. Example: &#x60;vehicleIds&#x3D;1234,5678&#x60; | 

### Return type

[**VehicleStatsResponse**](VehicleStatsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicleStatsFeed

> VehicleStatsListResponse GetVehicleStatsFeed(ctx).Types(types).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()

Follow a feed of vehicle stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	types := []string{"obdOdometerMeters"} // []string | The stat type you want this endpoint to return information on. Currently only one stat type is accepted per request.  - `engineStates`: The state of the engine (`Off`, `On`, `Idle`). - `fuelPercents`: The engine fuel level in percentage points (e.g. `99`, `50`, etc). - `obdOdometerMeters`: The odometer reading according to on-board diagnostics. If Samsara does not have diagnostic coverage for a particular vehicle, the value for this stat type will be omitted. In these cases, we recommend using `gpsOdometerMeters`. - `gpsOdometerMeters`: The odometer reading according to GPS calculations. This calculation is based off GPS distance traveled and a manual odometer offset for a given vehicle, specified by the user in Samsara's dashboard UI or through the `odometerMeters` field in the [PATCH /fleet/vehicles/{id}](#operation/updateVehicleById) endpoint. `gpsOdometerMeters` is equal to the manual offset plus the GPS distance traveled since the offset was set. The value for this stat type will be omitted if a manual offset is not provided for a given vehicle. *A manual offset can only be provided when we do not have diagnostic coverage for a particular vehicle.* - `obdEngineSeconds`: The cumulative number of seconds the engine has run according to on-board diagnostics. - `gpsDistanceMeters`: The distance the vehicle has traveled since the gateway was installed based on GPS calculations.
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)
	tagIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of tag IDs. Example: `tagIds=1234,5678` (optional)
	vehicleIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of vehicle IDs. Example: `vehicleIds=1234,5678` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetVehicleStatsFeed(context.Background()).Types(types).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVehicleStatsFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicleStatsFeed`: VehicleStatsListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVehicleStatsFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVehicleStatsFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **types** | **[]string** | The stat type you want this endpoint to return information on. Currently only one stat type is accepted per request.  - &#x60;engineStates&#x60;: The state of the engine (&#x60;Off&#x60;, &#x60;On&#x60;, &#x60;Idle&#x60;). - &#x60;fuelPercents&#x60;: The engine fuel level in percentage points (e.g. &#x60;99&#x60;, &#x60;50&#x60;, etc). - &#x60;obdOdometerMeters&#x60;: The odometer reading according to on-board diagnostics. If Samsara does not have diagnostic coverage for a particular vehicle, the value for this stat type will be omitted. In these cases, we recommend using &#x60;gpsOdometerMeters&#x60;. - &#x60;gpsOdometerMeters&#x60;: The odometer reading according to GPS calculations. This calculation is based off GPS distance traveled and a manual odometer offset for a given vehicle, specified by the user in Samsara&#39;s dashboard UI or through the &#x60;odometerMeters&#x60; field in the [PATCH /fleet/vehicles/{id}](#operation/updateVehicleById) endpoint. &#x60;gpsOdometerMeters&#x60; is equal to the manual offset plus the GPS distance traveled since the offset was set. The value for this stat type will be omitted if a manual offset is not provided for a given vehicle. *A manual offset can only be provided when we do not have diagnostic coverage for a particular vehicle.* - &#x60;obdEngineSeconds&#x60;: The cumulative number of seconds the engine has run according to on-board diagnostics. - &#x60;gpsDistanceMeters&#x60;: The distance the vehicle has traveled since the gateway was installed based on GPS calculations. | 
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 
 **tagIds** | **[]string** | A filter on the data based on this comma-separated list of tag IDs. Example: &#x60;tagIds&#x3D;1234,5678&#x60; | 
 **vehicleIds** | **[]string** | A filter on the data based on this comma-separated list of vehicle IDs. Example: &#x60;vehicleIds&#x3D;1234,5678&#x60; | 

### Return type

[**VehicleStatsListResponse**](VehicleStatsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicleStatsHistory

> VehicleStatsListResponse GetVehicleStatsHistory(ctx).StartTime(startTime).EndTime(endTime).Types(types).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()

Get historical vehicle stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	startTime := time.Now() // time.Time | A start time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00).
	endTime := time.Now() // time.Time | An end time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00).
	types := []string{"obdOdometerMeters"} // []string | The stat type you want this endpoint to return information on. Currently only one stat type is accepted per request.  - `engineStates`: The state of the engine (`Off`, `On`, `Idle`). - `fuelPercents`: The engine fuel level in percentage points (e.g. `99`, `50`, etc). - `obdOdometerMeters`: The odometer reading according to on-board diagnostics. If Samsara does not have diagnostic coverage for a particular vehicle, the value for this stat type will be omitted. In these cases, we recommend using `gpsOdometerMeters`. - `gpsOdometerMeters`: The odometer reading according to GPS calculations. This calculation is based off GPS distance traveled and a manual odometer offset for a given vehicle, specified by the user in Samsara's dashboard UI or through the `odometerMeters` field in the [PATCH /fleet/vehicles/{id}](#operation/updateVehicleById) endpoint. `gpsOdometerMeters` is equal to the manual offset plus the GPS distance traveled since the offset was set. The value for this stat type will be omitted if a manual offset is not provided for a given vehicle. *A manual offset can only be provided when we do not have diagnostic coverage for a particular vehicle.* - `obdEngineSeconds`: The cumulative number of seconds the engine has run according to on-board diagnostics. - `gpsDistanceMeters`: The distance the vehicle has traveled since the gateway was installed based on GPS calculations.
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)
	tagIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of tag IDs. Example: `tagIds=1234,5678` (optional)
	vehicleIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of vehicle IDs. Example: `vehicleIds=1234,5678` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetVehicleStatsHistory(context.Background()).StartTime(startTime).EndTime(endTime).Types(types).After(after).TagIds(tagIds).VehicleIds(vehicleIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVehicleStatsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicleStatsHistory`: VehicleStatsListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVehicleStatsHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVehicleStatsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **time.Time** | A start time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00). | 
 **endTime** | **time.Time** | An end time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00). | 
 **types** | **[]string** | The stat type you want this endpoint to return information on. Currently only one stat type is accepted per request.  - &#x60;engineStates&#x60;: The state of the engine (&#x60;Off&#x60;, &#x60;On&#x60;, &#x60;Idle&#x60;). - &#x60;fuelPercents&#x60;: The engine fuel level in percentage points (e.g. &#x60;99&#x60;, &#x60;50&#x60;, etc). - &#x60;obdOdometerMeters&#x60;: The odometer reading according to on-board diagnostics. If Samsara does not have diagnostic coverage for a particular vehicle, the value for this stat type will be omitted. In these cases, we recommend using &#x60;gpsOdometerMeters&#x60;. - &#x60;gpsOdometerMeters&#x60;: The odometer reading according to GPS calculations. This calculation is based off GPS distance traveled and a manual odometer offset for a given vehicle, specified by the user in Samsara&#39;s dashboard UI or through the &#x60;odometerMeters&#x60; field in the [PATCH /fleet/vehicles/{id}](#operation/updateVehicleById) endpoint. &#x60;gpsOdometerMeters&#x60; is equal to the manual offset plus the GPS distance traveled since the offset was set. The value for this stat type will be omitted if a manual offset is not provided for a given vehicle. *A manual offset can only be provided when we do not have diagnostic coverage for a particular vehicle.* - &#x60;obdEngineSeconds&#x60;: The cumulative number of seconds the engine has run according to on-board diagnostics. - &#x60;gpsDistanceMeters&#x60;: The distance the vehicle has traveled since the gateway was installed based on GPS calculations. | 
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 
 **tagIds** | **[]string** | A filter on the data based on this comma-separated list of tag IDs. Example: &#x60;tagIds&#x3D;1234,5678&#x60; | 
 **vehicleIds** | **[]string** | A filter on the data based on this comma-separated list of vehicle IDs. Example: &#x60;vehicleIds&#x3D;1234,5678&#x60; | 

### Return type

[**VehicleStatsListResponse**](VehicleStatsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddresses

> ListAddressesResponse ListAddresses(ctx).Limit(limit).After(after).TagIds(tagIds).Execute()

List all addresses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int64(789) // int64 | The limit for how many objects will be in the response. Default and max for this value is 512 objects. (optional) (default to 512)
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)
	tagIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of tag IDs. Example: `tagIds=1234,5678` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListAddresses(context.Background()).Limit(limit).After(after).TagIds(tagIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAddresses`: ListAddressesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The limit for how many objects will be in the response. Default and max for this value is 512 objects. | [default to 512]
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 
 **tagIds** | **[]string** | A filter on the data based on this comma-separated list of tag IDs. Example: &#x60;tagIds&#x3D;1234,5678&#x60; | 

### Return type

[**ListAddressesResponse**](ListAddressesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContacts

> ListContactsResponse ListContacts(ctx).Limit(limit).After(after).Execute()

List all contacts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int64(789) // int64 | The limit for how many objects will be in the response. Default and max for this value is 512 objects. (optional) (default to 512)
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListContacts(context.Background()).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContacts`: ListContactsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The limit for how many objects will be in the response. Default and max for this value is 512 objects. | [default to 512]
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 

### Return type

[**ListContactsResponse**](ListContactsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDrivers

> ListDriversResponse ListDrivers(ctx).IsDeactivated(isDeactivated).Limit(limit).After(after).TagIds(tagIds).UpdatedAfterTime(updatedAfterTime).CreatedAfterTime(createdAfterTime).Execute()

List all drivers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	isDeactivated := true // bool | If value is true, only drivers that are deactivated will appear in the response. This parameter will default to false if not provided (fetching only active drivers). (optional)
	limit := int64(789) // int64 | The limit for how many objects will be in the response. Default and max for this value is 512 objects. (optional) (default to 512)
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)
	tagIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of tag IDs. Example: `tagIds=1234,5678` (optional)
	updatedAfterTime := time.Now() // time.Time | A filter on data to have an updated at time after or equal to this specified time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00). (optional)
	createdAfterTime := time.Now() // time.Time | A filter on data to have a created at time after or equal to this specified time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListDrivers(context.Background()).IsDeactivated(isDeactivated).Limit(limit).After(after).TagIds(tagIds).UpdatedAfterTime(updatedAfterTime).CreatedAfterTime(createdAfterTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListDrivers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDrivers`: ListDriversResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListDrivers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDriversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDeactivated** | **bool** | If value is true, only drivers that are deactivated will appear in the response. This parameter will default to false if not provided (fetching only active drivers). | 
 **limit** | **int64** | The limit for how many objects will be in the response. Default and max for this value is 512 objects. | [default to 512]
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 
 **tagIds** | **[]string** | A filter on the data based on this comma-separated list of tag IDs. Example: &#x60;tagIds&#x3D;1234,5678&#x60; | 
 **updatedAfterTime** | **time.Time** | A filter on data to have an updated at time after or equal to this specified time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00). | 
 **createdAfterTime** | **time.Time** | A filter on data to have a created at time after or equal to this specified time in RFC 3339 format. Millisecond precision and timezones are supported. (Examples: 2019-06-13T19:08:25Z, 2019-06-13T19:08:25.455Z, OR 2015-09-15T14:00:12-04:00). | 

### Return type

[**ListDriversResponse**](ListDriversResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTags

> ListTagsResponse ListTags(ctx).Limit(limit).After(after).Execute()

List all tags



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int64(789) // int64 | The limit for how many objects will be in the response. Default and max for this value is 512 objects. (optional) (default to 512)
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListTags(context.Background()).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTags`: ListTagsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The limit for how many objects will be in the response. Default and max for this value is 512 objects. | [default to 512]
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 

### Return type

[**ListTagsResponse**](ListTagsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserRoles

> ListUserRolesResponse ListUserRoles(ctx).Limit(limit).After(after).Execute()

List all user roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int64(789) // int64 | The limit for how many objects will be in the response. Default and max for this value is 512 objects. (optional) (default to 512)
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListUserRoles(context.Background()).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListUserRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserRoles`: ListUserRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListUserRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The limit for how many objects will be in the response. Default and max for this value is 512 objects. | [default to 512]
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 

### Return type

[**ListUserRolesResponse**](ListUserRolesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> ListUsersResponse ListUsers(ctx).Limit(limit).After(after).Execute()

List all users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int64(789) // int64 | The limit for how many objects will be in the response. Default and max for this value is 512 objects. (optional) (default to 512)
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListUsers(context.Background()).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: ListUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The limit for how many objects will be in the response. Default and max for this value is 512 objects. | [default to 512]
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 

### Return type

[**ListUsersResponse**](ListUsersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVehicles

> ListVehiclesResponse ListVehicles(ctx).Limit(limit).After(after).TagIds(tagIds).Execute()

List all vehicles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int64(789) // int64 | The limit for how many objects will be in the response. Default and max for this value is 512 objects. (optional) (default to 512)
	after := "after_example" // string | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. (optional)
	tagIds := []string{"Inner_example"} // []string | A filter on the data based on this comma-separated list of tag IDs. Example: `tagIds=1234,5678` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListVehicles(context.Background()).Limit(limit).After(after).TagIds(tagIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListVehicles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVehicles`: ListVehiclesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListVehicles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVehiclesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The limit for how many objects will be in the response. Default and max for this value is 512 objects. | [default to 512]
 **after** | **string** | If specified, this should be the endCursor value from the previous page of results. When present, this request will return the next page of results that occur immediately after the previous page of results. | 
 **tagIds** | **[]string** | A filter on the data based on this comma-separated list of tag IDs. Example: &#x60;tagIds&#x3D;1234,5678&#x60; | 

### Return type

[**ListVehiclesResponse**](ListVehiclesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceTag

> TagResponse ReplaceTag(ctx, id).Tag(tag).Execute()

Update a tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique identifier for the tag.
	tag := *openapiclient.NewReplaceTagRequest() // ReplaceTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ReplaceTag(context.Background(), id).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ReplaceTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceTag`: TagResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ReplaceTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | [**ReplaceTagRequest**](ReplaceTagRequest.md) |  | 

### Return type

[**TagResponse**](TagResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddress

> AddressResponse UpdateAddress(ctx, id).Address(address).Execute()

Update an address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | ID of the Address. This can either be the Samsara-provided ID or an external ID. External IDs are customer-specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: `key:value`. For example, `crmId:abc123`
	address := *openapiclient.NewUpdateAddressRequest() // UpdateAddressRequest | The address fields to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateAddress(context.Background(), id).Address(address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAddress`: AddressResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Address. This can either be the Samsara-provided ID or an external ID. External IDs are customer-specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: &#x60;key:value&#x60;. For example, &#x60;crmId:abc123&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **address** | [**UpdateAddressRequest**](UpdateAddressRequest.md) | The address fields to update. | 

### Return type

[**AddressResponse**](AddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> ContactResponse UpdateContact(ctx, id).Contact(contact).Execute()

Update a contact



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique identifier for the contact.
	contact := *openapiclient.NewUpdateContactRequest() // UpdateContactRequest | Updates to the contact.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateContact(context.Background(), id).Contact(contact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContact`: ContactResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the contact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contact** | [**UpdateContactRequest**](UpdateContactRequest.md) | Updates to the contact. | 

### Return type

[**ContactResponse**](ContactResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDriver

> DriverResponse UpdateDriver(ctx, id).Driver(driver).Execute()

Update a driver



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | ID of the driver. This can either be the Samsara-specified ID, or an external ID. External IDs are customer specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: `key:value`. For example, `payrollId:ABFS18600`
	driver := *openapiclient.NewUpdateDriverRequest() // UpdateDriverRequest | Updates to the driver properties.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateDriver(context.Background(), id).Driver(driver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateDriver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDriver`: DriverResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateDriver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the driver. This can either be the Samsara-specified ID, or an external ID. External IDs are customer specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: &#x60;key:value&#x60;. For example, &#x60;payrollId:ABFS18600&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDriverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **driver** | [**UpdateDriverRequest**](UpdateDriverRequest.md) | Updates to the driver properties. | 

### Return type

[**DriverResponse**](DriverResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UserResponse UpdateUser(ctx, id).User(user).Execute()

Update a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique identifier for the user.
	user := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest | Updates to the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateUser(context.Background(), id).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**UpdateUserRequest**](UpdateUserRequest.md) | Updates to the user. | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVehicle

> VehicleResponse UpdateVehicle(ctx, id).Vehicle(vehicle).Execute()

Update a vehicle



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | ID of the vehicle. This can either be the Samsara-specified ID, or an external ID. External IDs are customer specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: `key:value`. For example, `maintenanceId:250020`
	vehicle := *openapiclient.NewUpdateVehicleRequest() // UpdateVehicleRequest | Fields that can be patched on a vehicle.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateVehicle(context.Background(), id).Vehicle(vehicle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateVehicle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVehicle`: VehicleResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateVehicle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the vehicle. This can either be the Samsara-specified ID, or an external ID. External IDs are customer specified key-value pairs created in the POST or PATCH requests of this resource. To specify an external ID as part of a path parameter, use the following format: &#x60;key:value&#x60;. For example, &#x60;maintenanceId:250020&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vehicle** | [**UpdateVehicleRequest**](UpdateVehicleRequest.md) | Fields that can be patched on a vehicle. | 

### Return type

[**VehicleResponse**](VehicleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

