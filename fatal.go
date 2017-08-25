package nestederror

import (
	"log"
	"strings"
)

// Fatalln join all errors with ' -> ' as a separator then print it with log.Fatalln.
func Fatalln(err error) {
	msg := strings.Join(Flatten(err).Strings(), " -> ")
	log.Fatalln(msg)
}
