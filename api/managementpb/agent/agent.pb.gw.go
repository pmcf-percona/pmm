// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: managementpb/agent/agent.proto

/*
Package agentv1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package agentv1beta1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_Agent_ListAgents_0(ctx context.Context, marshaler runtime.Marshaler, client AgentClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListAgentRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListAgents(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Agent_ListAgents_0(ctx context.Context, marshaler runtime.Marshaler, server AgentServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListAgentRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ListAgents(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterAgentHandlerServer registers the http handlers for service Agent to "mux".
// UnaryRPC     :call AgentServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterAgentHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterAgentHandlerServer(ctx context.Context, mux *runtime.ServeMux, server AgentServer) error {
	mux.Handle("POST", pattern_Agent_ListAgents_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/agent.v1beta1.Agent/ListAgents", runtime.WithHTTPPathPattern("/v1/management/Agent/List"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Agent_ListAgents_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Agent_ListAgents_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterAgentHandlerFromEndpoint is same as RegisterAgentHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAgentHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAgentHandler(ctx, mux, conn)
}

// RegisterAgentHandler registers the http handlers for service Agent to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAgentHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAgentHandlerClient(ctx, mux, NewAgentClient(conn))
}

// RegisterAgentHandlerClient registers the http handlers for service Agent
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "AgentClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "AgentClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "AgentClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterAgentHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AgentClient) error {
	mux.Handle("POST", pattern_Agent_ListAgents_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/agent.v1beta1.Agent/ListAgents", runtime.WithHTTPPathPattern("/v1/management/Agent/List"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Agent_ListAgents_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Agent_ListAgents_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

var pattern_Agent_ListAgents_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "management", "Agent", "List"}, ""))

var forward_Agent_ListAgents_0 = runtime.ForwardResponseMessage
