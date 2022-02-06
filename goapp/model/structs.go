package model

import "time"

//create request_post_api type
type RequestFE struct {
	TimeRange string `json:"timeRange"`
}
type FundArr struct {
	Data []FundData `json:"data"`
}
type FundData struct {
	Mstar_id           string  `json:"mstar_id"`
	Thailand_fund_code string  `json:"thailand_fund_code"`
	Nav_return         float64 `json:"nav_return"`
	Nav                float64 `json:"nav"`
	Nav_date           string  `json:"nav_date"`
	Avg_return         float64 `json:"avg_return"`
	Go_date            time.Time
}

type ResFundArr struct {
	Data []ResFundData `json:"data"`
}
type ResFundData struct {
	Name        string  `json:"name"`
	RankOfFund  int     `json:"rank_of_fund"`
	UpdatedDate string  `json:"updated_date"`
	Performance float64 `json:"performance"`
	Price       float64 `json:"price"`
}
