package main

func main() {

}

func numJewelsInStones(J string, S string) int {
	jewels := make(map[byte]int)
	total := 0
	for i := range J {
		jewels[J[i]] = 0
	}
	for i := range S {
		if _, ok := jewels[S[i]]; ok {
			total++
		}
	}
	return total
}
