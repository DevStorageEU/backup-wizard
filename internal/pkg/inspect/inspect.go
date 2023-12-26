package inspect

import (
	"bwizard/internal/pkg/ssh"
	ssh2 "golang.org/x/crypto/ssh"
)

type DeviceInspection struct {
	Hostname string
	CPU      string
	Ram      string
}

// Device inspects a device regarding hostname, cpu, memory and disk size
func Device(client *ssh.Client) (*DeviceInspection, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	defer func(session *ssh2.Session) {
		err := session.Close()
		if err != nil {
			panic(err)
		}
	}(session)

	hostnameCmd := "hostname"
	hostnameOutput, err := session.CombinedOutput(hostnameCmd)
	if err != nil {
		return nil, err
	}

	cpuCmd := "lscpu | grep 'Model name' | cut -f 2 -d \":\" | awk '{$1=$1}1'"
	cpuOutput, err := session.CombinedOutput(cpuCmd)
	if err != nil {
		return nil, err
	}

	ramCmd := "cat /proc/meminfo | grep 'MemTotal' | cut -f 2 -d \":\" | awk '{$1=$1}1'"
	ramOutput, err := session.CombinedOutput(ramCmd)
	if err != nil {
		return nil, err
	}

	return &DeviceInspection{
		Hostname: string(hostnameOutput),
		CPU:      string(cpuOutput),
		Ram:      string(ramOutput),
	}, nil
}
