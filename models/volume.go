package models

type Volume struct {
	Slices []Slice
	VX     int
	VY     int
	VZ     int
	Min    int
	Max    int
}

func (v *Volume) AddSlice(s Slice) {
	v.Slices = append(v.Slices, s)
}

func (v Volume) GetMin() int {
	if len(v.Slices) == 0 {
		return 0
	}
	firstSlice := v.Slices[0]
	if len(firstSlice.Points) == 0 {
		return 0
	}
	firstPoint := v.Slices[0].Points[0]
	result := firstPoint.Value
	for i := range v.Slices {
		sli := v.Slices[i]
		for j := range sli.Points {
			ps := sli.Points[j]
			if result > ps.Value {
				result = ps.Value
			}
		}
	}
	return result
}

func (v Volume) GetMax() int {
	if len(v.Slices) == 0 {
		return 0
	}
	firstSlice := v.Slices[0]
	if len(firstSlice.Points) == 0 {
		return 0
	}
	firstPoint := v.Slices[0].Points[0]
	result := firstPoint.Value
	for i := range v.Slices {
		sli := v.Slices[i]
		for j := range sli.Points {
			ps := sli.Points[j]
			if result < ps.Value {
				result = ps.Value
			}
		}
	}
	return result
}
