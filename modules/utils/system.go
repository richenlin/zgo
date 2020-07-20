package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/NebulousLabs/fastrand"
)

// UUID uuid
func UUID(length int64) string {
	ele := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	ele, _ = Random(ele)
	uuid := ""
	var i int64
	for i = 0; i < length; i++ {
		uuid += ele[fastrand.Intn(59)]
	}
	return uuid
}

// Random random
func Random(strings []string) ([]string, error) {
	for i := len(strings) - 1; i > 0; i-- {
		num := fastrand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := make([]string, 0)
	for i := 0; i < len(strings); i++ {
		str = append(str, strings[i])
	}
	return str, nil
}

// CompressedContent compressed comtent
func CompressedContent(h *template.HTML) {
	st := strings.Split(string(*h), "\n")
	var ss []string
	for i := 0; i < len(st); i++ {
		st[i] = strings.TrimSpace(st[i])
		if st[i] != "" {
			ss = append(ss, st[i])
		}
	}
	*h = template.HTML(strings.Join(ss, "\n"))
}

// ReplaceNth replace nth
func ReplaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

// InArray in array
func InArray(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// WrapURL warp url
func WrapURL(u string) string {
	uarr := strings.Split(u, "?")
	if len(uarr) < 2 {
		return url.QueryEscape(strings.Replace(u, "/", "_", -1))
	}
	v, err := url.ParseQuery(uarr[1])
	if err != nil {
		return url.QueryEscape(strings.Replace(u, "/", "_", -1))
	}
	return url.QueryEscape(strings.Replace(uarr[0], "/", "_", -1)) + "?" +
		strings.Replace(v.Encode(), "%7B%7B.Id%7D%7D", "{{.Id}}", -1)
}

// JSON json
func JSON(a interface{}) string {
	if a == nil {
		return ""
	}
	b, _ := json.Marshal(a)
	return string(b)
}

// ParseBool parse bool
func ParseBool(s string) bool {
	b1, _ := strconv.ParseBool(s)
	return b1
}

// ParseFloat32 parse float
func ParseFloat32(f string) float32 {
	s, _ := strconv.ParseFloat(f, 32)
	return float32(s)
}

// IsJSON is json
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// CopyMap copy map
func CopyMap(m map[string]string) map[string]string {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		panic(err)
	}
	var cm map[string]string
	err = dec.Decode(&cm)
	if err != nil {
		panic(err)
	}
	return cm
}

// CompareVersion compare version
func CompareVersion(src, toCompare string) bool {
	if toCompare == "" {
		return false
	}

	exp, _ := regexp.Compile(`-(.*)`)
	src = exp.ReplaceAllString(src, "")
	toCompare = exp.ReplaceAllString(toCompare, "")

	srcs := strings.Split(src, "v")
	srcArr := strings.Split(srcs[1], ".")
	op := ">"
	srcs[0] = strings.TrimSpace(srcs[0])
	if InArray([]string{">=", "<=", "=", ">", "<"}, srcs[0]) {
		op = srcs[0]
	}

	toCompare = strings.Replace(toCompare, "v", "", -1)

	if op == "=" {
		return srcs[1] == toCompare
	}

	if srcs[1] == toCompare && (op == "<=" || op == ">=") {
		return true
	}

	toCompareArr := strings.Split(strings.Replace(toCompare, "v", "", -1), ".")
	for i := 0; i < len(srcArr); i++ {
		v, err := strconv.Atoi(srcArr[i])
		if err != nil {
			return false
		}
		vv, err := strconv.Atoi(toCompareArr[i])
		if err != nil {
			return false
		}
		switch op {
		case ">", ">=":
			if v < vv {
				return true
			} else if v > vv {
				return false
			} else {
				continue
			}
		case "<", "<=":
			if v > vv {
				return true
			} else if v < vv {
				return false
			} else {
				continue
			}
		}
	}

	return false
}

const (
	// Byte byte
	Byte = 1
	// KByte kbyte
	KByte = Byte * 1024
	// MByte mbyte
	MByte = KByte * 1024
	// GByte gbyte
	GByte = MByte * 1024
	// TByte tbyte
	TByte = GByte * 1024
	// PByte pbyte
	PByte = TByte * 1024
	// EByte ebyte
	EByte = PByte * 1024
)

var bytesSizeTable = map[string]uint64{
	"b":  Byte,
	"kb": KByte,
	"mb": MByte,
	"gb": GByte,
	"tb": TByte,
	"pb": PByte,
	"eb": EByte,
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+" %s", val, suffix)
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1024, sizes)
}

// FileExist file exist
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// TimeSincePro calculates the time interval and generate full user-friendly string.
func TimeSincePro(then time.Time, m map[string]string) string {
	now := time.Now()
	diff := now.Unix() - then.Unix()

	if then.After(now) {
		return "future"
	}

	var timeStr, diffStr string
	for {
		if diff == 0 {
			break
		}

		diff, diffStr = computeTimeDiff(diff, m)
		timeStr += ", " + diffStr
	}
	return strings.TrimPrefix(timeStr, ", ")
}

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

func computeTimeDiff(diff int64, m map[string]string) (int64, string) {
	diffStr := ""
	switch {
	case diff <= 0:
		diff = 0
		diffStr = "now"
	case diff < 2:
		diff = 0
		diffStr = "1 " + m["second"]
	case diff < 1*Minute:
		diffStr = fmt.Sprintf("%d "+m["seconds"], diff)
		diff = 0

	case diff < 2*Minute:
		diff -= 1 * Minute
		diffStr = "1 " + m["minute"]
	case diff < 1*Hour:
		diffStr = fmt.Sprintf("%d "+m["minutes"], diff/Minute)
		diff -= diff / Minute * Minute

	case diff < 2*Hour:
		diff -= 1 * Hour
		diffStr = "1 " + m["hour"]
	case diff < 1*Day:
		diffStr = fmt.Sprintf("%d "+m["hours"], diff/Hour)
		diff -= diff / Hour * Hour

	case diff < 2*Day:
		diff -= 1 * Day
		diffStr = "1 " + m["day"]
	case diff < 1*Week:
		diffStr = fmt.Sprintf("%d "+m["days"], diff/Day)
		diff -= diff / Day * Day

	case diff < 2*Week:
		diff -= 1 * Week
		diffStr = "1 " + m["week"]
	case diff < 1*Month:
		diffStr = fmt.Sprintf("%d "+m["weeks"], diff/Week)
		diff -= diff / Week * Week

	case diff < 2*Month:
		diff -= 1 * Month
		diffStr = "1 " + m["month"]
	case diff < 1*Year:
		diffStr = fmt.Sprintf("%d "+m["months"], diff/Month)
		diff -= diff / Month * Month

	case diff < 2*Year:
		diff -= 1 * Year
		diffStr = "1 " + m["year"]
	default:
		diffStr = fmt.Sprintf("%d "+m["years"], diff/Year)
		diff = 0
	}
	return diff, diffStr
}
