package lengthOfLongestSubstring

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	//i = 0
	var count []int
	c := make(map[int32]string)
	c2 := make(map[int32]string)
	for _, v := range s {
		if c[v] != "" {
			count = append(count, len(c))
			if len(c2) == 0 {
				c2 = make(map[int32]string)
			}
			for k2, v2 := range c {
				if v2 == c[v] {
					c2 = make(map[int32]string)
					continue
				} else {
					c2[k2] = v2
				}
			}
			c2[v] = string(rune(v))

			c = c2
		} else {
			c[v] = string(rune(v))
		}
	}
	count = append(count, len(c))

	return max(count)

}

func max(l []int) (maxc int) {
	maxc = l[0]
	for _, v := range l {
		if v > maxc {
			maxc = v
		}
	}
	return
}

func main() {

	fmt.Println(lengthOfLongestSubstring2("abcabcbb"))

}

func lengthOfLongestSubstring2(s string) int {
	StrSli := make([]string, 0, len(s))
	for _, v := range s {
		StrSli = append(StrSli, string(rune(v)))
	}
	var count []int
	c := make([]string, 0, len(s))
	c2 := make([]string, 0, len(s))
	for _, v := range StrSli {
		if index, isIn := Inslice(v, c); isIn {
			count = append(count, len(c))
			c2 = c[index+1:]
			c2 = append(c2, v)
			c = c2
		} else {
			c = append(c, v)
		}
	}
	count = append(count, len(c))

	return max(count)

}

func Inslice(s string, sli []string) (index int, t bool) {

	for k, v := range sli {
		if s == v {
			return k, true
		} else {
			continue
		}

	}
	return 0, false
}
