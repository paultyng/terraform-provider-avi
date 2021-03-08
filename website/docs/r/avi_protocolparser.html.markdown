############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_protocolparser"
sidebar_current: "docs-avi-resource-protocolparser"
description: |-
  Creates and manages Avi ProtocolParser.
---

# avi_protocolparser

The ProtocolParser resource allows the creation and management of Avi ProtocolParser

## Example Usage

```hcl
resource "avi_protocolparser" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the protocol parser. Field introduced in 18.2.3.
* `parser_code` - (Required) Command script provided inline. Field introduced in 18.2.3.
* `description` - (Optional) Description of the protocol parser. Field introduced in 18.2.3.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `tenant_ref` - (Optional) Tenant uuid of the protocol parser. It is a reference to an object of type tenant. Field introduced in 18.2.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the protocol parser. Field introduced in 18.2.3.

