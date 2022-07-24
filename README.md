# addon-iso20022

Addition to the repository [iso20022-messages-for-go](https://github.com/prog-nov/iso20022-messages-for-go).
GO structures, XSD schemas for ISO 20022 financial messages.

## iso20022-go

Go structures for financial messages that were not included in the release for [iso20022-messages-for-go](https://github.com/prog-nov/iso20022-messages-for-go).

This is a very large container of GO files for financial messages, so it is divided into 10 sub-repositories. 

### iso20022-go: Navigate through 10 subrepositories. 

* [iso20022-go-prefix-A](https://github.com/prog-nov/iso20022-go-prefix-A)
> ISO 20022 financial message structures, **prefix A** (Only A)
* [iso20022-go-prefix-B](https://github.com/prog-nov/iso20022-go-prefix-B)
> ISO 20022 financial message structures, **prefix B** (A, B, C)
* [iso20022-go-prefix-C2](https://github.com/prog-nov/iso20022-go-prefix-C2)
> ISO 20022 financial message structures, **prefix C2** (Only C)
* [iso20022-go-prefix-D](https://github.com/prog-nov/iso20022-go-prefix-D)
> ISO 20022 financial message structures, **prefix D** (C, D, E, F)
* [iso20022-go-prefix-F](https://github.com/prog-nov/iso20022-go-prefix-F)
> ISO 20022 financial message structures, **prefix F** (F, G, H, I, L, M)
* [iso20022-go-prefix-M](https://github.com/prog-nov/iso20022-go-prefix-M)
> ISO 20022 financial message structures, **prefix M** (M, O, P)
* [iso20022-go-prefix-P](https://github.com/prog-nov/iso20022-go-prefix-P)
> ISO 20022 financial message structures, **prefix P** (P, Q, R)
* [iso20022-go-prefix-R](https://github.com/prog-nov/iso20022-go-prefix-R)
> ISO 20022 financial message structures, **prefix R** (R, S)
* [iso20022-go-prefix-S](https://github.com/prog-nov/iso20022-go-prefix-S)
> ISO 20022 financial message structures, **prefix S** (S, T)
* [iso20022-go-prefix-T](https://github.com/prog-nov/iso20022-go-prefix-T)
> ISO 20022 financial message structures, **prefix T** (T, U, V, W, Y)

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
