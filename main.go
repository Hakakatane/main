package main

import (
	"fmt"
	"example.com/building"
	"example.com/structure"
	"example.com/hotel"
)

func MakeSlogan(c hotel.Company) {
	c.Slogan()
}

func main() {
	b := building.NewBuilding(2020, "USA", 1000000.0)
	s := structure.NewStructure(b, 5000.0, 10, "47 street")
	h := hotel.NewHotel(s, "Plaza", 5, 200.0)
	MakeSlogan(&h)

	fmt.Println("Информация об отеле:")
	fmt.Println("  Название:", h.GetName())
	fmt.Println("  Звезды:", h.GetStars())
	fmt.Println("  Цена за ночь:", h.GetPricePerNight())
	fmt.Println("  Доступность:", h.GetAvailability())
	fmt.Println("  Адрес:", h.Structure.GetAddress())
	fmt.Println("  Площадь:", h.Structure.GetArea())
	fmt.Println("  Этажей:", h.Structure.GetFloors())
	fmt.Println("  Год постройки:", h.Structure.Building.GetYearBuilt())
	fmt.Println("  Страна:", h.Structure.Building.GetCountry())
	fmt.Println("  Стоимость постройки:", h.Structure.Building.GetCost())
	fmt.Println("  Статус постройки:", h.Structure.Building.GetStatus())
	fmt.Println("  Состояние здания:", h.Structure.GetCondition())

	// Изменяем состояние отеля
	h.NoRooms()
	fmt.Println("  Доступность:", h.GetAvailability())

	// Изменяем год постройки
	h.Structure.Building.YearBuilt = 2021
	fmt.Println("  Год постройки:", h.Structure.Building.GetYearBuilt())

	// Устанавливаем свой статус
	h.SetAvailability("Подготовка к открытию")
	fmt.Println("  Доступность:", h.GetAvailability())

	b2 := building.NewBuilding(1990, "France", 5000000.0)
	s2 := structure.NewStructure(b2, 10000.0, 20, "13 Charle st")
	h2 := hotel.NewHotel(s2, "Fiorano", 4, 350.0)

	fmt.Println("\nИнформация о втором отеле:")
	fmt.Println("  Название:", h2.GetName())
	fmt.Println("  Доступность:", h2.GetAvailability())
	fmt.Println("  Год постройки:", h2.Structure.Building.GetYearBuilt())

	h2.Close()
}
