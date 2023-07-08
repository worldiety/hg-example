// Package helloworld is a very small domain core without any I/O adapters, so its fine to just put it right
// into the bounded context folder!?
package helloworld

import "fmt"

// SayHello greets the given persona.
func SayHello(whom string) string {
	return fmt.Sprintf("hello %s", whom)
}
