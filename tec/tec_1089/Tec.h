#ifndef TEC_H
#define TEC_H

#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <stdint.h>
#include <math.h>

int32_t Address;
int8_t Buf[25];
int32_t Inst = 1;
int ComPortNr = 1;
int BaudRate = 57600;

typedef struct Config
{
    float CurrentLimitation;
    float VoltageLimitation;
    float CurrentErrorThreshold;
    float VoltageErrorThreshold;
    float PeltierMaxCurrent;
    float PeltierDeltaTemperature;
} Config;

#endif
