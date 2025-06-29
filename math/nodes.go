package math

import (
	"math"

	"github.com/EliCDavis/polyform/generator"
	"github.com/EliCDavis/polyform/nodes"
	"github.com/EliCDavis/polyform/refutil"
	"github.com/EliCDavis/vector"
)

func init() {
	factory := &refutil.TypeFactory{}

	refutil.RegisterType[Round](factory)

	refutil.RegisterType[nodes.Struct[CircumferenceNode]](factory)

	refutil.RegisterType[nodes.Struct[DifferenceNodeData[int]]](factory)
	refutil.RegisterType[nodes.Struct[DifferenceNodeData[float64]]](factory)
	refutil.RegisterType[nodes.Struct[DifferencesToArrayNodeData[int]]](factory)
	refutil.RegisterType[nodes.Struct[DifferencesToArrayNodeData[float64]]](factory)

	refutil.RegisterType[nodes.Struct[SumNodeData[int]]](factory)
	refutil.RegisterType[nodes.Struct[SumNodeData[float64]]](factory)

	refutil.RegisterType[nodes.Struct[SumArraysNodeData[int]]](factory)
	refutil.RegisterType[nodes.Struct[SumArraysNodeData[float64]]](factory)

	refutil.RegisterType[nodes.Struct[AddToArrayNodeData[int]]](factory)
	refutil.RegisterType[nodes.Struct[AddToArrayNodeData[float64]]](factory)

	refutil.RegisterType[nodes.Struct[DivideNodeData[int]]](factory)
	refutil.RegisterType[nodes.Struct[DivideNodeData[float64]]](factory)

	refutil.RegisterType[nodes.Struct[MultiplyNodeData[float64]]](factory)
	refutil.RegisterType[nodes.Struct[MultiplyNodeData[int]]](factory)
	refutil.RegisterType[nodes.Struct[MultiplyToArrayNodeData[float64]]](factory)
	refutil.RegisterType[nodes.Struct[MultiplyToArrayNodeData[int]]](factory)

	refutil.RegisterType[nodes.Struct[InverseNodeData[float64]]](factory)
	refutil.RegisterType[nodes.Struct[InverseNodeData[int]]](factory)

	refutil.RegisterType[nodes.Struct[NegateNode[int]]](factory)
	refutil.RegisterType[nodes.Struct[NegateNode[float64]]](factory)

	refutil.RegisterType[nodes.Struct[DoubleNode[int]]](factory)
	refutil.RegisterType[nodes.Struct[DoubleNode[float64]]](factory)

	refutil.RegisterType[nodes.Struct[HalfNode[int]]](factory)
	refutil.RegisterType[nodes.Struct[HalfNode[float64]]](factory)

	refutil.RegisterType[nodes.Struct[OneNode]](factory)
	refutil.RegisterType[nodes.Struct[ZeroNode]](factory)

	refutil.RegisterType[nodes.Struct[MinNode[int]]](factory)
	refutil.RegisterType[nodes.Struct[MinNode[float64]]](factory)
	refutil.RegisterType[nodes.Struct[MinArrayNode[int]]](factory)
	refutil.RegisterType[nodes.Struct[MinArrayNode[float64]]](factory)
	refutil.RegisterType[nodes.Struct[MaxNode[int]]](factory)
	refutil.RegisterType[nodes.Struct[MaxNode[float64]]](factory)
	refutil.RegisterType[nodes.Struct[MaxArrayNode[int]]](factory)
	refutil.RegisterType[nodes.Struct[MaxArrayNode[float64]]](factory)

	refutil.RegisterType[nodes.Struct[IntToFloatNode]](factory)
	refutil.RegisterType[nodes.Struct[FloatToIntNode]](factory)

	generator.RegisterTypes(factory)
}

// ============================================================================

type OneNode struct{}

func (cn OneNode) Int() nodes.StructOutput[int] {
	return nodes.NewStructOutput(1)
}

func (cn OneNode) Float64() nodes.StructOutput[float64] {
	return nodes.NewStructOutput(1.)
}

func (cn OneNode) Description() string {
	return "Just the number 1"
}

// ============================================================================

type ZeroNode struct{}

func (cn ZeroNode) Int() nodes.StructOutput[int] {
	return nodes.NewStructOutput(0)
}

func (cn ZeroNode) Float64() nodes.StructOutput[float64] {
	return nodes.NewStructOutput(0.)
}

func (cn ZeroNode) Description() string {
	return "Just the number 0"
}

// ============================================================================

type DoubleNode[T vector.Number] struct {
	In nodes.Output[T] `description:"The number to double"`
}

func (cn DoubleNode[T]) Int() nodes.StructOutput[int] {
	return nodes.NewStructOutput(int(nodes.TryGetOutputValue(cn.In, 0)) * 2)
}

func (cn DoubleNode[T]) Float64() nodes.StructOutput[float64] {
	return nodes.NewStructOutput(float64(nodes.TryGetOutputValue(cn.In, 0)) * 2)
}

func (cn DoubleNode[T]) Description() string {
	return "Doubles the number provided"
}

// ============================================================================

type HalfNode[T vector.Number] struct {
	In nodes.Output[T] `description:"The number to halve"`
}

func (cn HalfNode[T]) Int() nodes.StructOutput[int] {
	return nodes.NewStructOutput(int(float64(nodes.TryGetOutputValue(cn.In, 0)) * 0.5))
}

func (cn HalfNode[T]) Float64() nodes.StructOutput[float64] {
	return nodes.NewStructOutput(float64(nodes.TryGetOutputValue(cn.In, 0)) * 0.5)
}

func (cn HalfNode[T]) Description() string {
	return "Divides the number in half"
}

// ============================================================================

type IntToFloatNode struct {
	In nodes.Output[int]
}

func (cn IntToFloatNode) Out() nodes.StructOutput[float64] {
	return nodes.NewStructOutput(float64(nodes.TryGetOutputValue(cn.In, 0)))
}

type FloatToIntNode struct {
	In nodes.Output[float64]
}

func (cn FloatToIntNode) Out() nodes.StructOutput[int] {
	return nodes.NewStructOutput(int(nodes.TryGetOutputValue(cn.In, 0)))
}

// ============================================================================

type NegateNode[T vector.Number] struct {
	In nodes.Output[T] `description:"The number to take the additive inverse of"`
}

func (cn NegateNode[T]) Out() nodes.StructOutput[T] {
	return nodes.NewStructOutput(nodes.TryGetOutputValue(cn.In, 0) * -1)
}

func (cn NegateNode[T]) Description() string {
	return "The additive inverse of an element x, denoted −x, is the element that when added to x, yields the additive identity, 0"
}

// ============================================================================
type InverseNodeData[T vector.Number] struct {
	In nodes.Output[T] `description:"The number to take the inverse of"`
}

func (cn InverseNodeData[T]) Additive() nodes.StructOutput[T] {
	return nodes.NewStructOutput(nodes.TryGetOutputValue(cn.In, 0) * -1)
}

func (cn InverseNodeData[T]) AdditiveDescription() string {
	return "The additive inverse of an element x, denoted −x, is the element that when added to x, yields the additive identity, 0"
}

func (cn InverseNodeData[T]) Multiplicative() nodes.StructOutput[T] {
	v := nodes.TryGetOutputValue(cn.In, 0)
	if v == 0 {
		var t T
		return nodes.NewStructOutput(t)
	}
	return nodes.NewStructOutput(1. / v)
}

func (cn InverseNodeData[T]) MultiplicativeDescription() string {
	return "The multiplicative inverse for a number x, denoted by 1/x or x^−1, is a number which when multiplied by x yields the multiplicative identity, 1"
}

// ============================================================================

type DivideNodeData[T vector.Number] struct {
	Dividend nodes.Output[T]
	Divisor  nodes.Output[T]
}

func (DivideNodeData[T]) Description() string {
	return "Dividend / Divisor"
}

func (cn DivideNodeData[T]) Out() nodes.StructOutput[T] {
	var empty T
	a := nodes.TryGetOutputValue(cn.Dividend, empty)
	b := nodes.TryGetOutputValue(cn.Divisor, empty)

	// TODO: Eeeeehhhhhhhhhhhhhhhhhhhhh
	if b == 0 {
		return nodes.NewStructOutput(empty)
	}

	return nodes.NewStructOutput(a / b)
}

// ============================================================================

type Round = nodes.Struct[RoundNodeData]

type RoundNodeData struct {
	In nodes.Output[float64]
}

func (cn RoundNodeData) Int() nodes.StructOutput[int] {
	if cn.In == nil {
		return nodes.NewStructOutput(0)
	}
	return nodes.NewStructOutput(int(math.Round(cn.In.Value())))
}

func (cn RoundNodeData) Float() nodes.StructOutput[float64] {
	if cn.In == nil {
		return nodes.NewStructOutput(0.)
	}
	return nodes.NewStructOutput(math.Round(cn.In.Value()))
}

// ============================================================================

type CircumferenceNode struct {
	Radius nodes.Output[float64]
}

func (cn CircumferenceNode) Description() string {
	return "Circumference of a circle"
}

func (cn CircumferenceNode) Int() nodes.StructOutput[int] {
	if cn.Radius == nil {
		return nodes.NewStructOutput(0)
	}
	return nodes.NewStructOutput(int(math.Round(cn.Radius.Value() * 2 * math.Pi)))
}

func (cn CircumferenceNode) Float() nodes.StructOutput[float64] {
	if cn.Radius == nil {
		return nodes.NewStructOutput(0.)
	}
	return nodes.NewStructOutput(math.Round(cn.Radius.Value() * 2 * math.Pi))
}
