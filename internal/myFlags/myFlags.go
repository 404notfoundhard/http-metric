package myflags

import (
	"errors"
	"fmt"
	"strings"
)

type ListenAddres struct {
	Host string `env:"ADDRESS"`
	Port string
}

func (la *ListenAddres) String() string {
	return la.Host + ":" + la.Port
}

func (la *ListenAddres) Set(s string) error {
	hp := strings.Split(s, ":")
	if len(hp) != 2 {
		return errors.New("need address in a form host:port")
	}
	port := hp[1]
	if port == "" {
		return fmt.Errorf("port must be set")
	}
	la.Host = hp[0]
	la.Port = port
	return nil
}
