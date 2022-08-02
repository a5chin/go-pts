package pts_test

import (
	"pts"
	"testing"
)

var (
	empty pts.Points = pts.Points{}
	path  string     = "assets/sample.pts"
)

func TestReadPts(t *testing.T) {
	points := pts.ReadPTS(path)

	if &points == &empty {
		t.Errorf("Cannot read %v.", path)
	}
}

func TestNorm(t *testing.T) {
	points := pts.ReadPTS(path)
	var x_normed, _ interface{} = points.Array[0].Norm()

	if _, x_ok := x_normed.(float64); !x_ok {
		t.Errorf("Cannot calclate Norm.")
	}
}

func TestMean(t *testing.T) {
	points := pts.ReadPTS(path)
	var x_mean, _ interface{} = points.Mean()

	if _, x_ok := x_mean.(float64); !x_ok {
		t.Errorf("Cannot calclate Mean.")
	}
}

func TestVar(t *testing.T) {
	points := pts.ReadPTS(path)
	var x_var, _ interface{} = points.Var()

	if _, x_ok := x_var.(float64); !x_ok {
		t.Errorf("Cannot calclate Var.")
	}
}

func TestStd(t *testing.T) {
	points := pts.ReadPTS(path)
	var x_var, _ interface{} = points.Std()

	if _, x_ok := x_var.(float64); !x_ok {
		t.Errorf("Cannot calclate Std.")
	}
}
