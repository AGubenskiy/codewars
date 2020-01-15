import (
	"strings"
)
func ToCamelCase(s string) string {
	r := strings.NewReplacer("_", " ",	"-", " " )
	ss:=strings.Split(r.Replace(s), " ")
	for i:=1;i<len(ss);i++{ss[i]=strings.Title(ss[i])}
	return strings.Join(ss, "")
}
