// Copyright 2012 The Compositions Authors. All rights reserved. See the LICENSE file.

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

// Clr returns Centered log ratio transform of a dataset of compositions.
func Clr(in, out Matrix64) {
	rows, cols := in.Dims()
	r, c := out.Dims()
	if !(r == rows && c == cols) {
		panic("bad dimensions")
	}

	for i := 0; i < rows; i++ {
		meanln := 0.0
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

/*
Additive log ratio transform of a dataset of compositions.

Arguments
data 	a composition, not necessarily closed
alr 		the alr-transform of a composition, thus a (D-1)-dimensional real vector

Details

The alr-transform maps a composition in the D-part Aitchison-simplex non-isometrically to a D-1 dimensonal euclidian vector, treating the last part as common denominator of the others. The data can then be analysed in this transformation by all classical multivariate analysis tools not relying on a distance. The interpretation of the results is relatively simple, since the relation to the original D-1 first parts is preserved. However distance is an extremely relevant concept in most types of analysis, where a clr or ilr transformation should be preferred.
The additive logratio transform is given by

alr(x)_i := ln(x_i/x_D)

References

Aitchison, J. (1986) The Statistical Analysis of Compositional Data Monographs on Statistics and Applied Probability. Chapman & Hall Ltd., London (UK). 416p.
*/

// Alr returns Additive log ratio transform of a dataset of compositions.
func Alr(in, out Matrix64) {
	//alr matrix has one column less than data matrix
	rows, cols := in.Dims()
	r, c := out.Dims()
	if !(r == rows && c == cols-1) {
		panic("bad dimensions")
	}

	lastcol := cols
	for i := 0; i < rows; i++ {
		for j := 0; j < cols-1; j++ {

			out[i][j] = ln(in[i][j] / in[i][lastcol])
		}
	}
}

// ToDo: inv_alr(Matrix *in, Matrix *out)

/*
Closure of a composition
Description

Closes compositions to sum up to one , by dividing each part by the sum.

Arguments
in 	matrix of compositions
out		matrix of compositions closed to one

Details

The closure operation is given by

clo(x) := ( x_i / sum_j(x_i))

clo generates a composition without assigning one of the compositional classes acomp or rcomp. Note that after computing the closed-to-one version, obtaining a version closed to any other value is done by simple multiplication.

clo can be used to unclass compositions.
References

Aitchison, J. (1986) The Statistical Analysis of Compositional Data Monographs on Statistics and Applied Probability. Chapman & Hall Ltd., London (UK). 416p. 
*/

// Clo closes compositions to sum up to one , by dividing each part by the sum.
func Clo(in, out Matrix64) {
	rows, cols := in.Dims()
	r, c := out.Dims()
	if !(r == rows && c == cols) {
		panic("bad dimensions")
	}

	for i := 0; i < rows; i++ {
		sum := 0.0
		for j := 0; j < cols; j++ {
			sum += in[i][j]
		}

		for j := 0; j < cols; j++ {
			out[i][j] = in[i][j] / sum
		}
	}
}

/*
Centered planar transform
Compute the centered planar transform of a (dataset of) compositions and its inverse.

Arguments
in 	a composition or a data.matrix of compositions, not necessarily closed
out 	the cpt-transform of a composition or a data matrix of cpt-transforms of compositions. It is checked that the z sum up to 0.

The cpt-transform maps a composition in the D-part real-simplex isometrically to a D-1 dimensional euclidian vector space, identified with a plane parallel to the simplex but passing through the origin. However the transformation is not injective and does not even reach the whole plane. Thus resulting covariance matrices are always singular.

The data can then be analysed in this transformed space by all classical multivariate analysis tools not relying on a full rank of the covariance matrix. See ipt and apt for alternatives. The interpretation of the results is relatively easy since the relation of each transformed component to the original parts is preserved.

The centered planar transform is given by

cpt(x)_i = clo(x)_i - 1/D

References

van den Boogaart, K.G. and R. Tolosana-Delgado (2007) "compositions": a unified R package to analyze Compositional Data, Computers & Geosciences. (in press).
*/

// Cpt computes the Centered planar transform of a dataset of compositions.
func Cpt(in, out Matrix64) {
	rows, cols := in.Dims()
	r, c := out.Dims()
	if !(r == rows && c == cols) {
		panic("bad dimensions")
	}

	Clo(in, out)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			out[i][j] = out[i][j] - 1/float64(cols)
		}
	}
}

/*
Additive planar transform
Compute the additive planar transform of a (dataset of) compositions or its inverse.

Arguments:
in 	a composition or a matrix of compositions, not necessarily closed
out 	the apt-transform of a composition or a matrix of alr-transforms of compositions

Details:
The apt-transform maps a composition in the D-part real-simplex linearly to a D-1 dimensional euclidian vector. Although the transformation does not reach the whole R^{D-1}, resulting covariance matrices are typically of full rank.
The data can then be analysed in this transformation by all classical multivariate analysis tools not relying on distances. See cpt and ipt for alternatives. The interpretation of the results is easy since the relation to the first D-1 original variables is preserved.

The additive planar transform is given by
apt(x)_i := clo(x)_i, i=1,...,D-1
*/

// Apt computes the additive planar transform of a dataset of compositions.
func Apt(in, out Matrix64) {
	rows, cols := in.Dims()
	r, c := out.Dims()
	if !(r == rows && c == cols) {
		panic("bad dimensions")
	}

	for i := 0; i < rows; i++ {
		sum := 0.0
		for j := 0; j < cols-1; j++ { // D-1
			sum += in[i][j]
		}

		for j := 0; j < cols; j++ { //D

			out[i][j] = in[i][j] / sum
		}
	}
}
