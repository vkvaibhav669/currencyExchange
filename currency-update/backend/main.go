package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type APIResponse struct {
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
	Base      string `json:"base"`
	Date      string `json:"date"`
	Rates     struct {
		Aed float64 `json:"AED"`
		Afn float64 `json:"AFN"`
		All float64 `json:"ALL"`
		Amd float64 `json:"AMD"`
		Ang float64 `json:"ANG"`
		Aoa float64 `json:"AOA"`
		Ars float64 `json:"ARS"`
		Aud float64 `json:"AUD"`
		Awg float64 `json:"AWG"`
		Azn float64 `json:"AZN"`
		Bam float64 `json:"BAM"`
		Bbd float64 `json:"BBD"`
		Bdt float64 `json:"BDT"`
		Bgn float64 `json:"BGN"`
		Bhd float64 `json:"BHD"`
		Bif float64 `json:"BIF"`
		Bmd float64 `json:"BMD"`
		Bnd float64 `json:"BND"`
		Bob float64 `json:"BOB"`
		Brl float64 `json:"BRL"`
		Bsd float64 `json:"BSD"`
		Btc float64 `json:"BTC"`
		Btn float64 `json:"BTN"`
		Bwp float64 `json:"BWP"`
		Byn float64 `json:"BYN"`
		Byr float64 `json:"BYR"`
		Bzd float64 `json:"BZD"`
		Cad float64 `json:"CAD"`
		Cdf float64 `json:"CDF"`
		Chf float64 `json:"CHF"`
		Clf float64 `json:"CLF"`
		Clp float64 `json:"CLP"`
		Cny float64 `json:"CNY"`
		Cop float64 `json:"COP"`
		Crc float64 `json:"CRC"`
		Cuc float64 `json:"CUC"`
		Cup float64 `json:"CUP"`
		Cve float64 `json:"CVE"`
		Czk float64 `json:"CZK"`
		Djf float64 `json:"DJF"`
		Dkk float64 `json:"DKK"`
		Dop float64 `json:"DOP"`
		Dzd float64 `json:"DZD"`
		Egp float64 `json:"EGP"`
		Ern float64 `json:"ERN"`
		Etb float64 `json:"ETB"`
		Eur int     `json:"EUR"`
		Fjd float64 `json:"FJD"`
		Fkp float64 `json:"FKP"`
		Gbp float64 `json:"GBP"`
		Gel float64 `json:"GEL"`
		Ggp float64 `json:"GGP"`
		Ghs float64 `json:"GHS"`
		Gip float64 `json:"GIP"`
		Gmd float64 `json:"GMD"`
		Gnf float64 `json:"GNF"`
		Gtq float64 `json:"GTQ"`
		Gyd float64 `json:"GYD"`
		Hkd float64 `json:"HKD"`
		Hnl float64 `json:"HNL"`
		Hrk float64 `json:"HRK"`
		Htg float64 `json:"HTG"`
		Huf float64 `json:"HUF"`
		Idr float64 `json:"IDR"`
		Ils float64 `json:"ILS"`
		Imp float64 `json:"IMP"`
		Inr float64 `json:"INR"`
		Iqd float64 `json:"IQD"`
		Irr float64 `json:"IRR"`
		Isk float64 `json:"ISK"`
		Jep float64 `json:"JEP"`
		Jmd float64 `json:"JMD"`
		Jod float64 `json:"JOD"`
		Jpy float64 `json:"JPY"`
		Kes float64 `json:"KES"`
		Kgs float64 `json:"KGS"`
		Khr float64 `json:"KHR"`
		Kmf float64 `json:"KMF"`
		Kpw float64 `json:"KPW"`
		Krw float64 `json:"KRW"`
		Kwd float64 `json:"KWD"`
		Kyd float64 `json:"KYD"`
		Kzt float64 `json:"KZT"`
		Lak float64 `json:"LAK"`
		Lbp float64 `json:"LBP"`
		Lkr float64 `json:"LKR"`
		Lrd float64 `json:"LRD"`
		Lsl float64 `json:"LSL"`
		Ltl float64 `json:"LTL"`
		Lvl float64 `json:"LVL"`
		Lyd float64 `json:"LYD"`
		Mad float64 `json:"MAD"`
		Mdl float64 `json:"MDL"`
		Mga float64 `json:"MGA"`
		Mkd float64 `json:"MKD"`
		Mmk float64 `json:"MMK"`
		Mnt float64 `json:"MNT"`
		Mop float64 `json:"MOP"`
		Mro float64 `json:"MRO"`
		Mur float64 `json:"MUR"`
		Mvr float64 `json:"MVR"`
		Mwk float64 `json:"MWK"`
		Mxn float64 `json:"MXN"`
		Myr float64 `json:"MYR"`
		Mzn float64 `json:"MZN"`
		Nad float64 `json:"NAD"`
		Ngn float64 `json:"NGN"`
		Nio float64 `json:"NIO"`
		Nok float64 `json:"NOK"`
		Npr float64 `json:"NPR"`
		Nzd float64 `json:"NZD"`
		Omr float64 `json:"OMR"`
		Pab float64 `json:"PAB"`
		Pen float64 `json:"PEN"`
		Pgk float64 `json:"PGK"`
		Php float64 `json:"PHP"`
		Pkr float64 `json:"PKR"`
		Pln float64 `json:"PLN"`
		Pyg float64 `json:"PYG"`
		Qar float64 `json:"QAR"`
		Ron float64 `json:"RON"`
		Rsd float64 `json:"RSD"`
		Rub float64 `json:"RUB"`
		Rwf float64 `json:"RWF"`
		Sar float64 `json:"SAR"`
		Sbd float64 `json:"SBD"`
		Scr float64 `json:"SCR"`
		Sdg float64 `json:"SDG"`
		Sek float64 `json:"SEK"`
		Sgd float64 `json:"SGD"`
		Shp float64 `json:"SHP"`
		Sle float64 `json:"SLE"`
		Sll float64 `json:"SLL"`
		Sos float64 `json:"SOS"`
		Srd float64 `json:"SRD"`
		Std float64 `json:"STD"`
		Svc float64 `json:"SVC"`
		Syp float64 `json:"SYP"`
		Szl float64 `json:"SZL"`
		Thb float64 `json:"THB"`
		Tjs float64 `json:"TJS"`
		Tmt float64 `json:"TMT"`
		Tnd float64 `json:"TND"`
		Top float64 `json:"TOP"`
		Try float64 `json:"TRY"`
		Ttd float64 `json:"TTD"`
		Twd float64 `json:"TWD"`
		Tzs float64 `json:"TZS"`
		Uah float64 `json:"UAH"`
		Ugx float64 `json:"UGX"`
		Usd float64 `json:"USD"`
		Uyu float64 `json:"UYU"`
		Uzs float64 `json:"UZS"`
		Vef float64 `json:"VEF"`
		Ves float64 `json:"VES"`
		Vnd float64 `json:"VND"`
		Vuv float64 `json:"VUV"`
		Wst float64 `json:"WST"`
		Xaf float64 `json:"XAF"`
		Xag float64 `json:"XAG"`
		Xau float64 `json:"XAU"`
		Xcd float64 `json:"XCD"`
		Xdr float64 `json:"XDR"`
		Xof float64 `json:"XOF"`
		Xpf float64 `json:"XPF"`
		Yer float64 `json:"YER"`
		Zar float64 `json:"ZAR"`
		Zmk float64 `json:"ZMK"`
		Zmw float64 `json:"ZMW"`
		Zwl float64 `json:"ZWL"`
	} `json:"rates"`
}

var db *sql.DB

func main() {

	fmt.Println("Hello from currency exchanger program ! ")
	//dbURL := "postgres://root:admin123@db:3306/your_dbname?sslmode=disable"
	exp := "http://api.exchangeratesapi.io/v1/latest?access_key=02b5335542f0861caa327132f0484caa&format=1&base=EUR&symbol=USD,AMD,GBP"
	//exchangeAPI := "http://api.exchangeratesapi.io/v1/latest?access_key=02b5335542f0861caa327132f0484caa&format=1 &base = GBP & symbols = USD,EUR,GBP"
	fmt.Println(exp)

	resp, err := http.Get(exp)
	if err != nil {
		// handle error
		fmt.Println(resp)
	}
	//fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("err")
	}
	//fmt.Println(string(body))
	fmt.Println(string(body))

	//defer resp.Body.Close()
	var apiResults []APIResponse
	if err := json.Unmarshal(body, &apiResults); err != nil {
		log.Fatal("Failed to parse API response:", err)
	}

	// Insert API results into the SQL database
	for _, result := range apiResults {
		if err := insertAPIResult(db, result); err != nil {
			log.Printf("Failed to insert API result with ID %d: %v", result.Timestamp, err)
		}
	}
	defer resp.Body.Close()

}

func insertAPIResult(db *sql.DB, result APIResponse) error {
	_, err := db.Exec("INSERT INTO api_results (success, timestamp, base, date, rates) VALUES ($1, $2, $3, $4, $5)",
		result.Success, result.Timestamp, result.Base, result.Date, result.Rates)
	return err
}

// 	var err error
// 	db, err = sql.Open("postgres", dbURL)
// 	if err != nil {
// 		log.Fatal("Failed to connect to the database:", err)
// 	}
// 	defer db.Close()

// 	http.HandleFunc("/api/exchange-rates", handleExchangeRates)
// 	log.Println("Server started at :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }

// func handleExchangeRates(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	rates, err := getExchangeRates()
// 	if err != nil {
// 		http.Error(w, "Failed to fetch exchange rates", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(rates)
// }

// func getExchangeRates() ([]ExchangeRate, error) {
// 	rows, err := db.Query("SELECT * FROM exchange_rates ORDER BY last_update DESC LIMIT 10")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var rates []ExchangeRate
// 	for rows.Next() {
// 		var rate ExchangeRate
// 		err := rows.Scan(&rate.ID, &rate.Currency, &rate.Rate, &rate.LastUpdate)
// 		if err != nil {
// 			return nil, err
// 		}
// 		rates = append(rates, rate)
// 	}

// 	return rates, nil
// }
