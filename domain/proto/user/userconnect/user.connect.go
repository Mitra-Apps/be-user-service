// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/user/user.proto

package userconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	user "github.com/Mitra-Apps/be-user-service/domain/proto/user"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "proto.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceGetUsersProcedure is the fully-qualified name of the UserService's GetUsers RPC.
	UserServiceGetUsersProcedure = "/proto.UserService/GetUsers"
	// UserServiceLoginProcedure is the fully-qualified name of the UserService's Login RPC.
	UserServiceLoginProcedure = "/proto.UserService/Login"
	// UserServiceRegisterProcedure is the fully-qualified name of the UserService's Register RPC.
	UserServiceRegisterProcedure = "/proto.UserService/Register"
	// UserServiceCreateRoleProcedure is the fully-qualified name of the UserService's CreateRole RPC.
	UserServiceCreateRoleProcedure = "/proto.UserService/CreateRole"
	// UserServiceGetRoleProcedure is the fully-qualified name of the UserService's GetRole RPC.
	UserServiceGetRoleProcedure = "/proto.UserService/GetRole"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userServiceServiceDescriptor          = user.File_proto_user_user_proto.Services().ByName("UserService")
	userServiceGetUsersMethodDescriptor   = userServiceServiceDescriptor.Methods().ByName("GetUsers")
	userServiceLoginMethodDescriptor      = userServiceServiceDescriptor.Methods().ByName("Login")
	userServiceRegisterMethodDescriptor   = userServiceServiceDescriptor.Methods().ByName("Register")
	userServiceCreateRoleMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("CreateRole")
	userServiceGetRoleMethodDescriptor    = userServiceServiceDescriptor.Methods().ByName("GetRole")
)

// UserServiceClient is a client for the proto.UserService service.
type UserServiceClient interface {
	GetUsers(context.Context, *connect.Request[user.GetUsersRequest]) (*connect.Response[user.GetUsersResponse], error)
	Login(context.Context, *connect.Request[user.UserLoginRequest]) (*connect.Response[user.SuccessResponse], error)
	Register(context.Context, *connect.Request[user.UserRegisterRequest]) (*connect.Response[user.SuccessResponse], error)
	CreateRole(context.Context, *connect.Request[user.Role]) (*connect.Response[user.SuccessResponse], error)
	GetRole(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[user.SuccessResponse], error)
}

// NewUserServiceClient constructs a client for the proto.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		getUsers: connect.NewClient[user.GetUsersRequest, user.GetUsersResponse](
			httpClient,
			baseURL+UserServiceGetUsersProcedure,
			connect.WithSchema(userServiceGetUsersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		login: connect.NewClient[user.UserLoginRequest, user.SuccessResponse](
			httpClient,
			baseURL+UserServiceLoginProcedure,
			connect.WithSchema(userServiceLoginMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		register: connect.NewClient[user.UserRegisterRequest, user.SuccessResponse](
			httpClient,
			baseURL+UserServiceRegisterProcedure,
			connect.WithSchema(userServiceRegisterMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createRole: connect.NewClient[user.Role, user.SuccessResponse](
			httpClient,
			baseURL+UserServiceCreateRoleProcedure,
			connect.WithSchema(userServiceCreateRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getRole: connect.NewClient[emptypb.Empty, user.SuccessResponse](
			httpClient,
			baseURL+UserServiceGetRoleProcedure,
			connect.WithSchema(userServiceGetRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	getUsers   *connect.Client[user.GetUsersRequest, user.GetUsersResponse]
	login      *connect.Client[user.UserLoginRequest, user.SuccessResponse]
	register   *connect.Client[user.UserRegisterRequest, user.SuccessResponse]
	createRole *connect.Client[user.Role, user.SuccessResponse]
	getRole    *connect.Client[emptypb.Empty, user.SuccessResponse]
}

// GetUsers calls proto.UserService.GetUsers.
func (c *userServiceClient) GetUsers(ctx context.Context, req *connect.Request[user.GetUsersRequest]) (*connect.Response[user.GetUsersResponse], error) {
	return c.getUsers.CallUnary(ctx, req)
}

// Login calls proto.UserService.Login.
func (c *userServiceClient) Login(ctx context.Context, req *connect.Request[user.UserLoginRequest]) (*connect.Response[user.SuccessResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// Register calls proto.UserService.Register.
func (c *userServiceClient) Register(ctx context.Context, req *connect.Request[user.UserRegisterRequest]) (*connect.Response[user.SuccessResponse], error) {
	return c.register.CallUnary(ctx, req)
}

// CreateRole calls proto.UserService.CreateRole.
func (c *userServiceClient) CreateRole(ctx context.Context, req *connect.Request[user.Role]) (*connect.Response[user.SuccessResponse], error) {
	return c.createRole.CallUnary(ctx, req)
}

// GetRole calls proto.UserService.GetRole.
func (c *userServiceClient) GetRole(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.SuccessResponse], error) {
	return c.getRole.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the proto.UserService service.
type UserServiceHandler interface {
	GetUsers(context.Context, *connect.Request[user.GetUsersRequest]) (*connect.Response[user.GetUsersResponse], error)
	Login(context.Context, *connect.Request[user.UserLoginRequest]) (*connect.Response[user.SuccessResponse], error)
	Register(context.Context, *connect.Request[user.UserRegisterRequest]) (*connect.Response[user.SuccessResponse], error)
	CreateRole(context.Context, *connect.Request[user.Role]) (*connect.Response[user.SuccessResponse], error)
	GetRole(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[user.SuccessResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceGetUsersHandler := connect.NewUnaryHandler(
		UserServiceGetUsersProcedure,
		svc.GetUsers,
		connect.WithSchema(userServiceGetUsersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceLoginHandler := connect.NewUnaryHandler(
		UserServiceLoginProcedure,
		svc.Login,
		connect.WithSchema(userServiceLoginMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceRegisterHandler := connect.NewUnaryHandler(
		UserServiceRegisterProcedure,
		svc.Register,
		connect.WithSchema(userServiceRegisterMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceCreateRoleHandler := connect.NewUnaryHandler(
		UserServiceCreateRoleProcedure,
		svc.CreateRole,
		connect.WithSchema(userServiceCreateRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetRoleHandler := connect.NewUnaryHandler(
		UserServiceGetRoleProcedure,
		svc.GetRole,
		connect.WithSchema(userServiceGetRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceGetUsersProcedure:
			userServiceGetUsersHandler.ServeHTTP(w, r)
		case UserServiceLoginProcedure:
			userServiceLoginHandler.ServeHTTP(w, r)
		case UserServiceRegisterProcedure:
			userServiceRegisterHandler.ServeHTTP(w, r)
		case UserServiceCreateRoleProcedure:
			userServiceCreateRoleHandler.ServeHTTP(w, r)
		case UserServiceGetRoleProcedure:
			userServiceGetRoleHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) GetUsers(context.Context, *connect.Request[user.GetUsersRequest]) (*connect.Response[user.GetUsersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.UserService.GetUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) Login(context.Context, *connect.Request[user.UserLoginRequest]) (*connect.Response[user.SuccessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.UserService.Login is not implemented"))
}

func (UnimplementedUserServiceHandler) Register(context.Context, *connect.Request[user.UserRegisterRequest]) (*connect.Response[user.SuccessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.UserService.Register is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateRole(context.Context, *connect.Request[user.Role]) (*connect.Response[user.SuccessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.UserService.CreateRole is not implemented"))
}

func (UnimplementedUserServiceHandler) GetRole(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[user.SuccessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.UserService.GetRole is not implemented"))
}
