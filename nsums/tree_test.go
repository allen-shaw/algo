package nsums

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	root := buildTree([]int{2, 1, 4})
	ans := InorderTravverse(root)
	fmt.Println(ans)

}
