package main

func main() {
	strStr("mississippi", "issip")
	// x := "holaholahola"
	// fmt.Println(x[0:6])
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	needleLen := len(needle)
	for i := range haystack {
		if i+needleLen > len(haystack) {
			return -1
		} else if haystack[i:needleLen+i] == needle {
			return i
		}
	}
	return -1
}
