package api

import (
	"library/falcon/service"
)


type Api struct {
	*service.FalconApi
}

func New() *Api {

	service := &service.FalconApi{
	}

	return &Api{service}
}
