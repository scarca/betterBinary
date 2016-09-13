package betterBinary

func Binbuf(x uint64, b []byte){
	i:=7
	for x>0{
		b[i] = byte(x)
		x>>=8
		i--
	}
}
