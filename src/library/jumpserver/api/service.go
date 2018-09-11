package api

import (
	"library/jumpserver/service"
)


type Api struct {
	*service.JumpserverApi
}

func New() *Api {

	service := &service.JumpserverApi{
	}

	return &Api{service}
}
