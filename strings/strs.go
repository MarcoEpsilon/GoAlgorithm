package strings

/*
暴力查找子串位置
*/
func NormalSubString(pattern, text string) int {
	i, j := 0, 0
	for ; i < len(text) && j < len(pattern); {
		if text[i] == pattern[j] {
			i++;j++
		} else {
			i = i - j + 1;j = 0
		}
	}
	if j == len(pattern) {
		return i - len(pattern)
	} else {
		return -1
	}
}

/*
KMP求Next数组
*/
func Next(pattern string) (next []int) {
	if len(pattern) == 0 {
		return nil
	} else if len(pattern) == 1 {
		return make([]int, 1)
	}
	next = make([]int, len(pattern))
	next[0] = 0
	next[1] = 0
	j := 0 // 回退点
	i := 1 // 文本扫描指针
	for ; i < len(pattern) - 1; {
		if pattern[i] == pattern[j] {
			j++; i++; next[i] = j
		} else {
			//不能再回退
			if j == 0 {
				i++;next[i] = 0
			}
			j = next[j]
		}
	}
	return next
}

func KmpSubString(pattern, text string) int {
	if len(pattern) == 0 {
		return -1
	}
	i := 0 //文本指针
	j := 0 // next数组指针
	next := Next(pattern)
	for ; i < len(text) && j < len(pattern); {
		if text[i] == pattern[j] {
			i++;j++
		} else {
			if j == 0 { i++ }
			j = next[j]
		}
	}
	if j == len(pattern) {
		return i - j
	}
	return -1
}

func expectBeforeNotExistElement(index int, text string) (s string) {
	// unchecked index
	runes := make([]rune, 0)
	for i, r := range text {
		if i != index {
			runes = append(runes, r)
		}
	}
	return string(runes)
}

func Permutation(text string) []string {
	if len(text) == 0 {
		return nil
	} else if len(text) == 1 {
		return []string{ text }
	}
	ret := make([]string, 0)
	for i, r := range text {
		strs := expectBeforeNotExistElement(i, text)
		after := Permutation(strs)
		for _, v := range after {
			rus := append(append(make([]rune, 0), r), []rune(v)...)
			ret = append(ret, string(rus))
		}
	}
	return ret
}

/*
动态规划版本
*/
func MaxCommonSubString(left, right string) string {
	lookup := make([][]int, len(left) + 1)
	for i, _ := range lookup {
		lookup[i] = make([]int, len(right) + 1)
	}
	maxLength := 0
	maxIndex := 0
	for i := 1; i <= len(left); i++ {
		for j := 1; j <= len(right); j++ {
			if left[i - 1] == right[j - 1] {
				lookup[i][j] = lookup[i - 1][j - 1] + 1
				if lookup[i][j] > maxLength {
					maxLength = lookup[i][j]
					maxIndex = j
				}
			} else {
				lookup[i][j] = 0
			}
		}
	}
	if maxLength == 0 {
		return ""
	} else {
		return right[maxIndex - maxLength:maxIndex]
	}
}