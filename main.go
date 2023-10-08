package main

import (
	"fmt"
	"time"
	"wood/wood-pkg/util"
)

func main() {
	result := util.CalcWorkDays(time.Unix(1696089600, 0), time.Unix(1696608000, 0))
	fmt.Println(result)
}
