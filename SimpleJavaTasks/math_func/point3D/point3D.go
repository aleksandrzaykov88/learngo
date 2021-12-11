//point3D package describes point in 3D.
//It also provides some methods to work with it.
package main

import "fmt"

//Point3D is a point in 3D.
type Point3D struct {
	x, y, z float64
}

//NewPoint3D construts the point.
func NewPoint3D() *Point3D {
	return &Point3D{}
}

func (p *Point3D) SetX(x float64) {
	p.x = x
}

func (p *Point3D) SetY(y float64) {
	p.y = y
}

func (p *Point3D) SetZ(z float64) {
	p.z = z
}

func (p *Point3D) X() float64 {
	return p.x
}

func (p *Point3D) Y() float64 {
	return p.y
}

func (p *Point3D) Z() float64 {
	return p.z
}

//PrintPoint outputs point coordinates.
func (p *Point3D) PrintPoint() {
	fmt.Printf("Point coordinates: X=%.02f Y=%.02f Z=%.02f", p.x, p.y, p.z)
}

//pointsEquals returns true if coordinates of two input points are equal.
func pointsEquals(A, B Point3D) bool {
	if A.x == B.x && A.y == B.y && A.z == B.z {
		return true
	}
	return false
}
