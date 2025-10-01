package main

const MOD int = 1_000_000_007

func zigZagArrays(n int, l int, r int) (ans int) {
	m := r - l + 1
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		res := 0
		if i%2 > 0 {
			for j := m - 1; j >= 0; j-- {
				now := dp[j]
				dp[j] = res % MOD
				res += now
			}
		} else {
			for j, v := range dp {
				dp[j] = res % MOD
				res += v
			}
		}
	}

	for _, v := range dp {
		ans += v
		ans %= MOD
	}
	ans *= 2
	ans %= MOD
	return
}

func main() {

}
