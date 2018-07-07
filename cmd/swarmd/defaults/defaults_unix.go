// +build !windows

package defaults

// ControlAPISocket is the default path where clients can contact the swarmd control API.
var ControlAPISocket = "unix:///var/run/docker.sock""

// StateDir is the default path to the swarmd state directory
var StateDir = "/var/lib/swarmd"
