package server

import (
	"fmt"
	"go-crud/entity"
	"go-crud/usecase/termcondition"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type termConditionHandler struct {
	service termcondition.Service
}

func newTermConditionHanlder(service termcondition.Service) *termConditionHandler {
	return &termConditionHandler{service: service}
}

func (h *termConditionHandler) GetAllData() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		result, err := h.service.GetAllTermCondition()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.JSON(http.StatusOK, gin.H{"Data": result})
		}
	}
}

func (h *termConditionHandler) CreateData() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var requestBody entity.TermAndCondition
		err := ctx.BindJSON(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
		} else {
			switch {
			case requestBody.Title == "":
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "title is required"})
			case requestBody.TitleEn == "":
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "titleEn is required"})
			case requestBody.Content == "":
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "content is required"})
			case requestBody.ContentEn == "":
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "contentEn is required"})
			default:
				err := h.service.InsertTermCondition(requestBody.Title, requestBody.TitleEn, requestBody.Content, requestBody.ContentEn)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
				} else {
					ctx.JSON(http.StatusOK, gin.H{"message": "Sucess"})
				}
			}
		}
	}
}

func (h *termConditionHandler) GetDataById() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var requestBody entity.TermAndCondition
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
		} else {
			switch {
			case requestBody.Id == 0:
				c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
			default:
				result, err := h.service.GetTermConditionById(strconv.Itoa(int(requestBody.Id)))
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
				} else {
					c.JSON(http.StatusOK, gin.H{"Data": result})
				}
			}
		}
	}
}

func (h *termConditionHandler) UpdateDataById() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var requestBody entity.TermAndCondition
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
		} else {
			switch {
			case requestBody.Id == 0:
				c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
			default:
				err := h.service.UpdateTermCondition(strconv.Itoa(int(requestBody.Id)), requestBody.Title, requestBody.TitleEn, requestBody.Content, requestBody.ContentEn)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
				} else {
					c.JSON(http.StatusOK, gin.H{"message": "Sucess"})
				}
			}
		}
	}
}

func (h *termConditionHandler) DeleteData() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var requestBody entity.TermAndCondition
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
		} else {
			switch {
			case requestBody.Id == 0:
				c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
			default:
				err := h.service.DeleteTermCondition(strconv.Itoa(int(requestBody.Id)))
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
				} else {
					c.JSON(http.StatusOK, gin.H{"message": "Sucess"})
				}
			}
		}
	}
}
