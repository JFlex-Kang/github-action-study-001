package main

import "fmt"

var version string		// 빌드 시 ldflags 플래그 경유하여 버저을 저장하기 위한 변수

func main() {
	fmt.Printf("Example %s \n", version)
}