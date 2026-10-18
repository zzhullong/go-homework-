package homework01

import (
	"fmt"
	"sort"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// TODO: implement

	res := 0
	for index ,value := range nums{
		fmt.Printf("序号：%v,值:%v\n",index,value)
		
		res ^= value
	}
	fmt.Printf("返回结果值是:%v",res)
	return res
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// TODO: implement
	
	if x<0 {
		return false
	}

	if x==0 {
		return true
	}
	
	if x>0 {



		silec := []int{}
		
		for x>0 {
			
			
			y := x%10

			silec = append(silec, y)
			x = x/10
		}

		len := len(silec)

		for i := 0; i < len/2; i++ {

			fmt.Printf("序号：%v",i)
			if silec[i] != silec[len-1-i]{

				return false
			}


		}

		return  true
	}
	
	
	return true

	
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// TODO: implement

	    m := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    stack := []rune{}

    for _, char := range s {
        // 如果是右括号
        if left, ok := m[char]; ok {
            // 栈为空 或者 栈顶元素不等于对应的左括号
            if len(stack) == 0 || stack[len(stack)-1] != left {
                return false
            }
            // 弹出栈顶
            stack = stack[:len(stack)-1]
        } else {
            // 左括号，入栈
            stack = append(stack, char)
        }
    }
    // 循环结束栈必须是空的
    return len(stack) == 0                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     



	return false
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// TODO: implement
	if len(strs) == 0 {
		return ""
	}
	// 拿第一个字符串当作参考
	base := strs[0]

	// 遍历基准字符串每一个位置
	for i := range base {
		char := base[i]
		// 和剩下所有字符串对比第i位
		for j := 1; j < len(strs); j++ {
			// 如果这个字符串长度不够i 或者字符不一致，直接返回前面部分
			if i >= len(strs[j]) || strs[j][i] != char {
				return base[:i]
			}
		}
	}
	return base
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// TODO: implement
    for i := len(digits)-1; i >=0; i-- {
        if digits[i] <9 {
            digits[i]++
            return digits
        }
        // 等于9，置0，继续进位
        digits[i] = 0
    }
    // 走到这里说明全部是9，前面补1
    res := make([]int, len(digits)+1)
    res[0] = 1
    return res
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// TODO: implement
		if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// TODO: implement
		if len(intervals) == 0 {
		return [][]int{}
	}
	// 按区间start升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1] // 结果里最后一个区间
		cur := intervals[i]
		if cur[0] <= last[1] {
			// 重叠，合并，更新末尾
			last[1] = max(last[1], cur[1])
		} else {
			// 不重叠，直接加入结果
			res = append(res, cur)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// TODO: implement
	    m := make(map[int]int) // key:数组的值，value:下标
    for idx, v := range nums {
        need := target - v
        if pos, ok := m[need]; ok {
            // 找到了，返回两个下标
            return []int{pos, idx}
        }
        // 没找到，存入map
        m[v] = idx
    }
    // 题目保证一定有解，这里不会走到
    return nil
}
