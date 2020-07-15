package blockchain

import (
	"encoding/hex"
	"fmt"
)

func HexDemo(){
	address :="AGc9NrdF5MuMJpkFfZ3MWKa67ds6H2fzud"
	addrBytes, _ :=hex.DecodeString(address)
	format, _ :=fmt.Printf("%x",addrBytes)
	fmt.Println(format)
}

/**
45 as 0010 1101 in binary
0010 = 2 (base 10) = 2 (base 16)
1101 = 13 (base 10) = d (base 16)

Therefore: 45 = 0010 1101 = 0x2d  Byte to hex

 */


