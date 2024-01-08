package funcs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"weather_info_search/model"
)

func GetWeatherData(apiURL string) (model.Weather, error) {
	var weatherData model.Weather

	// APIからデータを取得
	resp, err := http.Get(apiURL)
	if err != nil {
		return weatherData, fmt.Errorf("error making request to OpenWeatherMap API: %v", err)
	}
	defer resp.Body.Close()

	// レスポンスボディを読み込む
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weatherData, fmt.Errorf("error reading response body: %v", err)
	}

	// JSONデコード
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return weatherData, fmt.Errorf("error decoding JSON: %v", err)
	}

	return weatherData, nil
}

func GetCityNameByID(cityID string) string {
	switch cityID {
	case "2130037":
		return "北海道"
	case "1850144":
		return "東京"
	case "1860291":
		return "神奈川"
	case "1856057":
		return "名古屋"
	case "1853909":
		return "大阪"
	case "1863967":
		return "福岡"
	case "1894616":
		return "沖縄"
	default:
		return "Unknown"
	}
}
