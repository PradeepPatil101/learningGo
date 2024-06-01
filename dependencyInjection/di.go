// The main purpose of this is to show that Greet can be used in application as well as for testing

package dependencyInjection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// func main() {
// 	Greet(os.Stdout, "Elodie")
// }
