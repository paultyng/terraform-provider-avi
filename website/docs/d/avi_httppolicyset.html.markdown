<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `cloud_config_cksum` - Checksum of cloud configuration for pool. Internally set by cloud connector. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `created_by` - Creator name. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `description` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `geo_db_ref` - Geo database. It is a reference to an object of type geodb. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `http_request_policy` - Http request policy for the virtual service. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `http_response_policy` - Http response policy for the virtual service. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `http_security_policy` - Http security policy for the virtual service. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ip_reputation_db_ref` - Ip reputation database. It is a reference to an object of type ipreputationdb. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `is_internal_policy` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - Name of the http policy set. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the http policy set. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

