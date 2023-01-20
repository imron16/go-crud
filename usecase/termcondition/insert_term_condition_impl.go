package termcondition

import (
	"errors"
	"go-crud/entity"
	"go-crud/shared/constant"
)

func (s *service) InsertTermCondition(title, titleEn, content, contentEn string) error {
	result, err := s.termAndConditionRepo.FindTermConditionByTitle(title)
	if err != nil {
		return errors.New(constant.InternalServerError)
	}

	if len(result.Title) > 0 {
		return errors.New(constant.TitleAlreadyExist)
	}

	termConditions := entity.TermAndCondition{}
	termConditions.Title = title
	termConditions.TitleEn = titleEn
	termConditions.Content = content
	termConditions.ContentEn = contentEn

	err = s.termAndConditionRepo.SaveOrUpdateTermCondition(termConditions)
	if err != nil {
		return err
	}

	return nil

}
