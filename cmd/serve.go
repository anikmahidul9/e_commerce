package cmd

import "ecommerce/config"
import "ecommerce/rest"

func Serve() {
	cnf := config.GetConfig()

	rest.Server(cnf)
}