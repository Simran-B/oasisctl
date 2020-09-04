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

package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/arangodb-managed/oasisctl/cmd/crypto"
	"github.com/arangodb-managed/oasisctl/tests"
)

func TestDeleteCryptoInvalidFlag(t *testing.T) {
	args := []string{"delete", "cacertificate", "--invalid"}
	compare := `^Error: unknown flag: --invalid
Usage:
  oasisctl delete cacertificate [flags]

Flags:
  -c, --cacertificate-id string   Identifier of the CA certificate
  -h, --help                      help for cacertificate
  -o, --organization-id string    Identifier of the organization.*
  -p, --project-id string         Identifier of the project.*

Global Flags:
      --endpoint string   API endpoint of the ArangoDB Oasis (default ".*")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis

.*.unknown.flag:.--invalid
$`
	out, err := tests.RunCommand(args)
	require.Error(t, err)
	assert.True(t, tests.CompareOutput(out, []byte(compare)))
}
