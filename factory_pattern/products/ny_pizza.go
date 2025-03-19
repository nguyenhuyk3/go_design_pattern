package products

import "fmt"

/*
* This struct will implement the pizza interface in interfaces folder
 */
type NYStyleCheesePizza struct {
	Name string
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{Name: "New York Style Cheese Pizza"}
}

func (p *NYStyleCheesePizza) Prepare() {
	fmt.Println("Chuẩn bị:", p.Name)
}

func (p *NYStyleCheesePizza) Bake() {
	fmt.Println("Nướng:", p.Name)
}

func (p *NYStyleCheesePizza) Cut() {
	fmt.Println("Cắt pizza theo lát mỏng:", p.Name)
}

func (p *NYStyleCheesePizza) Box() {
	fmt.Println("Đóng hộp:", p.Name)
}

func (p *NYStyleCheesePizza) GetName() string {
	return p.Name
}
