package main

import (
	"fmt"
	"net/http"

	"github.com/Haidar1528/mongodb_mux_golang/helper"
	"github.com/Haidar1528/mongodb_mux_golang/router"
)

func main() {
	fmt.Println("Server is starting ...")
	r := router.Router()
	err := http.ListenAndServe(":4000", r)
	fmt.Println("Server is listening at 4000")

	helper.HandleErr(err)
}
