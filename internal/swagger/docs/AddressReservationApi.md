# \AddressReservationApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](AddressReservationApi.md#Get) | **Get** /address/{addressId}/reservation | Get a list of address reservations
[**Post**](AddressReservationApi.md#Post) | **Post** /address/{addressId}/reservation/{id}/issue | Issues authorizations for an address reservation
[**Post_0**](AddressReservationApi.md#Post_0) | **Post** /address/{addressId}/reservation/{id}/revoke | Revoke authorizations for an address reservation


# **Get**
> []AddressReservation Get(ctx, addressId)
Get a list of address reservations



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressId** | **int32**| The address id | 

### Return type

[**[]AddressReservation**](AddressReservation.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, addressId, id)
Issues authorizations for an address reservation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressId** | **int32**| The address id | 
  **id** | **string**| The address reservation id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_0**
> Post_0(ctx, addressId, id)
Revoke authorizations for an address reservation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressId** | **int32**| The address id | 
  **id** | **string**| The address reservation id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

