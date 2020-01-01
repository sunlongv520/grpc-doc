package services

import (
	"context"
	"time"
)

type UserService struct {
}
//普通方法
func(*UserService) GetUserScore(ctx context.Context, in *UserScoreRequest)  (*UserScoreResponse, error){
	var score int32=101
	users:=make([]*UserInfo,0)
	for _,user:=range in.Users{
		user.UserScore=score
		score++
		users=append(users,user)
	}
	return &UserScoreResponse{Users:users},nil
}
func(*UserService) 	GetUserScoreByServerStream(in *UserScoreRequest,stream UserService_GetUserScoreByServerStreamServer) error {
	var score int32=101
	users:=make([]*UserInfo,0)
	for index,user:=range in.Users{
		user.UserScore=score
		score++
		users=append(users,user)

		if (index+1) % 2==0 && index>0{
			err:=stream.Send(&UserScoreResponse{Users:users})
			if err!=nil{
				return err
			}
			users=(users)[0:0]
		}
		time.Sleep(time.Second*1)
	}
	if len(users)>0{
		err:=stream.Send(&UserScoreResponse{Users:users})
		if err!=nil{
			return err
		}
	}
	return nil
}

