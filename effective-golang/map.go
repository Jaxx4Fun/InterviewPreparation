package effective

import "log"

func offset(tz string) int {
	timeZone := map[string]int{
		"a": 0,
	}
	delete(timeZone, "PDT") // Now on Standard Time

	// 检查ok，确定是否存在key
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}
