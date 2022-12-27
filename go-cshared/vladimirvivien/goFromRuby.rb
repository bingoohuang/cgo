require 'ffi_c'

module GoFuncs
  extend FFI::Library
  ffi_lib './awesome.so'
  attach_function :GoHTTP, [:string], :string
end

puts GoFuncs.GoHTTP("https://www.baidu.com/")

# ruby goFromRuby.rb
# https://github.com/bkendzior/golang_shared_libraries
