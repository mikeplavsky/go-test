package main

import "fmt"
import "unsafe"

func main() {

	str := "A Go variable"
	addr := unsafe.Pointer(&str)
	addr1 := &str

	fmt.Printf("The address of str is 0x%x\n", addr)
	fmt.Printf("The address of str is 0x%x\n", addr1)

	str2 := (*string)(addr)

	fmt.Printf("String constructed from pointer: %s\n", str2)
	fmt.Printf("String constructed from pointer: %s\n", *str2)

	address := uintptr(addr)
	address += 4

	str3 := (*string)(unsafe.Pointer(address))
	fmt.Printf("The address of stri3 is 0x%x\n", address)

	fmt.Printf("String constructed  from pointer: %s\n", *str3)

}
