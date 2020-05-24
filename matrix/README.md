Small tests with matrix inspired by

_Mathematics from the birth of numbers_ By _Jan Gullberg_ pp 646-647

This implements the first method described in the book to calculate the
determinant of a square matrix.

Usage:

```
$ go test ./matrix/ -v
=== RUN   TestDet
    TestDet: matrix_test.go:52: det({3 3 [[2 0 4] [3 1 0] [4 2 2]]}) is 12
    TestDet: matrix_test.go:52: det({3 3 [[-2 2 -3] [-1 1 3] [2 0 -1]]}) is 18
    TestDet: matrix_test.go:52: det({3 3 [[1 2 3] [4 5 6] [7 8 9]]}) is 0
    TestDet: matrix_test.go:52: det({4 4 [[1 3 5 9] [1 3 1 7] [4 3 9 7] [5 2 0 9]]}) is -376
--- PASS: TestDet (0.00s)
PASS
ok      _/home/bhesmans/git/go_play/matrix/matrix       0.003s
```
