package first_test

// 输⼊⼀个年份(比如：2021)，判断该年份是否为闰年，是则输出"yes"，否则输出"no"
// 闰年条件:
// 1. 普通闰年：公历年份是4的倍数，且不是100的倍数的，为闰年
// 2. 公历年份是整百数的，必须是400的倍数才是闰年（如1900年不是闰年，2000年是闰年）。
func IsleapYear(year int64) bool {
	if (year%4 == 0 && year%100 != 0) || (year%100 == 0 && year%400 == 0) {
		return true
	}
	return false
}
