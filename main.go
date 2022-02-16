<<<<<<< HEAD
package main

func main() {
	println(add(100,100))
}

func add(n, add int) int {
	return n+add
}

func multiply(n, multiplyBy int) int {
	result := n
	for i := 1; i < multiplyBy; i++ {
		result = add(result, n)
	}
	
	return result
=======
package main

func main() {
	println(add(100,100))
}

func add(n, add int) int {
	return n+add
}

func multiply(n, multiplyBy int) int {
	result := n
	for i := 1; i < multiplyBy; i++ {
		result = add(result, n)
	}
	return result
>>>>>>> 38e00decfea55bb50fc074c81e078f8e0771f3a5
}