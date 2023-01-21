package marching_test

import (
	"testing"

	"github.com/EliCDavis/polyform/modeling"
	"github.com/EliCDavis/polyform/modeling/marching"
	"github.com/EliCDavis/vector"
)

var meshResult modeling.Mesh

func BenchmarkSphere(b *testing.B) {
	cubesPerUnit := 10.
	var mesh modeling.Mesh
	for n := 0; n < b.N; n++ {
		canvas := marching.NewMarchingCanvas(cubesPerUnit)

		canvas.AddField(
			marching.Sphere(vector.Vector3Zero(), 2., 1.),
		)

		mesh = canvas.March(0)
	}

	meshResult = mesh
}