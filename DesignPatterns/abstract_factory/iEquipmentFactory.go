package main

import "fmt"

type iEquipmentFactory interface {
	makeSword() iSword
	makeShield() iShield
}

func getEquipmentFactory(smith string) (iEquipmentFactory, error) {
	if smith == "elf" {
		return &elf{}, nil
	}

	if smith == "dwarf" {
		return &dwarf{}, nil
	}

	return nil, fmt.Errorf("There is no such smith! Wraaa!")
}
