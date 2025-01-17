package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
	"gorm.io/gorm"
)

func GetCustomerAssociation(customers []model.Customer) {
	// For each seller
	for i := 0; i < len(customers); i++ {
		GetOneCustomerAssociation(&customers[i])
	}
}

func GetOneCustomerAssociation(customer *model.Customer) {
	// Check connection
	con := dbController.GetConnection()

	// with user data
	con.Model(&customer).Association("User").Find(&customer.User)

	// with cart
	con.Model(&customer).Association("Cart").Find(&customer.Cart)

	// with list of transaction history
	con.Model(&customer).Association("ListHistory").Find(&customer.ListHistory)
}

func ReadAllCustomers(custId []string) []model.Customer {
	// Check connection
	con := dbController.GetConnection()

	// Get customers from database
	var customers []model.Customer

	// Extra query by id
	if custId != nil {
		con.Find(&customers, custId[0])
	} else {
		con.Find(&customers)
	}

	return customers
}

func ReadCustomerAddress(custId int) string {
	// Check connection
	con := dbController.GetConnection()

	// Get customers from database
	var address string

	// Extra query by id
	con.Raw("SELECT address FROM customers WHERE user_id = ?", custId).Scan(&address)

	return address
}

func CreateCustomer(customer model.Customer) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&customer)

	return result
}

func DeleteCustomerByCustId(customerId string) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	//delete Customer on db by id
	var cust model.Customer
	result := con.Where("user_id = ?", customerId).Delete(&cust)
	return result
}

func SaveCustomer(cust model.Customer) *gorm.DB {
	con := dbController.GetConnection()
	result := con.Save(&cust)
	return result
}
