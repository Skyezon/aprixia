package datasource

import (
	"aprixa/utils"
	"database/sql"
	"fmt"
    _ "github.com/lib/pq"
)

type Database interface{
    ConnectDb() error
    Close() error
    InsertData(urlData UrlData) error
    GetData(short_url string) (urlData UrlData, err error)
    Ping() error
    IncrementRedirectCount(short_url string) error
}

type PostgresDB struct{
    Db *sql.DB
}

//TODO : probably will need model for this
type UrlData struct {
    ShortUrl string
    RealUrl string
    CreateAt string
    RedirectCount int
}

func NewDB() (Database, error){
    conf, err := utils.GetConfig()
    if err != nil {
        return nil,err
    }
    switch conf.DB.DB_TYPE {
        case "postgres" :{
            var db Database
            db = &PostgresDB{}
            db.ConnectDb()
            return db,nil
        }
        //insert here to add more db type &  adjust env config db type
        default :{
            return nil , fmt.Errorf("invalid db type")
        }
    }
}

func (db * PostgresDB)Ping() error{
     if db.Db != nil {
        return db.Db.Ping()
    }
   return nil 
}

func (db *PostgresDB)ConnectDb() (err error){
    conf,err := utils.GetConfig()
    psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.DB.DB_HOST, conf.DB.DB_PORT, conf.DB.DB_USER, conf.DB.DB_PASSWORD, conf.DB.DB_NAME)
         
        // open database
    db.Db, err = sql.Open("postgres", psqlconn)
    if err != nil {
        return  err
    } 
    return  nil 
}

func (db *PostgresDB)Close() (err error){
    if db.Db != nil {
        db.Db.Close()
    }
    return nil
}

func (db *PostgresDB)InsertData(urlData UrlData) error{
    query := fmt.Sprint(`INSERT INTO urldata VALUES ($1,$2,$3,$4);`)
    stmt, err := db.Db.Prepare(query)
    if err != nil {
       return  fmt.Errorf("something went wrong when preparing : %v",err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(urlData.RealUrl,urlData.ShortUrl,urlData.CreateAt,urlData.RedirectCount)
    if err != nil {
        return err
    }
   
    return nil
}


// 3 states : success, success empty, err
func (db *PostgresDB)GetData(short_url string) (urlData UrlData,err error) {
    query := fmt.Sprint(`SELECT * FROM urldata where short_url = $1`)
    stmt, err := db.Db.Prepare(query)
    if err != nil {
       return UrlData{}, fmt.Errorf("something went wrong when preparing : %v",err)
    }
    defer stmt.Close()
    row := stmt.QueryRow(short_url)
    if err != nil {
        return UrlData{},err
    }
    if err := row.Scan(&urlData.RealUrl,&urlData.ShortUrl,&urlData.CreateAt,&urlData.RedirectCount); err != nil{

        return UrlData{}, err 
    }
    return urlData, nil
    
}

func (db *PostgresDB) IncrementRedirectCount(short_url string) error{
    urlData, err := db.GetData(short_url)
    if err != nil {
        return err 
    }
    query := fmt.Sprint("UPDATE urldata set redirect_count = $1 WHERE short_url = $2")
    stmt, err := db.Db.Prepare(query)
    if err != nil {
       return  fmt.Errorf("something went wrong when preparing : %v",err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(urlData.RedirectCount + 1,short_url)
    if err != nil {
        return err
    }
    return nil
}
