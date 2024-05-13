package config

import (
	"github.com/spf13/viper"
)

const (
	envPath = ".env"
)

type Configurations struct {
	DBhost        string
	DBuser        string
	DBname        string
	DBport        string
	DBpassword    string
	Secreet       string
	AdminEmail    string
	AdminPassword string
	Port          string
}

// loading all the configurations for running the API
func LoadConfigurationss() (Configurations, error) {

	var config Configurations

	//setting the path of the env
	viper.SetConfigFile(envPath)

	// reading the env file
	err := viper.ReadInConfig()
	if err != nil {
		return Configurations{}, err
	}

	// getting the values from env
	config.Secreet = viper.GetString("jwt_secret")
	config.Port = viper.GetString("port")
	config.AdminEmail = viper.GetString("admin_email")
	config.AdminPassword = viper.GetString("admin_password")
	config.DBhost = viper.GetString("dbhost")
	config.DBport = viper.GetString("dbport")
	config.DBname = viper.GetString("dbname")
	config.DBpassword = viper.GetString("dbpassword")
	config.DBuser = viper.GetString("dbuser")

	return config, nil
}
