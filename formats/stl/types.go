package stl

import (
	"bytes"
	"io"

	"github.com/EliCDavis/polyform/generator"
	"github.com/EliCDavis/polyform/generator/artifact"
	"github.com/EliCDavis/polyform/modeling"
	"github.com/EliCDavis/polyform/nodes"
	"github.com/EliCDavis/polyform/refutil"
)

func init() {
	factory := &refutil.TypeFactory{}
	refutil.RegisterType[ReadNode](factory)
	refutil.RegisterType[ArtifactNode](factory)
	generator.RegisterTypes(factory)
}

type ReadNode = nodes.Struct[ReadNodeData]

type ReadNodeData struct {
	Data nodes.Output[[]byte]
}

func (gad ReadNodeData) Out() nodes.StructOutput[modeling.Mesh] {
	if gad.Data == nil {
		return nodes.NewStructOutput(modeling.EmptyMesh(modeling.TriangleTopology))
	}

	data := gad.Data.Value()
	if len(data) == 0 {
		return nodes.NewStructOutput(modeling.EmptyMesh(modeling.TriangleTopology))
	}

	cloud, err := ReadMesh(bytes.NewReader(data))
	if err != nil {
		out := nodes.NewStructOutput(modeling.EmptyMesh(modeling.TriangleTopology))
		out.LogError(err)
		return out
	}

	return nodes.NewStructOutput(*cloud)
}

// ============================================================================

type Artifact struct {
	Mesh modeling.Mesh
}

func (sa Artifact) Write(w io.Writer) error {
	return WriteMesh(w, sa.Mesh)
}

func (Artifact) Mime() string {
	return "application/sla"
}

type ArtifactNode = nodes.Struct[ArtifactNodeData]

type ArtifactNodeData struct {
	In nodes.Output[modeling.Mesh]
}

func (pn ArtifactNodeData) Out() nodes.StructOutput[artifact.Artifact] {
	return nodes.NewStructOutput[artifact.Artifact](Artifact{Mesh: pn.In.Value()})
}
