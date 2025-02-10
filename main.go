package main

import (
	"fmt"

	"github.com/fatali-fataliyev/exman/pkg/group"
)

func main() {

	fmt.Println("Welcome to ExMan")
	fmt.Println("====")
	var groupName string
	fmt.Print("Enter group name: ")
	fmt.Scanln(&groupName)

	createdGroup, err := group.CreateGroup(groupName)
	if err != nil {
		fmt.Println("failed to give name to group: ", err)
		return
	}

	var membersCount int

	fmt.Printf("Approximately how many people will be in the group?: ")
	fmt.Scan(&membersCount)

	for i := 0; i < membersCount; i++ {
		member, err := group.AddMember()
		if err != nil {
			fmt.Println("failed to add member: ", err)
		}
		createdGroup.Members = append(createdGroup.Members, member)
	}

	fmt.Println(createdGroup)
}
