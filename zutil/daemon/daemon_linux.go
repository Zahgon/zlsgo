package daemon

import (
	"errors"
	"strings"
	"text/template"
)

type (
	linuxSystemService struct {
		detect      func() bool
		interactive func() bool
		new         func(i Iface, c *Config) (ServiceIface, error)
		name        string
	}
	systemd struct {
		i Iface
		*Config
	}
)

const (
	optionReloadSignal = "ReloadSignal"
	optionPIDFile      = "PIDFile"
)

var errNoUserServiceSystemd = errors.New("user services are not supported on systemd")

func init() {
	chooseSystem(linuxSystemService{
		name:   "linux-systemd",
		detect: isSystemd,
		interactive: func() bool {
			is, _ := isInteractive()
			return is
		},
		new: newSystemdService,
	})
}

func (sc linuxSystemService) String() string { _ = "STUB: not implemented"; return "" }

func (sc linuxSystemService) Detect() bool { _ = "STUB: not implemented"; return false }

func (sc linuxSystemService) Interactive() bool { _ = "STUB: not implemented"; return false }

func (sc linuxSystemService) New(i Iface, c *Config) (s ServiceIface, err error) {
	_ = "STUB: not implemented"
	return *new(ServiceIface), nil
}

func isInteractive() (bool, error) { _ = "STUB: not implemented"; return false, nil }

var tf = map[string]interface{}{
	"cmd": func(s string) string {
		return `"` + strings.Replace(s, `"`, `\"`, -1) + `"`
	},
	"cmdEscape": func(s string) string {
		return strings.Replace(s, " ", `\x20`, -1)
	},
}

func isSystemd() bool { _ = "STUB: not implemented"; return false }

func newSystemdService(i Iface, c *Config) (ServiceIface, error) {
	_ = "STUB: not implemented"
	return *new(ServiceIface), nil
}

func (s *systemd) String() string { _ = "STUB: not implemented"; return "" }

func (s *systemd) configPath() (cp string, err error) { _ = "STUB: not implemented"; return "", nil }

func (s *systemd) template() *template.Template { _ = "STUB: not implemented"; return nil }

func (s *systemd) Install() error { _ = "STUB: not implemented"; return nil }

func (s *systemd) Uninstall() error { _ = "STUB: not implemented"; return nil }

func (s *systemd) Run() (err error) { _ = "STUB: not implemented"; return nil }

func (s *systemd) Start() error { _ = "STUB: not implemented"; return nil }

func (s *systemd) Stop() error { _ = "STUB: not implemented"; return nil }

func (s *systemd) Restart() error { _ = "STUB: not implemented"; return nil }

func (s *systemd) Status() string { _ = "STUB: not implemented"; return "" }

const systemdScript = `[Unit]
Description={{.Description}}
ConditionFileIsExecutable={{.Path|cmdEscape}}

[Service]
ExecStart={{.Path|cmdEscape}}{{range .Arguments}} {{.|cmd}}{{end}}
{{if .RootDir}}RootDirectory={{.RootDir|cmd}}{{end}}
{{if .WorkingDir}}WorkingDirectory={{.WorkingDir|cmdEscape}}{{end}}
{{if .UserName}}User={{.UserName}}{{end}}
{{if .ReloadSignal}}ExecReload=/bin/kill -{{.ReloadSignal}} "$MAINPID"{{end}}
{{if .PIDFile}}PIDFile={{.PIDFile|cmd}}{{end}}
Restart=on-failure
RestartSec=1
EnvironmentFile=-/etc/sysconfig/{{.Name}}
KillMode=process

[Install]
WantedBy=multi-user.target
`
