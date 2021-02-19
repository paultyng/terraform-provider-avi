############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_gslbservice"
sidebar_current: "docs-avi-datasource-gslbservice"
description: |-
  Get information of Avi GslbService.
---

# avi_gslbservice

This data source is used to to get avi_gslbservice objects.

## Example Usage

```hcl
data "avi_gslbservice" "foo_gslbservice" {
    uuid = "gslbservice-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search GslbService by name.
* `uuid` - (Optional) Search GslbService by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `application_persistence_profile_ref` - The federated application persistence associated with gslbservice site persistence functionality. It is a reference to an object of type applicationpersistenceprofile. Field introduced in 17.2.1.
* `controller_health_status_enabled` - Gs member's overall health status is derived based on a combination of controller and datapath health-status inputs. Note that the datapath status is determined by the association of health monitor profiles. Only the controller provided status is determined through this configuration.
* `created_by` - Creator name. Field introduced in 17.1.2.
* `description` - User defined description for the object.
* `domain_names` - Fully qualified domain name of the gslb service. Minimum of 1 items required.
* `down_response` - Response to the client query when the gslb service is down.
* `enabled` - Enable or disable the gslb service. If the gslb service is enabled, then the vips are sent in the dns responses based on reachability and configured algorithm. If the gslb service is disabled, then the vips are no longer available in the dns response.
* `groups` - Select list of pools belonging to this gslb service. Minimum of 1 items required.
* `health_monitor_refs` - Verify vs health by applying one or more health monitors. Active monitors generate synthetic traffic from dns service engine and to mark a vs up or down based on the response. It is a reference to an object of type healthmonitor. Maximum of 6 items allowed.
* `health_monitor_scope` - Health monitor probe can be executed for all the members or it can be executed only for third-party members. This operational mode is useful to reduce the number of health monitor probes in case of a hybrid scenario. In such a case, avi members can have controller derived status while non-avi members can be probed by via health monitor probes in dataplane. Enum options - GSLB_SERVICE_HEALTH_MONITOR_ALL_MEMBERS, GSLB_SERVICE_HEALTH_MONITOR_ONLY_NON_AVI_MEMBERS.
* `hm_off` - This field is an internal field and is used in se. Field introduced in 18.2.2.
* `is_federated` - This field indicates that this object is replicated across gslb federation. Field introduced in 17.1.3.
* `labels` - Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `min_members` - The minimum number of members to distribute traffic to. Allowed values are 1-65535. Special values are 0 - 'disable'. Field introduced in 17.2.4.
* `name` - Name for the gslb service.
* `num_dns_ip` - Number of ip addresses of this gslb service to be returned by the dns service. Enter 0 to return all ip addresses. Allowed values are 1-20. Special values are 0- 'return all ip addresses'.
* `pool_algorithm` - The load balancing algorithm will pick a gslb pool within the gslb service list of available pools. Enum options - GSLB_SERVICE_ALGORITHM_PRIORITY, GSLB_SERVICE_ALGORITHM_GEO. Field introduced in 17.2.3.
* `resolve_cname` - This field indicates that for a cname query, respond with resolved cnames in the additional section with a records. Field introduced in 18.2.5.
* `site_persistence_enabled` - Enable site-persistence for the gslbservice. Field introduced in 17.2.1.
* `tenant_ref` - It is a reference to an object of type tenant.
* `ttl` - Ttl value (in seconds) for records served for this gslb service by the dns service. Allowed values are 0-86400. Unit is sec.
* `use_edns_client_subnet` - Use the client ip subnet from the edns option as source ipaddress for client geo-location and consistent hash algorithm. Default is true. Field introduced in 17.1.1.
* `uuid` - Uuid of the gslb service.
* `wildcard_match` - Enable wild-card match of fqdn  if an exact match is not found in the dns table, the longest match is chosen by wild-carding the fqdn in the dns request. Default is false. Field introduced in 17.1.1.

