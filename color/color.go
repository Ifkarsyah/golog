package color

const (
	Reset = "\033[0m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
	Purple = "\033[35m"
)

func Ize(color string, msg string) string {
	return color + msg + Reset
}


