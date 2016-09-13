package betterBinary

import(
	"testing"
	"fmt"
)

func TestBinbuf(t *testing.T){
	var x uint64 = 0xaabbccdd00112233
	var c [8]byte
	Binbuf(x, &c)
	fmt.Printf("% x\n",c)
}

