# go-cshared

Calling Go Functions from Other Languages using C Shared Libraries

1. [vladimirvivien/go-cshared-examples](./vladimirvivien/) [origin repo](https://github.com/vladimirvivien/go-cshared-examples)

## thanks

1. [draffensperger/go-interlang](https://github.com/draffensperger/go-interlang)
2. [johncburns1/Golang-.NET-SharedC](https://github.com/johncburns1/Golang-.NET-SharedC)
3. [Calling Go from C++](https://darkcoding.net/software/building-shared-libraries-in-go-part-2/)
4. [Using Go Dynamic and Static Shared libraries in C/C++ Programs](https://xuri.me/2022/04/15/go-shared-libraries.html)

## blog

1. [Using Go's runtime/cgo to pass values between C and Go](blog/runtime-cgo.md)

## Calling Go from C++

concat/main.go:

```go
package main

import "C"

//export Concat
func Concat(sIn string, bIn []byte, bOut []byte) {
    n := copy(bOut, sIn)
    copy(bOut[n:], bIn)
}

func main() {}
```

1. `go build -buildmode=c-shared -o libconcat.so concat`

concat.cpp

```cpp
#include <vector>
#include <string>
#include <iostream>
#include "libconcat.h"

int main() {
    std::string s_in {"Hello "};
    std::vector<char> v_in {'W', 'o', 'r', 'l', 'd'};
    std::vector<char> v_out(11);

    GoString go_s_in{&s_in[0], static_cast<GoInt>(s_in.size())};
    GoSlice go_v_in{
        v_in.data(),
        static_cast<GoInt>(v_in.size()),
        static_cast<GoInt>(v_in.size()),
    };
    GoSlice go_v_out{
        v_out.data(),
        static_cast<GoInt>(v_out.size()),
        static_cast<GoInt>(v_out.size()),
    };

    Concat(go_s_in, go_v_in, go_v_out);

    for(auto& c : v_out) {
        std::cout << c;
    }
    std::cout << '\n';
}
```

```sh
g++ --std=c++14 concat.cpp -o concat -lconcat
./concat
```

## Go build 模式之 c-archive，c-shared，linkshared

[blog](https://davidchan0519.github.io/2019/04/05/go-buildmode-c/)

本文用实例介绍这 3 种模式的构建方法；

go 程序中需要被导出的方法 Hello ()：

```go
//hello.go
package main

import "C"
import "fmt"

func main() {
 Hello("hello")
}

//export Hello
func Hello(name string) {
 fmt.Println("output:",name)
}
```

注意事项：

1. `//export Hello`，这是约定，所有需要导出给 C 调用的函数，必须通过注释添加这个构建信息，否则不会构建生成 C 所需的头文件；
2. 导出函数的命名的首字母是否大写，不受 go 规则的影响，大小写均可；
3. 头部必须添加 import “C”;

### c-archive 模式

`$ go build --buildmode=c-archive hello.go`

编译生成 hello.a 和 C 头文件 hello.h

在所生成的 hello.h 的头文件中，我们可以看到 Go 的 Hello() 函数的定义：

```cpp
#ifdef __cplusplus
extern "C" {
#endif
...
extern void Hello(GoString p0);
...
#ifdef __cplusplus
}
#endif
```

在头文件中 GoString 的定义如下：

```cpp
typedef struct {
  const char *p;
  ptrdiff_t   n;
} _GoString_;


typedef _GoString_ GoString;
```

然后我们可以在 hello.c 中引用该头文件，并链接 Go 编译的静态库：

```c
#include <stdio.h>
#include <string.h>
#include "hello.h"

int main(int argc, char *argv[]){
    char *type = argv[1];
    char buf[] = "test for ";
    char input[16]={0};

    sprintf(input, "%s %s", buf, type);

    GoString str;
    str.p = input;
    str.n = strlen(input);

    Hello(str);
    return 0;
}
```

然后，构建 C 程序：

`$ gcc  -o hello_static hello.c hello.a -lpthread`

### c-shared 模式

`$ go build --buildmode=c-shared -o libhello.so hello.go`

编译生成 libhello.so

构建 c 程序，并链接动态库

`$ gcc  -o hello_dynamic hello.c  -L./ -lhello -lpthread`

1. `-L` 指定链接库为当前目录
1. `-l` 指定所需要连接的动态库名称，去掉 lib 前缀和.so 后缀

### linkshared 模式

该模式主要用于为 go 程序编译动态链接库

```sh
go install -buildmode=shared std
go build -linkshared -o hello_linkshared hello.go
```

编译生成 hello_linkshared 可执行程序

附上实例程序编译脚本

```sh
# !/bin/bash

# static
go build --buildmode=c-archive hello.go
gcc  -o hello_static hello.c hello.a -lpthread
echo "=== run static binary ==="
./hello_static static

# dynamic
go build --buildmode=c-shared -o libhello.so hello.go
gcc  -o hello_dynamic hello.c -L./ -lhello -lpthread
echo "=== run dynamic binary ==="
./hello_dynamic dynamic

# linkshared for go
sudo /usr/local/go/bin/go install -buildmode=shared std
sudo /usr/local/go/bin/go build -linkshared -o hello_linkshared hello.go
echo "=== run linkshared binary ==="
ls -lh ./hello_linkshared
./hello_linkshared

# default , static for go
go build -o hello_static_go hello.go
ls -lh  ./hello_static_go
./hello_static_go
```
