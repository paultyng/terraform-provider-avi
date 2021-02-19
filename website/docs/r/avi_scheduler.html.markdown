############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_scheduler"
sidebar_current: "docs-avi-resource-scheduler"
description: |-
  Creates and manages Avi Scheduler.
---

# avi_scheduler

The Scheduler resource allows the creation and management of Avi Scheduler

## Example Usage

```hcl
resource "avi_scheduler" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of scheduler.
* `backup_config_ref` - (Optional) Backup configuration to be executed by this scheduler. It is a reference to an object of type backupconfiguration.
* `enabled` - (Optional) Boolean flag to set enabled.
* `end_date_time` - (Optional) Scheduler end date and time.
* `frequency` - (Optional) Frequency at which custom scheduler will run. Allowed values are 0-60.
* `frequency_unit` - (Optional) Unit at which custom scheduler will run. Enum options - SCHEDULER_FREQUENCY_UNIT_MIN, SCHEDULER_FREQUENCY_UNIT_HOUR, SCHEDULER_FREQUENCY_UNIT_DAY, SCHEDULER_FREQUENCY_UNIT_WEEK, SCHEDULER_FREQUENCY_UNIT_MONTH.
* `run_mode` - (Optional) Scheduler run mode. Enum options - RUN_MODE_PERIODIC, RUN_MODE_AT, RUN_MODE_NOW.
* `run_script_ref` - (Optional) Control script to be executed by this scheduler. It is a reference to an object of type alertscriptconfig.
* `scheduler_action` - (Optional) Define scheduler action. Enum options - SCHEDULER_ACTION_RUN_A_SCRIPT, SCHEDULER_ACTION_BACKUP.
* `start_date_time` - (Optional) Scheduler start date and time.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

