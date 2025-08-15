package common

import (
	"time"

	"github.com/moov-io/fednow20022/pkg/fednow"
)

type CAMTReportType string
type ReportType string
type TransactionStatusCode string
type InstrumentPropCodeType string
type CdtDbtInd string
type ReportStatus string
type WorkingDayType string
type PaymentSystemType string
type BalanceType string
type CreditLineType string
type TransactionCode string
type ProprietaryAgentType string
type ModifyMode string
type AccountType string
type SettlementMethod string
type ExternalCashClearingSystem string
type LocalInstrumentType string
type ChargeBearerType string
type CodeOrProprietaryType string
type RemittanceDeliveryMethod string
type ProxyType string
type FundEventType string
type StatusReasonInformationCode string
type CopyDuplicateCode string
type ContactMethod string
type MissingOrIncorrectInformationCode string
type InvestigationExecutionConfirmationCode string
type ReturnReasonCode string
type MessageHandlingCode string
type PaymentMethodCode string
type CancellationRequestResponse string
type ReasonCodeType string
type InvestigationExecutionConfirmCode string

const (
	BusinessProcessingDate WorkingDayType = "BPRD"
)
const (
	Respondent ProprietaryAgentType = "RESP"
)
const (
	FedNowServiceLI LocalInstrumentType = "FDNA" // FedNow Service Local Instrument
)
const (
	FedNowService ExternalCashClearingSystem = "FDN" // FedNow Service
)
const (
	EveryDay ReportType = "EDAY"
	Intraday ReportType = "IDAY"
)
const (
	Credit CdtDbtInd = "CRDT"
	Debit  CdtDbtInd = "DBIT"
)
const (
	Cheque         PaymentMethodCode = "CHK"
	CreditTransfer PaymentMethodCode = "TRF"
)
const (
	MasterAccount       AccountType = "M"
	SingleRoutingNumber AccountType = "S"
)
const (
	TelephoneNumber ProxyType = "TELN" // Telephone Number
	EmailAddress    ProxyType = "EMAL" // Email Address
)
const (
	FedNowProcessingSuccessful        MessageHandlingCode = "TS01"
	ParticipantReceiptAcknowledgement MessageHandlingCode = "TS02"
)
const (
	CancelledAsPerRequest       CancellationRequestResponse = "CNCL" // Cancelled As Per Request
	CancellationRequestRejected CancellationRequestResponse = "RJCR" // Cancellation Request Rejected
	CancellationRequestPending  CancellationRequestResponse = "PDCR" // Cancellation Request Pending
)
const (
	ClearingSystem   SettlementMethod = "CLRG"
	CoverMethod      SettlementMethod = "COVE"
	InstructedAgent  SettlementMethod = "INDA"
	InstructingAgent SettlementMethod = "INGA"
)
const (
	ChargeBearerSLEV   ChargeBearerType = "SLEV" // Sender Pays All Charges
	ChargeBearerRECV   ChargeBearerType = "RECV" // Receiver Pays All Charges
	ChargeBearerSHAR   ChargeBearerType = "SHAR" // Shared Charges
	ChargeBearerDEBT   ChargeBearerType = "DEBT" // Shared Charges
	ChargeBearerCREDIT ChargeBearerType = "CRED" // Shared Charges
)
const (
	ReturnRequestAccepted          InvestigationExecutionConfirmCode = "IPAY"
	ReturnRequestRejected          InvestigationExecutionConfirmCode = "RJCR"
	ReturnRequestPending           InvestigationExecutionConfirmCode = "PDCR"
	PartiallyExecutedReturnRequest InvestigationExecutionConfirmCode = "PECR"
)
const (
	ContactEmail             ContactMethod = "EMAIL"
	ContactPhone             ContactMethod = "PHON"
	ContactMail              ContactMethod = "MAIL"
	ContactFax               ContactMethod = "FAXX" // Fax
	ContactLetter            ContactMethod = "LETT"
	ContactMobileOrCellPhone ContactMethod = "CELL"
)
const (
	IDUP InvestigationExecutionConfirmationCode = "IDUP" // Instruction Is Duplicate - It is recommended to provide a reference of the original message in the Resolution Related Information component
	INFO InvestigationExecutionConfirmationCode = "INFO" // Additional Information Sent – It is recommended to provide the Case Identification of the Information Request in the Case Identification of the Additional Payment Information message (camt.028)
	IPAY InvestigationExecutionConfirmationCode = "IPAY" // Payment Initiated – Correction Transaction component may be present with reference information of the payment instruction
	NINF InvestigationExecutionConfirmationCode = "NINF" // No Further Information Available
	PDNG InvestigationExecutionConfirmationCode = "PDNG" // Information Request Pending
)
const (
	ClosedAccountNumber           ReasonCodeType = "AC04" // Closed Account Number
	AwaitingDebitAuthority        ReasonCodeType = "ADAC" // Awaiting Debit Authority
	AgentDecision                 ReasonCodeType = "AGNT" // Agent Decision
	TransInsufficientFunds        ReasonCodeType = "AM04" // Insufficient Funds
	AlreadyReturned               ReasonCodeType = "ARDT" // Already Returned - It is recommended to provide the reference of the return payment
	AwaitingReply                 ReasonCodeType = "ARPL" // Awaiting Reply
	CustomerDecision              ReasonCodeType = "CUST" // Customer Decision
	LegalDecision                 ReasonCodeType = "LEGL" // Legal Decision
	TransNarrative                ReasonCodeType = "NARR" // Narrative - Must be followed by the reason in free-formatted text in Additional Information
	NoAnswerFromCustomer          ReasonCodeType = "NOAS" // No Answer From Customer
	NoOriginalTransactionReceived ReasonCodeType = "NOOR" // No Original Transaction Received
	PassedToNextAgent             ReasonCodeType = "PTNA" // Passed To Next Agent
	RequestingDebitAuthority      ReasonCodeType = "RQDA" // Requesting Debit Authority
)
const (
	InvalidOrMissingCreditorAccountType ReturnReasonCode = "AC14"
	DuplicatePaymentCode                ReturnReasonCode = "DUPL" // Duplicate Payment - It is recommended to provide the reference of the original message in Additional Information
	TooLowAmount                        ReturnReasonCode = "AM06" // Too Low Amount - It is suggested, if possible, to provide the expected amount in Additional Information
	WrongAmount                         ReturnReasonCode = "AM09" // Wrong Amount - It is suggested, if possible, to provide the expected amount in Additional Information
	UnrecognizedInitiatingParty         ReturnReasonCode = "BE05" // Unrecognized Initiating Party
	RequestedByCustomer                 ReturnReasonCode = "CUST" // Requested By Customer
	FollowingRefundRequest              ReturnReasonCode = "FOCR" // Following Refund Request - It is recommended to provide the reference of the Refund Request in Additional Information
	FraudulentPayment                   ReturnReasonCode = "FR01" // Fraudulent Payment
	NotSpecifiedReasonCustomer          ReturnReasonCode = "MS02" // Not Specified Reason Customer
	NotSpecifiedReasonAgent             ReturnReasonCode = "MS03" // Not Specified Reason Agent
	NarrativeReasonCode                 ReturnReasonCode = "NARR" // Narrative - Must be followed by the reason in free-formatted text in Additional Information
	RegulatoryReasonCode                ReturnReasonCode = "RR04" // Regulatory Reason
	ReturnUponUnableToApply             ReturnReasonCode = "RUTA" // Return Upon Unable To Apply - It is recommended to provide the reference of the investigation case in Additional Information
	UnduePayment                        ReturnReasonCode = "UPAY" // Undue Payment
)
const (
	MissingAmount                                   MissingOrIncorrectInformationCode = "MS06"
	MissingCreditor                                 MissingOrIncorrectInformationCode = "MS12"
	MissingCreditorAccount                          MissingOrIncorrectInformationCode = "MS13"
	MissingCreditorAgent                            MissingOrIncorrectInformationCode = "MS15"
	MissingDebtor                                   MissingOrIncorrectInformationCode = "MS03"
	MissingDebtorAccount                            MissingOrIncorrectInformationCode = "MS04"
	MissingDebtorAgent                              MissingOrIncorrectInformationCode = "MS05"
	MissingInstructedReimbursementAgent             MissingOrIncorrectInformationCode = "MS10"
	MissingInstructingReimbursementAgent            MissingOrIncorrectInformationCode = "MS09"
	MissingInstruction                              MissingOrIncorrectInformationCode = "MS14"
	MissingInstructionForCreditorAgent              MissingOrIncorrectInformationCode = "MS16"
	MissingInstructionForDebtorAgent                MissingOrIncorrectInformationCode = "MS17"
	MissingInstructionForNextAgent                  MissingOrIncorrectInformationCode = "MS02"
	MissingIntermediary                             MissingOrIncorrectInformationCode = "MS08"
	MissingRemittanceInformation                    MissingOrIncorrectInformationCode = "MS01"
	MissingSettlementAccount                        MissingOrIncorrectInformationCode = "MS07"
	MissingThirdReimbursementAgent                  MissingOrIncorrectInformationCode = "MS11"
	Narrative                                       MissingOrIncorrectInformationCode = "NARR"
	IncorrectCategoryPurpose                        MissingOrIncorrectInformationCode = "IN03"
	IncorrectChargeBearer                           MissingOrIncorrectInformationCode = "IN17"
	IncorrectCreditor                               MissingOrIncorrectInformationCode = "IN13"
	IncorrectCreditorAccount                        MissingOrIncorrectInformationCode = "IN14"
	IncorrectCreditorAddress                        MissingOrIncorrectInformationCode = "IN39"
	IncorrectCreditorAgentAccount                   MissingOrIncorrectInformationCode = "IN12"
	IncorrectDebtor                                 MissingOrIncorrectInformationCode = "IN07"
	IncorrectDebtorAccount                          MissingOrIncorrectInformationCode = "IN08"
	IncorrectDebtorAddress                          MissingOrIncorrectInformationCode = "IN38"
	IncorrectInstructedReimbursementAgent           MissingOrIncorrectInformationCode = "IN09"
	IncorrectInstructingReimbursementAgent          MissingOrIncorrectInformationCode = "IN37"
	IncorrectInstructionForCreditorAgent            MissingOrIncorrectInformationCode = "IN19"
	IncorrectInstructionForNextAgent                MissingOrIncorrectInformationCode = "IN18"
	IncorrectInterbankSettlementAmount              MissingOrIncorrectInformationCode = "IN06"
	IncorrectInterbankSettlementDate                MissingOrIncorrectInformationCode = "IN05"
	IncorrectPaymentClearingChannel                 MissingOrIncorrectInformationCode = "IN11"
	IncorrectPaymentPurpose                         MissingOrIncorrectInformationCode = "IN16"
	IncorrectPaymentServiceLevel                    MissingOrIncorrectInformationCode = "IN02"
	IncorrectRelatedReference                       MissingOrIncorrectInformationCode = "IN01"
	IncorrectRemittanceInformation                  MissingOrIncorrectInformationCode = "IN15"
	IncorrectReportEntryAmount                      MissingOrIncorrectInformationCode = "IN33"
	IncorrectRequestedExecutionDate                 MissingOrIncorrectInformationCode = "IN04"
	IncorrectSettlementAccount                      MissingOrIncorrectInformationCode = "IN36"
	IncorrectThirdReimbursementAgent                MissingOrIncorrectInformationCode = "IN10"
	InsufficientChargesDetails                      MissingOrIncorrectInformationCode = "MM34"
	InsufficientPaymentPurpose                      MissingOrIncorrectInformationCode = "MM35"
	MismatchCreditorAgentNameAccount                MissingOrIncorrectInformationCode = "MM22"
	MismatchCreditorNameAccount                     MissingOrIncorrectInformationCode = "MM20"
	MismatchDebtorNameAccount                       MissingOrIncorrectInformationCode = "MM21"
	PaymentExecutedCreditorAccountOrIdentification  MissingOrIncorrectInformationCode = "MM31"
	PaymentExecutedCreditorNameOrAddress            MissingOrIncorrectInformationCode = "MM32"
	PaymentExecutedDebtorAccountOrIdentification    MissingOrIncorrectInformationCode = "MM27"
	PaymentExecutedDebtorNameOrAddress              MissingOrIncorrectInformationCode = "MM28"
	PendingExecutionCreditorAccountOrIdentification MissingOrIncorrectInformationCode = "MM29"
	PendingExecutionCreditorNameOrAddress           MissingOrIncorrectInformationCode = "MM30"
	PendingExecutionDebtorAccountOrIdentification   MissingOrIncorrectInformationCode = "MM25"
	PendingExecutionDebtorNameOrAddress             MissingOrIncorrectInformationCode = "MM26"
)
const (
	//Ad hoc Fedwire Funds Service customized message.
	AdHoc                     FundEventType = "ADHC"
	ConnectionCheck           FundEventType = "PING"
	ConnectionPointDisconnect FundEventType = "FPCD"
	ConnectionPointReconnect  FundEventType = "FPCR"
	FedNowKeys                FundEventType = "FNKY"
	ParticipantOffline        FundEventType = "FPOF"
	ParticipantOnline         FundEventType = "FPON"
	ParticipantReactivation   FundEventType = "FPRE"
	SystemClosed              FundEventType = "CLSD"
	SystemExtension           FundEventType = "EXTN"
	SystemOpen                FundEventType = "OPEN"
	SystemRollover            FundEventType = "ROLL"
	TransactionLimitChange    FundEventType = "SITL"
)
const (
	AccountDebitCreditNotification                   CAMTReportType = "ADCN"
	AccountActivityDetailsReport                     CAMTReportType = "AADR"
	AccountActivityTotalsReport                      CAMTReportType = "AATR"
	AccountBalanceReport                             CAMTReportType = "ABAR"
	CorrespondentAccountActivityDetailsReport        CAMTReportType = "CADR"
	CorrespondentAccountActivityTotalsReport         CAMTReportType = "CATR"
	CorrespondentIntradayAccountActivityTotalsReport CAMTReportType = "CITR"
	IntradayAccountActivityTotalsReport              CAMTReportType = "IATR"
)
const (
	Copy          CopyDuplicateCode = "COPY"
	CopyDuplicate CopyDuplicateCode = "CODU"
	Duplicate     CopyDuplicateCode = "DUPL"
)
const (
	InsufficientFunds         StatusReasonInformationCode = "AM04" // The account does not have enough balance to process the transaction.
	DuplicateTransaction      StatusReasonInformationCode = "AM05" // The transaction appears to be a duplicate of another payment.
	WrongAccount              StatusReasonInformationCode = "AM09" // The account number provided is incorrect or does not exist.
	CreditorBankNotRegistered StatusReasonInformationCode = "CNOR" // The creditor’s bank is unknown or not part of the payment system.
	DebtorBankNotRegistered   StatusReasonInformationCode = "DNOR" // The debtor’s bank is unknown or not part of the payment system.
	InvalidFileFormat         StatusReasonInformationCode = "FF01" // The file submitted does not match the expected format.
	InvalidBIC                StatusReasonInformationCode = "RC01" // The Bank Identifier Code (BIC) provided is incorrect.
	MissingDebtorInfo         StatusReasonInformationCode = "RR01" // Mandatory information about the debtor is missing.
	MissingCreditorInfo       StatusReasonInformationCode = "RR02" // Mandatory information about the creditor is missing.
	CutOffTimeExceeded        StatusReasonInformationCode = "SL01" // The transaction was submitted after the cut-off time for processing.

	// Added values
	InvalidCreditorAccount            StatusReasonInformationCode = "AC03" // Invalid Creditor Account
	BlockedAccount                    StatusReasonInformationCode = "AC06" // Blocked Account
	ClosedCreditorAccount             StatusReasonInformationCode = "AC07" // Closed Creditor Account
	InconsistentEndCustomer           StatusReasonInformationCode = "BE01" // Inconsistent End Customer
	MissingOrIncorrectCreditorAddress StatusReasonInformationCode = "BE04" // Missing or Incorrect Creditor Address
	InvalidCreditorIdentification     StatusReasonInformationCode = "BE17" // Invalid Creditor Identification
	DuplicatePayment                  StatusReasonInformationCode = "DUPL" // Duplicate Payment - It is recommended to provide the reference of the original message in Additional Information
	NarrativeReason                   StatusReasonInformationCode = "NARR" // Narrative - Must be followed by the reason in free-formatted text in Additional Information
	RegulatoryReason                  StatusReasonInformationCode = "RR04" // Regulatory Reason
)
const (
	MessagesInProcess           TransactionStatusCode = "INPR"
	MessagesIntercepted         TransactionStatusCode = "ICPT"
	AcceptedTechnicalValidation TransactionStatusCode = "ACTC"
	AcceptedSettlementInProcess TransactionStatusCode = "ACSP"
	AcceptedWithChange          TransactionStatusCode = "ACWC"
	AcceptedCreditClearing      TransactionStatusCode = "ACCC"
	Sent                        TransactionStatusCode = "SENT"
	TransReceived               TransactionStatusCode = "RCVD"
	Rejected                    TransactionStatusCode = "RJCT"
	TransPending                TransactionStatusCode = "PDNG"
	Cancelled                   TransactionStatusCode = "CANC"
	AcceptedCustomerProfile     TransactionStatusCode = "ACCP"
	PartiallyAccepted           TransactionStatusCode = "PART"
	TransCredit                 TransactionStatusCode = "CRDT"
	TransDebit                  TransactionStatusCode = "DBIT"
	AcceptedSettlementCompleted TransactionStatusCode = "ACSC"

	NonvalueMessagesReceived TransactionStatusCode = "RCVD"
	NonvalueMessagesSent     TransactionStatusCode = "SENT"
	RejectedMessagesReceived TransactionStatusCode = "RJTR"
	RejectedMessagesSent     TransactionStatusCode = "RJTS"
	ValueMessagesReceived    TransactionStatusCode = "CRDT"
	ValueMessagesSent        TransactionStatusCode = "DBIT"

	Presented TransactionStatusCode = "PRES" // The transaction has been presented for settlement.
)
const (
	InstrumentCTRC                      InstrumentPropCodeType = "CTRC" // Credit Transfer (Proprietary Code)
	InstrumentDD                        InstrumentPropCodeType = "DD"   // Direct Debit
	InstrumentStraightThroughProcessing InstrumentPropCodeType = "STP"  // Straight Through Processing
	InstrumentNCT                       InstrumentPropCodeType = "NCT"  // National Credit Transfer
	InstrumentCTRD                      InstrumentPropCodeType = "CTRD" // National Credit Transfer
)
const (
	Book     ReportStatus = "BOOK"
	Pending  ReportStatus = "PDNG"
	Received ReportStatus = "RCVD"
	Settled  ReportStatus = "SETT"
)
const (
	PaymentSysUSABA PaymentSystemType = "USABA" // American Bankers Association (ABA) routing number system
	PaymentSysCHIPS PaymentSystemType = "CHIPS" // Clearing House Interbank Payments System
	PaymentSysSEPA  PaymentSystemType = "SEPA"  // Single Euro Payments Area
	PaymentSysRTGS  PaymentSystemType = "RTGS"  // Real-Time Gross Settlement
	PaymentSysSWIFT PaymentSystemType = "SWIFT" // Society for Worldwide Interbank Financial Telecommunication
	PaymentSysBACS  PaymentSystemType = "BACS"  // Bankers' Automated Clearing Services
)
const (
	AccountBalance                        BalanceType = "ABAL"
	AvailableBalanceFromAccountBalance    BalanceType = "AVAL"
	AvailableBalanceFromDaylightOverdraft BalanceType = "AVLD"
	DaylightOverdraftBalance              BalanceType = "DLOD"
	OpeningBalanceFinalBalanceLoaded      BalanceType = "OBFL"
	OpeningBalanceNotLoaded               BalanceType = "OBNL"
	OpeningBalancePriorDayBalanceLoaded   BalanceType = "OBPL"
)
const (
	CollateralAvailable                CreditLineType = "COLL"
	CollateralizedCapacity             CreditLineType = "CCAP"
	CollateralizedDaylightOverdrafts   CreditLineType = "CLOD"
	NetDebitCap                        CreditLineType = "NCAP"
	UncollateralizedDaylightOverdrafts CreditLineType = "ULOD"
)
const (
	AvailableAllOtherActivity        TransactionCode = "AVOT"
	FedNowFundsTransfers             TransactionCode = "FDNF"
	FedwireFundsTransfers            TransactionCode = "FDWF"
	FedwireSecuritiesTransfers       TransactionCode = "FDWS"
	MemoPostEntries                  TransactionCode = "MEMO"
	NationalSettlementServiceEntries TransactionCode = "NSSE"
	PrefundedACHCreditItems          TransactionCode = "FDAP"
	UnavailableAllOtherActivity      TransactionCode = "UVOT"
)
const (
	CodeCINV CodeOrProprietaryType = "CINV" // Invoice
	CodeCREQ CodeOrProprietaryType = "CREQ" // Credit Request
	CodeCNTR CodeOrProprietaryType = "CNTR" // Credit Note
	CodeDBTR CodeOrProprietaryType = "DBTR" // Debtor
	CodeCRED CodeOrProprietaryType = "CRED" // Credit
	CodeSCT  CodeOrProprietaryType = "SCT"  // SEPA Credit Transfer
	CodePAYM CodeOrProprietaryType = "PAYM" // Payment Message
	CodeRTGS CodeOrProprietaryType = "RTGS" // Real-Time Gross Settlement
	CodeRCLS CodeOrProprietaryType = "RCLS" // Reversal
	CodeRFF  CodeOrProprietaryType = "RFF"  // Reference
	CodeCMCN CodeOrProprietaryType = "CMCN" // Reference
)
const (
	Fax                       RemittanceDeliveryMethod = "FAXI" //Fax
	ElectronicDataInterchange RemittanceDeliveryMethod = "EDIC" //Electronic Data Interchange (EDI)
	UniformResourceIdentifier RemittanceDeliveryMethod = "URID" //Uniform Resource Identifier (URI)
	PostalMail                RemittanceDeliveryMethod = "POST" //Postal Mail
	Email                     RemittanceDeliveryMethod = "EMAL" //Email
)
const (
	All      ModifyMode = "ALLL"
	Changed  ModifyMode = "CHNG"
	Modified ModifyMode = "MODF"
)

type MessagePagenation struct {
	// PgNb (Page Number) indicates the current page of the report.
	// It is used for paginated messages where multiple pages exist.
	PageNumber string `json:"page_number,omitempty"`
	// LastPgInd (Last Page Indicator) specifies whether this is the last page of the report.
	// A value of 'true' means this is the final page, while 'false' means more pages follow.
	LastPageIndicator bool `json:"last_page_indicator,omitempty"`
}
type NumberAndSumOfTransactions struct {
	// NbOfNtries (Number of Entries) specifies the total count of transactions reported.
	// This value represents the total number of individual transactions included in the report.
	NumberOfEntries string `json:"number_of_entries,omitempty"`
	// Sum represents the total monetary value of all transactions included in the report.
	// It aggregates the values of individual transactions to provide a summary amount.
	Sum float64 `json:"sum,omitempty"`
}

type TotalsPerBankTransactionCode struct {
	// NbOfNtries (Number of Entries) specifies the total number of transactions for a given bank transaction code.
	// This helps in categorizing transactions based on their type.
	NumberOfEntries string `json:"number_of_entries,omitempty"`
	// It is used when the transaction code follows a bank-specific classification rather than a standard one.
	BankTransactionCode TransactionStatusCode `json:"bank_transaction_code,omitempty"` // Bank Transaction Code
}
type CurrencyAndAmount struct {
	//default: USD
	Currency string  `json:"currency,omitempty"` // Currency code, e.g., "USD" for US Dollar
	Amount   float64 `json:"amount,omitempty"`   // Amount in the specified currency
}
type Entry struct {
	// Amt (Amount) specifies the transaction amount along with the currency.
	// It represents the value of the transaction.
	Amount CurrencyAndAmount `json:"amount,omitempty"`
	// CdtDbtInd (Credit or Debit Indicator) specifies whether the transaction is a credit (CRDT) or debit (DBIT).
	CreditDebitIndicator CdtDbtInd `json:"credit_debit_indicator,omitempty"` // Credit or Debit Indicator
	// Sts (Status) represents the current status of the transaction entry.
	// It may indicate if the transaction is booked, pending, or in another state
	Status ReportStatus `json:"status,omitempty"` // Status of the transaction entry
	// BkTxCd (Bank Transaction Code) defines the type of transaction.
	// It provides structured information about the nature of the transaction (e.g., deposit, withdrawal, fee).
	BankTransactionCode TransactionStatusCode `json:"bank_transaction_code,omitempty"` // Bank Transaction Code
	// <MsgNmId> (Message Name Identification) specifies the identifier for the message type or
	// message version. In this case, "pacs.008.001.08" is the identifier for a specific type of
	// payment message, indicating it is a version 08 of the pacs.008 (Payment Initiation Request) message.
	MessageNameId string `json:"message_name_id,omitempty"` // Message Name Identification
	//Provides details on the entry.
	EntryDetails EntryDetail `json:"entry_details,omitempty"`
}
type EntryDetail struct {
	// MsgId (Message ID) represents the unique identifier for the message.
	// It is used to track the specific transaction message.
	MessageId string `json:"message_id,omitempty"` // Message ID
	// InstrId (Instruction ID) is an optional field representing a reference for the instruction associated with the transaction.
	// It can be used to identify a particular instruction within the system.
	InstructionId string `json:"instruction_id,omitempty"` // Instruction ID
	//Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	EndToEndId string `json:"end_to_end_id,omitempty"` // End-to-End Transaction ID
	// UETR (Unique End-to-End Transaction Reference) is an optional field providing a globally unique reference for the transaction.
	// It is typically used to track and identify a specific transaction across different systems.
	UniqueTransactionReference string `json:"unique_transaction_reference,omitempty"` // Unique Transaction Reference

	TransactionIdentification string `json:"transaction_identification,omitempty"` // Transaction Identification
	// ClrSysRef (Clearing System Reference) is an optional field used to provide a reference to the clearing system that processes the transaction.
	// It is specific to the system used for clearing and settlement of the transaction.
	ClearingSystemRef string `json:"clearing_system_ref,omitempty"` // Clearing System Reference
	// InstgAgt (Instructing Agent) represents the financial institution or branch that is instructing the transaction.
	// This is the party that initiates the transaction and sends the payment instructions.
	InstructingAgent Agent `json:"instructing_agent,omitempty"` // Instructing Agent
	// InstdAgt (Instructed Agent) represents the financial institution or branch that is instructed to process the transaction.
	// This is the party that receives the payment instructions and is responsible for executing the transaction.
	InstructedAgent Agent `json:"instructed_agent,omitempty"` // Instructed Agent
	// LclInstrm (Local Instrument) is an optional field that refers to a local instrument or payment method for the transaction.
	// It indicates how the transaction is to be processed (e.g., via a local payment system).

	ProprietaryAgent      ProprietaryAgent       `json:"proprietary_agent,omitzero"`        // Proprietary Agent
	LocalInstrumentChoice InstrumentPropCodeType `json:"local_instrument_choice,omitempty"` // Local Instrument Choice
	//Tp (Type) indicates the type of the related date. In this case, 'BPRD' could represent a specific type of related date, like business processing date.

	AcceptanceDateTime time.Time         `json:"acceptance_date_time,omitzero"` // Acceptance DateTime
	TypeAndDateTimes   []TypeAndDateTime `json:"type_and_date_times,omitempty"` // Related Dates and Times

	AdditionalTransactionInformation string `json:"additional_transaction_information,omitempty"` // Additional Transaction Information
}
type OriginalBusinessQuery struct {
	MessageIdentification     string    `json:"message_identification,omitempty"`      // Message Identification
	MessageNameIdentification string    `json:"message_name_identification,omitempty"` // Message Name Identification
	CreationDateTime          time.Time `json:"creation_date_time,omitzero"`           // Creation DateTime
}
type TypeAndDateTime struct {
	RelatedDatesProprietary WorkingDayType `json:"related_dates_proprietary,omitempty"` // Related Dates Proprietary Code
	RelatedDateTime         time.Time      `json:"related_date_time,omitzero"`          // Related DateTime
}
type Agent struct {
	BusinessIdCode     string            `json:"business_id_code,omitempty"`      // Business Identifier Code (BIC) of the financial institution
	PaymentSysCode     PaymentSystemType `json:"payment_sys_code,omitempty"`      // Payment System Code
	PaymentSysMemberId string            `json:"payment_sys_member_id,omitempty"` // Payment System Member ID
	BankName           string            `json:"bank_name,omitempty"`             // Name of the bank or financial institution
	PostalAddress      PostalAddress     `json:"postal_address,omitzero"`         // Postal address of the agent
	OtherTypeId        string            `json:"other_type_id,omitempty"`         // Other type identifier for the agent
}
type ProprietaryAgent struct {
	Type               string `json:"type,omitempty"`                  // Type of the proprietary agent
	PaymentSysMemberId string `json:"payment_sys_member_id,omitempty"` // Payment System Member ID of the proprietary agent
}
type PostalAddress struct {
	StreetName     string `json:"street_name,omitempty"`     // Street name of the address
	BuildingNumber string `json:"building_number,omitempty"` // Building number of the address
	BuildingName   string `json:"building_name,omitempty"`   // Building name of the address
	Floor          string `json:"floor,omitempty"`           // Floor number of the address
	RoomNumber     string `json:"room_number,omitempty"`     // Room number of the address
	PostBox        string `json:"post_box,omitempty"`        // Post box number of the address
	PostalCode     string `json:"postal_code,omitempty"`     // Postal code of the address
	TownName       string `json:"town_name,omitempty"`       // Town or city name of the address
	Subdivision    string `json:"subdivision,omitempty"`     // Subdivision or state of the address
	Country        string `json:"country,omitempty"`         // Country of the address
}
type Balance struct {
	//Specifies the nature of a balance.
	BalanceTypeId BalanceType `json:"balance_type_id,omitempty"`

	CdtLines []CreditLine `json:"credit_lines,omitempty"` // Credit lines associated with the balance
	//Amount of money of the cash balance.
	Amount CurrencyAndAmount `json:"amount,omitempty"`
	//Indicates whether the balance is a credit or a debit balance.
	CreditDebitIndicator CdtDbtInd `json:"credit_debit_indicator,omitempty"`
	//Indicates the date (and time) of the balance.
	DateTime time.Time `json:"date_time,omitzero"` // Date and time of the balance
}
type CreditLine struct {
	//Indicates whether or not the credit line is included in the balance.
	Included bool `json:"included,omitempty"` // Indicates if the credit line is included in the balance
	//Type of the credit line provided when multiple credit lines may be provided.
	Type CreditLineType `json:"type,omitempty"` // Type of the credit line
	//Amount of money of the cash balance.
	Amount CurrencyAndAmount `json:"amount,omitempty"` // Amount of the credit line
	//Indicates the date (and time) of the balance.
	DateTime time.Time `json:"date_time,omitzero"` // Date and time of the credit line
}
type TotalsPerBankTransaction struct {
	TotalNetEntries      float64                    `json:"total_net_entry_amount,omitempty"` // Total net entry amount for the bank transaction code
	CreditDebitIndicator CdtDbtInd                  `json:"credit_debit_indicator,omitempty"` // Credit or Debit Indicator
	CreditEntries        NumberAndSumOfTransactions `json:"credit_entries,omitempty"`         // Credit entries for the bank transaction code
	DebitEntries         NumberAndSumOfTransactions `json:"debit_entries,omitempty"`          // Debit entries for the bank transaction code
	BankTransactionCode  TransactionCode            `json:"bank_transaction_code,omitempty"`  // Bank Transaction Code
	Date                 time.Time                  `json:"date,omitzero"`                    // Date of the bank transaction
}
type TransactionDetailReference struct {
	TransactionId             string `json:"transaction_id,omitempty"`             // Transaction ID
	InstructionIdentification string `json:"instruction_identification,omitempty"` // Instruction Identification
	EndToEndIdentification    string `json:"end_to_end_identification,omitempty"`  // End-to-End Identification
	UETR                      string `json:"uetr,omitempty"`                       // Unique End-to-End Transaction Reference
}
type RelatedAgents struct {
	InstructingAgent Agent                `json:"instructing_agent,omitempty"` // Instructing Agent
	InstructedAgent  Agent                `json:"instructed_agent,omitempty"`  // Instructed Agent
	ProprietaryType  ProprietaryAgentType `json:"proprietary_type,omitempty"`  // Proprietary Type
	ProprietaryAgent Agent                `json:"proprietary_agent,omitzero"`  // Proprietary Agent
}
type RelatedDates struct {
	AcceptanceDateTime      time.Time      `json:"acceptance_date_time,omitzero"`       // Acceptance DateTime
	InterbankSettlementDate fednow.ISODate `json:"interbank_settlement_date,omitempty"` // Interbank Settlement Date
}
type PeriodDateAndTime struct {
	FromDate fednow.ISODate `json:"from_date,omitzero"` // Start date of the period
	ToDate   fednow.ISODate `json:"to_date,omitzero"`   // End
	FromTime time.Time      `json:"from_time,omitzero"` // Start time of the period
	ToTime   time.Time      `json:"to_time,omitzero"`   // End time of the period
	Type     ModifyMode     `json:"type,omitzero"`      // Type of the period (e.g., ALLL, CHNG, MODF)
}
type Assignments struct {
	Assigner Agent `json:"assigner,omitempty"` // Assigner Agent
	Assignee Agent `json:"assignee,omitempty"` // Assignee Agent
}
type DebtorAndCreditorAgent struct {
	DebtorAgent   Agent `json:"debtor_agent,omitempty"`   // Debtor Agent
	CreditorAgent Agent `json:"creditor_agent,omitempty"` // Creditor Agent
}
type GroupInformation struct {
	OriginalMessageIdentification     string    `json:"original_message_identification,omitempty"`      // Original Message Identification
	OriginalMessageNameIdentification string    `json:"original_message_name_identification,omitempty"` // Original Message Name Identification
	OriginalCreationDateTime          time.Time `json:"original_creation_date_time,omitzero"`           // Original Creation DateTime
}
type SettlementInformation struct {
	Method  SettlementMethod           `json:"method_code,omitempty"`  // Settlement Method Code
	Service ExternalCashClearingSystem `json:"service_code,omitempty"` // FedNow Service Code
}
type PaymentTypeInfo struct {
	LocalInstrument LocalInstrumentType `json:"local_instrument_code,omitempty"` // Local Instrument Code
	CategoryPurpose string              `json:"category_purpose,omitempty"`      // Category Purpose
	ServiceLevel    ChargeBearerType    `json:"service_level,omitempty"`         // Service Level
}
type TransactionParty struct {
	Agent                Agent         `json:"agent,omitzero"`                  // Agent representing the transaction party
	PartyName            string        `json:"party_name,omitzero"`             // Name of the transaction party
	PartyAddress         PostalAddress `json:"party_address,omitzero"`          // Address of the transaction party
	PartyAccountId       string        `json:"party_account_id,omitzero"`       // Account ID of the transaction party
	PartyAccountProxy    Proxy         `json:"party_account_proxy,omitzero"`    // Proxy for the transaction party's account
	PartyAgent           Agent         `json:"party_agent,omitzero"`            // Agent representing the transaction party
	UltimatePartyName    string        `json:"ultimate_party_name,omitzero"`    // Ultimate party name in the transaction
	UltimatePartyAddress PostalAddress `json:"ultimate_party_address,omitzero"` // Ultimate party address in the transaction
}
type RemittanceInformation struct {
	Unstructured                string                `json:"unstructured,omitempty"`                  // Unstructured remittance information
	ReferredDocumentInformation CodeOrProprietaryType `json:"referred_document_information,omitempty"` // Referred Document Information
	Number                      string                `json:"number,omitempty"`                        // Number of the referred document
	RelatedDate                 fednow.ISODate        `json:"related_date,omitempty"`                  // Related Date of
}
type RelatedRemittanceInformation struct {
	RemittanceIdentification string                   `json:"remittance_identification,omitempty"` // Remittance Identification
	Method                   RemittanceDeliveryMethod `json:"method,omitempty"`                    // Method of remittance delivery
	ElectronicAddress        string                   `json:"electronic_address,omitempty"`        // Electronic Address for remittance delivery
}
type Proxy struct {
	Type  ProxyType `json:"type,omitempty"`  // Type of the proxy (e.g., TELN for Telephone Number, EMAL for Email Address)
	Value string    `json:"value,omitempty"` // Value of the proxy (e.g., actual telephone number or email address)
}
type Reason struct {
	Code        StatusReasonInformationCode `json:"type,omitempty"`        // Type of the reason for the transaction status
	Proprietary string                      `json:"proprietary,omitempty"` // Code representing the reason
}
type ReturnReason struct {
	Code        ReturnReasonCode `json:"code,omitempty"`        // Code representing the return reason
	Info        string           `json:"info,omitempty"`        // Additional information about the return reason
	Proprietary string           `json:"proprietary,omitempty"` // Code representing the reason
}
type MarketPractice struct {
	Registry       string `json:"registry_code,omitempty"`       // Registry Code for the market practice
	Identification string `json:"identification_code,omitempty"` // Identification Code for the market practice
}
type ContactDetails struct {
	Name            string        `json:"name,omitempty"`             // Name of the contact person
	Email           string        `json:"email,omitempty"`            // Email address of the contact person
	Phone           string        `json:"phone,omitempty"`            // Phone number of the contact person
	PreferredMethod ContactMethod `json:"preferred_method,omitempty"` // Preferred method of contact
}
type MissingCodeAndInfo struct {
	Code MissingOrIncorrectInformationCode `json:"code,omitempty"` // Code representing the missing or incorrect information
	Info string                            `json:"info,omitempty"` // Additional information about the missing or incorrect information
}
type Creator struct {
	Name               string            `json:"name,omitzero"`                  // Name of the creator
	PaymentSysCode     PaymentSystemType `json:"payment_sys_code,omitzero"`      // Payment system code of the creator
	PaymentSysMemberId string            `json:"payment_sys_member_id,omitzero"` // Payment system
	Contact            ContactDetails    `json:"contact,omitzero"`               // Contact details of the creator
}
type PaymentInfomation struct {
	OriginalGroupInfo         GroupInformation           `json:"original_group_info,omitzero"`                  // Information about the original group
	OriginalTransactionDetail TransactionDetailReference `json:"original_transaction_detail,omitzero"`          // Reference to the original transaction detail
	SettlementAmount          CurrencyAndAmount          `json:"original_interbank_settlement_amount,omitzero"` // Amount of the original interbank settlement
	SettlementDate            fednow.ISODate             `json:"original_interbank_settlement_date,omitzero"`
}
type PartyContact struct {
	Name            string        `json:"name,omitempty"`
	PhoneNumber     string        `json:"phone_number,omitempty"`
	EmailAddress    string        `json:"email_address,omitempty"`
	MobileNumber    string        `json:"mobile_number,omitempty"`
	Department      string        `json:"department,omitempty"` // Department of the contact person
	PreferredMethod ContactMethod `json:"preferred_method,omitempty"`
}
