package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (order *Order) GetOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an Order")
}

func (order *Order) GetOrderById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an Order by Id")
}

func (order *Order) CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Order")
}

func (order *Order) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Order")
}

func (order *Order) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Order")
}
