package main

func main() {
	//s := "1[.]1[.]1[.]1"
	//s := "1.1.1.1"
	s := "255.255.10.1"
	println(defang(s))
}

func defang(s string) string {
	for i, idx := range _defang(s, 0) {
		s = s[:idx+(i*2)] + "[" + s[idx+(i*2):]
		s = s[:idx+2+(i*2)] + "]" + s[idx+2+(i*2):]
	}
	return s
}

func _defang(s string, i int) []int {
	if i == len(s) {
		return []int{}
	}
	if string(s[i]) == "." {
		return append([]int{i}, _defang(s, i+1)...)
	}
	return _defang(s, i+1)
}
