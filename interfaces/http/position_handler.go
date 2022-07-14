package http

import (
	"jobportal/constants"
	"jobportal/domain/dto"
	"jobportal/domain/model"
	"jobportal/usecase"

	"github.com/gin-gonic/gin"
)

type IPositionHandler interface {
	Search(c *gin.Context)
	Get(c *gin.Context)
}

type PositionHandler struct {
	positionUsecase usecase.IPositionUsecase
}

func NewPositionHandler(positionUsecase usecase.IPositionUsecase) IPositionHandler {
	return &PositionHandler{positionUsecase: positionUsecase}
}

func (handler *PositionHandler) Search(c *gin.Context) {
	var res dto.ResPositions
	var reqQueryParam model.ReqQueryParam

	if err := c.ShouldBindQuery(&reqQueryParam); err != nil {
		res.ResponseCode = "400"
		res.ResponseMessage = "Bad request"
		c.JSON(constants.ResponseCode[res.ResponseCode], res)
		return
	}

	res = handler.positionUsecase.Search(c.Request.Context(), reqQueryParam)

	c.JSON(constants.ResponseCode[res.ResponseCode], res)
}

func (handler *PositionHandler) Get(c *gin.Context) {
	var res dto.ResPosition
	var reqURIParam model.ReqURIParam

	if err := c.ShouldBindUri(&reqURIParam); err != nil {
		res.ResponseCode = "400"
		res.ResponseMessage = "Bad request"
		c.JSON(constants.ResponseCode[res.ResponseCode], res)
		return
	}

	res = handler.positionUsecase.Get(c.Request.Context(), reqURIParam.ID)
	c.JSON(constants.ResponseCode[res.ResponseCode], res)
}
