// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user.proto

package user

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	UserInfo(ctx context.Context, in *NameData, opts ...client.CallOption) (*Response, error)
	UpdateUserName(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*UpdateResp, error)
	UploadAvatar(ctx context.Context, in *UploadReq, opts ...client.CallOption) (*UploadResp, error)
	AuthUpdate(ctx context.Context, in *AuthReq, opts ...client.CallOption) (*CallResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) UserInfo(ctx context.Context, in *NameData, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.UserInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUserName(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*UpdateResp, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUserName", in)
	out := new(UpdateResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UploadAvatar(ctx context.Context, in *UploadReq, opts ...client.CallOption) (*UploadResp, error) {
	req := c.c.NewRequest(c.name, "User.UploadAvatar", in)
	out := new(UploadResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AuthUpdate(ctx context.Context, in *AuthReq, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "User.AuthUpdate", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	UserInfo(context.Context, *NameData, *Response) error
	UpdateUserName(context.Context, *UpdateReq, *UpdateResp) error
	UploadAvatar(context.Context, *UploadReq, *UploadResp) error
	AuthUpdate(context.Context, *AuthReq, *CallResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		UserInfo(ctx context.Context, in *NameData, out *Response) error
		UpdateUserName(ctx context.Context, in *UpdateReq, out *UpdateResp) error
		UploadAvatar(ctx context.Context, in *UploadReq, out *UploadResp) error
		AuthUpdate(ctx context.Context, in *AuthReq, out *CallResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) UserInfo(ctx context.Context, in *NameData, out *Response) error {
	return h.UserHandler.UserInfo(ctx, in, out)
}

func (h *userHandler) UpdateUserName(ctx context.Context, in *UpdateReq, out *UpdateResp) error {
	return h.UserHandler.UpdateUserName(ctx, in, out)
}

func (h *userHandler) UploadAvatar(ctx context.Context, in *UploadReq, out *UploadResp) error {
	return h.UserHandler.UploadAvatar(ctx, in, out)
}

func (h *userHandler) AuthUpdate(ctx context.Context, in *AuthReq, out *CallResponse) error {
	return h.UserHandler.AuthUpdate(ctx, in, out)
}
