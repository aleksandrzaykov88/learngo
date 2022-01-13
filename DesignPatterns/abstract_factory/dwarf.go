package main

type dwarf struct {
}

func (d *dwarf) makeSword() iSword {
	return &dwarfSword{
		sword: sword{
			mark:   "dwarf sword",
			damage: 5,
		},
	}
}

func (d *dwarf) makeShield() iShield {
	return &dwarfShield{
		shield: shield{
			mark:  "dwarf shield",
			armor: 10,
		},
	}
}
