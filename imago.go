package imago

// E exception
var (
	E        = &iexception{}
	G        = &igenerate{}
	Convert  = &iconvert{} // Convert 转换
	Crypto   = &icrypto{}  // Crypto 加密
	File     = &ifile{}    // File 文件
	Time     = &itime{}    // Time 时间
	monthMap = map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}
	weekMap = map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    0,
	}
)

// H hash
type (
	H          map[string]interface{}
	iexception struct{}
	igenerate  struct{}
	iconvert   struct{}
	icrypto    struct{}
	ifile      struct{}
	itime      struct{}
)
