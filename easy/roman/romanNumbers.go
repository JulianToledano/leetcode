package main

func main() {
	romanToInt("MMCCCXCIX")
}

func romanToInt(s string) int {
	romanNumbers := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	total := 0
	inRow := 1
	lastChar := string(s[0])
	temp := romanNumbers[lastChar]
	for c := range s[1:] {
		currentChar := string(s[c+1])
		if currentChar == lastChar {
			temp += romanNumbers[currentChar]
			inRow++
		} else {
			if romanNumbers[(currentChar)] > romanNumbers[(lastChar)] {
				lastChar = currentChar
				temp = romanNumbers[lastChar] - temp
				total += temp
				temp = 0
			} else {
				total += temp
				lastChar = currentChar
				temp = romanNumbers[lastChar]
			}

		}
	}
	total += temp
	return total
}
