package main

import "fmt"

func Work15() {
	fmt.Println(ErrorFunc())
	fmt.Println(TrueFunc())
}

func createHugeString() string {
	return "руаываоцжыоваоДдфта гр34шп шфукыи ыплкркы п9й4 рцкыгмр пгфцр ы9 8пщрцщкфгы рщгцрык мщгр4цщыр м24щрцымгггфциа8р м2щрцргщрпщфыдрвщрыфдпр 2йщ4цмт2 щгфцрщ РЩПЦИШПАДШ Гр ажарфугщ жА ЦПУЦЩАЖГ2РЦУЩЫПРИ 2ШГУЦруаываоцжыоваоДдфта гр34шп шфукыи ыплкркы п9й4 рцкыгмр пгфцр ы9 8пщрцщкфгы рщгцрык мщгр4цщыр м24щрцымгггфциа8р м2щрцргщрпщфыдрвщрыфдпр 2йщ4цмт2 щгфцрщ РЩПЦИШПАДШ Гр ажарфугщ жА ЦПУЦЩАЖГ2РЦУЩЫПРИ 2ШГУЦруаываоцжыоваоДдфта гр34шп шфукыи ыплкркы п9й4 рцкыгмр пгфцр ы9 8пщрцщкфгы рщгцрык мщгр4цщыр м24щрцымгггфциа8р м2щрцргщрпщфыдрвщрыфдпр 2йщ4цмт2 щгфцрщ РЩПЦИШПАДШ Гр ажарфугщ жА ЦПУЦЩАЖГ2РЦУЩЫПРИ 2ШГУЦруаываоцжыоваоДдфта гр34шп шфукыи ыплкркы п9й4 рцкыгмр пгфцр ы9 8пщрцщкфгы рщгцрык мщгр4цщыр м24щрцымгггфциа8р м2щрцргщрпщфыдрвщрыфдпр 2йщ4цмт2 щгфцрщ РЩПЦИШПАДШ Гр ажарфугщ жА ЦПУЦЩАЖГ2РЦУЩЫПРИ 2ШГУЦруаываоцжыоваоДдфта гр34шп шфукыи ыплкркы п9й4 рцкыгмр пгфцр ы9 8пщрцщкфгы рщгцрык мщгр4цщыр м24щрцымгггфциа8р м2щрцргщрпщфыдрвщрыфдпр 2йщ4цмт2 щгфцрщ РЩПЦИШПАДШ Гр ажарфугщ жА ЦПУЦЩАЖГ2РЦУЩЫПРИ 2ШГУЦЫВАЫВААЫВАЫАЫВАываываыаа"
}

func ErrorFunc() string {
	var justString string
	v := createHugeString()
	justString = v[:100] // Неверно образовывать срез строки, так как строка - это массив байтов, а символ может состоять из нескольких
	// Также в памяти будет храниться вся строка, хотя требуется только 100 её первых символов
	return justString
}

func TrueFunc() string {
	var justString string
	v := createHugeString()
	runeHugeString := []rune(v)[:100]                         // Преобразуем строку в срез рун и берём срез уже от него, каждый элемент будет хранить ровно один символ
	slicedRuneHugeString := make([]rune, len(runeHugeString)) // Создаём новый срез для копирования туда полученного, чтобы не хранить в памяти всю строку
	copy(slicedRuneHugeString, runeHugeString)                // Копируем
	justString = string(slicedRuneHugeString)                 // Преобразуем срез рун в строку
	return justString
}
