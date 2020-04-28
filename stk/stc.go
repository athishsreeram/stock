package main

import (
    "io/ioutil"
    "log"
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
	"time"
)

type Financials struct {
	Date                       string `json:"date"`
	Revenue                    string `json:"Revenue"`
	RevenueGrowth              string `json:"Revenue Growth"`
	CostOfRevenue              string `json:"Cost of Revenue"`
	GrossProfit                string `json:"Gross Profit"`
	RDExpenses                 string `json:"R&D Expenses"`
	SGAExpense                 string `json:"SG&A Expense"`
	OperatingExpenses          string `json:"Operating Expenses"`
	OperatingIncome            string `json:"Operating Income"`
	InterestExpense            string `json:"Interest Expense"`
	EarningsBeforeTax          string `json:"Earnings before Tax"`
	IncomeTaxExpense           string `json:"Income Tax Expense"`
	NetIncomeNonControllingInt string `json:"Net Income - Non-Controlling int"`
	NetIncomeDiscontinuedOps   string `json:"Net Income - Discontinued ops"`
	NetIncome                  string `json:"Net Income"`
	PreferredDividends         string `json:"Preferred Dividends"`
	NetIncomeCom               string `json:"Net Income Com"`
	EPS                        string `json:"EPS"`
	EPSDiluted                 string `json:"EPS Diluted"`
	WeightedAverageShsOut      string `json:"Weighted Average Shs Out"`
	WeightedAverageShsOutDil   string `json:"Weighted Average Shs Out (Dil)"`
	DividendPerShare           string `json:"Dividend per Share"`
	GrossMargin                string `json:"Gross Margin"`
	EBITDAMargin               string `json:"EBITDA Margin"`
	EBITMargin                 string `json:"EBIT Margin"`
	ProfitMargin               string `json:"Profit Margin"`
	FreeCashFlowMargin         string `json:"Free Cash Flow margin"`
	EBITDA                     string `json:"EBITDA"`
	EBIT                       string `json:"EBIT"`
	ConsolidatedIncome         string `json:"Consolidated Income"`
	EarningsBeforeTaxMargin    string `json:"Earnings Before Tax Margin"`
	NetProfitMargin            string `json:"Net Profit Margin"`
} 

type FST struct {
	Symbol     string `json:"symbol"`
	Financials []Financials `json:"financials"`
}

func main(){

	body := MakeRequest()
	var fst FST

	json.Unmarshal([]byte(body), &fst)

	for index, element := range fst.Financials {
		fmt.Println("++++++++")
		fmt.Println(index, "=>", element)
		PostRequest(element)
	}

}

func MakeRequest()[]byte{
	resp, err := http.Get("https://financialmodelingprep.com/api/v3/financials/income-statement/COST")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}


	return body
	
}

func PostRequest(reqFin Financials){
	url := "http://127.0.0.1:5601/api/console/proxy?path=%2Fcost%2F_doc&method=POST"

	dataFin, err := json.Marshal(reqFin)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(dataFin))
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataFin))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("kbn-xsrf", "true")

	// Create and Add cookie to request
	// cookie := http.Cookie{Name: "cookie_name", Value: "cookie_value"}
	// req.AddCookie(&cookie)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Validate cookie and headers are attached
	// fmt.Println(req.Cookies())
	fmt.Println(req.Header)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Printf("%s\n", body)
}