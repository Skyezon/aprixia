package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GenerateAliasHandler(w http.ResponseWriter, r *http.Request) {
    var data GenerateAliasRequest
    err := json.NewDecoder(r.Body).Decode(data)
    if err != nil {
         fmt.Errorf("unmarshall error : %v",err)
         return
    }

}

func RedirectAliasHandler(w http.ResponseWriter, r *http.Request) {
    var data RedirectAliasRequest
    err := json.NewDecoder(r.Body).Decode(data)
    if err != nil {
        fmt.Errorf("unmarshall error : %v",err)
        return
    }
}

func GetStatsHandler(w http.ResponseWriter, r * http.Request) {
   var data GetStatsRequest
   err := json.NewDecoder(r.Body).Decode(data)
   if err != nil {
        fmt.Errorf("unmarshall error: %v",err)
        return
   }
}


