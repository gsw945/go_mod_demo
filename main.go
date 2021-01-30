package main

import "rsc.io/quote"

func Hello() string {
    // return "Hello, world."
    return quote.Hello()
}


func main() {
	Hello()
}