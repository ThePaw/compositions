// Copyright 2012 The Composition Authors. All rights reserved. See the LICENSE file.

package compositions

/* 
Centered log ratio transform of a (dataset of) composition(s) and its inverse.

Arguments
data 		a composition or a data matrix of compositions, not necessarily closed
clr 		the clr-transform of a composition or a data matrix of clr-transforms of compositions, not necessarily centered (i.e. summing up to zero)

The clr-transform maps a composition in the D-part Aitchison-simplex isometrically to a D-dimensonal euclidian vector subspace: consequently, the transformation is not injective. Thus resulting covariance matrices are always singular.
The data can then be analysed in this transformation by all classical multivariate analysis tools not relying on a full rank of the covariance. See ilr and alr for alternatives. The interpretation of the results is relatively easy since the relation between each original part and a transformed variable is preserved.
The centered logratio transform is given by

clr(x) = (ln x - mean(ln x) ) The image of the clr is a vector with entries summing to 0. This hyperplane is also called the clr-plane.

Aitchison, J. (1986) The Statistical Analysis of Compositional Data, Monographs on Statistics and Applied Probability. Chapman & Hall Ltd., London (UK). 416p. 


*/

func Clr(in, out Matrix64) {
	var meanln float64
	rows, cols := in.Dims()
	r, c := out.Dims()
	if !(r == rows && c == cols) {
		panic("bad dimensions")
	}

	for i := 0; i < rows; i++ {
		meanln = 0
		for j := 0; j < cols; j++ {
			meanln += ln(in[i][j])
		}
		meanln /= float64(cols)
		for j := 0; j < cols; j++ {
			out[i][j] = ln(in[i][j]) - meanln
		}
	}
}

// ToDo: Inv_clr(in, out Matrix64)
