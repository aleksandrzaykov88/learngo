package main

type iShield interface {
	setMark(mark string)
	setArmor(armor int)
	getMark() string
	getArmor() int
}

type shield struct {
	mark  string
	armor int
}

func (s *shield) setMark(mark string) {
	s.mark = mark
}

func (s *shield) getMark() string {
	return s.mark
}

func (s *shield) setArmor(armor int) {
	s.armor = armor
}

func (s *shield) getArmor() int {
	return s.armor
}
