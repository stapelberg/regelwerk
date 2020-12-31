# \AddressApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](AddressApi.md#Get) | **Get** /address | Get a list of addresses
[**Post**](AddressApi.md#Post) | **Post** /address/{addressId} | Update an address


# **Get**
> []Address Get(ctx, )
Get a list of addresses



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Address**](Address.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, body, addressId)
Update an address



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddressUpdate**](AddressUpdate.md)| Address update representation | 
  **addressId** | **int32**| The address id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

