// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: probe/v1/probe.proto

/*
Package probe is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package probe

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
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_ProbeService_GetProbeData_0(ctx context.Context, marshaler runtime.Marshaler, client ProbeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProbeDataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetProbeData(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ProbeService_GetProbeData_0(ctx context.Context, marshaler runtime.Marshaler, server ProbeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProbeDataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetProbeData(ctx, &protoReq)
	return msg, metadata, err

}

func request_ProbeService_GetProbeDataBulk_0(ctx context.Context, marshaler runtime.Marshaler, client ProbeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBulkProbeDataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetProbeDataBulk(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ProbeService_GetProbeDataBulk_0(ctx context.Context, marshaler runtime.Marshaler, server ProbeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBulkProbeDataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetProbeDataBulk(ctx, &protoReq)
	return msg, metadata, err

}

func request_ProbeService_GetStreams_0(ctx context.Context, marshaler runtime.Marshaler, client ProbeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProbeDataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetStreams(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ProbeService_GetStreams_0(ctx context.Context, marshaler runtime.Marshaler, server ProbeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProbeDataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetStreams(ctx, &protoReq)
	return msg, metadata, err

}

func request_ProbeService_GetStream_0(ctx context.Context, marshaler runtime.Marshaler, client ProbeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetStreamRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["index"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "index")
	}

	protoReq.Index, err = runtime.Int32(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "index", err)
	}

	msg, err := client.GetStream(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ProbeService_GetStream_0(ctx context.Context, marshaler runtime.Marshaler, server ProbeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetStreamRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["index"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "index")
	}

	protoReq.Index, err = runtime.Int32(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "index", err)
	}

	msg, err := server.GetStream(ctx, &protoReq)
	return msg, metadata, err

}

func request_ProbeService_GetChapters_0(ctx context.Context, marshaler runtime.Marshaler, client ProbeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProbeDataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetChapters(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ProbeService_GetChapters_0(ctx context.Context, marshaler runtime.Marshaler, server ProbeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProbeDataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetChapters(ctx, &protoReq)
	return msg, metadata, err

}

func request_ProbeService_GetFormat_0(ctx context.Context, marshaler runtime.Marshaler, client ProbeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProbeDataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetFormat(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ProbeService_GetFormat_0(ctx context.Context, marshaler runtime.Marshaler, server ProbeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProbeDataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetFormat(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterProbeServiceHandlerServer registers the http handlers for service ProbeService to "mux".
// UnaryRPC     :call ProbeServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterProbeServiceHandlerFromEndpoint instead.
func RegisterProbeServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ProbeServiceServer) error {

	mux.Handle("POST", pattern_ProbeService_GetProbeData_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/probe.v1.ProbeService/GetProbeData", runtime.WithHTTPPathPattern("/probe/v1/data"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ProbeService_GetProbeData_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetProbeData_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetProbeData_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ProbeService_GetProbeDataBulk_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/probe.v1.ProbeService/GetProbeDataBulk", runtime.WithHTTPPathPattern("/probe/v1/data/bulk"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ProbeService_GetProbeDataBulk_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetProbeDataBulk_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetProbeDataBulk_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ProbeService_GetStreams_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/probe.v1.ProbeService/GetStreams", runtime.WithHTTPPathPattern("/probe/v1/streams"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ProbeService_GetStreams_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetStreams_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetStreams_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ProbeService_GetStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/probe.v1.ProbeService/GetStream", runtime.WithHTTPPathPattern("/probe/v1/streams/{index}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ProbeService_GetStream_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetStream_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetStream_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ProbeService_GetChapters_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/probe.v1.ProbeService/GetChapters", runtime.WithHTTPPathPattern("/probe/v1/chapters"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ProbeService_GetChapters_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetChapters_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetChapters_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ProbeService_GetFormat_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/probe.v1.ProbeService/GetFormat", runtime.WithHTTPPathPattern("/probe/v1/format"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ProbeService_GetFormat_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetFormat_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetFormat_0{resp}, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterProbeServiceHandlerFromEndpoint is same as RegisterProbeServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterProbeServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterProbeServiceHandler(ctx, mux, conn)
}

// RegisterProbeServiceHandler registers the http handlers for service ProbeService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterProbeServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterProbeServiceHandlerClient(ctx, mux, NewProbeServiceClient(conn))
}

// RegisterProbeServiceHandlerClient registers the http handlers for service ProbeService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ProbeServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ProbeServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ProbeServiceClient" to call the correct interceptors.
func RegisterProbeServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ProbeServiceClient) error {

	mux.Handle("POST", pattern_ProbeService_GetProbeData_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/probe.v1.ProbeService/GetProbeData", runtime.WithHTTPPathPattern("/probe/v1/data"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProbeService_GetProbeData_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetProbeData_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetProbeData_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ProbeService_GetProbeDataBulk_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/probe.v1.ProbeService/GetProbeDataBulk", runtime.WithHTTPPathPattern("/probe/v1/data/bulk"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProbeService_GetProbeDataBulk_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetProbeDataBulk_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetProbeDataBulk_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ProbeService_GetStreams_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/probe.v1.ProbeService/GetStreams", runtime.WithHTTPPathPattern("/probe/v1/streams"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProbeService_GetStreams_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetStreams_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetStreams_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ProbeService_GetStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/probe.v1.ProbeService/GetStream", runtime.WithHTTPPathPattern("/probe/v1/streams/{index}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProbeService_GetStream_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetStream_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetStream_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ProbeService_GetChapters_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/probe.v1.ProbeService/GetChapters", runtime.WithHTTPPathPattern("/probe/v1/chapters"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProbeService_GetChapters_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetChapters_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetChapters_0{resp}, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ProbeService_GetFormat_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/probe.v1.ProbeService/GetFormat", runtime.WithHTTPPathPattern("/probe/v1/format"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProbeService_GetFormat_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProbeService_GetFormat_0(annotatedContext, mux, outboundMarshaler, w, req, response_ProbeService_GetFormat_0{resp}, mux.GetForwardResponseOptions()...)

	})

	return nil
}

type response_ProbeService_GetProbeData_0 struct {
	proto.Message
}

func (m response_ProbeService_GetProbeData_0) XXX_ResponseBody() interface{} {
	response := m.Message.(*GetProbeDataResponse)
	return response.Data
}

type response_ProbeService_GetProbeDataBulk_0 struct {
	proto.Message
}

func (m response_ProbeService_GetProbeDataBulk_0) XXX_ResponseBody() interface{} {
	response := m.Message.(*GetBulkProbeDataResponse)
	return response.Data
}

type response_ProbeService_GetStreams_0 struct {
	proto.Message
}

func (m response_ProbeService_GetStreams_0) XXX_ResponseBody() interface{} {
	response := m.Message.(*GetStreamsResponse)
	return response.Streams
}

type response_ProbeService_GetStream_0 struct {
	proto.Message
}

func (m response_ProbeService_GetStream_0) XXX_ResponseBody() interface{} {
	response := m.Message.(*GetStreamResponse)
	return response.Stream
}

type response_ProbeService_GetChapters_0 struct {
	proto.Message
}

func (m response_ProbeService_GetChapters_0) XXX_ResponseBody() interface{} {
	response := m.Message.(*GetChaptersResponse)
	return response.Chapters
}

type response_ProbeService_GetFormat_0 struct {
	proto.Message
}

func (m response_ProbeService_GetFormat_0) XXX_ResponseBody() interface{} {
	response := m.Message.(*GetFormatResponse)
	return response.Format
}

var (
	pattern_ProbeService_GetProbeData_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"probe", "v1", "data"}, ""))

	pattern_ProbeService_GetProbeDataBulk_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"probe", "v1", "data", "bulk"}, ""))

	pattern_ProbeService_GetStreams_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"probe", "v1", "streams"}, ""))

	pattern_ProbeService_GetStream_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"probe", "v1", "streams", "index"}, ""))

	pattern_ProbeService_GetChapters_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"probe", "v1", "chapters"}, ""))

	pattern_ProbeService_GetFormat_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"probe", "v1", "format"}, ""))
)

var (
	forward_ProbeService_GetProbeData_0 = runtime.ForwardResponseMessage

	forward_ProbeService_GetProbeDataBulk_0 = runtime.ForwardResponseMessage

	forward_ProbeService_GetStreams_0 = runtime.ForwardResponseMessage

	forward_ProbeService_GetStream_0 = runtime.ForwardResponseMessage

	forward_ProbeService_GetChapters_0 = runtime.ForwardResponseMessage

	forward_ProbeService_GetFormat_0 = runtime.ForwardResponseMessage
)
