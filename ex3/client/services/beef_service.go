package services

import (
	"context"
	"fmt"
)

type BeefService interface {
	GetBeef() error
}

type beefService struct {
	beefSerivceClient BeefSerivceClient
}

func NewBeefService(beefSerivceClient BeefSerivceClient) BeefService {
	return beefService{beefSerivceClient}
}

func (base beefService) GetBeef() error {
	req := BeefRequest{}

	res, err := base.beefSerivceClient.GetBeef(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Printf("Response: %v", res)

	return nil
}
