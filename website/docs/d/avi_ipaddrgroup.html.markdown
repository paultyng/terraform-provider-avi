############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_ipaddrgroup"
sidebar_current: "docs-avi-datasource-ipaddrgroup"
description: |-
  Get information of Avi IpAddrGroup.
---

# avi_ipaddrgroup

This data source is used to to get avi_ipaddrgroup objects.

## Example Usage

```hcl
data "avi_ipaddrgroup" "foo_ipaddrgroup" {
    uuid = "ipaddrgroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search IpAddrGroup by name.
* `uuid` - (Optional) Search IpAddrGroup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `addrs` - Configure ip address(es).
* `apic_epg_name` - Populate ip addresses from members of this cisco apic epg.
* `country_codes` - Populate the ip address ranges from the geo database for this country.
* `description` - User defined description for the object.
* `ip_ports` - Configure (ip address, port) tuple(s).
* `labels` - Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `marathon_app_name` - Populate ip addresses from tasks of this marathon app.
* `marathon_service_port` - Task port associated with marathon service port. If marathon app has multiple service ports, this is required. Else, the first task port is used.
* `name` - Name of the ip address group.
* `prefixes` - Configure ip address prefix(es).
* `ranges` - Configure ip address range(s).
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the ip address group.

