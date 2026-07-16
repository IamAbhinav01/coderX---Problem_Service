package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() error {
	return godotenv.Load()
}

func init(){
	Load()
}

func GetInt(config_name string) int{
	value,ok := os.LookupEnv(config_name)

	if !ok{
		return 0
	}

	intValue,err := strconv.Atoi(value)

	if err!= nil{
		return 0
	}

	return intValue
}

func GetString(config_name string) string{
	value,ok := os.LookupEnv(config_name)

	if !ok{
		return ""
	}
	return value

}