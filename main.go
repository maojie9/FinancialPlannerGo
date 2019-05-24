package main

import (
	"fmt"

	generatedata "github.com/maojie9/FinancialPlannerGo/datasource"
)

func main() {
	var s = generatedata.GenerateCompanyInformation()

	//Show the raw data
	fmt.Println(s)

	//Generate P/B ratio
	generatePB(s)

}

func generatePB(s []generatedata.Company) {

	for i := 0; i < len(s); i++ {
		//Generate book value
		companyName := s[i]

		fmt.Println(companyName)
	}

}
