package world

import (
	"fmt"
)

type Handler struct{}
type Param struct {
	Name string
}

func (this *Handler) A(p *Param) (bool, string) {
	p.Name = "this is a Book "
	fmt.Printf("AAAAAAAAA %s AAAAAAAAAAA\n", p.Name)
	return true, "return message"
}
