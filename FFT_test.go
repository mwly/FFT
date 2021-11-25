package FFT

import "testing"

func TestFFT(t *testing.T) {

	result := FFT(make([]float64, 1))
	thing := make([]complex128, 1)

	for i, val := range result {
		if thing[i] != val {
			t.Fatal("one value was not equal")
		}
	}

}
