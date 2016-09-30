package lux

import "fmt"

func (x *Lux) write(f fmt.State) {
	f.Write(escape)
	f.Write([]byte{0x5b}) // "["

	for i := 0; i < len(x.styles); i++ {
		f.Write(attrs[x.styles[i]])
		if i < len(x.styles)-1 {
			f.Write([]byte{0x3b}) // ";"
		}
	}

	f.Write([]byte{0x6d}) // "m"
	f.Write([]byte(x.val))
	f.Write([]byte{0x1b, 0x5b, 0x30, 0x6d}) // "\033[0m"
}
