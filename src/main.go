package main

import (
	"fmt"

	"github.com/fatali-fataliyev/exman/pkg/group"
)

func main() {
	fmt.Println("a simple expenses manager between groups.")
	grp1, err := group.CreateGroup("Fatali")
	if err != nil {
		fmt.Println("failed to give name to group: ", err)
	}
	fmt.Println(grp1)
}
