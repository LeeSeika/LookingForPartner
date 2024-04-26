package server

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lookingforpartner/idl/pb/user"
	"lookingforpartner/pkg/discovery"
	"lookingforpartner/service/user/rpc/internal/handler"
	"lookingforpartner/service/user/rpc/internal/svc"
	"net"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	server *grpc.Server
	ctx    context.Context

	user.UnimplementedUserServiceServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
		server: grpc.NewServer(),
		ctx:    context.Background(),
	}
}

func (s *UserServer) UserLogin(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (s *UserServer) UserSignup(ctx context.Context, req *user.UserSignupRequest) (*user.UserSignupResponse, error) {
	signupHandler := handler.NewSignupHandler(s.svcCtx, ctx)
	return signupHandler.Signup(req)
}
func (s *UserServer) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}

func (s *UserServer) MustStart() {
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	name := s.svcCtx.Config.Server.Name
	etcdAddress := s.svcCtx.Config.Etcd.Host
	dialTimeout := s.svcCtx.Config.Etcd.DialTimeout
	srvTTL := s.svcCtx.Config.Etcd.TTL

	grpcAddress := fmt.Sprintf("%s:%d", s.svcCtx.Config.Server.Host, s.svcCtx.Config.Server.Port)

	server := s.server

	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, dialTimeout)
	srvInfo := discovery.Server{
		Name: name,
		Addr: grpcAddress,
	}

	user.RegisterUserServiceServer(server, s)
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(fmt.Sprintf("listen failed, err: %v", err))
	}
	if err := etcdRegister.Register(ctx, srvInfo, srvTTL); err != nil {
		panic(fmt.Sprintf("register failed, err: %v", err))
	}
	zap.L().Info("server started listening", zap.Any("grpc addrs", grpcAddress))
	if err := server.Serve(lis); err != nil {
		panic(fmt.Sprintf("serve failed, err: %v", err))
	}
}

func (s *UserServer) Stop() {
	s.server.Stop()
}
