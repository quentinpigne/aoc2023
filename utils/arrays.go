package utils

func Sum(data []int) int {
	sum := 0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	return sum
}

func Max(data []int) int {
	max := 0
	for i := 0; i < len(data); i++ {
		if max < int(data[i]) {
			max = int(data[i])
		}
	}
	return max
}

func StackBy[T any](data []T, stackSize int) [][]T {
	var result [][]T
	for i := 0; i < len(data); i++ {
		if i%stackSize == 0 {
			result = append(result, make([]T, 0, stackSize))
		}
		result[i/stackSize] = append(result[i/stackSize], data[i])
	}
	return result
}

func Map[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))
	for _, e := range data {
		res = append(res, f(e))
	}
	return res
}

func Filter[T any](data []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, e := range data {
		if f(e) {
			res = append(res, e)
		}
	}
	return res
}
