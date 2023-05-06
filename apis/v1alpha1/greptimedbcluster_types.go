// Copyright 2022 Greptime Team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StorageRetainPolicyType string

const (
	// RetainStorageRetainPolicyTypeRetain is the default options.
	// The storage(PVCs) will be retained when the cluster is deleted.
	RetainStorageRetainPolicyTypeRetain StorageRetainPolicyType = "Retain"

	// RetainStorageRetainPolicyTypeDelete specify that the storage will be deleted when the associated StatefulSet delete.
	RetainStorageRetainPolicyTypeDelete StorageRetainPolicyType = "Delete"
)

// ComponentKind is the kind of the component in the cluster.
type ComponentKind string

const (
	FrontendComponentKind ComponentKind = "frontend"
	DatanodeComponentKind ComponentKind = "datanode"
	MetaComponentKind     ComponentKind = "meta"
)

// ClusterPhase define the phase of the cluster.
type ClusterPhase string

const (
	// ClusterStarting means the controller start to create cluster.
	ClusterStarting ClusterPhase = "Starting"

	// ClusterRunning means all the components of cluster is ready.
	ClusterRunning ClusterPhase = "Running"

	// ClusterError means some kind of error happen in reconcile.
	ClusterError ClusterPhase = "Error"

	// ClusterTerminating means the cluster is terminating.
	ClusterTerminating ClusterPhase = "Terminating"
)

// SlimPodSpec is a slimmed down version of corev1.PodSpec.
// Most of the fields in SlimPodSpec are copied from corev1.PodSpec.
type SlimPodSpec struct {
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// NodeSelector field is from 'corev1.PodSpec.NodeSelector'.
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// List of initialization containers belonging to the pod.
	// Init containers are executed in order prior to containers being started. If any
	// init container fails, the pod is considered to have failed and is handled according
	// to its restartPolicy. The name for an init container or normal container must be
	// unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling
	// by finding the highest request/limit for each resource type, and then using the max of
	// that value or the sum of the normal containers. Limits are applied to init containers
	// in a similar fashion.
	// Init containers cannot currently be added or removed.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	// InitContainers field is from 'corev1.PodSpec.InitContainers'.
	// +optional
	InitContainers []corev1.Container `json:"initContainers,omitempty"`

	// Restart policy for all containers within the pod.
	// One of Always, OnFailure, Never.
	// Default to Always.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	// RestartPolicy field is from 'corev1.PodSpec.RestartPolicy'.
	// +optional
	RestartPolicy corev1.RestartPolicy `json:"restartPolicy,omitempty"`

	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
	// Value must be non-negative integer. The value zero indicates stop immediately via
	// the kill signal (no opportunity to shut down).
	// If this value is nil, the default grace period will be used instead.
	// The grace period is the duration in seconds after the processes running in the pod are sent
	// a termination signal and the time when the processes are forcibly halted with a kill signal.
	// Set this value longer than the expected cleanup time for your process.
	// Defaults to 30 seconds.
	// TerminationGracePeriodSeconds field is from 'corev1.PodSpec.TerminationGracePeriodSeconds'.
	// +optional
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty"`

	// Optional duration in seconds the pod may be active on the node relative to
	// StartTime before the system will actively try to mark it failed and kill associated containers.
	// Value must be a positive integer.
	// ActiveDeadlineSeconds field is from 'corev1.PodSpec.ActiveDeadlineSeconds'.
	// +optional
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty"`

	// Set DNS policy for the pod.
	// Defaults to "ClusterFirst".
	// Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.
	// DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.
	// To have DNS options set along with hostNetwork, you have to specify DNS policy
	// explicitly to 'ClusterFirstWithHostNet'.
	// DNSPolicy field is from 'corev1.PodSpec.DNSPolicy'.
	// +optional
	DNSPolicy corev1.DNSPolicy `json:"dnsPolicy,omitempty"`

	// ServiceAccountName is the name of the ServiceAccount to use to run this pod.
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// ServiceAccountName field is from 'corev1.PodSpec.ServiceAccountName'.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// Host networking requested for this pod. Use the host's network namespace.
	// If this option is set, the ports that will be used must be specified.
	// Default to false.
	// HostNetwork field is from 'corev1.PodSpec.HostNetwork'.
	// +optional
	HostNetwork bool `json:"hostNetwork,omitempty"`

	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
	// If specified, these secrets will be passed to individual puller implementations for them to use.
	// More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	// ImagePullSecrets field is from 'corev1.PodSpec.ImagePullSecrets'.
	// +optional
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`

	// If specified, the pod's scheduling constraints
	// Affinity field is from 'corev1.PodSpec.Affinity'.
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// If specified, the pod will be dispatched by specified scheduler.
	// If not specified, the pod will be dispatched by default scheduler.
	// SchedulerName field is from 'corev1.PodSpec.SchedulerName'.
	// +optional
	SchedulerName string `json:"schedulerName,omitempty"`

	// For most time, there is one main container in a pod(frontend/meta/datanode).
	// If specified, additional containers will be added to the pod as sidecar containers.
	// +optional
	AdditionalContainers []corev1.Container `json:"additionalContainers,omitempty"`

	// List of volumes that can be mounted by containers belonging to the pod.
	// +optional
	Volumes []Volume `json:"volumes,omitempty"`
}

// Volume represents a named volume in a pod that may be accessed by any container in the pod.
type Volume struct {
	// name of the volume.
	// Must be a DNS_LABEL and unique within the pod.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// volumeSource represents the location and type of the mounted volume.
	// If not specified, the Volume is implied to be an EmptyDir.
	// This implied behavior is deprecated and will be removed in a future version.
	VolumeSource `json:",inline" protobuf:"bytes,2,opt,name=volumeSource"`
}

type VolumeSource struct {
	// secret represents a secret that should populate this volume.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	// +optional
	Secret *SecretVolumeSource `json:"secret,omitempty"`
}

// The contents of the target Secret's Data field will be presented in a volume
// as files using the keys in the Data field as the file names.
// Secret volumes support ownership management and SELinux relabeling.
type SecretVolumeSource struct {
	// secretName is the name of the secret in the pod's namespace to use.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	// +optional
	SecretName string `json:"secretName,omitempty" protobuf:"bytes,1,opt,name=secretName"`
	// items If unspecified, each key-value pair in the Data field of the referenced
	// Secret will be projected into the volume as a file whose name is the
	// key and content is the value. If specified, the listed keys will be
	// projected into the specified paths, and unlisted keys will not be
	// present. If a key is specified which is not present in the Secret,
	// the volume setup will error unless it is marked optional. Paths must be
	// relative and may not contain the '..' path or start with '..'.
	// +optional
	Items []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
	// defaultMode is Optional: mode bits used to set permissions on created files by default.
	// Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
	// YAML accepts both octal and decimal values, JSON requires decimal values
	// for mode bits. Defaults to 0644.
	// Directories within the path are not affected by this setting.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	DefaultMode *int32 `json:"defaultMode,omitempty" protobuf:"bytes,3,opt,name=defaultMode"`
	// optional field specify whether the Secret or its keys must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

// Maps a string key to a path within a volume.
type KeyToPath struct {
	// key is the key to project.
	Key string `json:"key" protobuf:"bytes,1,opt,name=key"`

	// path is the relative path of the file to map the key to.
	// May not be an absolute path.
	// May not contain the path element '..'.
	// May not start with the string '..'.
	Path string `json:"path" protobuf:"bytes,2,opt,name=path"`
	// mode is Optional: mode bits used to set permissions on this file.
	// Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
	// YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
	// If not specified, the volume defaultMode will be used.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	Mode *int32 `json:"mode,omitempty" protobuf:"varint,3,opt,name=mode"`
}

// MainContainerSpec describes the specification of the main container of a pod.
// Most of the fields of MainContainerSpec are from 'corev1.Container'.
type MainContainerSpec struct {
	// The main container image name of the component.
	// +required
	Image string `json:"image,omitempty"`

	// The resource requirements of the main container.
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// Entrypoint array. Not executed within a shell.
	// The container image's ENTRYPOINT is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
	// produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
	// of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// Command field is from 'corev1.Container.Command'.
	// +optional
	Command []string `json:"command,omitempty"`

	// Arguments to the entrypoint.
	// The container image's CMD is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
	// produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
	// of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// Args field is from 'corev1.Container.Args'.
	// +optional
	Args []string `json:"args,omitempty"`

	// Container's working directory.
	// If not specified, the container runtime's default will be used, which
	// might be configured in the container image.
	// Cannot be updated.
	// WorkingDir field is from 'corev1.Container.WorkingDir'.
	// +optional
	WorkingDir string `json:"workingDir,omitempty"`

	// List of environment variables to set in the container.
	// Cannot be updated.
	// Env field is from 'corev1.Container.Env'.
	// +optional
	Env []corev1.EnvVar `json:"env,omitempty"`

	// Periodic probe of container liveness.
	// Container will be restarted if the probe fails.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// LivenessProbe field is from 'corev1.Container.LivenessProbe'.
	// +optional
	LivenessProbe *corev1.Probe `json:"livenessProbe,omitempty"`

	// Periodic probe of container service readiness.
	// Container will be removed from service endpoints if the probe fails.
	// ReadinessProbe field is from 'corev1.Container.LivenessProbe'.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	ReadinessProbe *corev1.Probe `json:"readinessProbe,omitempty"`

	// Actions that the management system should take in response to container lifecycle events.
	// Cannot be updated.
	// Lifecycle field is from 'corev1.Container.Lifecycle'.
	// +optional
	Lifecycle *corev1.Lifecycle `json:"lifecycle,omitempty"`

	// Image pull policy.
	// One of Always, Never, IfNotPresent.
	// Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	// ImagePullPolicy field is from 'corev1.Container.ImagePullPolicy'.
	// +optional
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// Pod volumes to mount into the container's filesystem.
	// Cannot be updated.
	// +optional
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`
}

// VolumeMount describes a mounting of a Volume within a container.
type VolumeMount struct {
	// This must match the Name of a Volume.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	// Defaults to false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"`
	// Path within the container at which the volume should be mounted.  Must
	// not contain ':'.
	MountPath string `json:"mountPath" protobuf:"bytes,3,opt,name=mountPath"`
	// Path within the volume from which the container's volume should be mounted.
	// Defaults to "" (volume's root).
	// +optional
	SubPath string `json:"subPath,omitempty" protobuf:"bytes,4,opt,name=subPath"`
	// Expanded path within the volume from which the container's volume should be mounted.
	// Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.
	// Defaults to "" (volume's root).
	// SubPathExpr and SubPath are mutually exclusive.
	// +optional
	SubPathExpr string `json:"subPathExpr,omitempty" protobuf:"bytes,6,opt,name=subPathExpr"`
}

// PodTemplateSpec defines the template for a pod of cluster.
type PodTemplateSpec struct {
	// The annotations to be created to the pod.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// The labels to be created to the pod.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// MainContainer defines the specification of the main container of the pod.
	// +optional
	MainContainer *MainContainerSpec `json:"main,omitempty"`

	// SlimPodSpec defines the desired behavior of the pod.
	// +optional
	SlimPodSpec `json:",inline"`
}

// ComponentSpec is the common specification for all components(frontend/meta/datanode).
type ComponentSpec struct {
	// The number of replicas of the components.
	// +required
	// +kubebuilder:validation:Minimum=1
	Replicas int32 `json:"replicas"`

	// The content of the configuration file of the component in TOML format.
	// +optional
	Config string `json:"config,omitempty"`

	// Template defines the pod template for the component, if not specified, the pod template will use the default value.
	// +optional
	Template *PodTemplateSpec `json:"template,omitempty"`
}

// MetaSpec is the specification for meta component.
type MetaSpec struct {
	ComponentSpec `json:",inline"`

	// +optional
	ServicePort int32 `json:"servicePort,omitempty"`

	// +optional
	EtcdEndpoints []string `json:"etcdEndpoints,omitempty"`

	// More meta settings can be added here...
}

// StorageSpec will generate PVC.
type StorageSpec struct {
	// The name of the storage.
	// +optional
	Name string `json:"name,omitempty"`

	// The name of the storage class to use for the volume.
	// +optional
	StorageClassName *string `json:"storageClassName,omitempty"`

	// The size of the storage.
	// +optional
	// +kubebuilder:validation:Pattern=(^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$)
	StorageSize string `json:"storageSize,omitempty"`

	// The mount path of the storage in datanode container.
	// +optional
	MountPath string `json:"mountPath,omitempty"`

	// The PVCs will retain or delete when the cluster is deleted, default to Retain.
	// +optional
	// +kubebuilder:validation:Enum:={"Retain", "Delete"}
	// +kubebuilder:default:="Retain"
	StorageRetainPolicy StorageRetainPolicyType `json:"storageRetainPolicy,omitempty"`
}

type ServiceSpec struct {
	// type determines how the Service is exposed.
	// +optional
	Type corev1.ServiceType `json:"type,omitempty"`

	// Additional annotations for the service
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Additional labels for the service
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// loadBalancerClass is the class of the load balancer implementation this Service belongs to.
	// +optional
	LoadBalancerClass *string `json:"loadBalancerClass,omitempty"`
}

// FrontendSpec is the specification for frontend component.
type FrontendSpec struct {
	ComponentSpec `json:",inline"`

	// +optional
	Service ServiceSpec `json:"service,omitempty"`

	// The TLS configurations of the frontend.
	// +optional
	TLS *TLSSpec `json:"tls,omitempty"`

	// More frontend settings can be added here...
}

type TLSSpec struct {
	// The secret name of the TLS certificate, and it must be in the same namespace of the cluster.
	// The secret must contain keys named ca.crt, tls.crt and tls.key.
	// +optional
	SecretName string `json:"secretName,omitempty"`

	// The mouth path of certificate in frontend container.
	// +optional
	CertificateMountPath string `json:"certificateMountPath,omitempty"`
}

// DatanodeSpec is the specification for datanode component.
type DatanodeSpec struct {
	ComponentSpec `json:",inline"`

	Storage StorageSpec `json:"storage,omitempty"`
	// More datanode settings can be added here...
}

// InitializerSpec is the init container to set up components configurations before running the container.
type InitializerSpec struct {
	// +optional
	Image string `json:"image,omitempty"`
}

// StorageProvider defines the storage provider for the cluster. The data will be stored in the storage.
type StorageProvider struct {
	S3    *S3StorageProvider    `json:"s3,omitempty"`
	Local *LocalStorageProvider `json:"local,omitempty"`
}

type S3StorageProvider struct {
	// The data will be stored in the bucket.
	// +optional
	Bucket string `json:"bucket,omitempty"`

	// The region of the bucket.
	// +optional
	Region string `json:"region,omitempty"`

	// The endpoint of the bucket.
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// The secret of storing the credentials of access key id and secret access key.
	// The secret must be the same namespace with the GreptimeDBCluster resource.
	// +optional
	SecretName string `json:"secretName,omitempty"`

	// The prefix path of the data in the bucket.
	// +optional
	Prefix string `json:"prefix,omitempty"`
}

type LocalStorageProvider struct {
	// The local directory to store the data.
	Directory string `json:"directory,omitempty"`
}

// GreptimeDBClusterSpec defines the desired state of GreptimeDBCluster
type GreptimeDBClusterSpec struct {
	// Base is the base pod template for all components and can be overridden by template of individual component.
	// +optional
	Base *PodTemplateSpec `json:"base,omitempty"`

	// Frontend is the specification of frontend node.
	// +optional
	Frontend *FrontendSpec `json:"frontend"`

	// Meta is the specification of meta node.
	// +optional
	Meta *MetaSpec `json:"meta"`

	// Datanode is the specification of datanode node.
	// +optional
	Datanode *DatanodeSpec `json:"datanode"`

	// +optional
	HTTPServicePort int32 `json:"httpServicePort,omitempty"`

	// +optional
	GRPCServicePort int32 `json:"grpcServicePort,omitempty"`

	// +optional
	MySQLServicePort int32 `json:"mysqlServicePort,omitempty"`

	// +optional
	PostgresServicePort int32 `json:"postgresServicePort,omitempty"`

	// +optional
	OpenTSDBServicePort int32 `json:"openTSDBServicePort,omitempty"`

	// +optional
	EnableInfluxDBProtocol bool `json:"enableInfluxDBProtocol,omitempty"`

	// +optional
	PrometheusMonitor *PrometheusMonitor `json:"prometheusMonitor,omitempty"`

	// +optional
	// The version of greptimedb.
	Version string `json:"version,omitempty"`

	// +optional
	Initializer *InitializerSpec `json:"initializer,omitempty"`

	// +optional
	StorageProvider *StorageProvider `json:"storage,omitempty"`

	// More cluster settings can be added here.
}

// PrometheusMonitor defines the PodMonitor configuration.
type PrometheusMonitor struct {
	// Enable a prometheus PodMonitor
	// +optional
	Enabled bool `json:"enabled,omitempty"`

	// HTTP path to scrape for metrics.
	// +optional
	Path string `json:"path,omitempty"`

	// Name of the pod port this endpoint refers to. Mutually exclusive with targetPort.
	// +optional
	Port string `json:"port,omitempty"`

	// Interval at which metrics should be scraped
	// +optional
	Interval string `json:"interval,omitempty"`

	// HonorLabels chooses the metrics labels on collisions with target labels.
	// +optional
	HonorLabels bool `json:"honorLabels,omitempty"`

	// Prometheus PodMonitor selector.
	// +optional
	LabelsSelector map[string]string `json:"labelsSelector,omitempty"`
}

// GreptimeDBClusterStatus defines the observed state of GreptimeDBCluster
type GreptimeDBClusterStatus struct {
	Frontend FrontendStatus `json:"frontend,omitempty"`
	Meta     MetaStatus     `json:"meta,omitempty"`
	Datanode DatanodeStatus `json:"datanode,omitempty"`

	// +optional
	ClusterPhase ClusterPhase `json:"clusterPhase,omitempty"`

	// +optional
	Conditions []GreptimeDBClusterCondition `json:"conditions,omitempty"`
}

// GreptimeDBClusterCondition describes the state of a deployment at a certain point.
type GreptimeDBClusterCondition struct {
	// Type of deployment condition.
	Type GreptimeDBConditionType `json:"type"`

	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`

	// The last time this condition was updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`

	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`

	// A human-readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty"`
}

type FrontendStatus struct {
	Replicas      int32 `json:"replicas"`
	ReadyReplicas int32 `json:"readyReplicas"`
}

type MetaStatus struct {
	Replicas      int32 `json:"replicas"`
	ReadyReplicas int32 `json:"readyReplicas"`

	// +optional
	EtcdEndpoints []string `json:"etcdEndpoints,omitempty"`
}

type DatanodeStatus struct {
	Replicas      int32 `json:"replicas"`
	ReadyReplicas int32 `json:"readyReplicas"`
}

type GreptimeDBConditionType string

// These are valid conditions of a GreptimeDBCluster.
const (
	// GreptimeDBClusterReady indicates that the GreptimeDB cluster is ready to serve requests.
	// Every component in the cluster are all ready.
	GreptimeDBClusterReady GreptimeDBConditionType = "Ready"

	// GreptimeDBClusterProgressing indicates that the GreptimeDB cluster is progressing.
	GreptimeDBClusterProgressing GreptimeDBConditionType = "Progressing"
)

func NewCondition(conditionType GreptimeDBConditionType, conditionStatus corev1.ConditionStatus,
	reason, message string) *GreptimeDBClusterCondition {
	condition := GreptimeDBClusterCondition{
		Type:               conditionType,
		Status:             conditionStatus,
		LastTransitionTime: metav1.Now(),
		LastUpdateTime:     metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
	return &condition
}

func (in *GreptimeDBClusterStatus) GetCondition(conditionType GreptimeDBConditionType) *GreptimeDBClusterCondition {
	for i := range in.Conditions {
		c := in.Conditions[i]
		if c.Type == conditionType {
			return &c
		}
	}
	return nil
}

func (in *GreptimeDBClusterStatus) SetCondition(condition GreptimeDBClusterCondition) {
	currentCondition := in.GetCondition(condition.Type)
	if currentCondition != nil &&
		currentCondition.Status == condition.Status &&
		currentCondition.Reason == condition.Reason {
		currentCondition.LastUpdateTime = condition.LastUpdateTime
		return
	}

	if currentCondition != nil && currentCondition.Status == condition.Status {
		condition.LastTransitionTime = currentCondition.LastTransitionTime
	}

	newConditions := in.filterOutCondition(in.Conditions, condition.Type)
	in.Conditions = append(newConditions, condition)
}

func (in *GreptimeDBClusterStatus) filterOutCondition(conditions []GreptimeDBClusterCondition, conditionType GreptimeDBConditionType) []GreptimeDBClusterCondition {
	var newConditions []GreptimeDBClusterCondition
	for _, c := range conditions {
		if c.Type == conditionType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=gtc
// +kubebuilder:printcolumn:name="FRONTEND",type="integer",JSONPath=".status.frontend.readyReplicas"
// +kubebuilder:printcolumn:name="DATANODE",type="integer",JSONPath=".status.datanode.readyReplicas"
// +kubebuilder:printcolumn:name="META",type="integer",JSONPath=".status.meta.readyReplicas"
// +kubebuilder:printcolumn:name="PHASE",type=string,JSONPath=".status.clusterPhase"
// +kubebuilder:printcolumn:name="AGE",type=date,JSONPath=".metadata.creationTimestamp"

// GreptimeDBCluster is the Schema for the greptimedbclusters API
type GreptimeDBCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GreptimeDBClusterSpec   `json:"spec,omitempty"`
	Status GreptimeDBClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GreptimeDBClusterList contains a list of GreptimeDBCluster
type GreptimeDBClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GreptimeDBCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GreptimeDBCluster{}, &GreptimeDBClusterList{})
}
