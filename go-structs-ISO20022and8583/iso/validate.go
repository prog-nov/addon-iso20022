package iso

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

import (
	"encoding/xml"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/figassis/bankiso/iso20022/acmt"
	"github.com/figassis/bankiso/iso20022/auth"
	"github.com/figassis/bankiso/iso20022/caaa"
	"github.com/figassis/bankiso/iso20022/caam"
	"github.com/figassis/bankiso/iso20022/cain"
	"github.com/figassis/bankiso/iso20022/camt"
	"github.com/figassis/bankiso/iso20022/catm"
	"github.com/figassis/bankiso/iso20022/catp"
	"github.com/figassis/bankiso/iso20022/head"
	"github.com/figassis/bankiso/iso20022/pacs"
	"github.com/figassis/bankiso/iso20022/pain"
	"github.com/figassis/bankiso/iso20022/remt"
	//"github.com/figassis/bankiso/iso20022/secl"
	//"github.com/figassis/bankiso/iso20022/semt"
	//"github.com/figassis/bankiso/iso20022/sese"
	//"github.com/figassis/bankiso/iso20022/admi"
	//"github.com/figassis/bankiso/iso20022/colr"
	//"github.com/figassis/bankiso/iso20022/fxtr"
	//"github.com/figassis/bankiso/iso20022/reda"
	//"github.com/figassis/bankiso/iso20022/seev"
	//"github.com/figassis/bankiso/iso20022/setr"
	//"github.com/figassis/bankiso/iso20022/tsin"
	//"github.com/figassis/bankiso/iso20022/tsmt"
	//"github.com/figassis/bankiso/iso20022/tsrv"
)

var ISO20022Registry map[string]interface{} = map[string]interface{}{
	"acmt.007.001.02": &acmt.Document00700102{},
	"acmt.008.001.02": &acmt.Document00800102{},
	"acmt.009.001.02": &acmt.Document00900102{},
	"acmt.010.001.02": &acmt.Document01000102{},
	"acmt.011.001.02": &acmt.Document01100102{},
	"acmt.012.001.02": &acmt.Document01200102{},
	"acmt.013.001.02": &acmt.Document01300102{},
	"acmt.014.001.02": &acmt.Document01400102{},
	"acmt.015.001.02": &acmt.Document01500102{},
	"acmt.016.001.02": &acmt.Document01600102{},
	"acmt.017.001.02": &acmt.Document01700102{},
	"acmt.018.001.02": &acmt.Document01800102{},
	"acmt.019.001.02": &acmt.Document01900102{},
	"acmt.020.001.02": &acmt.Document02000102{},
	"acmt.021.001.02": &acmt.Document02100102{},
	"acmt.022.001.02": &acmt.Document02200102{},
	"acmt.023.001.02": &acmt.Document02300102{},
	"acmt.024.001.02": &acmt.Document02400102{},
	"auth.001.001.01": &auth.Document00100101{},
	"auth.002.001.01": &auth.Document00200101{},
	"auth.003.001.01": &auth.Document00300101{},
	"auth.018.001.01": &auth.Document01800101{},
	"auth.019.001.01": &auth.Document01900101{},
	"auth.020.001.01": &auth.Document02000101{},
	"auth.021.001.01": &auth.Document02100101{},
	"auth.022.001.01": &auth.Document02200101{},
	"auth.023.001.01": &auth.Document02300101{},
	"auth.024.001.01": &auth.Document02400101{},
	"auth.025.001.01": &auth.Document02500101{},
	"auth.026.001.01": &auth.Document02600101{},
	"auth.027.001.01": &auth.Document02700101{},
	"caaa.001.001.05": &caaa.Document00100105{},
	"caaa.002.001.05": &caaa.Document00200105{},
	"caaa.003.001.05": &caaa.Document00300105{},
	"caaa.004.001.05": &caaa.Document00400105{},
	"caaa.005.001.05": &caaa.Document00500105{},
	"caaa.006.001.05": &caaa.Document00600105{},
	"caaa.007.001.05": &caaa.Document00700105{},
	"caaa.008.001.05": &caaa.Document00800105{},
	"caaa.009.001.05": &caaa.Document00900105{},
	"caaa.010.001.05": &caaa.Document01000105{},
	"caaa.011.001.05": &caaa.Document01100105{},
	"caaa.012.001.05": &caaa.Document01200105{},
	"caaa.013.001.05": &caaa.Document01300105{},
	"caaa.014.001.05": &caaa.Document01400105{},
	"caaa.015.001.05": &caaa.Document01500105{},
	"caaa.016.001.03": &caaa.Document01600103{},
	"caaa.017.001.03": &caaa.Document01700103{},
	"caam.001.001.02": &caam.Document00100102{},
	"caam.002.001.02": &caam.Document00200102{},
	"caam.003.001.02": &caam.Document00300102{},
	"caam.004.001.02": &caam.Document00400102{},
	"caam.005.001.02": &caam.Document00500102{},
	"caam.006.001.02": &caam.Document00600102{},
	"caam.007.001.01": &caam.Document00700101{},
	"caam.008.001.01": &caam.Document00800101{},
	"caam.009.001.02": &caam.Document00900102{},
	"caam.010.001.02": &caam.Document01000102{},
	"caam.011.001.01": &caam.Document01100101{},
	"caam.012.001.01": &caam.Document01200101{},
	"cain.001.001.01": &cain.Document00100101{},
	"cain.002.001.01": &cain.Document00200101{},
	"cain.003.001.01": &cain.Document00300101{},
	"cain.004.001.01": &cain.Document00400101{},
	"cain.005.001.01": &cain.Document00500101{},
	"cain.006.001.01": &cain.Document00600101{},
	"cain.007.001.01": &cain.Document00700101{},
	"cain.008.001.01": &cain.Document00800101{},
	"cain.009.001.01": &cain.Document00900101{},
	"cain.010.001.01": &cain.Document01000101{},
	"cain.011.001.01": &cain.Document01100101{},
	"cain.012.001.01": &cain.Document01200101{},
	"cain.013.001.01": &cain.Document01300101{},
	"camt.026.001.04": &camt.Document02600104{},
	"camt.027.001.04": &camt.Document02700104{},
	"camt.028.001.06": &camt.Document02800106{},
	"camt.029.001.06": &camt.Document02900106{},
	"camt.030.001.04": &camt.Document03000104{},
	"camt.031.001.04": &camt.Document03100104{},
	"camt.032.001.03": &camt.Document03200103{},
	"camt.033.001.04": &camt.Document03300104{},
	"camt.034.001.04": &camt.Document03400104{},
	"camt.035.001.03": &camt.Document03500103{},
	"camt.036.001.03": &camt.Document03600103{},
	"camt.037.001.04": &camt.Document03700104{},
	"camt.038.001.03": &camt.Document03800103{},
	"camt.039.001.04": &camt.Document03900104{},
	"camt.052.001.06": &camt.Document05200106{},
	"camt.053.001.06": &camt.Document05300106{},
	"camt.054.001.06": &camt.Document05400106{},
	"camt.055.001.05": &camt.Document05500105{},
	"camt.056.001.05": &camt.Document05600105{},
	"camt.057.001.05": &camt.Document05700105{},
	"camt.058.001.05": &camt.Document05800105{},
	"camt.059.001.05": &camt.Document05900105{},
	"camt.060.001.03": &camt.Document06000103{},
	"camt.086.001.02": &camt.Document08600102{},
	//"camt.087.001.03": &camt.Document08700103{}, //had a problem generating message definition. Got amigous message definition RequestToModifyPaymentV03
	"catm.001.001.05": &catm.Document00100105{},
	"catm.002.001.05": &catm.Document00200105{},
	"catm.003.001.05": &catm.Document00300105{},
	"catm.004.001.04": &catm.Document00400104{},
	"catm.005.001.02": &catm.Document00500102{},
	"catm.006.001.02": &catm.Document00600102{},
	"catm.007.001.01": &catm.Document00700101{},
	"catm.008.001.01": &catm.Document00800101{},
	"catp.001.001.02": &catp.Document00100102{},
	"catp.002.001.02": &catp.Document00200102{},
	"catp.003.001.02": &catp.Document00300102{},
	"catp.004.001.02": &catp.Document00400102{},
	"catp.005.001.02": &catp.Document00500102{},
	"catp.006.001.02": &catp.Document00600102{},
	"catp.007.001.02": &catp.Document00700102{},
	"catp.008.001.02": &catp.Document00800102{},
	"catp.009.001.02": &catp.Document00900102{},
	"catp.010.001.02": &catp.Document01000102{},
	"catp.011.001.02": &catp.Document01100102{},
	"catp.012.001.01": &catp.Document01200101{},
	"catp.013.001.01": &catp.Document01300101{},
	"catp.014.001.01": &catp.Document01400101{},
	"catp.015.001.01": &catp.Document01500101{},
	"catp.016.001.01": &catp.Document01600101{},
	"catp.017.001.01": &catp.Document01700101{},
	"head.001.001.01": &head.Document00100101{},
	"pacs.002.001.07": &pacs.Document00200107{},
	"pacs.003.001.06": &pacs.Document00300106{},
	"pacs.004.001.06": &pacs.Document00400106{},
	"pacs.007.001.06": &pacs.Document00700106{},
	"pacs.008.001.06": &pacs.Document00800106{},
	"pacs.009.001.06": &pacs.Document00900106{},
	"pacs.010.001.02": &pacs.Document01000102{},
	"pain.001.001.07": &pain.Document00100107{},
	"pain.002.001.07": &pain.Document00200107{},
	"pain.007.001.06": &pain.Document00700106{},
	"pain.008.001.06": &pain.Document00800106{},
	"pain.009.001.04": &pain.Document00900104{},
	"pain.010.001.04": &pain.Document01000104{},
	"pain.011.001.04": &pain.Document01100104{},
	"pain.012.001.04": &pain.Document01200104{},
	"pain.013.001.05": &pain.Document01300105{},
	"pain.014.001.05": &pain.Document01400105{},
	"remt.001.001.02": &remt.Document00100102{},
	"remt.002.001.01": &remt.Document00200101{},
}

func makeISO20022(code, message string) (result Message, err error) {

	val, ok := ISO20022Registry[code]
	if !ok {
		err = fmt.Errorf("Invalid ISO20022 code %s", code)
		return
	}

	if err = xml.Unmarshal([]byte(message), &val); err != nil {
		return
	}

	if result, ok = val.(Message); !ok {
		err = errors.New("Invalid ISO20022 message")
	}
	return
}

//Basic XML version 1.0 validation via header regex
//Eg. if XML, process as iso20022, else, check if it's an iso8583 message
func isXML(message string) bool {
	//Check if string is valid XML
	header1 := regexp.QuoteMeta("<?xml version='1.0' encoding='UTF-8'?>")
	header2 := regexp.QuoteMeta("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	valid, err := regexp.MatchString("^("+header1+"|"+header2+")", message)

	if err != nil {
		fmt.Println("xml.isXML: Could not process file. ", err)
		return false
	}

	if valid {
		//@TODO: check if XML is well formed
		return true
	}

	return false
}

//Returns a boolean and the message type, eg. true, "acmt"
//@TODO: implement full xsd schema validation
func ValidateISO20022(message string) (domain, code string, err error) {

	isXml := isXML(message)
	if !isXml {
		return "", "", errors.New("message is not valid XML")
	}

	v := Document{Format: ""}

	//@TODO To improve performance, consider using a string split or regex match instead
	if err = xml.Unmarshal([]byte(message), &v); err != nil {
		return
	}

	//Match all iso20022 message types
	search := `^urn:iso:std:iso:20022:tech:xsd:(acmt|admi|auth|caaa|caam|camt|catm|catp|colr|fxtr|pacs|pain|reda|remt|secl|seev|semt|sese|setr|tsin|tsmt|tsrv|head)\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}$`
	match, err := regexp.MatchString(search, v.Format)
	if err != nil {
		return
	}

	if !match {
		return "", "", errors.New("Input is not a valid iso20022 message")
	}

	//Eg. pain.009.001.04
	splitCode := strings.Split(v.Format, "xsd:")
	if len(splitCode) < 2 {
		return "", "", fmt.Errorf("%s is not a valid iso20022 message format", v.Format)
	}
	code = splitCode[1]

	splitCode = strings.Split(code, ".")
	if len(splitCode) != 4 {
		return "", "", fmt.Errorf("%s is not a valid iso20022 message code", code)
	}

	//Eg. pain
	domain = splitCode[0]

	return
}

//Return ISO Message struct containing message content as well as
//identifiers to facilitate type casting/assertion and message handling
func Decode(message string) (result Message, domain, code string, err error) {

	if domain, code, err = ValidateISO20022(message); err != nil {
		return
	}

	if result, err = makeISO20022(code, message); err != nil {
		return
	}

	return
}
