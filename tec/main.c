/*==============================================================================*/
/** @file       main.c
    @brief      This file is not part of the MeComAPI. Just simple Demo Application.
    @author     Meerstetter Engineering GmbH: Marc Luethi / Thomas Braun
    @version    v0.42

    This is the TOP file of the MeComAPI Demo Application.
    The Demo Application is being used to test the MeComAPI and
    should show the user how the MeComAPI shoud be used.

    The Demo Application has been compiled an the following Product:
    "Microsoft Visual Studio 2010"

    The Demo Application might uses some C++ functions.
    The MeComAPI is written in C (ANSI C99).

    Please refer to the Document 5170-MeComAPI for more information.

*/
/*==============================================================================*/
/*                          IMPORT                                              */
/*==============================================================================*/
#include <stdio.h>
#include <string.h>
#include "ConsoleIO.h"
#include "ComPort.h"
#include "MeCom.h"

/*==============================================================================*/
/*                          DEFINITIONS/DECLARATIONS                            */
/*==============================================================================*/

/*==============================================================================*/
/*                          STATIC FUNCTION PROTOTYPES                          */
/*==============================================================================*/
int DemoFunc();
static int32_t MenuSelection(void);
static void TestAllCommonGetFunctions(uint8_t Address);
static void TestAllLDDGetFunctions(uint8_t Address);
static void TestAllTECGetFunctions(uint8_t Address);

/*==============================================================================*/
/*                          EXTERN VARIABLES                                    */
/*==============================================================================*/

/*==============================================================================*/
/*                          STATIC  VARIABLES                                   */
/*==============================================================================*/

/*==============================================================================*/
/** @brief      Main function of the Demp Application
 *
*/
int main2()
{


    DemoFunc();

    return 0;
 
 
    printf("* Demo Application of the Meerstetter Engineering Communication Protocol API *\n\n");
    printf("If you have any questions, please do not hesitate to contact us under:\ncontact@meerstetter.ch or www.meerstetter.ch\n\n");

    printf("This Demo Application generates a Communication Log file 'ComLog.txt'.\n");
    printf("It is possible to use this Demo Application with or without ComInterface.\n");

    if(ConsoleIO_YesNo("Do you want to open a Comport? (press enter for default)", 1))
    {
        ComPort_Open(ConsoleIO_IntInput("Please enter ComPort number", 0, 50, 1), 
            ConsoleIO_IntInput("Please enter ComPort Speed", 4800, 1000000, 57600));
    }
    else
    {
        printf("No Comport has been opened, this will result in some communications timeouts.\n");
        printf("It is even though possible to see the resulting Host communication\nin the Communication Log file\n\n");
    }
    int32_t Address = ConsoleIO_IntInput("Please Enter the Device Address", 0, 255, 0);
    
   

    while(1)
    {
        switch(MenuSelection())
        {
            case 1: 
            {                
                int8_t Buf[25];
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_GetIdentString(Address, Buf)) 
                {
                    ConsoleIO_SetColor(ConsoleIO_Green);
                    printf("Ident String: %s\n", Buf);
                }          
            }
            break; //--------------------------------------------------------------------------------------
            case 2: 
            {
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_ResetDevice(Address))
                {
                    ConsoleIO_SetColor(ConsoleIO_Green);
                    printf("Device Reset OK.\n");
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 3: 
            {
                int32_t ParId   = ConsoleIO_IntInput("Please Enter Parameter ID", 0, 65535, 0);
                int32_t Inst    = ConsoleIO_IntInput("Please Enter Instance", 1, 255, 1);
                MeParLongFields Fields;
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_ParValuel(Address, ParId, Inst, &Fields, MeGet))
                {
                    ConsoleIO_SetColor(ConsoleIO_Green);
                    printf("Parameter ID: %d; Instance: %d; Value: %d\n", ParId, Inst, Fields.Value);
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 4: 
            {
                int32_t ParId   = ConsoleIO_IntInput("Please Enter Parameter ID", 0, 65535, 0);
                int32_t Inst    = ConsoleIO_IntInput("Please Enter Instance", 1, 255, 1);
                MeParFloatFields Fields;
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_ParValuef(Address, ParId, Inst, &Fields, MeGet))
                {
                    ConsoleIO_SetColor(ConsoleIO_Green);
                    printf("Parameter ID: %d; Instance: %d; Value: %f\n", ParId, Inst, Fields.Value);
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 5: 
            {
                int32_t ParId   = ConsoleIO_IntInput("Please Enter Parameter ID", 0, 65535, 0);
                int32_t Inst    = ConsoleIO_IntInput("Please Enter Instance", 1, 255, 1);
                MeParLongFields Fields;
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_ParValuel(Address, ParId, Inst, &Fields, MeGetLimits))
                {
                    ConsoleIO_SetColor(ConsoleIO_Reset);
                    Fields.Value = ConsoleIO_IntInput("Please Enter the new long Value", Fields.Min, Fields.Max, 0);
                
                    ConsoleIO_SetColor(ConsoleIO_Red);
                    if(MeCom_ParValuel(Address, ParId, Inst, &Fields, MeSet))
                    {
                        ConsoleIO_SetColor(ConsoleIO_Green);
                        printf("Parameter ID: %d; Instance: %d; New Value: %d\n", ParId, Inst, Fields.Value);
                    }
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 6: 
            {
                int32_t ParId   = ConsoleIO_IntInput("Please Enter Parameter ID", 0, 65535, 0);
                int32_t Inst    = ConsoleIO_IntInput("Please Enter Instance", 1, 255, 1);
                MeParFloatFields Fields;
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_ParValuef(Address, ParId, Inst, &Fields, MeGetLimits))
                {
                    ConsoleIO_SetColor(ConsoleIO_Reset);
                    Fields.Value = ConsoleIO_FloatInput("Please Enter the new float Value", Fields.Min, Fields.Max, 0);
                
                    ConsoleIO_SetColor(ConsoleIO_Red);
                    if(MeCom_ParValuef(Address, ParId, Inst, &Fields, MeSet))
                    {
                        ConsoleIO_SetColor(ConsoleIO_Green);
                        printf("Parameter ID: %d; Instance: %d; New Value: %f\n", ParId, Inst, Fields.Value);
                    }
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 10: 
            {
                MeParLongFields Fields;
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_COM_DeviceType(Address, &Fields, MeGet))
                {
                    ConsoleIO_SetColor(ConsoleIO_Green);
                    printf("Device Type: Value: %d\n", Fields.Value);
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 11: 
            {
                MeParLongFields Fields;
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_COM_SerialNumber(Address, &Fields, MeGet))
                {
                    ConsoleIO_SetColor(ConsoleIO_Green);
                    printf("Serial Number: Value: %d\n", Fields.Value);
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 20: 
            {
                MeParFloatFields Fields;
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_LDD_Mon_LaserDiodeCurrent(Address, &Fields, MeGet))
                {
                    ConsoleIO_SetColor(ConsoleIO_Green);
                    printf("LDD Laser Diode Current: Value: %f\n", Fields.Value);
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 21: 
            {
                MeParFloatFields Fields;
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_LDD_Ope_CurrentCW(Address, &Fields, MeGetLimits))
                {
                    ConsoleIO_SetColor(ConsoleIO_Reset);
                    Fields.Value = ConsoleIO_FloatInput("Please Enter New Current CW", Fields.Min, Fields.Max, 0);
                    ConsoleIO_SetColor(ConsoleIO_Red);
                    if(MeCom_LDD_Ope_CurrentCW(Address, &Fields, MeSet))
                    {
                        ConsoleIO_SetColor(ConsoleIO_Green);
                        printf("LDD Current CW: New Value: %f\n", Fields.Value);
                    }
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 22: 
            {
                MeParLongFields Fields;
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_LDD_Ope_EnableInputSource(Address, &Fields, MeGetLimits))
                {
                    ConsoleIO_SetColor(ConsoleIO_Reset);
                    Fields.Value = ConsoleIO_IntInput("Please Enter Enable Input Source", Fields.Min, Fields.Max, 0);
                    ConsoleIO_SetColor(ConsoleIO_Red);
                    if(MeCom_LDD_Ope_EnableInputSource(Address, &Fields, MeSet))
                    {
                        ConsoleIO_SetColor(ConsoleIO_Green);
                        printf("LDD Enable Input Source: New Value: %d\n", Fields.Value);
                    }
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 30: 
            {
                int32_t Inst    = ConsoleIO_IntInput("Please Enter Instance", 1, 255, 1);
                MeParFloatFields Fields;
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_TEC_Mon_ObjectTemperature(Address, Inst, &Fields, MeGet))
                {
                    ConsoleIO_SetColor(ConsoleIO_Green);
                    printf("TEC Object Temperature: Value: %f\n", Fields.Value);
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 31: 
            {
                MeParFloatFields Fields;
                int32_t Inst    = ConsoleIO_IntInput("Please Enter Instance", 1, 255, 1);
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_TEC_Tem_TargetObjectTemp(Address, Inst, &Fields, MeGetLimits))
                {
                    ConsoleIO_SetColor(ConsoleIO_Reset);
                    Fields.Value = ConsoleIO_FloatInput("Please Enter New Target Object Temp", Fields.Min, Fields.Max, 0);
                    ConsoleIO_SetColor(ConsoleIO_Red);
                    if(MeCom_TEC_Tem_TargetObjectTemp(Address, Inst, &Fields, MeSet))
                    {
                        ConsoleIO_SetColor(ConsoleIO_Green);
                        printf("TEC Object Temperature: New Value: %f\n", Fields.Value);
                    }
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 32: 
            {
                MeParLongFields Fields;
                int32_t Inst    = ConsoleIO_IntInput("Please Enter Instance", 1, 255, 1);
                ConsoleIO_SetColor(ConsoleIO_Red);
                if(MeCom_TEC_Ope_OutputStageEnable(Address, Inst, &Fields, MeGetLimits))
                {
                    ConsoleIO_SetColor(ConsoleIO_Reset);
                    Fields.Value = ConsoleIO_IntInput("Please Enter Output Stage Enable Status", Fields.Min, Fields.Max, 0);
                    ConsoleIO_SetColor(ConsoleIO_Red);
                    if(MeCom_TEC_Ope_OutputStageEnable(Address, Inst, &Fields, MeSet))
                    {
                        ConsoleIO_SetColor(ConsoleIO_Green);
                        printf("TEC Output Enable Status: New Value: %d\n", Fields.Value);
                    }
                }
            }
            break; //--------------------------------------------------------------------------------------
            case 50: 
            {
                TestAllCommonGetFunctions(Address);
            }
            break; //--------------------------------------------------------------------------------------
            case 51: 
            {
                TestAllLDDGetFunctions(Address);
            }
            break; //--------------------------------------------------------------------------------------
            case 52: 
            {
                TestAllTECGetFunctions(Address);
            }
            break; //--------------------------------------------------------------------------------------
            default:
                ConsoleIO_SetColor(ConsoleIO_Red);
                printf("Demo Function not available\n");
            break; //--------------------------------------------------------------------------------------
        }
        ConsoleIO_SetColor(ConsoleIO_Reset);
    }
    return 0;
}


int DemoFunc()
{
if(ConsoleIO_YesNo("Do you want to open a Comport? (press enter for default)", 1))
    {
        ComPort_Open(ConsoleIO_IntInput("Please enter ComPort number", 0, 50, 0), 
            ConsoleIO_IntInput("Please enter ComPort Speed", 4800, 1000000, 57600));
    }
    else
    {
        printf("No Comport has been opened, this will result in some communications timeouts.\n");
        printf("It is even though possible to see the resulting Host communication\nin the Communication Log file\n\n");
    }
    int32_t Address = ConsoleIO_IntInput("Please Enter the Device Address", 0, 255, 0);
    
    int8_t Buf[25];
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_GetIdentString(Address, Buf)) 
    {
               ConsoleIO_SetColor(ConsoleIO_Green);
               printf("Ident String: %s\n", Buf);
    }
    return 0;
}

/*==============================================================================*/
/** @brief      Menu selection function
 *
*/
static int32_t MenuSelection(void)
{
    int32_t option;
    ConsoleIO_SetColor(ConsoleIO_Cyan);
    printf("\n\n*********************** MAIN MENU ***********************\n");
    printf(" 1-Get Firmware Identification String (?IF Query)\n");
    printf(" 2-Reset Device (RS Set)\n");
    printf(" 3-Get int Parameter Value (?VR Query)\n");
    printf(" 4-Get float Parameter Value (?VR Query)\n");
    printf(" 5-Set int Parameter Value (?VS Query)\n");
    printf(" 6-Set float Parameter Value (?VS Query)\n");
    printf("10-Get Device Type\n");
    printf("11-Get Serial Number\n");
    printf("20-Get LDD Laser Diode Current\n");
    printf("21-Set LDD Current CW\n");
    printf("22-Set LDD Enable Input Source\n");
    printf("30-Get TEC Object Temperature\n");
    printf("31-Set TEC Target Object Temperature\n");
    printf("32-Set TEC Output Stage Enable Status\n");
    printf("50-Run all Common Get Functions (Test function)\n");
    printf("51-Run all LDD Get Functions (Test function)\n");
    printf("52-Run all TEC Get Functions (Test function)\n");
    printf("Exit with CTRL-C\n");
    ConsoleIO_SetColor(ConsoleIO_Reset);
    option = ConsoleIO_IntInput("\nPlease Select the Demo Function", 1, 100, 1);
    ConsoleIO_Clear();
    return option;
}


/*==============================================================================*/
/** @brief      Test function for all Common Get Functions
 *
*/
static void TestAllCommonGetFunctions(uint8_t Address)
{
    MeParLongFields  lFields;

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
}
/*==============================================================================*/
/** @brief      Test function for all LDD Get Functions
 *
*/
static void TestAllLDDGetFunctions(uint8_t Address)
{
    MeParLongFields  lFields;
    MeParFloatFields fFields;

    //Tab Monitor
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_DeviceType(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_DeviceType: %d\n", lFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_SerialNumber(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_SerialNumber: %d\n", lFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_HardwareVersion(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_HardwareVersion: %d\n", lFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_FirmwareVersion(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_FirmwareVersion: %d\n", lFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_FirmwareBuild(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_FirmwareBuild: %d\n", lFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_FPGAVersion(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_FPGAVersion: %d\n", lFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_LaserDiodeCurrent(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_LaserDiodeCurrent: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_LaserDiodeVoltage(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_LaserDiodeVoltage: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_LaserDiodeTemperature(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_LaserDiodeTemperature: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_PhotoDiodeCurrent(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_PhotoDiodeCurrent: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_LaserPower(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_LaserPower: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_LaserDiodecurrentCW(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_LaserDiodecurrentCW: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_LaserDiodeCurrentActual(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_LaserDiodeCurrentActual: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_LaserDiodeVoltageActual(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_LaserDiodeVoltageActual: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_LaserDiodeCurrentPulse(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_LaserDiodeCurrentPulse: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_LaserDiodeVoltagePulse(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_LaserDiodeVoltagePulse: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_DriverInputVoltage(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_DriverInputVoltage: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_10VInternalSupply(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_10VInternalSupply: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_3_3VInternalSupply(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_3_3VInternalSupply: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_1_2VInternalSupply(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_1_2VInternalSupply: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_ErrorNumber(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_ErrorNumber: %d\n", lFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_ErrorInstance(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_ErrorInstance: %d\n", lFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_ErrorParameter(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_ErrorParameter: %d\n", lFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_BuckConverter1Current(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_BuckConverter1Current: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_BuckConverter2Current(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_BuckConverter2Current: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_BuckConverter3Current(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_BuckConverter3Current: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_BasePlateTemperature(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_BasePlateTemperature: %f\n", fFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_DriverStatus(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_DriverStatus: %d\n", lFields.Value);
    }
     
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Mon_ParameterSystemFlashStatus(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Mon_ParameterSystemFlashStatus: %d\n", lFields.Value);
    }
   
    //Tab: Operation Control
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_CurrentInputSource(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_CurrentInputSource: %d\n", lFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_CurrentCW(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_CurrentCW: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_CurrentHigh(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_CurrentHigh: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_CurrentLow(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_CurrentLow: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_HighTime(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_HighTime: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_LowTime(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_LowTime: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_RiseTime(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_RiseTime: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_FallTime(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_FallTime: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_Synchronisation(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_Synchronisation: %d\n", lFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_PulseInputSource(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_PulseInputSource: %d\n", lFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_PulseHighTime(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_PulseHighTime: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_PulseLowTime(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_PulseLowTime: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Ope_EnableInputSource(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Ope_EnableInputSource: %d\n", lFields.Value);
    }

    //Tab: LaserPower Control
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_InputSource(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_InputSource: %d\n", lFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_LP_CW(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_LP_CW: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_LP_High(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_LP_High: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_LP_Low(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_LP_Low: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_HighTime(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_HighTime: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_LowTime(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_LowTime: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_RiseTime(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_RiseTime: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_FallTime(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_FallTime: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_Kp(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_Kp: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_Ti(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_Ti: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_Td(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_Td: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_SlopeLimit(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_SlopeLimit: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_CurrentLimiterStartValue(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_CurrentLimiterStartValue: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_CurrentLimiterRamp(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_CurrentLimiterRamp: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Las_LP_SystemScale(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Las_LP_SystemScale: %f\n", fFields.Value);
    }
    //Tab: Settings
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_Kp(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_Kp: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_Ti(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_Ti: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_Td(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_Td: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_AnalogCurrentFactor(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_AnalogCurrentFactor: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_CurrentLimitMax(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_CurrentLimitMax: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_CurrentLimitMin(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_CurrentLimitMin: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_MaxCurrentError(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_MaxCurrentError: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_SlopeLimit(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_SlopeLimit: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_PBCResFunc(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_PBCResFunc: %d\n", lFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_CommunicationWatchdog(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_CommunicationWatchdog: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_DeviceAddress(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_DeviceAddress: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_RS485CH1BaudRate(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_RS485CH1BaudRate: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_RS485CH1ResponseDelay(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_RS485CH1ResponseDelay: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_LaserDiodeTempLowerErrorThreshold(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_LaserDiodeTempLowerErrorThreshold: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_LaserDiodeTempUpperErrorThreshold(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_LaserDiodeTempUpperErrorThreshold: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_NTCLowerPointTemperature(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_NTCLowerPointTemperature: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_NTCLowerPointResistance(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_NTCLowerPointResistance: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_NTCMiddlePointTemperature(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_NTCMiddlePointTemperature: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_NTCMiddlePointResistance(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_NTCMiddlePointResistance: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_NTCUpperPointTemperature(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_NTCUpperPointTemperature: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Set_NTCUpperPointResistance(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Set_NTCUpperPointResistance: %f\n", fFields.Value);
    }

    //Tab: Expert
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_LaserDiodeTempADCCalibrationOffset(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_LaserDiodeTempADCCalibrationOffset: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_LaserDiodeTempADCCalibrationGain(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_LaserDiodeTempADCCalibrationGain: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_LaserPowerMeasurementRs(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_LaserPowerMeasurementRs: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_CurrentMeasurementOffset(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_CurrentMeasurementOffset: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_CurrentMeasurementGain(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_CurrentMeasurementGain: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_LaserPowerMeasurementOffset(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_LaserPowerMeasurementOffset: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_LaserPowerMeasurementGain(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_LaserPowerMeasurementGain: %f\n", fFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_ParallelFunctionType(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_ParallelFunctionType: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_ParallelRS485Ch(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_ParallelRS485Ch: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_ParallelNrOfSlaves(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_ParallelNrOfSlaves: %d\n", lFields.Value);
    }

    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Exp_ParallelSlaveID(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Exp_ParallelSlaveID: %d\n", lFields.Value);
    }
    
    //Not in Service Software displayed Parameters
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Oth_CurrentWave_Current(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Oth_CurrentWave_Current: %f\n", fFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Oth_CurrentWave_Pulse(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Oth_CurrentWave_Pulse: %d\n", lFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Oth_CurrentWave_Enable(Address, &lFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Oth_CurrentWave_Enable: %d\n", lFields.Value);
    }
    
    ConsoleIO_SetColor(ConsoleIO_Red);
    if(MeCom_LDD_Oth_CurrentWave_Light(Address, &fFields, MeGet)) 
    {
        ConsoleIO_SetColor(ConsoleIO_Green);
        printf("MeCom_LDD_Oth_CurrentWave_Light: %f\n", fFields.Value);
    }
}
/*==============================================================================*/
/** @brief      Test function for all TEC Get Functions
 *
*/
static void TestAllTECGetFunctions(uint8_t Address)
{
    MeParLongFields  lFields;
    MeParFloatFields fFields;

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

}

