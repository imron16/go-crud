package mysql

import (
	"errors"
	"go-crud/domain/term_and_condition"
	"go-crud/entity"
	"go-crud/shared/constant"

	"github.com/jinzhu/gorm"
)

type termAndConditionRepo struct {
	db *gorm.DB
}

func SetupTermAndConditionRepo(db *gorm.DB) term_and_condition.Repository {
	return &termAndConditionRepo{db: db}
}

func (a *termAndConditionRepo) All() (out []entity.TermAndCondition, err error) {
	err = a.db.Model(&entity.TermAndCondition{}).Find(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New(constant.NoFindAllTermCondition)
	}
	return out, err
}

func (a *termAndConditionRepo) SaveOrUpdateTermCondition(req entity.TermAndCondition) error {
	operation := a.db.Model(entity.TermAndCondition{}).Save(&req).Error
	if operation != nil {
		return errors.New(constant.CantSaveUpdateData)
	}
	return operation
}

func (a *termAndConditionRepo) FindTermConditionById(id string) (out entity.TermAndCondition, err error) {
	err = a.db.Model(&entity.TermAndCondition{}).Where("id = ? ", id).First(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New(constant.CantFindTermCondition)
	}
	return out, err
}

func (a *termAndConditionRepo) FindTermConditionByTitle(title string) (out entity.TermAndCondition, err error) {
	err = a.db.Model(&entity.TermAndCondition{}).Where("title = ? ", title).First(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New(constant.CantFindTermCondition)
	}
	return out, err
}

func (a *termAndConditionRepo) DeleteTermsCondition(id string) error {
	err := a.db.Where("id = ?", id).Delete(entity.TermAndCondition{}).Error
	if err != nil {
		return errors.New(constant.CantDeleteData)
	}
	return err
}
