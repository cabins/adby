package core

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

func deviceService(service string, operate string) {
	cmd := exec.Command("adb", "shell", "svc", service, operate)
	cmd.Run()
}

// WifiOn - Turn on the wifi
func WifiOn() {
	exec.Command("adb", "shell", "svc", "wifi", "enable").Run()
}

// WifiOff - Turn off the wifi
func WifiOff() {
	exec.Command("adb", "shell", "svc", "wifi", "disable").Run()
}

// WifiPrefer - Set the wifi as prefer choice
func WifiPrefer() {
	exec.Command("adb", "shell", "svc", "wifi", "prefer").Run()
}

// DataOn - Turn on the data connection
func DataOn() {
	exec.Command("adb", "shell", "svc", "data", "enable").Run()
}

// DataOff - Turn off the data connection
func DataOff() {
	exec.Command("adb", "shell", "svc", "data", "disable").Run()
}

// DataPrefer - Set the data connection as prefer choice
func DataPrefer() {
	exec.Command("adb", "shell", "svc", "data", "prefer").Run()
}

// StatusBarExpand - Expand the status bar
func StatusBarExpand() {
	exec.Command("adb", "shell", "cmd", "statusbar", "expand-notifications").Run()
}

// StatusBarCollapse - Collapse the status bar
func StatusBarCollapse() {
	exec.Command("adb", "shell", "cmd", "statusbar", "collapse").Run()
}

// StatusBarSettings - settings for status bar
func StatusBarSettings() {
	exec.Command("adb", "shell", "cmd", "statusbar", "expand-settings").Run()
}

func StatusBarForbid() {
	exec.Command("adb", "shell", "settings", "put", "global", "policy_control", "immersive.full=*").Run()
}

func StatusBarResume() {
	exec.Command("adb", "shell", "settings", "put", "global", "policy_control", "null").Run()
}

// DeviceCurrentTime - get the current time
func DeviceCurrentTime() string {
	timeString, err := exec.Command("adb", "shell", "date '+%F %T'").Output()
	if err != nil {
		log.Println("获取时间出错")
		return ""
	}

	return strings.TrimSpace(string(timeString))
}

// DeviceAutoTime - set the device's time to auto-setting
func DeviceAutoTime() {
	exec.Command("adb", "shell", "settings", "put", "global", "auto_time", "1").Run()
}

// DeviceTimeChange - Change the time of the device
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

type screen struct {
	Brightness       string `json:"brightness"`
	BrightnessMode   string `json:"brightness_mode"`
	ScreenSize       string `json:"screen_size"`
	Density          string `jsong:"density"`
	ScreenOffTimeout string `json"screen_off_timeout"`
}

func (s screen) PrintAsJSON() {
	res, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Println(err)
		fmt.Println(s)
		return
	}
	fmt.Println(string(res))
}

func (s screen) PrintAsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"分辨率", "屏幕密度", "亮度", "是否自动亮度", "息屏时间(ms)"})
	table.Append([]string{s.ScreenSize, s.Density, s.Brightness, s.BrightnessMode, s.ScreenOffTimeout})
	table.Render()
}

func GetScreenInfo() screen {
	brightness := ScreenBrightness("")
	brightnessMode := ScreenBrightMode(-1)
	screenSize := ScreenGetSize()
	density := ScreenGetDensity()
	screenOffTimeout := ScreenOffTimeout(-1)

	return screen{
		Brightness:       brightness,
		BrightnessMode:   brightnessMode,
		ScreenSize:       screenSize,
		Density:          density,
		ScreenOffTimeout: screenOffTimeout,
	}
}

func ScreenBrightness(b string) string {
	if b == "" {
		brightness, err := exec.Command("adb", "shell", "settings", "get", "system", "screen_brightness").Output()
		if err != nil {
			log.Println(err)
			brightness = []byte("Unknown")
		}
		return strings.TrimSpace(string(brightness))
	}

	_, err := strconv.Atoi(b)
	if err != nil {
		return err.Error()
	}

	exec.Command("adb", "shell", "settings", "put", "system", "screen_brightness", b).Run()
	return ScreenBrightness("")
}

func ScreenBrightMode(b int) string {
	if b != 1 && b != 0 {
		brightnessMode, err := exec.Command("adb", "shell", "settings", "get", "system", "screen_brightness_mode").Output()
		if err != nil {
			log.Println(err)
			brightnessMode = []byte("Unknown")
		}
		return strings.TrimSpace(string(brightnessMode))
	}
	exec.Command("adb", "shell", "settings", "put", "system", "screen_brightness_mode", strconv.Itoa(b)).Run()
	return ScreenBrightMode(-1)
}

func ScreenGetSize() string {
	screenSize, err := exec.Command("adb", "shell", "wm", "size").Output()
	if err != nil {
		log.Println(err)
		screenSize = []byte("Unknown")
	}

	return strings.TrimSpace(strings.Split(string(screenSize), ":")[1])
}

func ScreenGetDensity() string {
	screenSize, err := exec.Command("adb", "shell", "wm", "density").Output()
	if err != nil {
		log.Println(err)
		screenSize = []byte("Unknown")
	}

	return strings.TrimSpace(strings.Split(string(screenSize), ":")[1])
}

func ScreenOffTimeout(t int) string {
	if t == -1 {
		screenOffTimeout, err := exec.Command("adb", "shell", "settings", "get", "system", "screen_off_timeout").Output()
		if err != nil {
			log.Println(err)
			screenOffTimeout = []byte("Unknown")
		}
		return strings.TrimSpace(string(screenOffTimeout))
	}

	exec.Command("adb", "shell", "settings", "put", "system", "screen_off_timeout", strconv.Itoa(t)).Run()
	return ScreenOffTimeout(-1)
}
