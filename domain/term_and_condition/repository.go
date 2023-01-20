package term_and_condition

import (
	"go-crud/entity"
)

// create interface
type Repository interface {
	All() (out []entity.TermAndCondition, err error)
	SaveOrUpdateTermCondition(req entity.TermAndCondition) error
	FindTermConditionById(id string) (out entity.TermAndCondition, err error)
	FindTermConditionByTitle(title string) (out entity.TermAndCondition, err error)
	DeleteTermsCondition(id string) error
}
