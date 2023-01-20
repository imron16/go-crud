package container

import (
	"go-crud/shared/database"

	"go-crud/infrastructure/mysql"

	"go-crud/usecase/termcondition"

	"github.com/jinzhu/gorm"
)

type Container struct {
	TermConditionService termcondition.Service
}

func SetupContainer(tx *gorm.DB) Container {

	database.Migrate(tx)
	termConditionService := termcondition.NewTermsService(mysql.SetupTermAndConditionRepo(tx))
	return Container{TermConditionService: termConditionService}
}
