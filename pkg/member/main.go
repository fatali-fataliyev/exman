package member

import "fmt"

type Member struct {
	ID           int
	name         string
	surname      string
	paidCategory string
	amount       float64
}

func createMember(ID int, name string, surname string, piadCategory string, amount float64) (Member, error) {

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
	}, nil

}
