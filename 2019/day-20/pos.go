package main

type Pos struct {
	I, J int
}

func (p Pos) Up() Pos {
	return Pos{p.I-1, p.J}
}

func (p Pos) Down() Pos {
	return Pos{p.I+1, p.J}
}

func (p Pos) Left() Pos {
	return Pos{p.I, p.J-1}
}

func (p Pos) Right() Pos {
	return Pos{p.I, p.J+1}
}