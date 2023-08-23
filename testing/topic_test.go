package testing

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	nums := []int{5, 3, 8, 4, 2, 7}
	nums = QuickSort(nums)
	assert.Equal(t, nums, []int{2, 3, 4, 5, 7, 8})
}

func BenchmarkQuickSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		// 生成随机数组
		nums := rand.Perm(1000000)

		b.StartTimer()
		QuickSort(nums)
		b.StopTimer()
	}
}
