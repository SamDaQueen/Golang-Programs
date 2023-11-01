package main

import (
	"fmt"
	"time"
)

/* Race conditions

When programs run concurrently, there is interleaving involved. This
interleaving can be more granular than just source code, but at machine code
granularity, which can cause inconsistent results. The reading from a register,
performing an arithmetic calculation, as well as saving to memory can be in a
different order for multiple variables.

This causes the output of the program to be inconsistent as it depends on
the interleaving, which is not controlled by the programmer. This issue is
as known as race conditions. We have to ensure that the output does not depend
on the non-deterministic ordering of the machine code.

In the example below, the value of x within the goroutines is often printed
incorrectly as the calculation from the other goroutine occurs before the print
statement. The order in which the goroutines are executed also changes the final
output. Even though routine2() is called after routine1() in source code, there
are times when the the arithmetic instruction of routine 2 executes before
routine 1, giving the output 12.

*/

var x int = 2

func main() {
	go routine1()
	go routine2()
	time.Sleep(time.Millisecond)
	// will print 8 if order is r1, r2 and 12 otherwise
	fmt.Printf("Output: %d\n", x)
}

func routine1() {
	x = x * 3
	println("r1", x)
}

func routine2() {
	x = x + 2
	println("r2", x)
}