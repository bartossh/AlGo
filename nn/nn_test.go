package nn

import (
	"fmt"
	"math"
	"testing"

	"gotest.tools/assert"
)

func TestCopy(t *testing.T) {
	benchCases := [][]int{
		{10, 20},
		{100, 20},
		{100, 200},
		{1000, 200},
		{1000, 2000},
	}
	for _, bc := range benchCases {
		t.Run(fmt.Sprintf("copy matrix %v-%v", bc[0], bc[1]), func(t *testing.T) {
			n0 := NewMatrix(bc[0], bc[1])
			n0.Randomize()
			n1 := NewMatrix(bc[0], bc[1])
			Copy(n1, n0)
			for i := range n1.values {
				assert.Equal(t, n0.values[i], n1.values[i])
			}
		})
	}
}

func BenchmarkCopy(b *testing.B) {
	benchCases := [][]int{
		{10, 20},
		{100, 20},
		{100, 200},
		{1000, 200},
		{1000, 2000},
	}
	for _, bc := range benchCases {
		b.Run(fmt.Sprintf("copy matrix %v-%v", bc[0], bc[1]), func(b *testing.B) {
			n0 := NewMatrix(bc[0], bc[1])
			n0.Randomize()
			n1 := NewMatrix(bc[0], bc[1])
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				Copy(n1, n0)
			}
		})
	}
}

func TestSumMatrix(t *testing.T) {
	type testDef struct {
		a      Matrix
		b      Matrix
		result Matrix
	}

	testCases := []testDef{
		{
			a: Matrix{
				rows: 3,
				cols: 3,
				values: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
			},
			b: Matrix{
				rows: 3,
				cols: 3,
				values: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
			},
			result: Matrix{
				rows: 3,
				cols: 3,
				values: []float64{
					2, 4, 6,
					8, 10, 12,
					14, 16, 18,
				},
			},
		},
		{
			a: Matrix{
				rows: 2,
				cols: 3,
				values: []float64{
					1, 2, 3,
					4, 5, 6,
				},
			},
			b: Matrix{
				rows: 2,
				cols: 3,
				values: []float64{
					7, 8, 9,
					10, 11, 12,
				},
			},
			result: Matrix{
				rows: 2,
				cols: 3,
				values: []float64{
					8, 10, 12,
					14, 16, 18,
				},
			},
		},
		{
			a: Matrix{
				rows: 3,
				cols: 4,
				values: []float64{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 10, 11, 12,
				},
			},
			b: Matrix{
				rows: 3,
				cols: 4,
				values: []float64{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 10, 11, 12,
				},
			},
			result: Matrix{
				rows: 3,
				cols: 4,
				values: []float64{
					2, 4, 6, 8,
					10, 12, 14, 16,
					18, 20, 22, 24,
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %v", i), func(t *testing.T) {
			err := Sum(tc.a, tc.b)
			assert.NilError(t, err)
			for i := range tc.a.values {
				assert.Equal(t, true, tc.a.values[i] == tc.result.values[i])
			}

		})
	}

}

type matrixDef struct {
	rows, cols int
}
type testCase struct {
	dst, a, b matrixDef
}

var testCases = []testCase{
	{
		dst: matrixDef{5, 5},
		a:   matrixDef{5, 10},
		b:   matrixDef{10, 5},
	},
	{
		dst: matrixDef{50, 50},
		a:   matrixDef{50, 100},
		b:   matrixDef{100, 50},
	},
	{
		dst: matrixDef{50, 50},
		a:   matrixDef{50, 151},
		b:   matrixDef{151, 50},
	},
	{
		dst: matrixDef{100, 100},
		a:   matrixDef{100, 1000},
		b:   matrixDef{1000, 100},
	},
	{
		dst: matrixDef{1000, 1000},
		a:   matrixDef{1000, 500},
		b:   matrixDef{500, 1000},
	},
}

var testCasesFailure = []testCase{
	{
		dst: matrixDef{5, 6},
		a:   matrixDef{5, 10},
		b:   matrixDef{10, 5},
	},
	{
		dst: matrixDef{50, 50},
		a:   matrixDef{50, 50},
		b:   matrixDef{100, 50},
	},
	{
		dst: matrixDef{50, 50},
		a:   matrixDef{50, 151},
		b:   matrixDef{150, 50},
	},
	{
		dst: matrixDef{100, 100},
		a:   matrixDef{1000, 1000},
		b:   matrixDef{1000, 100},
	},
	{
		dst: matrixDef{1000, 1000},
		a:   matrixDef{1000, 2000},
		b:   matrixDef{1000, 1000},
	},
}

func TestMatrixDotProductSuccess(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("success dot for dst %v a %v b %v ", tc.dst, tc.a, tc.b), func(t *testing.T) {
			dst := NewMatrix(tc.dst.rows, tc.dst.cols)
			dst.Randomize()
			dstCopy := NewMatrix(tc.dst.rows, tc.dst.cols)
			err := Copy(dstCopy, dst)
			assert.NilError(t, err)
			a := NewMatrix(tc.a.rows, tc.a.cols)
			a.Randomize()
			b := NewMatrix(tc.b.rows, tc.b.cols)
			b.Randomize()
			err = Dot(dst, a, b)
			assert.NilError(t, err)
			for i := range dst.values {
				assert.Equal(t, false, dst.values[i] == dstCopy.values[i])
			}
		})
	}
}

func TestMatrixDotProductCorrectness(t *testing.T) {
	type testDef struct {
		a      Matrix
		b      Matrix
		result Matrix
	}

	testCases := []testDef{
		{
			a: Matrix{
				rows: 3,
				cols: 3,
				values: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
			},
			b: Matrix{
				rows: 3,
				cols: 2,
				values: []float64{
					1, 2,
					3, 4,
					5, 6,
				},
			},
			result: Matrix{
				rows: 3,
				cols: 2,
				values: []float64{
					22, 28,
					49, 64,
					76, 100,
				},
			},
		},
		{
			a: Matrix{
				rows: 2,
				cols: 3,
				values: []float64{
					1, 2, 3,
					4, 5, 6,
				},
			},
			b: Matrix{
				rows: 3,
				cols: 2,
				values: []float64{
					7, 8,
					9, 10,
					11, 12,
				},
			},
			result: Matrix{
				rows: 2,
				cols: 2,
				values: []float64{
					58, 64,
					139, 154,
				},
			},
		},
		{
			a: Matrix{
				rows: 3,
				cols: 4,
				values: []float64{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 10, 11, 12,
				},
			},
			b: Matrix{
				rows: 4,
				cols: 2,
				values: []float64{
					1, 2,
					3, 4,
					5, 6,
					7, 8,
				},
			},
			result: Matrix{
				rows: 2,
				cols: 2,
				values: []float64{
					50, 60,
					114, 140,
					178, 220,
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %v", i), func(t *testing.T) {
			dst := NewMatrix(tc.a.rows, tc.b.cols)
			err := Dot(dst, tc.a, tc.b)
			assert.NilError(t, err)
			for i := range dst.values {
				assert.Equal(t, true, dst.values[i] == tc.result.values[i])
			}

		})
	}

}

func TestMatrixDotProductFailure(t *testing.T) {
	for _, tc := range testCasesFailure {
		t.Run(fmt.Sprintf("failure dot for dst %v a %v b %v ", tc.dst, tc.a, tc.b), func(t *testing.T) {
			dst := NewMatrix(tc.dst.rows, tc.dst.cols)
			a := NewMatrix(tc.a.rows, tc.a.cols)
			a.Randomize()
			b := NewMatrix(tc.b.rows, tc.b.cols)
			b.Randomize()
			err := Dot(dst, a, b)
			assert.ErrorContains(t, err, "wrong")
		})
	}
}

func BenchmarkMatrixDotProduct(b *testing.B) {
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("dot dst %v a %v b %v ", tc.dst, tc.a, tc.b), func(b *testing.B) {
			dst := NewMatrix(tc.dst.rows, tc.dst.cols)
			a := NewMatrix(tc.a.rows, tc.a.cols)
			a.Randomize()
			c := NewMatrix(tc.b.rows, tc.b.cols)
			c.Randomize()
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				Dot(dst, a, c)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	m := NewMatrix(5, 5)
	m.Print("m zero")
	m.Randomize()
	m.Print("m randomize")
	m.Activate(Sigmoid)
	m.Print("m sigmoid")
}

func TestActivate(t *testing.T) {
	type testDef struct {
		given  Matrix
		result Matrix
		act    func(a float64) float64
	}

	testCases := []testDef{
		{
			given: Matrix{
				rows: 2,
				cols: 3,
				values: []float64{
					0.5, -1, 2,
					-0.1, 0.3, -0.5,
				},
			},
			result: Matrix{
				rows: 2,
				cols: 3,
				values: []float64{
					0.6224593312018546, 0.2689414213699951, 0.8807970779778823,
					0.47502081252106, 0.574442516811659, 0.3775406687981454,
				},
			},
			act: Sigmoid,
		},
		{
			given: Matrix{
				rows: 3,
				cols: 4,
				values: []float64{
					0.500000, -1.000000, 2.000000, 0.300000,
					-0.100000, 0.300000, -0.500000, 1.200000,
					-2.000000, 1.500000, -0.700000, 0.000000,
				},
			},
			result: Matrix{
				rows: 3,
				cols: 4,
				values: []float64{
					0.622459, 0.268941, 0.880797, 0.574443,
					0.475021, 0.574443, 0.377541, 0.768524,
					0.119203, 0.817574, 0.331812, 0.500000,
				},
			},
			act: Sigmoid,
		},
		{
			given: Matrix{
				rows: 3,
				cols: 4,
				values: []float64{
					0.500000, -1.000000, 2.000000, 0.300000,
					-0.100000, 0.300000, -0.500000, 1.200000,
					-2.000000, 1.500000, -0.700000, 0.000000,
				},
			},
			result: Matrix{
				rows: 3,
				cols: 4,
				values: []float64{
					0.462117, -0.761594, 0.964028, 0.291313,
					-0.099668, 0.291313, -0.462117, 0.833654,
					-0.964028, 0.905148, -0.604367, 0.000000,
				},
			},
			act: Tanh,
		},
		{
			given: Matrix{
				rows: 3,
				cols: 4,
				values: []float64{
					0.500000, -1.000000, 2.000000, 0.300000,
					-0.100000, 0.300000, -0.500000, 1.200000,
					-2.000000, 1.500000, -0.700000, 0.000000,
				},
			},
			result: Matrix{
				rows: 3,
				cols: 4,
				values: []float64{
					0.500000, -0.006321, 2.000000, 0.300000,
					0.000100000, 0.300000, -0.003935, 1.200000,
					-0.0086437, 1.500000, -0.005142, 0.000000,
				},
			},
			act: Elu,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("activation test %v", i), func(t *testing.T) {
			tc.given.Activate(tc.act)
			for j := range tc.given.values {
				assert.Equal(t, true, int(tc.given.values[j]*1000) == int(tc.result.values[j]*1000))
			}
		})
	}
}

func TestConvolve(t *testing.T) {
	filter := NewMatrix(3, 3)
	filter.SetAt(0, 0, 1.0)
	filter.SetAt(0, 1, 0.0)
	filter.SetAt(0, 2, -1.0)
	filter.SetAt(1, 0, 1.0)
	filter.SetAt(1, 1, 0.0)
	filter.SetAt(1, 2, -1.0)
	filter.SetAt(2, 0, 1.0)
	filter.SetAt(2, 1, 0.0)
	filter.SetAt(2, 2, -1.0)
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("convolve dst %v a %v b %v ", tc.dst, tc.a, tc.b), func(t *testing.T) {
			dst := NewMatrix(tc.dst.rows, tc.dst.cols)
			a := NewMatrix(tc.a.rows, tc.a.cols)
			a.Randomize()
			b := NewMatrix(tc.b.rows, tc.b.cols)
			b.Randomize()
			err := Dot(dst, a, b)
			assert.NilError(t, err)
			err = dst.Convolve(filter)
			assert.NilError(t, err)
		})
	}
}

func BenchmarkConvolve(b *testing.B) {
	filter := NewMatrix(3, 3)
	filter.SetAt(0, 0, 1.0)
	filter.SetAt(0, 1, 0.0)
	filter.SetAt(0, 2, -1.0)
	filter.SetAt(1, 0, 1.0)
	filter.SetAt(1, 1, 0.0)
	filter.SetAt(1, 2, -1.0)
	filter.SetAt(2, 0, 1.0)
	filter.SetAt(2, 1, 0.0)
	filter.SetAt(2, 2, -1.0)
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("testing dot for dst %v a %v b %v ", tc.dst, tc.a, tc.b), func(b *testing.B) {
			dst := NewMatrix(tc.dst.rows, tc.dst.cols)
			a := NewMatrix(tc.a.rows, tc.a.cols)
			a.Randomize()
			c := NewMatrix(tc.b.rows, tc.b.cols)
			c.Randomize()
			err := Dot(dst, a, c)
			assert.NilError(b, err)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = dst.Convolve(filter)
				assert.NilError(b, err)
			}
		})
	}
}

func TestNNForward(t *testing.T) {
	architecture := []Schema{
		{Size: 10, Activation: ReluActivation},
		{Size: 20, Activation: ReluActivation},
		{Size: 10, Activation: ReluActivation},
		{Size: 5, Activation: ReluActivation},
		{Size: 3, Activation: ReluActivation},
	}

	nn, err := NewNN(architecture)
	assert.NilError(t, err)
	nn.Randomize()
	in := NewMatrix(1, architecture[0].Size)
	in.Randomize()
	output := nn.Output()
	for i := 0; i < 100; i++ {
		in.Randomize()
		err = nn.Input(in)
		assert.NilError(t, err)
		err = nn.Forward()
		assert.NilError(t, err)
		out := nn.Output()
		assert.Equal(t, len(out.values), len(output.values))
		for j := range out.values {
			assert.Equal(t, false, output.values[j] == out.values[j])
		}

		output = out
		err = nn.Input(in)
		assert.NilError(t, err)
		err = nn.Forward()
		assert.NilError(t, err)
		out = nn.Output()
		for j := range out.values {
			assert.Equal(t, true, output.values[j] == out.values[j])
		}
	}

}

func BenchmarkNNForwardWithInput(b *testing.B) {
	benchCase := [][]Schema{
		{
			{Size: 10, Activation: ReluActivation},
			{Size: 20, Activation: ReluActivation},
			{Size: 5, Activation: ReluActivation},
			{Size: 1, Activation: ReluActivation},
		},
		{
			{Size: 20, Activation: ReluActivation},
			{Size: 50, Activation: ReluActivation},
			{Size: 10, Activation: ReluActivation},
			{Size: 8, Activation: ReluActivation},
			{Size: 5, Activation: ReluActivation},
		},
		{
			{Size: 200, Activation: ReluActivation},
			{Size: 500, Activation: ReluActivation},
			{Size: 100, Activation: ReluActivation},
			{Size: 50, Activation: ReluActivation},
			{Size: 10, Activation: ReluActivation},
		},
	}

	for _, arch := range benchCase {
		b.Run(fmt.Sprintf("size factor %v", arch[0].Size), func(b *testing.B) {
			nn, err := NewNN(arch)
			assert.NilError(b, err)
			nn.Randomize()
			b.ResetTimer()
			in := NewMatrix(0, arch[0].Size)
			in.Randomize()
			for n := 0; n < b.N; n++ {
				err = nn.Input(in)
				assert.NilError(b, err)
				err = nn.Forward()
				assert.NilError(b, err)
			}
		})
	}

}

func TestCost(t *testing.T) {
	type testDef struct {
		arch     []Schema
		dataRows int
	}
	testCases := []testDef{
		{
			arch: []Schema{
				{Size: 10, Activation: ReluActivation},
				{Size: 20, Activation: ReluActivation},
				{Size: 5, Activation: ReluActivation},
				{Size: 1, Activation: ReluActivation},
			},
			dataRows: 100,
		},
		{
			arch: []Schema{
				{Size: 8, Activation: ReluActivation},
				{Size: 10, Activation: EluActivation},
				{Size: 2, Activation: SigmoidActivation},
				{Size: 12, Activation: ReluActivation},
			},
			dataRows: 78,
		},
		{
			arch: []Schema{
				{Size: 80, Activation: ReluActivation},
				{Size: 300, Activation: SigmoidActivation},
				{Size: 200, Activation: EluActivation},
				{Size: 100, Activation: LeakyReluActivation},
				{Size: 12, Activation: SigmoidActivation},
			},
			dataRows: 2000,
		},
		{
			arch: []Schema{
				{Size: 100, Activation: ReluActivation},
				{Size: 1000, Activation: SigmoidActivation},
				{Size: 500, Activation: EluActivation},
				{Size: 100, Activation: LeakyReluActivation},
				{Size: 50, Activation: SigmoidActivation},
			},
			dataRows: 100,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("cost %v test", i), func(t *testing.T) {
			nn, err := NewNN(tc.arch)
			assert.NilError(t, err)
			nn.Randomize()
			in := NewMatrix(tc.dataRows, tc.arch[0].Size)
			out := NewMatrix(tc.dataRows, tc.arch[len(tc.arch)-1].Size)
			in.Randomize()
			out.Randomize()
			nn.Input(in)
			nn.Forward()
			cost, err := nn.Cost(in, out)
			assert.NilError(t, err)
			assert.Equal(t, true, cost != 0)
		})
	}
}

func BenchmarkCost(b *testing.B) {
	type testDef struct {
		name     string
		arch     []Schema
		dataRows int
	}
	testCases := []testDef{
		{
			name: "tiny",
			arch: []Schema{
				{Size: 10, Activation: ReluActivation},
				{Size: 20, Activation: ReluActivation},
				{Size: 5, Activation: ReluActivation},
				{Size: 1, Activation: ReluActivation},
			},
			dataRows: 50,
		},
		{
			name: "small",
			arch: []Schema{
				{Size: 8, Activation: ReluActivation},
				{Size: 10, Activation: EluActivation},
				{Size: 2, Activation: SigmoidActivation},
				{Size: 12, Activation: ReluActivation},
			},
			dataRows: 100,
		},
		{
			name: "decent",
			arch: []Schema{
				{Size: 80, Activation: ReluActivation},
				{Size: 300, Activation: SigmoidActivation},
				{Size: 200, Activation: EluActivation},
				{Size: 100, Activation: LeakyReluActivation},
				{Size: 12, Activation: SigmoidActivation},
			},
			dataRows: 100,
		},
		{
			name: "large",
			arch: []Schema{
				{Size: 100, Activation: ReluActivation},
				{Size: 1000, Activation: SigmoidActivation},
				{Size: 500, Activation: EluActivation},
				{Size: 100, Activation: LeakyReluActivation},
				{Size: 50, Activation: SigmoidActivation},
			},
			dataRows: 200,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("cost %v %s benchmark", i, tc.name), func(b *testing.B) {
			nn, err := NewNN(tc.arch)
			assert.NilError(b, err)
			nn.Randomize()
			in := NewMatrix(tc.dataRows, tc.arch[0].Size)
			out := NewMatrix(tc.dataRows, tc.arch[len(tc.arch)-1].Size)
			in.Randomize()
			out.Randomize()
			nn.Input(in)
			nn.Forward()

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, err := nn.Cost(in, out)
				assert.NilError(b, err)
			}
		})
	}
}

func TestBackpropErrors(t *testing.T) {
	type testDef struct {
		name     string
		arch     []Schema
		dataRows int
	}
	testCases := []testDef{
		{
			name: "tiny",
			arch: []Schema{
				{Size: 10, Activation: ReluActivation},
				{Size: 20, Activation: ReluActivation},
				{Size: 5, Activation: ReluActivation},
				{Size: 1, Activation: ReluActivation},
			},
			dataRows: 50,
		},
		{
			name: "small",
			arch: []Schema{
				{Size: 8, Activation: ReluActivation},
				{Size: 10, Activation: EluActivation},
				{Size: 2, Activation: SigmoidActivation},
				{Size: 12, Activation: ReluActivation},
			},
			dataRows: 100,
		},
		{
			name: "decent",
			arch: []Schema{
				{Size: 80, Activation: ReluActivation},
				{Size: 300, Activation: SigmoidActivation},
				{Size: 200, Activation: EluActivation},
				{Size: 100, Activation: LeakyReluActivation},
				{Size: 12, Activation: SigmoidActivation},
			},
			dataRows: 100,
		},
		{
			name: "large",
			arch: []Schema{
				{Size: 100, Activation: ReluActivation},
				{Size: 1000, Activation: SigmoidActivation},
				{Size: 500, Activation: EluActivation},
				{Size: 100, Activation: LeakyReluActivation},
				{Size: 50, Activation: SigmoidActivation},
			},
			dataRows: 200,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("cost %v test", i), func(t *testing.T) {
			nn, err := NewNN(tc.arch)
			assert.NilError(t, err)
			nn.Randomize()
			in := NewMatrix(tc.dataRows, tc.arch[0].Size)
			out := NewMatrix(tc.dataRows, tc.arch[len(tc.arch)-1].Size)
			in.Randomize()
			out.Randomize()
			v, err := in.Row(0)
			assert.NilError(t, err)
			err = nn.Input(v)
			assert.NilError(t, err)
			err = nn.Forward()
			assert.NilError(t, err)
			err = nn.Backprop(in, out)
			assert.NilError(t, err)
		})
	}
}

func BenchmarkBackprop(b *testing.B) {
	type testDef struct {
		name     string
		arch     []Schema
		dataRows int
	}
	testCases := []testDef{
		{
			name: "tiny",
			arch: []Schema{
				{Size: 10, Activation: ReluActivation},
				{Size: 20, Activation: ReluActivation},
				{Size: 5, Activation: ReluActivation},
				{Size: 1, Activation: ReluActivation},
			},
			dataRows: 50,
		},
		{
			name: "small",
			arch: []Schema{
				{Size: 8, Activation: ReluActivation},
				{Size: 10, Activation: EluActivation},
				{Size: 2, Activation: SigmoidActivation},
				{Size: 12, Activation: ReluActivation},
			},
			dataRows: 100,
		},
		{
			name: "decent",
			arch: []Schema{
				{Size: 80, Activation: ReluActivation},
				{Size: 300, Activation: SigmoidActivation},
				{Size: 200, Activation: EluActivation},
				{Size: 100, Activation: LeakyReluActivation},
				{Size: 12, Activation: SigmoidActivation},
			},
			dataRows: 100,
		},
		{
			name: "large",
			arch: []Schema{
				{Size: 100, Activation: ReluActivation},
				{Size: 1000, Activation: SigmoidActivation},
				{Size: 500, Activation: EluActivation},
				{Size: 100, Activation: LeakyReluActivation},
				{Size: 50, Activation: SigmoidActivation},
			},
			dataRows: 200,
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("cost %v %s benchmark", i, tc.name), func(b *testing.B) {
			nn, err := NewNN(tc.arch)
			assert.NilError(b, err)
			nn.Randomize()
			in := NewMatrix(tc.dataRows, tc.arch[0].Size)
			out := NewMatrix(tc.dataRows, tc.arch[len(tc.arch)-1].Size)
			in.Randomize()
			out.Randomize()
			v, err := in.Row(0)
			assert.NilError(b, err)
			err = nn.Input(v)
			assert.NilError(b, err)
			err = nn.Forward()
			assert.NilError(b, err)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = nn.Backprop(in, out)
				assert.NilError(b, err)
			}
		})
	}
}

func TestBackpropAndLearning(t *testing.T) {
	type testDef struct {
		name string
		arch []Schema
		in   Matrix
		out  Matrix
	}
	testCases := []testDef{
		//	{
		//		name: "tiny",
		//		arch: []Schema{
		//			{Size: 10, Activation: TanhActivation},
		//			{Size: 30, Activation: TanhActivation},
		//			{Size: 5, Activation: TanhActivation},
		//			{Size: 1, Activation: TanhActivation},
		//		},
		//		in: Matrix{
		//			rows: 10,
		//			cols: 10,
		//			values: []float64{
		//				1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		//				2, 3, 4, 5, 6, 7, 8, 9, 0, 11,
		//				3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		//				4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		//				5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
		//				6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		//				7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		//				8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
		//				9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
		//				10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		//			},
		//		},
		//		out: Matrix{
		//			rows: 10,
		//			cols: 1,
		//			values: []float64{
		//				1,
		//				2,
		//				3,
		//				4,
		//				5,
		//				6,
		//				7,
		//				8,
		//				9,
		//				10,
		//			},
		//		},
		//	},
		{
			name: "simplistic",
			arch: []Schema{
				{Size: 2, Activation: SigmoidActivation},
				{Size: 4, Activation: SigmoidActivation},
				{Size: 1, Activation: SigmoidActivation},
			},
			in: Matrix{
				rows: 4,
				cols: 2,
				values: []float64{
					0, 0,
					0, 1,
					1, 0,
					1, 1,
				},
			},
			out: Matrix{
				rows: 4,
				cols: 1,
				values: []float64{
					0,
					1,
					1,
					0,
				},
			},
		},
	}

	const epochs = 200000
	const learningRate float64 = 0.1

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("cost %v test", i), func(t *testing.T) {
			nn, err := NewNN(tc.arch)
			assert.NilError(t, err)
			nn.Randomize()
			tc.in.Normalize(tc.in.Min(), tc.in.Max())
			tc.out.Normalize(tc.out.Min(), tc.out.Max())
			tc.in.Print("in")
			tc.out.Print("out")
			maxCost := math.MaxFloat64
		epochLoop:
			for epoch := 0; epoch < epochs; epoch++ {
				err = nn.Backprop(tc.in, tc.out)
				assert.NilError(t, err)
				err = nn.Learn(learningRate)
				assert.NilError(t, err)
				cost, err := nn.Cost(tc.in, tc.out)
				if err != nil {
					break epochLoop
				}
				maxCost = cost
			}

			fmt.Printf("max cost: %.4f \n", maxCost)
			for i := 0; i < tc.in.rows; i++ {
				row, err := tc.in.Row(i)
				assert.NilError(t, err)
				nn.Input(row)
				nn.Forward()
				out, _ := tc.out.Row(i)
				fmt.Printf(" %v -----------------\n", i)
				out.Print("expected")
				nn.Output().Print("caclulated")
				fmt.Printf(" %v =================\n", i)
			}
		})
	}
}
