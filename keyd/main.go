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
	"os"
	"os/exec"
)

var (
	model string
	hooks = map[string]string{
		"H3-2s":   "/data/zzy/init.sh",         // 移动光猫-兴隆魁
		"F610GV9": "/opt/upt/apps/zzy/init.sh", // 电信光猫-张北(空闲)
	}
)

// xlkTVbox 兴隆魁移动电视盒子
func xlkTVBox() {
	_ = exec.Command("/data/zzy/init.sh").Start()
	_ = exec.Command("/system/bin/testagent.origin").Run()
}

// main
// example:
//   agent
//   agent.hook.sh
//   agent.hook.origin
func main() {
	_ = exec.Command(os.Args[0] + ".hook.sh").Start()
	_ = exec.Command(os.Args[0]+".hook.origin", os.Args[1:]...).Run()
}
