package main

func calc(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	a, b := calc(10, 20)
	println(a, b)

}
