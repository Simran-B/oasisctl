//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
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
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package format

import (
	"fmt"

	data "github.com/arangodb-managed/apis/data/v1"
)

// Deployment returns a single deployment formatted for humans.
func Deployment(x *data.Deployment, creds *data.DeploymentCredentials, opts Options, showRootpassword bool) string {
	pwd := func(creds *data.DeploymentCredentials) string {
		if showRootpassword {
			return creds.GetPassword()
		}
		return "*** use '--show-root-password' to expose ***"
	}
	d := []kv{
		{"id", x.GetId()},
		{"name", x.GetName()},
		{"description", x.GetDescription()},
		{"region", x.GetRegionId()},
		{"version", x.GetVersion()},
		{"ipwhitelist", x.GetIpwhitelistId()},
		{"url", x.GetUrl()},
		{"paused", formatBool(opts, x.GetIsPaused())},
		{"created-at", formatTime(opts, x.GetCreatedAt())},
		{"deleted-at", formatTime(opts, x.GetDeletedAt(), "-")},
		{"expires-at", formatTime(opts, x.GetExpiration().GetExpiresAt(), "-")},
		{"ready", x.GetStatus().GetReady()},
		{"bootstrapped", x.GetStatus().GetBootstrapped()},
		{"created", x.GetStatus().GetCreated()},
		{"upgrading", x.GetStatus().GetUpgrading()},

		{"coordinators", x.GetServers().GetCoordinators()},
		{"coordinator-memory-size", fmt.Sprintf("%d%s", x.GetServers().GetCoordinatorMemorySize(), "GB")},
		{"dbservers", x.GetServers().GetDbservers()},
		{"dbserver-memory-size", fmt.Sprintf("%d%s", x.GetServers().GetDbserverMemorySize(), "GB")},
		{"dbserver-disk-size", fmt.Sprintf("%d%s", x.GetServers().GetDbserverDiskSize(), "GB")},

		{"bootstrapped-at", formatTime(opts, x.GetStatus().GetBootstrappedAt(), "-")},
		{"endpoint-url", x.GetStatus().GetEndpoint()},
		{"root-password", pwd(creds)},

		{"model", x.Model.Model},
		{"is-clone", x.GetIsClone()},
		{"clone-backup-id", x.GetCloneBackupId()},
	}
	if x.Model.Model != data.ModelFlexible {
		d = append(d,
			kv{"node-count", fmt.Sprintf("%d", x.Model.NodeCount)},
			kv{"node-disk-size", fmt.Sprintf("%d%s", x.Model.NodeDiskSize, "GB")},
			kv{"node-size-id", x.Model.NodeSizeId})
	}
	return formatObject(opts, d...)
}

// DeploymentList returns a list of deployments formatted for humans.
func DeploymentList(list []*data.Deployment, opts Options) string {
	return formatList(opts, list, func(i int) []kv {
		x := list[i]
		d := []kv{
			{"id", x.GetId()},
			{"name", x.GetName()},
			{"description", x.GetDescription()},
			{"region", x.GetRegionId()},
			{"version", x.GetVersion()},
			{"ipwhitelist", x.GetIpwhitelistId()},
			{"url", x.GetUrl()},
			{"paused", formatBool(opts, x.GetIsPaused())},
			{"created-at", formatTime(opts, x.GetCreatedAt())},
			{"model", x.Model.Model},
		}
		if x.Model.Model != data.ModelFlexible {
			d = append(d,
				kv{"node-count", fmt.Sprintf("%d", x.Model.NodeCount)},
				kv{"node-disk-size", fmt.Sprintf("%d%s", x.Model.NodeDiskSize, "GB")},
				kv{"node-size-id", x.Model.NodeSizeId})
		}
		return d
	}, false)
}
