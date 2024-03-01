package main

func LevenshteinDistanceFullMatrix(a, b string) int {
	m := len(a)
	n := len(b)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

func LevenshteinDistanceTwoRow(a, b string) int {
	m := len(a)
	n := len(b)

	prevRow := make([]int, n+1)
	currRow := make([]int, n+1)

	for j := 1; j <= n; j++ {
		prevRow[j] = j
	}

	for i := 1; i <= m; i++ {
		currRow[0] = i
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				currRow[j] = prevRow[j-1]
			} else {
				currRow[j] = 1 + min(currRow[j-1], prevRow[j], prevRow[j-1])
			}
		}
		copy(prevRow, currRow)
	}
	return currRow[n]
}

func LevenshteinDistanceOneRow(a, b string) int {
	m := len(a)
	n := len(b)

	row := make([]int, max(n, m)+1)

	for j := range row {
		row[j] = j
	}

	for i := 1; i <= m; i++ {
		prevDiagonal := row[0]
		row[0] = i
		for j := 1; j <= n; j++ {
			prevAbove := row[j]
			if a[i-1] == b[j-1] {
				row[j] = prevDiagonal
			} else {
				row[j] = 1 + min(row[j-1], prevAbove, prevDiagonal)
			}
			prevDiagonal = prevAbove
		}
	}
	return row[n]
}
