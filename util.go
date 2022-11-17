package zerolog

import "fmt"

const (
	DefaultLogSourceKey     = "logger"
	DefaultLogCommonSource  = "common"
	DefaultLogStorageSource = "storage"
	DefaultLogWebSource     = "web"
	DefaultLogSourceColor   = 90
)

// Colorize colors the string in log.
func Colorize(c int, m any) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, m)
}
