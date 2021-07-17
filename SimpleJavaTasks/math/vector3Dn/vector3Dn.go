//vector3D package describes vector in 3D.
//It also provides some methods to work with it.
package vector3D

import (
	"math"

	"github.com/aleksandrzaykov88/learngo/SimpleJavaTasks/math/point3D"
)

//Vector3D describes a vector in 3D space.
type Vector3D struct {
	x, y, z float64
}

//NewVector3D constructs the null-vector.
func NewVector3D() *Vector3D {
	return &Vector3D{}
}

//NewVector3DCoords constructs the vector by its coordinates.
func NewVector3DCoords(x, y, z float64) *Vector3D {
	return &Vector3D{x, y, z}
}

//NewVector3DPoints constructs the vector by two 3D points.
//Vector start and end, respectivelyю
func NewVector3DPoints(a, b point3D.Point3D) *Vector3D {
	return &Vector3D{b.X() - a.X(), b.Y() - a.Y(), b.Z() - a.Z()}
}

func (v *Vector3D) SetX(x float64) {
	v.x = x
}

func (v *Vector3D) SetY(y float64) {
	v.y = y
}

func (v *Vector3D) SetZ(z float64) {
	v.z = z
}

func (v *Vector3D) X() float64 {
	return v.x
}

func (v *Vector3D) Y() float64 {
	return v.y
}

func (v *Vector3D) Z() float64 {
	return v.z
}

//GetLength returns length of a vector.
func (v *Vector3D) GetLength() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2))
}

//VectorsEqual returns true if two vectors are equal.
func (v *Vector3D) VectorsEqual(i Vector3D) bool {
	if i.x == v.x && i.y == v.y && i.z == v.z {
		return true
	}
	return false
}
