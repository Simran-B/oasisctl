//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package data

import (
	"fmt"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	data "github.com/arangodb-managed/apis/data/v1"
	rm "github.com/arangodb-managed/apis/resourcemanager/v1"

	"github.com/arangodb-managed/oasis/cmd"
	"github.com/arangodb-managed/oasis/pkg/format"
	"github.com/arangodb-managed/oasis/pkg/selection"
)

var (
	// getDeploymentCmd fetches a deployment that the user has access to
	getDeploymentCmd = &cobra.Command{
		Use:   "deployment",
		Short: "Get a deployment the authenticated user has access to",
		Run:   getDeploymentCmdRun,
	}
	getDeploymentArgs struct {
		deploymentID   string
		organizationID string
		projectID      string
	}
)

func init() {
	cmd.InitCommand(
		cmd.GetCmd,
		getDeploymentCmd,
		func(c *cobra.Command, f *flag.FlagSet) {
			cargs := &getDeploymentArgs
			f.StringVarP(&cargs.deploymentID, "deployment-id", "d", cmd.DefaultDeployment(), "Identifier of the deployment")
			f.StringVarP(&cargs.organizationID, "organization-id", "o", cmd.DefaultOrganization(), "Identifier of the organization")
			f.StringVarP(&cargs.projectID, "project-id", "p", cmd.DefaultProject(), "Identifier of the project")
		},
	)
}

func getDeploymentCmdRun(c *cobra.Command, args []string) {
	// Validate arguments
	log := cmd.CLILog
	cargs := getDeploymentArgs
	deploymentID, argsUsed := cmd.OptOption("deployment-id", cargs.deploymentID, args, 0)
	cmd.MustCheckNumberOfArgs(args, argsUsed)

	// Connect
	conn := cmd.MustDialAPI()
	datac := data.NewDataServiceClient(conn)
	rmc := rm.NewResourceManagerServiceClient(conn)
	ctx := cmd.ContextWithToken()

	// Fetch deployment
	item := selection.MustSelectDeployment(ctx, log, deploymentID, cargs.projectID, cargs.organizationID, datac, rmc)

	// Show result
	fmt.Println(format.Deployment(item, cmd.RootArgs.Format))
}
