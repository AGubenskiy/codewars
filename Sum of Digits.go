/*
In this kata, you must create a digital root function.

A digital root is the recursive sum of all the digits in a number. Given n, take the sum of the digits of n. If that value has more than one digit, continue reducing in this way until a single-digit number is produced. This is only applicable to the natural numbers.
*/
import (
	"os"
	"strconv"
)
func DigitalRoot(n int) int {
	for n>9{
		var r int=0
		nstring:=strconv.Itoa(n)
		for i:=0;i<len(nstring);i++{
			x, _ :=strconv.Atoi(nstring[i:i+1])
			r=r+x
		}
		n=r
	}
	return n
}
