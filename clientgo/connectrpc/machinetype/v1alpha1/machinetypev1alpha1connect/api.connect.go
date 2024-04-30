// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: machinetype/v1alpha1/api.proto

package machinetypev1alpha1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"
	v1alpha1 "github.com/ironcore-dev/lifecycle-manager/api/proto/machinetype/v1alpha1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// MachineTypeServiceName is the fully-qualified name of the MachineTypeService service.
	MachineTypeServiceName = "machinetype.v1alpha1.MachineTypeService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MachineTypeServiceListMachineTypesProcedure is the fully-qualified name of the
	// MachineTypeService's ListMachineTypes RPC.
	MachineTypeServiceListMachineTypesProcedure = "/machinetype.v1alpha1.MachineTypeService/ListMachineTypes"
	// MachineTypeServiceScanProcedure is the fully-qualified name of the MachineTypeService's Scan RPC.
	MachineTypeServiceScanProcedure = "/machinetype.v1alpha1.MachineTypeService/Scan"
	// MachineTypeServiceUpdateMachineTypeStatusProcedure is the fully-qualified name of the
	// MachineTypeService's UpdateMachineTypeStatus RPC.
	MachineTypeServiceUpdateMachineTypeStatusProcedure = "/machinetype.v1alpha1.MachineTypeService/UpdateMachineTypeStatus"
	// MachineTypeServiceAddMachineGroupProcedure is the fully-qualified name of the
	// MachineTypeService's AddMachineGroup RPC.
	MachineTypeServiceAddMachineGroupProcedure = "/machinetype.v1alpha1.MachineTypeService/AddMachineGroup"
	// MachineTypeServiceRemoveMachineGroupProcedure is the fully-qualified name of the
	// MachineTypeService's RemoveMachineGroup RPC.
	MachineTypeServiceRemoveMachineGroupProcedure = "/machinetype.v1alpha1.MachineTypeService/RemoveMachineGroup"
	// MachineTypeServiceGetJobProcedure is the fully-qualified name of the MachineTypeService's GetJob
	// RPC.
	MachineTypeServiceGetJobProcedure = "/machinetype.v1alpha1.MachineTypeService/GetJob"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	machineTypeServiceServiceDescriptor                       = v1alpha1.File_machinetype_v1alpha1_api_proto.Services().ByName("MachineTypeService")
	machineTypeServiceListMachineTypesMethodDescriptor        = machineTypeServiceServiceDescriptor.Methods().ByName("ListMachineTypes")
	machineTypeServiceScanMethodDescriptor                    = machineTypeServiceServiceDescriptor.Methods().ByName("Scan")
	machineTypeServiceUpdateMachineTypeStatusMethodDescriptor = machineTypeServiceServiceDescriptor.Methods().ByName("UpdateMachineTypeStatus")
	machineTypeServiceAddMachineGroupMethodDescriptor         = machineTypeServiceServiceDescriptor.Methods().ByName("AddMachineGroup")
	machineTypeServiceRemoveMachineGroupMethodDescriptor      = machineTypeServiceServiceDescriptor.Methods().ByName("RemoveMachineGroup")
	machineTypeServiceGetJobMethodDescriptor                  = machineTypeServiceServiceDescriptor.Methods().ByName("GetJob")
)

// MachineTypeServiceClient is a client for the machinetype.v1alpha1.MachineTypeService service.
type MachineTypeServiceClient interface {
	ListMachineTypes(context.Context, *connect.Request[v1alpha1.ListMachineTypesRequest]) (*connect.Response[v1alpha1.ListMachineTypesResponse], error)
	Scan(context.Context, *connect.Request[v1alpha1.ScanRequest]) (*connect.Response[v1alpha1.ScanResponse], error)
	UpdateMachineTypeStatus(context.Context, *connect.Request[v1alpha1.UpdateMachineTypeStatusRequest]) (*connect.Response[v1alpha1.UpdateMachineTypeStatusResponse], error)
	AddMachineGroup(context.Context, *connect.Request[v1alpha1.AddMachineGroupRequest]) (*connect.Response[v1alpha1.AddMachineGroupResponse], error)
	RemoveMachineGroup(context.Context, *connect.Request[v1alpha1.RemoveMachineGroupRequest]) (*connect.Response[v1alpha1.RemoveMachineGroupResponse], error)
	GetJob(context.Context, *connect.Request[v1alpha1.GetJobRequest]) (*connect.Response[v1alpha1.GetJobResponse], error)
}

// NewMachineTypeServiceClient constructs a client for the machinetype.v1alpha1.MachineTypeService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMachineTypeServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MachineTypeServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &machineTypeServiceClient{
		listMachineTypes: connect.NewClient[v1alpha1.ListMachineTypesRequest, v1alpha1.ListMachineTypesResponse](
			httpClient,
			baseURL+MachineTypeServiceListMachineTypesProcedure,
			connect.WithSchema(machineTypeServiceListMachineTypesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		scan: connect.NewClient[v1alpha1.ScanRequest, v1alpha1.ScanResponse](
			httpClient,
			baseURL+MachineTypeServiceScanProcedure,
			connect.WithSchema(machineTypeServiceScanMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateMachineTypeStatus: connect.NewClient[v1alpha1.UpdateMachineTypeStatusRequest, v1alpha1.UpdateMachineTypeStatusResponse](
			httpClient,
			baseURL+MachineTypeServiceUpdateMachineTypeStatusProcedure,
			connect.WithSchema(machineTypeServiceUpdateMachineTypeStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		addMachineGroup: connect.NewClient[v1alpha1.AddMachineGroupRequest, v1alpha1.AddMachineGroupResponse](
			httpClient,
			baseURL+MachineTypeServiceAddMachineGroupProcedure,
			connect.WithSchema(machineTypeServiceAddMachineGroupMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		removeMachineGroup: connect.NewClient[v1alpha1.RemoveMachineGroupRequest, v1alpha1.RemoveMachineGroupResponse](
			httpClient,
			baseURL+MachineTypeServiceRemoveMachineGroupProcedure,
			connect.WithSchema(machineTypeServiceRemoveMachineGroupMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getJob: connect.NewClient[v1alpha1.GetJobRequest, v1alpha1.GetJobResponse](
			httpClient,
			baseURL+MachineTypeServiceGetJobProcedure,
			connect.WithSchema(machineTypeServiceGetJobMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// machineTypeServiceClient implements MachineTypeServiceClient.
type machineTypeServiceClient struct {
	listMachineTypes        *connect.Client[v1alpha1.ListMachineTypesRequest, v1alpha1.ListMachineTypesResponse]
	scan                    *connect.Client[v1alpha1.ScanRequest, v1alpha1.ScanResponse]
	updateMachineTypeStatus *connect.Client[v1alpha1.UpdateMachineTypeStatusRequest, v1alpha1.UpdateMachineTypeStatusResponse]
	addMachineGroup         *connect.Client[v1alpha1.AddMachineGroupRequest, v1alpha1.AddMachineGroupResponse]
	removeMachineGroup      *connect.Client[v1alpha1.RemoveMachineGroupRequest, v1alpha1.RemoveMachineGroupResponse]
	getJob                  *connect.Client[v1alpha1.GetJobRequest, v1alpha1.GetJobResponse]
}

// ListMachineTypes calls machinetype.v1alpha1.MachineTypeService.ListMachineTypes.
func (c *machineTypeServiceClient) ListMachineTypes(ctx context.Context, req *connect.Request[v1alpha1.ListMachineTypesRequest]) (*connect.Response[v1alpha1.ListMachineTypesResponse], error) {
	return c.listMachineTypes.CallUnary(ctx, req)
}

// Scan calls machinetype.v1alpha1.MachineTypeService.Scan.
func (c *machineTypeServiceClient) Scan(ctx context.Context, req *connect.Request[v1alpha1.ScanRequest]) (*connect.Response[v1alpha1.ScanResponse], error) {
	return c.scan.CallUnary(ctx, req)
}

// UpdateMachineTypeStatus calls machinetype.v1alpha1.MachineTypeService.UpdateMachineTypeStatus.
func (c *machineTypeServiceClient) UpdateMachineTypeStatus(ctx context.Context, req *connect.Request[v1alpha1.UpdateMachineTypeStatusRequest]) (*connect.Response[v1alpha1.UpdateMachineTypeStatusResponse], error) {
	return c.updateMachineTypeStatus.CallUnary(ctx, req)
}

// AddMachineGroup calls machinetype.v1alpha1.MachineTypeService.AddMachineGroup.
func (c *machineTypeServiceClient) AddMachineGroup(ctx context.Context, req *connect.Request[v1alpha1.AddMachineGroupRequest]) (*connect.Response[v1alpha1.AddMachineGroupResponse], error) {
	return c.addMachineGroup.CallUnary(ctx, req)
}

// RemoveMachineGroup calls machinetype.v1alpha1.MachineTypeService.RemoveMachineGroup.
func (c *machineTypeServiceClient) RemoveMachineGroup(ctx context.Context, req *connect.Request[v1alpha1.RemoveMachineGroupRequest]) (*connect.Response[v1alpha1.RemoveMachineGroupResponse], error) {
	return c.removeMachineGroup.CallUnary(ctx, req)
}

// GetJob calls machinetype.v1alpha1.MachineTypeService.GetJob.
func (c *machineTypeServiceClient) GetJob(ctx context.Context, req *connect.Request[v1alpha1.GetJobRequest]) (*connect.Response[v1alpha1.GetJobResponse], error) {
	return c.getJob.CallUnary(ctx, req)
}

// MachineTypeServiceHandler is an implementation of the machinetype.v1alpha1.MachineTypeService
// service.
type MachineTypeServiceHandler interface {
	ListMachineTypes(context.Context, *connect.Request[v1alpha1.ListMachineTypesRequest]) (*connect.Response[v1alpha1.ListMachineTypesResponse], error)
	Scan(context.Context, *connect.Request[v1alpha1.ScanRequest]) (*connect.Response[v1alpha1.ScanResponse], error)
	UpdateMachineTypeStatus(context.Context, *connect.Request[v1alpha1.UpdateMachineTypeStatusRequest]) (*connect.Response[v1alpha1.UpdateMachineTypeStatusResponse], error)
	AddMachineGroup(context.Context, *connect.Request[v1alpha1.AddMachineGroupRequest]) (*connect.Response[v1alpha1.AddMachineGroupResponse], error)
	RemoveMachineGroup(context.Context, *connect.Request[v1alpha1.RemoveMachineGroupRequest]) (*connect.Response[v1alpha1.RemoveMachineGroupResponse], error)
	GetJob(context.Context, *connect.Request[v1alpha1.GetJobRequest]) (*connect.Response[v1alpha1.GetJobResponse], error)
}

// NewMachineTypeServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMachineTypeServiceHandler(svc MachineTypeServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	machineTypeServiceListMachineTypesHandler := connect.NewUnaryHandler(
		MachineTypeServiceListMachineTypesProcedure,
		svc.ListMachineTypes,
		connect.WithSchema(machineTypeServiceListMachineTypesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	machineTypeServiceScanHandler := connect.NewUnaryHandler(
		MachineTypeServiceScanProcedure,
		svc.Scan,
		connect.WithSchema(machineTypeServiceScanMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	machineTypeServiceUpdateMachineTypeStatusHandler := connect.NewUnaryHandler(
		MachineTypeServiceUpdateMachineTypeStatusProcedure,
		svc.UpdateMachineTypeStatus,
		connect.WithSchema(machineTypeServiceUpdateMachineTypeStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	machineTypeServiceAddMachineGroupHandler := connect.NewUnaryHandler(
		MachineTypeServiceAddMachineGroupProcedure,
		svc.AddMachineGroup,
		connect.WithSchema(machineTypeServiceAddMachineGroupMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	machineTypeServiceRemoveMachineGroupHandler := connect.NewUnaryHandler(
		MachineTypeServiceRemoveMachineGroupProcedure,
		svc.RemoveMachineGroup,
		connect.WithSchema(machineTypeServiceRemoveMachineGroupMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	machineTypeServiceGetJobHandler := connect.NewUnaryHandler(
		MachineTypeServiceGetJobProcedure,
		svc.GetJob,
		connect.WithSchema(machineTypeServiceGetJobMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/machinetype.v1alpha1.MachineTypeService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MachineTypeServiceListMachineTypesProcedure:
			machineTypeServiceListMachineTypesHandler.ServeHTTP(w, r)
		case MachineTypeServiceScanProcedure:
			machineTypeServiceScanHandler.ServeHTTP(w, r)
		case MachineTypeServiceUpdateMachineTypeStatusProcedure:
			machineTypeServiceUpdateMachineTypeStatusHandler.ServeHTTP(w, r)
		case MachineTypeServiceAddMachineGroupProcedure:
			machineTypeServiceAddMachineGroupHandler.ServeHTTP(w, r)
		case MachineTypeServiceRemoveMachineGroupProcedure:
			machineTypeServiceRemoveMachineGroupHandler.ServeHTTP(w, r)
		case MachineTypeServiceGetJobProcedure:
			machineTypeServiceGetJobHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMachineTypeServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMachineTypeServiceHandler struct{}

func (UnimplementedMachineTypeServiceHandler) ListMachineTypes(context.Context, *connect.Request[v1alpha1.ListMachineTypesRequest]) (*connect.Response[v1alpha1.ListMachineTypesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("machinetype.v1alpha1.MachineTypeService.ListMachineTypes is not implemented"))
}

func (UnimplementedMachineTypeServiceHandler) Scan(context.Context, *connect.Request[v1alpha1.ScanRequest]) (*connect.Response[v1alpha1.ScanResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("machinetype.v1alpha1.MachineTypeService.Scan is not implemented"))
}

func (UnimplementedMachineTypeServiceHandler) UpdateMachineTypeStatus(context.Context, *connect.Request[v1alpha1.UpdateMachineTypeStatusRequest]) (*connect.Response[v1alpha1.UpdateMachineTypeStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("machinetype.v1alpha1.MachineTypeService.UpdateMachineTypeStatus is not implemented"))
}

func (UnimplementedMachineTypeServiceHandler) AddMachineGroup(context.Context, *connect.Request[v1alpha1.AddMachineGroupRequest]) (*connect.Response[v1alpha1.AddMachineGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("machinetype.v1alpha1.MachineTypeService.AddMachineGroup is not implemented"))
}

func (UnimplementedMachineTypeServiceHandler) RemoveMachineGroup(context.Context, *connect.Request[v1alpha1.RemoveMachineGroupRequest]) (*connect.Response[v1alpha1.RemoveMachineGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("machinetype.v1alpha1.MachineTypeService.RemoveMachineGroup is not implemented"))
}

func (UnimplementedMachineTypeServiceHandler) GetJob(context.Context, *connect.Request[v1alpha1.GetJobRequest]) (*connect.Response[v1alpha1.GetJobResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("machinetype.v1alpha1.MachineTypeService.GetJob is not implemented"))
}