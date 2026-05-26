package zcli

// SingleKillSignal returns a channel that will receive a value when the application
// receives a termination signal (such as SIGINT or SIGTERM).
// This can be used to implement graceful shutdown handling.
func SingleKillSignal() <-chan bool { _ = "STUB: not implemented"; return nil }
