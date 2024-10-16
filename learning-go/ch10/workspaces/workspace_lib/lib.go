package workspace_lib

import (
	"github.com/kesopeso/keso_go_add/v2"
)

func AddNums(a int, b int) int {
	return keso_go_add.Add(a, b)
}

func SubNums(a int, b int) int {
	return a - b
}
