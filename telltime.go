//package main creates an executable program
package main

//Import multiple package and alias for time package
import (
	"fmt"
	t "time"
)

func main() {
	// Print Hello World
	fmt.Println("Hello World!", t.Now())
}
