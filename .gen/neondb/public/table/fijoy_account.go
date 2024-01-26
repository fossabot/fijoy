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

var FijoyAccount = newFijoyAccountTable("public", "fijoy_account", "")

type fijoyAccountTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnString
	Name        postgres.ColumnString
	AccountType postgres.ColumnString
	Institution postgres.ColumnString
	WorkspaceID postgres.ColumnString
	Balance     postgres.ColumnFloat
	CreatedAt   postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FijoyAccountTable struct {
	fijoyAccountTable

	EXCLUDED fijoyAccountTable
}

// AS creates new FijoyAccountTable with assigned alias
func (a FijoyAccountTable) AS(alias string) *FijoyAccountTable {
	return newFijoyAccountTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FijoyAccountTable with assigned schema name
func (a FijoyAccountTable) FromSchema(schemaName string) *FijoyAccountTable {
	return newFijoyAccountTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FijoyAccountTable with assigned table prefix
func (a FijoyAccountTable) WithPrefix(prefix string) *FijoyAccountTable {
	return newFijoyAccountTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FijoyAccountTable with assigned table suffix
func (a FijoyAccountTable) WithSuffix(suffix string) *FijoyAccountTable {
	return newFijoyAccountTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFijoyAccountTable(schemaName, tableName, alias string) *FijoyAccountTable {
	return &FijoyAccountTable{
		fijoyAccountTable: newFijoyAccountTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newFijoyAccountTableImpl("", "excluded", ""),
	}
}

func newFijoyAccountTableImpl(schemaName, tableName, alias string) fijoyAccountTable {
	var (
		IDColumn          = postgres.StringColumn("id")
		NameColumn        = postgres.StringColumn("name")
		AccountTypeColumn = postgres.StringColumn("account_type")
		InstitutionColumn = postgres.StringColumn("institution")
		WorkspaceIDColumn = postgres.StringColumn("workspace_id")
		BalanceColumn     = postgres.FloatColumn("balance")
		CreatedAtColumn   = postgres.TimestampColumn("created_at")
		allColumns        = postgres.ColumnList{IDColumn, NameColumn, AccountTypeColumn, InstitutionColumn, WorkspaceIDColumn, BalanceColumn, CreatedAtColumn}
		mutableColumns    = postgres.ColumnList{NameColumn, AccountTypeColumn, InstitutionColumn, WorkspaceIDColumn, BalanceColumn, CreatedAtColumn}
	)

	return fijoyAccountTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		AccountType: AccountTypeColumn,
		Institution: InstitutionColumn,
		WorkspaceID: WorkspaceIDColumn,
		Balance:     BalanceColumn,
		CreatedAt:   CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
