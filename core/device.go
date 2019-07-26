package core

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func deviceService(service string, operate string) {
	cmd := exec.Command("adb", "shell", "svc", service, operate)
	cmd.Run()
}

func WifiOn() {
	exec.Command("adb", "shell", "svc", "wifi", "enable").Run()
}

func WifiOff() {
	exec.Command("adb", "shell", "svc", "wifi", "disable").Run()
}

func WifiPrefer() {
	exec.Command("adb", "shell", "svc", "wifi", "prefer").Run()
}

func DataOn() {
	exec.Command("adb", "shell", "svc", "data", "enable").Run()
}

func DataOff() {
	exec.Command("adb", "shell", "svc", "data", "disable").Run()
}

func DataPrefer() {
	exec.Command("adb", "shell", "svc", "data", "prefer").Run()
}

func StatusBarExpand() {
	exec.Command("adb", "shell", "cmd", "statusbar", "expand-notifications").Run()
}

func StatusBarCollapse() {
	exec.Command("adb", "shell", "cmd", "statusbar", "collapse").Run()
}

func StatusBarSettings() {
	exec.Command("adb", "shell", "cmd", "statusbar", "expand-settings").Run()
}

func StatusBarForbid() {
	exec.Command("adb", "shell", "settings", "put", "global", "policy_control", "immersive.full=*").Run()
}

func StatusBarResume() {
	exec.Command("adb", "shell", "settings", "put", "global", "policy_control", "null").Run()
}

func DeviceCurrentTime() string {
	timeString, err := exec.Command("adb", "shell", "date '+%F %T'").Output()
	if err != nil {
		log.Println("获取时间出错")
		return ""
	}

	return strings.TrimSpace(string(timeString))
}

func DeviceAutoTime() {
	exec.Command("adb", "shell", "settings", "put", "global", "auto_time", "1").Run()
}

func DeviceTimeChange(year, month, week, day, hour, minute int) {
	// 先获取到当前设备的时间
	currentTimeString := DeviceCurrentTime()
	t, err := time.Parse("2006-01-02 15:04:05", currentTimeString)
	if err != nil {
		log.Println("解析时间出错，不做任何修改", err)
		return
	}

	log.Println(year, month, day, hour, minute, week)
	// Add year, month, day
	t = t.AddDate(year, month, day+week*7)

	// Add Time
	t = t.Add(time.Duration(hour)*time.Hour + time.Duration(minute)*time.Minute)

	// Turn off the auto time-setting
	exec.Command("adb", "shell", "settings", "put", "global", "auto_time", "0").Run()

	// Set the new t (time) to device
	// 兼容6.0以下的机器
	// TODO

	// 对于6.0以上的机器
	err = exec.Command("adb", "shell", fmt.Sprintf("date %02d%02d%02d%02d%02d", t.Month(), t.Day(), t.Hour(), t.Minute(), t.Year())).Run()
	log.Println(err)
}

type DeviceInfo struct {
	SerialNo string `json:"serial_no"`
	Mac      string `json:"mac"`
	Battery  string `json:"battery"`
	Time     string `json:"time"`
}

func (device DeviceInfo) PrintAsJSON() {
	res, err := json.MarshalIndent(device, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(res))
}

func (device DeviceInfo) PrintAsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"SerialNo", "Mac", "Battery", "Time"})

	table.Append([]string{device.SerialNo, device.Mac, device.Battery, device.Time})
	table.Render()
}

func GetDeviceInfo() DeviceInfo {
	serialNo, err := exec.Command("adb", "get-serialno").Output()
	if err != nil {
		serialNo = []byte("")
	}

	mac, err := exec.Command("adb", "shell", "cat", "/sys/class/net/wlan0/address").Output()
	if err != nil {
		mac = []byte("")
	}

	battery, err := exec.Command("adb", "shell", "dumpsys", "battery", "|", "grep", "level").Output()
	if err != nil {
		battery = []byte("")
	}

	tTime := DeviceCurrentTime()

	return DeviceInfo{
		SerialNo: strings.TrimSpace(string(serialNo)),
		Mac:      strings.TrimSpace(string(mac)),
		Battery:  strings.TrimSpace(string(battery)),
		Time:     tTime,
	}
}
