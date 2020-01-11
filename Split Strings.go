func Solution(str string) []string {
var result []string
	input:=str
    if len(input)%2>0 {
	    input=input+"_"
    }
for i:=0; i<len(input)-1;i=i+2{
	result=append(result,input[i:i+2])
}
return result
}
