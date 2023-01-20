package termcondition

import (
	"errors"
	"go-crud/shared/constant"
	"strings"
)

func (s *service) UpdateTermCondition(id, title, titleEn, content, contentEn string) error {
	result, err := s.termAndConditionRepo.FindTermConditionById(id)
	if err != nil {
		return errors.New(constant.InternalServerError)
	}
	if len(strings.TrimSpace(result.Title)) == 0 {
		return errors.New(constant.TitleAlreadyExist)
	}
	result.Title = title
	result.TitleEn = titleEn
	result.Content = content
	result.ContentEn = contentEn

	err = s.termAndConditionRepo.SaveOrUpdateTermCondition(result)
	if err != nil {
		return err
	}
	return nil
}
