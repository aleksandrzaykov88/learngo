package main

type iSword interface {
	setMark(mark string)
	setDamage(damage int)
	getMark() string
	getDamage() int
}

type sword struct {
	mark   string
	damage int
}

func (s *sword) setMark(mark string) {
	s.mark = mark
}

func (s *sword) getMark() string {
	return s.mark
}

func (s *sword) setDamage(damage int) {
	s.damage = damage
}

func (s *sword) getDamage() int {
	return s.damage
}
