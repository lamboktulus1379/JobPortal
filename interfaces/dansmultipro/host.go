package dansmultipro

import (
	"context"
	"encoding/json"
	"fmt"
	"jobportal/interfaces"
	"jobportal/interfaces/dansmultipro/models"
	"strings"
)

type IDansmultiproHost interface {
	Search(ctx context.Context, req models.ReqQueryParam) ([]models.Position, error)
	Get(ctx context.Context, ID string) (models.Position, error)
}

type DansmultiproHost struct {
	id   string
	host string
}

func NewDansmultiproHost(host string) IDansmultiproHost {
	return &DansmultiproHost{host: host}
}

func (dansmultiproHost *DansmultiproHost) Search(ctx context.Context, req models.ReqQueryParam) ([]models.Position, error) {
	var res []models.Position

	endpoint := "/api/recruitment/positions.json"
	method := "GET"

	hostClient := interfaces.NewHost(dansmultiproHost.host, endpoint, method, nil, nil, req)
	byteData, statusCode, err := hostClient.HTTPGet()
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(byteData, &res); err != nil {
		return res, err
	}

	if statusCode < 200 && statusCode > 299 {
		return res, fmt.Errorf("Something occurred with server")
	}

	return res, nil
}

func (dansmultiproHost *DansmultiproHost) Get(ctx context.Context, ID string) (models.Position, error) {
	var res models.Position

	endpoint := "/api/recruitment/positions/{ID}"
	method := "GET"

	hostClient := interfaces.NewHost(dansmultiproHost.host, strings.ReplaceAll(endpoint, "{ID}", ID), method, nil, nil, nil)
	byteData, statusCode, err := hostClient.HTTPGet()
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(byteData, &res); err != nil {
		return res, err
	}

	if statusCode < 200 && statusCode > 299 {
		return res, fmt.Errorf("Something occurred with server")
	}

	return res, nil
}
