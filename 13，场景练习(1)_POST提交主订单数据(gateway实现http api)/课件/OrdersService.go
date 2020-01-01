package services

import (
	"context"
	"fmt"

)

type OrdersService struct {

}

func(this *OrdersService)NewOrder(ctx context.Context,orderRequest *OrderRequest) (*OrderResponse, error)   {
	fmt.Println(orderRequest.OrderMain)
	return &OrderResponse{
		Status:"OK",
		Message:"success",
	},nil
}
