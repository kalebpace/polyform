package gausops

import (
	"bufio"
	"bytes"

	"github.com/EliCDavis/polyform/formats/ply"
	"github.com/EliCDavis/polyform/formats/spz"
	"github.com/EliCDavis/polyform/modeling"
	"github.com/EliCDavis/polyform/nodes"
)

type LoaderNode = nodes.Struct[modeling.Mesh, LoaderNodeData]

type LoaderNodeData struct {
	Data nodes.NodeOutput[[]byte]
}

func (pn LoaderNodeData) Process() (modeling.Mesh, error) {
	bufReader := bufio.NewReader(bytes.NewReader(pn.Data.Value()))

	splat, err := ply.ReadMesh(bufReader)
	if err != nil {
		return modeling.EmptyPointcloud(), err
	}

	return *splat, nil

	// header, err := ply.ReadHeader(bufReader)
	// if err != nil {
	// 	return modeling.EmptyPointcloud(), err
	// }

	// reader := header.BuildReader(bufReader)
	// plyMesh, err := reader.ReadMesh(ply.GuassianSplatVertexAttributesNoHarmonics)
	// if err != nil {
	// 	return modeling.EmptyPointcloud(), err
	// }
	// return *plyMesh, err
}

type SpzLoaderNode = nodes.Struct[modeling.Mesh, SpzLoaderNodeData]

type SpzLoaderNodeData struct {
	Data nodes.NodeOutput[[]byte]
}

func (pn SpzLoaderNodeData) Process() (modeling.Mesh, error) {
	bufReader := bufio.NewReader(bytes.NewReader(pn.Data.Value()))

	header, err := spz.Read(bufReader)
	if err != nil {
		// panic(err)
		return modeling.EmptyPointcloud(), err
	}

	return header.Mesh, err
}
