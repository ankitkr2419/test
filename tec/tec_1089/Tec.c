#include "Tec.h"
#include "ConsoleIO.h"
#include "ComPort.h"
#include "MeCom.h"

int32_t Address;
int8_t Buf[25];
int32_t Inst = 1;
int ComPortNr = 0;
int BaudRate = 57600;         

int initiateTEC()
{
    MeParFloatFields Fields;
    MeParLongFields Longs;

    ComPort_Open(ComPortNr, BaudRate);
    
    Address = 2;

    if(MeCom_GetIdentString(Address, Buf)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("Ident String: %s\n", Buf);
    }

    Fields.Value = 4;
    if(MeCom_TEC_Ope_CurrentLimitation(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Ope_CurrentLimitation(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Ope_CurrentLimitation value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

    Fields.Value = 20;
    if(MeCom_TEC_Ope_VoltageLimitation(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Ope_VoltageLimitation(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Ope_VoltageLimitation value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

    Fields.Value = 4.8;
    if(MeCom_TEC_Ope_CurrentErrorThreshold(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Ope_CurrentErrorThreshold(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Ope_CurrentErrorThreshold value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

    Fields.Value = 21;
    if(MeCom_TEC_Ope_VoltageErrorThreshold(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Ope_VoltageErrorThreshold(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Ope_VoltageErrorThreshold value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

    // Might need to set Proximity to 0
    Fields.Value = 0;
    if(MeCom_TEC_Tem_ProximityWidth(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Tem_ProximityWidth(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Tem_ProximityWidth value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }


    // Peltier Full Control
    Longs.Value = 0;

    if(MeCom_TEC_Tem_ModelizationMode(Address, Inst, &Longs, MeGetLimits)){
        MeCom_TEC_Tem_ModelizationMode(Address, Inst, &Longs, MeSet);
    } else{
        printf("TEC MeCom_TEC_Tem_ModelizationMode value out of range( %d, %d).", Longs.Max, Longs.Min);
        return -1;
    }

    // Temperature Controller
    Longs.Value = 2;

    if(MeCom_TEC_Ope_OutputStageInputSelection(Address, Inst, &Longs, MeGetLimits)){
        MeCom_TEC_Ope_OutputStageInputSelection(Address, Inst, &Longs, MeSet);
    } else{
        printf("TEC MeCom_TEC_Ope_OutputStageInputSelection value out of range( %d, %d).", Longs.Max, Longs.Min);
        return -1;
    }

/*
    Fields.Value = 44.40232;
    if(MeCom_TEC_Tem_Kp(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Tem_Kp(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Tem_Kp value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

    Fields.Value = 0.7976229;
    if(MeCom_TEC_Tem_Ti(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Tem_Ti(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Tem_Ti value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

    Fields.Value = 0.1914295;
    if(MeCom_TEC_Tem_Td(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Tem_Td(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Tem_Td value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

    Fields.Value = 0.3;
    if(MeCom_TEC_Tem_DPartDampPT1(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Tem_DPartDampPT1(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Tem_DPartDampPT1 value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }
*/
    // Peltier Max Current

    Fields.Value = 7;
    if(MeCom_TEC_Tem_PeltierMaxCurrent(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Tem_PeltierMaxCurrent(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Tem_PeltierMaxCurrent value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

    Fields.Value = 75;
    if(MeCom_TEC_Tem_PeltierDeltaTemperature(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Tem_PeltierDeltaTemperature(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC MeCom_TEC_Tem_PeltierDeltaTemperature value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

    // 0 means cooling
    Longs.Value = 0;

    if(MeCom_TEC_Tem_PeltierPositiveCurrentIs(Address, Inst, &Longs, MeGetLimits)){
        MeCom_TEC_Tem_PeltierPositiveCurrentIs(Address, Inst, &Longs, MeSet);
    } else{
        printf("TEC MeCom_TEC_Tem_PeltierPositiveCurrentIs value out of range( %d, %d).", Longs.Max, Longs.Min);
        return -1;
    }

    // Stage Enable
    Longs.Value = 1;

    if (MeCom_TEC_Ope_OutputStageEnable(Address, Inst, &Longs , MeSet)){
        printf("MeCom_TEC_Ope_OutputStageEnable %d", Longs.Value);
    } else{
        printf("TEC MeCom_TEC_Ope_OutputStageEnable value out of range( %d, %d).", Longs.Max, Longs.Min);
        return -1;
    }

    // 5. Read Actual Output Current and Voltage
    int i = 0;
    while(i != 3){
            i++;
            sleep(1); 

            if(MeCom_TEC_Mon_ActualOutputCurrent(Address, Inst, &Fields, MeGet))
            {
                printf("TEC MeCom_TEC_Mon_ActualOutputCurrent: Value: %f\n", Fields.Value);
            } else {
                printf("TEC MeCom_TEC_Mon_ActualOutputCurrent value couldn't be read");
                return -1;
            }

            if(MeCom_TEC_Mon_ActualOutputVoltage(Address, Inst, &Fields, MeGet))
            {
                printf("TEC MeCom_TEC_Mon_ActualOutputVoltage: Value: %f\n", Fields.Value);
            } else {
                printf("TEC MeCom_TEC_Mon_ActualOutputVoltage value couldn't be read");
                return -1;
            }
    }

    return 0;
}


int setTempAndRamp(double target, double ramp)
{   
    MeParFloatFields Fields;
  
    // 1. Setup Target Temp

    Fields.Value = target;
    printf("TEC Object Temperature: New Value: %f\n", Fields.Value);

    // check for limit
    if(MeCom_TEC_Tem_TargetObjectTemp(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Tem_TargetObjectTemp(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC Object Temperature value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

    if (ramp == 0){
        goto skipRamp;
    }
    Fields.Value = ramp;
    printf("TEC Object Ramp Up Rate: New Value: %f\n", Fields.Value);

    // check for limit

    if(MeCom_TEC_Tem_CoarseTempRamp(Address, Inst, &Fields, MeGetLimits)){
        MeCom_TEC_Tem_CoarseTempRamp(Address, Inst, &Fields, MeSet);
    } else{
        printf("TEC Object Ramp Up Rate value out of range( %f, %f).", Fields.Max, Fields.Min);
        return -1;
    }

skipRamp:
    while(1){
            usleep(200 * 1000);

            if(MeCom_TEC_Mon_ObjectTemperature(Address, Inst, &Fields, MeGet))
            {
                printf("TEC Object Temperature: Value: %f\n", Fields.Value);
            } else {
                printf("TEC Object Temperature value couldn't be read");
                return -1;
            }

            // Play of +- 1
            if ( ( (target + 0.25) >= Fields.Value ) && ( (target - 0.25) <= Fields.Value ) ) {
                printf("TEC Object Temperature Reached:: %f, target: %f\n", Fields.Value, target);
                return 0;
            }
    }
    return 0;
}


double getObjectTemp(){
    MeParFloatFields Fields;

getTemp:
    if(MeCom_TEC_Mon_ObjectTemperature(Address, Inst, &Fields, MeGet))
    {
        // TEC will only be valid for temperatures above Room Temp
        if (Fields.Value >= 10) {
            return Fields.Value; 
        }
        // sleep for 200ms
        usleep(200 * 1000);
        if (isnan(Fields.Value) ) {
            printf("TEC Object Temperature value is NaN");
            return -1;
        }
        goto getTemp;
    } else {
        printf("TEC Object Temperature value couldn't be read");
        return -1;
    }
    return Fields.Value;
}


int checkForErrorState()
{

    MeParLongFields Longs; 
    if(MeCom_COM_DeviceStatus(Address, &Longs, MeGet)){
        // print only if its not Run
        if (Longs.Value != 2 && Longs.Value < 11 && Longs.Value >= 0){
            printf("TEC MeCom_COM_DeviceStatus :%d\n", Longs.Value);
        }
    } else{
        printf("TEC Reading MeCom_COM_DeviceStatus error");
        return -1;
    }

// check for Error type
    if (Longs.Value == 3) {
        if(MeCom_COM_ErrorNumber(Address, &Longs, MeGet)){
            printf("TEC MeCom_COM_ErrorNumber :%d\n", Longs.Value);
            return 3;
        } else{
            printf("TEC Reading MeCom_COM_ErrorNumber error");
            return -1;
        }
    }
    return 0;
}

int autoTune(){
    MeParLongFields Longs;
    MeParFloatFields Fields;

    Longs.Value = 1;
    if(MeCom_TEC_Oth_AtmAutoTuningStart(Address, Inst, &Longs, MeGetLimits)){
        MeCom_TEC_Oth_AtmAutoTuningStart(Address, Inst, &Longs, MeSet);
        printf("TEC MeCom_TEC_Oth_AtmAutoTuningStart :%d\n", Longs.Value);
    } else{
        printf("TEC Reading MeCom_TEC_Oth_AtmAutoTuningStart error");
        return -1;
    }

    if(MeCom_TEC_Oth_AtmTuningStatus(Address, Inst, &Longs, MeGet)){
        printf("TEC MeCom_TEC_Oth_AtmTuningStatus :%d\n", Longs.Value);
    } else{
        printf("TEC Reading MeCom_TEC_Oth_AtmTuningStatus error");
        return -1;
    }

    // Other Longs value means its either idle or success or err
    while (Longs.Value > 0 && Longs.Value < 4){
        sleep(5);
        if(MeCom_TEC_Oth_AtmTuningStatus(Address, Inst, &Longs, MeGet)){
            printf("TEC MeCom_TEC_Oth_AtmTuningStatus :%d\n", Longs.Value);
        } else{
            printf("TEC Reading MeCom_TEC_Oth_AtmTuningStatus error");
        return -1;
        }

        //  Read Progress
        if(MeCom_TEC_Oth_AtmTuningProgress(Address, Inst, &Fields, MeGet)){
            printf("TEC MeCom_TEC_Oth_AtmTuningProgress :%f\n", Fields.Value);
        } else{
            printf("TEC Reading MeCom_TEC_Oth_AtmTuningProgress error");
        return -1;
        }
    }

    // Error
    if (Longs.Value == 10){
        printf("TEC Reading Auto Tuning error");
        return -1;
    // Success
    } else if (Longs.Value == 4){
        printf("TEC Auto Tuning SUCCESS");
        return 4;
    // Other
    } else {
        printf("TEC looks to be in idle state");
        return 0;
    }

}

int resetDevice(){
    if (MeCom_ResetDevice(Address)){
        printf("TEC Reset SUCCESS");
        initiateTEC();
    } else {
        printf("TEC Reset FAIL");
        return -1;
    }
    return 0;
}


int getAllTEC(){
    MeParLongFields  lFields;
    MeParFloatFields fFields;

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_COM_DeviceType(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_COM_DeviceType: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_COM_HardwareVersion(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_COM_HardwareVersion: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_COM_SerialNumber(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_COM_SerialNumber: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_COM_FirmwareVersion(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_COM_FirmwareVersion: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_COM_DeviceStatus(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_COM_DeviceStatus: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_COM_ErrorNumber(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_COM_ErrorNumber: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_COM_ErrorInstance(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_COM_ErrorInstance: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_COM_ErrorParameter(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_COM_ErrorParameter: %d\n", lFields.Value);
    }
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_COM_ParameterSystemFlashSaveOff(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_COM_ParameterSystemFlashSaveOff: %d\n", lFields.Value);
    }
        ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_COM_ParameterSystemFlashStatus(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_COM_ParameterSystemFlashStatus: %d\n", lFields.Value);
    }


    //Tab Monitor
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_ObjectTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_ObjectTemperature: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_SinkTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_SinkTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_TargetObjectTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_TargetObjectTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_RampNominalObjectTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_RampNominalObjectTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_ThermalPowerModelCurrent(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_ThermalPowerModelCurrent: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_ActualOutputCurrent(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_ActualOutputCurrent: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_ActualOutputVoltage(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_ActualOutputVoltage: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_PIDLowerLimitation(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_PIDLowerLimitation: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_PIDUpperLimitation(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_PIDUpperLimitation: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_PIDControlVariable(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_PIDControlVariable: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_ObjectSensorRawADCValue(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_ObjectSensorRawADCValue: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_SinkSensorRawADCValue(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_SinkSensorRawADCValue: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_ObjectSensorResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_ObjectSensorResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_SinkSensorResitance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_SinkSensorResitance: %f\n", fFields.Value);
    }
        
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_SinkSensorTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_SinkSensorTemperature: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_FirmwareVersion(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_FirmwareVersion: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_FirmwareBuildNumber(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_FirmwareBuildNumber: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_HardwareVersion(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_HardwareVersion: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_SerialNumber(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_SerialNumber: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_DriverInputVoltage(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_DriverInputVoltage: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_MedVInternalSupply(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_MedVInternalSupply: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_3_3VInternalSupply(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_3_3VInternalSupply: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_BasePlateTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_BasePlateTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_ErrorNumber(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_ErrorNumber: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_ErrorInstance(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_ErrorInstance: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_ErrorParameter(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_ErrorParameter: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_ParallelActualOutputCurrent(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_ParallelActualOutputCurrent: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_DriverStatus(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_DriverStatus: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_FanRelativeCoolingPower(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_FanRelativeCoolingPower: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_FanNominalFanSpeed(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_FanNominalFanSpeed: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_FanActualFanSpeed(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_FanActualFanSpeed: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_FanActualPwmLevel(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_FanActualPwmLevel: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_BasePlateTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_BasePlateTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Mon_TemperatureIsStable(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Mon_TemperatureIsStable: %d\n", lFields.Value);
    }
    
    //Tab Operation
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_OutputStageInputSelection(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_OutputStageInputSelection: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_OutputStageEnable(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_OutputStageEnable: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_SetStaticCurrent(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_SetStaticCurrent: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_SetStaticVoltage(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_SetStaticVoltage: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_CurrentLimitation(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_CurrentLimitation: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_VoltageLimitation(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_VoltageLimitation: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_CurrentErrorThreshold(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_CurrentErrorThreshold: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_VoltageErrorThreshold(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_VoltageErrorThreshold: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_GeneralOperatingMode(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_GeneralOperatingMode: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_DeviceAddress(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_DeviceAddress: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_RS485CH1BaudRate(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_RS485CH1BaudRate: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_RS485CH1ResponseDelay(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_RS485CH1ResponseDelay: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Ope_ComWatchDogTimeout(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Ope_ComWatchDogTimeout: %f\n", fFields.Value);
    }

    //Tab Temperature Control
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_TargetObjectTemp(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_TargetObjectTemp: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_CoarseTempRamp(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_CoarseTempRamp: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_ProximityWidth(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_ProximityWidth: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_Kp(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_Kp: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_Ti(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_Ti: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_Td(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_Td: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_DPartDampPT1(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_DPartDampPT1: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_ModelizationMode(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_ModelizationMode: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_PeltierMaxCurrent(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_PeltierMaxCurrent: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_PeltierMaxVoltage(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_PeltierMaxVoltage: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_PeltierCoolingCapacity(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_PeltierCoolingCapacity: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_PeltierDeltaTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_PeltierDeltaTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_PeltierPositiveCurrentIs(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_PeltierPositiveCurrentIs: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_ResistorResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_ResistorResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Tem_ResistorMaxCurrent(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Tem_ResistorMaxCurrent: %f\n", fFields.Value);
    }
    
    //Tab Object Temperature
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_TemperatureOffset(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_TemperatureOffset: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_TemperatureGain(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_TemperatureGain: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_LowerErrorThreshold(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_LowerErrorThreshold: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_UpperErrorThreshold(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_UpperErrorThreshold: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_MaxTempChange(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_MaxTempChange: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_NTCLowerPointTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_NTCLowerPointTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_NTCLowerPointResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_NTCLowerPointResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_NTCMiddlePointTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_NTCMiddlePointTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_NTCMiddlePointResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_NTCMiddlePointResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_NTCUpperPointTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_NTCUpperPointTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_NTCUpperPointResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_NTCUpperPointResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_StabilityTemperatureWindow(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_StabilityTemperatureWindow: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_StabilityMinTimeInWindow(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_StabilityMinTimeInWindow: %f\n", fFields.Value);
    }

    if(MeCom_TEC_Obj_StabilityMaxStabiTime(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_StabilityMaxStabiTime: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_MeasLowestResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_MeasLowestResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_MeasHighestResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_MeasHighestResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_MeasTempAtLowestResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_MeasTempAtLowestResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Obj_MeasTempAtHighestResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Obj_MeasTempAtHighestResistance: %f\n", fFields.Value);
    }
    
    //Tab Sink Temperature
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_TemperatureOffset(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_TemperatureOffset: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_TemperatureGain(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_TemperatureGain: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_LowerErrorThreshold(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_LowerErrorThreshold: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_UpperErrorThreshold(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_UpperErrorThreshold: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_MaxTempChange(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_MaxTempChange: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_NTCLowerPointTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_NTCLowerPointTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_NTCLowerPointResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_NTCLowerPointResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_NTCMiddlePointTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_NTCMiddlePointTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_NTCMiddlePointResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_NTCMiddlePointResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_NTCUpperPointTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_NTCUpperPointTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_NTCUpperPointResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_NTCUpperPointResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_SinkTemperatureSelection(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_SinkTemperatureSelection: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_FixedTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_FixedTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_MeasLowestResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_MeasLowestResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_MeasHighestResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_MeasHighestResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_MeasTempAtLowestResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_MeasTempAtLowestResistance: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Sin_MeasTempAtHighestResistance(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Sin_MeasTempAtHighestResistance: %f\n", fFields.Value);
    }

    //Tab Expert: Sub Tab Temperature Measurement
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_ObjMeasPGAGain(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_ObjMeasPGAGain: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_ObjMeasCurrentSource(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_ObjMeasCurrentSource: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_ObjMeasADCRs(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_ObjMeasADCRs: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_ObjMeasADCCalibOffset(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_ObjMeasADCCalibOffset: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_ObjMeasADCCalibGain(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_ObjMeasADCCalibGain: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_ObjMeasSensorTypeSelection(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_ObjMeasSensorTypeSelection: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_SinMeasADCRv(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_SinMeasADCRv: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_SinMeasADCVps(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_SinMeasADCVps: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_SinMeasADCCalibOffset(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_SinMeasADCCalibOffset: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_SinMeasADCCalibGain(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_SinMeasADCCalibGain: %f\n", fFields.Value);
    }

    //Tab Expert: Sub Tab Display
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_DisplayType(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_DisplayType: %d\n", lFields.Value);
    }
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_DisplayLineDefText(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_DisplayLineDefText: %d\n", lFields.Value);
    }
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_DisplayLineAltText(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_DisplayLineAltText: %d\n", lFields.Value);
    }
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_DisplayLineAltMode(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_DisplayLineAltMode: %d\n", lFields.Value);
    }
    //Tab Expert: Sub Tab PBC
    for(int32_t Inst=1; Inst<=8; Inst++)
    {
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_PbcFunction(Address, Inst, &lFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_PbcFunction%d: %d\n", Inst, lFields.Value);
        }
    }
    for(int32_t Inst=1; Inst<=2; Inst++)
    {
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_ChangeButtonLowTemperature(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_ChangeButtonLowTemperature%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_ChangeButtonHighTemperature(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_ChangeButtonHighTemperature%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_ChangeButtonStepSize(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_ChangeButtonStepSize%d: %f\n", Inst, fFields.Value);
        }
    }

    //Tab Expert: Sub Tab Fan
    for(int32_t Inst=1; Inst<=2; Inst++)
    {
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanControlEnable(Address, Inst, &lFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanControlEnable%d: %d\n", Inst, lFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanActualTempSource(Address, Inst, &lFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanActualTempSource%d: %d\n", Inst, lFields.Value);
        }

        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanTargetTemp(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanTargetTemp%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanTempKp(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanTempKp%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanTempTi(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanTempTi%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanTempTd(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanTempTd%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanSpeedMin(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanSpeedMin%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanSpeedMax(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanSpeedMax%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanSpeedKp(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanSpeedKp%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanSpeedTi(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanSpeedTi%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanSpeedTd(Address, Inst, &fFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanSpeedTd%d: %f\n", Inst, fFields.Value);
        }
        ConsoleIO_SetColor(ConsoleIO_Red);
        if(MeCom_TEC_Exp_FanSpeedBypass(Address, Inst, &lFields, MeGet)) 
        {
            ConsoleIO_SetColor(ConsoleIO_Green);
            printf("MeCom_TEC_Exp_FanSpeedBypass%d: %d\n", Inst, lFields.Value);
        }
    }
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_PwmFrequency(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_PwmFrequency: %d\n", lFields.Value);
    }
    
    //Tab Expert: Sub Tab Misc
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_MiscActObjectTempSource(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_MiscActObjectTempSource: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_MiscDelayTillReset(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_MiscDelayTillReset: %d\n", lFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Exp_MiscError108Delay(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Exp_MiscError108Delay: %d\n", lFields.Value);
    }

    //Other Parameters (Not directly displayed in the Service Software)
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_LiveEnable(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_LiveEnable: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_LiveSetCurrent(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_LiveSetCurrent: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_LiveSetVoltage(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_LiveSetVoltage: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_SineRampStartPoint(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_SineRampStartPoint: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_ObjectTargetTempSourceSelection(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_ObjectTargetTempSourceSelection: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_ObjectTargetTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_ObjectTargetTemperature: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmAutoTuningStart(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmAutoTuningStart: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmAutoTuningCancel(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmAutoTuningCancel: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmThermalModelSpeed(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmThermalModelSpeed: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmTuningParameter2A(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmTuningParameter2A: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmTuningParameter2D(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmTuningParameter2D: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmTuningParameterKu(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmTuningParameterKu: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmTuningParameterTu(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmTuningParameterTu: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmPIDParameterKp(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmPIDParameterKp: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmPIDParameterTi(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmPIDParameterTi: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmPIDParameterTd(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmPIDParameterTd: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmSlowPIParameterKp(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmSlowPIParameterKp: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmSlowPIParameterTi(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmSlowPIParameterTi: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmPIDDPartDamping(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmPIDDPartDamping: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmCoarseTempRamp(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmCoarseTempRamp: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmProximityWidth(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmProximityWidth: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmTuningStatus(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmTuningStatus: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_AtmTuningProgress(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_AtmTuningProgress: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_LutTableStart(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_LutTableStart: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_LutTableStop(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_LutTableStop: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_LutTableStatus(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_LutTableStatus: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_LutCurrentTableLine(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_LutCurrentTableLine: %d\n", lFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_LutTableIDSelection(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_LutTableIDSelection: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_LutNrOfRepetitions(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_LutNrOfRepetitions: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_PbcEnableFunction(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_PbcEnableFunction: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_PbcSetOutputToPushPull(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_PbcSetOutputToPushPull: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_PbcSetOutputStates(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_PbcSetOutputStates: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_PbcReadInputStates(Address, 1, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_PbcReadInputStates: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_TEC_Oth_ExternalActualObjectTemperature(Address, 1, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_TEC_Oth_ExternalActualObjectTemperature: %f\n", fFields.Value);
    }
    return 0;
}