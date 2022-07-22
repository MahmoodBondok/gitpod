// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package main

import (
	"fmt"

	"github.com/gitpod-io/gitpod/authorizer/pkg/dbgen"
)

var user = &dbgen.TypeSpec{
	Name:     "user",
	Table:    "d_b_user",
	IDColumn: "id",
	Relations: []dbgen.Relation{
		{
			Name: "owner",
			Targets: []dbgen.RelationTarget{
				dbgen.RelationSelf{},
			},
		},
		{
			Name: "writer",
			Targets: []dbgen.RelationTarget{
				dbgen.RelationSelf{},
				dbgen.RelationRef("owner"),
			},
		},
		{
			Name: "reader",
			Targets: []dbgen.RelationTarget{
				dbgen.RelationSelf{},
				dbgen.RelationRef("writer"),
			},
		},
	},
}

var workspace = &dbgen.TypeSpec{
	Name:     "workspace",
	Table:    "d_b_workspace",
	IDColumn: "id",
	Relations: []dbgen.Relation{
		{
			Name: "owner",
			Targets: []dbgen.RelationTarget{
				dbgen.RelationTable{
					Target: user,
					Column: "ownerId",
				},
			},
		},
		{
			Name: "access",
			Targets: []dbgen.RelationTarget{
				dbgen.RelationSelf{},
				dbgen.RelationRef("owner"),
			},
		},
		{
			Name: "writer",
			Targets: []dbgen.RelationTarget{
				dbgen.RelationSelf{},
				dbgen.RelationRef("owner"),
			},
		},
		{
			Name: "reader",
			Targets: []dbgen.RelationTarget{
				dbgen.RelationSelf{},
				dbgen.RelationRef("access"),
				dbgen.RelationRef("writer"),
			},
		},
	},
}

var workspaceInstance = &dbgen.TypeSpec{
	Name:     "workspace_instance",
	Table:    "d_b_workspace_instance",
	IDColumn: "id",
	Relations: []dbgen.Relation{
		{
			Name: "owner",
			Targets: []dbgen.RelationTarget{
				dbgen.RelationRemoteRef{
					Target: workspace,
					Name:   "owner",
				},
			},
		},
		{
			Name: "access",
			Targets: []dbgen.RelationTarget{
				dbgen.RelationRef("owner"),
				dbgen.RelationRemoteRef{
					Target: workspace,
					Name:   "access",
				},
			},
		},
	},
}

func main() {
	sess := dbgen.NewSession("main_test")
	sess.Generate(user)
	sess.Generate(workspace)
	sess.Generate(workspaceInstance)
	sess.Commit()
	fmt.Println(sess)
}
