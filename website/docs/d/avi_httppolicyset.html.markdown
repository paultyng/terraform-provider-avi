############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_httppolicyset"
sidebar_current: "docs-avi-datasource-httppolicyset"
description: |-
  Get information of Avi HTTPPolicySet.
---

# avi_httppolicyset

This data source is used to to get avi_httppolicyset objects.

## Example Usage

```hcl
data "avi_httppolicyset" "foo_httppolicyset" {
    uuid = "httppolicyset-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search HTTPPolicySet by name.
* `uuid` - (Optional) Search HTTPPolicySet by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_config_cksum` - Checksum of cloud configuration for pool. Internally set by cloud connector.
* `created_by` - Creator name.
* `description` - User defined description for the object.
* `geo_db_ref` - Geo database. It is a reference to an object of type geodb. Field introduced in 21.1.1.
* `http_request_policy` - Http request policy for the virtual service.
* `http_response_policy` - Http response policy for the virtual service.
* `http_security_policy` - Http security policy for the virtual service.
* `ip_reputation_db_ref` - Ip reputation database. It is a reference to an object of type ipreputationdb. Field introduced in 20.1.3.
* `is_internal_policy` - Boolean flag to set is_internal_policy.
* `labels` - Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `name` - Name of the http policy set.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the http policy set.

