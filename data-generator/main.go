package main

func main() {
	const COUNT = 1000
	const DATA_DIR = "../csv/"

	WriteCustomersToCSV(generateRandomCustomers(COUNT), DATA_DIR+"customers.csv")
	WriteProductsToCSV(generateRandomProducts(COUNT), DATA_DIR+"products.csv")
	WriteSalesToCSV(generateRandomSales(COUNT), DATA_DIR+"sales.csv")
	WriteSuppliersToCSV(generateRandomSuppliers(COUNT), DATA_DIR+"suppliers.csv")
}
