package handler

import (
	"fmt"
	"net/http"
	"os"

	"weather_info_search/funcs"

	"github.com/gin-gonic/gin"
)

// Run Web Server
func Serve() {
	r := gin.Default()

	// templates配下のHTMLをロード
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handleHome)
	r.GET("/weather", handleGetWeatherInfo)

	r.Run()
}

// 初期表示
func handleHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// 天気情報取得
func handleGetWeatherInfo(ctx *gin.Context) {
	// idクエリパラメーターを取得
	cityID := ctx.Query("id")

	// OpenWeatherMap APIに対するリクエストURLを構築
	apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%s&appid=%s", cityID, os.Getenv("OPENWEATHERMAP_API_KEY"))

	// APIからデータを取得
	weatherData, err := funcs.GetWeatherData(apiURL)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// HTMLテンプレートに渡すデータを生成
	data := gin.H{
		"CityName":    funcs.GetCityNameByID(cityID),
		"Temp":        fmt.Sprintf("%.2f", weatherData.Main.Temp-273.16),
		"Weather":     weatherData.Weather[0].Main,
		"Description": weatherData.Weather[0].Description,
	}

	// HTMLテンプレートを読み込み、データを埋め込んでHTMLを生成
	ctx.HTML(http.StatusOK, "result.html", data)
}
