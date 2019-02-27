package find_deepest_triplet

import (
	"testing"
)

func TestLenLessThanThree(t *testing.T) {
	t.Parallel()
	t.Run("Test with zero length array", func(t *testing.T) {
		s := NewSolution()
		answer := s.solution([]int{})

		if answer != -1 {
			t.Errorf("Zero length array should return -1, but got %d\n", answer)
			return
		}
	})

	t.Run("Test with one length array", func(t *testing.T) {
		s := NewSolution()
		answer := s.solution([]int{1})

		if answer != -1 {
			t.Errorf("One length array should return -1, but got %d\n", answer)
			return
		}
	})

	t.Run("Test with two length array", func(t *testing.T) {
		s := NewSolution()
		answer := s.solution([]int{2, 1})

		if answer != -1 {
			t.Errorf("Two length array should return -1, but got %d\n", answer)
			return
		}
	})
}

func TestSameNumbers(t *testing.T) {
	s := NewSolution()

	t.Run("Test with same numbers array", func(t *testing.T) {
		answer := s.solution([]int{1, 1, 1, 1, 1, 1, 1})
		if answer != -1 {
			t.Errorf("Same numbers array should return -1, but got %d\n", answer)
			return
		}
	})
}

func TestNoPits(t *testing.T) {
	t.Run("Test with no pits", func(t *testing.T) {
		s := NewSolution()
		answer := s.solution([]int{1, 2, 3, 4, 3, 2, 1})

		if answer != -1 {
			t.Errorf("Array with no pits should return -1, but got %d\n", answer)
			return
		}
	})
}

func TestLargeArray(t *testing.T) {
	s := NewSolution()

	t.Run("Test with large array", func(t *testing.T) {
		question := make([]int, 1000000)
		base := 100000000
		var min, max int

		for i := range question {
			if i <= len(question)/2 {
				question[i] = base - i*200
			} else {
				question[i] = (i + 1) * 100
			}

			if i == 0 {
				max = question[i]
				min = question[i]
			} else {
				if question[i] < min {
					min = question[i]
				}

				if question[i] > max {
					max = question[i]
				}
			}
		}

		answer := s.solution(question)

		if answer != max-min {
			t.Errorf("Large array should return %d, but got %d\n", max-min, answer)
			return
		}
	})
}

func TestSolutions(t *testing.T) {
	s := NewSolution()
	var answer int

	t.Run("Test with different array", func(t *testing.T) {
		answer = s.solution([]int{0, 1, 3, -2, 0, 1, 0, -3, 2, 3})
		if answer != 4 {
			t.Errorf("Should return 4, but got %d\n", answer)
			return
		}

		answer = s.solution([]int{5, 3, 1, -2, 0, 1, 10, -10, 2, 3})
		if answer != 13 {
			t.Errorf("Should return 13, but got %d\n", answer)
			return
		}

		answer = s.solution([]int{5, 3, 1, -2, 0, 1, 10, -10, 2})
		if answer != 12 {
			t.Errorf("Should return 12, but got %d\n", answer)
			return
		}

		answer = s.solution([]int{10, 8, 7, 0, -1, 20, -2, 0, 1, 0, -12})
		if answer != 11 {
			t.Errorf("Should return 11, but got %d\n", answer)
			return
		}

		answer = s.solution([]int{10, 8, 7, 0, -1, 20, -2, 0, 1, 0, -12, 5})
		if answer != 13 {
			t.Errorf("Should return 13, but got %d\n", answer)
			return
		}

		answer = s.solution([]int{10, 10, 8, 7, 0, -1, -1, -1, 20, 20, 20, -2, 0, 1, 1, 0, -12, -12, 5, 5})
		if answer != 13 {
			t.Errorf("Should return 13, but got %d\n", answer)
			return
		}

		answer = s.solution([]int{10, 8, 7, 0, -1, 20, -2, 0, 1, 0, -8, 5})
		if answer != 11 {
			t.Errorf("Should return 11, but got %d\n", answer)
			return
		}
	})
}
