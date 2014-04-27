// Chinese Lunar Calendar Package.
package lunar

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	// Since 1900~2050
	// base data about chinese year informace
	// 保存公历农历之间的转换信息:以任意一年作为起点，
	// 把从这一年起若干年(依需要而定)的农历信息保存起来。
	// 要保存一年的信息，只要两个信息就够了: 1)农历每个月的大小;2)今年是否有闰月，闰几月以及闰月的大小。
	// 用一个整数来保存这些信息就足够了。
	// 具体的方法是:用一位来表示一个月的大小，大月记为1，小月记为0，
	// 这样就用掉12位(无闰月)或13位(有闰月)，再用高四位来表示闰月的月份，没有闰月记为0。
	// ※-----例----: 2000年的信息数据是0xc96，化成二进制就是110010010110B，
	// 表示的含义是:1、2、5、8、10、11月大，其余月份小。 2001年的农历信息数据是0x1a95(因为闰月，所以有13位)，
	// 具体的就是1、2、4、5、8、10、12月大， 其余月份小(0x1a95=1101010010101B)，
	// 4月的后面那一个0表示的是闰4月小，接着的那个1表示5月大。 这样就可以用一个数组来保存这些信息。在这里用数组lunarInfo[]来保存这些信息
	lunarTable = []int{
		0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260,
		0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2,
		0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255,
		0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977,
		0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40,
		0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970,
		0x06566, 0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0,
		0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950,
		0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4,
		0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557,
		0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5d0,
		0x14573, 0x052d0, 0x0a9a8, 0x0e950, 0x06aa0,
		0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570,
		0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0,
		0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4,
		0x0d250, 0x0d558, 0x0b540, 0x0b5a0, 0x195a6,
		0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a,
		0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570,
		0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50,
		0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0,
		0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552,
		0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5,
		0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9,
		0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930,
		0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60,
		0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530,
		0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0,
		0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45,
		0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577,
		0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0,
	}
	GanTable            = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	ZhiTable            = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	AnimalTable         = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	lunarMonthNameTable = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "腊"}
	monthStr1           = []string{"初", "十", "廿", "卅"}
	monthStr2           = []string{"日", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	BaseYear            = 1900
)

//Solar structure
type Solar struct {
	time.Time
}

func NewSolar(year, month, day, hour, min, sec int) *Solar {
	t := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC)
	return &Solar{t}
}

func NewSolarNow() *Solar {
	return &Solar{time.Now()}
}

func (s *Solar) String() string {
	return fmt.Sprintf("%d年%d月%d日 %d时%d分%d秒", s.Year(), s.Month(), s.Day(), s.Hour(), s.Minute(), s.Second())
}

//Luanr structure
type Lunar struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
}

func NewLunar(year, month, day, hour, min, sec int) *Lunar {
	return &Lunar{year, month, day, hour, min, sec}
}

func NewLunarNow() *Lunar {
	return NewSolarNow().Convert()
}

func (l *Lunar) String() string {
	return fmt.Sprintf("%s%s%s %d时%d分%d秒", YearString(l.Year()), MonthString(l.Month()), DayString(l.Day()), l.Hour(), l.Minute(), l.Second())
}

func (s *Solar) Convert() *Lunar {
	var i int
	var leap int
	var isLeap bool
	var temp int

	var day, dayCyl int
	var month, monCyl int
	var year, yearCyl int

	base := time.Date(BaseYear, 1, 31, 0, 0, 0, 0, time.UTC)
	offset := int(s.Sub(base).Hours() / 24) //offset days

	dayCyl = offset + 40
	monCyl = 14

	for i = BaseYear; i < 2050 && offset > 0; i++ {
		temp = YearDays(i)
		offset -= temp
		monCyl += 12
	}

	if offset < 0 {
		offset += temp
		i--
		monCyl -= 12
	}

	year = i
	yearCyl = i - 1864

	leap = LeapMonth(i) //闰哪个月
	isLeap = false

	for i := 1; i < 13 && offset > 0; i++ {
		//leap month
		if leap > 0 && i == (leap+1) && this.isLeap == false {
			i--
			isLeap = true
			temp = LeapDays(this.year)
		} else {
			temp = MonthDays(this.year, i)
		}
		//reset leap month
		if isLeap == true && i == (leap+1) {
			isLeap = false
		}
		offset -= temp
		if isLeap == false {
			monCyl++
		}
	}

	if offset == 0 && leap > 0 && i == (leap+1) {
		if isLeap {
			isLeap = false
		} else {
			isLeap = true
			i--
			monCyl--
		}
	}

	if offset < 0 {
		offset += temp
		i--
		monCyl--
	}
	month = i
	day = offset + 1

	return &Lunar{year, month, day, s.Hour(), s.Minute(), s.Second()}
}

func (l *Lunar) Convert() *Solar {
	lyear := l.Year()
	lmonth = l.Month()
	lday := l.Day()
	offset := 0
	leap := IsLeap(lyear)

	// increment year
	for i := BaseYear; i < lyear; i++ {
		offset += YearDays(i)
	}

	// increment month
	// add days in all months up to the current month
	for i := 1; i < lmonth; i++ {
		// add extra days for leap month
		if i == LeapMonth(lyear) {
			offset += LeapDays(lyear)
		}
		offset += MonthDays(lyear, i)
	}
	// if current month is leap month, add days in normal month
	isLeapMonth := (LeapMonth(lyear) == lmonth)
	if leap && isLeapMonth {
		offset += MonthDays(lyear, i)
	}
	// increment
	offset += lday - 1

	base := time.Date(BaseYear, 1, 31, 0, 0, 0, 0, time.UTC)
	solar := base.Add(time.Duration(offset * 86400000))

	year := solar.Year()
	month := solar.Month()
	day := solar.Day()
	return NewSolar(year, month, day, l.Hour(), l.Minute(), l.Second())
}

/*
 * Lunar Methods
 */
func (l *Lunar) Year() {
	return l.year
}

func (l *Lunar) Month() {
	return l.month
}

func (l *Lunar) Day() {
	return l.day
}

func (l *Lunar) Hour() {
	return l.hour
}

func (l *Lunar) Minute() {
	return l.minute
}

func (l *Lunar) Second() {
	return l.second
}

func (l *Lunar) Festival(fm FestivalMap) (string, error) {
	m := fmt.Sprintf("%2d", l.month)
	d := fmt.Sprintf("%2d", l.day)

	return fm.Get(m + d)
}

/*
 * 24 JieQi
 */
const JieQiTable = []string{
	"小寒", "大寒", "立春", "雨水", "惊蛰", "春分",
	"清明", "谷雨", "立夏", "小满", "芒种", "夏至",
	"小暑", "大暑", "立秋", "处暑", "白露", "秋分",
	"寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
}

var JieQiTableBase = []int{4, 19, 3, 18, 4, 19, 4, 19, 4, 20, 4, 20, 6, 22, 6, 22, 6, 22, 7, 22, 6, 21, 6, 21}
var JieQiTableIdx = "0123415341536789:;<9:=<>:=1>?012@015@015@015AB78CDE8CD=1FD01GH01GH01IH01IJ0KLMN;LMBEOPDQRST0RUH0RVH0RWH0RWM0XYMNZ[MB\\]PT^_ST`_WH`_WH`_WM`_WM`aYMbc[Mde]Sfe]gfh_gih_Wih_WjhaWjka[jkl[jmn]ope]qph_qrh_sth_W"
var JieQiTableOffset = "211122112122112121222211221122122222212222222221222122222232222222222222222233223232223232222222322222112122112121222211222122222222222222222222322222112122112121222111211122122222212221222221221122122222222222222222222223222232222232222222222222112122112121122111211122122122212221222221221122122222222222222221211122112122212221222211222122222232222232222222222222112122112121111111222222112121112121111111222222111121112121111111211122112122112121122111222212111121111121111111111122112122112121122111211122112122212221222221222211111121111121111111222111111121111111111111111122112121112121111111222111111111111111111111111122111121112121111111221122122222212221222221222111011111111111111111111122111121111121111111211122112122112121122211221111011111101111111111111112111121111121111111211122112122112221222211221111011111101111111110111111111121111111111111111122112121112121122111111011111121111111111111111011111111112111111111111011111111111111111111221111011111101110111110111011011111111111111111221111011011101110111110111011011111101111111111211111001011101110111110110011011111101111111111211111001011001010111110110011011111101111111110211111001011001010111100110011011011101110111110211111001011001010011100110011001011101110111110211111001010001010011000100011001011001010111110111111001010001010011000111111111111111111111111100011001011001010111100111111001010001010000000111111000010000010000000100011001011001010011100110011001011001110111110100011001010001010011000110011001011001010111110111100000010000000000000000011001010001010011000111100000000000000000000000011001010001010000000111000000000000000000000000011001010000010000000"

// y年的第n个节气为几日(从0,即小寒算起)
func JieQi(year, n int) int {
	charcodeAt := int(JieQiTableIdx[year-BaseYear])
	offset := strconv.Atoi(JieQiTableOffset[(charcodeAt-48)*24+n])
	//return JieQiTableBase[n] + JieQiTableOffset.charAt((JieQiTableIdx.charCodeAt(year-BaseYear)-48)*24+n)
	return JieQiTableBase + offset
}

/*
 * Common Methods
 */

func IsLeap(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}

//the total days of this year
func YearDays(year int) int {
	sum := 348
	for i := 0x8000; i > 0x8; i >>= 1 {
		if lunarTable[year-BaseYear] & i {
			sum += 1
		}
	}
	return sum + LeapDays(year)
}

//which month leaps in this year?
//return 1-12(if there is one) or 0(no leap month).
func LeapMonth(year int) int {
	return lunarTable[year-BaseYear] & 0xf
}

//the days of this year's leap month
func LeapDays(year int) int {
	if LeapMonth(year) {
		if lunarTable[year-BaseYear] & 0x10000 {
			return 30
		}
		return 29
	}
	return 0
}

//the days of the m-th month of this year
func MonthDays(year, month int) int {
	if lunarTable[year-BaseYear] & (0x10000 >> month) {
		return 30
	}
	return 29
}

func YearString(year int) string {
	return string(year) + "年"
}

func MonthString(month int) string {
	return lunarMonthNameTable[month%12] + "月"
}

func DayString(day int) (s string) {
	switch day {
	case 10:
		s = "初十"
	case 20:
		s = "二十"
	case 30:
		s = "三十"
	default:
		s = monthStr1[int(day/10)]
		s += monthStr2[day%10]
	}
	return
}

/*
 * Utils
 */

// Tian Gan
func Gan(x int) string {
	return GanTable[x%10]
}

// Di Zhi
func Zhi(x int) string {
	return ZhiTable[x%12]
}

// Tian Gan & Di Zhi
func GanZhi(x int) string {
	return GanTable[x%10] + ZhiTable[x%12]
}

// Sheng Xiao
func AnimalYear(year int) string {
	return AnimalTable[((year - BaseYear) % 12)]
}

type FestivalMap map[string]string

func NewFestivalMap() FestivalMap {
	return make(FestivalMap)
}

func (fm FestivalMap) Add(key, val string) {
	fm[key] = val
}

func (fm FestivalMap) Del(key string) {
	delete(fm, key)
}

func (fm FestivalMap) Get(key string) (string, error) {
	desc, ok := fm[key]
	if ok {
		return desc, nil
	}
	return "", errors.New("NotFound")
}

func (fm FestivalMap) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	file.Close()
	for k, v := range fm {
		file.WriteString(k, " ", v)
	}
	return nil
}

func NewFestivalsFromFile(filename string) FestivalMap {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	fest := NewFestivalMap()
	r, err := bufio.NewReader(file)
	for {
		buf, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		line := strings.Trim(string(buf))
		items := strings.Split(line, " ")
		date := items[0]
		desc := items[1]
		fest.Add(date, desc)
	}
	return fest
}

const (
	SolarFestivals = FestivalMap{
		"0101": "元旦",
		"0214": "情人节",
		"0308": "妇女节",
		"0312": "植树节",
		"0401": "愚人节",
		"0422": "地球日",
		"0501": "劳动节",
		"0504": "青年节",
		"0531": "无烟日",
		"0601": "儿童节",
		"0606": "爱眼日",
		"0701": "建党日",
		"0707": "抗战纪念日",
		"0801": "建军节",
		"0910": "教师节",
		"0918": "九·一八事变纪念日",
		"1001": "国庆节",
		"1031": "万圣节",
		"1111": "光棍节",
		"1201": "艾滋病日",
		"1213": "南京大屠杀纪念日",
		"1224": "平安夜",
		"1225": "圣诞节",
	}
	LunarFestivals = FestivalMap{
		"0101": "春节",
		"0115": "元宵节",
		"0202": "龙抬头",
		"0505": "端午节",
		"0707": "七夕",
		"0715": "中元节",
		"0815": "中秋节",
		"0909": "重阳节",
		"1208": "腊八节",
		"1223": "小年",
		"0100": "除夕",
	}
)