package myApp

import (
	"fmt"

	"github.com/alamgir-ahosain/e-commerce-project/config"
	"github.com/alamgir-ahosain/e-commerce-project/internal"
)

func Serve() {

	//load config
	cnf := config.GetConfig()
	fmt.Println(cnf.Version)
	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.HttpPort)

	internal.Start(cnf)
}
