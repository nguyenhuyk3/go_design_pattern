package products

import "fmt"

type ChicagoStyleCheesePizza struct {
	Name string
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{Name: "Chicago Style Cheese Pizza"}
}

func (p *ChicagoStyleCheesePizza) Prepare() {
	fmt.Println("Chuẩn bị:", p.Name)
}

func (p *ChicagoStyleCheesePizza) Bake() {
	fmt.Println("Nướng:", p.Name)
}

func (p *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cắt pizza theo lát vuông:", p.Name)
}

func (p *ChicagoStyleCheesePizza) Box() {
	fmt.Println("Đóng hộp:", p.Name)
}

func (p *ChicagoStyleCheesePizza) GetName() string {
	return p.Name
}
