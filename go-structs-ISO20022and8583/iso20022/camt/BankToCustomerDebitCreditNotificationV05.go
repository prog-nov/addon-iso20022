package camt

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document05400105 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.054.001.05 Document"`
	Message *BankToCustomerDebitCreditNotificationV05 `xml:"BkToCstmrDbtCdtNtfctn"`
}

func (d *Document05400105) AddMessage() *BankToCustomerDebitCreditNotificationV05 {
	d.Message = new(BankToCustomerDebitCreditNotificationV05)
	return d.Message
}

// Scope
// The BankToCustomerDebitCreditNotification message is sent by the account servicer to an account owner or to a party authorised by the account owner to receive the message. It can be used to inform the account owner, or authorised party, of single or multiple debit and/or credit entries reported to the account.
// Usage
// The BankToCustomerDebitCreditNotification message can contain reports for more than one account. It provides information for cash management and/or reconciliation.
// The BankToCustomerDebitCreditNotification message can be used to :
// - report pending and booked items;
// - notify one or more debit entries;
// - notify one or more credit entries;
// - notify a combination of debit and credit entries.
// It can include underlying details of transactions that have been included in the entry.
// It is possible that the receiver of the message is not the account owner, but a party entitled by the account owner to receive the account information (also known as recipient).
// It does not contain balance information.
type BankToCustomerDebitCreditNotificationV05 struct {

	// Common information for the message.
	GroupHeader *iso20022.GroupHeader58 `xml:"GrpHdr"`

	// Notifies debit and credit entries for the account.
	Notification []*iso20022.AccountNotification11 `xml:"Ntfctn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (b *BankToCustomerDebitCreditNotificationV05) AddGroupHeader() *iso20022.GroupHeader58 {
	b.GroupHeader = new(iso20022.GroupHeader58)
	return b.GroupHeader
}

func (b *BankToCustomerDebitCreditNotificationV05) AddNotification() *iso20022.AccountNotification11 {
	newValue := new(iso20022.AccountNotification11)
	b.Notification = append(b.Notification, newValue)
	return newValue
}

func (b *BankToCustomerDebitCreditNotificationV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	b.SupplementaryData = append(b.SupplementaryData, newValue)
	return newValue
}
func ( d *Document05400105 ) String() (result string, ok bool) { return }
