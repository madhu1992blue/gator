package config
import (
	"os"
	"bytes"
	"log"
	"encoding/json"
)


func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Couldn't determine the users home dir while loading config: %v", err)
		return  "",err
	}
	configPath := homeDir + "/.gatorconfig.json"
	return configPath, nil

}


type Config struct {
	DBUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}


func Read() (*Config, error) {
	
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	var config Config
	if err := decoder.Decode(&config); err!=nil {
		return nil, err
	}
	return &config, nil
}


func (c *Config) SetUser(newUsername string) error {	
	c.CurrentUserName = newUsername
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}
	
	configBytes, err := json.Marshal(c)
	if err != nil {
		log.Printf("Couldn't convert config to JSON: %v", err)
		return err
	}
	err = os.WriteFile(configPath, configBytes, 0600)
	if err != nil {
		log.Printf("Failed to save the username in the config: %v", err) 	
		return err
	}
	return nil

}
