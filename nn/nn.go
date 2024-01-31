package nn

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
)

type ActivationType string

const (
	SigmoidActivation   ActivationType = "sigmoid"
	TanhActivation      ActivationType = "tanh"
	ReluActivation      ActivationType = "relu"
	LeakyReluActivation ActivationType = "leaky_relu"
	EluActivation       ActivationType = "elu"
)

const alpha float64 = 0.01

func randFloat64() float64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(1<<53))
	if err != nil {
		panic(err)
	}
	return float64(nBig.Int64()) / (1 << 53)
}

// Neutral activation function does nothing to the input value.
func Neutral(a float64) float64 { return a }

// Sigmoid calculate sigmoid function for given input.
func Sigmoid(a float64) float64 { return 1.0 / (1.0 + math.Exp(-a)) } // Derivative of the sigmoid activation function

// SigmoidDeactivate calculates sigmoid deactivation value.
func SigmoidDeactivate(a float64) float64 {
	return a * (1 - a)
}

// Than calculates hyperbolic tangent function for given input.
func Tanh(a float64) float64 { return (math.Exp(a) - math.Exp(-a)) / (math.Exp(a) + math.Exp(-a)) }

// TanhDeactivate calculates tangent deactivation value.
func TanhDeactivate(a float64) float64 {
	return 1 - a*a
}

// Relu calculates rectified linear unit function for given input.
func Relu(a float64) float64 {
	return math.Max(0, a)
}

// ReLUDeactivate calculates rectified linear unit  deactivation value.
func ReLUDeactivate(a float64) float64 {
	if a > 0 {
		return 1
	}
	return 0
}

// LeakyRelu calculates leaky rectified linear unit function for given input.
// Alpha value is constant and equal to 0.01
func LeakyRelu(a float64) float64 {
	if a > 0 {
		return a
	}
	return a * alpha
}

// LeakyReLUDeactivate calculates leaky rectified linear unit deactivation value.
func LeakyReLUDeactivate(a float64) float64 {
	if a > 0 {
		return 1
	}
	return alpha
}

// Elu calculates exponential linear unit function for given input.
func Elu(a float64) float64 {
	if a >= 0 {
		return a
	}
	return alpha * (math.Exp(a) - 1)
}

// EluDeactivate calculates exponential linear unit deactivation value.
func EluDeactivate(x float64) float64 {
	if x >= 0 {
		return 1
	}
	return Elu(x) + alpha
}

// Sotmax calculates softmax function on slice of float64 values.
// Returns slice with the same size as input slice.
func Softmax(x []float64) []float64 {
	var sum float64
	result := make([]float64, len(x))

	for _, val := range x {
		sum += math.Exp(val)
	}

	for i, val := range x {
		result[i] = math.Exp(val) / sum
	}

	return result
}

func selectActivationFunction(t ActivationType) ActivationFunction {
	switch t {
	case SigmoidActivation:
		return Sigmoid
	case TanhActivation:
		return Tanh
	case ReluActivation:
		return Relu
	case LeakyReluActivation:
		return LeakyRelu
	case EluActivation:
		return Elu
	default:
		return Neutral
	}
}

func selectDeactivationActivationFunction(t ActivationType) ActivationFunction {
	switch t {
	case SigmoidActivation:
		return SigmoidDeactivate
	case TanhActivation:
		return TanhDeactivate
	case ReluActivation:
		return ReLUDeactivate
	case LeakyReluActivation:
		return LeakyReLUDeactivate
	case EluActivation:
		return EluDeactivate
	default:
		return Neutral
	}
}

// ActivationFunction is an activation signature function.
// Activations functions avaliable in this package are: Sigmoid, Tanh, Relu, LeakyRelu, Elu.
type ActivationFunction func(a float64) float64

// Matrix is a 2D matrix.
type Matrix struct {
	values []float64
	rows   int
	cols   int
}

// NewMatrix creates new Matrix with given number of rows and columns.
func NewMatrix(rows, cols int) Matrix {
	return Matrix{
		values: make([]float64, rows*cols),
		rows:   rows,
		cols:   cols,
	}
}

// Min finds the minimum number in matrix.
func (m Matrix) Min() float64 {
	min := math.MaxFloat64
	for _, v := range m.values {
		if v < min {
			min = v
		}
	}
	return min
}

// Max finds the mximum number in matrix.
func (m Matrix) Max() float64 {
	max := -math.MaxFloat64
	for _, v := range m.values {
		if v > max {
			max = v
		}
	}
	return max
}

// Normalize normalizes the values in the matrix.
func (m Matrix) Normalize(min, max float64) {
	for i := range m.values {
		m.values[i] = (m.values[i] - min) / (max - min)
	}
}

// UnNormalize reverts normalization.
// Values may differ a little from origin because of loating point rounding happening in normalization and unnormalizaton.
func (m Matrix) UnNormalize(min, max float64) {
	for i := range m.values {
		m.values[i] = (m.values[i] + min) * (max - min)
	}
}

// At returns value at given row and column or if outside of range returns an error.
func (m Matrix) At(row, col int) (float64, error) {
	if row >= m.rows || row < 0 {
		return 0.0, fmt.Errorf("wrong row index")
	}
	if col >= m.cols || col < 0 {
		return 0.0, fmt.Errorf("wrong column index")
	}
	return m.values[col+row*m.cols], nil
}

// SetAt sets value at given row and column or if outside of range returns an error.
func (m Matrix) SetAt(row, col int, v float64) error {
	if row >= m.rows || row < 0 {
		return fmt.Errorf("wrong row index")
	}
	if col >= m.cols || col < 0 {
		return fmt.Errorf("wrong column index")
	}
	m.values[col+row*m.cols] = v
	return nil
}

// Randomize rendomizes the matrix values.
func (m Matrix) Randomize() error {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.SetAt(i, j, randFloat64())
		}
	}
	return nil
}

// Zero zeros the matrix values.
func (m Matrix) Zero() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.SetAt(i, j, 0.0)
		}
	}
}

// Activate activates all values in the matrix.
func (m Matrix) Activate(actf ActivationFunction) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			v, _ := m.At(i, j)
			m.SetAt(i, j, actf(v))
		}
	}
}

// Convolve function applies filter convolution into matrix.
// Filter matrix should have the same number of rows and columns and be of smaller size than
// the convolution receiver matrix.
func (m Matrix) Convolve(filter Matrix) error {
	if filter.rows != filter.cols {
		return fmt.Errorf(
			"expected the same number of rows and columns, received rows [ %v ] cols [ %v ]",
			filter.rows, filter.cols,
		)
	}
	if filter.rows >= m.rows {
		return fmt.Errorf(
			"expected matrix of bigger size than the filter, received rows in matrix [ %v ], in filter [ %v ]",
			m.rows, filter.rows,
		)
	}
	if filter.cols >= m.cols {
		return fmt.Errorf(
			"expected matrix of bigger size than the filter, received columns in matrix [ %v ], in filter [ %v ]",
			m.cols, filter.cols,
		)
	}

	padding := filter.cols / 2

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			// apply filter
			sum := 0.0
			for k := -padding; k <= padding; k++ {
				for l := -padding; l <= padding; l++ {
					r := i + k
					c := j + l
					if r >= 0 && r < m.rows && c >= 0 && c < m.cols {
						mV, err := m.At(r, c)
						if err != nil {
							// unreachable
							return fmt.Errorf("unreachable, %w", err)
						}
						mF, err := filter.At(k+padding, l+padding)
						if err != nil {
							// unreachable
							return fmt.Errorf("unreachable, %w", err)
						}
						sum += mV * mF
					}
				}
			}
			m.SetAt(i, j, sum)
		}
	}
	return nil
}

// Row returns copy of a matrix row at row index as a matrix with single row.
func (m Matrix) Row(row int) (Matrix, error) {
	var nm Matrix
	if row < 0 || row >= m.rows {
		return nm, fmt.Errorf("exceeded row number, expected row in range [ 0, %v ], received %v", m.rows, row)
	}
	nm.cols = m.cols
	nm.rows = 1
	nm.values = make([]float64, m.cols)
	copy(nm.values, m.values[row*m.cols:row*m.cols+m.cols])
	return nm, nil
}

// Print prints matrix into stdout.
func (m Matrix) Print(name string) {
	fmt.Printf("%s = [\n", name)
	for i := 0; i < m.rows; i++ {
		fmt.Printf("  ")
		for j := 0; j < m.cols; j++ {
			v, _ := m.At(i, j)
			fmt.Printf("%.3f, ", v)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("]\n")
}

// Dot calculates dot product of two matrices and saves the product in the destination matrix.
// Input matrix a needs to have number of columns corresponding to matrix b number of rows otherwise error is returned.
// Destination matrix dst needs to have number of rows corresponding to matrix a number of rows otherwise error is returned.
// Destination matrix dst needs to have number of columns corresponding to matrix b number of columns otherwise error is returned.
func Dot(dst, a, b Matrix) error {
	if a.cols != b.rows {
		return fmt.Errorf(
			"wrong size of matrices, matrix a cols [ %v ] doesn't equal to matrix b rows [ %v ]",
			a.cols, b.rows,
		)
	}
	if dst.rows != a.rows {
		return fmt.Errorf(
			"wrong size of matrices, matrix a rows [ %v ] doesn't equal to matrix dst rows [ %v ]",
			a.rows, dst.rows,
		)
	}
	if dst.cols != b.cols {
		return fmt.Errorf(
			"wrong size of matrices, matrix a cols [ %v ] doesn't equal to matrix dst cols [ %v ]",
			b.cols, dst.cols,
		)
	}

	for i := 0; i < dst.rows; i++ {
		for j := 0; j < dst.cols; j++ {
			dst.SetAt(i, j, 0.0)
			for k := 0; k < a.cols; k++ {
				dstV, err := dst.At(i, j)
				if err != nil {
					return err
				}
				aV, err := a.At(i, k)
				if err != nil {
					// unreachable
					return fmt.Errorf("unreachable, %w", err)
				}
				bV, err := b.At(k, j)
				if err != nil {
					// unreachable
					return fmt.Errorf("unreachable, %w", err)
				}
				if err := dst.SetAt(i, j, dstV+aV*bV); err != nil {
					return fmt.Errorf("unreachable, %w", err)
				}
			}
		}
	}
	return nil
}

// Copy copies matrix from src to dst.
// Matrices shall have the same size (number of rows and columns shall match) or error is returned.
func Copy(dst, src Matrix) error {
	if dst.rows != src.rows {
		return fmt.Errorf(
			"unmatching number of rows, dst [ %v ], src [ %v]", dst.rows, src.rows)
	}
	if dst.cols != src.cols {
		return fmt.Errorf(
			"unmatching number of columns, dst [ %v ], src [ %v]", dst.cols, src.cols)
	}
	if len(dst.values) != len(src.values) {
		return fmt.Errorf(
			"unmatching underlining slice len, dst [ %v ], src [ %v]", len(dst.values), len(src.values))

	}
	copy(dst.values, src.values)
	return nil
}

// Sum sums src and dst matrices to dst.
// Matrices shall have the same size (number of rows and columns shall match) or error is returned.
func Sum(dst, src Matrix) error {
	if dst.rows != src.rows {
		return fmt.Errorf(
			"unmatching number of rows, dst [ %v ], src [ %v]", dst.rows, src.rows)
	}
	if dst.cols != src.cols {
		return fmt.Errorf(
			"unmatching number of columns, dst [ %v ], src [ %v]", dst.cols, src.cols)
	}

	for i := 0; i < src.rows; i++ {
		for j := 0; j < src.cols; j++ {
			sV, err := src.At(i, j)
			if err != nil {
				// unreachable
				return fmt.Errorf("unreachable, %w", err)
			}
			dV, err := dst.At(i, j)
			if err != nil {
				// unreachable
				return fmt.Errorf("unreachable, %w", err)
			}
			err = dst.SetAt(i, j, sV+dV)
			if err != nil {
				// unreachable
				return fmt.Errorf("unreachable, %w", err)
			}
		}
	}
	return nil
}

// Layer is a NN layer.
type layer struct {
	ws    Matrix // weights
	bs    Matrix // biases
	as    Matrix // activations
	act   ActivationFunction
	deact ActivationFunction
}

// Schema describes the layer schema and can be decoded from json or yaml file.
type Schema struct {
	Size       int            `json:"size" yaml:"size"`
	Activation ActivationType `json:"activates_type" yaml:"activates_type"`
}

// NN is a Neural Network.
type NN struct {
	arch []layer
	bpNN *NN
}

// NewNN creates new NN based on the architecture.
func NewNN(architecture []Schema) (NN, error) {
	var nn NN
	if len(architecture) < 3 {
		return nn, errors.New("expecting at list 3 layers, input layer, hidden layer(s), output layer")
	}

	nn.arch = make([]layer, len(architecture))
	nn.arch[0] = layer{as: NewMatrix(1, architecture[0].Size), act: selectActivationFunction(architecture[0].Activation)}
	for i := range architecture {
		actf := selectActivationFunction(architecture[i].Activation)
		dactf := selectDeactivationActivationFunction(architecture[i].Activation)
		switch i {
		case 0:
			continue
		case len(architecture) - 1:
			actf = Neutral
			dactf = Neutral
		default:
		}
		nn.arch[i].act = actf
		nn.arch[i].deact = dactf
		nn.arch[i-1].ws = NewMatrix(architecture[i-1].Size, architecture[i].Size)
		nn.arch[i-1].bs = NewMatrix(1, architecture[i].Size)
		nn.arch[i].as = NewMatrix(1, architecture[i].Size)
	}

	return nn, nil
}

func (nn NN) zero() {
	for _, l := range nn.arch {
		l.as.Zero()
		l.ws.Zero()
		l.bs.Zero()
	}
}

func (nn *NN) createNewBackpropNN() {
	bpnn := new(NN)
	bpnn.arch = make([]layer, len(nn.arch))
	for i := range nn.arch {
		bpnn.arch[i].as = NewMatrix(nn.arch[i].as.rows, nn.arch[i].as.cols)
		bpnn.arch[i].ws = NewMatrix(nn.arch[i].ws.rows, nn.arch[i].ws.cols)
		bpnn.arch[i].bs = NewMatrix(nn.arch[i].bs.rows, nn.arch[i].bs.cols)
	}
	nn.bpNN = bpnn
}

func (nn NN) validateInOutSize(in, out Matrix) error {
	if in.cols != nn.arch[0].as.cols {
		return fmt.Errorf(
			"expected input number of columns [ %v ], received  [ %v ]",
			nn.arch[0].as.cols,
			in.cols,
		)
	}
	if out.cols != nn.arch[len(nn.arch)-1].as.cols {
		return fmt.Errorf(
			"expected output number of columns [ %v ], received  [ %v ]",
			nn.arch[len(nn.arch)-1].as.cols,
			in.cols,
		)
	}
	if in.rows != out.rows {
		return fmt.Errorf(
			"expected input and output to match with number of rows, input [ %v ] rows, out [ %v ] rows",
			in.rows,
			out.rows,
		)
	}
	return nil
}

// Randomize randomizes the all layers matrices values.
func (nn NN) Randomize() {
	for _, l := range nn.arch {
		l.as.Randomize()
		l.ws.Randomize()
		l.bs.Randomize()
	}
}

// Input places input in to the input layer or if input matrix doesn't match returns error.
func (nn *NN) Input(in Matrix) error {
	return Copy(nn.arch[0].as, in)
}

// Output returns output matrix.
func (nn NN) Output() Matrix {
	m := NewMatrix(nn.arch[len(nn.arch)-1].as.rows, nn.arch[len(nn.arch)-1].as.cols)
	Copy(m, nn.arch[len(nn.arch)-1].as)
	return m
}

// Forward applies feed forward action on neural network.
func (nn NN) Forward() error {
	for i := 0; i < len(nn.arch)-1; i++ {
		if err := Dot(nn.arch[i+1].as, nn.arch[i].as, nn.arch[i].ws); err != nil {
			return fmt.Errorf("unreachable, %w", err)
		}
		if err := Sum(nn.arch[i+1].as, nn.arch[i].bs); err != nil {
			return fmt.Errorf("unreachable, %w", err)
		}
		nn.arch[i+1].as.Activate(nn.arch[i+1].act)
	}
	return nil
}

// Cost function calculates the cost.
func (nn NN) Cost(in, out Matrix) (float64, error) {
	var cost float64
	if err := nn.validateInOutSize(in, out); err != nil {
		return cost, err
	}

	for i := 0; i < in.rows; i++ {
		inRow, err := in.Row(i)
		if err != nil {
			return 0.0, fmt.Errorf("unreachable, %w", err)
		}
		if err := nn.Input(inRow); err != nil {
			return 0.0, fmt.Errorf("unreachable, %w", err)
		}
		if err := nn.Forward(); err != nil {
			return 0.0, err
		}
		for j := 0; j < out.cols; j++ {
			outV, err := nn.arch[len(nn.arch)-1].as.At(0, j)
			if err != nil {
				return 0.0, err
			}
			testV, err := out.At(i, j)
			if err != nil {
				return 0.0, err
			}
			d := outV - testV
			cost += d * d
		}
	}

	cost /= float64(in.rows)
	if math.IsNaN(cost) {
		return 0.0, fmt.Errorf("cost is NaN")
	}
	return cost, nil
}

// Backprop performs back propagation on NN.
func (nn *NN) Backprop(in, out Matrix) error {
	switch nn.bpNN {
	case nil:
		nn.createNewBackpropNN()
	default:
		nn.bpNN.zero()
	}

	if err := nn.validateInOutSize(in, out); err != nil {
		return err
	}

	for i := 0; i < in.rows; i++ {
		s := i * in.cols
		copy(nn.arch[0].as.values, in.values[s:s+in.cols])

		if err := nn.Forward(); err != nil {
			return err
		}

		for _, l := range nn.bpNN.arch {
			l.as.Zero()
		}

		for j := 0; j < out.cols; j++ {
			outV, err := nn.arch[len(nn.arch)-1].as.At(0, j)
			if err != nil {
				return err
			}
			testV, err := out.At(i, j)
			if err != nil {
				return err
			}
			nn.bpNN.arch[len(nn.bpNN.arch)-1].as.SetAt(0, j, outV-testV)
		}

		for idx := len(nn.arch) - 1; idx > 0; idx-- {
			for j := 0; j < nn.arch[idx].as.cols; j++ {
				val, err := nn.arch[idx].as.At(0, j)
				if err != nil {
					return err
				}

				dVal, err := nn.bpNN.arch[idx].as.At(0, j)
				if err != nil {
					return err
				}

				qVal := nn.arch[idx].deact(val)

				bsVal, err := nn.bpNN.arch[idx-1].bs.At(0, j)
				if err != nil {
					return err
				}

				err = nn.bpNN.arch[idx-1].bs.SetAt(0, j, bsVal+2.0*dVal*qVal)
				if err != nil {
					return err
				}

				for k := 0; k < nn.arch[idx-1].as.cols; k++ {
					pVal, err := nn.arch[idx-1].as.At(0, k)
					if err != nil {
						return err
					}

					wVal, err := nn.arch[idx-1].ws.At(k, j)
					if err != nil {
						return err
					}

					wsVal, err := nn.bpNN.arch[idx-1].ws.At(k, j)
					if err != nil {
						return err
					}

					err = nn.bpNN.arch[idx-1].ws.SetAt(k, j, wsVal+2.0*dVal*qVal*pVal)
					if err != nil {
						return err
					}

					asVal, err := nn.bpNN.arch[idx-1].as.At(0, k)
					if err != nil {
						return err
					}
					if math.IsNaN(asVal) {
						// panic("is NaN")
						return fmt.Errorf("is NaN")
					}
					err = nn.bpNN.arch[idx-1].as.SetAt(0, k, asVal+2.0*dVal*qVal*wVal)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	for i := 0; i < len(nn.bpNN.arch)-1; i++ {
		for j := 0; j < nn.bpNN.arch[i].ws.rows; j++ {
			for k := 0; k < nn.bpNN.arch[i].ws.cols; k++ {
				val, err := nn.bpNN.arch[i].ws.At(j, k)
				if err != nil {
					return err
				}

				err = nn.bpNN.arch[i].ws.SetAt(j, k, val/float64(out.rows))
				if err != nil {
					return err
				}
			}
		}
		for k := 0; k < nn.bpNN.arch[i].bs.cols; k++ {
			val, err := nn.bpNN.arch[i].bs.At(0, k)
			if err != nil {
				return err
			}

			err = nn.bpNN.arch[i].bs.SetAt(0, k, val/float64(out.rows))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ClearBackprop clears back propagation removing it from memory.
// Learning is impossible after back propagation is cleared.
// Use it only after learning process is finished and there is a need to save memory.
func (nn *NN) ClearBackprop() {
	nn.bpNN = nil
}

// Learn performs learning by applying back propagation with given learning rate.
// Requires back propagation to be performed first.
func (nn NN) Learn(r float64) error {
	if nn.bpNN == nil {
		return errors.New("back propagation is cleared, cannot perform learning")
	}
	for i := 0; i < len(nn.arch)-1; i++ {
		for j := 0; j < nn.arch[i].ws.rows; j++ {
			for k := 0; k < nn.arch[i].ws.cols; k++ {
				wVal, err := nn.arch[i].ws.At(j, k)
				if err != nil {
					return fmt.Errorf("unreachable, %w", err)
				}
				bpwVal, err := nn.bpNN.arch[i].ws.At(j, k)
				if err != nil {
					return fmt.Errorf("unreachable, %w", err)
				}
				err = nn.arch[i].ws.SetAt(j, k, wVal-r*bpwVal)
				if err != nil {
					return fmt.Errorf("unreachable, %w", err)
				}
			}
		}

		for k := 0; k < nn.arch[i].bs.cols; k++ {
			bVal, err := nn.arch[i].bs.At(0, k)
			if err != nil {
				return fmt.Errorf("unreachable, %w", err)
			}
			bpbVal, err := nn.bpNN.arch[i].bs.At(0, k)
			if err != nil {
				return fmt.Errorf("unreachable, %w", err)
			}
			err = nn.arch[i].bs.SetAt(0, k, bVal-r*bpbVal)
			if err != nil {
				return fmt.Errorf("unreachable, %w", err)
			}
		}
	}

	return nil
}

// PrintActivationLayer prints activation layer from NN architecture at given index.
func (nn NN) PrintActivationLayer(idx int) {
	if idx < 0 || idx > len(nn.arch) {
		fmt.Printf("index out of architecture size, expected 0 to %v, received %v\n", len(nn.arch)-1, idx)
	}
	nn.arch[idx].as.Print(fmt.Sprintf("activation layer %v", idx))
}

// PrintWeightsLayer prints weights layer from NN architecture at given index.
func (nn NN) PrintWeightsLayer(idx int) {
	if idx < 0 || idx > len(nn.arch)-1 {
		fmt.Printf("index out of architecture size, expected 0 to %v, received %v\n", len(nn.arch)-1, idx)
	}
	nn.arch[idx].ws.Print(fmt.Sprintf("weights layer %v", idx))
}

// PrintBiasLayer prints bias layer from NN architecture at given index.
func (nn NN) PrintBiasLayer(idx int) {
	if idx < 0 || idx > len(nn.arch)-1 {
		fmt.Printf("index out of architecture size, expected 0 to %v, received %v\n", len(nn.arch)-1, idx)
	}
	nn.arch[idx].bs.Print(fmt.Sprintf("bias layer %v", idx))
}
