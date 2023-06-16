package logstash

import (
	"errors"
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"net"
	"time"
)

type Logstash struct {
	Hostname   string
	Port       int
	Connection *net.TCPConn
	Timeout    int
}

// New func
func New(hostname string, port int, timeout int) *Logstash {
	l := Logstash{}
	l.Hostname = hostname
	l.Port = port
	l.Connection = nil
	l.Timeout = timeout
	return &l
}

// NewList func
func NewList(listAddr []string, timeout int) []*Logstash {
	resList := make([]*Logstash, 0)

	for _, addr := range listAddr {
		l := &Logstash{}

		host, port, err := base.ParseURL(addr)
		if err != nil {
			fmt.Println(fmt.Sprintf("Address invalid: %s, error: %+v", addr, err))

			continue
		}

		l.Hostname = host
		l.Port = port
		l.Connection = nil
		l.Timeout = timeout

		resList = append(resList, l)
	}

	return resList
}

// Dump func
func (l *Logstash) Dump() {
	fmt.Println("Hostname:   ", l.Hostname)
	fmt.Println("Port:       ", l.Port)
	fmt.Println("Connection: ", l.Connection)
	fmt.Println("Timeout:    ", l.Timeout)
}

// String func
func (l *Logstash) String() string {
	return fmt.Sprintf("Hostname: %s, Port: %d, Timeout: %d", l.Hostname, l.Port, l.Timeout)
}

// SetTimeouts func
func (l *Logstash) SetTimeouts() {
	deadline := time.Now().Add(time.Duration(l.Timeout) * time.Second)
	l.Connection.SetDeadline(deadline)
	l.Connection.SetWriteDeadline(deadline)
	l.Connection.SetReadDeadline(deadline)
}

// Connect func
func (l *Logstash) Connect() (*net.TCPConn, error) {
	var connection *net.TCPConn

	service := fmt.Sprintf("%s:%d", l.Hostname, l.Port)
	addr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		return connection, err
	}

	connection, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		return connection, err
	}

	if connection != nil {
		l.Connection = connection
		l.Connection.SetLinger(0) // default -1
		l.Connection.SetNoDelay(true)
		l.Connection.SetKeepAlive(true)
		l.Connection.SetKeepAlivePeriod(time.Duration(5) * time.Second)
		l.SetTimeouts()
	}

	return connection, err
}

// Writeln func
func (l *Logstash) Writeln(message string) error {
	var err = errors.New("tcp connection is nil")
	message = fmt.Sprintf("%s\n", message)
	if l.Connection != nil {
		n, err := l.Connection.Write([]byte(message))
		if err != nil {
			fmt.Println(fmt.Sprintf("Gostash::Writeln - Error: %+v", err))

			myAddr := l.Connection.RemoteAddr()
			l.Connection.Close()
			l.Connection = nil

			_, conErr := l.Connect()
			fmt.Println(fmt.Sprintf("Gostash::Writeln - Re-Connecting to [%+v] conErr: %+v", myAddr, conErr))

			return conErr
		}

		fmt.Println(fmt.Sprintf("Gostash::Writeln - Connection write n: %d", n))

		// Successful write! Let's extend the timeout.
		l.SetTimeouts()

		return nil
	}

	l.Connect()

	return err
}
