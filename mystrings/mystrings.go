package mystrings


func Reverse(s string) string {
	result := ""
	
	for _, v := range s {
		result += string(v)
	}
	fmt.Println(result)
	return result
}

Reverse("balls")
