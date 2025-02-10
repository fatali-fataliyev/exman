package Category

import "fmt"

type Category struct {
	name         string
	amount       float64
	paidMemberID int
}

func CreateCategory(name string, amount float64, paidMemberID int) (Category, error) {
	if len(name) > 30 {
		return Category{}, fmt.Errorf("Max allowed category name length is 30 characters.")
	}

	return Category{
		name,
		amount,
		paidMemberID,
	}, nil

}
