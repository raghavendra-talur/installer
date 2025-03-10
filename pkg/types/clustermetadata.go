package types

import (
	"github.com/openshift/installer/pkg/types/aws"
	"github.com/openshift/installer/pkg/types/azure"
	"github.com/openshift/installer/pkg/types/baremetal"
	"github.com/openshift/installer/pkg/types/gcp"
	"github.com/openshift/installer/pkg/types/ibmcloud"
	"github.com/openshift/installer/pkg/types/kubevirt"
	"github.com/openshift/installer/pkg/types/libvirt"
	"github.com/openshift/installer/pkg/types/openstack"
	"github.com/openshift/installer/pkg/types/ovirt"
	"github.com/openshift/installer/pkg/types/vsphere"
)

// ClusterMetadata contains information
// regarding the cluster that was created by installer.
type ClusterMetadata struct {
	// ClusterName is the name for the cluster.
	ClusterName string `json:"clusterName"`
	// ClusterID is a globally unique ID that is used to identify an Openshift cluster.
	ClusterID string `json:"clusterID"`
	// InfraID is an ID that is used to identify cloud resources created by the installer.
	InfraID                 string `json:"infraID"`
	ClusterPlatformMetadata `json:",inline"`
}

// ClusterPlatformMetadata contains metadata for platfrom.
type ClusterPlatformMetadata struct {
	AWS       *aws.Metadata       `json:"aws,omitempty"`
	OpenStack *openstack.Metadata `json:"openstack,omitempty"`
	Libvirt   *libvirt.Metadata   `json:"libvirt,omitempty"`
	Azure     *azure.Metadata     `json:"azure,omitempty"`
	GCP       *gcp.Metadata       `json:"gcp,omitempty"`
	IBMCloud  *ibmcloud.Metadata  `json:"ibmcloud,omitempty"`
	BareMetal *baremetal.Metadata `json:"baremetal,omitempty"`
	Ovirt     *ovirt.Metadata     `json:"ovirt,omitempty"`
	VSphere   *vsphere.Metadata   `json:"vsphere,omitempty"`
	Kubevirt  *kubevirt.Metadata  `json:"kubevirt,omitempty"`
}

// Platform returns a string representation of the platform
// (e.g. "aws" if AWS is non-nil).  It returns an empty string if no
// platform is configured.
func (cpm *ClusterPlatformMetadata) Platform() string {
	if cpm == nil {
		return ""
	}
	if cpm.AWS != nil {
		return aws.Name
	}
	if cpm.Libvirt != nil {
		return libvirt.Name
	}
	if cpm.OpenStack != nil {
		return openstack.Name
	}
	if cpm.Azure != nil {
		return azure.Name
	}
	if cpm.GCP != nil {
		return gcp.Name
	}
	if cpm.IBMCloud != nil {
		return ibmcloud.Name
	}
	if cpm.BareMetal != nil {
		return "baremetal"
	}
	if cpm.Ovirt != nil {
		return ovirt.Name
	}
	if cpm.VSphere != nil {
		return vsphere.Name
	}
	if cpm.Kubevirt != nil {
		return kubevirt.Name
	}
	return ""
}
