# Jotang 2020 #10 做甜点的怎么能不懂顾客的想法呢！！！

## 前置条件

- Go语言语法基础
- 计算机网络基础（其实啥也妹有也能做···）

## 过程

### 初步搭建

```go
// verson 0.1
package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler函数，接收responseWriter结构和http里的Request结构，功能是处理收到的请求
func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Golang yyds！")
}

func main() {
	// handleFunc函数，起到路由效果，确定该路径下应该使用哪个处理函数
	http.HandleFunc("/", handler)
	fmt.Println("service port:8001")
	// err为设置监听端口返回的结果，如果失败会有内容
	err := http.ListenAndServe(":8001",nil)
	if err != nil {
		// 本行是print在命令行中的
		log.Fatal("ListenAndServe: ", err)
	}
}
```

实现效果如下：

cmd命令行：

![image-20200927154639315](C:\Users\15567\AppData\Roaming\Typora\typora-user-images\image-20200927154639315.png)

网页访问结果：

![image-20200927154721779](C:\Users\15567\AppData\Roaming\Typora\typora-user-images\image-20200927154721779.png)

### 任务一

