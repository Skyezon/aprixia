package handler

//request

type GenerateAliasRequest struct {
    long_url string
} 

type RedirectAliasRequest struct {
    short_url string
}

type GetStatsRequest struct {
    short_url string
}

//response

type GenerateAliasResponse struct {
    short_url string
}

type RedirectAliasResponse struct {
    long_url string
}

type GetStatsResponse struct {
    redirect_count int
    create_at string
}
