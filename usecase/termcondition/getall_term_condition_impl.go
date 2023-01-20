package termcondition

import "go-crud/entity"

func (s *service) GetAllTermCondition() ([]entity.TermAndCondition, error) {
	var out []entity.TermAndCondition
	result, err := s.termAndConditionRepo.All()
	if err != nil {
		return out, err
	}
	out = result
	return out, err
}
