package cfdporomedia

import (
	"math"
	"strconv"
)

//Naar and Wygal's model
func WaterRemPerm(Sw float64, Swc float64) float64 {
	Krw := math.Pow((Sw-Swc)/(1-Swc), 4)
	return Krw
}

//Naar and Wygal's model
func OilRemPerm(Sw float64, Sg float64, So float64, Swc float64, Sor float64) float64 {
	So3 := math.Pow(So, 3)
	Kro := So3 * (1 - Sg + 2*Sw - 3*Swc) / math.Pow((1-Swc), 4)
	if So <= Sor {
		Kro = 0
	}
	return Kro
}

//Naar and Wygal's model
func GasRemPerm(Sw float64, Sg float64, Swc float64, Sgc float64) float64 {
	Sg3 := math.Pow(Sg, 3)
	Krg := Sg3 * (2 - Sg - 2*Swc) / math.Pow((1-Swc), 4)
	if Sg <= Sgc {
		Krg = 0
	}
	return Krg
}

type SlightlyCompressibleRelPerm struct {
	CartGrid CartGrid
}

func NewSlightlyCompressibleRelPerm(CartGrid CartGrid) SlightlyCompressibleRelPerm {
	return SlightlyCompressibleRelPerm{
		CartGrid: CartGrid,
	}
}

func (impl *SlightlyCompressibleRelPerm) SetSlightlyCompressibleRelPerm() {

	counter := -1
	i := 0
	j := 0
	k := 0
	for k = 0; k < impl.CartGrid.Nz; k++ {
		for j = 0; j < impl.CartGrid.Ny; j++ {
			for i = 0; i < impl.CartGrid.Nx; i++ {
				counter++
				index := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k)
				item := impl.CartGrid.Blocks[index]
				item.FluidData.RelPermData.Kro_old = 1
				item.FluidData.RelPermData.Kro_new = 1
				item.FluidData.RelPermData.Kro_itr = 1
				impl.CartGrid.Blocks[index] = item
			}
		}
	}
}
