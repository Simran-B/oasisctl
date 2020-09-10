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
// Author Gergely Brautigam
//
// +build e2e

package iam

import (
	"testing"

	"github.com/arangodb-managed/oasisctl/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	common "github.com/arangodb-managed/apis/common/v1"
	iam "github.com/arangodb-managed/apis/iam/v1"
	"github.com/arangodb-managed/oasisctl/cmd"
)

func TestCreateRole(t *testing.T) {
	cmd.RootCmd.PersistentPreRun(nil, nil)
	ctx := cmd.ContextWithToken()
	conn := cmd.MustDialAPI()
	iamc := iam.NewIAMServiceClient(conn)
	org, err := tests.GetDefaultOrganization()
	require.NoError(t, err)

	testRole := "testRole"
	defer func() {
		list, err := iamc.ListRoles(ctx, &common.ListOptions{ContextId: org})
		if err != nil {
			t.Log(err)
		}

		// We only delete the test role
		for _, role := range list.GetItems() {
			if role.GetName() == testRole {
				if _, err := iamc.DeleteRole(ctx, &common.IDOptions{Id: role.GetId()}); err != nil {
					t.Log(err)
				}
				break
			}
		}
	}()

	compare := `^Success!
Id          \d+
Name        ` + testRole + `
Description 
Predefined  false
Permissions 
Url         /Organization/\d+/Role/\d+
Created-At  .*
Deleted-At  -
$`

	args := []string{"create", "role", "--name=" + testRole}
	out, err := tests.RunCommand(args)
	require.NoError(t, err)
	assert.True(t, tests.CompareOutput(out, []byte(compare)))
}
