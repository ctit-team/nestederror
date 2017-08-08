package nestederror

import (
	"fmt"
	"log"
	"strings"
)

// Fatalln join all errors with ' -> ' as a separator then print it with log.Fatalln.
func Fatalln(err error, format string, args ...interface{}) {
	const sep = " -> "

	prefix := fmt.Sprintf(format, args...)
	joined := strings.Join(Flatten(err).Strings(), sep)

	log.Fatalln(prefix + sep + joined)
}
