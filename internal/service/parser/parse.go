package parser

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"

	categoryModel "github.com/acronix0/XML-Parser-Golang/internal/repository/category/model"
	productModel "github.com/acronix0/XML-Parser-Golang/internal/repository/product/model"
)

func (s *ParserService) Parse(ctx context.Context, filepath string) error{
	file, err := os.Open(s.firstFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	secondFileData, err := s.parseSecondFile(ctx)
	if err != nil {
		return err
	}

	categories, products, err := s.parseFirstFile(ctx,secondFileData)
	if err != nil{
		return err
	}
	for pr:=range products{
		fmt.Println(pr)
	}
	for ct:=range categories{
		fmt.Println(ct)
	}
	return nil
}

func (s *ParserService)parseFirstFile(ctx context.Context, productsChan <-chan map[string]productModel.Product ) ([]categoryModel.Category,map[string]productModel.Product,error){
	var (
		category categoryModel.Category
		categories []categoryModel.Category
		product productModel.Product
		products map[string]productModel.Product
	)
	file, err := os.Open(s.firstFilePath)
	if err != nil {
		return nil,nil,err
	}
	defer file.Close()
	decoder := xml.NewDecoder(file)
	for t, _:=decoder.Token(); t != nil; t, _=decoder.Token(){
		switch se:=t.(type){
		case xml.StartElement:
			if se.Name.Local == "Группа" {
				err = decoder.DecodeElement(&category, &se)
				if err == nil {
					categories = append(categories, category)
				}else{
					s.log.Info("error while parse category element from xml")
				}
			}	

			if se.Name.Local == "Товар" {
				err = decoder.DecodeElement(&product, &se)
				if err == nil {
					if products == nil {
						products = <-productsChan
					}
					tempProduct := products[product.Article] 
					product.Price = tempProduct.Price
					product.Stock = tempProduct.Stock
					products[product.Article] = product
				}
			}
		}
			
	}
	for pr:=range products{
		fmt.Println(pr)
	}
	return categories, products, nil
}
func (s *ParserService)parseSecondFile(ctx context.Context) (<-chan map[string]productModel.Product, error){

	c := make(chan map[string]productModel.Product, 1)
	products := make(map[string]productModel.Product)
	go func() {
			defer close(c)
			file, err := os.Open(s.secondFilePath)
			if err != nil {
				return 
			}
			decoder := xml.NewDecoder(file)
			var product productModel.Product
			for t, _ := decoder.Token(); t != nil; t, _=decoder.Token(){
				switch se := t.(type){
				case xml.StartElement:
					if se.Name.Local == "Предложение"{
						err = decoder.DecodeElement(&product, &se)
						if err != nil {
							return 
						}
						products[product.Article] = product
					}
				}

		}
		c <- products
		
	}()
	
	for pr:=range products{
		fmt.Println(pr)
	}
	return c,nil


}