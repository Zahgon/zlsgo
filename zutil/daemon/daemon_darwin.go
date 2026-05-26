package daemon

import (
	"github.com/sohaha/zlsgo/zerror"
)

type (
	darwinSystem         struct{}
	darwinLaunchdService struct {
		i Iface
		*Config
		userService bool
	}
)

const version = "darwin-launchd"

var interactive = false

func (darwinSystem) String() string { _ = "STUB: not implemented"; return "" }

func (darwinSystem) Detect() bool { _ = "STUB: not implemented"; return false }

func (darwinSystem) Interactive() bool { _ = "STUB: not implemented"; return false }

func (darwinSystem) New(i Iface, c *Config) (s ServiceIface, err error) {
	_ = "STUB: not implemented"
	return *new(ServiceIface), nil
}

func init() {
	var err error
	chooseSystem(darwinSystem{})
	interactive, err = isInteractive()
	zerror.Panic(err)
}

func isInteractive() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s *darwinLaunchdService) String() string { _ = "STUB: not implemented"; return "" }

func (s *darwinLaunchdService) getHomeDir() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *darwinLaunchdService) getServiceFilePath() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *darwinLaunchdService) Install() error { _ = "STUB: not implemented"; return nil }

// ~/Library/LaunchAgents exists

func (s *darwinLaunchdService) Uninstall() error { _ = "STUB: not implemented"; return nil }

func (s *darwinLaunchdService) Start() error { _ = "STUB: not implemented"; return nil }

func (s *darwinLaunchdService) Stop() error { _ = "STUB: not implemented"; return nil }

func (s *darwinLaunchdService) Status() string { _ = "STUB: not implemented"; return "" }

func (s *darwinLaunchdService) Restart() error { _ = "STUB: not implemented"; return nil }

func (s *darwinLaunchdService) Run() error { _ = "STUB: not implemented"; return nil }

var launchdConfig = `<?xml version='1.0' encoding='UTF-8'?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN"
"http://www.apple.com/DTDs/PropertyList-1.0.dtd" >
<plist version='1.0'>
<dict>
<key>Label</key><string>{{html .Name}}</string>
<key>ProgramArguments</key>
<array>
        <string>{{html .Path}}</string>
{{range .Config.Arguments}}
        <string>{{html .}}</string>
{{end}}
</array>
{{if .UserName}}<key>UserName</key><string>{{html .UserName}}</string>{{end}}
{{if .RootDir}}<key>RootDirectory</key><string>{{html .RootDir}}</string>{{end}}
{{if .WorkingDir}}<key>WorkingDirectory</key><string>{{html .WorkingDir}}</string>{{end}}
<key>SessionCreate</key><{{bool .SessionCreate}}/>
<key>KeepAlive</key><{{bool .KeepAlive}}/>
<key>RunAtLoad</key><{{bool .RunAtLoad}}/>
<key>Disabled</key><false/>
</dict>
</plist>
`

// <key>StandardOutPath</key>
// <string>/tmp/zlsgo/{{html .Name}}.log</string>
// <key>StandardErrorPath</key>
// <string>/tmp/zlsgo/{{html .Name}}.err</string>
