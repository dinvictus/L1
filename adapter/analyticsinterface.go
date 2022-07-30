package adapter

type Analytics interface { // Интерфейс, который нужен клиенту
	AnalyzeData(xmlData string) string
}
