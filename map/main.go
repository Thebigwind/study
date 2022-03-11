package main

import (
	"fmt"
	"math"
)

/*
Go 语言中读取 map 有两种语法：带 comma 和 不带 comma。当要查询的 key 不在 map 里，带 comma 的用法会返回一个 bool 型变量提示 key 是否在 map 中；
而不带 comma 的语句则会返回一个 key 类型的零值。如果 key 是 int 型就会返回 0，如果 key 是 string 类型，就会返回空字符串。
*/
func main() {
	Data := make(map[string]struct{}, 0)
	fmt.Printf("Data:%+v\n", Data)
	//ageMap := make(map[string]int)
	//ageMap["qcrao"] = 18
	//
	//// 不带 comma 用法
	//age1 := ageMap["stefno"]
	//fmt.Println(age1)
	//
	//// 带 comma 用法
	//age2, ok := ageMap["stefno"]
	//fmt.Println(age2, ok)

	//test()

	test3()
}

func test() {
	m := make(map[float64]int)
	m[1.4] = 1
	m[2.4] = 2
	m[math.NaN()] = 3
	m[math.NaN()] = 3

	for k, v := range m {
		fmt.Printf("[%v, %d] ", k, v)
	}

	fmt.Printf("\nk: %v, v: %d\n", math.NaN(), m[math.NaN()])
	fmt.Printf("k: %v, v: %d\n", 2.400000000001, m[2.400000000001])
	fmt.Printf("k: %v, v: %d\n", 2.4000000000000000000000001, m[2.4000000000000000000000001])

	fmt.Println(math.NaN() == math.NaN())
}

func test2() {
	aa := make(map[int][]string)

	value, ok := aa[4]
	if ok {
		fmt.Println("xx")
		fmt.Println(value)
	} else {
		fmt.Println("oo")
		fmt.Println(value)
	}

	//aaV := make([]string,0)
	aa[4] = append(aa[4], "3")

	if ok {
		fmt.Println("xxxx")
		fmt.Println(aa[4])
	} else {
		fmt.Println("oooo")
		fmt.Println(aa[4])
	}
}

func test3() {
	var aa = map[string]string{
		"0":   "CN",
		"1":   "HK",
		"2":   "MO",
		"3":   "TW",
		"4":   "AF",
		"5":   "AL",
		"6":   "DZ",
		"7":   "AD",
		"8":   "AO",
		"9":   "AI",
		"10":  "AG",
		"11":  "AR",
		"12":  "AM",
		"13":  "SH",
		"14":  "AU",
		"15":  "AT",
		"16":  "AZ",
		"17":  "BS",
		"18":  "BH",
		"19":  "BD",
		"20":  "BB",
		"21":  "BY",
		"22":  "BE",
		"23":  "BZ",
		"24":  "BJ",
		"25":  "BM",
		"26":  "BO",
		"27":  "BW",
		"28":  "BR",
		"29":  "BN",
		"30":  "BG",
		"31":  "BF",
		"32":  "BU",
		"33":  "BI",
		"34":  "CM",
		"35":  "CA",
		"36":  "KY",
		"37":  "CF",
		"38":  "TD",
		"39":  "CL",
		"40":  "CO",
		"41":  "CD",
		"42":  "CG",
		"43":  "CK",
		"44":  "CR",
		"45":  "CU",
		"46":  "CY",
		"47":  "CZ",
		"48":  "DK",
		"49":  "DJ",
		"50":  "DO",
		"51":  "AS",
		"52":  "EC",
		"53":  "EG",
		"54":  "SV",
		"55":  "EE",
		"56":  "ET",
		"57":  "FJ",
		"58":  "FI",
		"59":  "FR",
		"60":  "GF",
		"61":  "PF",
		"62":  "GA",
		"63":  "GM",
		"64":  "GE",
		"65":  "DE",
		"66":  "GH",
		"67":  "GI",
		"68":  "GR",
		"69":  "GD",
		"70":  "GU",
		"71":  "GT",
		"72":  "GN",
		"73":  "GY",
		"74":  "HT",
		"75":  "HN",
		"76":  "HU",
		"77":  "IS",
		"78":  "IN",
		"79":  "ID",
		"80":  "IR",
		"81":  "IQ",
		"82":  "IE",
		"83":  "IL",
		"84":  "IT",
		"85":  "",
		"86":  "JM",
		"87":  "JP",
		"88":  "JO",
		"89":  "KH",
		"90":  "KZ",
		"91":  "KE",
		"92":  "KR",
		"93":  "KW",
		"94":  "KG",
		"95":  "LA",
		"96":  "LV",
		"97":  "LB",
		"98":  "LS",
		"99":  "LR",
		"100": "LY",
		"101": "LI",
		"102": "LT",
		"103": "LU",
		"104": "MG",
		"105": "MW",
		"106": "MY",
		"107": "MV",
		"108": "ML",
		"109": "MT",
		"111": "MQ",
		"112": "MU",
		"113": "MX",
		"114": "MD",
		"115": "MC",
		"116": "MN",
		"117": "MS",
		"118": "MA",
		"119": "MZ",
		"120": "NA",
		"121": "NR",
		"122": "NP",
		"123": "",
		"124": "NL",
		"125": "NZ",
		"126": "NI",
		"127": "NE",
		"128": "NG",
		"129": "KR",
		"130": "NO",
		"131": "OM",
		"132": "PK",
		"133": "PA",
		"134": "PG",
		"135": "PY",
		"136": "PE",
		"137": "PH",
		"138": "PL",
		"139": "PT",
		"140": "PR",
		"141": "QA",
		"142": "RE",
		"143": "RO",
		"144": "RU",
		"145": "LC",
		"146": "VC",
		"147": "SM",
		"148": "ST",
		"149": "SA",
		"150": "SN",
		"151": "SC",
		"152": "SL",
		"153": "SG",
		"154": "SK",
		"155": "SI",
		"156": "SB",
		"157": "SO",
		"158": "SS",
		"159": "ZA",
		"160": "ES",
		"161": "LK",
		"164": "SD",
		"165": "SR",
		"166": "SZ",
		"167": "SE",
		"168": "CH",
		"169": "SY",
		"170": "TJ",
		"171": "TZ",
		"172": "TH",
		"173": "",
		"174": "TG",
		"175": "TO",
		"176": "TT",
		"177": "TN",
		"178": "TR",
		"179": "TM",
		"180": "UG",
		"181": "UA",
		"182": "AE",
		"183": "GB",
		"184": "US",
		"185": "UY",
		"186": "UZ",
		"187": "VE",
		"188": "VN",
		"189": "YE",
		"191": "ZM",
		"192": "ZW",
	}

	str := ""
	bb := make(map[string]string)
	for k, v := range aa {
		bb[v] = k
	}
	//fmt.Printf("%+v",bb)
	for k, v := range bb {
		str += "\"" + k + "\"" + ":" + "\"" + v + "\"" + ","
	}
	fmt.Printf(str)
}
