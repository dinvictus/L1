package main

import (
	a "adapter"
	"fmt"
)

func Work21() {
	dataProvider := a.DataProvider{}                                    // Источник данных в формате xml
	dataXML := dataProvider.GetDataXML()                                // Данные в формате xml
	var analytics a.Analytics = &a.XmlJsonAdapter{Coefficient: 1.65423} // Объявляем адаптер через интерфейс
	analyticsData := analytics.AnalyzeData(dataXML)                     // Анализируем данные
	fmt.Println(analyticsData)
}
