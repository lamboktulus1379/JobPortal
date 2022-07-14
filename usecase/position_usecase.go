package usecase

import (
	"context"
	"fmt"
	"jobportal/domain/dto"
	"jobportal/domain/model"
	"jobportal/interfaces/dansmultipro"
	"jobportal/interfaces/dansmultipro/models"
)

type IPositionUsecase interface {
	Search(ctx context.Context, reqQueryParam model.ReqQueryParam) dto.ResPositions
	Get(ctx context.Context, ID string) dto.ResPosition
}

type PositionUsecase struct {
	dansmultiproHost dansmultipro.IDansmultiproHost
}

func NewPositionUsecase(dansmultiproHost dansmultipro.IDansmultiproHost) IPositionUsecase {
	return &PositionUsecase{dansmultiproHost: dansmultiproHost}
}

func (postionUsecase *PositionUsecase) Search(ctx context.Context, reqQueryParam model.ReqQueryParam) dto.ResPositions {
	var res dto.ResPositions

	req := models.ReqQueryParam{
		Description: reqQueryParam.Description,
		Location:    reqQueryParam.Location,
		FullTime:    reqQueryParam.FullTime,
		Page:        reqQueryParam.Page,
	}
	fmt.Println("Description:", req.Description)

	positions, err := postionUsecase.dansmultiproHost.Search(ctx, req)
	if err != nil {
		res.ResponseCode = "500"
		res.ResponseMessage = err.Error()
		return res
	}

	positionDtos := make([]dto.Position, 0)
	for _, position := range positions {
		if position.ID != "" {
			positionDto := dto.Position{
				ID:          position.ID,
				Type:        position.Type,
				URL:         position.URL,
				CreatedAt:   position.CreatedAt,
				Company:     position.Company,
				CompanyURL:  position.CompanyURL,
				Location:    position.Location,
				Title:       position.Title,
				HowToApply:  position.HowToApply,
				CompanyLogo: position.CompanyLogo,
			}

			positionDtos = append(positionDtos, positionDto)
		}

	}

	pagination := dto.Pagination{
		PageNumber:  reqQueryParam.Page,
		TotalRecord: len(positionDtos),
	}

	res.Pagination = &pagination
	res.Data = &positionDtos

	res.ResponseCode = "200"
	res.ResponseMessage = "Success"

	return res
}

func (postionUsecase *PositionUsecase) Get(ctx context.Context, ID string) dto.ResPosition {
	var res dto.ResPosition

	position, err := postionUsecase.dansmultiproHost.Get(ctx, ID)
	if err != nil {
		res.ResponseCode = "500"
		res.ResponseMessage = err.Error()
		return res
	}

	positionDto := dto.Position{
		ID:          position.ID,
		Type:        position.Type,
		URL:         position.URL,
		CreatedAt:   position.CreatedAt,
		Company:     position.Company,
		CompanyURL:  position.CompanyURL,
		Location:    position.Location,
		Title:       position.Title,
		HowToApply:  position.HowToApply,
		CompanyLogo: position.CompanyLogo,
	}

	res.Data = &positionDto
	res.ResponseCode = "200"
	res.ResponseMessage = "Success"

	return res
}
