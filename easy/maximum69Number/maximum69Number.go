package main

func main() {

}

func maximum69Number(num int) int {

	if num/1000 == 6 {
		return num + 3000
	} else if num%1000/100 == 6 {
		return num + 300
	} else if num%100/10 == 6 {
		return num + 30
	} else if num%10 == 6 {
		return num + 3
	}

	return num
}
