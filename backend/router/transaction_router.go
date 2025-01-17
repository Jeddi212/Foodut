package router

import (
	trController "github.com/Foodut/backend/modules/transaction/rest-api/controller"
	usrController "github.com/Foodut/backend/modules/user/rest-api/controller"
	"github.com/gorilla/mux"
)

func TransactionRouter(router *mux.Router) {

	//  TRANSACTION
	//
	//- Get All Transaction
	router.HandleFunc("/transactions", usrController.AuthenticateMinimumLevel(trController.GetAllTransactions, 3)).Methods("GET")

	//- Get All Transaction Detail
	router.HandleFunc("/transactions-d", trController.GetAllTransactionsDecentralize).Methods("GET")

	//- Get Customer Transaction
	router.HandleFunc("/transactions-c", usrController.Authenticate(trController.GetCustomerTransactions, 1)).Methods("GET")

	//- Get Customer Transaction Detail
	router.HandleFunc("/transactions-cd", trController.GetCustomerTransactionsDecentralize).Methods("GET")

	//- Get All Order From Seller
	router.HandleFunc("/orders", trController.GetAllOrders).Methods("GET")

	//- Get Carts With Availability
	router.HandleFunc("/cart-a", trController.GetCartWithAvailability).Methods("GET")

	//- Insert Cart
	router.HandleFunc("/cart", usrController.Authenticate(trController.PostToCart, 1)).Methods("POST")

	//- Insert Transaction
	router.HandleFunc("/transactions", usrController.AuthenticateMinimumLevel(trController.PostTransaction, 1)).Methods("POST")

	//- Update Cart
	router.HandleFunc("/cart", usrController.Authenticate(trController.UpdateCart, 1)).Methods("PUT")

	//- Delete Carts
	router.HandleFunc("/cart", trController.DeleteCarts).Methods("DELETE")

	//- Delete Specific Product from a Cart (Update w/ delete spesific item)
	router.HandleFunc("/cart-specific", trController.DeleteSpesificProductFromCart).Methods("DELETE")

	//- Update Transaction
	// router.HandleFunc("/transactions/{transaction_id}", trController.UpdateTransaction).Methods("PUT")

	//- Delete Transaction
	router.HandleFunc("/transactions/{id}", usrController.AuthenticateMinimumLevel(trController.DeleteTransaction, 3)).Methods("DELETE")
}
