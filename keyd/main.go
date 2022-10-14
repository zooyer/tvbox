/**
 * @Author: zzy
 * @Email: zhangzhongyuan@didiglobal.com
 * @Description:
 * @File: main.go
 * @Package: main
 * @Version: 1.0.0
 * @Date: 2022/10/1 17:22
 */

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"time"
	"unsafe"

	"github.com/zooyer/embed/log"
	"gopkg.in/yaml.v3"
)

type Hook struct {
	Key string `yaml:"key"`
	Cmd string `yaml:"cmd"`
}

type Config struct {
	Shell  string     `yaml:"shell"`
	Device string     `yaml:"device"`
	Hooks  []Hook     `yaml:"hooks"`
	Log    log.Config `yaml:"log"`
}

type Timeval struct {
	Sec  uint32
	USec uint32
}

type InputEvent struct {
	Time  Timeval
	Type  uint16
	Code  uint16
	Value uint32
}

func IsLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb

	return b == 0x04
}

func execCommand(shell, cmd string) {
	if cmd == "" {
		return
	}

	if shell == "" {
		shell = "/system/bin/sh"
	}

	_ = exec.Command(shell, "-c", cmd).Run()
}

func main() {
	// 1. 读取配置文件
	data, err := os.ReadFile("./keyd.yaml")
	if err != nil {
		panic(err)
	}

	// 2. 解析配置文件
	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	// 3. 初始化日志
	log.Init(&config.Log)

	// 4. 判断设备文件是否存在
	for {
		if _, err = os.Stat(config.Device); err != nil {
			log.ZError("stat", config.Device, "error:", err.Error())
			time.Sleep(time.Second)
			continue
		}
		break
	}

	// 5. 建立hook索引
	var index = make(map[string]string)
	for _, hook := range config.Hooks {
		index[hook.Key] = hook.Cmd
	}

	// 6. 打开设备文件(模拟getevent)
	file, err := os.Open(config.Device)
	if err != nil {
		log.ZError("open", config.Device, "error:", err.Error())
		return
	}
	defer file.Close()

	// 7. 读取输入事件并处理
	var msg = make([]byte, 24)
	for {
		_, err := file.Read(msg)
		if err != nil {
			log.ZError("read", config.Device, "error:", err.Error())
			time.Sleep(time.Second)
			continue
		}

		var event InputEvent
		var order binary.ByteOrder = binary.BigEndian
		if IsLittleEndian() {
			order = binary.LittleEndian
		}

		_ = binary.Read(bytes.NewReader(msg), order, &event)
		var key = fmt.Sprintf("%04x %04x %08x", event.Type, event.Code, event.Value)

		log.ZTrace("key:", key)

		if key == "0000 0000 00000000" {
			continue
		}

		if cmd := index[key]; cmd != "" {
			go execCommand(config.Shell, cmd)
		}
	}
}
