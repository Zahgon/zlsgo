package cron

import (
	"regexp"
	"strings"
	"sync"
	"time"
)

// copyright https://github.com/gorhill/cronexpr

type (
	// Expression represents a parsed cron expression
	// Contains the parsing results of each field in the expression, used to calculate the next execution time
	Expression struct {
		daysOfWeek             map[int]bool // Set of days of the week
		lastWeekDaysOfWeek     map[int]bool // Set of last specific weekday of the month
		specificWeekDaysOfWeek map[int]bool // Set of specific weekdays
		daysOfMonth            map[int]bool // Set of days of the month
		workdaysOfMonth        map[int]bool // Set of workdays of the month
		monthList              []int        // List of months
		actualDaysOfMonthList  []int        // List of actual days of the month
		secondList             []int        // List of seconds
		hourList               []int        // List of hours
		minuteList             []int        // List of minutes
		yearList               []int        // List of years
		lastWorkdayOfMonth     bool         // Whether it's the last workday of the month
		daysOfMonthRestricted  bool         // Whether days of month are restricted
		lastDayOfMonth         bool         // Whether it's the last day of the month
		daysOfWeekRestricted   bool         // Whether days of week are restricted
	}
	cronDirective struct {
		kind  int
		first int
		last  int
		step  int
		sbeg  int
		send  int
	}
)

const (
	none = 0
	one  = 1
	span = 2
	all  = 3
)

var (
	dowNormalizedOffsets = [][]int{
		{1, 8, 15, 22, 29},
		{2, 9, 16, 23, 30},
		{3, 10, 17, 24, 31},
		{4, 11, 18, 25},
		{5, 12, 19, 26},
		{6, 13, 20, 27},
		{7, 14, 21, 28},
	}
	genericDefaultList = []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	}
	yearDefaultList = []int{
		1970, 1971, 1972, 1973, 1974, 1975, 1976, 1977, 1978, 1979,
		1980, 1981, 1982, 1983, 1984, 1985, 1986, 1987, 1988, 1989,
		1990, 1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998, 1999,
		2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009,
		2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019,
		2020, 2021, 2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029,
		2030, 2031, 2032, 2033, 2034, 2035, 2036, 2037, 2038, 2039,
		2040, 2041, 2042, 2043, 2044, 2045, 2046, 2047, 2048, 2049,
		2050, 2051, 2052, 2053, 2054, 2055, 2056, 2057, 2058, 2059,
		2060, 2061, 2062, 2063, 2064, 2065, 2066, 2067, 2068, 2069,
		2070, 2071, 2072, 2073, 2074, 2075, 2076, 2077, 2078, 2079,
		2080, 2081, 2082, 2083, 2084, 2085, 2086, 2087, 2088, 2089,
		2090, 2091, 2092, 2093, 2094, 2095, 2096, 2097, 2098, 2099,
	}
	numberTokens = map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"00": 0, "01": 1, "02": 2, "03": 3, "04": 4, "05": 5, "06": 6, "07": 7, "08": 8, "09": 9,
		"10": 10, "11": 11, "12": 12, "13": 13, "14": 14, "15": 15, "16": 16, "17": 17, "18": 18, "19": 19,
		"20": 20, "21": 21, "22": 22, "23": 23, "24": 24, "25": 25, "26": 26, "27": 27, "28": 28, "29": 29,
		"30": 30, "31": 31, "32": 32, "33": 33, "34": 34, "35": 35, "36": 36, "37": 37, "38": 38, "39": 39,
		"40": 40, "41": 41, "42": 42, "43": 43, "44": 44, "45": 45, "46": 46, "47": 47, "48": 48, "49": 49,
		"50": 50, "51": 51, "52": 52, "53": 53, "54": 54, "55": 55, "56": 56, "57": 57, "58": 58, "59": 59,
		"1970": 1970, "1971": 1971, "1972": 1972, "1973": 1973, "1974": 1974, "1975": 1975, "1976": 1976, "1977": 1977, "1978": 1978, "1979": 1979,
		"1980": 1980, "1981": 1981, "1982": 1982, "1983": 1983, "1984": 1984, "1985": 1985, "1986": 1986, "1987": 1987, "1988": 1988, "1989": 1989,
		"1990": 1990, "1991": 1991, "1992": 1992, "1993": 1993, "1994": 1994, "1995": 1995, "1996": 1996, "1997": 1997, "1998": 1998, "1999": 1999,
		"2000": 2000, "2001": 2001, "2002": 2002, "2003": 2003, "2004": 2004, "2005": 2005, "2006": 2006, "2007": 2007, "2008": 2008, "2009": 2009,
		"2010": 2010, "2011": 2011, "2012": 2012, "2013": 2013, "2014": 2014, "2015": 2015, "2016": 2016, "2017": 2017, "2018": 2018, "2019": 2019,
		"2020": 2020, "2021": 2021, "2022": 2022, "2023": 2023, "2024": 2024, "2025": 2025, "2026": 2026, "2027": 2027, "2028": 2028, "2029": 2029,
		"2030": 2030, "2031": 2031, "2032": 2032, "2033": 2033, "2034": 2034, "2035": 2035, "2036": 2036, "2037": 2037, "2038": 2038, "2039": 2039,
		"2040": 2040, "2041": 2041, "2042": 2042, "2043": 2043, "2044": 2044, "2045": 2045, "2046": 2046, "2047": 2047, "2048": 2048, "2049": 2049,
		"2050": 2050, "2051": 2051, "2052": 2052, "2053": 2053, "2054": 2054, "2055": 2055, "2056": 2056, "2057": 2057, "2058": 2058, "2059": 2059,
		"2060": 2060, "2061": 2061, "2062": 2062, "2063": 2063, "2064": 2064, "2065": 2065, "2066": 2066, "2067": 2067, "2068": 2068, "2069": 2069,
		"2070": 2070, "2071": 2071, "2072": 2072, "2073": 2073, "2074": 2074, "2075": 2075, "2076": 2076, "2077": 2077, "2078": 2078, "2079": 2079,
		"2080": 2080, "2081": 2081, "2082": 2082, "2083": 2083, "2084": 2084, "2085": 2085, "2086": 2086, "2087": 2087, "2088": 2088, "2089": 2089,
		"2090": 2090, "2091": 2091, "2092": 2092, "2093": 2093, "2094": 2094, "2095": 2095, "2096": 2096, "2097": 2097, "2098": 2098, "2099": 2099,
	}
	monthTokens = map[string]int{
		`1`: 1, `jan`: 1, `january`: 1,
		`2`: 2, `feb`: 2, `february`: 2,
		`3`: 3, `mar`: 3, `march`: 3,
		`4`: 4, `apr`: 4, `april`: 4,
		`5`: 5, `may`: 5,
		`6`: 6, `jun`: 6, `june`: 6,
		`7`: 7, `jul`: 7, `july`: 7,
		`8`: 8, `aug`: 8, `august`: 8,
		`9`: 9, `sep`: 9, `september`: 9,
		`10`: 10, `oct`: 10, `october`: 10,
		`11`: 11, `nov`: 11, `november`: 11,
		`12`: 12, `dec`: 12, `december`: 12,
	}
	dowTokens = map[string]int{
		`0`: 0, `sun`: 0, `sunday`: 0,
		`1`: 1, `mon`: 1, `monday`: 1,
		`2`: 2, `tue`: 2, `tuesday`: 2,
		`3`: 3, `wed`: 3, `wednesday`: 3,
		`4`: 4, `thu`: 4, `thursday`: 4,
		`5`: 5, `fri`: 5, `friday`: 5,
		`6`: 6, `sat`: 6, `saturday`: 6,
		`7`: 0,
	}
)

func atoi(s string) int { _ = "STUB: not implemented"; return 0 }

type fieldDescriptor struct {
	atoi         func(string) int
	name         string
	valuePattern string
	defaultList  []int
	min          int
	max          int
}

var (
	secondDescriptor = fieldDescriptor{
		name:         "second",
		min:          0,
		max:          59,
		defaultList:  genericDefaultList[0:60],
		valuePattern: `0?[0-9]|[1-5][0-9]`,
		atoi:         atoi,
	}
	minuteDescriptor = fieldDescriptor{
		name:         "minute",
		min:          0,
		max:          59,
		defaultList:  genericDefaultList[0:60],
		valuePattern: `0?[0-9]|[1-5][0-9]`,
		atoi:         atoi,
	}
	hourDescriptor = fieldDescriptor{
		name:         "hour",
		min:          0,
		max:          23,
		defaultList:  genericDefaultList[0:24],
		valuePattern: `0?[0-9]|1[0-9]|2[0-3]`,
		atoi:         atoi,
	}
	domDescriptor = fieldDescriptor{
		name:         "day-of-month",
		min:          1,
		max:          31,
		defaultList:  genericDefaultList[1:32],
		valuePattern: `0?[1-9]|[12][0-9]|3[01]`,
		atoi:         atoi,
	}
	monthDescriptor = fieldDescriptor{
		name:         "month",
		min:          1,
		max:          12,
		defaultList:  genericDefaultList[1:13],
		valuePattern: `0?[1-9]|1[012]|jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec|january|february|march|april|march|april|june|july|august|september|october|november|december`,
		atoi: func(s string) int {
			return monthTokens[s]
		},
	}
	dowDescriptor = fieldDescriptor{
		name:         "day-of-week",
		min:          0,
		max:          6,
		defaultList:  genericDefaultList[0:7],
		valuePattern: `0?[0-7]|sun|mon|tue|wed|thu|fri|sat|sunday|monday|tuesday|wednesday|thursday|friday|saturday`,
		atoi: func(s string) int {
			return dowTokens[s]
		},
	}
	yearDescriptor = fieldDescriptor{
		name:         "year",
		min:          1970,
		max:          2099,
		defaultList:  yearDefaultList[:],
		valuePattern: `19[789][0-9]|20[0-9]{2}`,
		atoi:         atoi,
	}
	layoutWildcard            = `^\*$|^\?$`
	layoutValue               = `^(%value%)$`
	layoutRange               = `^(%value%)-(%value%)$`
	layoutWildcardAndInterval = `^\*/(\d+)$`
	layoutValueAndInterval    = `^(%value%)/(\d+)$`
	layoutRangeAndInterval    = `^(%value%)-(%value%)/(\d+)$`
	layoutLastDom             = `^l$`
	layoutWorkdom             = `^(%value%)w$`
	layoutLastWorkdom         = `^lw$`
	layoutDowOfLastWeek       = `^(%value%)l$`
	layoutDowOfSpecificWeek   = `^(%value%)#([1-5])$`
	fieldFinder               = regexp.MustCompile(`\S+`)
	entryFinder               = regexp.MustCompile(`[^,]+`)
	layoutRegexp              = make(map[string]*regexp.Regexp)
	layoutRegexpLock          sync.Mutex
	cronNormalizer            = strings.NewReplacer(
		"@yearly", "0 0 0 1 1 * *",
		"@annually", "0 0 0 1 1 * *",
		"@monthly", "0 0 0 1 * * *",
		"@weekly", "0 0 0 * * 0 *",
		"@daily", "0 0 0 * * * *",
		"@hourly", "0 0 * * * * *")
)

func (expr *Expression) secondFieldHandler(s string) error { _ = "STUB: not implemented"; return nil }

func (expr *Expression) minuteFieldHandler(s string) error { _ = "STUB: not implemented"; return nil }

func (expr *Expression) hourFieldHandler(s string) error { _ = "STUB: not implemented"; return nil }

func (expr *Expression) monthFieldHandler(s string) error { _ = "STUB: not implemented"; return nil }

func (expr *Expression) yearFieldHandler(s string) error { _ = "STUB: not implemented"; return nil }

func genericFieldHandler(s string, desc fieldDescriptor) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (expr *Expression) dowFieldHandler(s string) error { _ = "STUB: not implemented"; return nil }

// `5L`

// `5#3`

func (expr *Expression) domFieldHandler(s string) error { _ = "STUB: not implemented"; return nil }

// days of month map
// work days of month map

// `L`

// `LW`

// `15W`

func populateOne(values map[int]bool, v int) { _ = "STUB: not implemented"; return }

func populateMany(values map[int]bool, min, max, step int) { _ = "STUB: not implemented"; return }

func toList(set map[int]bool) []int { _ = "STUB: not implemented"; return nil }

func genericFieldParse(s string, desc fieldDescriptor) ([]*cronDirective, error) {
	_ = "STUB: not implemented"
	// At least one entry must be present
	return nil, nil
}

// `*`

// `5`

// `5-20`

// `*/2`

// `5/2`

// `5-20/2`

// No behavior for this one, let caller deal with it

func makeLayoutRegexp(layout, value string) *regexp.Regexp { _ = "STUB: not implemented"; return nil }

// ParseNextTime parses a cron expression and calculates the next execution time.
func ParseNextTime(cronLine string) (nextTime time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Parse parses a cron expression string, returning an Expression object.
func Parse(cronLine string) (*Expression, error) { _ = "STUB: not implemented"; return nil, nil }

// ignore fields beyond 7th

// second field (optional)

// minute field

// hour field

// day of month field

// month field

// day of week field

// year field

// Next returns the next time point after the given time that matches the cron expression.
func (expr *Expression) Next(fromTime time.Time) time.Time {
	_ = "STUB: not implemented"
	// Special case
	return *new(time.Time)
}

// year

// month

// day of month

// hour

// minute

// second

// If we reach this point, there is nothing better to do
// than to move to the next second

// NextN returns n time points after the given time that match the cron expression.
func (expr *Expression) NextN(fromTime time.Time, n uint) []time.Time {
	_ = "STUB: not implemented"
	return nil
}

func (expr *Expression) nextYear(t time.Time) time.Time {
	_ = "STUB: not implemented"
	// Find index at which item in list is greater or equal to
	// candidate year
	return *new(time.Time)
}

// Year changed, need to recalculate actual days of month

func (expr *Expression) nextMonth(t time.Time) time.Time {
	_ = "STUB: not implemented"
	// Find index at which item in list is greater or equal to
	// candidate month
	return *new(time.Time)
}

// Month changed, need to recalculate actual days of month

func (expr *Expression) nextDayOfMonth(t time.Time) time.Time {
	_ = "STUB: not implemented"
	// Find index at which item in list is greater or equal to
	// candidate day of month
	return *new(time.Time)
}

func (expr *Expression) nextHour(t time.Time) time.Time {
	_ = "STUB: not implemented"
	// Find index at which item in list is greater or equal to
	// candidate hour
	return *new(time.Time)
}

func (expr *Expression) nextMinute(t time.Time) time.Time {
	_ = "STUB: not implemented"
	// Find index at which item in list is greater or equal to
	// candidate minute
	return *new(time.Time)
}

func (expr *Expression) nextSecond(t time.Time) time.Time {
	_ = "STUB: not implemented"
	// nextSecond() assumes all other fields are exactly matched
	// to the cron expression
	return *new(time.Time)
}

// Find index at which item in list is greater or equal to
// candidate second

func (expr *Expression) calculateActualDaysOfMonth(year, month int) []int {
	_ = "STUB: not implemented"
	return nil
}

// As per crontab man page (http://linux.die.net/man/5/crontab#):
//  "The day of a command's execution can be specified by two
//  "fields - day of month, and day of week. If both fields are
//  "restricted (ie, aren't *), the command will be run when
//  "either field matches the current time"

// If both fields are not restricted, all days of the month are a hit

// day-of-month != `*`

// Last day of month

// Last work day of month

// Days of month

// Ignore days beyond end of month

// Work days of month
// As per Wikipedia: month boundaries are not crossed.

// Ignore days beyond end of month

// day-of-week != `*`

// How far first sunday is from first day of month

// days of week
//  offset : (7 - day_of_week_of_1st_day_of_month)
//  target : 1 + (7 * week_of_month) + (offset + day_of_week) % 7

// days of week of specific week in the month
//  offset : (7 - day_of_week_of_1st_day_of_month)
//  target : 1 + (7 * week_of_month) + (offset + day_of_week) % 7

// Last days of week of the month

func workdayOfMonth(targetDom, lastDom time.Time) int {
	_ = "STUB: not implemented"
	// If saturday, then friday
	// If sunday, then monday
	return 0
}
