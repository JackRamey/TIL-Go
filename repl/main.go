package main

import "fmt"

type Object struct {
	ID string
}

func main() {
	objects := []*Object{nil, {ID: "1"}}
	fmt.Println("objects: ", objects)

	objects = []*Object{}
	objects = append(objects, nil)
	objects = append(objects, &Object{ID: "1"})
	fmt.Println("objects: ", objects)

	var os []*Object
	fmt.Println("objects: ", os)
}
