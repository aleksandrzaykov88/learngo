package main

import (
	"fmt"

	"github.com/aleksandrzaykov88/learngo/SimpleJavaTasks/math_func/vector3D"
)

//Vector3DProcessor is a type whitch provides some popular operations with vectors.
type Vector3DProcessor struct {
	v, w vector3D.Vector3D
}

//IsCollinear returns true if two vectors are collinear.
func (p *Vector3DProcessor) IsCollinear() bool {
	if p.v.X()/p.w.X() == p.v.Y()/p.w.Y() && p.v.Z()/p.w.Z() == p.v.Y()/p.w.Y() && p.v.Z()/p.w.Z() == p.v.X()/p.w.X() {
		return true
	}
	return false
}

//VectorProduct returns vector product of two vectors.
func (p *Vector3DProcessor) VectorProduct() vector3D.Vector3D {
	return *vector3D.NewVector3DCoords(p.v.Y()*p.w.Z()-p.w.Y()*p.v.Z(), p.w.X()*p.v.Z()-p.v.X()*p.w.Z(), p.v.X()*p.w.Y()-p.w.X()*p.v.Y())
}

//ScalarProduct returns scalar product of two vectors.
func (p *Vector3DProcessor) ScalarProduct() float64 {
	return p.v.X()*p.w.X() + p.v.Y()*p.w.Y() + p.v.Z()*p.w.Z()
}

//DifVectors returns defference of two vectors.
func (p *Vector3DProcessor) DifVectors() vector3D.Vector3D {
	return *vector3D.NewVector3DCoords(p.v.X()-p.w.X(), p.v.Y()-p.w.Y(), p.v.Z()-p.w.Z())
}

//SumVectors returns sum of two vectors.
func (p *Vector3DProcessor) SumVectors() vector3D.Vector3D {
	return *vector3D.NewVector3DCoords(p.v.X()+p.w.X(), p.v.Y()+p.w.Y(), p.v.Z()+p.w.Z())
}

func main() {
	v := vector3D.NewVector3DCoords(2, 2, 2)
	w := vector3D.NewVector3DCoords(4, 4, 1)
	vp := Vector3DProcessor{*v, *w}
	fmt.Println(vp.IsCollinear())
}
