package termcondition

import "go-crud/entity"

func (s *service) GetTermConditionById(id string) (entity.TermAndCondition, error) {
	var out *entity.TermAndCondition
	result, err := s.termAndConditionRepo.FindTermConditionById(id)
	if err != nil {
		return *out, err
	}
	out = &result
	return *out, err
}
