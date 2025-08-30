package main

import (
	"fmt"

	//	"github.com/alamgir-ahosain/e-commerce-project/cmd/myApp"
	"github.com/alamgir-ahosain/e-commerce-project/cmd/myApp"
	"github.com/alamgir-ahosain/e-commerce-project/config"
)

/* ___status code
   200 : ok
   201 : create server resource
   400 : bad request
   404 : not found
   500 : internal server error
*/

func main() {

	cnf := config.GetConfig()
	fmt.Println(cnf.Version)
	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.HttpPort)

	myApp.Serve()

}
