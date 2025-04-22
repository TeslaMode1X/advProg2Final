package grpc

import (
	"context"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/user"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/crypto"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/security"
	_ "github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceGrpc struct {
	user.UnimplementedUserServiceServer
	userService interfaces.UserService
}

func NewUserServiceGrpc(userService interfaces.UserService) *UserServiceGrpc {
	return &UserServiceGrpc{
		userService: userService,
	}
}

func (g *UserServiceGrpc) UserLogin(ctx context.Context, req *user.RequestUserLogin) (*user.UserResponse, error) {
	const op = "service.grpc.UserLogin"

	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "Email and password are required")
	}

	id, err := g.userService.UserLoginService(req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	token, err := security.CreateToken(id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to generate token")
	}

	return &user.UserResponse{
		Id:    id.String(),
		Token: token,
	}, nil
}

func (g *UserServiceGrpc) UserRegistration(ctx context.Context, req *user.RequestUserRegistration) (*user.UserResponse, error) {
	const op = "grpc.user.UserRegistration"

	if req.Email == "" || req.Password == "" || req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "Username, email and password are required")
	}

	hashedPassword, err := crypto.HashPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to hash password")
	}

	userModel := model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	id, err := g.userService.UserRegisterService(userModel)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.UserResponse{
		Id: id.String(),
	}, nil
}

func (g *UserServiceGrpc) UserGetById(ctx context.Context, req *user.RequestUserGetById) (*user.ResponseUserGetById, error) {
	const op = "grpc.user.UserGetById"

	model, err := g.userService.UserGetByIdService(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.ResponseUserGetById{
		Id:    model.ID.String(),
		Name:  model.Username,
		Email: model.Email,
	}, nil
}

func (g *UserServiceGrpc) UserDeleteById(ctx context.Context, req *user.RequestUserGetById) (*user.Empty, error) {
	var empty user.Empty

	err := g.userService.UserDeleteByIdService(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty, nil
}

func (g *UserServiceGrpc) UserExists(ctx context.Context, req *user.RequestUserGetById) (*user.ResponseUserExists, error) {
	const op = "grpc.user.UserExists"

	exists, err := g.userService.UserExistsService(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.ResponseUserExists{Exists: exists}, nil
}

func (g *UserServiceGrpc) UserChangePassword(ctx context.Context, req *user.RequestUserChangePassword) (*user.Empty, error) {
	const op = "grpc.user.UserChangePassword"

	err := g.userService.UserUpdatePasswordService(req.Id, req.OldPassword, req.NewPassword)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.Empty{}, nil
}
