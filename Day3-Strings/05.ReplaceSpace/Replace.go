
func replaceSpace(s string) string {
	var str string = ""
	for _, v := range s {
		if v == ' ' {
			str += "%20"
		} else {
			str += string(v)
		}
	}
	return str
}

/*
import "strings"

func replaceSpace(s string) string{
	return strings.Replace(s, " ", "%20f", -1)
}
*/