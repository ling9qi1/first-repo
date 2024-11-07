package main

import (
	"fmt" // 导入 Go 标准库的 fmt 包

	"github.com/gorilla/mux" // 导入第三方的 mux 路由包
)

func main() {
	fmt.Println("Hello, Go!") // 使用 fmt 包中的 Println 函数

	// 使用 gorilla/mux 包中的功能来创建路由
	r := mux.NewRouter()
	fmt.Println(r) // 这里只是简单的输出路由对象，通常你会配置路由器并启动服务
}
