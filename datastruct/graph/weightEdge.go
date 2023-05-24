package graph

/*
带权边
*/
type WeightEdge struct {
	VertexI int
	VertexJ int
	Weight  int
}

func NewWeightEdge(out int, in int, weight int) WeightEdge {
	return WeightEdge{
		VertexI: out,
		VertexJ: in,
		Weight:  weight,
	}
}
