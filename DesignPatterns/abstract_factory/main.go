package main

import "fmt"

func main() {
	elfFactory, _ := getEquipmentFactory("elf")
	dwarfFactory, _ := getEquipmentFactory("dwarf")

	dwarfSword := dwarfFactory.makeSword()
	dwarfShield := dwarfFactory.makeShield()

	elfSword := elfFactory.makeSword()
	elfShield := elfFactory.makeShield()

	printSwordDetails(dwarfSword)
	printShieldDetails(dwarfShield)

	printSwordDetails(elfSword)
	printShieldDetails(elfShield)
}

func printSwordDetails(s iSword) {
	fmt.Printf("Mark: %s", s.getMark())
	fmt.Println()
	fmt.Printf("Damage: %d", s.getDamage())
	fmt.Println()
}

func printShieldDetails(s iShield) {
	fmt.Printf("Mark: %s", s.getMark())
	fmt.Println()
	fmt.Printf("Armor: %d", s.getArmor())
	fmt.Println()
}
