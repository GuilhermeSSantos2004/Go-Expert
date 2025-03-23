package main

func main() {
	for i := 0; i < 10; i++ {
	 	println(i)
	}

	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		println(k, v)
	}

	fruit := []string{"Aplle", "Banana", "orange"}
	for k, _ := range fruit { // '_' => id blank identifier 
		println(k)
	}


	i := 0 // it's simmillar a While
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Hello, World!")  //Infinite loop (Laço infinito)
	}
}