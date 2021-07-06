package main

// Config in another package will not work
type Config struct {
	_    [0]int
	Name string
	Size int
}

func main() {
	_ = Config{[0]int{}, "bar", 123} // this line not work
	_ = Config{Name: "bar", Size: 123}
}
