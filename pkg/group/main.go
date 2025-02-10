package group

import (
	"fmt"
	"time"

	Category "github.com/fatali-fataliyev/exman/pkg/category"
	"github.com/fatali-fataliyev/exman/pkg/member"
	"golang.org/x/exp/rand"
)

type Group struct {
	Name         string
	Members      []member.Member
	MembersCount int
	CreationDate string
}

func CreateGroup(groupName string) (Group, error) {
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
		Name:         groupName,
		Members:      groupMembers,
		MembersCount: len(groupMembers) + 1,
		CreationDate: createdDate,
	}, nil

}

func AddMember() (member.Member, error) {
	var name, surname, currency string
	var paidCategory Category.Category
	var ID = rand.Intn(1000) + 1
	var amount float64

	fmt.Print("Enter member name: " + "\n")
	fmt.Scan(&name)
	fmt.Print("Enter member surname: " + "\n")
	fmt.Scan(&surname)
	fmt.Print("Enter member paid category: " + "\n")
	fmt.Scan(&paidCategory)
	fmt.Print("Enter member paid amount: ")
	fmt.Scan(&amount)
	fmt.Print("Enter member paid amount currency(usd, eur, azn, rub...): ")
	fmt.Scan(&currency)

	newMember, err := member.CreateMember(ID, name, surname, paidCategory, amount, currency)

	if err != nil {
		return member.Member{}, fmt.Errorf("failed to create member: %w", err)
	}

	return newMember, nil
}

func getLendPersons() {

}
