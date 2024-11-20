package features

//lint:file-ignore U1000 we want to introduce this feature flag unused.

var (
	// SourcedAutogeneratedIntegrations enables adding a "source" to autogenerated integrations.
	// TODO(ROX-20353): if this is enabled by default, make sure to address sensor reconciliation of ImageIntegration resources.
	SourcedAutogeneratedIntegrations = registerFeature("Enable autogenerated integrations with cluster/namespace/secret source", "ROX_SOURCED_AUTOGENERATED_INTEGRATIONS", unchangeableInProd)

	// VulnMgmtWorkloadCVEs enables APIs and UI pages for the VM Workload CVE enhancements
	VulnMgmtWorkloadCVEs = registerFeature("Vuln Mgmt Workload CVEs", "ROX_VULN_MGMT_WORKLOAD_CVES", enabled, unchangeableInProd)

	// StoreEventHashes stores the hashes of successfully processed objects we receive from Sensor into the database
	StoreEventHashes = registerFeature("Store Event Hashes", "ROX_STORE_EVENT_HASHES", enabled, unchangeableInProd)

	// PreventSensorRestartOnDisconnect enables a new behavior in Sensor where it avoids restarting when the gRPC connection with Central ends.
	PreventSensorRestartOnDisconnect = registerFeature("Prevent Sensor restart on disconnect", "ROX_PREVENT_SENSOR_RESTART_ON_DISCONNECT", enabled, unchangeableInProd)

	// ComplianceEnhancements enables APIs and UI pages for Compliance 2.0
	ComplianceEnhancements = registerFeature("Compliance enhancements", "ROX_COMPLIANCE_ENHANCEMENTS", enabled)

	// ActiveVulnMgmt defines if the active vuln mgmt feature is enabled
	ActiveVulnMgmt = registerFeature("Enable Active Vulnerability Management", "ROX_ACTIVE_VULN_MGMT")

	// UnifiedCVEDeferral enables APIs and UI pages for unified deferral workflow.
	UnifiedCVEDeferral = registerFeature("Enable new unified Vulnerability deferral workflow", "ROX_VULN_MGMT_UNIFIED_CVE_DEFERRAL", enabled)

	// ClusterAwareDeploymentCheck enables roxctl deployment check to check deployments on the cluster level.
	ClusterAwareDeploymentCheck = registerFeature("Enables cluster level check for the 'roxctl deployment check' command.", "ROX_CLUSTER_AWARE_DEPLOYMENT_CHECK", enabled)

	// SensorReconciliationOnReconnect enables sensors to support reconciliation when reconnecting
	SensorReconciliationOnReconnect = registerFeature("Enable Sensors to support reconciliation on reconnect", "ROX_SENSOR_RECONCILIATION", enabled)

	// AuthMachineToMachine allows to exchange ID tokens for Central tokens without requiring user interaction.
	AuthMachineToMachine = registerFeature("Enable Auth Machine to Machine functionalities", "ROX_AUTH_MACHINE_TO_MACHINE", enabled)

	// PolicyCriteriaModal enables a modal for selecting policy criteria when editing a policy
	PolicyCriteriaModal = registerFeature("Enable modal to select policy criteria when editing a policy", "ROX_POLICY_CRITERIA_MODAL")

	// SensorDeploymentBuildOptimization enables a performance improvement by skipping deployments processing when no dependency or spec changed
	SensorDeploymentBuildOptimization = registerFeature("Enables a performance improvement by skipping deployments processing when no dependency or spec changed", "ROX_DEPLOYMENT_BUILD_OPTIMIZATION", enabled)

	// SensorCapturesIntermediateEvents enables sensor to capture intermediate events when it is disconnected from central
	SensorCapturesIntermediateEvents = registerFeature("Enables sensor to capture intermediate events when it is disconnected from central", "ROX_CAPTURE_INTERMEDIATE_EVENTS", enabled)

	// ScannerV4 indicates Scanner V4 is installed and should be used as the default image scanner in Central/Sensor.
	ScannerV4 = registerFeature("Enables Scanner V4 runtime functionality", "ROX_SCANNER_V4")

	// ScannerV4MultiBundle enables Scanner V4 to consume vulnerabilities using multi-bundle archives.
	ScannerV4MultiBundle = registerFeature("Enables Scanner V4 to consume vulnerabilities using multi-bundle archives", "ROX_SCANNER_V4_MULTI_BUNDLE", enabled)

	// VulnMgmtLegacySnooze enables APIs and UI for the legacy VM 1.0 "snooze CVE" functionality in the new VM 2.0 sections
	VulnMgmtLegacySnooze = registerFeature("Enables the ability to snooze Node and Platform CVEs in VM 2.0", "ROX_VULN_MGMT_LEGACY_SNOOZE")

	// ComplianceReporting enables support for compliance reporting.
	ComplianceReporting = registerFeature("Enable support for V2 compliance reporting", "ROX_COMPLIANCE_REPORTING", enabled)

	// UnqualifiedSearchRegistries enables support for unqualified search registries and short name aliases.
	UnqualifiedSearchRegistries = registerFeature("Enable support for unqualified search registries and short name aliases", "ROX_UNQUALIFIED_SEARCH_REGISTRIES")

	// ComplianceRemediationV2 enables the remediation feature of the compliance v2 integration.
	ComplianceRemediationV2 = registerFeature("Enable Compliance Remediation feature", "ROX_COMPLIANCE_REMEDIATION", enabled)

	// SensorAggregateDeploymentReferenceOptimization enables a performance improvement by aggregating deployment references when the same reference is queued for processing
	SensorAggregateDeploymentReferenceOptimization = registerFeature("Enables a performance improvement by aggregating deployment references when the same reference is queued for processing", "ROX_AGGREGATE_DEPLOYMENT_REFERENCE_OPTIMIZATION")

	// ACSCSEmailNotifier enables support for the ACSCS email notifier integratioon
	ACSCSEmailNotifier = registerFeature("Enable support for ACSCS Email notifier type", "ROX_ACSCS_EMAIL_NOTIFIER", enabled)

	// AttemptManifestDigest enables attempting to pull manifest digests from registres that historically did not
	// support it but now appear to (ie: Nexus and RHEL).
	AttemptManifestDigest = registerFeature("Enables attempts to pull manifest digests for all registry integrations", "ROX_ATTEMPT_MANIFEST_DIGEST", enabled)

	// VulnMgmt2GA enables support for migration changes for VM 2.0 GA
	VulnMgmt2GA = registerFeature("Enables support for migration changes for VM 2.0 GA", "ROX_VULN_MGMT_2_GA", enabled)

	// VulnMgmtAdvancedFilters enables support for advanced filters for VM 2.0 GA
	VulnMgmtAdvancedFilters = registerFeature("Enables support for advanced filters for VM 2.0 GA", "ROX_VULN_MGMT_ADVANCED_FILTERS", enabled)

	// DelegateWatchedImageReprocessing when set to enabled reprocessing of watched images may be delegated to secured clusters based
	// on the delegated scanning config.
	DelegateWatchedImageReprocessing = registerFeature("Enables delegating scans for watched images during reprocessing", "ROX_DELEGATE_WATCHED_IMAGE_REPROCESSING", enabled)

	// CollectorRuntimeConfig enables filtering of runtime events.
	CollectorRuntimeConfig = registerFeature("Enable collector runtime configuration", "ROX_COLLECTOR_RUNTIME_CONFIG")

	// SensorSingleScanPerImage when set to enabled forces Sensor to allow only a single scan per image to be active at any given
	// time. Will only have an affect if UnqualifiedSearchRegistries is also enabled.
	// TODO(ROX-24641): Remove dependency on the UnqualifiedSearchRegistries feature so that this is enabled by default.
	SensorSingleScanPerImage = registerFeature("Sensor will only allow a single active scan per image", "ROX_SENSOR_SINGLE_SCAN", enabled)

	// MicrosoftSentinelNotifier enables the Microsoft Sentinel notifier.
	MicrosoftSentinelNotifier = registerFeature("Enable the Microsoft Sentinel notifier", "ROX_MICROSOFT_SENTINEL", enabled)

	// ScanScheduleReportJobs enables support for compliance scan schedule report jobs
	ScanScheduleReportJobs = registerFeature("Enables support for compliance scan schedule report jobs", "ROX_SCAN_SCHEDULE_REPORT_JOBS", enabled)

	// PlatformComponents introduces the concept of platform collections and filtered views across the app.
	PlatformComponents = registerFeature("Introduce the concept of platform collections and filtered views across the app", "ROX_PLATFORM_COMPONENTS", enabled)

	// Display NVD CVSS score in UI.
	NvdCvssUI = registerFeature("Display NVD CVSS score in UI", "ROX_NVD_CVSS_UI", enabled)

	// Display clusters page patternfly redesign.
	ClustersPageMigrationUI = registerFeature("Display clusters page patternfly redesign", "ROX_CLUSTERS_PAGE_MIGRATION_UI")

	// ClusterRegistrationSecrets enables support for Cluster Registration Secrets (CRS), the next-gen init-bundles.
	ClusterRegistrationSecrets = registerFeature("Enable support for Cluster Registration Secrets (CRS)", "ROX_CLUSTER_REGISTRATION_SECRETS", enabled)

	// SensorPullSecretsByName when set to enabled will cause Sensor to capture pull secrets by secret name and registry host instead of just
	// registry host.
	SensorPullSecretsByName = registerFeature("Sensor will capture pull secrets by name and registry host instead of just registry host", "ROX_SENSOR_PULL_SECRETS_BY_NAME", enabled)

	// ExternalIPs enables storing detailed discovered external IPs
	ExternalIPs = registerFeature("Central will work with discovered external IPs", "ROX_EXTERNAL_IPS")

	// NetworkGraphExternalIPs enables displaying external (discovered) entities in the network graph
	NetworkGraphExternalIPs = registerFeature("Enables display of external IPs in the network graph UI", "ROX_NETWORK_GRAPH_EXTERNAL_IPS")

	// ScannerV4RedHatCVEs enables displaying CVEs instead of RHSAs/RHEAs/RHBAs in the place of fixed vulnerabilities affected Red Hat products.
	// TODO(ROX-26672): Remove this once we can show both CVEs and RHSAs in the UI + reports.
	ScannerV4RedHatCVEs = registerFeature("Scanner V4 will output CVEs instead of RHSAs/RHBAs/RHEAs for fixed Red Hat vulnerabilities", "ROX_SCANNER_V4_RED_HAT_CVES")

	// ScannerV4ReIndex enables Scanner V4 manifest re-indexing.
	ScannerV4ReIndex = registerFeature("Scanner V4 will re-index and delete unused manifests", "ROX_SCANNER_V4_REINDEX", enabled)

	// Display RHSA/RHBA/RHEA advisory separately from associated CVE.
	CVEAdvisorySeparation = registerFeature("Display RHSA/RHBA/RHEA advisory separately from associated CVE", "ROX_CVE_ADVISORY_SEPARATION")
)
