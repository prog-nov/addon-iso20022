package main

import "encoding/xml"

type Document struct {
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
