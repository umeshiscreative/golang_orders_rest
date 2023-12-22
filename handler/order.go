package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (order *Order) GetOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an Order")
}
