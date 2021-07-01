package main

import "fmt"

type Point3D struct {
	x, y, z float64
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

func main() {
	var A, B Point3D
	A.SetX(1)
	A.SetY(1)
	A.SetZ(1)
	B.SetX(1)
	B.SetY(1)
	B.SetZ(1)
	fmt.Println(pointsEquals(A, B))
	fmt.Println(A.X())
}
