package services

import (
	"context"
)

type OrdersService struct {

}

func(this *OrdersService)NewOrder(ctx context.Context,orderRequest *OrderRequest) (*OrderResponse, error)   {
	err:= orderRequest.OrderMain.Validate()
	if err!=nil{
		return &OrderResponse{
			Status:"error",
			Message:err.Error(),
		},nil
	}
	return &OrderResponse{
		Status:"OK",
		Message:"success",
	},nil
}
