package main

import (
	"fmt"
	"math"
	"sort"
)

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9。因为 nums[0] + nums[1] = 2 + 7 = 9，所以返回 [0, 1]
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok {
			return []int{m[target-nums[i]], i}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}

/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	count := len(nums)
	ret := make([][]int, 0)
	for i := 0; i < count-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l := i + 1
		r := count - 1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l-1] == nums[l] {
					l++
				}
				r--
				for l < r && nums[r+1] == nums[r] {
					r--
				}
			}
		}
	}
	return ret
}

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在唯一答案。
例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 */
func threeSumClosest(nums []int, target int) int {
	closest := math.MaxInt32
	count := len(nums)
	sort.Ints(nums)
	for i := 0; i < count-2; i++ {
		l := i + 1
		r := count - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < target {
				if target-sum < abs(target, closest) {
					closest = sum
				}
				l++
			} else if sum > target {
				if sum-target < abs(target, closest) {
					closest = sum
				}
				r--
			} else {
				return target
			}
		}
	}
	return closest
}

func abs(a, b int) int {
	c := a - b
	if c < 0 {
		return -1 * c
	}
	return c
}

/**
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？
找出所有满足条件且不重复的四元组。

注意：
答案中不可以包含重复的四元组。

示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
 */
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	count := len(nums)
	for i := 0; i < count-3; i++ {
		for j := i + 1; j < count-2; j++ {
			tmpTarget := target - nums[i] - nums[j]
			l := j + 1
			r := count - 1
			for l < r {
				sum := nums[l] + nums[r]
				if sum < tmpTarget {
					l++
				} else if sum > tmpTarget {
					r--
				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					r--
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}
			// 注意去重
			for j < count-2 && nums[j] == nums[j+1] {
				j++
			}
		}
		// 注意去重
		for i < count-3 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}

/**
给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。

为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。

例如:

输入:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

输出:
2

解释:
两个元组如下:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
 */
func fourSumCount(A []int, B []int, C []int, D []int) int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	length := len(A)
	ret := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			sum := A[i] + B[j]
			m1[sum]++
			sum = C[i] + D[j]
			m2[sum]++
		}
	}
	for key, val := range m1 {
		if _, ok := m2[0-key]; ok {
			ret += val * m2[0-key]
		}
	}
	return ret
}

func main() {
	fmt.Println(fourSumCount([]int{-1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}
