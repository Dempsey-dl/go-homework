package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func singleNumber(nums []int) int {
	// var a int
	// for i := range nums {
	// 	a ^= nums[i]
	// }

	// return a

	map1 := make(map[int]int)
	for _, k := range nums {
		map1[k]++
	}
	for i, j := range map1 {
		if j == 1 {
			return i
		}
	}
	return 0
}

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := x
	sum := 0
	for num > 0 {
		sum = sum*10 + num%10
		fmt.Printf("sum %d =  %d + %d\n", sum, sum*10, num%10)
		num /= 10
		fmt.Println("num", num/10)
	}
	return sum == x
}

func isPalindrome1(x int) bool {
	// 特殊情况处理
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		fmt.Printf("revertedNumber(%d)= revertedNumber*10(%d) + x%%10(%d)\n", revertedNumber, revertedNumber*10, x%10)
		revertedNumber = revertedNumber*10 + x%10
		fmt.Printf("x /= 10(%d)\n", x)
		x /= 10

	}
	// 当数字长度为奇数时，可以通过 revertedNumber/10 去掉中间的数字
	return x == revertedNumber || x == revertedNumber/10
}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/
func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}

		}
	}
	return prefix
}

/*
删除排序数组中的重复项
*/
func rmmoveReNum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}

	}
	return slow + 1
}

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}

	}
	return append([]int{1}, digits...)
}

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merge := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := merge[len(merge)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			merge = append(merge, current)
		}

	}
	return merge
}

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
*/
func targetMerge(num []int, target int) []int {
	if len(num) == 0 {
		return []int{0}
	}

	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i]+num[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, v := range nums {
		if j, ok := hashMap[target-v]; ok {
			return []int{j, i}
		}
		hashMap[v] = i
	}
	return nil

}
func main() {
	// nums := []int{1, 2, 3, 4, 5, 5, 1, 2, 3}
	// a := singleNumber(nums)
	// fmt.Println("单一值", a)

	// b := isPalindrome1(23124)
	// fmt.Println("回文数", b)

	// if isValid("(") {
	// 	fmt.Println("valid")
	// } else {
	// 	fmt.Println("invalid")
	// }

	// strs := []string{"str2123", "str1234", "str12345"}
	// fmt.Println(longestCommonPrefix(strs))

	// num := []int{1, 3, 5, 6, 8, 9}
	// fmt.Println(plusOne(num))

	// intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// fmt.Println(merge(intervals))

	// nums := []int{9, 10, 20, 231}
	// fmt.Println(targetMerge(nums, 30))
}
