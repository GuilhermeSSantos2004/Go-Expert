package main

func main() {
	a := 1
	b := 2
	c := 3

	if a < b {
		println(b)
	} else {
		println(a)
	}

	if a > b && c > a {
		print(" > b && c > a")
	}

	if a > b || c > a {
		print(" > b || c > a")
	}


	switch a {
		case 1:
			print('a')
		case 2:
			print('b')
		case 3: 
			print('c')
		default:
			print("Default: D") 
	}

}
