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

package iam

import (
	"fmt"

	"github.com/spf13/cobra"

	common "github.com/arangodb-managed/apis/common/v1"
	iam "github.com/arangodb-managed/apis/iam/v1"
	rm "github.com/arangodb-managed/apis/resourcemanager/v1"

	"github.com/arangodb-managed/oasisctl/cmd"
	"github.com/arangodb-managed/oasisctl/pkg/format"
	"github.com/arangodb-managed/oasisctl/pkg/selection"
)

var (
	// listRolesCmd fetches roles of the given organization
	listRolesCmd = &cobra.Command{
		Use:   "roles",
		Short: "List all roles of the given organization",
		Run:   listRolesCmdRun,
	}
	listRolesArgs struct {
		organizationID string
	}
)

func init() {
	cmd.ListCmd.AddCommand(listRolesCmd)
	f := listRolesCmd.Flags()
	f.StringVarP(&listRolesArgs.organizationID, "organization-id", "o", cmd.DefaultOrganization(), "Identifier of the organization")
}

func listRolesCmdRun(c *cobra.Command, args []string) {
	// Validate arguments
	log := cmd.CLILog
	cargs := listRolesArgs
	organizationID, argsUsed := cmd.OptOption("organization-id", cargs.organizationID, args, 0)
	cmd.MustCheckNumberOfArgs(args, argsUsed)

	// Connect
	conn := cmd.MustDialAPI()
	iamc := iam.NewIAMServiceClient(conn)
	rmc := rm.NewResourceManagerServiceClient(conn)
	ctx := cmd.ContextWithToken()

	// Fetch organization
	org := selection.MustSelectOrganization(ctx, log, organizationID, rmc)

	// Fetch roles in organization
	list, err := iamc.ListRoles(ctx, &common.ListOptions{ContextId: org.GetId()})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to list roles")
	}

	// Show result
	fmt.Println(format.RoleList(list.Items, cmd.RootArgs.Format))
}
