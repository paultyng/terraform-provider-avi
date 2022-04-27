<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_clusterclouddetails"
sidebar_current: "docs-avi-datasource-clusterclouddetails"
description: |-
  Get information of Avi ClusterCloudDetails.
---

# avi_clusterclouddetails

This data source is used to to get avi_clusterclouddetails objects.

## Example Usage

```hcl
data "avi_clusterclouddetails" "foo_clusterclouddetails" {
    uuid = "clusterclouddetails-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ClusterCloudDetails by name.
* `uuid` - (Optional) Search ClusterCloudDetails by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `azure_info` - Azure info to configure cluster_vip on the controller. Field introduced in 17.2.5. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - Field introduced in 17.2.5. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.2.5. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Field introduced in 17.2.5. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

