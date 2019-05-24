package main

import (
	"fmt"

	generatedata "github.com/maojie9/FinancialPlannerGo/datasource"
)

func main() {
	var s = generatedata.GenerateCompanyInformation()
	fmt.Println(s)
}
