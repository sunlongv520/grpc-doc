package services

import (
	"context"
	"fmt"
)

type OrdersService struct {

}

func(this *OrdersService)NewOrder(ctx context.Context,orderMain *OrderMain) (*OrderResponse, error)   {
	fmt.Println(orderMain)
	return &OrderResponse{
		Status:"OK",
		Message:"success",
	},nil
}