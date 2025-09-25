package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create an order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list orders")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get order by id")
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("udate order by id")
}

func (o *Order) DelteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("detle order by id")
}
