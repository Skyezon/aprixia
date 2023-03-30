package service

import (
    "aprixa/datasource"
    "math/rand"
    "time"
    "unicode"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
//will return 6 character alphanumeric

func generate6Alphanumeric() string {
    rand.Seed(time.Now().UnixNano())
    alphaFlag := false
    numericFlag := false
    b := make([]byte, 6)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
        if numericFlag == false && unicode.IsDigit(rune( b[i])) {
            numericFlag = true
        }
        if alphaFlag == false && unicode.IsLetter(rune(b[i])){
            alphaFlag = true
        }
    }
    if alphaFlag && numericFlag {
        return string(b)
    }else{
        return generate6Alphanumeric()
    }
}

func UrlShorterner(longUrl string) (string,error) {
    shortUrl := generate6Alphanumeric()
    var db datasource.Database
    db,err := datasource.NewDB()
    if err != nil {
        return "",err
    }
   defer db.Close() 
    //TODO : probably going to need redis for better performance
    err = db.InsertData(datasource.UrlData{ 
        RealUrl : longUrl,
        ShortUrl: shortUrl,
        CreateAt: time.Now().Format("2006-01-02 15:04:05"),
        RedirectCount : 0})
        if err != nil {
            //if err == unique constraint then generate again
            return "",err
        }
        return shortUrl, nil
    }

    func GetUrlData(shortUrl string) (datasource.UrlData,error){
        var db datasource.Database
        db,err := datasource.NewDB()
        if err != nil {
            return datasource.UrlData{},err
        }
       defer db.Close() 
        urlData, err :=  db.GetData(shortUrl)
        if err != nil {
            return datasource.UrlData{},err
        }
        return urlData,nil
    }

    func GetRedirect(shortUrl string)(string,error){
        urlData, err := GetUrlData(shortUrl)
        if err != nil {
            return "",err
        }
        var db datasource.Database
        db,err = datasource.NewDB()
        if err != nil {
            return "",err
        }
       defer db.Close() 
        err = db.IncrementRedirectCount(shortUrl)

        if err != nil {
            return "",err
        }
        return urlData.RealUrl,nil 
    }
