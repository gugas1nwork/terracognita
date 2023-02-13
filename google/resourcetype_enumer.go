// Code generated by "enumer -type ResourceType -addprefix google_ -transform snake -linecomment"; DO NOT EDIT.

package google

import (
	"fmt"
	"strings"
)

const _ResourceTypeName = "google_compute_instancegoogle_compute_firewallgoogle_compute_networkgoogle_compute_health_checkgoogle_compute_instance_groupgoogle_compute_instance_iam_policygoogle_compute_backend_bucketgoogle_compute_backend_servicegoogle_compute_ssl_certificategoogle_compute_target_http_proxygoogle_compute_target_https_proxygoogle_compute_url_mapgoogle_compute_global_forwarding_rulegoogle_compute_forwarding_rulegoogle_compute_diskgoogle_compute_addressgoogle_compute_attached_diskgoogle_compute_autoscalergoogle_compute_global_addressgoogle_compute_imagegoogle_compute_instance_group_managergoogle_compute_instance_templategoogle_compute_managed_ssl_certificategoogle_compute_network_endpoint_groupgoogle_compute_routegoogle_compute_security_policygoogle_compute_service_attachmentgoogle_compute_snapshotgoogle_compute_ssl_policygoogle_compute_subnetworkgoogle_compute_target_grpc_proxygoogle_compute_target_instancegoogle_compute_target_poolgoogle_compute_target_ssl_proxygoogle_compute_target_tcp_proxygoogle_dns_managed_zonegoogle_dns_record_setgoogle_dns_policygoogle_project_iam_custom_rolegoogle_billing_subaccountgoogle_sql_database_instancegoogle_sql_databasegoogle_storage_bucketgoogle_storage_bucket_iam_policygoogle_filestore_instancegoogle_container_clustergoogle_container_node_poolgoogle_redis_instancegoogle_logging_metricgoogle_monitoring_alert_policygoogle_monitoring_groupgoogle_monitoring_notification_channelgoogle_monitoring_uptime_check_config"

var _ResourceTypeIndex = [...]uint16{0, 23, 46, 68, 95, 124, 158, 187, 217, 247, 279, 312, 334, 371, 401, 420, 442, 470, 495, 524, 544, 581, 613, 651, 688, 708, 738, 771, 794, 819, 844, 876, 906, 932, 963, 994, 1017, 1038, 1055, 1085, 1110, 1138, 1157, 1178, 1210, 1235, 1259, 1285, 1306, 1327, 1357, 1380, 1418, 1455}

const _ResourceTypeLowerName = "google_compute_instancegoogle_compute_firewallgoogle_compute_networkgoogle_compute_health_checkgoogle_compute_instance_groupgoogle_compute_instance_iam_policygoogle_compute_backend_bucketgoogle_compute_backend_servicegoogle_compute_ssl_certificategoogle_compute_target_http_proxygoogle_compute_target_https_proxygoogle_compute_url_mapgoogle_compute_global_forwarding_rulegoogle_compute_forwarding_rulegoogle_compute_diskgoogle_compute_addressgoogle_compute_attached_diskgoogle_compute_autoscalergoogle_compute_global_addressgoogle_compute_imagegoogle_compute_instance_group_managergoogle_compute_instance_templategoogle_compute_managed_ssl_certificategoogle_compute_network_endpoint_groupgoogle_compute_routegoogle_compute_security_policygoogle_compute_service_attachmentgoogle_compute_snapshotgoogle_compute_ssl_policygoogle_compute_subnetworkgoogle_compute_target_grpc_proxygoogle_compute_target_instancegoogle_compute_target_poolgoogle_compute_target_ssl_proxygoogle_compute_target_tcp_proxygoogle_dns_managed_zonegoogle_dns_record_setgoogle_dns_policygoogle_project_iam_custom_rolegoogle_billing_subaccountgoogle_sql_database_instancegoogle_sql_databasegoogle_storage_bucketgoogle_storage_bucket_iam_policygoogle_filestore_instancegoogle_container_clustergoogle_container_node_poolgoogle_redis_instancegoogle_logging_metricgoogle_monitoring_alert_policygoogle_monitoring_groupgoogle_monitoring_notification_channelgoogle_monitoring_uptime_check_config"

func (i ResourceType) String() string {
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ResourceTypeNoOp() {
	var x [1]struct{}
	_ = x[ComputeInstance-(0)]
	_ = x[ComputeFirewall-(1)]
	_ = x[ComputeNetwork-(2)]
	_ = x[ComputeHealthCheck-(3)]
	_ = x[ComputeInstanceGroup-(4)]
	_ = x[ComputeInstanceIAMPolicy-(5)]
	_ = x[ComputeBackendBucket-(6)]
	_ = x[ComputeBackendService-(7)]
	_ = x[ComputeSSLCertificate-(8)]
	_ = x[ComputeTargetHTTPProxy-(9)]
	_ = x[ComputeTargetHTTPSProxy-(10)]
	_ = x[ComputeURLMap-(11)]
	_ = x[ComputeGlobalForwardingRule-(12)]
	_ = x[ComputeForwardingRule-(13)]
	_ = x[ComputeDisk-(14)]
	_ = x[ComputeAddress-(15)]
	_ = x[ComputeAttachedDisk-(16)]
	_ = x[ComputeAutoscaler-(17)]
	_ = x[ComputeGlobalAddress-(18)]
	_ = x[ComputeImage-(19)]
	_ = x[ComputeInstanceGroupManager-(20)]
	_ = x[ComputeInstanceTemplate-(21)]
	_ = x[ComputeManagedSSLCertificate-(22)]
	_ = x[ComputeNetworkEndpointGroup-(23)]
	_ = x[ComputeRoute-(24)]
	_ = x[ComputeSecurityPolicy-(25)]
	_ = x[ComputeServiceAttachment-(26)]
	_ = x[ComputeSnapshot-(27)]
	_ = x[ComputeSSLPolicy-(28)]
	_ = x[ComputeSubnetwork-(29)]
	_ = x[ComputeTargetGRPCProxy-(30)]
	_ = x[ComputeTargetInstance-(31)]
	_ = x[ComputeTargetPool-(32)]
	_ = x[ComputeTargetSSLProxy-(33)]
	_ = x[ComputeTargetTCPProxy-(34)]
	_ = x[DNSManagedZone-(35)]
	_ = x[DNSRecordSet-(36)]
	_ = x[DNSPolicy-(37)]
	_ = x[ProjectIAMCustomRole-(38)]
	_ = x[BillingSubaccount-(39)]
	_ = x[SQLDatabaseInstance-(40)]
	_ = x[SQLDatabase-(41)]
	_ = x[StorageBucket-(42)]
	_ = x[StorageBucketIAMPolicy-(43)]
	_ = x[FilestoreInstance-(44)]
	_ = x[ContainerCluster-(45)]
	_ = x[ContainerNodePool-(46)]
	_ = x[RedisInstance-(47)]
	_ = x[LoggingMetric-(48)]
	_ = x[MonitoringAlertPolicy-(49)]
	_ = x[MonitoringGroup-(50)]
	_ = x[MonitoringNotificationChannel-(51)]
	_ = x[MonitoringUptimeCheckConfig-(52)]
}

var _ResourceTypeValues = []ResourceType{ComputeInstance, ComputeFirewall, ComputeNetwork, ComputeHealthCheck, ComputeInstanceGroup, ComputeInstanceIAMPolicy, ComputeBackendBucket, ComputeBackendService, ComputeSSLCertificate, ComputeTargetHTTPProxy, ComputeTargetHTTPSProxy, ComputeURLMap, ComputeGlobalForwardingRule, ComputeForwardingRule, ComputeDisk, ComputeAddress, ComputeAttachedDisk, ComputeAutoscaler, ComputeGlobalAddress, ComputeImage, ComputeInstanceGroupManager, ComputeInstanceTemplate, ComputeManagedSSLCertificate, ComputeNetworkEndpointGroup, ComputeRoute, ComputeSecurityPolicy, ComputeServiceAttachment, ComputeSnapshot, ComputeSSLPolicy, ComputeSubnetwork, ComputeTargetGRPCProxy, ComputeTargetInstance, ComputeTargetPool, ComputeTargetSSLProxy, ComputeTargetTCPProxy, DNSManagedZone, DNSRecordSet, DNSPolicy, ProjectIAMCustomRole, BillingSubaccount, SQLDatabaseInstance, SQLDatabase, StorageBucket, StorageBucketIAMPolicy, FilestoreInstance, ContainerCluster, ContainerNodePool, RedisInstance, LoggingMetric, MonitoringAlertPolicy, MonitoringGroup, MonitoringNotificationChannel, MonitoringUptimeCheckConfig}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:23]:           ComputeInstance,
	_ResourceTypeLowerName[0:23]:      ComputeInstance,
	_ResourceTypeName[23:46]:          ComputeFirewall,
	_ResourceTypeLowerName[23:46]:     ComputeFirewall,
	_ResourceTypeName[46:68]:          ComputeNetwork,
	_ResourceTypeLowerName[46:68]:     ComputeNetwork,
	_ResourceTypeName[68:95]:          ComputeHealthCheck,
	_ResourceTypeLowerName[68:95]:     ComputeHealthCheck,
	_ResourceTypeName[95:124]:         ComputeInstanceGroup,
	_ResourceTypeLowerName[95:124]:    ComputeInstanceGroup,
	_ResourceTypeName[124:158]:        ComputeInstanceIAMPolicy,
	_ResourceTypeLowerName[124:158]:   ComputeInstanceIAMPolicy,
	_ResourceTypeName[158:187]:        ComputeBackendBucket,
	_ResourceTypeLowerName[158:187]:   ComputeBackendBucket,
	_ResourceTypeName[187:217]:        ComputeBackendService,
	_ResourceTypeLowerName[187:217]:   ComputeBackendService,
	_ResourceTypeName[217:247]:        ComputeSSLCertificate,
	_ResourceTypeLowerName[217:247]:   ComputeSSLCertificate,
	_ResourceTypeName[247:279]:        ComputeTargetHTTPProxy,
	_ResourceTypeLowerName[247:279]:   ComputeTargetHTTPProxy,
	_ResourceTypeName[279:312]:        ComputeTargetHTTPSProxy,
	_ResourceTypeLowerName[279:312]:   ComputeTargetHTTPSProxy,
	_ResourceTypeName[312:334]:        ComputeURLMap,
	_ResourceTypeLowerName[312:334]:   ComputeURLMap,
	_ResourceTypeName[334:371]:        ComputeGlobalForwardingRule,
	_ResourceTypeLowerName[334:371]:   ComputeGlobalForwardingRule,
	_ResourceTypeName[371:401]:        ComputeForwardingRule,
	_ResourceTypeLowerName[371:401]:   ComputeForwardingRule,
	_ResourceTypeName[401:420]:        ComputeDisk,
	_ResourceTypeLowerName[401:420]:   ComputeDisk,
	_ResourceTypeName[420:442]:        ComputeAddress,
	_ResourceTypeLowerName[420:442]:   ComputeAddress,
	_ResourceTypeName[442:470]:        ComputeAttachedDisk,
	_ResourceTypeLowerName[442:470]:   ComputeAttachedDisk,
	_ResourceTypeName[470:495]:        ComputeAutoscaler,
	_ResourceTypeLowerName[470:495]:   ComputeAutoscaler,
	_ResourceTypeName[495:524]:        ComputeGlobalAddress,
	_ResourceTypeLowerName[495:524]:   ComputeGlobalAddress,
	_ResourceTypeName[524:544]:        ComputeImage,
	_ResourceTypeLowerName[524:544]:   ComputeImage,
	_ResourceTypeName[544:581]:        ComputeInstanceGroupManager,
	_ResourceTypeLowerName[544:581]:   ComputeInstanceGroupManager,
	_ResourceTypeName[581:613]:        ComputeInstanceTemplate,
	_ResourceTypeLowerName[581:613]:   ComputeInstanceTemplate,
	_ResourceTypeName[613:651]:        ComputeManagedSSLCertificate,
	_ResourceTypeLowerName[613:651]:   ComputeManagedSSLCertificate,
	_ResourceTypeName[651:688]:        ComputeNetworkEndpointGroup,
	_ResourceTypeLowerName[651:688]:   ComputeNetworkEndpointGroup,
	_ResourceTypeName[688:708]:        ComputeRoute,
	_ResourceTypeLowerName[688:708]:   ComputeRoute,
	_ResourceTypeName[708:738]:        ComputeSecurityPolicy,
	_ResourceTypeLowerName[708:738]:   ComputeSecurityPolicy,
	_ResourceTypeName[738:771]:        ComputeServiceAttachment,
	_ResourceTypeLowerName[738:771]:   ComputeServiceAttachment,
	_ResourceTypeName[771:794]:        ComputeSnapshot,
	_ResourceTypeLowerName[771:794]:   ComputeSnapshot,
	_ResourceTypeName[794:819]:        ComputeSSLPolicy,
	_ResourceTypeLowerName[794:819]:   ComputeSSLPolicy,
	_ResourceTypeName[819:844]:        ComputeSubnetwork,
	_ResourceTypeLowerName[819:844]:   ComputeSubnetwork,
	_ResourceTypeName[844:876]:        ComputeTargetGRPCProxy,
	_ResourceTypeLowerName[844:876]:   ComputeTargetGRPCProxy,
	_ResourceTypeName[876:906]:        ComputeTargetInstance,
	_ResourceTypeLowerName[876:906]:   ComputeTargetInstance,
	_ResourceTypeName[906:932]:        ComputeTargetPool,
	_ResourceTypeLowerName[906:932]:   ComputeTargetPool,
	_ResourceTypeName[932:963]:        ComputeTargetSSLProxy,
	_ResourceTypeLowerName[932:963]:   ComputeTargetSSLProxy,
	_ResourceTypeName[963:994]:        ComputeTargetTCPProxy,
	_ResourceTypeLowerName[963:994]:   ComputeTargetTCPProxy,
	_ResourceTypeName[994:1017]:       DNSManagedZone,
	_ResourceTypeLowerName[994:1017]:  DNSManagedZone,
	_ResourceTypeName[1017:1038]:      DNSRecordSet,
	_ResourceTypeLowerName[1017:1038]: DNSRecordSet,
	_ResourceTypeName[1038:1055]:      DNSPolicy,
	_ResourceTypeLowerName[1038:1055]: DNSPolicy,
	_ResourceTypeName[1055:1085]:      ProjectIAMCustomRole,
	_ResourceTypeLowerName[1055:1085]: ProjectIAMCustomRole,
	_ResourceTypeName[1085:1110]:      BillingSubaccount,
	_ResourceTypeLowerName[1085:1110]: BillingSubaccount,
	_ResourceTypeName[1110:1138]:      SQLDatabaseInstance,
	_ResourceTypeLowerName[1110:1138]: SQLDatabaseInstance,
	_ResourceTypeName[1138:1157]:      SQLDatabase,
	_ResourceTypeLowerName[1138:1157]: SQLDatabase,
	_ResourceTypeName[1157:1178]:      StorageBucket,
	_ResourceTypeLowerName[1157:1178]: StorageBucket,
	_ResourceTypeName[1178:1210]:      StorageBucketIAMPolicy,
	_ResourceTypeLowerName[1178:1210]: StorageBucketIAMPolicy,
	_ResourceTypeName[1210:1235]:      FilestoreInstance,
	_ResourceTypeLowerName[1210:1235]: FilestoreInstance,
	_ResourceTypeName[1235:1259]:      ContainerCluster,
	_ResourceTypeLowerName[1235:1259]: ContainerCluster,
	_ResourceTypeName[1259:1285]:      ContainerNodePool,
	_ResourceTypeLowerName[1259:1285]: ContainerNodePool,
	_ResourceTypeName[1285:1306]:      RedisInstance,
	_ResourceTypeLowerName[1285:1306]: RedisInstance,
	_ResourceTypeName[1306:1327]:      LoggingMetric,
	_ResourceTypeLowerName[1306:1327]: LoggingMetric,
	_ResourceTypeName[1327:1357]:      MonitoringAlertPolicy,
	_ResourceTypeLowerName[1327:1357]: MonitoringAlertPolicy,
	_ResourceTypeName[1357:1380]:      MonitoringGroup,
	_ResourceTypeLowerName[1357:1380]: MonitoringGroup,
	_ResourceTypeName[1380:1418]:      MonitoringNotificationChannel,
	_ResourceTypeLowerName[1380:1418]: MonitoringNotificationChannel,
	_ResourceTypeName[1418:1455]:      MonitoringUptimeCheckConfig,
	_ResourceTypeLowerName[1418:1455]: MonitoringUptimeCheckConfig,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:23],
	_ResourceTypeName[23:46],
	_ResourceTypeName[46:68],
	_ResourceTypeName[68:95],
	_ResourceTypeName[95:124],
	_ResourceTypeName[124:158],
	_ResourceTypeName[158:187],
	_ResourceTypeName[187:217],
	_ResourceTypeName[217:247],
	_ResourceTypeName[247:279],
	_ResourceTypeName[279:312],
	_ResourceTypeName[312:334],
	_ResourceTypeName[334:371],
	_ResourceTypeName[371:401],
	_ResourceTypeName[401:420],
	_ResourceTypeName[420:442],
	_ResourceTypeName[442:470],
	_ResourceTypeName[470:495],
	_ResourceTypeName[495:524],
	_ResourceTypeName[524:544],
	_ResourceTypeName[544:581],
	_ResourceTypeName[581:613],
	_ResourceTypeName[613:651],
	_ResourceTypeName[651:688],
	_ResourceTypeName[688:708],
	_ResourceTypeName[708:738],
	_ResourceTypeName[738:771],
	_ResourceTypeName[771:794],
	_ResourceTypeName[794:819],
	_ResourceTypeName[819:844],
	_ResourceTypeName[844:876],
	_ResourceTypeName[876:906],
	_ResourceTypeName[906:932],
	_ResourceTypeName[932:963],
	_ResourceTypeName[963:994],
	_ResourceTypeName[994:1017],
	_ResourceTypeName[1017:1038],
	_ResourceTypeName[1038:1055],
	_ResourceTypeName[1055:1085],
	_ResourceTypeName[1085:1110],
	_ResourceTypeName[1110:1138],
	_ResourceTypeName[1138:1157],
	_ResourceTypeName[1157:1178],
	_ResourceTypeName[1178:1210],
	_ResourceTypeName[1210:1235],
	_ResourceTypeName[1235:1259],
	_ResourceTypeName[1259:1285],
	_ResourceTypeName[1285:1306],
	_ResourceTypeName[1306:1327],
	_ResourceTypeName[1327:1357],
	_ResourceTypeName[1357:1380],
	_ResourceTypeName[1380:1418],
	_ResourceTypeName[1418:1455],
}

// ResourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResourceTypeString(s string) (ResourceType, error) {
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResourceType values", s)
}

// ResourceTypeValues returns all values of the enum
func ResourceTypeValues() []ResourceType {
	return _ResourceTypeValues
}

// ResourceTypeStrings returns a slice of all String values of the enum
func ResourceTypeStrings() []string {
	strs := make([]string, len(_ResourceTypeNames))
	copy(strs, _ResourceTypeNames)
	return strs
}

// IsAResourceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResourceType) IsAResourceType() bool {
	for _, v := range _ResourceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
