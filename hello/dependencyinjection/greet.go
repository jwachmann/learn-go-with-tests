package dependencyinjection

import (
	"fmt"
	"io"
)

// Greet greets a person by the given name
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}
