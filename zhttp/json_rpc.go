package zhttp

import (
	"crypto/tls"
	"net/http"
	"net/rpc"
	"time"
)

type JSONRPC struct {
	client  *rpc.Client
	path    string
	address string
	options JSONRPCOptions
}

func (j *JSONRPC) Call(serviceMethod string, args interface{}, reply interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JSONRPC) Go(serviceMethod string, args interface{}, reply interface{}, done chan *rpc.Call) *rpc.Call {
	_ = "STUB: not implemented"
	return nil
}

func (j *JSONRPC) Close() error { _ = "STUB: not implemented"; return nil }

func (j *JSONRPC) connect() error { _ = "STUB: not implemented"; return nil }

type JSONRPCOptions struct {
	TlsConfig  *tls.Config
	Header     http.Header
	Timeout    time.Duration
	RetryDelay time.Duration
	Retry      bool
}

func NewJSONRPC(address string, path string, opts ...func(o *JSONRPCOptions)) (client *JSONRPC, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
