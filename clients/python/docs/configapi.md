# ConfigApi

## lakefs\_client.ConfigApi

All URIs are relative to [http://localhost/api/v1](http://localhost/api/v1)

| Method | HTTP request | Description |
| :--- | :--- | :--- |
| [**get\_storage\_config**](configapi.md#get_storage_config) | **GET** /config/storage |  |
| [**setup**](configapi.md#setup) | **POST** /setup\_lakefs | setup lakeFS and create a first user |

## **get\_storage\_config**

> StorageConfig get\_storage\_config\(\)

retrieve lakeFS storage configuration

### Example

* Basic Authentication \(basic\_auth\):
* Api Key Authentication \(cookie\_auth\):
* Bearer \(JWT\) Authentication \(jwt\_token\):

  \`\`\`python

  import time

  import lakefs\_client

  from lakefs\_client.api import config\_api

  from lakefs\_client.model.error import Error

  from lakefs\_client.model.storage\_config import StorageConfig

  from pprint import pprint

  **Defining the host is optional and defaults to** [**http://localhost/api/v1**](http://localhost/api/v1)\*\*\*\*

  **See configuration.py for a list of all supported configuration parameters.**

  configuration = lakefs\_client.Configuration\(

    host = "[http://localhost/api/v1](http://localhost/api/v1)"

  \)

## The client must configure the authentication and authorization parameters

## in accordance with the API server security policy.

## Examples for each auth method are provided below, use the example that

## satisfies your auth use case.

## Configure HTTP basic authorization: basic\_auth

configuration = lakefs\_client.Configuration\( username = 'YOUR\_USERNAME', password = 'YOUR\_PASSWORD' \)

## Configure API key authorization: cookie\_auth

configuration.api\_key\['cookie\_auth'\] = 'YOUR\_API\_KEY'

## Uncomment below to setup prefix \(e.g. Bearer\) for API key, if needed

## configuration.api\_key\_prefix\['cookie\_auth'\] = 'Bearer'

## Configure Bearer authorization \(JWT\): jwt\_token

configuration = lakefs\_client.Configuration\( access\_token = 'YOUR\_BEARER\_TOKEN' \)

## Enter a context with an instance of the API client

with lakefs\_client.ApiClient\(configuration\) as api\_client:

```text
# Create an instance of the API class
api_instance = config_api.ConfigApi(api_client)

# example, this endpoint has no required or optional parameters
try:
    api_response = api_instance.get_storage_config()
    pprint(api_response)
except lakefs_client.ApiException as e:
    print("Exception when calling ConfigApi->get_storage_config: %s\n" % e)
```

```text
### Parameters
This endpoint does not need any parameter.

### Return type

[**StorageConfig**](StorageConfig.md)

### Authorization

[basic_auth](../README.md#basic_auth), [cookie_auth](../README.md#cookie_auth), [jwt_token](../README.md#jwt_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | lakeFS storage configuration |  -  |
**401** | Unauthorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **setup**
> CredentialsWithSecret setup(setup)

setup lakeFS and create a first user

### Example

```python
import time
import lakefs_client
from lakefs_client.api import config_api
from lakefs_client.model.credentials_with_secret import CredentialsWithSecret
from lakefs_client.model.setup import Setup
from lakefs_client.model.error import Error
from pprint import pprint
# Defining the host is optional and defaults to http://localhost/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = lakefs_client.Configuration(
    host = "http://localhost/api/v1"
)


# Enter a context with an instance of the API client
with lakefs_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = config_api.ConfigApi(api_client)
    setup = Setup(
        username="username_example",
        key=AccessKeyCredentials(
            access_key_id="AKIAIOSFODNN7EXAMPLE",
            secret_access_key="wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
        ),
    ) # Setup | 

    # example passing only required values which don't have defaults set
    try:
        # setup lakeFS and create a first user
        api_response = api_instance.setup(setup)
        pprint(api_response)
    except lakefs_client.ApiException as e:
        print("Exception when calling ConfigApi->setup: %s\n" % e)
```

### Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **setup** | [**Setup**](setup.md) |  |  |

### Return type

[**CredentialsWithSecret**](credentialswithsecret.md)

### Authorization

No authorization required

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
| :--- | :--- | :--- |


**200** \| user created successfully \| - \| **400** \| bad request \| - \| **409** \| setup was already called \| - \| **0** \| Internal Server Error \| - \|

[\[Back to top\]](configapi.md) [\[Back to API list\]](../#documentation-for-api-endpoints) [\[Back to Model list\]](../#documentation-for-models) [\[Back to README\]](../)

