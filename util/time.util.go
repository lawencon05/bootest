package util

import "time"

const layoutDef = "2006-01-02 15:04:05"

func TimeToStr(tm time.Time, layout ...string) string {
	if len(layout) > 0 {
		return tm.Format(layout[0])
	}
	return tm.Format(layoutDef)
}

func StrToTime(tmStr string, layout ...string) (time.Time, error) {
	if len(layout) > 0 {
		return time.Parse(layout[0], tmStr)
	}
	return time.Parse(layoutDef, tmStr)
}
