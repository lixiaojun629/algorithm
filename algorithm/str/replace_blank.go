package str

func replaceSpace(s string) string {
	targetLen := len(s)
	for _, c := range s {
		if c == ' ' {
			targetLen += 2
		}
	}
	target := make([]byte, targetLen)
	for i, j := targetLen-1, len(s)-1; i >= 0; {
		if s[j] == ' ' {
			target[i] = '0'
			target[i-1] = '2'
			target[i-2] = '%'
			i -= 2
		} else {
			target[i] = s[j]
		}
		i--
		j--
	}
	return string(target)
}
