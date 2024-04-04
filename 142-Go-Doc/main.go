// Package main to see code in command line
package main

// AddSum  add unlimited number of values of type int
func AddSum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

//go doc - To see the code of documentation

/*Doc (usually run as go doc) accepts zero, one or two arguments.

Zero arguments:

    go doc

Show the documentation for the package in the current directory.

One argument:

    go doc <pkg>
    go doc <sym>[.<methodOrField>]
    go doc [<pkg>.]<sym>[.<methodOrField>]
    go doc [<pkg>.][<sym>.]<methodOrField>
*/

//example
// go doc fmt
