package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("sample.xml")
	if err != nil {
		panic(err)
	}
	var document *Document
	err = xml.Unmarshal(dat, &document)
	if err != nil {
		panic(err)
	}
	paymentOut, err := paymentEnricher(document)
	if err != nil {
		panic(err)
	}
	fmt.Println("Output is ", paymentOut.SettlementAmount)
	fmt.Println("Currency is ", paymentOut.SettlementCurcy)
	pay, err := json.Marshal(paymentOut)
	if err != nil{
		panic(err)
	}
	var payOut PaymentOutput
	err = json.Unmarshal(pay, &payOut)
	if err != nil {
		panic(err)
	}
	fmt.Println("Test" , payOut.SettlementAmount)

	fmt.Printf("%+v", pay)
	err = ioutil.WriteFile("output.json", pay, 0644)
	fmt.Printf("%+v", paymentOut)
}

func paymentEnricher(doc *Document) (PaymentOutput, error) {
	output := PaymentOutput{
		Credttm:          doc.FIToFICstmrCdtTrf.GrpHdr.CreDtTm,
		SettlementAmount: doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Text,
		SettlementCurcy:  doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Ccy,
		Desc:             "Test",
	}
 	return output,nil
}