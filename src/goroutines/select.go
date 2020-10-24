package main

import "fmt"

// for {
//     select {
//         case a = <- inputChannel:
//             fmt.println(a)
//         case b <- c:
//             fmt.println(b)
//         case <- abort:
//             return // break the infinite loop
//     }
// }

func main() {
	fmt.Println("Hello World")
}
