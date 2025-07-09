package common

type ElementHelper struct {
	Title         string
	Rules         string
	Type          string
	Documentation string
}

type PostalAddressHelper struct {
	StreetName     ElementHelper
	BuildingNumber ElementHelper
	BuildingName   ElementHelper
	Floor          ElementHelper
	RoomNumber     ElementHelper
	PostalCode     ElementHelper
	TownName       ElementHelper
	Subdivision    ElementHelper
	Country        ElementHelper
}

func BuildPostalAddressHelper() PostalAddressHelper {
	return PostalAddressHelper{
		StreetName: ElementHelper{
			Title:         "Street Name",
			Rules:         "",
			Type:          `Max70Text (based on string) minLength: 1 maxLength: 70`,
			Documentation: `Name of a street or thoroughfare.`,
		},
		BuildingNumber: ElementHelper{
			Title:         "Building Number",
			Rules:         "",
			Type:          `Max16Text (based on string) minLength: 1 maxLength: 16`,
			Documentation: `Number that identifies the position of a building on a street.`,
		},
		BuildingName: ElementHelper{
			Title:         "Building Name",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Name of the building or house.`,
		},
		Floor: ElementHelper{
			Title:         "Floor",
			Rules:         "",
			Type:          `Max70Text (based on string) minLength: 1 maxLength: 70`,
			Documentation: `Floor or storey within a building.`,
		},
		RoomNumber: ElementHelper{
			Title:         "Room Number",
			Rules:         "",
			Type:          `Max70Text (based on string) minLength: 1 maxLength: 70`,
			Documentation: `Building room number.`,
		},
		PostalCode: ElementHelper{
			Title:         "Postal Code",
			Rules:         "",
			Type:          `Max16Text (based on string) minLength: 1 maxLength: 16`,
			Documentation: `Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.`,
		},
		TownName: ElementHelper{
			Title:         "Town Name",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Name of a built-up area, with defined boundaries, and a local government.`,
		},
		Subdivision: ElementHelper{
			Title:         "Country Sub Division",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Identifies a subdivision of a country such as state, region, county.`,
		},
		Country: ElementHelper{
			Title:         "Country",
			Rules:         "",
			Type:          `CountryCode (based on string) pattern: [A-Z]{2,2}`,
			Documentation: `Nation with its own government.`,
		},
	}
}

type AgentHelper struct {
	BusinessIdCode     ElementHelper
	PaymentSysCode     ElementHelper
	PaymentSysMemberId ElementHelper
	BankName           ElementHelper
	PostalAddress      PostalAddressHelper
	OtherTypeId        ElementHelper
}

func BuildAgentHelper() AgentHelper {
	return AgentHelper{
		BusinessIdCode: ElementHelper{
			Title:         "BICFI",
			Rules:         "",
			Type:          `BICFIDec2014Identifier (based on string) pattern: [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1} identificationScheme: SWIFT; BICIdentifier`,
			Documentation: `Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".`,
		},
		PaymentSysCode: ElementHelper{
			Title:         "Clearing System Identification Code",
			Rules:         "",
			Type:          `PaymentSystemType(PaymentSysUSABA, PaymentSysCHIPS, PaymentSysSEPA, PaymentSysRTGS, PaymentSysSWIFT, PaymentSysBACS)`,
			Documentation: `Identification of a clearing system, in a coded form as published in an external list.`,
		},
		PaymentSysMemberId: ElementHelper{
			Title:         "Member Identification",
			Rules:         "",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification of a member of a clearing system.`,
		},
		BankName: ElementHelper{
			Title:         "Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which an agent is known and which is usually used to identify that agent.`,
		},
		OtherTypeId: ElementHelper{
			Title:         "Other Type Id",
			Rules:         `Must be the Inquiry Routing Number Usage: It may be a subaccount routing number or the master account routing number. This element in conjunction with the balance type code in the Account\Type element will determine the information reported in the Account Balance Report.`,
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification assigned by an institution.`,
		},
	}
}

type ProprietaryAgentHelper struct {
	Type               ElementHelper
	PaymentSysMemberId ElementHelper
}

func BuildProprietaryAgentHelper() ProprietaryAgentHelper {
	return ProprietaryAgentHelper{
		Type: ElementHelper{
			Title:         "Type",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Type of agent, in a coded form as published in an external list.`,
		},
		PaymentSysMemberId: ElementHelper{
			Title:         "Member Identification",
			Rules:         "",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification of a member of a clearing system.`,
		},
	}
}

type SequenceRangeHelper struct {
	FromSeq ElementHelper
	ToSeq   ElementHelper
}

func BuildSequenceRangeHelper() SequenceRangeHelper {
	return SequenceRangeHelper{
		FromSeq: ElementHelper{
			Title:         "From Sequence",
			Rules:         "Sequence number in From Sequence must be less than or equal to sequence number in To Sequence.",
			Type:          `SequenceNumber_FedwireFunds_1 (based on decimal) minInclusive: 000001 maxInclusive: 999999 pattern: [0-9]{6,6}`,
			Documentation: `Start sequence of the range.`,
		},
		ToSeq: ElementHelper{
			Title:         "To Sequence",
			Rules:         "Sequence number in From Sequence must be less than or equal to sequence number in To Sequence.",
			Type:          `SequenceNumber_FedwireFunds_1 (based on decimal) minInclusive: 000001 maxInclusive: 999999 pattern: [0-9]{6,6}`,
			Documentation: `End sequence of the range.`,
		},
	}
}

type MessagePagenationHelper struct {
	PageNumber        ElementHelper
	LastPageIndicator ElementHelper
}

func BuildMessagePagenationHelper() MessagePagenationHelper {
	return MessagePagenationHelper{
		PageNumber: ElementHelper{
			Title:         "Page Number",
			Rules:         "",
			Type:          `Max5NumericText (based on string) pattern: [0-9]{1,5}`,
			Documentation: `Page number.`,
		},
		LastPageIndicator: ElementHelper{
			Title:         "Last Page Indicator",
			Rules:         "",
			Type:          `YesNoIndicator (based on boolean) meaningWhenFalse: No meaningWhenTrue: Yes`,
			Documentation: `Indicates the last page.`,
		},
	}
}

type NumberAndSumOfTransactionsHelper struct {
	NumberOfEntries ElementHelper
	Sum             ElementHelper
}

func BuildNumberAndSumOfTransactionsHelper() NumberAndSumOfTransactionsHelper {
	return NumberAndSumOfTransactionsHelper{
		NumberOfEntries: ElementHelper{
			Title:         "Number Of Entries",
			Rules:         "",
			Type:          `Max15NumericText (based on string) pattern: [0-9]{1,15}`,
			Documentation: `Number of individual entries included in the report.`,
		},
		Sum: ElementHelper{
			Title:         "Sum",
			Rules:         "",
			Type:          `DecimalNumber (based on decimal) totalDigits: 18 fractionDigits: 17`,
			Documentation: `Total of all individual entries included in the report.`,
		},
	}
}

type CurrencyAndAmountHelper struct {
	Currency ElementHelper
	Amount   ElementHelper
}

func BuildCurrencyAndAmountHelper() CurrencyAndAmountHelper {
	return CurrencyAndAmountHelper{
		Currency: ElementHelper{
			Title:         "Currency",
			Rules:         "",
			Type:          `ActiveOrHistoricCurrencyCode (based on string) pattern: [A-Z]{3,3}`,
			Documentation: `Medium of exchange of currency.`,
		},
		Amount: ElementHelper{
			Title:         "Amount",
			Rules:         "",
			Type:          `DecimalNumber (based on decimal) totalDigits: 18 fractionDigits: 17`,
			Documentation: `The number of fractional digits (or minor unit of currency) must comply with ISO 4217. Note: The decimal separator is a dot.`,
		},
	}
}

type OriginalBusinessQueryHelper struct {
	MessageIdentification     ElementHelper
	MessageNameIdentification ElementHelper
	CreationDateTime          ElementHelper
}

func BuildOriginalBusinessQueryHelper() OriginalBusinessQueryHelper {
	return OriginalBusinessQueryHelper{
		MessageIdentification: ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `This is the Message Identification of the account report request message sent by the FedNow participant to request the account activity details report. `,
		},
		MessageNameIdentification: ElementHelper{
			Title:         "Message Name Identification",
			Rules:         "",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `This is the Message Identification Name of the account reporting request message sent by the FedNow participant to request the account activity details report (i.e., camt.060.001.05 or a subsequent version of the message as it is introduced in a FedNow Service future release).`,
		},
		CreationDateTime: ElementHelper{
			Title:         "Creation DateTime",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `This is the Creation Date Time of the account report request message sent by the FedNow participant to request the account activity details report. Time is in 24-hour clock format and includes the offset against the Coordinated Universal Time (UTC). `,
		},
	}
}

type TypeAndDateTimeHelper struct {
	RelatedDatesProprietary ElementHelper
	RelatedDateTime         ElementHelper
}

func BuildTypeAndDateTimeHelper() TypeAndDateTimeHelper {
	return TypeAndDateTimeHelper{
		RelatedDatesProprietary: ElementHelper{
			Title:         "Related Dates Proprietary",
			Rules:         "",
			Type:          `WorkingDayType`,
			Documentation: `Specifies the type of date.`,
		},
		RelatedDateTime: ElementHelper{
			Title:         "Related Dates Proprietary",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date in ISO format.`,
		},
	}
}

type EntryDetailHelper struct {
	MessageId                        ElementHelper
	InstructionId                    ElementHelper
	EndToEndId                       ElementHelper
	UniqueTransactionReference       ElementHelper
	TransactionIdentification        ElementHelper
	ClearingSystemRef                ElementHelper
	InstructingAgent                 AgentHelper
	InstructedAgent                  AgentHelper
	LocalInstrumentChoice            ElementHelper
	AcceptanceDateTime               ElementHelper
	TypeAndDateTimes                 TypeAndDateTimeHelper
	AdditionalTransactionInformation ElementHelper
}

func BuildEntryDetailHelper() EntryDetailHelper {
	return EntryDetailHelper{
		MessageId: ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `IMAD_FedwireFunds_1 (based on string) minLength: 22 maxLength: 22 pattern: [0-9]{8}[A-Z0-9]{8}[0-9]{6}`,
			Documentation: `Point to point reference, as assigned by the instructing party of the underlying message.`,
		},
		InstructionId: ElementHelper{
			Title:         "Instruction Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction. Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.`,
		},
		EndToEndId: ElementHelper{
			Title:         "EndToEndId",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.`,
		},
		UniqueTransactionReference: ElementHelper{
			Title:         "UETR",
			Rules:         "",
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide an end-to-end reference of a payment transaction.`,
		},
		TransactionIdentification: ElementHelper{
			Title:         "Transaction Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.`,
		},
		ClearingSystemRef: ElementHelper{
			Title:         "Clearing System Reference",
			Rules:         "",
			Type:          `OMAD_FedwireFunds_1 (based on string) pattern: [A-Z0-9]{34,34}`,
			Documentation: `Unique reference, as assigned by a clearing system, to unambiguously identify the instruction. Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.`,
		},
		InstructingAgent: BuildAgentHelper(),
		InstructedAgent:  BuildAgentHelper(),
		LocalInstrumentChoice: ElementHelper{
			Title:         "Local Instrument Choice",
			Rules:         "",
			Type:          `InstrumentPropCodeType(InstrumentCTRC, InstrumentDD ...)`,
			Documentation: `User community specific instrument. Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.`,
		},
		AcceptanceDateTime: ElementHelper{
			Title:         "Acceptance Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.`,
		},
		TypeAndDateTimes: BuildTypeAndDateTimeHelper(),
		AdditionalTransactionInformation: ElementHelper{
			Title:         "Additional Transaction Information",
			Rules:         "",
			Type:          `Max500Text (based on string) minLength: 1 maxLength: 500`,
			Documentation: `Further details of the transaction.`,
		},
	}
}

type EntryHelper struct {
	Amount               CurrencyAndAmountHelper
	CreditDebitIndicator ElementHelper
	Status               ElementHelper
	BankTransactionCode  ElementHelper
	MessageNameId        ElementHelper
	EntryDetails         EntryDetailHelper
}

func BuildEntryHelper() EntryHelper {
	return EntryHelper{
		Amount: BuildCurrencyAndAmountHelper(),
		CreditDebitIndicator: ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `CdtDbtInd(Credit, Debit)`,
			Documentation: `Indicates whether the transaction is a credit or a debit transaction.`,
		},
		Status: ElementHelper{
			Title:         "Status",
			Rules:         "",
			Type:          `ReportStatus(Book, Pending, Received, Settled)`,
			Documentation: `Status of an entry on the books of the account servicer.`,
		},
		BankTransactionCode: ElementHelper{
			Title:         "Bank Transaction Code",
			Rules:         "",
			Type:          `TransactionStatusCode(MessagesInProcess, MessagesIntercepted ...)`,
			Documentation: `Set of elements used to fully identify the type of underlying transaction resulting in an entry.`,
		},
		MessageNameId: ElementHelper{
			Title:         "Message Name Identification",
			Rules:         "",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `Specifies the message name identifier of the message that will be used to provide additional details.`,
		},
		EntryDetails: BuildEntryDetailHelper(),
	}
}

type PartyIdentifyHelper struct {
	Name    ElementHelper
	Address PostalAddressHelper
}

func BuildPartyIdentifyHelper() PartyIdentifyHelper {
	return PartyIdentifyHelper{
		Name: ElementHelper{
			Title:         "Name",
			Rules:         "If BIC is not present, then Name must be present. Postal address information may be required under applicable law. Even when not required, it is strongly recommended to include this information to the extent possible.",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which a party is known and which is usually used to identify that party.`,
		},
		Address: BuildPostalAddressHelper(),
	}
}

type NumberAndStatusOfTransactionsHelper struct {
	NumberOfEntries ElementHelper
	Status          ElementHelper
}

func BuildNumberAndStatusOfTransactionsHelper() NumberAndStatusOfTransactionsHelper {
	return NumberAndStatusOfTransactionsHelper{
		NumberOfEntries: ElementHelper{
			Title:         "Number Of Entries",
			Rules:         "",
			Type:          `Max15NumericText (based on string) pattern: [0-9]{1,15}`,
			Documentation: `Number of individual entries for the bank transaction code.`,
		},
		Status: ElementHelper{
			Title:         "Status",
			Rules:         "",
			Type:          `TransactionStatusCode(MessagesInProcess, MessagesIntercepted...)`,
			Documentation: `Proprietary bank transaction code to identify the underlying transaction.`,
		},
	}
}

type FiniancialInstitutionIdHelper struct {
	BusinessId             ElementHelper
	ClearingSystemId       ElementHelper
	ClearintSystemMemberId ElementHelper
	Name                   ElementHelper
	Address                PostalAddressHelper
}

func BuildFiniancialInstitutionIdHelper() FiniancialInstitutionIdHelper {
	return FiniancialInstitutionIdHelper{
		BusinessId: ElementHelper{
			Title:         "BICFI",
			Rules:         "",
			Type:          `BICFIDec2014Identifier (based on string) pattern: [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1} identificationScheme: SWIFT; BICIdentifier`,
			Documentation: `Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".`,
		},
		ClearingSystemId: ElementHelper{
			Title:         "Clearing System Identification Code",
			Rules:         "",
			Type:          `PaymentSystemType(PaymentSysUSABA, PaymentSysCHIPS, PaymentSysSEPA, PaymentSysRTGS, PaymentSysSWIFT, PaymentSysBACS)`,
			Documentation: `Identification of a clearing system, in a coded form as published in an external list.`,
		},
		ClearintSystemMemberId: ElementHelper{
			Title:         "Member Identification",
			Rules:         "",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification of a member of a clearing system.`,
		},
		Name: ElementHelper{
			Title:         "Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which an agent is known and which is usually used to identify that agent.`,
		},
		Address: BuildPostalAddressHelper(),
	}
}

type CreditLineHelper struct {
	Included ElementHelper
	Type     ElementHelper
	Amount   CurrencyAndAmountHelper
	DateTime ElementHelper
}

func BuildCreditLineHelper() CreditLineHelper {
	return CreditLineHelper{
		Included: ElementHelper{
			Title:         "Included",
			Rules:         "",
			Type:          `Boolean (based on string)`,
			Documentation: `Indicates whether or not the credit line is included in the balance.`,
		},
		Type: ElementHelper{
			Title:         "Type",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Type of the credit line provided when multiple credit lines may be provided.`,
		},
		Amount: BuildCurrencyAndAmountHelper(),
		DateTime: ElementHelper{
			Title:         "Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Indicates the date (and time) of the balance.`,
		},
	}
}

type TotalsPerBankTransactionHelper struct {
	TotalNetEntries      ElementHelper
	CreditDebitIndicator ElementHelper
	CreditEntries        NumberAndSumOfTransactionsHelper
	DebitEntries         NumberAndSumOfTransactionsHelper
	BankTransactionCode  ElementHelper
	Date                 ElementHelper
}

func BuildTotalsPerBankTransactionHelper() TotalsPerBankTransactionHelper {
	return TotalsPerBankTransactionHelper{
		TotalNetEntries: ElementHelper{
			Title:         "Number Of Entries",
			Rules:         "",
			Type:          `Max15NumericText (based on string) pattern: [0-9]{1,15}`,
			Documentation: `Number of individual entries included in the report.`,
		},
		CreditDebitIndicator: ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `CdtDbtInd(Credit, Debit)`,
			Documentation: `Indicates whether the balance is a credit or a debit balance.`,
		},
		CreditEntries: BuildNumberAndSumOfTransactionsHelper(),
		DebitEntries:  BuildNumberAndSumOfTransactionsHelper(),
		BankTransactionCode: ElementHelper{
			Title:         "Bank Transaction Code",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the bank transaction code to which the entry refers.`,
		},
		Date: ElementHelper{
			Title:         "Date",
			Rules:         "",
			Type:          `ISODate (based on string)`,
			Documentation: `Date at which the transaction was executed.`,
		},
	}
}

type TransactionDetailReferenceHelper struct {
	TransactionId             ElementHelper
	InstructionIdentification ElementHelper
	EndToEndIdentification    ElementHelper
	UETR                      ElementHelper
}

func BuildTransactionDetailReferenceHelper() TransactionDetailReferenceHelper {
	return TransactionDetailReferenceHelper{
		TransactionId: ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the instructing party of the underlying message.`,
		},
		InstructionIdentification: ElementHelper{
			Title:         "Instruction Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.`,
		},
		EndToEndIdentification: ElementHelper{
			Title:         "End To End Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.`,
		},
		UETR: ElementHelper{
			Title:         "UETR",
			Rules:         "",
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide an end-to-end reference of a payment transaction.`,
		},
	}
}

type RelatedAgentsHelper struct {
	InstructingAgent AgentHelper
	InstructedAgent  AgentHelper
	ProprietaryType  ElementHelper
	ProprietaryAgent AgentHelper
}

func BuildRelatedAgentsHelper() RelatedAgentsHelper {
	return RelatedAgentsHelper{
		InstructingAgent: BuildAgentHelper(),
		InstructedAgent:  BuildAgentHelper(),
		ProprietaryType: ElementHelper{
			Title:         "Proprietary Type",
			Rules:         "",
			Type:          `This element contains the code 'RESP' (Respondent) to identify the related agent is a respondent of the receiver of the Account Debit Credit Notification message.`,
			Documentation: `Type of agent, in a coded form as published in an external list.`,
		},
		ProprietaryAgent: BuildAgentHelper(),
	}
}

type RelatedDatesHelper struct {
	AcceptanceDateTime      ElementHelper
	InterbankSettlementDate ElementHelper
}

func BuildRelatedDatesHelper() RelatedDatesHelper {
	return RelatedDatesHelper{
		AcceptanceDateTime: ElementHelper{
			Title:         "Acceptance Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.`,
		},
		InterbankSettlementDate: ElementHelper{
			Title:         "Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
	}
}

type TotalsPerBankTransactionCodeHelper struct {
	NumberOfEntries     ElementHelper
	BankTransactionCode ElementHelper
}

func BuildTotalsPerBankTransactionCodeHelper() TotalsPerBankTransactionCodeHelper {
	return TotalsPerBankTransactionCodeHelper{
		NumberOfEntries: ElementHelper{
			Title:         "Number of Entries",
			Rules:         "",
			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
			Documentation: `Number of individual entries for the bank transaction code.`,
		},
		BankTransactionCode: ElementHelper{
			Title:         "Bank Transaction Code",
			Rules:         "",
			Type:          `TransactionStatusCode(MessagesInProcess, MessagesIntercepted ...)`,
			Documentation: `Bank transaction code in a proprietary form, as defined by the issuer.`,
		},
	}
}

type BalanceHelper struct {
	BalanceTypeId        ElementHelper
	CreditLine           CreditLineHelper
	Amount               CurrencyAndAmountHelper
	CreditDebitIndicator ElementHelper
	DateTime             ElementHelper
}

func BuildBalanceHelper() BalanceHelper {
	return BalanceHelper{
		BalanceTypeId: ElementHelper{
			Title:         "Balance Type Id",
			Rules:         "",
			Type:          `BalanceType(AccountBalance, AvailableBalanceFromAccountBalance ...)`,
			Documentation: `Specifies the nature of a balance.`,
		},
		CreditLine: BuildCreditLineHelper(),
		Amount:     BuildCurrencyAndAmountHelper(),
		CreditDebitIndicator: ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `CdtDbtInd(Credit, Debit)`,
			Documentation: `Indicates whether the balance is a credit or a debit balance.`,
		},
		DateTime: ElementHelper{
			Title:         "Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Indicates the date (and time) of the balance.`,
		},
	}
}

type PeriodDateAndTimeHelper struct {
	FromDate ElementHelper
	ToDate   ElementHelper
	FromTime ElementHelper
	ToTime   ElementHelper
	Type     ElementHelper
}

func BuildPeriodDateAndTimeHelper() PeriodDateAndTimeHelper {
	return PeriodDateAndTimeHelper{
		FromDate: ElementHelper{
			Title:         "From Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Start date of the range.`,
		},
		ToDate: ElementHelper{
			Title:         "To Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `End date of the range.`,
		},
		FromTime: ElementHelper{
			Title:         "From Time",
			Rules:         "",
			Type:          `ISOTime (based on time)`,
			Documentation: `Time at which the time span starts.`,
		},
		ToTime: ElementHelper{
			Title:         "To Time",
			Rules:         "",
			Type:          `ISOTime (based on time)`,
			Documentation: `Time at which the time span ends.`,
		},
		Type: ElementHelper{
			Title:         "Type",
			Rules:         "",
			Type:          `ModifyMode`,
			Documentation: `Specifies that the query requests that only items that have changed since the last query be returned.`,
		},
	}
}

type AssignmentsHelper struct {
	Assigner AgentHelper
	Assignee AgentHelper
}

func BuildAssignmentsHelper() AssignmentsHelper {
	return AssignmentsHelper{
		Assigner: BuildAgentHelper(),
		Assignee: BuildAgentHelper(),
	}
}

type GroupInformationHelper struct {
	OriginalMessageIdentification     ElementHelper
	OriginalMessageNameIdentification ElementHelper
	OriginalCreationDateTime          ElementHelper
}

func BuildGroupInformationHelper() GroupInformationHelper {
	return GroupInformationHelper{
		OriginalMessageIdentification: ElementHelper{
			Title:         "Original Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `This is the Message Identification of the account report request message sent by the FedNow participant to request the account activity details report.`,
		},
		OriginalMessageNameIdentification: ElementHelper{
			Title:         "Original Message Name Identification",
			Rules:         "",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `This is the Message Identification Name of the account reporting request message sent by the FedNow participant to request the account activity details report (i.e., camt.060.001.05 or a subsequent version of the message as it is introduced in a FedNow Service future release).`,
		},
		OriginalCreationDateTime: ElementHelper{
			Title:         "Original Creation DateTime",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `This is the Creation Date Time of the account report request message sent by the FedNow participant to request the account activity details report. Time is in 24-hour clock format and includes the offset against the Coordinated Universal Time (UTC).`,
		},
	}
}

type SettlementInformationHelper struct {
	Method  ElementHelper
	Service ElementHelper
}

func BuildSettlementInformationHelper() SettlementInformationHelper {
	return SettlementInformationHelper{
		Method: ElementHelper{
			Title:         "Method",
			Rules:         "",
			Type:          `SettlementMethodType`,
			Documentation: `Specifies the method of settlement.`,
		},
		Service: ElementHelper{
			Title:         "Service",
			Rules:         "",
			Type:          `ExternalCashClearingSystem`,
			Documentation: `Specifies the service used for settlement.`,
		},
	}
}

type PaymentTypeInfoHelper struct {
	LocalInstrument ElementHelper
	CategoryPurpose ElementHelper
}

func BuildPaymentTypeInfoHelper() PaymentTypeInfoHelper {
	return PaymentTypeInfoHelper{
		LocalInstrument: ElementHelper{
			Title:         "Local Instrument",
			Rules:         "",
			Type:          `LocalInstrumentType`,
			Documentation: `User community specific instrument. Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.`,
		},
		CategoryPurpose: ElementHelper{
			Title:         "Category Purpose",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 3`,
			Documentation: `Category purpose, in a proprietary form.`,
		},
	}
}

type TransactionPartyHelper struct {
	Agent          AgentHelper
	PartyName      ElementHelper
	PartyAddress   PostalAddressHelper
	PartyAccountId ElementHelper
	PartyAgent     AgentHelper
}

func BuildTransactionPartyHelper() TransactionPartyHelper {
	return TransactionPartyHelper{
		Agent: BuildAgentHelper(),
		PartyName: ElementHelper{
			Title:         "Party Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which a party is known and which is usually used to identify that party.`,
		},
		PartyAddress: BuildPostalAddressHelper(),
		PartyAccountId: ElementHelper{
			Title:         "Party Account Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Identification of the account of the party.`,
		},
		PartyAgent: BuildAgentHelper(),
	}
}

type RemittanceInformationHelper struct {
	ReferredDocumentInformation ElementHelper
	Number                      ElementHelper
	RelatedDate                 ElementHelper
}

func BuildRemittanceInformationHelper() RemittanceInformationHelper {
	return RemittanceInformationHelper{
		ReferredDocumentInformation: ElementHelper{
			Title:         "Referred Document Information",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Reference to a document that is referred to in the remittance information.`,
		},
		Number: ElementHelper{
			Title:         "Number",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Number of the document that is referred to in the remittance information.`,
		},
		RelatedDate: ElementHelper{
			Title:         "Related Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date related to the document that is referred to in the remittance information.`,
		},
	}
}

type RelatedRemittanceInformationHelper struct {
	RemittanceIdentification ElementHelper
	Method                   ElementHelper
	ElectronicAddress        ElementHelper
}

func BuildRelatedRemittanceInformationHelper() RelatedRemittanceInformationHelper {
	return RelatedRemittanceInformationHelper{
		RemittanceIdentification: ElementHelper{
			Title:         "Remittance Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Identification of the remittance information.`,
		},
		Method: ElementHelper{
			Title:         "Method",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Method of remittance information.`,
		},
		ElectronicAddress: ElementHelper{
			Title:         "Electronic Address",
			Rules:         "",
			Type:          `Max2048Text (based on string) minLength: 1 maxLength: 2048`,
			Documentation: `Electronic address for remittance information.`,
		},
	}
}

type ReasonHelper struct {
	Code        ElementHelper
	Proprietary ElementHelper
}

func BuildReasonHelper() ReasonHelper {
	return ReasonHelper{
		Code: ElementHelper{
			Title:         "Code",
			Rules:         "",
			Type:          `StatusReasonInformationCode`,
			Documentation: `Reason for the status, as published in an external reason code list.`,
		},
		Proprietary: ElementHelper{
			Title:         "Proprietary",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Reason for the status, in a proprietary form.`,
		},
	}
}

type ReturnReasonHelper struct {
	Code        ElementHelper
	Proprietary ElementHelper
}

func BuildReturnReasonHelper() ReturnReasonHelper {
	return ReturnReasonHelper{
		Code: ElementHelper{
			Title:         "Code",
			Rules:         "",
			Type:          `ReturnReasonCode`,
			Documentation: `Reason for the return, as published in an external reason code list.`,
		},
		Proprietary: ElementHelper{
			Title:         "Proprietary",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Reason for the return, in a proprietary form.`,
		},
	}
}

type MarketPracticeHelper struct {
	Identification ElementHelper
	Registry       ElementHelper
}

func BuildMarketPracticeHelper() MarketPracticeHelper {
	return MarketPracticeHelper{
		Identification: ElementHelper{
			Title:         "Identification",
			Rules:         "",
			Type:          `MarketPracticeIdentification_FedNow_1 (based on string) minLength: 13 maxLength: 17 pattern: frb([.]{1,1})fednow([.]{1,1})(([a-z]{3,3}[.]{1,1})){0,1}01`,
			Documentation: `Identifier which unambiguously identifies, within the implementation specification registry, the implementation specification to which the ISO 20022 message is compliant. This can be done via a URN. It can also contain a version number or date.`,
		},
		Registry: ElementHelper{
			Title:         "Registry",
			Rules:         "",
			Type:          `Max350Text (based on string) minLength: 1 maxLength: 350`,
			Documentation: `Name of the implementation specification registry in which the implementation specification of the ISO 20022 message is maintained. For example, "MyStandards".`,
		},
	}
}

type ContactDetailsHelper struct {
	Name            ElementHelper
	Email           ElementHelper
	Phone           ElementHelper
	PreferredMethod ElementHelper
}

func BuildContactDetailsHelper() ContactDetailsHelper {
	return ContactDetailsHelper{
		Name: ElementHelper{
			Title:         "Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which a party is known and which is usually used to identify that party.`,
		},
		Email: ElementHelper{
			Title:         "Email",
			Rules:         "",
			Type:          `Max2048Text (based on string) minLength: 1 maxLength: 2048`,
			Documentation: `Address for electronic mail (e-mail).`,
		},
		Phone: ElementHelper{
			Title:         "Phone",
			Rules:         "",
			Type:          `PhoneNumber (based on string) pattern: \+[0-9]{1,3}-[0-9()+\-]{1,30}`,
			Documentation: `Collection of information that identifies a phone number, as defined by telecom services.`,
		},
		PreferredMethod: ElementHelper{
			Title:         "Preferred Method",
			Rules:         "",
			Type:          `ContactMethod`,
			Documentation: `Preferred method used to reach the contact.`,
		},
	}
}

type MissingCodeAndInfoHelper struct {
	Code ElementHelper
	Info ElementHelper
}

func BuildMissingCodeAndInfoHelper() MissingCodeAndInfoHelper {
	return MissingCodeAndInfoHelper{
		Code: ElementHelper{
			Title:         "Code",
			Rules:         "",
			Type:          `MissingOrIncorrectInformationCode`,
			Documentation: `Specifies the missing information in a coded form.`,
		},
		Info: ElementHelper{
			Title:         "Information",
			Rules:         "",
			Type:          `Max500Text (based on string) minLength: 1 maxLength: 500`,
			Documentation: `Further details about the incorrect information.`,
		},
	}
}

type CreatorHelper struct {
	Name               ElementHelper
	PaymentSysCode     ElementHelper
	PaymentSysMemberId ElementHelper
	Contact            ContactDetailsHelper
}

func BuildCreatorHelper() CreatorHelper {
	return CreatorHelper{
		Name: ElementHelper{
			Title:         "Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name of the creator of the message.`,
		},
		PaymentSysCode: ElementHelper{
			Title:         "Payment System Code",
			Rules:         "",
			Type:          `PaymentSystemType`,
			Documentation: `Payment system code of the creator.`,
		},
		PaymentSysMemberId: ElementHelper{
			Title:         "Payment System Member ID",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Payment system member ID of the creator.`,
		},
		Contact: BuildContactDetailsHelper(),
	}
}

type PaymentInfomationHelper struct {
	OriginalGroupInfo         GroupInformationHelper
	OriginalTransactionDetail TransactionDetailReferenceHelper
	SettlementAmount          CurrencyAndAmountHelper
	SettlementDate            ElementHelper
}

func BuildPaymentInfomationHelper() PaymentInfomationHelper {
	return PaymentInfomationHelper{
		OriginalGroupInfo:         BuildGroupInformationHelper(),
		OriginalTransactionDetail: BuildTransactionDetailReferenceHelper(),
		SettlementAmount:          BuildCurrencyAndAmountHelper(),
		SettlementDate: ElementHelper{
			Title:         "Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `The date of the original interbank settlement.`,
		},
	}
}
