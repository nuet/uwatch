package api

import (
	"library/cmdb/service"
)


type Api struct {
	*service.CmdbApi
}

func New() *Api {

	service := &service.CmdbApi{
	}

	return &Api{service}
}
