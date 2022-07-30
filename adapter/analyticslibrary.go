package adapter

import (
	"encoding/json"
	"fmt"
)

type AnalyticsLibrary struct { // Закрытая библиотека для анализа данных
	Coefficient float32
}

func (obj AnalyticsLibrary) AnalyzeDataJson(data string) string { // Метод для анализа данных
	if !json.Valid([]byte(data)) {
		return "Not valid json data"
	}
	fmt.Println(data)
	return fmt.Sprint("Analytics: ", float32(len(data))*obj.Coefficient)
}
