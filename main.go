package main

import (
	"fmt"
	"root/module/warriors"
)

func main() {
	soldier1 := warriors.New("Marko", 0)
	soldier2 := warriors.New("Marija", 5)

	warriors.Damage(soldier1, soldier2, "heavy attack")

	soldier2.Painkillers()
	soldier2.Painkillers()

	warriors.Damage(soldier2, soldier1, "lite attack")

	fmt.Printf("Marko ima %d zdravlja, a Marija ima %d zdravnja", soldier1.GetHelth(), soldier2.GetHelth())
}
