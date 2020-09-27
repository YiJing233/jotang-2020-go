# Jotang 2020 #10 做甜点的怎么能不懂顾客的想法呢！！！

## 前置条件

- Go语言语法基础
- 计算机网络基础（其实啥也妹有也能做···）

## 过程

### 初步搭建

```go
# version 0.1
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Golang yyds！")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("service port:8001")
	err := http.ListenAndServe(":8001",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

实现效果如下：

cmd命令行：

![image-20200927154639315](C:\Users\15567\AppData\Roaming\Typora\typora-user-images\image-20200927154639315.png)

网页访问结果：

![image-20200927154721779](C:\Users\15567\AppData\Roaming\Typora\typora-user-images\image-20200927154721779.png)

