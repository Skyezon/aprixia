package database

type DataSource interface{
    InsertData(real_url string, short_url string, create_at string, redirect_count int) error
    GetStats(short_url string) (redirect_count int, create_at string)
    GetRealUrl(short_url string) 
    GetData(short_url string)
}

type UrlData struct {
    ShortUrl string
    RealUrl string
    CreateAt string
    RedirectCount int
}
