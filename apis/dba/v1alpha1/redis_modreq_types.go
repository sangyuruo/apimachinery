/*
Copyright The KubeDB Authors.

Licensed under the Apache License, ModificationRequest 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceCodeRedisModificationRequest     = "rdmodreq"
	ResourceKindRedisModificationRequest     = "RedisModificationRequest"
	ResourceSingularRedisModificationRequest = "redismodificationrequest"
	ResourcePluralRedisModificationRequest   = "redismodificationrequests"
)

// RedisModificationRequest defines a Redis database version.

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=redismodificationrequests,singular=redismodificationrequest,shortName=rdmodreq,scope=Cluster,categories={datastore,kubedb,appscode}
type RedisModificationRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              RedisModificationRequestSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            RedisModificationRequestStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// RedisModificationRequestSpec is the spec for redis version
type RedisModificationRequestSpec struct {
	Version string              `json:"version,omitempty" protobuf:"bytes,1,opt,name=version"`
	Redis   *v1.ObjectReference `json:"mongodb,omitempty" protobuf:"bytes,2,opt,name=redis"`
}

// RedisModificationRequestStatus is the status for redis version
type RedisModificationRequestStatus struct {
	// Conditions applied to the request, such as approval or denial.
	// +optional
	Conditions []RedisModificationRequestCondition `json:"conditions,omitempty" protobuf:"bytes,1,rep,name=conditions"`
	// observedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,2,opt,name=observedGeneration"`
}

type RedisModificationRequestCondition struct {
	// request approval state, currently Approved or Denied.
	Type RequestConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=RequestConditionType"`

	// brief reason for the request state
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,2,opt,name=reason"`

	// human readable message with details about the request state
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`

	// timestamp for the last update to this condition
	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty" protobuf:"bytes,4,opt,name=lastUpdateTime"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisModificationRequestList is a list of RedisModificationRequests
type RedisModificationRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of RedisModificationRequest CRD objects
	Items []RedisModificationRequest `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
