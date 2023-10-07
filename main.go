package main

func main() {
	const COUNT = 10
	WriteCustomersToCSV(generateRandomCustomers(COUNT), "csv/customers.csv")
	WriteProductsToCSV(generateRandomProducts(COUNT), "csv/products.csv")
	WriteSalesToCSV(generateRandomSales(COUNT), "csv/sales.csv")
	WriteSuppliersToCSV(generateRandomSuppliers(COUNT), "csv/suppliers.csv")
}
