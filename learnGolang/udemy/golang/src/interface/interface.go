package main

import "fmt"

type Notebook struct {
	content []byte
}

func (nb *Notebook) Write(p []byte) (n int, err error) {
	nb.content = append(nb.content, p...)
	return len(p), nil
}

func (nb *Notebook) String() string {
	return string(nb.content)
}

// satisfaction

type Sword struct {
	name string
}

func (s Sword) detail() string {
	return "Super cold  Iced sword "
}
func (s Sword) cost() int {
	return 10
}

type Bow struct {
	name string
}

func (b Bow) detail() string {
	return "Magic Fire bow "
}
func (b Bow) cost() int {
	return 10
}

type Item interface {
	cost() int
}

type Weapon interface {
	detail() string
	Item
}

func attack(w Weapon) {

	fmt.Printf("attack by : %s.Item cost %d\n", w.detail(), w.cost())

}
func main() {
	nb := Notebook{}
	fmt.Println("Before :", nb)
	fmt.Println("After :", &nb, "Hello")
	fmt.Println("-------------")

	sword := Sword{name: "Icy-sword"}
	bow := Bow{name: "Fire-Bow"}
	var w Weapon
	w = sword

	attack(sword)
	attack(bow)
	attack(w)
}
