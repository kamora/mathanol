package mathanol

import "iter"

func Rem[T Unsigned](value T, div T) iter.Seq[T] {
	return func(yield func(T) bool) {

		for value > 0 {
			rem := value % div

			if !yield(rem) {
				return
			}

			value = (value - rem) / div
		}
	}
}

func Clamp[T Number](v T, low T, high T) T {
	if v < low {
		return low
	}

	if v > high {
		return high
	}

	return v
}
