package calsum

import (
	"bufio"
	"io"
	"os/exec"
)

/*

static void calSum(int c) {
	int sum = 0;
	for(int i=0; i<=c; i++ ){
        sum=sum+i;
    }
}

*/
// #cgo LDFLAGS: -lstdc++
import "C"

func CgocalSum(c int) {
	C.calSum(C.int(c))
}

func GoCalSum(c int) {
	sum := 0
	for i := 0; i <= c; i++ {
		sum += i
	}
}

func revoke(in, out chan string) error {
	cmd := exec.Command("./a.out")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case input, ok := <-in:
				if !ok {
					return
				}
				io.WriteString(stdin, input+"\n")
			}
		}
	}()

	go func() {
		result := ""
		r := bufio.NewReader(stdout)
		for {
			line, _ := r.ReadString('\n')
			if line == "EOF\n" {
				out <- result
				result = ""
			} else {
				result += line
			}
		}
	}()

	return cmd.Start()
}
