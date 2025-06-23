package AccountBalanceReport_camt_052_001_08

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type TotalsPerBankTransactionCodeHelper struct {
	NumberOfEntries     common.ElementHelper
	BankTransactionCode common.ElementHelper
}

func BuildTotalsPerBankTransactionCodeHelper() TotalsPerBankTransactionCodeHelper {
	return TotalsPerBankTransactionCodeHelper{
		NumberOfEntries: common.ElementHelper{
			Title:         "Number of Entries",
			Rules:         "",
			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
			Documentation: `Number of individual entries for the bank transaction code.`,
		},
		BankTransactionCode: common.ElementHelper{
			Title:         "Bank Transaction Code",
			Rules:         "",
			Type:          `TransactionStatusCode(MessagesInProcess, MessagesIntercepted ...)`,
			Documentation: `Bank transaction code in a proprietary form, as defined by the issuer.`,
		},
	}
}

type BalanceHelper struct {
	BalanceTypeId        common.ElementHelper
	CreditLine           common.CreditLineHelper
	Amount               common.CurrencyAndAmountHelper
	CreditDebitIndicator common.ElementHelper
	DateTime             common.ElementHelper
}
func BuildBalanceHelper() BalanceHelper{
	return BalanceHelper{
		BalanceTypeId: common.ElementHelper{
			Title:         "Balance Type Id",
			Rules:         "",
			Type:          `BalanceType(AccountBalance, AvailableBalanceFromAccountBalance ...)`,
			Documentation: `Specifies the nature of a balance.`,
		},
		CreditLine: common.BuildCreditLineHelper(),
		Amount:   common.BuildCurrencyAndAmountHelper(),
		CreditDebitIndicator: common.ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `CdtDbtInd(Credit, Debit)`,
			Documentation: `Indicates whether the balance is a credit or a debit balance.`,
		},
		DateTime: common.ElementHelper{
			Title:         "Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Indicates the date (and time) of the balance.`,
		},
	}
}

type MessageHelper struct {
	MessageId             common.ElementHelper
	CreatedDateTime       common.ElementHelper
	Pagenation            common.MessagePagenationHelper
	OriginalBusinessQuery common.OriginalBusinessQueryHelper // Optional field for original business query
	ReportId              common.ElementHelper
	ReportCreateDateTime  common.ElementHelper
	AccountOtherId        common.ElementHelper
	RelateAccountOtherId  common.ElementHelper
	Balances              BalanceHelper
	TransactionsSummary   common.TotalsPerBankTransactionHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: common.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message. Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.`,
		},
		CreatedDateTime: common.ElementHelper{
			Title:         "Message Identification",
			Rules:         "This is the calendar date and time in New York City (Eastern Time) when the message is created by the Fedwire Funds Service application. Time is in 24-hour clock format and includes the offset against the Coordinated Universal Time (UTC).",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		Pagenation:            common.BuildMessagePagenationHelper(),
		OriginalBusinessQuery: common.BuildOriginalBusinessQueryHelper(), // Optional field for original business query
		ReportId: common.ElementHelper{
			Title:         "Report Type Id",
			Rules:         "",
			Type:          `ReportType(EveryDay, Intraday)`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account report.`,
		},
		ReportCreateDateTime: common.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "This is the Fedwire Funds Service funds-transfer business day. Note: Time will be defaulted to 00:00:00 in New York City (Eastern Time) with the offset against the Coordinated Universal Time (UTC).",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the report was created.`,
		},
		AccountOtherId: common.ElementHelper{
			Title:         "Account Other Id",
			Rules:         "his is the routing number to which the activity report relates.",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification assigned by an institution.`,
		},
		RelateAccountOtherId: common.ElementHelper{
			Title:         "Related Account Other Id",
			Rules:         "",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `This component only occurs in a Correspondent Account Activity Details Report (CADR) to identify the correspondent's respondent for which activity details are reported. The component is not present in a regular Account Activity Details Report (AADR).`,
		},
		Balances: BuildBalanceHelper(),
		TransactionsSummary: common.BuildTotalsPerBankTransactionHelper(),
	}

}
