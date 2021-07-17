package main

import (
	"fmt"

	"github.com/aleksandrzaykov88/learngo/SimpleJavaTasks/point3D"
)

type Vector3D struct {
	x, y, z float64
}

func NewVector3D() *Vector3D {
	return &Vector3D{}
}

func NewVector3DCoords(x, y, z float64) *Vector3D {
	return &Vector3D{x, y, z}
}

func NewVector3DPoints(a, b point3D.Point3D) *Vector3D {
	return &Vector3D{b.X() - a.X(), b.Y() - a.Y(), a.Z() - b.Z()}
}

func main() {
	v := Vector3D{1, 2, 3}
	fmt.Println(v)
}
