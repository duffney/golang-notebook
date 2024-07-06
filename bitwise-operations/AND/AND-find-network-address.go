package main

import (
	"net"
	"fmt"
	"strconv"
)

func main() {
	ipBytes := []byte{192,168,1,18}
	subnetBytes := []byte{255,255,255,0}

	binaryIp := ""
	for _, byte := range ipBytes {
		binaryIp += fmt.Sprintf("%08b", byte)
	}
	
	fmt.Println("IP address:", binaryIp)

	binaryMask := ""
	for _, byte := range subnetBytes {
		binaryMask += fmt.Sprintf("%08b", byte)
	}

	fmt.Println("Subnet Mask:", binaryMask)

	binaryNetwork := ""
	for i := 0; i < 32; i++ {
		if binaryIp[i] == '1' && binaryMask[i] == '1' {
			binaryNetwork += "1"
		} else {
			binaryNetwork += "0"
		}
	}


	fmt.Println("NetworkAddress:" , binaryNetwork)
	
	left := 0
	right := 8
	var octets[]byte

	for left < right {
		oct, _ := strconv.ParseInt(binaryNetwork[left:right], 2, 64)
		octets = append(octets, byte(oct))
		left = right
		if right < 24 {
			right = right+8
		} else {
			oct,_ := strconv.ParseInt(binaryNetwork[right:], 2, 64)
			octets = append(octets, byte(oct))
		}
	}

	netAddr :=  net.IP(octets)
	fmt.Println("Network with Net:", netAddr.String())
}
