package main

func main() {
	println(biggerIsGreater("ab"))
	println(biggerIsGreater("bb"))
	println(biggerIsGreater("hefg"))
	println(biggerIsGreater("dhck"))
	println(biggerIsGreater("dkhc"))
}

func biggerIsGreater(w string) string {
	s := []byte(w)
	n := len(s)
	if n <= 1 {
		return "no answer"
	}

	// Step 1: Find the largest index i such that s[i] < s[i+1]
	i := n - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}
	if i < 0 {
		return "no answer"
	}

	// Step 2: Find the smallest j > i such that s[j] > s[i]
	j := n - 1
	for s[j] <= s[i] {
		j--
	}

	// Step 3: Swap s[i] and s[j]
	s[i], s[j] = s[j], s[i]

	// Step 4: Reverse the suffix starting from i+1
	reverse(s[i+1:])

	return string(s)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
