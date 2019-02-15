package composite

import "fmt"

type MenuComponent interface {
	Name() string
	Description() string
	Price() float32
	Print()

	Add(MenuComponent)
	Remove(int)
	Child(int) MenuComponent
}

type MenuItem struct {
	name        string
	description string
	price       float32
}

func NewMenuItem(name, description string, price float32) MenuComponent {
	return &MenuItem{
		name:        name,
		description: description,
		price:       price,
	}
}

func (m *MenuItem) Name() string {
	return m.name
}

func (m *MenuItem) Description() string {
	return m.description
}

func (m *MenuItem) Price() float32 {
	return m.price
}

func (m *MenuItem) Print() {
	fmt.Printf("  %s, $%.2f\n", m.name, m.price)
	fmt.Printf("    -- %s\n", m.description)
}

func (m *MenuItem) Add(MenuComponent) {
	panic("not implement")
}

func (m *MenuItem) Remove(int) {
	panic("not implement")
}

func (m *MenuItem) Child(int) MenuComponent {
	panic("not implement")
}

type Menu struct {
	name        string
	description string
	children    []MenuComponent
}

func NewMenu(name, description string) MenuComponent {
	return &Menu{
		name:        name,
		description: description,
	}
}

func (m *Menu) Name() string {
	return m.name
}

func (m *Menu) Description() string {
	return m.description
}

func (m *Menu) Price() (price float32) {
	for _, v := range m.children {
		price += v.Price()
	}
	return
}

func (m *Menu) Print() {
	fmt.Printf("%s, %s, $%.2f\n", m.name, m.description, m.Price())
	fmt.Println("------------------------")
	for _, v := range m.children {
		v.Print()
	}
	fmt.Println()
}

func (m *Menu) Add(c MenuComponent) {
	m.children = append(m.children, c)
}

func (m *Menu) Remove(idx int) {
	m.children = append(m.children[:idx], m.children[idx+1:]...)
}

func (m *Menu) Child(idx int) MenuComponent {
	return m.children[idx]
}
