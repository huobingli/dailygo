package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func TimeParseYYYYMMDD(in string, sub string) (out time.Time, err error) {
	layout := "2006" + sub + "01" + sub + "02"
	out, err = time.ParseInLocation(layout, in, time.Local)
	if err != nil {
		return
	}
	return
}

func getCurDay() (date int) {
	curTime := time.Now()
	year := curTime.Year()
	month := int(curTime.Month())
	day := curTime.Day()
	return year*10000 + month*100 + day
}

// 获取CPU使用情况
func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

// 获取内存使用情况
func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

// 获取磁盘使用情况
func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}

func main() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Print("get Partitions failed, err:%v\n", err)
		return
	}

	for _, part := range parts {
		fmt.Print("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Print("disk info:used:%f free:%f\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	strOut := ""
	for k, v := range ioStat {
		strtmp := fmt.Sprintf("%v:%v\n", k, v)
		strOut += strtmp
	}

	fmt.Println(strOut)
}
