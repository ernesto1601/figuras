package figuras

import "fmt"

type Cuadrado struct {
	Nombre string
	Lado   float32
}

func (c *Cuadrado) CalcularArea() float32 {
	return c.Lado * c.Lado
}
func (c *Cuadrado) CalcularPerimetro() float32 {
	return (2 * c.Lado) + (2 * c.Lado)
}
func (c *Cuadrado) ImprimeNombre() {
	fmt.Println(c.Nombre)
}
