/*
Illustrates how to get arguments from the command-line.
Usage:

    args 10.0
will print out the number you entered. If not given any
argument, will print out usage info.
*/

#include <stdio.h>

int calSum(int c) {
	int sum = 0;
	for(int i=0; i<=c; i++ ){
        sum=sum+i;
    }

    return sum;
}


int main(int argc, char *argv[]){
    int arg;
    int result;
    while(1) {
        fscanf(stdin, "%d", &arg);
        if (arg == 0) {
            break;
        }
        int result = calSum(arg);
        printf("%d\nEOF\n", result);
        fflush(stdout);
    }

//    int arg;
//    int result;
//
//    if ( argc != 2 ) {
//        printf( "usage: %s float \n", argv[0] );
//    }
//    else {
//        sscanf(argv[1], "%d", &arg); // reads command-line argument
//        result = calSum(arg);
//        printf("%d\nEOF\n", result);
//    }
//
//    return(0);
}

