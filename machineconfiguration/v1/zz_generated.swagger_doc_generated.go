package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ContainerRuntimeConfig = map[string]string{
	"": "ContainerRuntimeConfig describes a customized Container Runtime configuration.",
}

func (ContainerRuntimeConfig) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfig
}

var map_ContainerRuntimeConfigCondition = map[string]string{
	"":                   "ContainerRuntimeConfigCondition defines the state of the ContainerRuntimeConfig",
	"type":               "type specifies the state of the operator's reconciliation functionality.",
	"status":             "status of the condition, one of True, False, Unknown.",
	"lastTransitionTime": "lastTransitionTime is the time of the last update to the current status object.",
	"reason":             "reason is the reason for the condition's last transition.  Reasons are PascalCase",
	"message":            "message provides additional information about the current condition. This is only to be consumed by humans.",
}

func (ContainerRuntimeConfigCondition) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfigCondition
}

var map_ContainerRuntimeConfigList = map[string]string{
	"": "ContainerRuntimeConfigList is a list of ContainerRuntimeConfig resources",
}

func (ContainerRuntimeConfigList) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfigList
}

var map_ContainerRuntimeConfigSpec = map[string]string{
	"": "ContainerRuntimeConfigSpec defines the desired state of ContainerRuntimeConfig",
}

func (ContainerRuntimeConfigSpec) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfigSpec
}

var map_ContainerRuntimeConfigStatus = map[string]string{
	"":                   "ContainerRuntimeConfigStatus defines the observed state of a ContainerRuntimeConfig",
	"observedGeneration": "observedGeneration represents the generation observed by the controller.",
	"conditions":         "conditions represents the latest available observations of current state.",
}

func (ContainerRuntimeConfigStatus) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfigStatus
}

var map_ContainerRuntimeConfiguration = map[string]string{
	"":            "ContainerRuntimeConfiguration defines the tuneables of the container runtime",
	"pidsLimit":   "pidsLimit specifies the maximum number of processes allowed in a container",
	"logLevel":    "logLevel specifies the verbosity of the logs based on the level it is set to. Options are fatal, panic, error, warn, info, and debug.",
	"logSizeMax":  "logSizeMax specifies the Maximum size allowed for the container log file. Negative numbers indicate that no size limit is imposed. If it is positive, it must be >= 8192 to match/exceed conmon's read buffer.",
	"overlaySize": "overlaySize specifies the maximum size of a container image. This flag can be used to set quota on the size of container images. (default: 10GB)",
}

func (ContainerRuntimeConfiguration) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfiguration
}

var map_KubeletConfig = map[string]string{
	"": "KubeletConfig describes a customized Kubelet configuration.",
}

func (KubeletConfig) SwaggerDoc() map[string]string {
	return map_KubeletConfig
}

var map_KubeletConfigCondition = map[string]string{
	"":                   "KubeletConfigCondition defines the state of the KubeletConfig",
	"type":               "type specifies the state of the operator's reconciliation functionality.",
	"status":             "status of the condition, one of True, False, Unknown.",
	"lastTransitionTime": "lastTransitionTime is the time of the last update to the current status object.",
	"reason":             "reason is the reason for the condition's last transition.  Reasons are PascalCase",
	"message":            "message provides additional information about the current condition. This is only to be consumed by humans.",
}

func (KubeletConfigCondition) SwaggerDoc() map[string]string {
	return map_KubeletConfigCondition
}

var map_KubeletConfigList = map[string]string{
	"": "KubeletConfigList is a list of KubeletConfig resources",
}

func (KubeletConfigList) SwaggerDoc() map[string]string {
	return map_KubeletConfigList
}

var map_KubeletConfigSpec = map[string]string{
	"": "KubeletConfigSpec defines the desired state of KubeletConfig",
}

func (KubeletConfigSpec) SwaggerDoc() map[string]string {
	return map_KubeletConfigSpec
}

var map_KubeletConfigStatus = map[string]string{
	"":                   "KubeletConfigStatus defines the observed state of a KubeletConfig",
	"observedGeneration": "observedGeneration represents the generation observed by the controller.",
	"conditions":         "conditions represents the latest available observations of current state.",
}

func (KubeletConfigStatus) SwaggerDoc() map[string]string {
	return map_KubeletConfigStatus
}

var map_MachineConfig = map[string]string{
	"": "MachineConfig defines the configuration for a machine",
}

func (MachineConfig) SwaggerDoc() map[string]string {
	return map_MachineConfig
}

var map_MachineConfigList = map[string]string{
	"": "MachineConfigList is a list of MachineConfig resources",
}

func (MachineConfigList) SwaggerDoc() map[string]string {
	return map_MachineConfigList
}

var map_MachineConfigPool = map[string]string{
	"": "MachineConfigPool describes a pool of MachineConfigs.",
}

func (MachineConfigPool) SwaggerDoc() map[string]string {
	return map_MachineConfigPool
}

var map_MachineConfigPoolCondition = map[string]string{
	"":                   "MachineConfigPoolCondition contains condition information for an MachineConfigPool.",
	"type":               "type of the condition, currently ('Done', 'Updating', 'Failed').",
	"status":             "status of the condition, one of ('True', 'False', 'Unknown').",
	"lastTransitionTime": "lastTransitionTime is the timestamp corresponding to the last status change of this condition.",
	"reason":             "reason is a brief machine readable explanation for the condition's last transition.",
	"message":            "message is a human readable description of the details of the last transition, complementing reason.",
}

func (MachineConfigPoolCondition) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolCondition
}

var map_MachineConfigPoolList = map[string]string{
	"": "MachineConfigPoolList is a list of MachineConfigPool resources",
}

func (MachineConfigPoolList) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolList
}

var map_MachineConfigPoolSpec = map[string]string{
	"":                      "MachineConfigPoolSpec is the spec for MachineConfigPool resource.",
	"machineConfigSelector": "machineConfigSelector specifies a label selector for MachineConfigs. Refer https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ on how label and selectors work.",
	"nodeSelector":          "nodeSelector specifies a label selector for Machines",
	"paused":                "paused specifies whether or not changes to this machine config pool should be stopped. This includes generating new desiredMachineConfig and update of machines.",
	"maxUnavailable":        "maxUnavailable specifies the percentage or constant number of machines that can be updating at any given time. default is 1.",
	"configuration":         "The targeted MachineConfig object for the machine config pool.",
}

func (MachineConfigPoolSpec) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolSpec
}

var map_MachineConfigPoolStatus = map[string]string{
	"":                        "MachineConfigPoolStatus is the status for MachineConfigPool resource.",
	"observedGeneration":      "observedGeneration represents the generation observed by the controller.",
	"configuration":           "configuration represents the current MachineConfig object for the machine config pool.",
	"machineCount":            "machineCount represents the total number of machines in the machine config pool.",
	"updatedMachineCount":     "updatedMachineCount represents the total number of machines targeted by the pool that have the CurrentMachineConfig as their config.",
	"readyMachineCount":       "readyMachineCount represents the total number of ready machines targeted by the pool.",
	"unavailableMachineCount": "unavailableMachineCount represents the total number of unavailable (non-ready) machines targeted by the pool. A node is marked unavailable if it is in updating state or NodeReady condition is false.",
	"degradedMachineCount":    "degradedMachineCount represents the total number of machines marked degraded (or unreconcilable). A node is marked degraded if applying a configuration failed..",
	"conditions":              "conditions represents the latest available observations of current state.",
}

func (MachineConfigPoolStatus) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolStatus
}

var map_MachineConfigPoolStatusConfiguration = map[string]string{
	"":       "MachineConfigPoolStatusConfiguration stores the current configuration for the pool, and optionally also stores the list of MachineConfig objects used to generate the configuration.",
	"source": "source is the list of MachineConfig objects that were used to generate the single MachineConfig object specified in `content`.",
}

func (MachineConfigPoolStatusConfiguration) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolStatusConfiguration
}

var map_MachineConfigSpec = map[string]string{
	"":           "MachineConfigSpec is the spec for MachineConfig",
	"osImageURL": "OSImageURL specifies the remote location that will be used to fetch the OS.",
	"config":     "Config is a Ignition Config object.",
}

func (MachineConfigSpec) SwaggerDoc() map[string]string {
	return map_MachineConfigSpec
}

// AUTO-GENERATED FUNCTIONS END HERE
