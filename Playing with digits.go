package kata
import (
	"math"
	"strconv"
)
func DigPow(n, p int) int {
var pr float64 =0
var t int
var str string
str=strconv.Itoa(n)
for i:=0;i<len(str);i++{
t, _ =strconv.Atoi(str[i:i+1])
	pr+=math.Pow(float64(t),float64(p+i))
}
if int(pr)%n>1 {return -1}else{
return int(pr)/n}
}
