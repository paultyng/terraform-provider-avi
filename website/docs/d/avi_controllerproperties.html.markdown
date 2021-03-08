############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_controllerproperties"
sidebar_current: "docs-avi-datasource-controllerproperties"
description: |-
  Get information of Avi ControllerProperties.
---

# avi_controllerproperties

This data source is used to to get avi_controllerproperties objects.

## Example Usage

```hcl
data "avi_controllerproperties" "foo_controllerproperties" {
    uuid = "controllerproperties-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ControllerProperties by name.
* `uuid` - (Optional) Search ControllerProperties by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allow_admin_network_updates` - Allow non-admin tenants to update admin vrfcontext and network objects. Field introduced in 18.2.7, 20.1.1.
* `allow_ip_forwarding` - Field introduced in 17.1.1.
* `allow_unauthenticated_apis` - Allow unauthenticated access for special apis.
* `allow_unauthenticated_nodes` - Boolean flag to set allow_unauthenticated_nodes.
* `api_idle_timeout` - Allowed values are 0-1440. Unit is min.
* `api_perf_logging_threshold` - Threshold to log request timing in portal_performance.log and server-timing response header. Any stage taking longer than 1% of the threshold will be included in the server-timing header. Field introduced in 18.1.4, 18.2.1. Unit is milliseconds.
* `appviewx_compat_mode` - Export configuration in appviewx compatibility mode. Field introduced in 17.1.1. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `async_patch_merge_period` - Period for which asynchronous patch requests are queued. Allowed values are 30-120. Special values are 0 - 'deactivated'. Field introduced in 18.2.11, 20.1.3. Unit is sec.
* `async_patch_request_cleanup_duration` - Duration for which asynchronous patch requests should be kept, after being marked as success or fail. Allowed values are 5-120. Field introduced in 18.2.11, 20.1.3. Unit is min.
* `attach_ip_retry_interval` - Unit is sec.
* `attach_ip_retry_limit` - Placeholder for description of property attach_ip_retry_limit of obj type controllerproperties field type integer  type int.
* `bm_use_ansible` - Use ansible for se creation in baremetal. Field introduced in 17.2.2.
* `cleanup_expired_authtoken_timeout_period` - Period for auth token cleanup job. Field introduced in 18.1.1. Unit is min.
* `cleanup_sessions_timeout_period` - Period for sessions cleanup job. Field introduced in 18.1.1. Unit is min.
* `cloud_reconcile` - Enable/disable periodic reconcile for all the clouds. Field introduced in 17.2.14,18.1.5,18.2.1.
* `cluster_ip_gratuitous_arp_period` - Period for cluster ip gratuitous arp job. Unit is min.
* `consistency_check_timeout_period` - Period for consistency check job. Field introduced in 18.1.1. Unit is min.
* `controller_resource_info_collection_period` - Periodically collect stats. Field introduced in 20.1.3. Unit is min.
* `crashed_se_reboot` - Unit is sec.
* `dead_se_detection_timer` - Unit is sec.
* `default_minimum_api_timeout` - Minimum api timeout value.if this value is not 60, it will be the default timeout for all apis that do not have a specific timeout.if an api has a specific timeout but is less than this value, this value will become the new timeout. Allowed values are 60-3600. Field introduced in 18.2.6. Unit is sec.
* `del_offline_se_after_reboot_delay` - The amount of time the controller will wait before deleting an offline se after it has been rebooted. For unresponsive ses, the total time will be  unresponsive_se_reboot + del_offline_se_after_reboot_delay. For crashed ses, the total time will be crashed_se_reboot + del_offline_se_after_reboot_delay. Field introduced in 20.1.5. Unit is sec.
* `dns_refresh_period` - Period for refresh pool and gslb dns job. Unit is min. Allowed in basic(allowed values- 60) edition, essentials(allowed values- 60) edition, enterprise edition.
* `dummy` - Placeholder for description of property dummy of obj type controllerproperties field type integer  type int.
* `edit_system_limits` - Allow editing of system limits. Keep in mind that these system limits have been carefully selected based on rigorous testing in our testig environments. Modifying these limits could destabilize your cluster. Do this at your own risk!. Field introduced in 20.1.1.
* `enable_api_sharding` - This setting enables the controller leader to shard api requests to the followers (if any). Field introduced in 18.1.5, 18.2.1.
* `enable_memory_balancer` - Enable/disable memory balancer. Field introduced in 17.2.8.
* `fatal_error_lease_time` - Unit is sec.
* `federated_datastore_cleanup_duration` - Federated datastore will not cleanup diffs unless they are at least this duration in the past. Field introduced in 20.1.1. Unit is hours.
* `file_object_cleanup_period` - Period for file object cleanup job. Field introduced in 20.1.1. Unit is min.
* `max_dead_se_in_grp` - Placeholder for description of property max_dead_se_in_grp of obj type controllerproperties field type integer  type int.
* `max_pcap_per_tenant` - Maximum number of pcap files stored per tenant.
* `max_se_spawn_interval_delay` - Maximum delay possible to add to se_spawn_retry_interval after successive se spawn failure. Field introduced in 20.1.1. Unit is sec.
* `max_seq_attach_ip_failures` - Maximum number of consecutive attach ip failures that halts vs placement. Field introduced in 17.2.2.
* `max_seq_vnic_failures` - Placeholder for description of property max_seq_vnic_failures of obj type controllerproperties field type integer  type int.
* `max_threads_cc_vip_bg_worker` - Maximum number of threads in threadpool used by cloud connector ccvipbgworker. Allowed values are 1-100. Field introduced in 20.1.3.
* `permission_scoped_shared_admin_networks` - Network and vrfcontext objects from the admin tenant will not be shared to non-admin tenants unless admin permissions are granted. Field introduced in 18.2.7, 20.1.1.
* `persistence_key_rotate_period` - Period for rotate app persistence keys job. Allowed values are 1-1051200. Special values are 0 - 'disabled'. Unit is min. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition.
* `portal_request_burst_limit` - Burst limit on number of incoming requests. 0 to disable. Field introduced in 20.1.1.
* `portal_request_rate_limit` - Maximum average number of requests allowed per second. 0 to disable. Field introduced in 20.1.1. Unit is per_second.
* `portal_token` - Token used for uploading tech-support to portal. Field introduced in 16.4.6,17.1.2.
* `process_locked_useraccounts_timeout_period` - Period for process locked user accounts job. Field introduced in 18.1.1. Unit is min.
* `process_pki_profile_timeout_period` - Period for process pki profile job. Field introduced in 18.1.1. Unit is min.
* `query_host_fail` - Unit is sec.
* `resmgr_log_caching_period` - Period for each cycle of log caching in resource manager. At the end of each cycle, the in memory cached log history will be cleared. Field introduced in 20.1.5. Unit is sec.
* `safenet_hsm_version` - Version of the safenet package installed on the controller. Field introduced in 16.5.2,17.2.3.
* `se_create_timeout` - Unit is sec.
* `se_failover_attempt_interval` - Interval between attempting failovers to an se. Unit is sec.
* `se_from_marketplace` - This setting decides whether se is to be deployed from the cloud marketplace or to be created by the controller. The setting is applicable only when byol license is selected. Enum options - MARKETPLACE, IMAGE. Field introduced in 18.1.4, 18.2.1.
* `se_offline_del` - Unit is sec.
* `se_spawn_retry_interval` - Default retry period before attempting another service engine spawn in se group. Field introduced in 20.1.1. Unit is sec.
* `se_vnic_cooldown` - Unit is sec.
* `se_vnic_gc_wait_time` - Duration to wait after last vnic addition before proceeding with vnic garbage collection. Used for testing purposes. Field introduced in 20.1.4. Unit is sec.
* `secure_channel_cleanup_timeout` - Period for secure channel cleanup job. Unit is min.
* `secure_channel_controller_token_timeout` - Unit is min.
* `secure_channel_se_token_timeout` - Unit is min.
* `seupgrade_copy_pool_size` - This parameter defines the number of simultaneous se image downloads in a segroup. It is used to pace the se downloads so that controller network/cpu bandwidth is a bounded operation. A value of 0 will disable the pacing scheme and all the se(s) in the segroup will attempt to download the image. Field introduced in 18.2.6.
* `seupgrade_fabric_pool_size` - Pool size used for all fabric commands during se upgrade.
* `seupgrade_segroup_min_dead_timeout` - Time to wait before marking segroup upgrade as stuck. Unit is sec.
* `shared_ssl_certificates` - Ssl certificates in the admin tenant can be used in non-admin tenants. Field introduced in 18.2.5.
* `ssl_certificate_expiry_warning_days` - Number of days for ssl certificate expiry warning. Unit is days.
* `unresponsive_se_reboot` - Unit is sec.
* `upgrade_dns_ttl` - Time to account for dns ttl during upgrade. This is in addition to vs_scalein_timeout_for_upgrade in se_group. Field introduced in 17.1.1. Unit is sec. Allowed in basic(allowed values- 5) edition, essentials(allowed values- 5) edition, enterprise edition.
* `upgrade_fat_se_lease_time` - Amount of time controller waits for a large-sized se (>=128gb memory) to reconnect after it is rebooted during upgrade. Field introduced in 18.2.10, 20.1.1. Unit is sec.
* `upgrade_lease_time` - Amount of time controller waits for a regular-sized se (<128gb memory) to reconnect after it is rebooted during upgrade. Starting 18.2.10/20.1.1, the default time has increased from 360 seconds to 600 seconds. Unit is sec.
* `upgrade_se_per_vs_scale_ops_txn_time` - This parameter defines the upper-bound value of the vs scale-in or vs scale-out operation executed in the sescalein and sescale context. User can tweak this parameter to a higher value if the segroup gets suspended due to sescalein or sescaleout timeout failure typically associated with high number of vs(es) scaled out. Field introduced in 18.2.10, 20.1.1. Unit is sec.
* `user_agent_cache_config` - Configuration for user-agent cache used in bot management. Field introduced in 21.1.1.
* `uuid` - Unique object identifier of the object.
* `vnic_op_fail_time` - Unit is sec.
* `vs_apic_scaleout_timeout` - Time to wait for the scaled out se to become ready before marking the scaleout done, applies to apic configuration only. Unit is sec.
* `vs_awaiting_se_timeout` - Unit is sec.
* `vs_key_rotate_period` - Period for rotate vs keys job. Allowed values are 1-1051200. Special values are 0 - 'disabled'. Unit is min.
* `vs_scaleout_ready_check_interval` - Interval for checking scaleout_ready status while controller is waiting for scaleoutready rpc from the service engine. Field introduced in 18.2.2. Unit is sec.
* `vs_se_attach_ip_fail` - Time to wait before marking attach ip operation on an se as failed. Field introduced in 17.2.2. Unit is sec.
* `vs_se_bootup_fail` - Unit is sec.
* `vs_se_create_fail` - Unit is sec.
* `vs_se_ping_fail` - Unit is sec.
* `vs_se_vnic_fail` - Unit is sec.
* `vs_se_vnic_ip_fail` - Unit is sec.
* `warmstart_se_reconnect_wait_time` - Unit is sec.
* `warmstart_vs_resync_wait_time` - Timeout for warmstart vs resync. Field introduced in 18.1.4, 18.2.1. Unit is sec.

