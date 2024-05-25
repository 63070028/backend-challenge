package services

import (
	"context"
)

type BeefService interface {
	GetBeef() (*BeefResponse, error) 
}

type beefService struct {
	beefSerivceClient BeefSerivceClient
}

func NewBeefService(beefSerivceClient BeefSerivceClient) BeefService {
	return beefService{beefSerivceClient}
}

func (base beefService) GetBeef() (*BeefResponse, error) {
	req := BeefRequest{}

	res, err := base.beefSerivceClient.GetBeef(context.Background(), &req)
	if err != nil {
		return nil, err
	}

	return res, err
}
