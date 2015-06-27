package main

import (
       "golang.org/x/tour/wc"
       "strings"
)

func WordCount(s string) map[string]int {
     m := make(map[string]int)
     for _,a := range strings.Fields(s) {
     	 m[a] += 1
	 }
	 
	 return m
}

func main() {
     wc.Test(WordCount)
}
