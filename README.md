# addon-iso20022

Addition to the repository [iso20022-messages-for-go](https://github.com/prog-nov/iso20022-messages-for-go).
GO structures, XSD schemas for ISO 20022 financial messages.

## iso20022-go

Go structures for financial messages that were not included in the release for [iso20022-messages-for-go](https://github.com/prog-nov/iso20022-messages-for-go).

* AcceptanceResult6.go
* AcceptedReason1Choice.go
* AcceptedReason3Choice.go
* AcceptedReason4Choice.go
* AcceptedReason7Choice.go
* AcceptedReason8Choice.go
* AcceptedReason9Choice.go
* AcceptedReason10Choice.go
* AcceptedReason11Choice.go
* AcceptedReason12Choice.go
* AcceptedStatus1Choice.go
* AcceptedStatus3Choice.go
* AcceptedStatus4Choice.go
* AcceptedStatus7Choice.go
* AcceptedStatus8Choice.go
* AcceptedStatus9Choice.go
* AcceptedStatus10Choice.go
* AcceptedStatusReason1.go
* AcceptedStatusReason1Choice.go
* AcceptedStatusReason1Code.go
* AcceptedStatusReason3.go
* AcceptedStatusReason4.go
* AcceptedStatusReason7.go
* AcceptedStatusReason8.go
* AccountAndBalance11.go
* AccountAndBalance15.go
* AccountAndBalance16.go
* AccountAndBalance17.go
* AccountAndBalance21.go
* AccountAndBalance22.go
* OriginalPaymentInstruction12.go
* OriginalPaymentInstruction13.go
* OriginalPaymentInstruction14.go
* OriginalPaymentInstruction15.go
* OriginalPaymentInstruction16.go
* OriginalPaymentInstruction17.go
* OriginalPaymentInstruction18.go
* OriginalPaymentInstruction19.go
* OriginalPaymentInstruction20.go
* OriginalPaymentInstruction21.go
* OriginalPaymentInstruction22.go
* OriginalPaymentInstruction23.go
* OriginalPaymentInstruction24.go
* OriginalRequestInformation1.go
* OriginalTransactionReference1.go
* OriginalTransactionReference13.go
* OriginalTransactionReference14.go
* OriginalTransactionReference15.go
* OriginalTransactionReference16.go
* OriginalTransactionReference17.go
* OriginalTransactionReference18.go
* OriginalTransactionReference19.go
* OriginalTransactionReference20.go
* OriginalTransactionReference21.go
* and other. (see. `/iso20022-go`).

## go-structs-ISO20022and8583

Thanks to **Assis Ngolo** for this part of the repository.

Go implementation of the ISO 20022 and 8583 standards
This package consumes valid iso20022 XML strings and returns Go structs to be processed by an iso20022 application.

> see. `addon-iso20022/go-structs-ISO20022and8583`.

## messageprocessor-go

Go structs: PaymentInput.go and PaymentOutput.go.

### Example

``` xml type Document struct {
	XMLName           xml.Name `xml:"Document"`
	Text              string   `xml:",chardata"`
	Xmlns             string   `xml:"xmlns,attr"`
	FIToFICstmrCdtTrf struct {
		Text   string `xml:",chardata"`
		GrpHdr struct {
			Text              string `xml:",chardata"`
			MsgId             string `xml:"MsgId"`
			CreDtTm           string `xml:"CreDtTm"`
			NbOfTxs           string `xml:"NbOfTxs"`
			TtlIntrBkSttlmAmt struct {
				Text string `xml:",chardata"`
				Ccy  string `xml:"Ccy,attr"`
			} `xml:"TtlIntrBkSttlmAmt"`
			IntrBkSttlmDt string `xml:"IntrBkSttlmDt"`
			SttlmInf      struct {
				Text     string `xml:",chardata"`
				SttlmMtd string `xml:"SttlmMtd"`
				ClrSys   struct {
					Text  string `xml:",chardata"`
					Prtry string `xml:"Prtry"`
				} `xml:"ClrSys"`
			} `xml:"SttlmInf"`
			InstgAgt struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					BIC  string `xml:"BIC"`
				} `xml:"FinInstnId"`
			} `xml:"InstgAgt"`
			InstdAgt struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					BIC  string `xml:"BIC"`
				} `xml:"FinInstnId"`
			} `xml:"InstdAgt"`
		} `xml:"GrpHdr"`
		CdtTrfTxInf struct {
			Text  string `xml:",chardata"`
			PmtId struct {
				Text       string `xml:",chardata"`
				InstrId    string `xml:"InstrId"`
				EndToEndId string `xml:"EndToEndId"`
				TxId       string `xml:"TxId"`
			} `xml:"PmtId"`
			PmtTpInf struct {
				Text   string `xml:",chardata"`
				SvcLvl struct {
					Text string `xml:",chardata"`
					Cd   string `xml:"Cd"`
				} `xml:"SvcLvl"`
			} `xml:"PmtTpInf"`
			IntrBkSttlmAmt struct {
				Text string `xml:",chardata"`
				Ccy  string `xml:"Ccy,attr"`
			} `xml:"IntrBkSttlmAmt"`
			ChrgBr string `xml:"ChrgBr"`
			Dbtr   struct {
				Text    string `xml:",chardata"`
				Nm      string `xml:"Nm"`
				PstlAdr struct {
					Text string `xml:",chardata"`
					Ctry string `xml:"Ctry"`
				} `xml:"PstlAdr"`
				ID struct {
					Text  string `xml:",chardata"`
					OrgId struct {
						Text string `xml:",chardata"`
						Othr struct {
							Text string `xml:",chardata"`
							ID   string `xml:"Id"`
						} `xml:"Othr"`
					} `xml:"OrgId"`
				} `xml:"Id"`
			} `xml:"Dbtr"`
			DbtrAcct struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text string `xml:",chardata"`
					IBAN string `xml:"IBAN"`
				} `xml:"Id"`
			} `xml:"DbtrAcct"`
			DbtrAgt struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					BIC  string `xml:"BIC"`
				} `xml:"FinInstnId"`
			} `xml:"DbtrAgt"`
			CdtrAgt struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					BIC  string `xml:"BIC"`
				} `xml:"FinInstnId"`
			} `xml:"CdtrAgt"`
			Cdtr struct {
				Text    string `xml:",chardata"`
				Nm      string `xml:"Nm"`
				PstlAdr struct {
					Text string `xml:",chardata"`
					Ctry string `xml:"Ctry"`
				} `xml:"PstlAdr"`
				ID struct {
					Text  string `xml:",chardata"`
					OrgId struct {
						Text string `xml:",chardata"`
						Othr struct {
							Text string `xml:",chardata"`
							ID   string `xml:"Id"`
						} `xml:"Othr"`
					} `xml:"OrgId"`
				} `xml:"Id"`
			} `xml:"Cdtr"`
			CdtrAcct struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text string `xml:",chardata"`
					IBAN string `xml:"IBAN"`
				} `xml:"Id"`
			} `xml:"CdtrAcct"`
			RmtInf struct {
				Text  string `xml:",chardata"`
				Ustrd string `xml:"Ustrd"`
			} `xml:"RmtInf"`
		} `xml:"CdtTrfTxInf"`
	} `xml:"FIToFICstmrCdtTrf"`
}
```
