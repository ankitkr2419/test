package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Common struct {
	RoomTemperature float64  `json:"room_temperature" validate:"required,lte=30,gte=20"`
	ReceiverEmail   string `json:"receiver_email"`
	ReceiverName    string `json:"receiver_name"`
}

func SetCommonConfigValues(co Common) (err error) {

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,
		fmt.Sprintf("room_temperature: %.2f", GetRoomTemp()),
		fmt.Sprintf("receiver_email: %s", GetReceiverEmail()),
		fmt.Sprintf("receiver_name: %s", GetReceiverName()),
	)
	newString = append(newString,
		fmt.Sprintf("room_temperature: %.2f", co.RoomTemperature),
		fmt.Sprintf("receiver_email: %s", co.ReceiverEmail),
		fmt.Sprintf("receiver_name: %s", co.ReceiverName),
	)

	err = UpdateConfig(configPath)
	if err != nil {
		return
	}

	SetRoomTemp(co.RoomTemperature)
	SetReceiverEmail(co.ReceiverEmail)
	SetReceiverName(co.ReceiverName)
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

func SetRoomTemp(rT float64) {
	viper.Set("room_temp", rT)
}

func SetSecretKey(key string) {
	key = "SECRET_KEY"
	viper.Set(key, "123456qwerty")
}

func SetReceiverEmail(rE string) {
	viper.Set("receiver_email", rE)
}

func SetReceiverName(rN string) {
	viper.Set("receiver_name", rN)
}

func GetCommonConfigValues() Common {
	return Common{
		RoomTemperature: GetRoomTemp(),
		ReceiverEmail: GetReceiverEmail(),
		ReceiverName: GetReceiverName(),
	}
}