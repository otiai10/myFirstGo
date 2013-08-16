package fizzbuzz

import "fmt"

func Run() {
    fmt.Print("Let's start FizzBuzz!\n")
    for i := 0; i < 30; i++ {
        s := "--"
        if i % (3*5) == 0 {
            s = "FizzBuzz!!"
        } else if i % 5 == 0 {
            s = "Buzz"
        } else if i % 3 == 0 {
            s = "Fizz"
        } else {
            // Do Nothing
        }
        fmt.Printf("%d\t: %s\n", i, s)
    }
}