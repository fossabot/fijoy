//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var FijoyWorkspaceUser = newFijoyWorkspaceUserTable("public", "fijoy_workspace_user", "")

type fijoyWorkspaceUserTable struct {
	postgres.Table

	// Columns
	WorkspaceID postgres.ColumnString
	UserID      postgres.ColumnString
	Role        postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FijoyWorkspaceUserTable struct {
	fijoyWorkspaceUserTable

	EXCLUDED fijoyWorkspaceUserTable
}

// AS creates new FijoyWorkspaceUserTable with assigned alias
func (a FijoyWorkspaceUserTable) AS(alias string) *FijoyWorkspaceUserTable {
	return newFijoyWorkspaceUserTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FijoyWorkspaceUserTable with assigned schema name
func (a FijoyWorkspaceUserTable) FromSchema(schemaName string) *FijoyWorkspaceUserTable {
	return newFijoyWorkspaceUserTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FijoyWorkspaceUserTable with assigned table prefix
func (a FijoyWorkspaceUserTable) WithPrefix(prefix string) *FijoyWorkspaceUserTable {
	return newFijoyWorkspaceUserTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FijoyWorkspaceUserTable with assigned table suffix
func (a FijoyWorkspaceUserTable) WithSuffix(suffix string) *FijoyWorkspaceUserTable {
	return newFijoyWorkspaceUserTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFijoyWorkspaceUserTable(schemaName, tableName, alias string) *FijoyWorkspaceUserTable {
	return &FijoyWorkspaceUserTable{
		fijoyWorkspaceUserTable: newFijoyWorkspaceUserTableImpl(schemaName, tableName, alias),
		EXCLUDED:                newFijoyWorkspaceUserTableImpl("", "excluded", ""),
	}
}

func newFijoyWorkspaceUserTableImpl(schemaName, tableName, alias string) fijoyWorkspaceUserTable {
	var (
		WorkspaceIDColumn = postgres.StringColumn("workspace_id")
		UserIDColumn      = postgres.StringColumn("user_id")
		RoleColumn        = postgres.StringColumn("role")
		allColumns        = postgres.ColumnList{WorkspaceIDColumn, UserIDColumn, RoleColumn}
		mutableColumns    = postgres.ColumnList{RoleColumn}
	)

	return fijoyWorkspaceUserTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		WorkspaceID: WorkspaceIDColumn,
		UserID:      UserIDColumn,
		Role:        RoleColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
