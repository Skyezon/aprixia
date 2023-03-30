package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
    DB Database `json:"database"`
}
 type  Database struct {
        DB_NAME string `json:"db_name"`
        DB_PASSWORD string `json:"db_password"`
        DB_HOST string `json:"db_host"`
        DB_PORT string `json:"db_port"`
        DB_USER string `json:"db_user"`
        DB_TYPE string `json:"db_type"`
    }

func GetConfig() (Config,error){
    content, err := ioutil.ReadFile("./env.json")
    if err != nil {
        fmt.Printf("fail to read env.json : %v",err)
        return Config{}, err
    }
    
    var config Config
    err = json.Unmarshal(content,&config) 

    if err != nil {
        fmt.Printf("fail to unmarshall : %v",err)
        return Config{}, err
    }
    return config, err
}

func ValidateMethodPost(r *http.Request, w http.ResponseWriter) bool {
    if r.Method != http.MethodPost {
         http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return false
     }
    return true
}

func ValidateMethodGet(r *http.Request, w http.ResponseWriter) bool {
    if r.Method != http.MethodGet {
         http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
         return false
    }
    return true
}

