package main

import (
	"fmt"

	"github.com/aleksandrzaykov88/learngo/SimpleJavaTasks/math/vector3Dn"
)

type Vector3DProcessor struct {
	v, w vector3Dn.Vector3D
}

func (p *Vector3DProcessor) SumVectors() vector3Dn.Vector3D {
	return vector3D.NewVector3DCoords(p.v.X()+p.w.X(), p.v.Y()+p.w.Y(), p.v.Z()+p.w.Z())
}

func main() {
	v := vector3Dn.NewVector3DCoords(1, 3, 7)
	w := vector3Dn.NewVector3DCoords(33, 11, 27)
	vp := Vector3DProcessor{v, w}
	fmt.Println(vp.SumVectors())
}
