package zcli

import (
	"flag"
)

// errorText formats a string with red color for error messages.
func errorText(msg string) string { _ = "STUB: not implemented"; return "" }

// tipText formats a string with green color for tips and success messages.
func tipText(msg string) string { _ = "STUB: not implemented"; return "" }

// warnText formats a string with yellow color for warning messages.
func warnText(msg string) string { _ = "STUB: not implemented"; return "" }

// showText formats a string with light grey color for regular informational text.
func showText(msg string) string { _ = "STUB: not implemented"; return "" }

// Help displays usage information for the current command or subcommand and exits the program.
// If a subcommand is active, it shows help for that specific subcommand.
func Help() { _ = "STUB: not implemented"; return }

// numOfGlobalFlags returns the number of global flags defined in the application.
func numOfGlobalFlags() (count int) { _ = "STUB: not implemented"; return 0 }

// Error prints a formatted error message, optionally shows a help tip, and exits the program with status code 1.
func Error(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// showRequired displays a list of required flags that must be provided by the user.
func showRequired(_ *flag.FlagSet, requiredFlags []string) { _ = "STUB: not implemented"; return }

// showSubcommandUsage displays the usage information for a subcommand using its flag set.
func showSubcommandUsage(fs *flag.FlagSet, _ *cmdCont) {
	_ = "STUB: not implemented"

	// showLogo displays the application logo if one is defined.
	// Returns true if a logo was displayed, false otherwise.
	return
}

func showLogo() bool { _ = "STUB: not implemented"; return false }

// showFlagsHelp displays help information for the help flag unless help display is disabled.
func showFlagsHelp() { _ = "STUB: not implemented"; return }

// showDescription displays the application name and version if defined.
// The logoOk parameter indicates whether a logo was already displayed to avoid redundancy.
// Returns true if a description was displayed, false otherwise.
func showDescription(logoOk bool) bool { _ = "STUB: not implemented"; return false }

// showVersion displays the application version if defined.
// Returns true if a version was displayed, false otherwise.
func showVersion() bool { _ = "STUB: not implemented"; return false }

// showVersionNum displays version information.
// If info is true, it shows detailed build information including Go version and build time.
// Otherwise, it only shows the version number.
func showVersionNum(info bool) { _ = "STUB: not implemented"; return }

//noinspection GoBoolExpressions

//noinspection GoBoolExpressions

//noinspection GoBoolExpressions

// showHeadr displays the application header, which may include the logo, name, and version.
func showHeadr() { _ = "STUB: not implemented"; return }

// argsIsHelp checks if any of the provided arguments is a help flag (-h, -help, --help).
// If found, it sets the global help flag to true.
func argsIsHelp(args []string) { _ = "STUB: not implemented"; return }

// CheckErr handles errors by logging them and optionally exiting the program.
// If exit is true, it calls os.Exit(1) after logging the error.
// It has special handling for service system detection errors.
func CheckErr(err error, exit ...bool) { _ = "STUB: not implemented"; return }

// isDetach checks if the provided argument is a detach flag (D or detach).
// Returns true if it is a detach flag, false otherwise.
func isDetach(a string) bool { _ = "STUB: not implemented"; return false }

// IsSudo checks if the current process is running with sudo/administrator privileges.
// Returns true if running with elevated privileges, false otherwise.
func IsSudo() bool { _ = "STUB: not implemented"; return false }

// IsDoubleClickStartUp detects if the application was started by double-clicking
// rather than from a command line.
// Returns true if started by double-click, false otherwise.
func IsDoubleClickStartUp() bool { _ = "STUB: not implemented"; return false }

// LockInstance ensures only one instance of the application can run at a time.
// Returns a cleanup function and a boolean indicating whether the lock was acquired.
// If the lock was not acquired (ok is false), another instance is already running.
func LockInstance() (clean func(), ok bool) { _ = "STUB: not implemented"; return nil, false }
