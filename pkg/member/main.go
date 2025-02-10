package member

import (
	"fmt"

	Category "github.com/fatali-fataliyev/exman/pkg/category"
)

type Member struct {
	ID           int
	name         string
	surname      string
	paidCategory Category.Category
	amount       float64
	currency     string
}

func CreateMember(ID int, name string, surname string, piadCategory Category.Category, amount float64, currency string) (Member, error) {
	if ID < 0 {
		return Member{}, fmt.Errorf("ID numbers should be positive")
	}
	if len(name) > 20 {
		return Member{}, fmt.Errorf("maximum allowed name length is 20 characters")
	}
	if len(surname) > 30 {
		return Member{}, fmt.Errorf("maximum allowed name length is 30 characters")
	}

	return Member{
		ID,
		name,
		surname,
		piadCategory,
		amount,
		currency,
	}, nil

}
