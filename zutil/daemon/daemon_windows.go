package daemon

import (
	"sync"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type (
	windowsSystem  struct{}
	windowsService struct {
		i            Iface
		stopStartErr error
		*Config
		errSync sync.Mutex
	}
)

const version = "windows-service"

var interactive = false

func init() {
	var err error
	chooseSystem(windowsSystem{})
	interactive, err = svc.IsAnInteractiveSession()
	if err != nil {
		panic(err)
	}
}

func (windowsSystem) String() string { _ = "STUB: not implemented"; return "" }

func (windowsSystem) Detect() bool { _ = "STUB: not implemented"; return false }

func (windowsSystem) Interactive() bool { _ = "STUB: not implemented"; return false }

func (windowsSystem) New(i Iface, c *Config) (ServiceIface, error) {
	_ = "STUB: not implemented"
	return *new(ServiceIface), nil
}

func (w *windowsService) String() string { _ = "STUB: not implemented"; return "" }

func (w *windowsService) setError(err error) { _ = "STUB: not implemented"; return }

func (w *windowsService) getError() error { _ = "STUB: not implemented"; return nil }

func (w *windowsService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (bool, uint32) {
	_ = "STUB: not implemented"
	return false, 0
}

func (w *windowsService) Install() error { _ = "STUB: not implemented"; return nil }

func (w *windowsService) Uninstall() error { _ = "STUB: not implemented"; return nil }

func (w *windowsService) Run() error { _ = "STUB: not implemented"; return nil }

func (w *windowsService) Start() error { _ = "STUB: not implemented"; return nil }

func (w *windowsService) Stop() error { _ = "STUB: not implemented"; return nil }

func (w *windowsService) Restart() error { _ = "STUB: not implemented"; return nil }

func (w *windowsService) Status() string { _ = "STUB: not implemented"; return "" }

func (w *windowsService) forceKeep(processId uint32) error { _ = "STUB: not implemented"; return nil }

func (w *windowsService) stopWait(s *mgr.Service) error { _ = "STUB: not implemented"; return nil }

func connect() (*mgr.Mgr, error) { _ = "STUB: not implemented"; return nil, nil }

func getStopTimeout() time.Duration {
	_ = "STUB: not implemented"
	// For default and paths see https://support.microsoft.com/en-us/kb/146092
	return *new(time.Duration)
}
