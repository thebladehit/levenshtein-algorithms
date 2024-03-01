# levenshtein-algorithms

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