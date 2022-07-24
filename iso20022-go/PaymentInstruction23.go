package iso20022

// Instruction to pay an amount of money to an ultimate beneficiary, on behalf of an originator. This instruction may have to be forwarded several times to complete the settlement chain.
type PaymentInstruction23 struct {

	// Reference assigned by a sending party to unambiguously identify the payment information block within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod7Code `xml:"PmtMtd"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment. If payment by cheque, the date when the cheque must be generated by the bank.
	//
	// Usage: This is the date on which the debtor's account(s) is (are) to be debited.
	RequestedExecutionDate *DateAndDateTimeChoice `xml:"ReqdExctnDt"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Account used to process charges associated with a transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CreditTransferTransaction []*CreditTransferTransaction22 `xml:"CdtTrfTx"`
}

func (p *PaymentInstruction23) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction23) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod7Code)(&value)
}

func (p *PaymentInstruction23) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction23) AddRequestedExecutionDate() *DateAndDateTimeChoice {
	p.RequestedExecutionDate = new(DateAndDateTimeChoice)
	return p.RequestedExecutionDate
}

func (p *PaymentInstruction23) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction23) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction23) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction23) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction23) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction23) AddCreditTransferTransaction() *CreditTransferTransaction22 {
	newValue := new(CreditTransferTransaction22)
	p.CreditTransferTransaction = append(p.CreditTransferTransaction, newValue)
	return newValue
}
