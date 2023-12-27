package ssh

import (
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"strconv"
)

type Credentials struct {
	Username       string
	PrivateKeyPath string
}

type Client struct {
	host   string
	port   uint16
	config *ssh.ClientConfig
	client *ssh.Client
}

// NewClient returns a new SSH client
func NewClient(host string, port uint16, credentials *Credentials) (*Client, error) {
	privateKey, err := os.ReadFile(credentials.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: credentials.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		// TODO: Only use that for development. Implement a proper host key check
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return &Client{
		host:   host,
		port:   port,
		config: config,
	}, nil
}

// Connect starts a SSH session to the server
func (c *Client) Connect() error {
	client, err := ssh.Dial("tcp", net.JoinHostPort(c.host, strconv.Itoa(int(c.port))), c.config)
	if err != nil {
		return err
	}

	c.client = client

	return nil
}

// Close closes the connection to the server
func (c *Client) Close() error {
	if c.client == nil {
		return nil
	}

	return c.client.Close()
}

// NewSession returns a new session to run command in
func (c *Client) NewSession() (*ssh.Session, error) {
	return c.client.NewSession()
}

func (c *Client) RunCommand(cmd string) ([]byte, error) {
	session, err := c.NewSession()
	if err != nil {
		return nil, err
	}

	defer func(session *ssh.Session) {
		_ = session.Close()
	}(session)

	return session.CombinedOutput(cmd)
}
