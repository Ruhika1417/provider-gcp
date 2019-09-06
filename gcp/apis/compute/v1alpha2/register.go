/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// +kubebuilder:object:generate=true
// +groupName=compute.gcp.crossplane.io
// +versionName=v1alpha2

package v1alpha2

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
)

// Packaage type metadata.
const (
	Group   = "compute.gcp.crossplane.io"
	Version = "v1alpha2"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// GKECluster type metadata.
var (
	GKEClusterKind             = reflect.TypeOf(GKECluster{}).Name()
	GKEClusterKindAPIVersion   = GKEClusterKind + "." + SchemeGroupVersion.String()
	GKEClusterGroupVersionKind = SchemeGroupVersion.WithKind(GKEClusterKind)
)

// GKEClusterClass type metadata.
var (
	GKEClusterClassKind             = reflect.TypeOf(GKEClusterClass{}).Name()
	GKEClusterClassKindAPIVersion   = GKEClusterClassKind + "." + SchemeGroupVersion.String()
	GKEClusterClassGroupVersionKind = SchemeGroupVersion.WithKind(GKEClusterClassKind)
)

// Network type metadata.
var (
	NetworkKind             = reflect.TypeOf(Network{}).Name()
	NetworkKindAPIVersion   = NetworkKind + "." + SchemeGroupVersion.String()
	NetworkGroupVersionKind = SchemeGroupVersion.WithKind(NetworkKind)
)

// Subnetwork type metadata.
var (
	SubnetworkKind             = reflect.TypeOf(Subnetwork{}).Name()
	SubnetworkKindAPIVersion   = SubnetworkKind + "." + SchemeGroupVersion.String()
	SubnetworkGroupVersionKind = SchemeGroupVersion.WithKind(SubnetworkKind)
)

func init() {
	SchemeBuilder.Register(&GKECluster{}, &GKEClusterList{})
	SchemeBuilder.Register(&GKEClusterClass{}, &GKEClusterClassList{})
	SchemeBuilder.Register(&Network{}, &NetworkList{})
	SchemeBuilder.Register(&Subnetwork{}, &SubnetworkList{})
}