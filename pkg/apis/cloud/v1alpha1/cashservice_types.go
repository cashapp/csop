package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	autoscaling "k8s.io/api/autoscaling/v2beta2"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CashServiceSpec defines the desired state of CashService
// +k8s:openapi-gen=true
type CashServiceSpec struct {
	// +kubebuilder:validation:Enum=Standard;Jobs
	// +kubebuilder:Default=Standard
	Kind string `json:"kind"`

	// labels to be applied to all resources
	Labels map[string]string `json:"labels,omitempty"`
	// container contains the configuration required to setup the container
	Container CashServiceContainer `json:"container"`
	// autoscaling contains the configuration used to setup the horizontal pod
	// autoscaler
	Autoscaling CashServiceAutoscaling `json:"autoscaling"`
	// rollingUpdate specifies the behaviour of a rolling update of the service
	RollingUpdate appsv1.RollingUpdateDeployment `json:"rollingUpdate,omitempty"`
	// environment contains the environment variables that will be accessible
	// to the service containers
	Environment corev1.ConfigMap `json:"environment,omitempty"`
	// networkPolicy configures the acceptable sources and destinations of
	// network traffic
	NetworkPolicy netv1.NetworkPolicy `json:"networkPolicy,omitempty"`
}

type CashServiceAutoscaling struct {
	// minReplicas is the lower limit for the number of replicas to which the
	// autoscaler can scale down.  It defaults to 1 pod.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:Default=1
	// +optional
	MinReplicas int32 `json:"minReplicas,omitempty"`
	// maxReplicas is the upper limit for the number of replicas to which the
	// autoscaler can scale up.
	// It cannot be less that minReplicas.
	MaxReplicas int32 `json:"maxReplicas"`
	// metrics contains the specifications for which to use to calculate the
	// desired replica count (the maximum replica count across all metrics will
	// be used).  The desired replica count is calculated multiplying the
	// ratio between the target value and the current value by the current
	// number of pods.  Ergo, metrics used must decrease as the pod count is
	// increased, and vice-versa.  See the individual metric source types for
	// more information about how each type of metric must respond.
	// If not set, the default metric will be set to 80% average CPU utilization.
	// +optional
	Metrics []autoscaling.MetricSpec `json:"metrics,omitempty"`
}

type CashServiceContainer struct {
	// deployment contains the configuration required to deploy an image to the
	// service container.
	Deployment CashServiceDeployment `json:"deployment"`
	// resouces contains the resource requests and limits to be applied to the
	// service container.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	// debug contains the configuration related to the deployment of a debug
	// container in the pod.
	Debug CashServiceDebug `json:"debug,omitempty"`
}

type CashServiceDeployment struct {
	// +kubebuilder:validation:Enum=Manual;Automatic
	// +kubebuilder:Default=Manual
	Mode string `json:"mode"`
	// manual contains the configuration required to manually deploy a service
	// using an image URL.
	// +optional
	Manual CashServiceManualDeployment `json:"manual,omitempty"`
	// automatic contains the configuration required to automatically deploy
	// the most recent image from a docker repository.
	// +optional
	Automatic CashServiceAutomaticDeployment `json:"automatic,omitempty"`
}

type CashServiceManualDeployment struct {
	//  imageURL specifies the image to deploy in the service container.
	ImageURL string `json:"imageURL"`
}

type CashServiceAutomaticDeployment struct {
	// respository specifies a url to the respository from which the most
	// recent image should be pulled.
	Repository string `json:"repository"`
}

type CashServiceDebug struct {
	// enabled indicates whether a debug container should be included within
	// the pod.
	// +kubebuilder:Default=false
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// imageURL specifes the image to deploy in the debug container.
	// +optional
	ImageURL string `json:"imageURL,omitempty"`
}

// CashServiceStatus defines the observed state of CashService
// +k8s:openapi-gen=true
type CashServiceStatus struct {
	// deployments contains the status of each of the deployments managed by
	// the CashService.
	Deployments map[string]appsv1.DeploymentStatus `json:"deployments,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CashService is the Schema for the cashservices API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=cashservices,scope=Namespaced
type CashService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CashServiceSpec   `json:"spec,omitempty"`
	Status CashServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CashServiceList contains a list of CashService
type CashServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CashService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CashService{}, &CashServiceList{})
}
