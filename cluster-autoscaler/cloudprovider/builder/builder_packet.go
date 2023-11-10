//go:build packet
// +build packet

/*
Copyright 2019 The Kubernetes Authors.

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

package builder

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/equinixmetal"
	"k8s.io/autoscaler/cluster-autoscaler/config"
)

// AvailableCloudProviders supported by the cloud provider builder.
var AvailableCloudProviders = []string{
	equinixmetal.ProviderName,
	equinixmetal.EquinixMetalProviderName,
}

// DefaultCloudProvider for Packet-only build is Packet.
const DefaultCloudProvider = cloudprovider.EquinixMetalProviderName

func buildCloudProvider(opts config.AutoscalingOptions, do cloudprovider.NodeGroupDiscoveryOptions, rl *cloudprovider.ResourceLimiter) cloudprovider.CloudProvider {
	switch opts.CloudProviderName {
	case equinixmetal.ProviderName, cloudprovider.EquinixMetalProviderName:
		return equinixmetal.BuildCloudProvider(opts, do, rl)
	case cloudprovider.EquinixMetalProviderName:
		return equinixmetal.BuildCloudProvider(opts, do, rl)
	}

	return nil
}
