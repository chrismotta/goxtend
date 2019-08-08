package muebles

import "fmt"

// Sillas es una silla
type Sillas struct {
	Name  string
	Price float32
	Tela  string
}

// Mesas es una mesa
type Mesas struct {
	Name   string
	Price  float32
	Madera string
}

// Metodos es una interface de Sillas y Mesas
type Metodos interface {
	GetPrice()
}

// CallGetPrice llama a getPrice
func CallGetPrice(element Metodos) {
	element.GetPrice()
}

// SetMaterial setea el material
func (element *Sillas) SetMaterial(material string) {
	element.Tela = material
}

// GetPrice devuelve el precio
func (element Sillas) GetPrice() {
	fmt.Printf("El precio es %f!!\n", element.Price)
}

// GetPrice devuelve el precio
func (element Mesas) GetPrice() {
	fmt.Printf("El precio es %f!!\n", element.Price)
}
