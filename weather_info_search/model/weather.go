package model

// 天気の詳細情報
type WeatherInfo struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// 気温の詳細情報
type TempInfo struct {
	Temp float64 `json:"temp"`
}

// 天気情報
type Weather struct {
	Name    string        `json:"name"`
	Weather []WeatherInfo `json:"weather"`
	Main    TempInfo      `json:"main"`
}
