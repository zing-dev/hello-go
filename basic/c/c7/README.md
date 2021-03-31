# `CGO`

## `go` 调用 `c` 函数

### 调用原生函数

> 参数
`go 方式 表现 c 类型 C.类型`

- 引入头文件
- `imoport C`
- C.函数名()调用

### 调用自定义函数
```
/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
        printf("%s", s);
}
*/
```

- 引入头文件
- 定义函数
- `imoport C`
- C.函数名()调用
