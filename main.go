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

	grp1, err := group.CreateGroup(groupName)
	if err != nil {
		fmt.Println("failed to give name to group: ", err)
		return
	}
	fmt.Println(grp1)
}
