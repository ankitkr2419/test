#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <termios.h>
#include <time.h>
#include <unistd.h>
#include <errno.h>
#include <stdbool.h>
#include <pthread.h>
#include <sys/types.h>
#include <sys/stat.h>
#include "MePort.h"


#define DEVICE "/dev/ttyUSB"
static int FDSerial;

static void WriteDataToDebugFile(char *prefix, char *in);
static void* recvData(void* arg);
static int Baudrate(int Speed);


void ComPort_Open(int PortNr, int Speed)
{
	struct termios Settings;
	char DeviceName[30];
	sprintf(DeviceName, "%s%d", DEVICE, PortNr);

	FDSerial = open(DeviceName, O_RDWR | O_NOCTTY);

	if (FDSerial == -1)
	{
		printf("ComPort %s opening faild!\n", DeviceName);
		switch(errno)
		{
			case EACCES:
				printf("Have no permission to open %s\n", DeviceName); break;
			case ENOENT:
				printf("There is no Device %s\n", DeviceName); break;
			default:
				printf("Unknown Error: ERRNO is %d\n", errno); break;
		}
		return;
	}
	printf("ComPort %s Opened!\n", DeviceName);

	//Load Port Settings
	tcgetattr(FDSerial, &Settings);

	//Set Baudrate
	cfsetspeed(&Settings, Baudrate(Speed));

	//Set Read and Local
	Settings.c_cflag |= (CLOCAL | CREAD);

	//set: 8N1
	Settings.c_cflag &= ~PARENB;
	Settings.c_cflag &= ~CSTOPB;
	Settings.c_cflag &= ~CSIZE;
	Settings.c_cflag |= CS8;

	//Read Raw Data
	Settings.c_lflag &= ~(ICANON | ISIG);
	Settings.c_iflag |= IGNPAR;
	//Disable Software Flow Control
	Settings.c_iflag &= ~(IXON | IXOFF | IXANY);

	//Write Raw Data
	Settings.c_oflag &= ~OPOST;

	//Clear Buffer
	tcflush(FDSerial, TCIOFLUSH);
	//Write Settings to ComPort
	tcsetattr(FDSerial, TCSANOW, &Settings);

	//Receive Task
	pthread_t thread;
	pthread_create(&thread, NULL, &recvData, NULL);
}

void ComPort_Close(void)
{
	close(FDSerial);
}

void ComPort_Send(char *in)
{
	WriteDataToDebugFile("OUT: ", in);

	size_t size;
	int len = strlen(in);
	size = write(FDSerial, in, len);
	if(size != len) return; // Data Write Error
}

static void* recvData(void* arg)
{
	char Buffer[100];
	ssize_t bytes_read;

	while(true)
	{
		bytes_read = read(FDSerial, Buffer, sizeof(Buffer)-1);
		Buffer[bytes_read] = 0;
		WriteDataToDebugFile("IN:  ", Buffer);
		MePort_ReceiveByte((int8_t*)Buffer);
	}
	pthread_exit((void *)pthread_self());
}

static void WriteDataToDebugFile(char *prefix, char *in)
{
	FILE *fd;
	char tecPathBuf[100];
	char comPathBuf[100];
	char fullPath[130];
	time_t seconds; 
	struct stat st = {0};

	sprintf(tecPathBuf, "%s/%s", getenv("HOME"), "cpagent/utils/tec");

	if (stat(tecPathBuf, &st) == -1) {
		mkdir(tecPathBuf, 0755);
	}

	// length of seconds is 10, we will only use first 6 digits, 
	// thus tec Com logs will be generated along 10000 secs
    seconds = time(NULL);
	seconds = seconds/10000;
	sprintf(comPathBuf, "ComLog_%ld.log", seconds);

	sprintf(fullPath, "%s/%s",tecPathBuf, comPathBuf);

	fd = fopen(fullPath, "a+");
	if(fd >= 0)
	{
		int len = strcspn(in, "\r\n");
		seconds = time(NULL);
		char format[40] = "";
		sprintf(format, "%%d : %%s%%.%ds\n", len);
		fprintf(fd, format, seconds, prefix, in);
		fclose(fd);
	}
}

int Baudrate(int Speed)
{
	switch(Speed)
	{
		case 1200: return B1200;
		case 1800: return B1800;
		case 2400: return B2400;
		case 4800: return B4800;
		case 9600: return B9600;
		case 19200: return B19200;
		case 38400: return B38400;
		case 57600: return B57600;
		case 115200: return B115200;
		case 230400: return B230400;
		case 460800: return B460800;
		case 500000: return B500000;
		case 576000: return B576000;
		case 921600: return B921600;
		case 1000000: return B1000000;
		default:
			printf("Baudrate not supported\n");
			return B0;
	}
	return B0;
}
