package kata
import (
	"math"
)
func ListSquared(m, n int)  [][]int {
	var  p float64 =0
	var res [][]int=make([][]int, 0)
	for x:=m;x<n+1;x++{
		p=0
		for i:=1;i<x+1;i++{
			if x%i==0{
				p+=float64(i*i)
			}
		}
if math.Sqrt(p)-math.Round(math.Sqrt(p))==0{
res=append(res,[]int{x,int(p)})
}
	}
return res
}
