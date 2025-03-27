package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	xForwardedForHeader        = "x-forwarded-for"
	userAgentHeader            = "user-agent"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	meta := &Metadata{}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgent := md.Get(grpcGatewayUserAgentHeader); len(userAgent) > 0 {
			meta.UserAgent = userAgent[0]
		}
		if clientIP := md.Get(xForwardedForHeader); len(clientIP) > 0 {
			meta.ClientIP = clientIP[0]
		}
		if userAgent := md.Get(userAgentHeader); len(userAgent) > 0 {
			meta.UserAgent = userAgent[0]
		}
	}
	if p, ok := peer.FromContext(ctx); ok {
		meta.ClientIP = p.Addr.String()
	}
	return meta
}
