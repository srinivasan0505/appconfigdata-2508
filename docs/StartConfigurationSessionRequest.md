

# StartConfigurationSessionRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**applicationIdentifier** | **String** | The application ID or the application name. |  |
|**environmentIdentifier** | **String** | The environment ID or the environment name. |  |
|**configurationProfileIdentifier** | **String** | The configuration profile ID or the configuration profile name. |  |
|**requiredMinimumPollIntervalInSeconds** | **Integer** | Sets a constraint on a session. If you specify a value of, for example, 60 seconds, then the client that established the session can&#39;t call &lt;a&gt;GetLatestConfiguration&lt;/a&gt; more frequently than every 60 seconds. |  [optional] |



