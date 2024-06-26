// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"github.com/go-logr/logr"
)

type LogInterceptor struct {
	logger *slog.Logger
}

func NewLoggerInterceptor(log *slog.Logger) connect.Interceptor {
	return &LogInterceptor{logger: log}
}

func (l *LogInterceptor) WrapUnary(unaryFunc connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		log := l.logger.With("peer", req.Peer(), "endpoint", req.Spec().Procedure)
		reqCtx := logr.NewContextWithSlogLogger(ctx, log)
		response, err := unaryFunc(reqCtx, req)
		if err != nil {
			log.Error("request failed", "error", err.Error())
			return response, err
		}
		log.Debug("request finished", "response", response.Any())
		return response, err
	}
}

func (l *LogInterceptor) WrapStreamingClient(clientFunc connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		return &streamingClientInterceptor{
			StreamingClientConn: clientFunc(ctx, spec),
			logger:              l.logger,
		}
	}
}

func (l *LogInterceptor) WrapStreamingHandler(handlerFunc connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		return handlerFunc(ctx, &streamingHandlerInterceptor{
			StreamingHandlerConn: conn,
			logger:               l.logger,
		})
	}
}

type streamingClientInterceptor struct {
	connect.StreamingClientConn
	logger *slog.Logger
}

func (s *streamingClientInterceptor) Spec() connect.Spec {
	// TODO implement me
	panic("implement me")
}

func (s *streamingClientInterceptor) Peer() connect.Peer {
	// TODO implement me
	panic("implement me")
}

// func (s *streamingClientInterceptor) Send(a any) error {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingClientInterceptor) RequestHeader() http.Header {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingClientInterceptor) CloseRequest() error {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingClientInterceptor) Receive(a any) error {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingClientInterceptor) ResponseHeader() http.Header {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingClientInterceptor) ResponseTrailer() http.Header {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingClientInterceptor) CloseResponse() error {
// 	// TODO implement me
// 	panic("implement me")
// }

type streamingHandlerInterceptor struct {
	connect.StreamingHandlerConn
	logger *slog.Logger
}

// func (s *streamingHandlerInterceptor) Spec() connect.Spec {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingHandlerInterceptor) Peer() connect.Peer {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingHandlerInterceptor) Receive(a any) error {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingHandlerInterceptor) RequestHeader() http.Header {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingHandlerInterceptor) Send(a any) error {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (s *streamingHandlerInterceptor) ResponseHeader() http.Header {
// 	// TODO implement me
// 	panic("implement me")
// }

func (s *streamingHandlerInterceptor) ResponseTrailer() http.Header {
	// TODO implement me
	panic("implement me")
}
