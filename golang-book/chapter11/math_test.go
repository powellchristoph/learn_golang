package math

import "testing"

type testdata struct {
    values []float64
    average float64
    min float64
    max float64
}

var tests = []testdata{
    { []float64{1, 2}, 1.5, 1, 2},
    { []float64{1, 1, 1, 1, 1, 1}, 1, 1, 1},
    { []float64{-1, 1}, 0, -1, 1},
    { []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5.5, 1, 10},
}

func TestAverage(t *testing.T) {
    for _, data := range tests {
        v := Average(data.values)
        if v != data.average {
            t.Error(
                "For", data.values,
                "expected", data.average,
                "got", v,
            )
        }
    }
}

func TestMin(t *testing.T) {
    for _, data := range tests {
        v := Min(data.values)
        if v != data.min {
            t.Error(
                "For", data.values,
                "expected", data.min,
                "got", v,
            )
        }
    }
}

func TestMax(t *testing.T) {
    for _, data := range tests {
        v := Max(data.values)
        if v != data.max {
            t.Error(
                "For", data.values,
                "expected", data.max,
                "got", v,
            )
        }
    }
}
//func TestAverage(t *testing.T) {
//    var v float64
//    v = Average([]float64{1, 2})
//    if v != 1.5 {
//        t.Error("Expected 1.5, got ", v)
//    }
//}
