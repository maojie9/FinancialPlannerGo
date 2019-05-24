package generatedata

type company struct {
	companyName              string
	companyCurrentStockPrice float32
	companyFinancialResults  []financialResults
}

type financialResults struct {
	year                  int
	totalShares           int64
	totalRevenue          float64
	netProfit             float64
	totalDividendPerShare float32
	totalEquities         float64
}

/*
	This company is based on the list of 10 blue chip companies published in Dr
	Wealth article
	https://www.drwealth.com/blue-chip-stocks/

	NOTE: #9 Global Logistic is delisted

	Source of financial result is based on SGX stock screener only just for reference
*/

func generateLatestStockPrice() []float32 {
	var stockPrice []float32

	stockPrice = append(stockPrice, 25.67) //#1 DBS Group Holdings Ltd
	stockPrice = append(stockPrice, 11.07) //#2 Oversea-Chinese Banking Corporation Limited
	stockPrice = append(stockPrice, 24.63) //#3 United Overseas Bank Limited (Singapore)
	stockPrice = append(stockPrice, 3.15)  //#4 Singapore Telecommunications Limited
	stockPrice = append(stockPrice, 62.37) //#5 Jardine Matheson Holdings Limited
	stockPrice = append(stockPrice, 6.84)  //#6 Hongkong Land Holdings Limited
	stockPrice = append(stockPrice, 6.21)  //#7 Keppel Corporation Limited
	stockPrice = append(stockPrice, 25.67) //#8 Jardine Strategic Holdings Limited
	stockPrice = append(stockPrice, 3.37)  //#9 Global Logistic Properties Limited
	stockPrice = append(stockPrice, 0.79)  //#10 Thai Beverage Public Company Limited

	return stockPrice

}

/*
	Learning point for creating function is that it must start with Capital Letter in the first
	letter then it can export to other packages

*/

func GenerateCompanyInformation() []company {

	var listOfTop10BlueChipCompanies []company

	//Get the latest stock price
	stockPrice := generateLatestStockPrice()

	//#1 DBS Group Holdings Ltd

	var dbs company
	dbs.companyName = "DBS Group Holdings Ltd"
	dbs.companyCurrentStockPrice = stockPrice[0]
	dbs.companyFinancialResults = append(dbs.companyFinancialResults, financialResults{2018, 2559460000, 13183000000, 5653000000, 1.20, 49045000000})

	listOfTop10BlueChipCompanies = append(listOfTop10BlueChipCompanies, dbs)

	//#2 Oversea-Chinese Banking Corporation Limited

	var ocbc company
	ocbc.companyName = "Oversea-Chinese Banking Corporation Limited"
	ocbc.companyCurrentStockPrice = stockPrice[1]
	ocbc.companyFinancialResults = append(ocbc.companyFinancialResults, financialResults{2018, 4211140000, 9701901000, 4675226000, 0.43, 42136870000})

	listOfTop10BlueChipCompanies = append(listOfTop10BlueChipCompanies, ocbc)

	//#3 United Overseas Bank Limited (Singapore)

	var uob company
	uob.companyName = "United Overseas Bank Limited (Singapore)"
	uob.companyCurrentStockPrice = stockPrice[2]
	uob.companyFinancialResults = append(uob.companyFinancialResults, financialResults{2018, 1671350000, 9116000000, 4008000000, 1.20, 37812792000})

	listOfTop10BlueChipCompanies = append(listOfTop10BlueChipCompanies, uob)

	//#4 Singapore Telecommunications Limited

	var singtel company
	singtel.companyName = "Singapore Telecommunications Limited"
	singtel.companyCurrentStockPrice = stockPrice[3]
	singtel.companyFinancialResults = append(singtel.companyFinancialResults, financialResults{2018, 16344330000, 17531800000, 5430300000, 0.175, 29256800000})

	listOfTop10BlueChipCompanies = append(listOfTop10BlueChipCompanies, singtel)

	//#5 Jardine Matheson Holdings Limited

	var Jardine1 company
	Jardine1.companyName = "Jardine Matheson Holdings Limited"
	Jardine1.companyCurrentStockPrice = stockPrice[4]
	Jardine1.companyFinancialResults = append(Jardine1.companyFinancialResults, financialResults{2018, 376000000, 58661744452.44, 6211438269.08, 2.34, 36336549669.29})

	listOfTop10BlueChipCompanies = append(listOfTop10BlueChipCompanies, Jardine1)

	//#6 Hongkong Land Holdings Limited

	var hongkong company
	hongkong.companyName = "Hongkong Land Holdings Limited"
	hongkong.companyCurrentStockPrice = stockPrice[5]
	hongkong.companyFinancialResults = append(hongkong.companyFinancialResults, financialResults{2018, 2341800000, 3676692714.62, 3389360572.18, 30.35, 52888150135.53})

	listOfTop10BlueChipCompanies = append(listOfTop10BlueChipCompanies, hongkong)

	//#7 Keppel Corporation Limited

	var keppel company
	keppel.companyName = "Keppel Corporation Limited"
	keppel.companyCurrentStockPrice = stockPrice[6]
	keppel.companyFinancialResults = append(keppel.companyFinancialResults, financialResults{2018, 1824890000, 5964780000, 943830000, 0.30, 11278210000})

	listOfTop10BlueChipCompanies = append(listOfTop10BlueChipCompanies, keppel)

	//#8 Jardine Strategic Holdings Limited

	var Jardine2 company
	Jardine2.companyName = "Jardine Strategic Holdings Limited"
	Jardine2.companyCurrentStockPrice = stockPrice[7]
	Jardine2.companyFinancialResults = append(Jardine2.companyFinancialResults, financialResults{2018, 1108000000, 47029774672.56, 5994937546.99, 0.47, 43480539903.03})

	listOfTop10BlueChipCompanies = append(listOfTop10BlueChipCompanies, Jardine2)

	//#9 Global Logistic Properties Limited

	var global company
	global.companyName = " Global Logistic Properties Limited"
	global.companyCurrentStockPrice = stockPrice[8]
	global.companyFinancialResults = append(global.companyFinancialResults, financialResults{})

	listOfTop10BlueChipCompanies = append(listOfTop10BlueChipCompanies, global)

	//#10 Thai Beverage Public Company Limited

	var thai company
	thai.companyName = "Thai Beverage Public Company Limited"
	thai.companyCurrentStockPrice = stockPrice[9]
	thai.companyFinancialResults = append(thai.companyFinancialResults, financialResults{2018, 25116000000, 9906040555.35, 799129222.74, 0.017, 5217106430.45})

	listOfTop10BlueChipCompanies = append(listOfTop10BlueChipCompanies, thai)

	return listOfTop10BlueChipCompanies

}
