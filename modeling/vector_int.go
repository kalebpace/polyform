package modeling

import (
	"fmt"
	"math"

	"github.com/EliCDavis/vector"
)

type VectorInt struct {
	X int
	Y int
	Z int
}

func Vector3ToInt(v vector.Vector3, power int) VectorInt {
	newPower := math.Pow10(power)
	return VectorInt{
		X: int(math.Round(v.X() * newPower)),
		Y: int(math.Round(v.Y() * newPower)),
		Z: int(math.Round(v.Z() * newPower)),
	}
}

func (v VectorInt) Sub(other VectorInt) VectorInt {
	return VectorInt{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v VectorInt) ToRegularVector() vector.Vector3 {
	return vector.NewVector3(
		float64(v.X),
		float64(v.Y),
		float64(v.Z),
	)
}

func (v VectorInt) String() string {
	return fmt.Sprintf("%d, %d, %d", v.X, v.Y, v.Z)
}