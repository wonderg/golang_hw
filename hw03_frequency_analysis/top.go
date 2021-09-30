package main

//package hw03frequencyanalysis
import (
	"fmt"
	"strings"
)



func Top10(s string) []string {
	countedList := []struct{
		word string
		counter int
	}
	list := strings.Fields(s)
	for q := 0; q < len(list)-1; q++ {
		if list[q] == list[q+1] {
			fmt.Println("s:", list[q])
		}
	}
	return nil
}

func main() {
	s := "cat and dog, one dog,two cats and one man"
	Top10(s)

}
