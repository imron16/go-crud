package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type InsertTermConditionRequest struct {
	Title     string `json:"title" validate:"required"`
	TitleEn   string `json:"titleEn" validate:"required"`
	Content   string `json:"content" validate:"required"`
	ContentEn string `json:"contentEn" validate:"required"`
}

func NewAdminRegisterRequest(in *gin.Context) (out InsertTermConditionRequest, err error) {
	body, _ := ioutil.ReadAll(in.Request.Body)
	err = json.Unmarshal(body, &out)
	if err != nil {
		return out, errors.New("BAD_PROCESSING")
	}

	return out, err
}
