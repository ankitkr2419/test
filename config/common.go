package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Common struct {
	RoomTemperature   float64 `json:"room_temperature" validate:"required,lte=30,gte=20"`
	ReceiverEmail     string  `json:"receiver_email"`
	ReceiverName      string  `json:"receiver_name"`
	SerialNumber      string  `json:"serial_number"`
	MachineVersion    string  `json:"machine_version"`
	SoftwareVersion   string  `json:"software_version"`
	ContactNumber     string  `json:"contact_number"`
	ManufacturingYear string  `json:"manufacturing_year"`
}

func SetCommonConfigValues(co Common) (err error) {

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,
		fmt.Sprintf("room_temp: %.2f", GetRoomTemp()),
		fmt.Sprintf("receiver_email: %s", GetReceiverEmail()),
		fmt.Sprintf("receiver_name: %s", GetReceiverName()),
		fmt.Sprintf("serial_number: %s", GetSerialNumber()),
		fmt.Sprintf("machine_version: %s", GetMachineVersion()),
		fmt.Sprintf("software_version: %s", GetSoftwareVersion()),
		fmt.Sprintf("contact_number: %s", GetContactNumber()),
		fmt.Sprintf("manufacturing_year: %s", GetManufacturingYear()),
	)
	newString = append(newString,
		fmt.Sprintf("room_temp: %.2f", co.RoomTemperature),
		fmt.Sprintf("receiver_email: %s", co.ReceiverEmail),
		fmt.Sprintf("receiver_name: %s", co.ReceiverName),
		fmt.Sprintf("serial_number: %s", co.SerialNumber),
		fmt.Sprintf("machine_version: %s", co.MachineVersion),
		fmt.Sprintf("software_version: %s", co.SoftwareVersion),
		fmt.Sprintf("contact_number: %s", co.ContactNumber),
		fmt.Sprintf("manufacturing_year: %s", co.ManufacturingYear),
	)

	err = UpdateConfig(configPath)
	if err != nil {
		return
	}

	SetRoomTemp(co.RoomTemperature)
	SetReceiverEmail(co.ReceiverEmail)
	SetReceiverName(co.ReceiverName)
	SetSerialNumber(co.SerialNumber)
	SetMachineVersion(co.MachineVersion)
	SetSoftwareVersion(co.SoftwareVersion)
	SetContactNumber(co.ContactNumber)
	SetManufacturingYear(co.ManufacturingYear)
	return
}

func AppName() string {
	if appName == "" {
		appName = ReadEnvString("APP_NAME")
	}
	return appName
}

func AppPort() int {
	if appPort == 0 {
		appPort = ReadEnvInt("APP_PORT")
	}
	return appPort
}

func GetSendGridAPIKey() string {
	key := "SENDGRID_API_KEY"
	checkIfSet(key)
	return viper.GetString(key)
}

func GetSecretKey() string {
	key := "SECRET_KEY"
	checkIfSet(key)
	return viper.GetString(key)
}

func GetRoomTemp() float64 {
	return ReadEnvFloat("room_temp")
}

func GetReceiverName() string {
	key := "receiver_name"
	checkIfSet(key)
	return viper.GetString(key)
}

func GetReceiverEmail() string {
	key := "receiver_email"
	checkIfSet(key)
	return viper.GetString(key)
}

func GetSerialNumber() string {
	key := "serial_number"
	checkIfSet(key)
	return viper.GetString(key)
}

func GetMachineVersion() string {
	key := "machine_version"
	checkIfSet(key)
	return viper.GetString(key)
}

func GetSoftwareVersion() string {
	key := "software_version"
	checkIfSet(key)
	return viper.GetString(key)
}

func GetContactNumber() string {
	key := "contact_number"
	checkIfSet(key)
	return viper.GetString(key)
}

func GetManufacturingYear() string {
	key := "manufacturing_year"
	checkIfSet(key)
	return viper.GetString(key)
}

func SetRoomTemp(rT float64) {
	viper.Set("room_temp", rT)
}

func SetSecretKey(value string) {
	key := "SECRET_KEY"
	viper.Set(key, value)
}

func SetReceiverEmail(rE string) {
	viper.Set("receiver_email", rE)
}

func SetReceiverName(rN string) {
	viper.Set("receiver_name", rN)
}

func SetSerialNumber(sN string) {
	viper.Set("serial_number", sN)
}

func SetMachineVersion(mV string) {
	viper.Set("machine_version", mV)
}

func SetSoftwareVersion(sV string) {
	viper.Set("software_version", sV)
}

func SetContactNumber(cN string) {
	viper.Set("contact_number", cN)
}

func SetManufacturingYear(mY string) {
	viper.Set("manufacturing_year", mY)
}

func GetCommonConfigValues() Common {
	return Common{
		RoomTemperature:   GetRoomTemp(),
		ReceiverEmail:     GetReceiverEmail(),
		ReceiverName:      GetReceiverName(),
		SerialNumber:      GetSerialNumber(),
		MachineVersion:    GetMachineVersion(),
		SoftwareVersion:   GetSoftwareVersion(),
		ContactNumber:     GetContactNumber(),
		ManufacturingYear: GetManufacturingYear(),
	}
}
