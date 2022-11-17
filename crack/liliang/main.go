/**
 * @Author: zzy
 * @Email: zhangzhongyuan@didiglobal.com
 * @Description:
 * @File: main.go
 * @Package: liliang
 * @Version: 1.0.0
 * @Date: 2022/11/17 12:21
 */

package main

import "github.com/gin-gonic/gin"

// M310H
// 型号：新魔百和M310H
// 破解方式1：连接绑定宽带，进入IPTV后，切换成wifi连接到PC。PC抓包分析请求地址和端口和URI，并设置地址后启动hook下载文件服务
// 破解方式2：TTL开启adbd
func M310H() {
	engine := gin.Default()
	engine.StaticFile("/upload/apk/app_hjkhtvmanagerRelease.apk", "/Users/zhangzhongyuan/Downloads/dangbeimarket_4.4.0_288_yunji.apk")
	engine.Run(":18082")
}

func main() {
	M310H()
}
