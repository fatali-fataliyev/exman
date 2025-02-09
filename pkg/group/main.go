package group

import (
	"fmt"
	"time"
)

type Group struct {
	name         string
	members      int //after retype to User type.
	creationDate string
}

func CreateGroup(groupName string) (Group, error) {
	//add user functionality.
	//add current time to group.

	if groupName == "" {
		return Group{}, fmt.Errorf("group name is empty")
	}
	if len(groupName) > 30 {
		return Group{}, fmt.Errorf("group name is long, maximum allowed group name length is 30 characters")
	}

	// members := make(map[string]any)

	currentTime := time.Now()
	createdDate := currentTime.Format("2006/01/02")

	return Group{
		name:         groupName,
		members:      0,
		creationDate: createdDate,
	}, nil

}

func addMember() {

}
