package common

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

// EndPoints a list of endpoints by marketpalceID
var EndPoints = map[string]string{
	"ATVPDKIKX0DER":  "www.amazon.com",    // 美国
	"A1AM78C64UM0Y8": "www.amazon.com.mx", // 墨西哥
	"A2Q3Y263D00KWC": "www.amazon.com.br", // 巴西
	"A2EUQ1WTGCTBG2": "www.amazon.ca",     // 加拿大
	"A2VIGQ35RCS4UG": "www.amazon.ae",     // 阿拉伯联合酋长国
	"A1PA6795UKMFR9": "www.amazon.de",     // 德国
	"A1RKKUPIHCS9HS": "www.amazon.es",     // 西班牙
	"A13V1IB3VIYZZH": "www.amazon.fr",     //法国
	"A1F83G8C2ARO7P": "amazon.co.uk",      // 英国
	"A21TJRUUN4KGV":  "www.amazon.in",     // 印度
	"APJ6JRA9NG5V4":  "www.amazon.it",     // 意大利
	"A33AVAJ2PDY3EV": "www.amazon.com.tr", // 土耳其
	"A39IBJ37TRP1C6": "www.amazon.com.au", // 澳大利亚
	"A1VC38T7YXB528": "www.amazon.jp",     // 小日本
	"AAHKV2X7AFYLW":  "amazon.cn",         // 中国
}

//GetDomainName Get DomainName by marketpalceID
func GetDomainName(marketpalceID string) string {
	for _marketplaceID, endPoint := range EndPoints {
		if _marketplaceID == marketpalceID {
			return endPoint
		}
	}
	return ""
}

// UUID UUID
func UUID() (guid string) {
	u, _ := uuid.NewV4()
	guid = strings.ToUpper(u.String())
	return
}

//CombineString  combine string  like python join function
func CombineString(connectChr string, strList ...string) (output string) {
	if len(strList) == 0 {
		output = ""
	} else {
		output = strList[0]
		for _, chr := range strList[1:] {
			if chr != "" {
				output += fmt.Sprintf("%s%s", connectChr, chr)
			}
		}
	}
	return
}

// UUIDByte uuidbyte
func UUIDByte() []byte {
	u, _ := uuid.NewV4()
	return u.Bytes()
}

// H string to []byte
func H(str string) (b []byte) {
	str = regexp.MustCompile(`\s+`).ReplaceAllString(str, "")
	b, _ = hex.DecodeString(str)
	return
}

// ToFixed 取四舍五入
func ToFixed(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	fv := 0.0000000001 + f //对浮点数产生.xxx999999999 计算不准进行处理
	return math.Floor(fv*shift+.5) / shift
}

//HashString  string -> hash string
func HashString(srcString string) (output string) {
	h := sha1.New()
	h.Write([]byte(srcString))
	output = hex.EncodeToString(h.Sum(nil))
	return
}

//GetPreDate  获取当前前几天的起始日期
func GetPreDate(n int, loc *time.Location) (output time.Time) {
	t1 := time.Now().AddDate(0, 0, -n)
	output = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, loc)
	return
}

// StringInSlice StringInSlice
func StringInSlice(s string, list []string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}

func StringInSlicePit(s string, list []*string) bool {
	for _, item := range list {
		if *item == s {
			return true
		}
	}
	return false
}

func DateRange(from, to string) (dateRange []string) {
	layout := "2006-01-02"
	fromDate, _ := time.Parse(layout, from)
	endDate, _ := time.Parse(layout, to)
	if fromDate.After(endDate) {
		return
	}
	date := fromDate
	for {
		dateRange = append(dateRange, date.Format(layout))
		if date.Equal(endDate) {
			break
		}
		date = date.AddDate(0, 0, 1)
	}
	return
}

// StringDedup string 去重
func StringDedup(arr []string) (dedup []string) {
	mapSet := make(map[string]bool)
	for _, v := range arr {
		mapSet[v] = true
	}
	for k := range mapSet {
		dedup = append(dedup, k)
	}
	return
}

func monthInterval(y int, m time.Month) (firstDay, lastDay time.Time) {
	firstDay = time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	lastDay = time.Date(y, m+1, 1, 0, 0, 0, -1, time.UTC)
	return firstDay, lastDay
}

// MonthInterval MonthInterval
func MonthInterval(year, month int) (fromDate, endDate string) {
	currentYear, currentMonth, _ := time.Now().Date()

	layout := "2006-01-02"
	firstDay, lastDay := monthInterval(year, time.Month(month))
	fromDate, endDate = firstDay.Format(layout), lastDay.Format(layout)
	if currentYear == year && int(currentMonth) == month {
		// 前一天
		endDate = time.Now().Format(layout)
		// endDate = time.Now().AddDate(0, 0, -1).Format(layout)
	}
	return
}

func FormatMonth(year, month int) string {
	return fmt.Sprintf("%d-%02d", year, month)
}

func CurMon() string {
	now := time.Now()
	return FormatMonth(now.Year(), int(now.Month()))
}

// Div Div numberator 分子  denominator 分母
func Div(numberator, denominator float64, decimal int) float64 {
	if denominator == 0 {
		return 0
	}
	v, _ := Decimal(float64(numberator/denominator), decimal)
	return v
}

// Div Div numberator 分子  denominator 分母
func DivInts(numberator, denominator int, decimal int) float64 {
	if denominator == 0 {
		return 0
	}
	return Div(float64(numberator), float64(denominator), 2)
}

func Decimal(value float64, accuracy int) (v float64, err error) {
	fmtString := fmt.Sprintf("%s.%df", "%", accuracy)
	return strconv.ParseFloat(fmt.Sprintf(fmtString, value), 64)
}

func String2Decimal(value string, accuracy int) (v float64, err error) {
	v, err = strconv.ParseFloat(value, 64)
	return Decimal(v, 2)
}

// DivWithPercent  numberator 分子  denominator 分母
func DivWithPercent(numberator, denominator float64, decimal int) string {
	return fmt.Sprintf("%f", Div(numberator, denominator, decimal))
}

// MonthDays MonthDays
func MonthDays(year, month int) (days int) {
	sm := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	em := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	return int(em.Sub(sm).Hours()) / 24
}

func SplitMonth(month, connectChar string) (y, m int, err error) {
	yearMonth := strings.Split(month, "-")
	if len(yearMonth) < 2 {
		err = fmt.Errorf("monthFormatErr")
		return
	}
	y, _ = strconv.Atoi(yearMonth[0])
	m, _ = strconv.Atoi(yearMonth[1])
	return
}

// SubDays 计算起始时间的相隔天数
func SubDays(end, start time.Time) int {
	_t1 := time.Date(end.Year(), time.Month(end.Month()), end.Day(), 0, 0, 0, 0, time.UTC)
	_t2 := time.Date(start.Year(), time.Month(start.Month()), start.Day(), 0, 0, 0, 0, time.UTC)
	days := _t1.Sub(_t2).Hours() / 24
	return int(days)
}
