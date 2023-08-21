package piscine

func ConcatParams(args []string) string {
	var result string
	argsLength := 0
	for range args {
		argsLength++
	}
	for i, word := range args {
		if i == argsLength-1 {
			result += word
		} else {
			result += word + "\n"
		}
	}
	return result
}
