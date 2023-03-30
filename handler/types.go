package handler

//request

type GenerateAliasRequest struct {
    LongUrl string `json:"long_url"`
} 

type RedirectAliasRequest struct {
    ShortUrl string `json:"short_url"`
}

type GetStatsRequest struct {
    ShortUrl string `json:"short_url"`
}

//response

type GenerateAliasResponse struct {
    ShortUrl string `json:"short_url"`
}

type RedirectAliasResponse struct {
    LongUrl string `json:"long_url"`
}

type GetStatsResponse struct {
    RedirectCount int `json:"redirect_count"`
    CreateAt string `json:"create_at"`
}

type ErrorResponse struct {
    Message string `json:"message"`
}
