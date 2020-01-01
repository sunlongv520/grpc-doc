package services

import (
	"context"
)

type ProdService struct {}

func(this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	  return &ProdResponse{ProdStock:28},nil
}
func(this *ProdService) GetProdStocks(ctx context.Context,size *QuerySize) (*ProdResponseList, error)  {
			Prodres:= []*ProdResponse{
					&ProdResponse{ProdStock:28},
					&ProdResponse{ProdStock:29},
					&ProdResponse{ProdStock:30},
					&ProdResponse{ProdStock:31},
			}
			return &ProdResponseList{
				Prodres:Prodres,
			},nil
}


