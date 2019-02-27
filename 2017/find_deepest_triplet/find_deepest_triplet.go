package find_deepest_triplet

/*
 * A non-empty zero-indexed array A consiting of N integers is given.
 * A pit in this array is any triplet of integers (P, Q, R) such that:
 *
 *  1. 0 <= P < Q < R < N;
 *  2. sequence [A[P], A[P+1], ..., A[Q]] is strictly decreasing,
 *     i.e. A[P] > A[P+1] > ... > A[Q];
 *  3. sequence A[Q], A[Q+1], ..., A[R] is strictly increasing,
 *     i.e. A[Q] < A[Q+1] < ... < A[R].
 *
 * The depth of a pit (P, Q, R) is the number min{A[P] - A[Q], A[R] - A[Q]}.
 *
 * For example, consider array A consisting of 10 elements such that:
 *
 * 	A[0] = 0
 * 	A[1] = 1
 * 	A[2] = 3
 * 	A[3] = -2
 * 	A[4] = 0
 * 	A[5] = 1
 * 	A[6] = 0
 * 	A[7] = -3
 * 	A[8] = 2
 * 	A[9] = 3
 *
 * Triplet (2, 3, 4) is one of pits in this array,
 * because sequence [A[2], A[3]] is strictly decreasing (3 > -2)
 * and sequence [A[3], A[4]] is strictly increasing (-2 < 0).
 * Its depth is min{A[2] - A[3], A[4] - A[3]} = 2.
 * Triplet (2, 3, 4) is another pit with depth 3.
 * Triplet (5, 7, 8) is yet another pit with depth 4.
 * There is no pit in this array deeper (i.e. having depth greater) than 4.
 *
 * Write a function:
 *
 * 	class Solution { public int solution(int[] A); }
 *
 * that, given a non-empty zero-indexed array A consisting of N integers,
 * returns the depth of the deepest pit in array A. The function should return -1
 * if there are no pits in array A.
 *
 * For example, consider array A consisting of 10 elements such that:
 *
 * 	A[0] = 0
 * 	A[1] = 1
 * 	A[2] = 3
 * 	A[3] = -2
 * 	A[4] = 0
 * 	A[5] = 1
 * 	A[6] = 0
 * 	A[7] = -3
 * 	A[8] = 2
 * 	A[9] = 3
 *
 * the function should return 4, as explained above.
 *
 * Assume that:
 * 	1. N is an integer withing the range [1..1,000,000];
 * 	2. each element of array A is an integer withing the range
 *     [[-100,00,000..100,000,000]].
 *
 * Complexity:
 * 	1. expected worst-case time complexity is O(N);
 * 	2. expected worst-case space complexity is O(N); beyond input storage
 *     (not counting the storage required for input arguments).
 *
 * Elements of input arrays can be modified.
 */

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

type Solution struct {
	maxDepth int
}

func NewSolution() *Solution {
	return &Solution{maxDepth: -1}
}

func (s *Solution) solution(A []int) int {
	// Not a validate pit
	if len(A) < 3 {
		return -1
	}

	s.maxDepth = -1

	var p, q, r int
	findP := true
	findQ := false
	findR := false

	p = MinInt
	q = MaxInt
	r = MinInt

	for i := 0; i < len(A); i++ {
		if findP {
			if A[i] >= p {
				// Still an increasing sequence, update p.
				p = A[i]
				continue
			}

			// A[i] < p, it's the start of a decreasing sequence.
			// p is found, start finding q.
			q = A[i]
			findP = false
			findQ = true
		} else if findQ {
			if A[i] <= q {
				// Still a decreasing sequence, update q.
				q = A[i]
				continue
			}

			// A[i] > q, it's the end of decreasing sequence,
			// (start of a increasing sequence)
			// q is found, start finding r.
			r = A[i]

			if i == len(A)-1 {
				// Reached the end of array, r must be current element.
				// Update max depth.
				s.updateMaxDepth(p, q, r)
			} else {
				findQ = false
				findR = true
			}
		} else if findR {
			if A[i] >= r {
				// Still an increasing sequence, update r.
				r = A[i]
				if i != len(A)-1 {
					// Only continue updating r if not reaching
					// the end of array yet.
					continue
				}
			}

			// Can be one of following conditions:
			// 	1. A[i] < r, it's the end of increasing sequence.
			// 	2. Has reached the end of array.
			// p, q, r are all been found, update max depth.
			s.updateMaxDepth(p, q, r)

			// Start finding next pit.
			p = r
			q = A[i]
			r = MinInt
			findQ = true
			findR = false
		}
	}

	return s.maxDepth
}

func (s *Solution) updateMaxDepth(p, q, r int) {
	pqDepth := p - q
	qrDepth := r - q
	var pqrMinDepth int

	if pqDepth < qrDepth {
		pqrMinDepth = pqDepth
	} else {
		pqrMinDepth = qrDepth
	}

	if s.maxDepth < pqrMinDepth {
		s.maxDepth = pqrMinDepth
	}
}
