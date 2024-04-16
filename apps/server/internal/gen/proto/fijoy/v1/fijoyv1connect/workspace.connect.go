// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: fijoy/v1/workspace.proto

package fijoyv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "fijoy/internal/gen/proto/fijoy/v1"
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
	// WorkspaceServiceName is the fully-qualified name of the WorkspaceService service.
	WorkspaceServiceName = "fijoy.v1.WorkspaceService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// WorkspaceServiceCreateWorkspaceProcedure is the fully-qualified name of the WorkspaceService's
	// CreateWorkspace RPC.
	WorkspaceServiceCreateWorkspaceProcedure = "/fijoy.v1.WorkspaceService/CreateWorkspace"
	// WorkspaceServiceGetWorkspacesProcedure is the fully-qualified name of the WorkspaceService's
	// GetWorkspaces RPC.
	WorkspaceServiceGetWorkspacesProcedure = "/fijoy.v1.WorkspaceService/GetWorkspaces"
	// WorkspaceServiceGetWorkspaceByIdProcedure is the fully-qualified name of the WorkspaceService's
	// GetWorkspaceById RPC.
	WorkspaceServiceGetWorkspaceByIdProcedure = "/fijoy.v1.WorkspaceService/GetWorkspaceById"
	// WorkspaceServiceGetWorkspaceByNamespaceProcedure is the fully-qualified name of the
	// WorkspaceService's GetWorkspaceByNamespace RPC.
	WorkspaceServiceGetWorkspaceByNamespaceProcedure = "/fijoy.v1.WorkspaceService/GetWorkspaceByNamespace"
	// WorkspaceServiceUpdateWorkspaceNameProcedure is the fully-qualified name of the
	// WorkspaceService's UpdateWorkspaceName RPC.
	WorkspaceServiceUpdateWorkspaceNameProcedure = "/fijoy.v1.WorkspaceService/UpdateWorkspaceName"
	// WorkspaceServiceUpdateWorkspaceNamespaceProcedure is the fully-qualified name of the
	// WorkspaceService's UpdateWorkspaceNamespace RPC.
	WorkspaceServiceUpdateWorkspaceNamespaceProcedure = "/fijoy.v1.WorkspaceService/UpdateWorkspaceNamespace"
	// WorkspaceServiceUpdatePrimaryCurrencyProcedure is the fully-qualified name of the
	// WorkspaceService's UpdatePrimaryCurrency RPC.
	WorkspaceServiceUpdatePrimaryCurrencyProcedure = "/fijoy.v1.WorkspaceService/UpdatePrimaryCurrency"
	// WorkspaceServiceUpdateLocaleProcedure is the fully-qualified name of the WorkspaceService's
	// UpdateLocale RPC.
	WorkspaceServiceUpdateLocaleProcedure = "/fijoy.v1.WorkspaceService/UpdateLocale"
	// WorkspaceServiceDeleteWorkspaceProcedure is the fully-qualified name of the WorkspaceService's
	// DeleteWorkspace RPC.
	WorkspaceServiceDeleteWorkspaceProcedure = "/fijoy.v1.WorkspaceService/DeleteWorkspace"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	workspaceServiceServiceDescriptor                        = v1.File_fijoy_v1_workspace_proto.Services().ByName("WorkspaceService")
	workspaceServiceCreateWorkspaceMethodDescriptor          = workspaceServiceServiceDescriptor.Methods().ByName("CreateWorkspace")
	workspaceServiceGetWorkspacesMethodDescriptor            = workspaceServiceServiceDescriptor.Methods().ByName("GetWorkspaces")
	workspaceServiceGetWorkspaceByIdMethodDescriptor         = workspaceServiceServiceDescriptor.Methods().ByName("GetWorkspaceById")
	workspaceServiceGetWorkspaceByNamespaceMethodDescriptor  = workspaceServiceServiceDescriptor.Methods().ByName("GetWorkspaceByNamespace")
	workspaceServiceUpdateWorkspaceNameMethodDescriptor      = workspaceServiceServiceDescriptor.Methods().ByName("UpdateWorkspaceName")
	workspaceServiceUpdateWorkspaceNamespaceMethodDescriptor = workspaceServiceServiceDescriptor.Methods().ByName("UpdateWorkspaceNamespace")
	workspaceServiceUpdatePrimaryCurrencyMethodDescriptor    = workspaceServiceServiceDescriptor.Methods().ByName("UpdatePrimaryCurrency")
	workspaceServiceUpdateLocaleMethodDescriptor             = workspaceServiceServiceDescriptor.Methods().ByName("UpdateLocale")
	workspaceServiceDeleteWorkspaceMethodDescriptor          = workspaceServiceServiceDescriptor.Methods().ByName("DeleteWorkspace")
)

// WorkspaceServiceClient is a client for the fijoy.v1.WorkspaceService service.
type WorkspaceServiceClient interface {
	CreateWorkspace(context.Context, *connect.Request[v1.CreateWorkspaceRequest]) (*connect.Response[v1.Workspace], error)
	GetWorkspaces(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.Workspaces], error)
	GetWorkspaceById(context.Context, *connect.Request[v1.GetWorkspaceByIdRequest]) (*connect.Response[v1.Workspace], error)
	GetWorkspaceByNamespace(context.Context, *connect.Request[v1.GetWorkspaceByNamespaceRequest]) (*connect.Response[v1.Workspace], error)
	UpdateWorkspaceName(context.Context, *connect.Request[v1.UpdateWorkspaceNameRequest]) (*connect.Response[v1.Workspace], error)
	UpdateWorkspaceNamespace(context.Context, *connect.Request[v1.UpdateWorkspaceNamespaceRequest]) (*connect.Response[v1.Workspace], error)
	UpdatePrimaryCurrency(context.Context, *connect.Request[v1.UpdatePrimaryCurrencyRequest]) (*connect.Response[v1.Workspace], error)
	UpdateLocale(context.Context, *connect.Request[v1.UpdateLocaleRequest]) (*connect.Response[v1.Workspace], error)
	DeleteWorkspace(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.Workspace], error)
}

// NewWorkspaceServiceClient constructs a client for the fijoy.v1.WorkspaceService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkspaceServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) WorkspaceServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &workspaceServiceClient{
		createWorkspace: connect.NewClient[v1.CreateWorkspaceRequest, v1.Workspace](
			httpClient,
			baseURL+WorkspaceServiceCreateWorkspaceProcedure,
			connect.WithSchema(workspaceServiceCreateWorkspaceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getWorkspaces: connect.NewClient[emptypb.Empty, v1.Workspaces](
			httpClient,
			baseURL+WorkspaceServiceGetWorkspacesProcedure,
			connect.WithSchema(workspaceServiceGetWorkspacesMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getWorkspaceById: connect.NewClient[v1.GetWorkspaceByIdRequest, v1.Workspace](
			httpClient,
			baseURL+WorkspaceServiceGetWorkspaceByIdProcedure,
			connect.WithSchema(workspaceServiceGetWorkspaceByIdMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getWorkspaceByNamespace: connect.NewClient[v1.GetWorkspaceByNamespaceRequest, v1.Workspace](
			httpClient,
			baseURL+WorkspaceServiceGetWorkspaceByNamespaceProcedure,
			connect.WithSchema(workspaceServiceGetWorkspaceByNamespaceMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		updateWorkspaceName: connect.NewClient[v1.UpdateWorkspaceNameRequest, v1.Workspace](
			httpClient,
			baseURL+WorkspaceServiceUpdateWorkspaceNameProcedure,
			connect.WithSchema(workspaceServiceUpdateWorkspaceNameMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateWorkspaceNamespace: connect.NewClient[v1.UpdateWorkspaceNamespaceRequest, v1.Workspace](
			httpClient,
			baseURL+WorkspaceServiceUpdateWorkspaceNamespaceProcedure,
			connect.WithSchema(workspaceServiceUpdateWorkspaceNamespaceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updatePrimaryCurrency: connect.NewClient[v1.UpdatePrimaryCurrencyRequest, v1.Workspace](
			httpClient,
			baseURL+WorkspaceServiceUpdatePrimaryCurrencyProcedure,
			connect.WithSchema(workspaceServiceUpdatePrimaryCurrencyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateLocale: connect.NewClient[v1.UpdateLocaleRequest, v1.Workspace](
			httpClient,
			baseURL+WorkspaceServiceUpdateLocaleProcedure,
			connect.WithSchema(workspaceServiceUpdateLocaleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteWorkspace: connect.NewClient[emptypb.Empty, v1.Workspace](
			httpClient,
			baseURL+WorkspaceServiceDeleteWorkspaceProcedure,
			connect.WithSchema(workspaceServiceDeleteWorkspaceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// workspaceServiceClient implements WorkspaceServiceClient.
type workspaceServiceClient struct {
	createWorkspace          *connect.Client[v1.CreateWorkspaceRequest, v1.Workspace]
	getWorkspaces            *connect.Client[emptypb.Empty, v1.Workspaces]
	getWorkspaceById         *connect.Client[v1.GetWorkspaceByIdRequest, v1.Workspace]
	getWorkspaceByNamespace  *connect.Client[v1.GetWorkspaceByNamespaceRequest, v1.Workspace]
	updateWorkspaceName      *connect.Client[v1.UpdateWorkspaceNameRequest, v1.Workspace]
	updateWorkspaceNamespace *connect.Client[v1.UpdateWorkspaceNamespaceRequest, v1.Workspace]
	updatePrimaryCurrency    *connect.Client[v1.UpdatePrimaryCurrencyRequest, v1.Workspace]
	updateLocale             *connect.Client[v1.UpdateLocaleRequest, v1.Workspace]
	deleteWorkspace          *connect.Client[emptypb.Empty, v1.Workspace]
}

// CreateWorkspace calls fijoy.v1.WorkspaceService.CreateWorkspace.
func (c *workspaceServiceClient) CreateWorkspace(ctx context.Context, req *connect.Request[v1.CreateWorkspaceRequest]) (*connect.Response[v1.Workspace], error) {
	return c.createWorkspace.CallUnary(ctx, req)
}

// GetWorkspaces calls fijoy.v1.WorkspaceService.GetWorkspaces.
func (c *workspaceServiceClient) GetWorkspaces(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.Workspaces], error) {
	return c.getWorkspaces.CallUnary(ctx, req)
}

// GetWorkspaceById calls fijoy.v1.WorkspaceService.GetWorkspaceById.
func (c *workspaceServiceClient) GetWorkspaceById(ctx context.Context, req *connect.Request[v1.GetWorkspaceByIdRequest]) (*connect.Response[v1.Workspace], error) {
	return c.getWorkspaceById.CallUnary(ctx, req)
}

// GetWorkspaceByNamespace calls fijoy.v1.WorkspaceService.GetWorkspaceByNamespace.
func (c *workspaceServiceClient) GetWorkspaceByNamespace(ctx context.Context, req *connect.Request[v1.GetWorkspaceByNamespaceRequest]) (*connect.Response[v1.Workspace], error) {
	return c.getWorkspaceByNamespace.CallUnary(ctx, req)
}

// UpdateWorkspaceName calls fijoy.v1.WorkspaceService.UpdateWorkspaceName.
func (c *workspaceServiceClient) UpdateWorkspaceName(ctx context.Context, req *connect.Request[v1.UpdateWorkspaceNameRequest]) (*connect.Response[v1.Workspace], error) {
	return c.updateWorkspaceName.CallUnary(ctx, req)
}

// UpdateWorkspaceNamespace calls fijoy.v1.WorkspaceService.UpdateWorkspaceNamespace.
func (c *workspaceServiceClient) UpdateWorkspaceNamespace(ctx context.Context, req *connect.Request[v1.UpdateWorkspaceNamespaceRequest]) (*connect.Response[v1.Workspace], error) {
	return c.updateWorkspaceNamespace.CallUnary(ctx, req)
}

// UpdatePrimaryCurrency calls fijoy.v1.WorkspaceService.UpdatePrimaryCurrency.
func (c *workspaceServiceClient) UpdatePrimaryCurrency(ctx context.Context, req *connect.Request[v1.UpdatePrimaryCurrencyRequest]) (*connect.Response[v1.Workspace], error) {
	return c.updatePrimaryCurrency.CallUnary(ctx, req)
}

// UpdateLocale calls fijoy.v1.WorkspaceService.UpdateLocale.
func (c *workspaceServiceClient) UpdateLocale(ctx context.Context, req *connect.Request[v1.UpdateLocaleRequest]) (*connect.Response[v1.Workspace], error) {
	return c.updateLocale.CallUnary(ctx, req)
}

// DeleteWorkspace calls fijoy.v1.WorkspaceService.DeleteWorkspace.
func (c *workspaceServiceClient) DeleteWorkspace(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.Workspace], error) {
	return c.deleteWorkspace.CallUnary(ctx, req)
}

// WorkspaceServiceHandler is an implementation of the fijoy.v1.WorkspaceService service.
type WorkspaceServiceHandler interface {
	CreateWorkspace(context.Context, *connect.Request[v1.CreateWorkspaceRequest]) (*connect.Response[v1.Workspace], error)
	GetWorkspaces(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.Workspaces], error)
	GetWorkspaceById(context.Context, *connect.Request[v1.GetWorkspaceByIdRequest]) (*connect.Response[v1.Workspace], error)
	GetWorkspaceByNamespace(context.Context, *connect.Request[v1.GetWorkspaceByNamespaceRequest]) (*connect.Response[v1.Workspace], error)
	UpdateWorkspaceName(context.Context, *connect.Request[v1.UpdateWorkspaceNameRequest]) (*connect.Response[v1.Workspace], error)
	UpdateWorkspaceNamespace(context.Context, *connect.Request[v1.UpdateWorkspaceNamespaceRequest]) (*connect.Response[v1.Workspace], error)
	UpdatePrimaryCurrency(context.Context, *connect.Request[v1.UpdatePrimaryCurrencyRequest]) (*connect.Response[v1.Workspace], error)
	UpdateLocale(context.Context, *connect.Request[v1.UpdateLocaleRequest]) (*connect.Response[v1.Workspace], error)
	DeleteWorkspace(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.Workspace], error)
}

// NewWorkspaceServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkspaceServiceHandler(svc WorkspaceServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	workspaceServiceCreateWorkspaceHandler := connect.NewUnaryHandler(
		WorkspaceServiceCreateWorkspaceProcedure,
		svc.CreateWorkspace,
		connect.WithSchema(workspaceServiceCreateWorkspaceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceGetWorkspacesHandler := connect.NewUnaryHandler(
		WorkspaceServiceGetWorkspacesProcedure,
		svc.GetWorkspaces,
		connect.WithSchema(workspaceServiceGetWorkspacesMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceGetWorkspaceByIdHandler := connect.NewUnaryHandler(
		WorkspaceServiceGetWorkspaceByIdProcedure,
		svc.GetWorkspaceById,
		connect.WithSchema(workspaceServiceGetWorkspaceByIdMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceGetWorkspaceByNamespaceHandler := connect.NewUnaryHandler(
		WorkspaceServiceGetWorkspaceByNamespaceProcedure,
		svc.GetWorkspaceByNamespace,
		connect.WithSchema(workspaceServiceGetWorkspaceByNamespaceMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceUpdateWorkspaceNameHandler := connect.NewUnaryHandler(
		WorkspaceServiceUpdateWorkspaceNameProcedure,
		svc.UpdateWorkspaceName,
		connect.WithSchema(workspaceServiceUpdateWorkspaceNameMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceUpdateWorkspaceNamespaceHandler := connect.NewUnaryHandler(
		WorkspaceServiceUpdateWorkspaceNamespaceProcedure,
		svc.UpdateWorkspaceNamespace,
		connect.WithSchema(workspaceServiceUpdateWorkspaceNamespaceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceUpdatePrimaryCurrencyHandler := connect.NewUnaryHandler(
		WorkspaceServiceUpdatePrimaryCurrencyProcedure,
		svc.UpdatePrimaryCurrency,
		connect.WithSchema(workspaceServiceUpdatePrimaryCurrencyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceUpdateLocaleHandler := connect.NewUnaryHandler(
		WorkspaceServiceUpdateLocaleProcedure,
		svc.UpdateLocale,
		connect.WithSchema(workspaceServiceUpdateLocaleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceDeleteWorkspaceHandler := connect.NewUnaryHandler(
		WorkspaceServiceDeleteWorkspaceProcedure,
		svc.DeleteWorkspace,
		connect.WithSchema(workspaceServiceDeleteWorkspaceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/fijoy.v1.WorkspaceService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WorkspaceServiceCreateWorkspaceProcedure:
			workspaceServiceCreateWorkspaceHandler.ServeHTTP(w, r)
		case WorkspaceServiceGetWorkspacesProcedure:
			workspaceServiceGetWorkspacesHandler.ServeHTTP(w, r)
		case WorkspaceServiceGetWorkspaceByIdProcedure:
			workspaceServiceGetWorkspaceByIdHandler.ServeHTTP(w, r)
		case WorkspaceServiceGetWorkspaceByNamespaceProcedure:
			workspaceServiceGetWorkspaceByNamespaceHandler.ServeHTTP(w, r)
		case WorkspaceServiceUpdateWorkspaceNameProcedure:
			workspaceServiceUpdateWorkspaceNameHandler.ServeHTTP(w, r)
		case WorkspaceServiceUpdateWorkspaceNamespaceProcedure:
			workspaceServiceUpdateWorkspaceNamespaceHandler.ServeHTTP(w, r)
		case WorkspaceServiceUpdatePrimaryCurrencyProcedure:
			workspaceServiceUpdatePrimaryCurrencyHandler.ServeHTTP(w, r)
		case WorkspaceServiceUpdateLocaleProcedure:
			workspaceServiceUpdateLocaleHandler.ServeHTTP(w, r)
		case WorkspaceServiceDeleteWorkspaceProcedure:
			workspaceServiceDeleteWorkspaceHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWorkspaceServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkspaceServiceHandler struct{}

func (UnimplementedWorkspaceServiceHandler) CreateWorkspace(context.Context, *connect.Request[v1.CreateWorkspaceRequest]) (*connect.Response[v1.Workspace], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("fijoy.v1.WorkspaceService.CreateWorkspace is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) GetWorkspaces(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.Workspaces], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("fijoy.v1.WorkspaceService.GetWorkspaces is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) GetWorkspaceById(context.Context, *connect.Request[v1.GetWorkspaceByIdRequest]) (*connect.Response[v1.Workspace], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("fijoy.v1.WorkspaceService.GetWorkspaceById is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) GetWorkspaceByNamespace(context.Context, *connect.Request[v1.GetWorkspaceByNamespaceRequest]) (*connect.Response[v1.Workspace], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("fijoy.v1.WorkspaceService.GetWorkspaceByNamespace is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) UpdateWorkspaceName(context.Context, *connect.Request[v1.UpdateWorkspaceNameRequest]) (*connect.Response[v1.Workspace], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("fijoy.v1.WorkspaceService.UpdateWorkspaceName is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) UpdateWorkspaceNamespace(context.Context, *connect.Request[v1.UpdateWorkspaceNamespaceRequest]) (*connect.Response[v1.Workspace], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("fijoy.v1.WorkspaceService.UpdateWorkspaceNamespace is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) UpdatePrimaryCurrency(context.Context, *connect.Request[v1.UpdatePrimaryCurrencyRequest]) (*connect.Response[v1.Workspace], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("fijoy.v1.WorkspaceService.UpdatePrimaryCurrency is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) UpdateLocale(context.Context, *connect.Request[v1.UpdateLocaleRequest]) (*connect.Response[v1.Workspace], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("fijoy.v1.WorkspaceService.UpdateLocale is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) DeleteWorkspace(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.Workspace], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("fijoy.v1.WorkspaceService.DeleteWorkspace is not implemented"))
}
