package litres

import (
	"encoding/json"
	"fmt"
	"go-litres/logger"
	"go-litres/models"
	"go-litres/tools"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)


func PutBookToList(id int, wg *sync.WaitGroup, books *[]models.Book) {
	defer wg.Done()
	url := fmt.Sprintf("https://api.litres.ru/foundation/api/arts/%d", id)
	jsonResponse := getJsonResponse(url)
	book := formatJsonResponseToBook(jsonResponse)
	*books = append(*books, book)

}

func getJsonResponse(url string) (result models.JsonResponse){
	client := http.Client{Timeout: 30 * time.Second}
	
	resp, err := client.Get(url)
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	tools.CheckErr(err)
	logger.Log.Printf("Status: %d - Парсим страницу %s", resp.StatusCode, url)
	return

}

func formatJsonResponseToBook(json models.JsonResponse) (book models.Book) {
	var isAudio bool
	switch json.Payload.Data.IsAudio {
	case 1: isAudio = true
	}

	year, _ := strconv.Atoi(strings.Split(json.Payload.Data.Year, "-")[0])

	return models.Book{
		Id: json.Payload.Data.Id,
		Name: json.Payload.Data.Name,
		Description: json.Payload.Data.Description,
		Image: "https://cv9.litres.ru" + json.Payload.Data.Image,
		Url: "https://cv9.litres.ru" +  json.Payload.Data.Url,
		OldPrice: json.Payload.Data.Price.OldPrice,
		NewPrice: json.Payload.Data.Price.NewPrice,
		Currency: json.Payload.Data.Price.Currency,
		MinAge: json.Payload.Data.MinAge,
		Language: json.Payload.Data.Language,
		Rating: json.Payload.Data.Rating.Code,
		Pages: json.Payload.Data.Page.Count,
		Year: year,
		IsAudio: isAudio,
	}
}