package main

import "github.com/alamgir-ahosain/e-commerce-project/cmd/myApp"

/* ___status code
   200 : ok
   201 : create server resource
   400 : bad request
   404 : not found
   500 : internal server error
*/

func main() {

	myApp.Serve()

}
