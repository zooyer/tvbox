/**
 * @Author: zzy
 * @Email: zhangzhongyuan@didiglobal.com
 * @Description:
 * @File: devices.go
 * @Package: devices
 * @Version: 1.0.0
 * @Date: 2022/10/8 09:47
 */

package devices

import (
	"os"
	"strings"
)

type InputID struct {
	Bus     uint16 `json:"bus"`
	Vendor  uint16 `json:"vendor"`
	Product uint16 `json:"product"`
	Version uint16 `json:"version"`
}

type InputAbsinfo struct {
	Value      int32 `json:"value"`
	Minimum    int32 `json:"minimum"`
	Maximum    int32 `json:"maximum"`
	Fuzz       int32 `json:"fuzz"`
	Flat       int32 `json:"flat"`
	Resolution int32 `json:"resolution"`
}

type InputKeymapEntry struct {
	Flags    uint8     `json:"flags"`
	Len      uint8     `json:"len"`
	Index    uint16    `json:"index"`
	Keycode  uint32    `json:"keycode"`
	Scancode [32]uint8 `json:"scancode"`
}

type InputMask struct {
	Type      uint32 `json:"type"`
	CodesSize uint32 `json:"codes_size"`
	CodesPtr  uint64 `json:"codes_ptr"`
}

type DeviceBitmasks struct {
	PROP uint16 `json:"PROP"`
	EV   uint16 `json:"EV"`
	KEY  string `json:"KEY"`
	REL  uint16 `json:"REL"`
	ABS  string `json:"ABS"`
	MSC  uint16 `json:"MSC"`
	LED  uint16 `json:"LED"`
}

type Device struct {
	ID       InputID `json:"id"`
	Name     string  `json:"name"`
	Phys     string  `json:"phys"`
	Sysfs    string  `json:"sysfs"`
	Uniq     string  `json:"uniq"`
	Handlers string  `json:"handlers"`
}

func ReadInputDevices() {
	data, err := os.ReadFile("/proc/bus/input/devices")
	if err != nil {
		return
	}

	devices := strings.Split(string(data), "\n\n")
	for _, device := range devices {
		if device = strings.TrimSpace(device); device == "" {
			continue
		}
	}
}
