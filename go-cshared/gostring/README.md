# GoString 在 C# 和 Cgo 中传递的正确方式

```go
package main

// #include <stdint.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

//export PrintHello10
func PrintHello10(a string, b *string) int32 {
	*b = fmt.Sprintf("你好: %s，今天11月18日，见此通知后，立即提前下班。员工需带电脑回家，以备居家办公。！！", a)
	return 0
}

func main() {}
```

```cs
using System;
using System.Runtime.InteropServices;
using System.Text;

namespace GoSharedDLL
{
    class Program
    {
        public struct GoString
        {
            public IntPtr p;

            public int length;

            public static implicit operator GoString(string s)

            {
                int len = Encoding.UTF8.GetByteCount(s);
                byte[] buffer = new byte[len + 1];
                Encoding.UTF8.GetBytes(s, 0, s.Length, buffer, 0);
                IntPtr nativeUtf8 = Marshal.AllocHGlobal(buffer.Length);
                Marshal.Copy(buffer, 0, nativeUtf8, buffer.Length);

                return new GoString { p = nativeUtf8, length = len };
            }

            public static implicit operator string(GoString s)

            {
                byte[] buffer = new byte[s.length];
                Marshal.Copy(s.p, buffer, 0, buffer.Length);
                return Encoding.UTF8.GetString(buffer);
            }
        }

        [
            DllImport(
                "shared.dll",
                CharSet = CharSet.Unicode,
                CallingConvention = CallingConvention.StdCall)
        ]
        public static extern int
        PrintHello10(GoString data, ref GoString result);

        static void Main(string[] args)
        {
            string res = args[0];
            Console.WriteLine("Input: " + res);

            GoString p10input = res;
            GoString p10 = new GoString();
            PrintHello10(p10input, ref p10);
            Console.WriteLine("PrintHello10 Returns: " + p10);
        }
    }
}
```

1. 编译 cgo dll: `go build -o shared.dll -buildmode=c-shared`，会生成 shared.dll 和 shared.h 文件
2. 编译 c# : `csc a.cs`，会生成 a.exe 文件
3. 运行 a.exc: `./a.exe 中华人民共和国万岁`

```log
PS C:\cgo> go build -o shared.dll -buildmode=c-shared
PS C:\cgo> ls
Mode                LastWriteTime     Length Name
----                -------------     ------ ----
-a---        2022-11-18     17:09       3561 a.cs
-a---        2022-11-18     17:20       5120 a.exe
-a---        2022-11-18     17:19       1159 a.go
-a---        2022-11-18     16:01         20 go.mod
-a---        2022-11-18     16:40    3303269 shared.dll
-a---        2022-11-18     16:40       2168 shared.h

PS C:\cgo> go build -o shared.dll -buildmode=c-shared
PS C:\cgo> ./a 孙悟空
Input: 孙悟空
PrintHello10 Returns: 你好: 孙悟空，今天11月18日，见此通知后，立即提前下班。员工需带电脑回家，以备居家办公。！！
```

## 感谢

1. [WIN7 用 CMD 编译简单的 C# 文件](https://blog.csdn.net/qq_37690519/article/details/87938895) WIN7上 找到 Windows 自带的 C# 编译器 csc.exe，路径一般是：C:\Windows\Microsoft.NET\Framework64\v4.0.30319（每个人的版本可能不同）
2. [Net 中调用 Go 语言 DLL 时使用 GoString 的正确姿势（解决.Net 中调用 Go 语言 DLL 时中文乱码和 Unicode 的异常问题）] (http://www.meilongkui.com/archives/1024)
3. [Conversion in .net: Native Utf-8 <-> Managed String](https://stackoverflow.com/questions/10773440/conversion-in-net-native-utf-8-managed-string)
