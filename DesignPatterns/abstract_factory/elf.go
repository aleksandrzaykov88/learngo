package main

type elf struct {
}

func (e *elf) makeSword() iSword {
	return &elfSword{
		sword: sword{
			mark:   "ELF SMITH",
			damage: 10,
		},
	}
}

func (e *elf) makeShield() iShield {
	return &elfShield{
		shield: shield{
			mark:  "ELF SMITH",
			armor: 5,
		},
	}
}
