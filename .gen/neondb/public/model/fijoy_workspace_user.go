//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type FijoyWorkspaceUser struct {
	WorkspaceID string `sql:"primary_key"`
	UserID      string `sql:"primary_key"`
	Role        string
}
