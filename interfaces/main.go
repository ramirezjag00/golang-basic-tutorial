package main

import (
	"fmt"
)

type dog interface {
        bark() string
        walk() string
        anything() string
}

func main() {
        fmt.Println(walkOrBark(2))
}

type dogs struct {
        action string
        sound string
        what string
}

func (d *dogs) walk() string {
        return d.action
}

func (d *dogs) bark() string {
        return d.sound
}

func (d *dogs) anything() string {
        return d.what
}

func walkOrBark(num int) string {
        var d dog
        a := dogs{"Walk walk walk", "Arf arf arf", "Nothing to do here"}
        d = &a

        if num == 1 {
                return d.walk()
        } else if num == 2 {
                return d.bark()
        } else {
                return d.anything()
        }
}
