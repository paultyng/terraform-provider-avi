############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_ssopolicy"
sidebar_current: "docs-avi-datasource-ssopolicy"
description: |-
  Get information of Avi SSOPolicy.
---

# avi_ssopolicy

This data source is used to to get avi_ssopolicy objects.

## Example Usage

```hcl
data "avi_ssopolicy" "foo_ssopolicy" {
    uuid = "ssopolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SSOPolicy by name.
* `uuid` - (Optional) Search SSOPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `authentication_policy` - Authentication policy settings. Field introduced in 18.2.1.
* `authorization_policy` - Authorization policy settings. Field introduced in 18.2.5.
* `labels` - Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `name` - Name of the sso policy. Field introduced in 18.2.3.
* `tenant_ref` - Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 18.2.3.
* `type` - Sso policy type. Enum options - SSO_TYPE_SAML, SSO_TYPE_PINGACCESS, SSO_TYPE_JWT, SSO_TYPE_LDAP. Field introduced in 18.2.5.
* `uuid` - Uuid of the sso policy. Field introduced in 18.2.3.

