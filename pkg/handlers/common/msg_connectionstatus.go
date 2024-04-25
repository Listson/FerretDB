// Copyright 2021 FerretDB Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"context"

	"github.com/FerretDB/FerretDB/pkg/clientconn/conninfo"
	"github.com/FerretDB/FerretDB/pkg/types"
	"github.com/FerretDB/FerretDB/pkg/util/must"
	"github.com/FerretDB/FerretDB/pkg/wire"
)

// MsgConnectionStatus is a common implementation of the connectionStatus command.
func MsgConnectionStatus(ctx context.Context, _ *wire.OpMsg) (*wire.OpMsg, error) {
	users := types.MakeArray(1)

	if username, _ := conninfo.Get(ctx).Auth(); username != "" {
		users.Append(must.NotFail(types.NewDocument(
			"user", username,
		)))
	}

	var reply wire.OpMsg
	must.NoError(reply.SetSections(wire.OpMsgSection{
		Documents: []*types.Document{must.NotFail(types.NewDocument(
			"authInfo", must.NotFail(types.NewDocument(
				"authenticatedUsers", users,
				"authenticatedUserRoles", must.NotFail(types.NewArray()),
				"authenticatedUserPrivileges", must.NotFail(types.NewArray()),
			)),
			"ok", float64(1),
		))},
	}))

	return &reply, nil
}