package stack


// https://leetcode.cn/problems/valid-parentheses/


func isValid(s string) bool {
	slice := []rune{}
	for _, v := range s{
		if v == '{' || v == '(' || v == '[' {
			slice = append(slice, v)
		} else {
			if len(slice) == 0 {
				return false
			}
			left := slice[len(slice) - 1]
			slice = slice[:len(slice) - 1]
			if left == '{' && v != '}' {
				return false
			}
			if left == '(' && v != ')' {
				return false
			}
			if left == '[' && v != ']' {
				return false
			}
		}
	}
	return len(slice) == 0
}


func isValid2(s string) bool {
	slice := []byte{}
	m := map[byte]byte{'{':'}', '(':')', '[':']'}
	for _, c := range s {
		cb := byte(c)
		if _, ok := m[cb]; ok {
			slice = append(slice, cb)
		} else {
			if len(slice) == 0 {
				return false
			}
			left := slice[len(slice) - 1]
			slice = slice[:len(slice) - 1]
			if cb != m[left] {
				return false
			}
		}
	}
	return len(slice) == 0
}