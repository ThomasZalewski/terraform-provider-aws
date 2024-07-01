---
subcategory: "Transfer Family"
layout: "aws"
page_title: "AWS: aws_transfer_connector"
description: |-
  Terraform data source for managing an AWS Transfer Family Connector.
---
<!---
TIP: A few guiding principles for writing documentation:
1. Use simple language while avoiding jargon and figures of speech.
2. Focus on brevity and clarity to keep a reader's attention.
3. Use active voice and present tense whenever you can.
4. Document your feature as it exists now; do not mention the future or past if you can help it.
5. Use accessible and inclusive language.
--->

# Data Source: aws_transfer_connector

Terraform data source for managing an AWS Transfer Family Connector.

## Example Usage

### Basic Usage

```terraform
data "aws_transfer_connector" "this" {
  connector_id = "c-xxxxxxxxxxxxxx"
}
```

## Argument Reference

The following arguments are required:

* `connector_id` - (Required) Unique identifier for connector


## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `access_role` - ARN of the AWS Identity and Access Managment role.
* `arn` - ARN of the Connector. 
* `as2_config` - Structure containing the parameters for an AS2 connector object. Contains the following attributes:

  
* `connector_id` - Unique identifier for the connector.
* `logging_role` -  ARN of the IAM role that allows a connector to turn on CLoudwatch logging for Amazon S3 events.
* `security_policy_name` - Name of security policy.
* `service_managed_egress_ip_addresses` - List of egress Ip addresses.
  * `` -




