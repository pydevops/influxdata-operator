package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file

// RestoreSpec defines the desired config of Restore
type RestoreSpec struct {
	Database          string        `json:"database"`
	RestoreToDatabase string        `json:"restoreTo"`
	BackupId          string        `json:"backupId"`
	Storage           BackupStorage `json:"storage"`
	PodName           string        `json:"podname"`
	ContainerName     string        `json:"containername"`
	Rp                string        `json:"rp"`
	NewRp             string        `json:"newRp"`
	Shard             string        `json:"shard"`
}

// RestoreStatus defines the observed state of Restore
type RestoreStatus struct {
	RunResult string `json:"result"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Restore is the Schema for the restores API
// +k8s:openapi-gen=true
type Restore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestoreSpec   `json:"spec,omitempty"`
	Status RestoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RestoreList contains a list of Restore
type RestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Restore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Restore{}, &RestoreList{})
}
