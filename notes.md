# GoTour 
Current progress --> https://tour.golang.org/flowcontrol/1

## Random
By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

Note: The environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number.

To see a different number, seed the number generator; see rand.Seed. Time is constant in the playground, so you will need to use something else as the seed.

``` go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

## Named return values
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

``` go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

## Short variable declarations
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

``` go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

```

## Basic types
Go's basic types are

``` go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type


# 002. 
go mod vendor - to get src code for deps

# Questions 

- When shift operators could be used? Provide some real life use cases https://stackoverflow.com/questions/5801008/go-and-operators/5801065


# Strings 
[src](https://www.bogotobogo.com/GoLang/GoLang_byte_and_rune.php)
byte and rune
Golang has integer types called byte and rune that are aliases for uint8 and int32 data types, respectively.

In Go, the byte and rune data types are used to distinguish characters from integer values.

In Go, there is no char data type. It uses byte and rune to represent character values.

The byte data type represents ASCII characters while the rune data type represents a more broader set of Unicode characters that are encoded in UTF-8 format.



```
fmt.Println(reflect.TypeOf('0'))
```
