package main

import (
	"fmt"

	"github.com/aleksandrzaykov88/learngo/SimpleJavaTasks/math_func/vector3D"
)

type Vector3DProcessor struct {
	v, w vector3D.Vector3D
}

func (p *Vector3DProcessor) SumVectors() vector3D.Vector3D {
	return *vector3D.NewVector3DCoords(p.v.X()+p.w.X(), p.v.Y()+p.w.Y(), p.v.Z()+p.w.Z())
}

func main() {
	v := vector3D.NewVector3DCoords(1, 3, 7)
	w := vector3D.NewVector3DCoords(33, 11, 27)
	vp := Vector3DProcessor{*v, *w}
	fmt.Println(vp.SumVectors())
}
