package goserver

import (
	"fmt"
)

func Serve() string {
	return fmt.Sprintf("This is the %s talking", "server")
}
