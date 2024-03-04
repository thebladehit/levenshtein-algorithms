# levenshtein-algorithms

## Introduction
The Levenshtein distance is a measure of the difference between two sequences. It is defined as the minimum number of single-character edits (insertions, deletions, or substitutions) required to change one word into the other. This metric is useful in applications such as spell checking, plagiarism detection, and DNA sequencing.

## Recursive approach
The recursive approach to calculating the Levenshtein distance directly implements the algorithm's definition. It compares each character of the two strings, recursively calculating the cost of insertions, deletions, and substitutions needed to align the strings.

## Iterative approach with Full Matrix
The iterative technique with a full matrix uses a 2D matrix to hold the intermediate results of the Levenshtein distance calculation. It begins with empty strings and iteratively fills the matrix row by row. It computes the minimum cost of insertions, deletions, and replacements based on the characters of both strings.

### How it works
The algorithm fills the matrix row by row, each cell representing the minimum edit distance between substrings. The final value at the bottom-right corner of the matrix is the Levenshtein distance between the two full strings.

#### Example: "ab", "abc"

|       | "" | "a" | "ab" | "abc" | 
|:-----:|:--:|:---:|:----:|:-----:|
|  ""   | 0  |  1  |  2   |   3   |
|  "a"  | 1  |  0  |  1   |   2   |
| "ab"  | 2  |  1  |  0   |   1   |

## Iterative approach with Two Rows
To reduce the space complexity of the full matrix approach, this method maintains only two rows at any time: the previous row and the current row being calculated.

### How it works
As the algorithm progresses, it updates these two rows iteratively. This approach significantly reduces memory usage while retaining the efficiency of the full matrix method.

## Iterative approach with One Row
Further optimizing for space, this method requires only a single row and two variables to store intermediate calculations. It leverages the fact that to calculate the current cell's value, only three adjacent values are needed:

- The cell directly above (prevAbove)
- The cell to the left (row[i-1])
- The diagonal cell (prevDiagonal)

### How it works
This approach iteratively updates the row, using additional variables to track necessary values from the previous iteration. This results in the most space-efficient calculation of the Levenshtein distance.

## Benchmarking
In order to run tests use:
```bash
go test -bench .
```

### Results
We use different length of input data in order to analyze the performance of the algorithms on different string lengths.

**Extra short string (0 symbols)**

![Extra short](./img/extrashort.png)

**Very short string (1 symbols)**

![Very short](./img/veryshort.png)

**Short string (3 symbols)**

![Short](./img/short.png)

**Medium string (6 symbols)**

![Medium](./img/medium.png)

**Long string (12 symbols)**

![Long](./img/long.png)

**Very long string (24 symbols)**

![Long](./img/verylong.png)