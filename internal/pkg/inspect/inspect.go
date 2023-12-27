package inspect

import (
	"bwizard/internal/pkg/ssh"
	"slices"
	"strings"
)

type DeviceInspection struct {
	Hostname        string
	CPU             string
	Ram             string
	Disks           []string
	Agent           string
	OperatingSystem string
}

// Device inspects a device regarding hostname, cpu, memory and disk size
func Device(client *ssh.Client) (*DeviceInspection, error) {
	hostnameCmd := "hostname"
	hostnameOutput, err := client.RunCommand(hostnameCmd)
	if err != nil {
		return nil, err
	}

	cpuCmd := "lscpu | grep 'Model name' | cut -f 2 -d \":\" | awk '{$1=$1}1'"
	cpuOutput, err := client.RunCommand(cpuCmd)
	if err != nil {
		return nil, err
	}

	ramCmd := "cat /proc/meminfo | grep 'MemTotal' | cut -f 2 -d \":\" | awk '{$1=$1}1'"
	ramOutput, err := client.RunCommand(ramCmd)
	if err != nil {
		return nil, err
	}

	disksCmd := "for disk in $(lsblk | grep -E \"^(sd[a-z]|vd[a-z])\" | awk '{print $1}'); do echo \"$disk: $(lsblk | grep \"^$disk \" | awk '{print $4}')\"; done"
	disksOutput, err := client.RunCommand(disksCmd)
	if err != nil {
		return nil, err
	}

	agentCmd := "ssh -V"
	agentOutput, err := client.RunCommand(agentCmd)
	if err != nil {
		return nil, err
	}

	osCmd := "cat /etc/os-release | grep 'PRETTY_NAME' | cut -f 2 -d \"=\" | awk '{$1=$1}1'"
	osOutput, err := client.RunCommand(osCmd)
	if err != nil {
		return nil, err
	}

	disks := slices.DeleteFunc(strings.Split(string(disksOutput), "\n"), func(e string) bool {
		return e == ""
	})

	return &DeviceInspection{
		Hostname:        removeNewLine(string(hostnameOutput)),
		CPU:             removeNewLine(string(cpuOutput)),
		Ram:             removeNewLine(string(ramOutput)),
		Agent:           removeNewLine(string(agentOutput)),
		Disks:           disks,
		OperatingSystem: removeNewLine(strings.ReplaceAll(string(osOutput), "\"", "")),
	}, nil
}

func removeNewLine(in string) string {
	return strings.ReplaceAll(in, "\n", "")
}
