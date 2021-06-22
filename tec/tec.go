package tec


/*
int DemoFunc();
#include <stdlib.h>
#include <time.h>
#include <fcntl.h>
#include <termios.h>
#include <unistd.h>
#include <errno.h>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#cgo CFLAGS  : -std=gnu99 -Wall -g -O3
#cgo LDFLAGS : -pthread -lrt
*/
import "C"

func ConnectTEC(){

	C.DemoFunc()
}
