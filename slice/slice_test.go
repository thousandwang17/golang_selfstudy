package test

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

// use origin slice address
func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

// make new address for new slice
func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024) // 1M
		ans = append(ans, f(origin))
	}
	printMem(t)
	_ = ans
}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generateWithOutCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func printMem(t *testing.T) {
	t.Helper()
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}

func BenchmarkGenerateWithCap(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generateWithCap(t.N)
	}
}

func BenchmarkGenerateWithOutCap(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generateWithOutCap(t.N)
	}
}

func TestLastCharsBySlice(t *testing.T) { testLastChars(t, lastNumsBySlice) }
func TestLastCharsByCopy(t *testing.T)  { testLastChars(t, lastNumsByCopy) }
