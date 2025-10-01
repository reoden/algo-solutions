package main

const mod int = 1_000_000_007

type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod
			}
		}
	}
	return c
}

// a^n * f0
func (a matrix) powMul(n int, f0 matrix) matrix {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func zigZagArrays(n int, l int, r int) (ans int) {
	m := r - l + 1
	mat := newMatrix(m*2, m*2)

	for i := range m {
		for j := range i {
			mat[i][j+m] = 1
		}
		for j := i + 1; j < m; j++ {
			mat[i+m][j] = 1
		}
	}

	baseMat := newMatrix(m*2, 1)
	for i := range baseMat {
		baseMat[i][0] = 1
	}
	result := mat.powMul(n-1, baseMat)
	for _, v := range result {
		ans += v[0]
		ans %= mod
	}

	return
}

func main() {

}
