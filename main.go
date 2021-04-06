// main.go

package main

func main() {
	a := App{}
	// You need to set your Username and Password here
	a.Initialize("root", "", "learn_golang")

	a.Run(":8080")
}
