package models

type Slice struct {
	Points    []Point
	Thickness int
	Index     int
	Width     int
	Height    int
}

func (s *Slice) AddPoint(x, y, value int) {
	var p Point
	p.X = x
	p.Y = y
	p.Value = value
	s.Points = append(s.Points, p)
}
