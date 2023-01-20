package termcondition

import (
	"go-crud/entity"
)

type Service interface {
	InsertTermCondition(title, titleEn, content, contentEn string) error
	UpdateTermCondition(id, title, titleEn, content, contentEn string) error
	GetAllTermCondition() ([]entity.TermAndCondition, error)
	GetTermConditionById(id string) (entity.TermAndCondition, error)
	DeleteTermCondition(id string) error
}
