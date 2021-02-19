############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_poolgroupdeploymentpolicy"
sidebar_current: "docs-avi-datasource-poolgroupdeploymentpolicy"
description: |-
  Get information of Avi PoolGroupDeploymentPolicy.
---

# avi_poolgroupdeploymentpolicy

This data source is used to to get avi_poolgroupdeploymentpolicy objects.

## Example Usage

```hcl
data "avi_poolgroupdeploymentpolicy" "foo_poolgroupdeploymentpolicy" {
    uuid = "poolgroupdeploymentpolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search PoolGroupDeploymentPolicy by name.
* `uuid` - (Optional) Search PoolGroupDeploymentPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `auto_disable_old_prod_pools` - It will automatically disable old production pools once there is a new production candidate.
* `description` - User defined description for the object.
* `evaluation_duration` - Duration of evaluation period for automatic deployment. Allowed values are 60-86400. Unit is sec.
* `labels` - Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `name` - The name of the pool group deployment policy.
* `rules` - List of list.
* `scheme` - Deployment scheme. Enum options - BLUE_GREEN, CANARY.
* `target_test_traffic_ratio` - Target traffic ratio before pool is made production. Allowed values are 1-100. Unit is ratio.
* `tenant_ref` - It is a reference to an object of type tenant.
* `test_traffic_ratio_rampup` - Ratio of the traffic that is sent to the pool under test. Test ratio of 100 means blue green. Allowed values are 1-100.
* `uuid` - Uuid of the pool group deployment policy.
* `webhook_ref` - Webhook configured with url that avi controller will pass back information about pool group, old and new pool information and current deployment rule results. It is a reference to an object of type webhook. Field introduced in 17.1.1.

