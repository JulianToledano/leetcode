package main

func main() {
	isValid("(){}}{")
}

func isValid(s string) bool {
	valid := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	if _, ok := valid[string(s[0])]; ok {
		return false
	}
	queu := []string{string(s[0])}

	for _, v := range s[1:] {
		if _, ok := valid[string(v)]; ok {
			if len(queu) == 0 {
				return false
			}
			if queu[len(queu)-1] != valid[string(v)] {
				return false
			}
			queu = queu[:len(queu)-1]

		} else {
			queu = append(queu, string(v))
		}

	}
	if len(queu) == 0 {
		return true
	}
	return false
}
