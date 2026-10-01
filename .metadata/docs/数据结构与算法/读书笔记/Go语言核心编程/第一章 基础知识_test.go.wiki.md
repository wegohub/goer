# 第一章 基础知识_test.go

## 描述

总结第一章基础知识

## 总结

- Go语言诞生背景：为了解决当下编程语言对"并发支持不友好"、“编译速度慢”、“编程复杂”这三个问题而诞生的。
- Go源代码特征
  - 源程序默认为UTF-8编码
  - 语句结尾的分号可以省略
  - 函数以func开头，开头的"{"必须在函数头所在行的尾部，不能单独一行(单独一行就是新的作用域了)
- Go语言的token分为：关键字、标识符、操作符、分隔符和字面常量
  ![image.png](../../../../static/upload/a057d799-03de-4974-8e27-52f977d1ff9b.png)

- Go的标识符必须是字母或者下划线，区分大小写，并且Unicode字符也可以作为标识符的构成，但是不推荐
- Go语言25个关键字：
  - break
  - default
  - func
  - interface
  - select
  - case 
  - defer
  - go 
  - map
  - struct
  - chan
  - else
  - goto
  - package
  - switch
  - const
  - fallthrough
  - if 
  - range
  - type
  - continue
  - for
  - import
  - return
  - var

在Go语言中，`len`函数用于返回不同类型数据结构的长度或元素个数。具体来说：

- 对于**切片（slice）**，`len`返回的是切片中元素的个数，而不是字节数。
- 对于**数组（array）**，`len`返回的是数组中元素的个数，而不是字节数。
- 对于**字符串（string）**，`len`返回的是字符串的字节数，而不是字符数。
- 对于**映射（map）**，`len`返回的是映射中键值对的数量。
- 对于**通道（channel）**，`len`返回的是通道中未读消息的数量。

以下是一些示例代码，展示`len`函数在不同类型上的使用：

```go
package main

import (
    "fmt"
)

func main() {
    // 切片
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(len(slice)) // 输出5

    // 数组
    array := [5]int{1, 2, 3, 4, 5}
    fmt.Println(len(array)) // 输出5

    // 字符串
    str := "不愿意"
    fmt.Println(len(str)) // 输出9（字节数）

    // 映射
    m := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
    }
    fmt.Println(len(m)) // 输出3

    // 通道
    ch := make(chan int, 5)
    ch <- 1
    ch <- 2
    fmt.Println(len(ch)) // 输出2
}
```

总结：
- `len(切片)`返回的是切片中元素的个数。
- `len(字符串)`返回的是字符串的字节数。

如果你需要获取字符串的字符数而不是字节数，可以使用`unicode/utf8`包中的`RuneCountInString`函数，如前所述。
