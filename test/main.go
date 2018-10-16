package main

type A struct {
	b B
}

type B interface {
	Get()
	Set()
}

type C struct {
}

func (c *C) Get() {

}

func (c *C) Set() {

}
func NewA() (a A) {
	c := C{}
	return c
}

func main() {

}
