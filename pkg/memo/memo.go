package memo

import "sync"

func Memoize1[T comparable, R any](fn func(T) R) func(T) R {
	cache := make(map[T]R)
	var mu sync.Mutex

	return func(arg T) R {
		mu.Lock()
		defer mu.Unlock()

		if val, ok := cache[arg]; ok {
			return val
		}
		result := fn(arg)
		cache[arg] = result
		return result
	}
}

func MemoRec[R any, A comparable](fn func(self func(A) R, a A) R) func(A) R {
	cache := map[A]R{}

	var memoized func(A) R
	memoized = func(a A) R {
		if v, ok := cache[a]; ok {
			return v
		}

		res := fn(memoized, a)
		cache[a] = res
		return res
	}
	return memoized
}

/*Example
fib := MemoRec(func(f func(int) int, n int) int {
	if n < 2 {
		return n
	}
	return f(n-1) + f(n-2)
})

fmt.Println(fib(45)) */