package kata

func Evaporator(content float64, evapPerDay int, threshold int) int {
    // your code
    momntvol:=content
	day:=0
	for momntvol>float64(threshold)*content/100 {
		momntvol=float64(100-evapPerDay)*momntvol/100
		day++

}
return day
}
