# Calling Go Functions from Other Languages using C Shared Libraries

This respository contains source examples for the article [*Calling Go Functions from Other Languages*](https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf#.n73as5d6d) (medium.com).  Using the `-buildmode=c-shared` build flag, the compiler outputs a standard shared object binary file (.so) exposing Go functions as a C-style APIs. This lets programmers create Go libraries that can be called from other languages including C, Python, Ruby, Node, and Java (see contributed example for Lua) as done in this repository.

## The Go Code

First, let us write the Go code. Assume that we have written an `awesome` Go library that we want to make available to other languages. There are four requirements to follow before compiling the code into a shared library:

* The package must be amain  package. The compiler will build the package and all of its dependencies into a single shared object binary.
* The source must import the pseudo-package “C”.
* Use the //export comment to annotate functions you wish to make accessible to other languages.
* An empty main function must be declared.

The following Go source exports four functions `Add`, `Cosine`, `Sort`, and `Log`. Admittedly, the awesome library is not that impressive. However, its diverse function signatures will help us explore type mapping implications.

File [awesome.go](./awesome/awesome.go)

The package is compiled using the `-buildmode=c-shared` build flag to create the shared object binary:

`cd awesome && go build -o awesome.so -buildmode=c-shared`

Upon completion, the compiler outputs two files: `awesome.h`, a C header file and `awesome.so`, the shared object file, shown below:

```sh
$ ls -hl
total 1.9M
-rw-rw-r-- 1 d5k d5k  550 Apr 15  2021 awesome.go
-rw-rw-r-- 1 d5k d5k 1.8K Nov  6 19:19 awesome.h
-rw-rw-r-- 1 d5k d5k 1.9M Nov  6 19:19 awesome.so
-rw-rw-r-- 1 d5k d5k   24 Nov  6 19:09 go.mod
```

Notice that the `.so` file is around 2 Mb, relatively large for such a small library. This is because the entire Go runtime machinery and dependent packages are crammed into a single shared object binary (similar to compiling a single static executable binary).

## The header file

[The header file](./awesome.h) defines C types mapped to Go compatible types using cgo semantics.

Next we compile the C code, specifying the shared object library:

`gcc -o client client1.c ./awesome.so`

When the resulting binary is executed, it links to the awesome.so library, calling the functions that were exported from Go as the output shows below.

```shell
$ ./client
awesome.Add(12,99) = 111
awesome.Cosine(1) = 0.540302
awesome.Sort(77,12,5,99,28,23): 5,12,23,28,77,99,
Hello from C!
```

## Dynamically Loaded

In this approach, the C code uses the dynamic link loader library (`libdl.so`) to dynamically load and bind exported symbols. It uses functions defined in `dhfcn.h` such as `dlopen` to open the library file, `dlsym` to look up a symbol, `dlerror` to retrieve errors, and `dlclose` to close the shared library file.

Because the binding and linking is done in your source code, this version is lengthier. However, it is doing the same thing as before, as highlighted in the following snippet (some print statements and error handling omitted).

File [client2.c](./client2.c)

In the previous code, we define our own subset of Go compatible C types `go_int`, `go_float`, `go_slice`, and `go_str`. We use `dlsym` to load symbols `Add`, `Cosine`, `Sort`, and `Log` and assign them to their respective function pointers. Next, we compile the code linking it with the `dl` library (not the awesome.so) as follows:

`gcc -o client client2.c -ldl`

When the code is executed, the C binary loads and links to shared library awesome.so producing the following output:

```shell
$ ./client
awesome.Add(12, 99) = 111
awesome.Cosine(1) = 0.540302
awesome.Sort(44,23,7,66,2): 2,7,23,44,66,
Hello from C!
```

## From Python

In Python things get a little easier. We use can use the `ctypes` [foreign function library](https://docs.python.org/3/library/ctypes.html) to call Go functions from the the awesome.so shared library as shown in the following snippet (some print statements are omitted).

File [client.py](./client.py)

Note the `lib` variable represents the loaded symbols from the shared object file. We also defined Python classes `GoString` and `GoSlice` to map to their respective C struct types. When the Python code is executed, it calls the Go functions in the shared object producing the following output:

```shell
$ python3 client.py
awesome.Add(12,99) = 111
awesome.Cosine(1) = 0.540302
awesome.Sort(74,4,122,9,12) = [4, 9, 12, 74, 122]
Hello Python!
log id 1
```

### Python CFFI (contributed)

The following example was contributed by [@sbinet](https://github.com/sbinet) (thank you!)

Python also has a portable CFFI library that works with Python2/Python3/pypy unchanged.  The
following example uses a C-wrapper to defined the exported Go types.  This makes
the python example less opaque and even easier to understand.

File [client-cffi.py](./client-cffi.py)

```shell
$ python3 client-cffi.py
awesome.Add(12,99) = 111
awesome.Cosine(1) = 0.540302
awesome.Sort(74,4,122,9,12) = [4, 9, 12, 74, 122]
Hello Python!
log id 1
```

## From Ruby

Calling Go functions from Ruby follows a similar pattern as above. We use the the [FFI gem](https://github.com/ffi/ffi) to dynamically load and call exported Go functions in the awesome.so shared object file as shown in the following snippet.

File [client.rb](./client.rb)

In Ruby, we must extend the `FFI` module to declare the symbols being loaded from the shared library. We use Ruby classes `GoSlice` and `GoString` to map the respective C structs. When we run the code it calls the exported Go functions as shown below:

```shell
$ ruby client.rb
awesome.Add(12, 99) = 111
awesome.Cosine(1) = 0.5403023058681398
awesome.Sort([92, 101, 3, 44, 7]) = [3, 7, 44, 92, 101]
Hello Ruby!
```

## From Node

For Node, we use a foreign function library called [node-ffi](https://github.com/node-ffi/node-ffi) (and a couple dependent packages) to dynamically load and call exported Go functions in the awesome.so shared object file as shown in the following snippet:

File [client.js](./client.js)

Node uses the `ffi` object to declare the loaded symbols from the shared library . We also use Node struct objects `GoSlice` and `GoString` to map to their respective C structs. When we run the code it calls the exported Go functions as shown below:

```shell
awesome.Add(12, 99) =  111
awesome.Cosine(1) =  0.5403023058681398
awesome.Sort([12,54,9,423,9] =  [ 0, 9, 12, 54, 423 ]
Hello Node!
```

## From Java

To call the exported Go functions from Java, we use the [Java Native Access library](https://github.com/java-native-access/jna) or JNA as shown in the following code snippet (with some statements omitted or abbreviated):

File [Client.java](./Client.java)

To use JNA, we define Java interface `Awesome` to represents the symbols loaded from the awesome.so shared library file. We also declare classes `GoSlice` and `GoString` to map to their respective C struct representations. When we compile and run the code, it calls the exported Go functions as shown below:

```shell
$ javac -cp jna.jar Client.java
$ java -cp .:jna.jar Client
awesome.Add(12, 99) = 111
awesome.Cosine(1.0) = 0.5403023058681398
awesome.Sort(53,11,5,2,88) = [2 5 11 53 88 ]
Hello Java!
```

## From Lua (contributed)

This example was contributed by [@titpetric](https://github.com/titpetric). See his insightful write up on [*Calling Go functions from LUA*](https://scene-si.org/2017/03/13/calling-go-functions-from-lua/).

The forllowing shows how to invoke exported Go functions from Lua. As before, it uses an FFI library to dynamically load the shared object file and bind to the exported function symbols.

File [client.lua](./client.lua)

When the example is executed, it produces the following:

```
$ luajit client.lua
awesome.Add(12, 99) = 111.000000
awesome.Cosine(1) = 0.540302
awesome.Sort([12,54,9,423,9] = 0, 9, 12, 54, 423
Hello LUA!
```

## From Julia (Contributed)

The following example was contributed by [@r9y9](https://github.com/r9y9). It shows how to invoke exported Go functions from the Julia language. As [documented here](https://docs.julialang.org/en/stable/manual/calling-c-and-fortran-code/), Julia has the capabilities to invoke exported functions from shared libraries similar to other languages discussed here.

File [client.jl](./client.jl)

When the example is executed, it produces the following:

```shell
$ julia client.jl
awesome.Add(12, 9) = 21
awesome.Cosine(1) = 0.5403023058681398
awesome.Sort([77, 12, 5, 99, 28, 23]) = [5, 12, 23, 28, 77, 99]
Hello from Julia!
```

## From Dart (Contributed)

The following example was contributed by @dpurfield. It shows how to invoke exported Go functions from the Dart language. As documented [here](https://dart.dev/guides/libraries/c-interop), Dart has the capability to invoke exported functions from shared libraries similar to other languages discussed here.

File [client.dart](./client.dart)

## From C\#

To call the exported Go functions from C# we use the [DllImportAttribute](https://docs.microsoft.com/en-us/dotnet/api/system.runtime.interopservices.dllimportattribute?view=netframework-4.8) attribute to dynamically load and call exported Go functions in the awesome.so shared object file as shown in the following snippet.

File [client.cs](./client.cs)

When the example is executed, it produces the following:

```sh
> dotnet run
awesome.Add(12,99) = 111
awesome.Cosine(1) = 0,5403023058681398
awesome.Sort(77,12,5,99,28,23): 5, 12, 23, 28, 77, 99
Hello from C#!
```

## Conclusion

This repo shows how to create a Go library that can be used from C, Python, Ruby, Node, Java, Lua, Julia. By compiling Go packages into C-style shared libraries, Go programmers have a powerful way of integrating their code with any modern language that supports dynamic loading and linking of shared object files.
