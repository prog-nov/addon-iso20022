package iso20022

const (
	FININSTNID             = "FinInstnId"
	STTLMMTD               = "SttlmMtd"
	INSTGAGT               = "InstgAgt"
	ORGNLGRPINF            = "OrgnlGrpInf"
	ORGNLINSTRID           = "OrgnlInstrId"
	ORGNLENDTOENDID        = "OrgnlEndToEndId"
	ORGNLTXID              = "OrgnlTxId"
	ORGNLINTRBKSTTLMAMT    = "OrgnlIntrBkSttlmAmt"
	RTRDINTRBKSTTLMAMT     = "RtrdIntrBkSttlmAmt"
	INTRBKSTTLMDT          = "IntrBkSttlmDt"
	RTRDINSTDAMT           = "RtrdInstdAmt"
	CHRGBR                 = "ChrgBr"
	ORGNLTXREF             = "OrgnlTxRef"
	MSGDEFIDR              = "MsgDefIdr"
	INSTDAGT               = "InstdAgt"
	INTRBKSTTLMAMT         = "IntrBkSttlmAmt"
	INSTDAMT               = "InstdAmt"
	XCHGRATE               = "XchgRate"
	DBTRAGT                = "DbtrAgt"
	CDTRAGT                = "CdtrAgt"
	ENVLP                  = "Envlp"
	SPLMTRYDATA            = "SplmtryData"
	FITOFICSTMRCDTTRF      = "FIToFICstmrCdtTrf"
	STTLMACCT              = "SttlmAcct"
	INSTRID                = "InstrId"
	ENDTOENDID             = "EndToEndId"
	TXID                   = "TxId"
	ORGNLGRPINFANDSTS      = "OrgnlGrpInfAndSts"
	JUSTFN                 = "Justfn"
	ORGNLMSGNMID           = "OrgnlMsgNmId"
	ORGNLMSGID             = "OrgnlMsgId"
	ORGNLGRPINFANDCXL      = "OrgnlGrpInfAndCxl"
	STTLMINF               = "SttlmInf"
	ISSR                   = "Issr"
	ASSGNR                 = "Assgnr"
	BIZMSGIDR              = "BizMsgIdr"
	CREDT                  = "CreDt"
	FITOFIPMTSTSRPT        = "FIToFIPmtStsRpt"
	MSGID                  = "MsgId"
	CREDTTM                = "CreDtTm"
	FIID                   = "FIId"
	BICFI                  = "BICFI"
	GRPHDR                 = "GrpHdr"
	TXINFANDSTS            = "TxInfAndSts"
	TXSTS                  = "TxSts"
	CHRGSINF               = "ChrgsInf"
	PLCANDNM               = "PlcAndNm"
	NBOFTXS                = "NbOfTxs"
	ORGNLCREDTTM           = "OrgnlCreDtTm"
	RTRID                  = "RtrId"
	RTRRSNINF              = "RtrRsnInf"
	ORGTR                  = "Orgtr"
	SVCLVL                 = "SvcLvl"
	PMTTPINF               = "PmtTpInf"
	DBTR                   = "Dbtr"
	PSTLADR                = "PstlAdr"
	STRTNM                 = "StrtNm"
	BLDGNB                 = "BldgNb"
	PSTCD                  = "PstCd"
	TWNNM                  = "TwnNm"
	ADRLINE                = "AdrLine"
	CTRY                   = "Ctry"
	FRMSGHEAD              = "Fr"
	TOMSGHEAD              = "To"
	CDTTRFTXINF            = "CdtTrfTxInf"
	AGT                    = "Agt"
	UNDRLYG                = "Undrlyg"
	MSSNGORINCRRCTINF      = "MssngOrIncrrctInf"
	UBLTOAPPLY             = "UblToApply"
	IDMSGHEAD              = "Id"
	CONF                   = "Conf"
	RSLTNOFINVSTGTN        = "RsltnOfInvstgtn"
	SIGNATUREMETHOD        = "SignatureMethod"
	CANONICALIZATIONMETHOD = "CanonicalizationMethod"
	NTFCTNOFCASEASSGNMT    = "NtfctnOfCaseAssgnmt"
	X509DATA               = "X509Data"
	KEYINFO                = "KeyInfo"
	SIGNATURE              = "Signature"
	CRETR                  = "Cretr"
	ASSGNE                 = "Assgne"
	CXLRSNINF              = "CxlRsnInf"
	FITOFIPMTCXLREQ        = "FIToFIPmtCxlReq"
	ORGNLINSTDAMT          = "OrgnlInstdAmt"
	ULTMTCDTR              = "UltmtCdtr"
	MOD                    = "Mod"
	REQTOMODFYPMT          = "ReqToModfyPmt"
)

// Description of ISO 20022 constants
// FININSTNID this is FinInstnId <FinInstnId>  FinancialInstitutionIdentification7 Based on the value in the Financial Institution Identification of Debtor Agent
// STTLMMTD this is SttlmMtd <SttlmMtd> SettlementMethod1Code Settlement Method - Method used to settle the (batch of) payment instructions.
// INSTGAGT this is InstgAgt <InstgAgt> BranchAndFinancialInstitutionIdentification4 Agent that instructs the next party in the chain to carry out the (set of) instruction(s)
// ORGNLGRPINF this is OrgnlGrpInf <OrgnlGrpInf> OriginalGroupInformation3 Original Group Information. Set of elements used to provide information on the original messsage.
// ORGNLINSTRID this is OrgnlInstrId <OrgnlInstrId> OriginalInstructionIdentification Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
// ORGNLENDTOENDID this is OrgnlEndToEndId <OrgnlEndToEndId> OriginalEndToEndIdentification Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction
// ORGNLTXID this is OrgnlTxId <OrgnlTxId> OriginalTransactionIdentification Unique identification, as assigned by the original first instructing agent, to umambiguously identify the transaction.
// ORGNLINTRBKSTTLMAMT this is OrgnlIntrBkSttlmAmt <OrgnlIntrBkSttlmAmt> ActiveOrHistoricCurrencyAndAmount Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
// RTRDINTRBKSTTLMAMT this is RtrdIntrBkSttlmAmt <RtrdIntrBkSttlmAmt> ActiveCurrencyAndAmount Returned Interbank Settlement Amount. Amount of money moved between the instructing agent and the instructed agent in the returned transaction.
// INTRBKSTTLMDT this is IntrBkSttlmDt <IntrBkSttlmDt> InterbankSettlementDate Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due
// RTRDINSTDAMT this is RtrdInstdAmt <RtrdInstdAmt> ActiveOrHistoricCurrencyAndAmount Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction
// CHRGBR this is ChrgBr <ChrgBr> ChargeBearerType1Code Specifies which party/parties will bear the charges associated with the processing of the payment transaction
// ORGNLTXREF this is OrgnlTxRef <OrgnlTxRef> OriginalTransactionReference13 Set of key elements used to identify the original transaction that is being referred to
// MSGDEFIDR this is MsgDefIdr <MsgDefIdr> MessageDefinitionIdentifier Contains the MessageIdentifier that defines the BusinessMessage. It must contain a MessageIdentifier published on the ISO 20022 website. AppHdr/MsgDefIdr
// INSTDAGT this is InstdAgt <InstdAgt> BranchAndFinancialInstitutionIdentification4 Instructed Agent. The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on
// INTRBKSTTLMAMT this is IntrBkSttlmAmt <IntrBkSttlmAmt> ActiveOrHistoricCurrencyAndAmount Interbank Settlement Amount. Amount of money moved between the instructing agent and the instructed agent.
// INSTDAMT this is InstdAmt <InstdAmt> ActiveOrHistoricCurrencyAndAmount Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
// XCHGRATE this is XchgRate <XchgRate> BaseOneRate Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
// DBTRAGT this is DbtrAgt <DbtrAgt> BranchAndFinancialInstitutionIdentification4 Debtor Agent. Financial institution servicing an account for the debtor.
// CDTRAGT this is CdtrAgt <CdtrAgt> BranchAndFinancialInstitutionIdentification4 Creditor Agent. Financial institution servicing an account for the creditor.
// ENVLP this is Envlp <Envlp> SupplementaryDataEnvelope1 Envelope. Technical element wrapping the supplementary data.
// SPLMTRYDATA this is SplmtryData <SplmtryData> SupplementaryData1 Additional information that can not be captured in the structured fields and/or any other specific block. Technical component that contains the validated supplementary data information. This technical envelope allows to segregate the supplementary data information from any other information.
// FITOFICSTMRCDTTRF this is FIToFICstmrCdtTrf <FIToFICstmrCdtTrf> FIToFICustomerCreditTransferV02 FITo FICustomer Credit Transfer. The FI2FI Customer Credit Transfer message is sent by the debtor's agent to the creditor's agent, directly or through other agents and/or a payment clearing and settlement system. It is used to move funds from a debtor's account to a creditor.
// STTLMACCT this is SttlmAcct <SttlmAcct> CashAccount16 Settlement Account.  A specific purpose account used to post debit and credit entries as a result of the transaction.
// INSTRID this is InstrId <InstrId> InstructionIdentification Instruction Identification. Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction
// ENDTOENDID this is EndToEndId <EndToEndId> EndToEndId End To End Identification. Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain
// TXID this is TxId <TxId> TransactionIdentification Transaction Identification. Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
// ORGNLGRPINFANDSTS this is OrgnlGrpInfAndSts <OrgnlGrpInfAndSts> OriginalGroupInformation20 Original Group Information And Status. Original group information concerning the group of transactions, to which the status report message refers to.
// JUSTFN this is Justfn <Justfn> CaseForwardingNotification3Code Justification. Justification for the forward action. Specifies the action requested when forwarding a case.
// ORGNLMSGNMID this is OrgnlMsgNmId <OrgnlMsgNmId> OriginalMessageNameIdentification Original Message Name Identification. Specifies the original message name identifier to which the message refers.
// ORGNLMSGID this is OrgnlMsgId <OrgnlMsgId> OriginalMessageIdentification "Original Message Identification. Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
// ORGNLGRPINFANDCXL this is OrgnlGrpInfAndCxl <OrgnlGrpInfAndCxl> OriginalGroupInformation23 Original Group Information And Cancellation. Set of elements used to provide information on the original messsage, to which the cancellation refers
// STTLMINF this is SttlmInf <SttlmInf> SettlementInformation13 Settlement Information. Specifies the details on how the settlement of the original transaction(s) between the instructing agent and the instructed agent was completed.
// ISSR this is Issr <Issr> Issuer Entity that assigns the credit reference type.
// ASSGNR this is Assgnr <Assgnr> Party7Choice Instructing Party. Party who assigns the case.
// BIZMSGIDR this is BizMsgIdr <BizMsgIdr> BusinessMessageIdentifier Business Message Identifier. Unambiguously identifies the Business Message to the MessagingEndpoint that has created the Business Message.
// CREDT this is CreDt <CreDt> ISONormalisedDateTime Date and time when this Business Message (header) was created. Note Times must be normalized, using the "Z" annotation.
// FITOFIPMTSTSRPT this is FIToFIPmtStsRpt <FIToFIPmtStsRpt> FIToFIPaymentStatusReportV03 FITo FIPayment Status Report. The PaymentInitiationStatusReport message is sent by an instructed agent to the payment initiator. It is used to inform this party about the positive or negative status of an instruction (either single, group or file). It is also used to report on a pending instruction.
// MSGID this is MsgId <MsgId> MessageIdentification Message Identification. Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message. Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
// CREDTTM this is CreDtTm <CreDtTm> ISODateTime Creation Date Time. Date and time at which the message was created.
// FIID this is FIId <FIId> FinancialInstitutionIdentification Financial Institution Identification. Set of elements used to uniquely and unambiguously identify a financial institution or a branch of a financial institution.
// BICFI this is BICFI <BICFI> BankIdentificationCode Business Identifier Code (BIC). BIC (Business Identifier Code) is an international standard for routing business transactions and identifying business parties.
// GRPHDR this is GrpHdr <GrpHdr> GroupHeader33 Group Header. Set of characteristics shared by all individual transactions included in the message.
// TXINFANDSTS this is TxInfAndSts <TxInfAndSts> PaymentTransactionInformation26 Transaction Information And Status. Information concerning the original transactions, to which the status report message refers
// TXSTS this is TxSts <TxSts> TransactionIndividualStatus3Code Transaction Status. Specifies the status of a transaction, in a coded form.
// CHRGSINF this is ChrgsInf <ChrgsInf> ChargesInformation5 Charges Information. Provides information on the charges related to the processing of the rejection of the instruction.
// PLCANDNM this is PlcAndNm <PlcAndNm> PlcAndNm Place And Name. Unambiguous reference to the location where the supplementary data must be inserted in the message instance.
// NBOFTXS this is NbOfTxs <NbOfTxs> Max15NumericText Number Of Transactions. Number of individual transactions contained in the message. For example <NbOfTxs>1</NbOfTxs>, <NbOfTxs>27</NbOfTxs>, <NbOfTxs>100</NbOfTxs>.
// ORGNLCREDTTM this is OrgnlCreDtTm <OrgnlCreDtTm> ISODateTime Original Creation Date Time. Original date and time at which the message was created.
// RTRID this is RtrId <RtrId> Max35Text Return Identification. Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the returned transaction. Usage: The instructing party is the party sending the return message and not the party that sent the original instruction that is being returned
// RTRRSNINF this is RtrRsnInf <RtrRsnInf> ReturnReasonInformation9 Return Reason Information. Set of elements used to provide detailed information on the return reason.
// ORGTR this is Orgtr <Orgtr> PartyIdentification32 Originator. Party that issues the status
// SVCLVL this is SvcLvl <SvcLvl> ServiceLevel8Choice Service Level. Agreement under which or rules under which the transaction should be processed.
// PMTTPINF this is PmtTpInf <PmtTpInf> PaymentTypeInformation21 Payment Type Information. Set of elements used to further specify the type of transaction.
// DBTR this is Dbtr <Dbtr> PartyIdentification32 Debtor. Party that owes an amount of money to the (ultimate) creditor.
// PSTLADR this is PstlAdr <PstlAdr> PostalAddress6 Postal Address. Information that locates and identifies a specific address, as defined by postal services.
// STRTNM this is StrtNm <StrtNm> Max70Text Street Name. Name of a street or thoroughfare.
// BLDGNB this is BldgNb <BldgNb> Max16Text Building Number. Number that identifies the position of a building on a street.
// PSTCD this is PstCd <PstCd> Max16Text Post Code. Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail
// TWNNM this is TwnNm <TwnNm> Max35Text Town Name. Name of a built-up area, with defined boundaries, and a local government.
// ADRLINE this is AdrLine <AdrLine> Max70Text Address Line. Information that locates and identifies a specific address, as defined by postal services, presented in free format text.
// CTRY this is Ctry <Ctry> CountryCode Country. Nation with its own government.
// FRMSGHEAD this is Fr <Fr> Party9Choice From. The sending MessagingEndpoint that has created this Business Message for the receiving MessagingEndpoint that will process this Business Message.
// TOMSGHEAD this is To <To> Party9Choice The MessagingEndpoint designated by the sending MessagingEndpoint to be the recipient who will ultimately process this Business Message.
// CDTTRFTXINF this is CdtTrfTxInf <CdtTrfTxInf> CreditTransferTransactionInformation11
// AGT this is Agt <Agt> BranchAndFinancialInstitutionIdentification5 Agent. Agent that takes the transaction charges or to which the transaction charges are due.
// UNDRLYG this is Undrlyg <Undrlyg> UnderlyingTransaction4Choice Underlying. Specifies the details of the underlying transaction on which the investigation is processed.
// MSSNGORINCRRCTINF this is MssngOrIncrrctInf <MssngOrIncrrctInf> MissingOrIncorrectInformation3 Missing Or Incorrect Information. Set of elements used to provide further information on the reason for the unable to apply investigation.
// UBLTOAPPLY this is UblToApply <UblToApply> UnableToApplyV07 The UnableToApply message is sent by a case creator or a case assigner to a case assignee. This message is used to initiate an investigation of a payment instruction that cannot be executed or reconciled.
// ID this is Id <Id> Max35Text Identification. Uniquely identifies the case.
// CONF this is Conf <Conf> ExternalInvestigationExecutionConfirmation1Co Confirmation. Specifies the status of the investigation, in a coded form.
// RSLTNOFINVSTGTN this is RsltnOfInvstgtn <RsltnOfInvstgtn> ResolutionOfInvestigationV08 Resolution Of Investigation. The ResolutionOfInvestigation message is sent by a case assignee to a case creator/case assigner. This message is used to inform of the resolution of a case, and optionally provides details about . - the corrective action undertaken by the case assignee.
// SIGNATUREMETHOD this is SignatureMethod <SignatureMethod> SignatureMethod The SignatureMethod property uses a string Uniform Resource Identifier (URI) to represents the <SignatureMethod> element of an XML digital signature. Signature Method: HMAC-SHA1, DSA with SHA1, RSA with SHA1, http://www.w3.org/2000/09 http://www.w3.org/2000/09/xmldsig#dsa-sha1/xmldsig#hmac-sha1 http://www.w3.org/2000/09/xmldsig#rsa-sha1
// CANONICALIZATIONMETHOD this is CanonicalizationMethod <CanonicalizationMethod> CanonicalizationMethod The CanonicalizationMethod property uses a string Uniform Resource Identifier (URI) to represent the <Canonicalization> element of an XML digital signature.
// NTFCTNOFCASEASSGNMT this is NtfctnOfCaseAssgnmt <NtfctnOfCaseAssgnmt> NotificationOfCaseAssignmentV05 Notification Of Case Assignment. The NotificationOfCaseAssignment message is sent by a case assignee to a case creator/case assigner.
// X509DATA this is X509Data <X509Data> KeyInfoX509Data Represents an <X509Data> subelement of an XMLDSIG or XML Encryption <KeyInfo> element.
// KEYINFO this is KeyInfo <KeyInfo> KeyInfo Represents an XML digital signature or XML encryption <KeyInfo> element.
// SIGNATURE this is Signature <Signature> Signature Represents the <Signature> element of an XML signature.
// CRETR this is Cretr <Cretr> Party35Choice Creator. Party that created the investigation case.
// ASSGNE this is Assgne <Assgne> Party35Choice Assignee. Party to which the case is assigned.
// CXLRSNINF this is CxlRsnInf <CxlRsnInf> CancellationReasonInformation3 Cancellation Reason Information. Set of elements used to provide detailed information on the cancellation reason.
// FITOFIPMTCXLREQ this is FIToFIPmtCxlReq <FIToFIPmtCxlReq> FIToFIPaymentCancellationRequestV01 FITo FIPayment Cancellation Request. This message allows initiating an investigation case when a payment transaction needs to be cancelled.
// ORGNLINSTDAMT this is OrgnlInstdAmt <OrgnlInstdAmt> OriginalInstructedAmount Original Instructed Amount. Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
// ULTMTCDTR this is UltmtCdtr <UltmtCdtr> PartyIdentification32 Ultimate Creditor. Ultimate party to which an amount of money is due.
// MOD this is Mod <Mod> RequestedModification7 Modification. Identifies the list of modifications requested.
// REQTOMODFYPMT this is ReqToModfyPmt <ReqToModfyPmt> RequestToModifyPaymentV05 Request To Modify Payment. The RequestToModifyPayment message is sent by a case creator/case assigner to a case assignee. This message is used to request the modification of characteristics of an original payment instruction.
