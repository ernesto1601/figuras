package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	Nombre string
	Radio  float32
}

func (cir *Circulo) CalcularArea() float32 {
	return math.Pi * (cir.Radio * cir.Radio)
}
func (cir *Circulo) CalcularPerimetro() float32 {
	return (2 * math.Pi) * cir.Radio
}

func (cir *Circulo) ImprimeNombre() {
	fmt.Println(cir.Nombre)
}
