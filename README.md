# go-ssm-aws
A thin go client that interfaces with [AWS SSM](https://www.amazonaws.cn/en/systems-manager/).

## Why this package?
This package is a thin wrapper over the `aws-sdk-go` and hides the complexity dealing with the GO AWS SDK.
A good use case for this package is when secure parameters for an application are stored in
[AWS Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html). 
During application startup this package can be used to fetch the params (typically DB Hostname, username, password etc) and can be used in the application.

## Example

### Installation

```bash
go get -u github.com/IamMayankThakur/go-ssm-aws
```

### Basic Usage

```go
    // If the following parameters are present in SSM:
    // /path/to/key  -> with value "value" 

    import ssm "github.com/IamMayankThakur/go-ssm-aws"

    cfg := ssm.Config {
        Enabled: true
        SecretsPath: "/path/to"
        Region: "us-east-1"
    }
    ssmClient, err := ssm.New(&cfg)
    if err != nil {
        return err
    }
    // And getting a specific value
    value, err := ssmClient.GetValueByName("key", true)
    // value should be "value"
```
