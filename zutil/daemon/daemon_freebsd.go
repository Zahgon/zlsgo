package daemon

type (
	freebsdRcdService struct {
		i Iface
		*Config
		userService bool
	}
	freebsdSystem struct{}
)

const version = "freebsd-rcd"

var interactive = false

func init() {
	var err error
	chooseSystem(freebsdSystem{})
	interactive, err = isInteractive()
	if err != nil {
		panic(err)
	}
}

func (freebsdSystem) String() string { _ = "STUB: not implemented"; return "" }

func (freebsdSystem) Detect() bool { _ = "STUB: not implemented"; return false }

func (freebsdSystem) Interactive() bool { _ = "STUB: not implemented"; return false }

func (freebsdSystem) New(i Iface, c *Config) (ServiceIface, error) {
	_ = "STUB: not implemented"
	return *new(ServiceIface), nil
}

func isInteractive() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s *freebsdRcdService) Status() string { _ = "STUB: not implemented"; return "" }

func (s *freebsdRcdService) String() string { _ = "STUB: not implemented"; return "" }

func (s *freebsdRcdService) getServiceFilePath() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *freebsdRcdService) Install() error { _ = "STUB: not implemented"; return nil }

//  ~/Library/LaunchAgents exists.

func (s *freebsdRcdService) Uninstall() error { _ = "STUB: not implemented"; return nil }

func (s *freebsdRcdService) Start() error { _ = "STUB: not implemented"; return nil }

func (s *freebsdRcdService) Stop() error { _ = "STUB: not implemented"; return nil }

func (s *freebsdRcdService) Restart() error { _ = "STUB: not implemented"; return nil }

func (s *freebsdRcdService) Run() error { _ = "STUB: not implemented"; return nil }

const rcdScriptOpsrampAgent = `. /etc/rc.subr

name="opsramp_agent"
rcvar="opsramp_agent_enable"
command="/opt/opsramp/agent/opsramp-agent service"
pidfile="/var/run/${name}.pid"

start_cmd="test_start"
stop_cmd="test_stop"
status_cmd="test_status"

test_start() {
        /usr/sbin/daemon -p ${pidfile} ${command}
}

test_status() {
        if [ -e ${pidfile} ]; then
                echo ${name} is running...
        else
                echo ${name} is not running.
        fi
}

test_stop() {
        if [ -e ${pidfile} ]; then
` +
	" kill `cat ${pidfile}`; " +
	`
        else
                echo ${name} is not running?
        fi
}

load_rc_config $name
run_rc_command "$1"
`

const rcdScriptOpsrampShield = `. /etc/rc.subr

name="opsramp_shield"
rcvar="opsramp_shield_enable"
command="/opt/opsramp/agent/bin/opsramp-shield service"
pidfile="/var/run/${name}.pid"

start_cmd="test_start"
stop_cmd="test_stop"
status_cmd="test_status"

test_start() {
        /usr/sbin/daemon -p ${pidfile} ${command}
}

test_status() {
        if [ -e ${pidfile} ]; then
                echo ${name} is running...
        else
                echo ${name} is not running.
        fi
}

test_stop() {
        if [ -e ${pidfile} ]; then
` +
	" kill `cat ${pidfile}`; " +
	`
        else
                echo ${name} is not running?
        fi
}

load_rc_config $name
run_rc_command "$1"
`
const rcdScriptAgentUninstall = `. /etc/rc.subr

name="agent_uninstall"
rcvar="agent_uninstall_enable"
command="/opt/opsramp/agent/bin/uninstall"
pidfile="/var/run/${name}.pid"

start_cmd="test_start"
stop_cmd="test_stop"
status_cmd="test_status"

test_start() {
        /usr/sbin/daemon -p ${pidfile} ${command}
}

test_status() {
        if [ -e ${pidfile} ]; then
                echo ${name} is running...
        else
                echo ${name} is not running.
        fi
}

test_stop() {
        if [ -e ${pidfile} ]; then
` +
	" kill `cat ${pidfile}`; " +
	`
        else
                echo ${name} is not running?
        fi
}

load_rc_config $name
run_rc_command "$1"
`
