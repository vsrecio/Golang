package main

func main() {
	abcGen()
}

func abcGen() {
	for i := byte('a');  i <= byte('z'); i++ {
		println(string(i))
	}
}
