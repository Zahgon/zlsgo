package zcli

import (
	"sync"

	"github.com/sohaha/zlsgo/zutil/daemon"
)

type (
	// app implements the daemon service interface for running the application as a service
	app struct {
		run    func()
		done   chan struct{}
		err    error
		mu     sync.Mutex
		status bool
	}
	// serviceStop implements the Cmd interface for stopping a service
	serviceStop struct {
	}
	// serviceStart implements the Cmd interface for starting a service
	serviceStart struct {
	}
	// serviceRestart implements the Cmd interface for restarting a service
	serviceRestart struct {
	}
	// serviceInstall implements the Cmd interface for installing a service
	serviceInstall struct {
	}
	// serviceUnInstall implements the Cmd interface for uninstalling a service
	serviceUnInstall struct {
	}
	// serviceStatus implements the Cmd interface for checking a service's status
	serviceStatus struct {
	}
)

var (
	service    daemon.ServiceIface
	serviceErr error
	once       sync.Once
)

// Start implements the daemon.ServiceIface Start method for the app type.
// It runs the application function in a goroutine and returns any error that occurs.
func (a *app) Start(daemon.ServiceIface) error { _ = "STUB: not implemented"; return nil }

// Stop implements the daemon.ServiceIface Stop method for the app type.
// It waits for the application to stop with a timeout of 30 seconds.
func (a *app) Stop(daemon.ServiceIface) error { _ = "STUB: not implemented"; return nil }

// return errors.New("forced timeout")

func (a *app) lastErr() error { _ = "STUB: not implemented"; return nil }

// Flags implements the Cmd interface for the serviceStatus command.
// It checks for service errors before proceeding.
func (*serviceStatus) Flags(_ *Subcommand) { _ = "STUB: not implemented"; return }

// Run implements the Cmd interface for the serviceStatus command.
// It displays the current status of the service.
func (*serviceStatus) Run(_ []string) { _ = "STUB: not implemented"; return }

// Flags implements the Cmd interface for the serviceInstall command.
// It checks for service errors before proceeding.
func (*serviceInstall) Flags(_ *Subcommand) { _ = "STUB: not implemented"; return }

// Run implements the Cmd interface for the serviceInstall command.
// It installs and starts the service.
func (*serviceInstall) Run(_ []string) { _ = "STUB: not implemented"; return }

// Flags implements the Cmd interface for the serviceUnInstall command.
// It checks for service errors before proceeding.
func (*serviceUnInstall) Flags(_ *Subcommand) { _ = "STUB: not implemented"; return }

// Run implements the Cmd interface for the serviceUnInstall command.
// It uninstalls the service from the system.
func (*serviceUnInstall) Run(_ []string) { _ = "STUB: not implemented"; return }

// Flags implements the Cmd interface for the serviceStart command.
// It checks for service errors before proceeding.
func (*serviceStart) Flags(_ *Subcommand) { _ = "STUB: not implemented"; return }

// Run implements the Cmd interface for the serviceStart command.
// It starts the service if it is not already running.
func (*serviceStart) Run(_ []string) { _ = "STUB: not implemented"; return }

// Flags implements the Cmd interface for the serviceStop command.
// It checks for service errors before proceeding.
func (*serviceStop) Flags(_ *Subcommand) { _ = "STUB: not implemented"; return }

// Run implements the Cmd interface for the serviceStop command.
// It stops the service if it is running.
func (*serviceStop) Run(_ []string) { _ = "STUB: not implemented"; return }

// Flags implements the Cmd interface for the serviceRestart command.
// It checks for service errors before proceeding.
func (*serviceRestart) Flags(_ *Subcommand) { _ = "STUB: not implemented"; return }

// Run implements the Cmd interface for the serviceRestart command.
// It restarts the service.
func (*serviceRestart) Run(_ []string) { _ = "STUB: not implemented"; return }

// LaunchServiceRun initializes a service with the given name and description,
// and runs it immediately. If the --detach flag is set, it will run the service
// in the background. If no service system is available, it runs the function directly.
func LaunchServiceRun(name string, description string, fn func(), config ...*daemon.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// LaunchService initializes a service with the given name, description, and run function.
// It also registers service management commands (install, uninstall, status, etc.).
// Returns the service interface and any error that occurred during initialization.
func LaunchService(name string, description string, fn func(), config ...*daemon.Config) (daemon.ServiceIface, error) {
	_ = "STUB: not implemented"
	return *new(daemon.ServiceIface), nil
}

// The file path is redirected to the current execution file path
