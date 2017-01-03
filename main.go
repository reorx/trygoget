package main

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

func main() {
	m := mock.Mock{}
	fmt.Println("vim-go", m)
}
