############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_applicationpersistenceprofile"
sidebar_current: "docs-avi-resource-applicationpersistenceprofile"
description: |-
  Creates and manages Avi ApplicationPersistenceProfile.
---

# avi_applicationpersistenceprofile

The ApplicationPersistenceProfile resource allows the creation and management of Avi ApplicationPersistenceProfile

## Example Usage

```hcl
resource "avi_applicationpersistenceprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A user-friendly name for the persistence profile.
* `persistence_type` - (Required) Method used to persist clients to the same server for a duration of time or a session. Enum options - PERSISTENCE_TYPE_CLIENT_IP_ADDRESS, PERSISTENCE_TYPE_HTTP_COOKIE, PERSISTENCE_TYPE_TLS, PERSISTENCE_TYPE_CLIENT_IPV6_ADDRESS, PERSISTENCE_TYPE_CUSTOM_HTTP_HEADER, PERSISTENCE_TYPE_APP_COOKIE, PERSISTENCE_TYPE_GSLB_SITE. Allowed in basic(allowed values- persistence_type_client_ip_address,persistence_type_http_cookie) edition, essentials(allowed values- persistence_type_client_ip_address,persistence_type_http_cookie) edition, enterprise edition.
* `app_cookie_persistence_profile` - (Optional) Specifies the application cookie persistence profile parameters.
* `description` - (Optional) User defined description for the object.
* `hdr_persistence_profile` - (Optional) Specifies the custom http header persistence profile parameters.
* `http_cookie_persistence_profile` - (Optional) Specifies the http cookie persistence profile parameters.
* `ip_persistence_profile` - (Optional) Specifies the client ip persistence profile parameters.
* `is_federated` - (Optional) This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 17.1.3.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `server_hm_down_recovery` - (Optional) Specifies behavior when a persistent server has been marked down by a health monitor. Enum options - HM_DOWN_PICK_NEW_SERVER, HM_DOWN_ABORT_CONNECTION, HM_DOWN_CONTINUE_PERSISTENT_SERVER. Allowed in basic(allowed values- hm_down_pick_new_server) edition, essentials(allowed values- hm_down_pick_new_server) edition, enterprise edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the persistence profile.

