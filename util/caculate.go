package util

import "time"

// 计算指定时间段内有多少个工作日
// begin 开始时间
// end 结束时间
func CalcWorkDays(begin, end time.Time) int {
	var currentTime = begin
	var workingCount int
	for {
		if currentTime.After(end) {
			break
		}
		// 周六周日
		if currentTime.Weekday() == time.Sunday || currentTime.Weekday() == time.Saturday {
			// nothing
		} else {
			workingCount++
		}
		currentTime = currentTime.Add(24 * time.Hour)
	}
	return workingCount
}
