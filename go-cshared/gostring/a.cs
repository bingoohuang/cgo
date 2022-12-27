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

        [
            DllImport(
                "shared.dll",
                CharSet = CharSet.Unicode,
                CallingConvention = CallingConvention.StdCall)
        ]
        public static extern int
        PrintHello20([In] byte[] data, ref IntPtr output);

        [
            DllImport(
                "shared.dll",
                CharSet = CharSet.Unicode,
                CallingConvention = CallingConvention.StdCall)
        ]
        public static extern IntPtr PrintHello40(byte[] data);

        [
            DllImport(
                "shared.dll",
                CharSet = CharSet.Unicode,
                CallingConvention = CallingConvention.StdCall)
        ]
        public static extern int
        PrintHello21([In] byte[] data, ref IntPtr output);

        [
            DllImport(
                "shared.dll",
                CharSet = CharSet.Unicode,
                CallingConvention = CallingConvention.StdCall)
        ]
        public static extern IntPtr PrintHello41(byte[] data);

        static void Main(string[] args)
        {
            string res = args[0];
            Console.WriteLine("Input: " + res);

            GoString p10input = res;
            GoString p10 = new GoString();
            PrintHello10(p10input, ref p10);
            Console.WriteLine("PrintHello10 Returns: " + p10);

            IntPtr output = IntPtr.Zero;
            IntPtr output1 = IntPtr.Zero;
            var input = Encoding.UTF8.GetBytes(res);

            var a = PrintHello40(input);
            Console
                .WriteLine("PrintHello40 Returns: " +
                Marshal.PtrToStringAnsi(a));

            var a1 = PrintHello41(input);
            Console
                .WriteLine("PrintHello41 Returns: " +
                Marshal.PtrToStringAnsi(a1));

            var i = PrintHello20(input, ref output);
            Console.WriteLine("PrintHello20 Returns: " + i);
            Console
                .WriteLine("Ref Val changed to: " +
                Marshal.PtrToStringAnsi(output, i));

            var i1 = PrintHello21(input, ref output1);
            Console.WriteLine("PrintHello21 Returns: " + i1);
            Console
                .WriteLine("Ref Val changed to: " +
                Marshal.PtrToStringAnsi(output1, i1));
        }
    }
}
