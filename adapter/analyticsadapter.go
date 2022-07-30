package adapter

import (
	"encoding/json"
	"encoding/xml"
)

type dataFormat struct { // Структура для преобразования xml в json
	ProductList []struct {
		Sku      string `xml:"sku" json:"sku"`
		Quantity int    `xml:"quantity" json:"quantity"`
	} `xml:"Product" json:"products"`
}

type XmlJsonAdapter struct { // Адаптер
	aLibrary    AnalyticsLibrary // Закрытая библиотека для анализа данных
	Coefficient float32          // Какой-то коофициент
}

func (adapter *XmlJsonAdapter) AnalyzeData(xmlData string) string { // Метод для анализа данных
	adapter.aLibrary = AnalyticsLibrary{adapter.Coefficient} // Создаём экземпляр библиотеки с коофициентом
	dataformat := dataFormat{}
	errXml := xml.Unmarshal([]byte(xmlData), &dataformat) // Данные xml  в структуру
	if errXml != nil {
		return errXml.Error()
	}
	jsonData, errJson := json.Marshal(dataformat) // Данные из структуры в json
	if errJson != nil {
		return errJson.Error()
	}
	return adapter.aLibrary.AnalyzeDataJson(string(jsonData)) // Анализ данных
}
