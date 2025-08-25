# DefaultApi

All URIs are relative to *http://appconfigdata.us-east-1.amazonaws.com*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getLatestConfiguration**](DefaultApi.md#getLatestConfiguration) | **GET** /configuration#configuration_token |  |
| [**startConfigurationSession**](DefaultApi.md#startConfigurationSession) | **POST** /configurationsessions |  |


<a id="getLatestConfiguration"></a>
# **getLatestConfiguration**
> GetLatestConfigurationResponse getLatestConfiguration(configurationToken, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



&lt;p&gt;Retrieves the latest deployed configuration. This API may return empty configuration data if the client already has the latest version. For more information about this API action and to view example CLI commands that show how to use it with the &lt;a&gt;StartConfigurationSession&lt;/a&gt; API action, see &lt;a href&#x3D;\&quot;http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration\&quot;&gt;Retrieving the configuration&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;. &lt;/p&gt; &lt;important&gt; &lt;p&gt;Note the following important information.&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;Each configuration token is only valid for one call to &lt;code&gt;GetLatestConfiguration&lt;/code&gt;. The &lt;code&gt;GetLatestConfiguration&lt;/code&gt; response includes a &lt;code&gt;NextPollConfigurationToken&lt;/code&gt; that should always replace the token used for the just-completed call in preparation for the next one. &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;GetLatestConfiguration&lt;/code&gt; is a priced call. For more information, see &lt;a href&#x3D;\&quot;https://aws.amazon.com/systems-manager/pricing/\&quot;&gt;Pricing&lt;/a&gt;.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;/important&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfigdata.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String configurationToken = "configurationToken_example"; // String | <p>Token describing the current state of the configuration session. To obtain a token, first call the <a>StartConfigurationSession</a> API. Note that every call to <code>GetLatestConfiguration</code> will return a new <code>ConfigurationToken</code> (<code>NextPollConfigurationToken</code> in the response) and <i>must</i> be provided to subsequent <code>GetLatestConfiguration</code> API calls.</p> <important> <p>This token should only be used once. To support long poll use cases, the token is valid for up to 24 hours. If a <code>GetLatestConfiguration</code> call uses an expired token, the system returns <code>BadRequestException</code>.</p> </important>
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      GetLatestConfigurationResponse result = apiInstance.getLatestConfiguration(configurationToken, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getLatestConfiguration");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **configurationToken** | **String**| &lt;p&gt;Token describing the current state of the configuration session. To obtain a token, first call the &lt;a&gt;StartConfigurationSession&lt;/a&gt; API. Note that every call to &lt;code&gt;GetLatestConfiguration&lt;/code&gt; will return a new &lt;code&gt;ConfigurationToken&lt;/code&gt; (&lt;code&gt;NextPollConfigurationToken&lt;/code&gt; in the response) and &lt;i&gt;must&lt;/i&gt; be provided to subsequent &lt;code&gt;GetLatestConfiguration&lt;/code&gt; API calls.&lt;/p&gt; &lt;important&gt; &lt;p&gt;This token should only be used once. To support long poll use cases, the token is valid for up to 24 hours. If a &lt;code&gt;GetLatestConfiguration&lt;/code&gt; call uses an expired token, the system returns &lt;code&gt;BadRequestException&lt;/code&gt;.&lt;/p&gt; &lt;/important&gt; | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**GetLatestConfigurationResponse**](GetLatestConfigurationResponse.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ThrottlingException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | BadRequestException |  -  |
| **483** | InternalServerException |  -  |

<a id="startConfigurationSession"></a>
# **startConfigurationSession**
> StartConfigurationSessionResponse startConfigurationSession(startConfigurationSessionRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Starts a configuration session used to retrieve a deployed configuration. For more information about this API action and to view example CLI commands that show how to use it with the &lt;a&gt;GetLatestConfiguration&lt;/a&gt; API action, see &lt;a href&#x3D;\&quot;http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration\&quot;&gt;Retrieving the configuration&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfigdata.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    StartConfigurationSessionRequest startConfigurationSessionRequest = new StartConfigurationSessionRequest(); // StartConfigurationSessionRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      StartConfigurationSessionResponse result = apiInstance.startConfigurationSession(startConfigurationSessionRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#startConfigurationSession");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **startConfigurationSessionRequest** | [**StartConfigurationSessionRequest**](StartConfigurationSessionRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**StartConfigurationSessionResponse**](StartConfigurationSessionResponse.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | ThrottlingException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | BadRequestException |  -  |
| **483** | InternalServerException |  -  |

