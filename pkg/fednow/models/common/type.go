package common

import "time"

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

const (
	BusinessProcessingDate WorkingDayType = "BPRD"
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

	ProprietaryAgent      ProprietaryAgent       `json:"proprietary_agent,omitempty"`       // Proprietary Agent
	LocalInstrumentChoice InstrumentPropCodeType `json:"local_instrument_choice,omitempty"` // Local Instrument Choice
	//Tp (Type) indicates the type of the related date. In this case, 'BPRD' could represent a specific type of related date, like business processing date.

	AcceptanceDateTime time.Time         `json:"acceptance_date_time,omitempty"` // Acceptance DateTime
	TypeAndDateTimes   []TypeAndDateTime `json:"type_and_date_times,omitempty"`  // Related Dates and Times

	AdditionalTransactionInformation string `json:"additional_transaction_information,omitempty"` // Additional Transaction Information
}
type OriginalBusinessQuery struct {
	MessageIdentification     string    `json:"message_identification,omitempty"`      // Message Identification
	MessageNameIdentification string    `json:"message_name_identification,omitempty"` // Message Name Identification
	CreationDateTime          time.Time `json:"creation_date_time,omitempty"`          // Creation DateTime
}
type TypeAndDateTime struct {
	RelatedDatesProprietary WorkingDayType `json:"related_dates_proprietary,omitempty"` // Related Dates Proprietary Code
	RelatedDateTime         time.Time      `json:"related_date_time,omitempty"`         // Related DateTime
}
type Agent struct {
	BusinessIdCode     string            `json:"business_id_code,omitempty"`      // Business Identifier Code (BIC) of the financial institution
	PaymentSysCode     PaymentSystemType `json:"payment_sys_code,omitempty"`      // Payment System Code
	PaymentSysMemberId string            `json:"payment_sys_member_id,omitempty"` // Payment System Member ID
	BankName           string            `json:"bank_name,omitempty"`             // Name of the bank or financial institution
	PostalAddress      PostalAddress     `json:"postal_address,omitempty"`        // Postal address of the agent
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
	DateTime time.Time `json:"date_time,omitempty"` // Date and time of the balance
}
type CreditLine struct {
	//Indicates whether or not the credit line is included in the balance.
	Included bool `json:"included,omitempty"` // Indicates if the credit line is included in the balance
	//Type of the credit line provided when multiple credit lines may be provided.
	Type CreditLineType `json:"type,omitempty"` // Type of the credit line
	//Amount of money of the cash balance.
	Amount CurrencyAndAmount `json:"amount,omitempty"` // Amount of the credit line
	//Indicates the date (and time) of the balance.
	DateTime time.Time `json:"date_time,omitempty"` // Date and time of the credit line
}
type TotalsPerBankTransaction struct {
	TotalNetEntries      float64                    `json:"total_net_entry_amount,omitempty"` // Total net entry amount for the bank transaction code
	CreditDebitIndicator CdtDbtInd                  `json:"credit_debit_indicator,omitempty"` // Credit or Debit Indicator
	CreditEntries        NumberAndSumOfTransactions `json:"credit_entries,omitempty"`         // Credit entries for the bank transaction code
	DebitEntries         NumberAndSumOfTransactions `json:"debit_entries,omitempty"`          // Debit entries for the bank transaction code
	BankTransactionCode  TransactionCode            `json:"bank_transaction_code,omitempty"`  // Bank Transaction Code
	Date                 time.Time                  `json:"date,omitempty"`                   // Date of the bank transaction
}
