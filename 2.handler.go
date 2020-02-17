package main

import (
	pb "github.com/ruandao/micro-shippy-user-service/proto/user"
	"golang.org/x/net/context"
)

type handler struct {
	repository Repository
}

func (srv *handler) Create(ctx context.Context, user *pb.User, resp *pb.Response) error {
	if err := srv.repository.Create(ctx, user); err != nil {
		return err
	}
	resp.User = user
	return nil
}

func (srv *handler) Get(ctx context.Context, user *pb.User, resp *pb.Response) error {
	storeUser, err := srv.repository.Get(ctx, user.Id)
	if err != nil {
		return err
	}
	resp.User = UnmarshalUser(storeUser)
	return nil
}

func (srv *handler) GetAll(ctx context.Context, _ *pb.Request, resp *pb.Response) error {
	storeUsers, err := srv.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	users := make([]*pb.User, 0, len(storeUsers))
	for _, user := range storeUsers {
		users = append(users, UnmarshalUser(user))
	}
	resp.Users = users
	return nil
}

func (srv *handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	_, err := srv.repository.GetByEmailAndPassword(ctx, MarshalUser(req))
	if err != nil {
		return err
	}
	resp.Token = "// todo"
	return nil
}

func (srv *handler) ValidateToken(context.Context, *pb.Token, *pb.Token) error {
	// get user from token
	return nil
}
