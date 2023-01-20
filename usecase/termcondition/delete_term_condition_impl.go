package termcondition

func (s *service) DeleteTermCondition(id string) error {
	err := s.termAndConditionRepo.DeleteTermsCondition(id)
	if err != nil {
		return err
	}
	return nil
}
