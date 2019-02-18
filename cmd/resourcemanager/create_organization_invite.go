//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package rm

import (
	"fmt"
	"github.com/arangodb-managed/oasis/cmd"

	"github.com/spf13/cobra"

	iam "github.com/arangodb-managed/apis/iam/v1"
	rm "github.com/arangodb-managed/apis/resourcemanager/v1"
	"github.com/arangodb-managed/oasis/pkg/format"
	"github.com/arangodb-managed/oasis/pkg/selection"
)

var (
	// createOrganizationInvite creates a new organization invite
	createOrganizationInviteCmd = &cobra.Command{
		Use:   "invite",
		Short: "Create a new invite to an organization",
		Run:   createOrganizationInviteCmdRun,
	}
	createOrganizationInviteArgs struct {
		email          string
		organizationID string
	}
)

func init() {
	createOrganizationCmd.AddCommand(createOrganizationInviteCmd)

	f := createOrganizationInviteCmd.Flags()
	f.StringVar(&createOrganizationInviteArgs.email, "email", "", "Email address of the person to invite")
	f.StringVarP(&createOrganizationInviteArgs.organizationID, "organization-id", "o", cmd.DefaultOrganization(), "Identifier of the organization to create the invite in")
}

func createOrganizationInviteCmdRun(c *cobra.Command, args []string) {
	// Validate arguments
	log := cmd.CLILog
	cargs := createOrganizationInviteArgs
	email, argsUsed := cmd.ReqOption("email", cargs.email, args, 0)
	cmd.MustCheckNumberOfArgs(args, argsUsed)

	// Connect
	conn := cmd.MustDialAPI()
	iamc := iam.NewIAMServiceClient(conn)
	rmc := rm.NewResourceManagerServiceClient(conn)
	ctx := cmd.ContextWithToken()

	// Fetch organization
	org := selection.MustSelectOrganization(ctx, log, cargs.organizationID, rmc)

	// Create invite
	result, err := rmc.CreateOrganizationInvite(ctx, &rm.OrganizationInvite{
		OrganizationId: org.GetId(),
		Email:          email,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create organization invite")
	}

	// Show result
	fmt.Println("Success!")
	fmt.Println(format.OrganizationInvite(ctx, result, iamc, cmd.RootArgs.Format))
}