package acmt

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00200102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.002.001.02 Document"`
	Message *AccountDetailsConfirmationV02 `xml:"AcctDtlsConfV02"`
}

func (d *Document00200102) AddMessage() *AccountDetailsConfirmationV02 {
	d.Message = new(AccountDetailsConfirmationV02)
	return d.Message
}

// Scope
// An account servicer, eg, a registrar, transfer agent or custodian bank sends the AccountDetailsConfirmation message to an account owner, eg, an investor to confirm the opening of an investment fund account, execution of an AccountModificationInstruction or to return information requested in a GetAccountDetails message.
// Usage
// The AccountDetailsConfirmation message is used to confirm the opening of an account, modification of an account or the provision of information requested in a previously sent GetAccountDetails message. The message contains detailed information relevant to the opened account.
// When the AccountDetailsConfirmation is used to confirm execution of an AccountModificationInstruction message, it contains the modified subsets of account details that were specified in the AccountModificationInstruction.
// When the AccountDetailsConfirmation is used to reply to a GetAccountDetails message, it returns the selected subsets of account details that were specified in the GetAccountDetails message.
type AccountDetailsConfirmationV02 struct {

	// Identifies the message.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Identifies a related order.
	OrderReference *iso20022.InvestmentFundOrder4 `xml:"OrdrRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Provide detailed information about the request or instruction which triggered this confirmation.
	ConfirmationDetails *iso20022.AccountManagementConfirmation1 `xml:"ConfDtls"`

	// Confirmation of the information related to a selected investment account.
	InvestmentAccount *iso20022.InvestmentAccount27 `xml:"InvstmtAcct,omitempty"`

	// Confirmation of information related to parties who are related to a selected investment account.
	AccountParties *iso20022.AccountParties5 `xml:"AcctPties,omitempty"`

	// Confirmation of information related to intermediaries who are related to a selected investment account.
	Intermediaries []*iso20022.Intermediary12 `xml:"Intrmies,omitempty"`

	// Placement agent for the hedge fund industry.
	Placement *iso20022.ReferredAgent1 `xml:"Plcmnt,omitempty"`

	// Eligibility conditions applicable when there is an allocation of new issues for hedge fund account opening.
	NewIssueAllocation *iso20022.NewIssueAllocation1 `xml:"NewIsseAllcn,omitempty"`

	// Confirmation of the information related to a savings plan that is related to a selected investment account.
	SavingsInvestmentPlan []*iso20022.InvestmentPlan4 `xml:"SvgsInvstmtPlan,omitempty"`

	// Confirmation of the information related to a withrawal plan that is related to a selected investment account.
	WithdrawalInvestmentPlan []*iso20022.InvestmentPlan4 `xml:"WdrwlInvstmtPlan,omitempty"`

	// Confirmation of the cash settlement standing instruction associated to the investment fund transaction.
	CashSettlement *iso20022.InvestmentFundCashSettlementInformation3 `xml:"CshSttlm,omitempty"`

	// Identifies documents to be provided for the account opening.
	ServiceLevelAgreement []*iso20022.DocumentToSend1 `xml:"SvcLvlAgrmt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountDetailsConfirmationV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	a.MessageIdentification = new(iso20022.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountDetailsConfirmationV02) AddOrderReference() *iso20022.InvestmentFundOrder4 {
	a.OrderReference = new(iso20022.InvestmentFundOrder4)
	return a.OrderReference
}

func (a *AccountDetailsConfirmationV02) AddRelatedReference() *iso20022.AdditionalReference3 {
	a.RelatedReference = new(iso20022.AdditionalReference3)
	return a.RelatedReference
}

func (a *AccountDetailsConfirmationV02) AddConfirmationDetails() *iso20022.AccountManagementConfirmation1 {
	a.ConfirmationDetails = new(iso20022.AccountManagementConfirmation1)
	return a.ConfirmationDetails
}

func (a *AccountDetailsConfirmationV02) AddInvestmentAccount() *iso20022.InvestmentAccount27 {
	a.InvestmentAccount = new(iso20022.InvestmentAccount27)
	return a.InvestmentAccount
}

func (a *AccountDetailsConfirmationV02) AddAccountParties() *iso20022.AccountParties5 {
	a.AccountParties = new(iso20022.AccountParties5)
	return a.AccountParties
}

func (a *AccountDetailsConfirmationV02) AddIntermediaries() *iso20022.Intermediary12 {
	newValue := new(iso20022.Intermediary12)
	a.Intermediaries = append(a.Intermediaries, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV02) AddPlacement() *iso20022.ReferredAgent1 {
	a.Placement = new(iso20022.ReferredAgent1)
	return a.Placement
}

func (a *AccountDetailsConfirmationV02) AddNewIssueAllocation() *iso20022.NewIssueAllocation1 {
	a.NewIssueAllocation = new(iso20022.NewIssueAllocation1)
	return a.NewIssueAllocation
}

func (a *AccountDetailsConfirmationV02) AddSavingsInvestmentPlan() *iso20022.InvestmentPlan4 {
	newValue := new(iso20022.InvestmentPlan4)
	a.SavingsInvestmentPlan = append(a.SavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV02) AddWithdrawalInvestmentPlan() *iso20022.InvestmentPlan4 {
	newValue := new(iso20022.InvestmentPlan4)
	a.WithdrawalInvestmentPlan = append(a.WithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV02) AddCashSettlement() *iso20022.InvestmentFundCashSettlementInformation3 {
	a.CashSettlement = new(iso20022.InvestmentFundCashSettlementInformation3)
	return a.CashSettlement
}

func (a *AccountDetailsConfirmationV02) AddServiceLevelAgreement() *iso20022.DocumentToSend1 {
	newValue := new(iso20022.DocumentToSend1)
	a.ServiceLevelAgreement = append(a.ServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV02) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
func ( d *Document00200102 ) String() (result string, ok bool) { return }
