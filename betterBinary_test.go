package betterBinary

import(
	"testing"
	"fmt"
)

func TestBinbuf(t *testing.T){
	var x uint64 = 0xaabbccdd00112233
	var c =make([]byte, 8)
	Binbuf(x, c)
	fmt.Printf("% x\n",c)
}

