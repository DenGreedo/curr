package converter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type apiResponse struct {
	Result string             `json:"result"`
	Rates  map[string]float64 `json:"rates"`
}

func FetchRates(base string) (map[string]float64, error) {
	base = strings.ToUpper(base)
	url := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", base)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API вернул статус %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	var data apiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON: %w", err)
	}

	if data.Result != "success" {
		return nil, fmt.Errorf("API сообщил об ошибке: %s", data.Result)
	}

	return data.Rates, nil
}

func Convert(amount float64, from, to string, rates map[string]float64) (float64, error) {
	from = strings.ToUpper(from)
	to = strings.ToUpper(to)

	if amount <= 0 {
		return 0, fmt.Errorf("сумма должна быть положительной")
	}

	rate, ok := rates[to]
	if !ok {
		return 0, fmt.Errorf("валюта %s не найдена в списке курсов", to)
	}

	return amount * rate, nil
}
