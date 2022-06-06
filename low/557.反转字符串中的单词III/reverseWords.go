package main

func reverseWords(s string) string {
	t := []byte(s)
	i := 0
	l := 0
	r := 0
	for i < len(t) && r+1 < len(t) && l+1 < len(t) {
		if t[l] == ' ' {
			l++
		}

		if t[r] == ' ' || t[r+1] != ' ' {
			r++
		}

		if (t[l] != ' ' && r == len(t)-1) || (t[l] != ' ' && t[r+1] == ' ') {
			for j := 0; j <= (r-l)/2; j++ {
				t[l+j], t[l+(r-l)-j] = t[l+(r-l)-j], t[l+j]
			}
			l = r
			l++
			r++
		}

		i++
	}

	return string(t)
}
