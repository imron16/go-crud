package termcondition

import (
	"go-crud/domain/term_and_condition"
)

type service struct {
	termAndConditionRepo term_and_condition.Repository
}

func NewTermsService(
	termAndConditionRepo term_and_condition.Repository,
) Service {
	return &service{termAndConditionRepo: termAndConditionRepo}
}
