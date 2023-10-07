package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/brianvoe/gofakeit"
)

type Customer struct {
	CustomerID int    `fake:"{number:1,10}"`
	FirstName  string `json:"first_name" fake:"{firstname}"`
	LastName   string `json:"last_name" fake:"{lastname}"`
	Email      string `json:"email" fake:"{email}"`
	Phone      string `json:"phone" fake:"{phone}"`
}

type Product struct {
	ProductID     int     `json:"product_id" fake:"{number:1,1000}"`
	ProductName   string  `json:"product_name" fake:"{word}"`
	Brand         string  `json:"brand" fake:"{word}"`
	Price         float64 `json:"price" fake:"{price}"`
	StockQuantity int     `json:"stock_quantity" fake:"{number:1,100}"`
	SupplierID    int     `json:"supplier_id" fake:"{number:1,1000}"`
}

type Sale struct {
	SaleID       int     `json:"sale_id" fake:"{number:1,1000}"`
	CustomerID   int     `json:"customer_id" fake:"{number:1,1000}"`
	ProductID    int     `json:"product_id" fake:"{number:1,1000}"`
	SaleDate     string  `json:"sale_date" fake:"{date}"`
	QuantitySold int     `json:"quantity_sold" fake:"{number:1,100}"`
	TotalPrice   float64 `json:"total_price" fake:"{price}"`
}

type Supplier struct {
	SupplierID    int    `json:"supplier_id" fake:"{number:1,1000}"`
	SupplierName  string `json:"supplier_name" fake:"{company}"`
	ContactPerson string `json:"contact_person" fake:"{fullname}"`
	Phone         string `json:"phone" fake:"{phone}"`
	Email         string `json:"email" fake:"{email}"`
}

func generateRandomCustomers(n int) []Customer {
	customers := make([]Customer, 0, n)
	for i := 0; i < n; i++ {
		customer := Customer{
			CustomerID: gofakeit.Number(1, 1000),
			FirstName:  gofakeit.FirstName(),
			LastName:   gofakeit.LastName(),
			Email:      gofakeit.Email(),
			Phone:      gofakeit.Phone(),
		}
		customers = append(customers, customer)
	}
	return customers
}

func generateRandomProducts(n int) []Product {
	products := make([]Product, 0, n)
	for i := 0; i < n; i++ {
		product := Product{
			ProductID:     gofakeit.Number(1, 1000),
			ProductName:   gofakeit.Word(),
			Brand:         gofakeit.Company(),
			Price:         gofakeit.Price(10, 1000),
			StockQuantity: gofakeit.Number(1, 100),
			SupplierID:    gofakeit.Number(1, 50),
		}
		products = append(products, product)
	}
	return products
}

func generateRandomSales(n int) []Sale {
	sales := make([]Sale, 0, n)
	for i := 0; i < n; i++ {
		sale := Sale{
			SaleID:       gofakeit.Number(1, 1000),
			CustomerID:   gofakeit.Number(1, 100),
			ProductID:    gofakeit.Number(1, 100),
			SaleDate:     gofakeit.Date().Format("2006-01-02"),
			QuantitySold: gofakeit.Number(1, 10),
			TotalPrice:   gofakeit.Price(10, 500),
		}
		sales = append(sales, sale)
	}
	return sales
}

func generateRandomSuppliers(n int) []Supplier {
	suppliers := make([]Supplier, 0, n)
	for i := 0; i < n; i++ {
		supplier := Supplier{
			SupplierID:    gofakeit.Number(1, 50),
			SupplierName:  gofakeit.Company(),
			ContactPerson: gofakeit.Name(),
			Phone:         gofakeit.Phone(),
			Email:         gofakeit.Email(),
		}
		suppliers = append(suppliers, supplier)
	}
	return suppliers
}

func WriteCustomersToCSV(customers []Customer, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, customer := range customers {
		err := writer.Write([]string{
			strconv.Itoa(customer.CustomerID),
			customer.FirstName,
			customer.LastName,
			customer.Email,
			customer.Phone,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteProductsToCSV(products []Product, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, product := range products {
		err := writer.Write([]string{
			strconv.Itoa(product.ProductID),
			product.ProductName,
			product.Brand,
			strconv.FormatFloat(product.Price, 'f', -1, 64),
			strconv.Itoa(product.StockQuantity),
			strconv.Itoa(product.SupplierID),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteSalesToCSV(sales []Sale, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, sale := range sales {
		err := writer.Write([]string{
			strconv.Itoa(sale.SaleID),
			strconv.Itoa(sale.CustomerID),
			strconv.Itoa(sale.ProductID),
			sale.SaleDate,
			strconv.Itoa(sale.QuantitySold),
			strconv.FormatFloat(sale.TotalPrice, 'f', -1, 64),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteSuppliersToCSV(suppliers []Supplier, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, supplier := range suppliers {
		err := writer.Write([]string{
			strconv.Itoa(supplier.SupplierID),
			supplier.SupplierName,
			supplier.ContactPerson,
			supplier.Phone,
			supplier.Email,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
