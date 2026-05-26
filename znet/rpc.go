package znet

// JSONRPCOption defines configuration options for the JSON-RPC handler.
type JSONRPCOption struct {
	DisabledHTTP bool // Disables HTTP method handling, only processes RPC calls
	Debug        bool // Enables debug mode with additional logging and method inspection
}

// JSONRPC creates a handler that processes JSON-RPC requests.
// It registers the provided receiver objects with the RPC server and returns a handler function
// that can be used with the router. The handler supports both HTTP and WebSocket transports.
func JSONRPC(rcvr map[string]interface{}, opts ...func(o *JSONRPCOption)) func(c *Context) {
	_ = "STUB: not implemented"
	return nil
}
