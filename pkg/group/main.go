package group

import (
	"fmt"
)

type Group struct {
	name         string
	members      int //after retype to User type.
	creationDate string
}

func CreateGroup(groupName string) (Group, error) {
	//creates group name
	//add user functionality.
	//add current time to group.

	if groupName == "" {
		return Group{}, fmt.Errorf("group name is empty, write a normal group name please.")
	}

	return Group{}, nil

}
