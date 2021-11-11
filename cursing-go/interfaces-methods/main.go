package main

import "fmt"

type retangulo struct {
	altura, largura float64
}

type cicle struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c cicle) area() float64 {
	return c.raio * c.raio * 3.14
}

type IForm interface {
	area() float64
}

func isArea(f IForm) float64 {
	return f.area()
}

func generical(value ...interface{}) {
	for _, v := range value {
		fmt.Println(v)
	}
}
func main() {
	//interfaces
	r := retangulo{altura: 10, largura: 5}
	c := cicle{raio: 5}

	fmt.Println(isArea(r))
	fmt.Println(isArea(c))
	//interfaces genericas
	generical(1, "Picanha Paulo", 1.333, true)

}
