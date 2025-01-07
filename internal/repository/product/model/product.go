package model

type Product struct{
	ID int
	Article string `xml:"Артикул"`
	Name string `xml:"Наименование"`
	CategoryId int`xml:"Группы>Ид"`
	Price float32`xml:"Цены>Цена>ЦенаЗаЕдиницу"`
	Stock int`xml:"Количество"`
}
