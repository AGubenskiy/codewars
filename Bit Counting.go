package kata
import (
	"strconv"
	"strings"
)

func CountBits(ui uint) int {
s:=strconv.FormatUint(uint64(ui),2)
return len(strings.Replace(s,"0","",-2))
}
