package main

import (
	"aprixa/datasource"
	"aprixa/handler"
	"aprixa/utils"
	"fmt"
	"log"
	"net/http"
)

func main()  {
    fmt.Println("running")
    if err := healthCheck();err != nil {
        log.Fatal(fmt.Sprintf("health check fail : %v",err))
    }
    fmt.Println("healthCheck completed")
    mux := http.NewServeMux()

    mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w,"Hello world")
    })
    
    mux.HandleFunc("/shorter",handler.GenerateAliasHandler)
    mux.HandleFunc("/redirect",handler.RedirectAliasHandler)
    mux.HandleFunc("/stats",handler.GetStatsHandler)

    fmt.Println("Listening on port :8080")
    http.ListenAndServe(":8080",mux)
}

func healthCheck()error{
    _, err := utils.GetConfig()
   if err != nil {
        return err
   }
//   fmt.Printf("ini liatin : %+v",conf)
    
   db , err := datasource.NewDB()
   if err != nil {
        return err
   }
    defer db.Close()
   err =  db.Ping() 
   if err != nil {
       return err
   }

   return nil

}
