package main

import (
	"fmt"

	"github.com/fatali-fataliyev/exman/pkg/group"
)

func main() {
	grp1, err := group.CreateGroup("tallinn 2024 holiday")
	if err != nil {
		fmt.Println("failed to give name to group: ", err)
		return
	}
	fmt.Println(grp1)
}
