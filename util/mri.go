package util

import "github.com/grayzone/three/models"
import "io/ioutil"

//import "github.com/astaxie/beego"

func LoadMRIRaw(filepath string, x, y, z int) models.Volume {
	var result models.Volume
	result.VX = x
	result.VY = y
	result.VZ = z

	b, _ := ioutil.ReadFile(filepath)
	index := 0
	for i := 0; i < z; i++ {
		var s models.Slice
		s.Thickness = 1
		s.Index = i
		s.Width = x
		s.Height = y
		for j := 0; j < y; j++ {
			for k := 0; k < x; k++ {
				s.AddPoint(k, j, int(b[index]))
				index++
			}
		}
		result.AddSlice(s)
	}
	//	result.Max = result.GetMax()
	//	result.Min = result.GetMin()

	//	beego.Info(result.Max, result.Min, result.VX, result.VY, result.VZ)

	return result
}
