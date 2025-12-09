package memo

import (
	"fmt"
	"sync"

	"golang.org/x/sync/singleflight"
)

// Memoize is a function that takes in a function and returns a version of the same function that memoizes previous results
func Memoize[T comparable, R any](fn func(T) R) func(T) R {
	cache := make(map[T]R)

	return func(arg T) R {
		if val, ok := cache[arg]; ok {
			return val
		}
		result := fn(arg)
		cache[arg] = result
		return result
	}
}

// Memoize is a function that takes in a recursive function and returns a version of the same function that memoizes previous results
func MemoizeRecursive[T comparable, R any](fn func(self func(T) R, a T) R) func(T) R {
	cache := map[T]R{}

	var memoized func(T) R
	memoized = func(a T) R {
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

// MemoizeSingleFlight returns a memoized version of the input function that is safe for use in go routines
func MemoizeSingleFlight[T comparable, R any](fn func(T) R) func(T) R {
	var (
		cache = make(map[T]R)
		mu    sync.RWMutex
		sf    singleflight.Group
	)

	return func(arg T) R {
		mu.RLock()
		if v, ok := cache[arg]; ok {
			mu.RUnlock()
			return v
		}
		mu.RUnlock()

		// Prevent duplicate work for concurrent requests
		val, _, _ := sf.Do(
			stringify(arg),
			func() (any, error) {
				res := fn(arg)

				mu.Lock()
				cache[arg] = res
				mu.Unlock()

				return res, nil
			},
		)

		return val.(R)
	}
}

// MemoizeSingleFlight returns a memoized version of the (recursive) input function that is safe for use in go routines
func MemoizeSingleFlightRecursive[T comparable, R any](fn func(self func(T) R, a T) R) func(T) R {
	var (
		cache = make(map[T]R)
		mu    sync.RWMutex
		sf    singleflight.Group
	)
	var memoized func(T) R

	memoized = func(a T) R {
		mu.RLock()
		if v, ok := cache[a]; ok {
			mu.RUnlock()
			return v
		}
		mu.RUnlock()

		// Prevent duplicate work for concurrent requests
		val, _, _ := sf.Do(
			stringify(a),
			func() (any, error) {
				res := fn(memoized, a)

				mu.Lock()
				cache[a] = res
				mu.Unlock()

				return res, nil
			},
		)

		return val.(R)
	}
	
	return memoized
}

func stringify[T comparable](v T) string {
	return fmt.Sprintf("%v", v)
}