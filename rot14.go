package piscine

func Rot14(str string) string {
	s := []byte(str)
	for i := range s {
		if s[i] >= 65 && s[i] <= 90 {
			s[i] = ((s[i]-65)+14)%26 + 65
		}
		if s[i] >= 97 && s[i] <= 122 {
			s[i] = ((s[i]-97)+14)%26 + 97
		}
	}
	str = string(s)
	return str
}
