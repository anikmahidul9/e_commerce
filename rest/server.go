package rest

import (
	"fmt"
	"ecommerce/config"
	middleware "ecommerce/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

func Server(cnf config.Config){
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux:= http.NewServeMux()
    initRoutes(mux,manager)
	wrappedMux := manager.WrapMux(mux)
	

	addr := ":"+strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server is running on port", addr)

	err := http.ListenAndServe(addr,wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}

}