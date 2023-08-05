package main

 //vucut kitle endeksi
 // formul :vkı= kg/(boy*boy)
 import (
 	"bufio"
 	"fmt"
 	"os"
 	"strconv"
 )

 func main() {
 	scanner := bufio.NewScanner(os.Stdin)

 	fmt.Print("kilonuzu giriniz ")
 	scanner.Scan()
 	kilo, _ := strconv.ParseFloat(scanner.Text(), 64)

 	fmt.Print("boyunuzu giriniz: ")
 	scanner.Scan()
 	boy, _ := strconv.ParseFloat(scanner.Text(), 64)

 	vki := kilo / (boy * boy)
 	fmt.Printf("vücüt kitele endeksizni = %f", vki)

 }
