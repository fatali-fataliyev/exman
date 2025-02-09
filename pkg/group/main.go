package group

import (
	"fmt"
	"time"

	"github.com/fatali-fataliyev/exman/pkg/member"
)

type Group struct {
	name         string
	members      []member.Member //after retype to User type.
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

	groupMembers := []member.Member{}

	currentTime := time.Now()
	createdDate := currentTime.Format("2006/01/02")

	return Group{
		name:         groupName,
		members:      groupMembers,
		creationDate: createdDate,
	}, nil

}

func AddMember() (member.Member, error) {
	var name, surname, paidCategory string
	var ID = 1
	var amount float64

	fmt.Print("Enter member name: " + "\n")
	fmt.Scan(&name)
	fmt.Print("Enter member surname: " + "\n")
	fmt.Scan(&surname)
	fmt.Print("Enter member paid category: " + "\n")
	fmt.Scan(&paidCategory)
	fmt.Print("Enter member paid amount: ")
	fmt.Scan(&amount)

	newMember, err := member.CreateMember(ID, name, surname, paidCategory, amount)

	if err != nil {
		return member.Member{}, fmt.Errorf("failed to create member: %w", err)
	}

	return newMember, nil
}
