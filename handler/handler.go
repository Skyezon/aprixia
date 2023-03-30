package handler

import (
	"aprixa/service"
	"aprixa/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func GenerateAliasHandler(w http.ResponseWriter, r *http.Request) {
    if !utils.ValidateMethodPost(r,w){
        return
    }

    var data GenerateAliasRequest
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
         fmt.Printf("unmarshall error : %v",err)
         http.Error(w,err.Error(),http.StatusInternalServerError)
        return
    }
    shorterUrl, err := service.UrlShorterner(data.LongUrl)
    if err != nil {
        fmt.Printf("fail to url shorter : %v",err)
         http.Error(w,err.Error(),http.StatusInternalServerError)
        return
    }
    res, err := json.Marshal(GenerateAliasResponse{ShortUrl: shorterUrl}) 
    if err != nil {
        fmt.Printf("fail to Marshal : %v",err)
         http.Error(w,err.Error(),http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func RedirectAliasHandler(w http.ResponseWriter, r *http.Request) {
    if !utils.ValidateMethodPost(r,w){
        return
    }

    var data RedirectAliasRequest
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        fmt.Printf("unmarshall error : %v",err)
         http.Error(w,err.Error(),http.StatusInternalServerError)
        return
    }
    realUrl, err := service.GetRedirect(data.ShortUrl)
    if err != nil {
        if err == sql.ErrNoRows{
            http.Error(w,fmt.Sprintf("Short link not found"),http.StatusNotFound)
            return
        }
         http.Error(w,err.Error(),http.StatusInternalServerError)
        return
    }
    res, err := json.Marshal(RedirectAliasResponse{LongUrl: realUrl})
    if err != nil {
         http.Error(w,err.Error(),http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetStatsHandler(w http.ResponseWriter, r * http.Request) {
    if !utils.ValidateMethodGet(r,w){
        return
    }

   var data GetStatsRequest
   err := json.NewDecoder(r.Body).Decode(&data)
   if err != nil {
        fmt.Printf("unmarshall error: %v",err)
         http.Error(w,err.Error(),http.StatusInternalServerError)
        return
   }
    urlData, err := service.GetUrlData(data.ShortUrl)
    if err != nil {
        if err == sql.ErrNoRows{
            http.Error(w,fmt.Sprintf("Short link not found"),http.StatusNotFound)
            return
        } 
        http.Error(w,err.Error(),http.StatusInternalServerError)
        return
    }
    res, err := json.Marshal(GetStatsResponse{RedirectCount: urlData.RedirectCount, CreateAt: urlData.CreateAt})
    if err != nil {
         http.Error(w,err.Error(),http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}


