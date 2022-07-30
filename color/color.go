package color

import "fmt"

const (
	defaultt = "\x1b[39m"
	red      = "\x1b[91m"
	green    = "\x1b[32m"
	blue     = "\x1b[94m"
	gray     = "\x1b[90m"
)

func Red(s string) string {
	return fmt.Sprintf("%s%s%s", red, s, defaultt)
}

func Green(s string) string {
	return fmt.Sprintf("%s%s%s", green, s, defaultt)
}

func Blue(s string) string {
	return fmt.Sprintf("%s%s%s", blue, s, defaultt)
}

func Gray(s string) string {
	return fmt.Sprintf("%s%s%s", gray, s, defaultt)
}
