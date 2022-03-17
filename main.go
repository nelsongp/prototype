package main

import (
	"fmt"
	"github.com/nelsongp/prototype/prototype"
)

func main() {
	color := "rojo"
	phones := map[string]string{"home": "123456", "work": "789456"}
	p1 := prototype.Prototype{"Alexys", 39, []string{"Deibis", "Carol"}, &color, phones}
	// copia del objeto original
	p2 := p1.Clone()

	// cambio de los valores al objeto p2
	p2.Age = 40
	p2.Name = "Carol"
	p2.Friends[0] = "Maria"
	p2.Friends[1] = "Pedro"

	// Estos cambios solo afectan al objeto p1
	color = "azul"
	phones["home"] = "147258"

	// muestra la información
	fmt.Println(p1.String())
	fmt.Println(p2.String())
}
