[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://pkg.go.dev/badge/github.com/moov-io/fednow20022.svg)](https://pkg.go.dev/github.com/moov-io/fednow20022)
[![Build Status](https://github.com/moov-io/fednow20022/workflows/Go/badge.svg)](https://github.com/moov-io/fednow20022/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/fednow20022/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/fednow20022)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/fednow20022)](https://goreportcard.com/report/github.com/moov-io/fednow20022)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/fednow20022?label=project%20size)](https://github.com/moov-io/fednow20022)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/fednow20022/master/LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/fednow20022)](https://github.com/moov-io/fednow20022)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)


## moov-io/fednow20022

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

This repository contains a subset of ISO 20022 messages for FedNow payments. ISO 20022 is a standard for electronic data interchange between financial institutions. It describes a metadata repository containing descriptions of messages and business processes, and a maintenance process for the repository content. The standard covers financial information transferred between financial institutions that includes payment transactions, securities trading and settlement information, credit and debit card transactions, and other financial information.

Refer to the [FedNow Explorer](https://explore.fednow.org/), the [FedNow Explorer resources](https://explore.fednow.org/resources) and the [FedNow Service group on MyStandards](https://www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/FedNow_Service) for more information. The FedNow ISO 20022 specifications are on MyStandards, which requires a login.

## Project Status

Go code is generated with [`moov-io/xsd2go`](https://github.com/moov-io/xsd2go) inside the `gen/` folder. See [Generate the Go code](#generate-the-go-code). Please star the project if you are interested in its progress. Feedback on this early version of ISO 20022 is appreciated and vital to its success. Please let us know if you encounter any bugs/unclear documentation or have feature suggestions by opening up an issue. Thanks!

## Naming Conventions

The XSDs define some naming that can be confusing.

| Go Package                     | Description                                                      |
|--------------------------------|------------------------------------------------------------------|
| `gen/fednow_incoming_external` | Messages sent INTO the FedNOW services - from a bank or vendor.  |
| `gen/fednow_outgoing_external` | Messages coming INTO a bank or vendor - from the FedNOW service. |

## Usage

This project is a Go library only. It does not have a CLI, a server or a Docker image.

### Installation

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.23.3 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go.

```
go get github.com/moov-io/fednow20022@latest
```

### Supported messages

Each message is a Go package in `gen/`. Each package has a `Document` type for the message and the XSD is in `xsd/`. The XSD files were last updated in August 2025.

| Message | Go package | Element in the envelope | Direction |
|---------|------------|-------------------------|-----------|
| admi.002.001.01 | `gen/admi_002_001_01` | `FedNowMessageReject` | Incoming and outgoing |
| admi.004.001.02 | `gen/admi_004_001_02` | `FedNowParticipantBroadcast` (incoming), `FedNowBroadcast` (outgoing) | Incoming and outgoing |
| admi.006.001.01 | `gen/admi_006_001_01` | `FedNowRetrievalRequest` | Incoming |
| admi.007.001.01 | `gen/admi_007_001_01` | `FedNowReceiptAcknowledgement` | Incoming and outgoing |
| admi.011.001.01 | `gen/admi_011_001_01` | `FedNowSystemResponse` | Outgoing |
| admi.998.001.02 | `gen/admi_998_001_02`, `gen/sup_FedNowParticipantFile_admi_998_001_02` | `FedNowParticipantFile` | Outgoing |
| camt.026.001.07 | `gen/camt_026_001_07` | `FedNowInformationRequest` | Incoming and outgoing |
| camt.028.001.09 | `gen/camt_028_001_09` | `FedNowAdditionalPaymentInformation` | Incoming and outgoing |
| camt.029.001.09 | `gen/camt_029_001_09` | `FedNowInformationRequestResponse`, `FedNowReturnRequestResponse`, `FedNowRequestForPaymentCancellationRequestResponse` | Incoming and outgoing |
| camt.052.001.08 | `gen/camt_052_001_08` | `FedNowAccountActivityDetailsReport`, `FedNowAccountActivityTotalsReport`, `FedNowAccountBalanceReport` | Outgoing |
| camt.054.001.08 | `gen/camt_054_001_08` | `AccountDebitCreditNotification` | Outgoing |
| camt.055.001.09 | `gen/camt_055_001_09` | `FedNowRequestForPaymentCancellationRequest` | Incoming and outgoing |
| camt.056.001.08 | `gen/camt_056_001_08` | `FedNowReturnRequest` | Incoming and outgoing |
| camt.060.001.05 | `gen/camt_060_001_05` | `FedNowAccountReportingRequest` | Incoming |
| head.001.001.02 | `gen/head_001_001_02` | `AppHdr` in each message | Incoming and outgoing |
| pacs.002.001.10 | `gen/pacs_002_001_10` | `FedNowPaymentStatus` | Incoming and outgoing |
| pacs.004.001.10 | `gen/pacs_004_001_10` | `FedNowPaymentReturn` | Incoming and outgoing |
| pacs.008.001.08 | `gen/pacs_008_001_08` | `FedNowCustomerCreditTransfer` | Incoming and outgoing |
| pacs.009.001.08 | `gen/pacs_009_001_08` | `FedNowInstitutionCreditTransfer` | Incoming and outgoing |
| pacs.028.001.03 | `gen/pacs_028_001_03` | `FedNowPaymentStatusRequest` | Incoming and outgoing |
| pain.013.001.07 | `gen/pain_013_001_07` | `FedNowRequestForPayment` | Incoming and outgoing |
| pain.014.001.07 | `gen/pain_014_001_07` | `FedNowRequestForPaymentResponse` | Incoming and outgoing |
| FedNow key exchange | `gen/FedNowKeyExchange` | `FedNowIncomingMessageSignatureManagement`, `FedNowOutgoingMessageSignatureManagement` | Incoming and outgoing |

### Envelopes

FedNow messages are sent in an envelope. Most messages have an `AppHdr` (head.001.001.02) and a `Document`. The key exchange messages do not.

- Use `fednow_incoming_external.FedNowIncoming` for a message that you send to the FedNow Service.
- Use `fednow_outgoing_external.FedNowOutgoing` for a message that the FedNow Service sends to you.

Set only one message field in `FedNowIncomingMessage` or `FedNowOutgoingMessage`.

### Quick start

Build a pacs.008 customer credit transfer and encode it:

```go
created := time.Now().In(fednow.Eastern())
msgID := fednow.MessageID(created, "111111111", "MyReference001")

msg := fednow_incoming_external.FedNowIncoming{
	FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
		FedNowCustomerCreditTransfer: &fednow_incoming_external.FedNowCustomerCreditTransfer{
			AppHdr: head_001_001_02.BusinessApplicationHeaderV02{
				BizMsgIdr: head_001_001_02.Max35Text(msgID),
				MsgDefIdr: "pacs.008.001.08",
				CreDt:     fednow.ISODateTime(created),
				// Fr, To, MktPrctc ...
			},
			Document: pacs_008_001_08.Document{
				FIToFICstmrCdtTrf: pacs_008_001_08.FIToFICustomerCreditTransferV08{
					GrpHdr: pacs_008_001_08.GroupHeader93{
						MsgId:   pacs_008_001_08.Max35Text(msgID),
						CreDtTm: fednow.ISODateTime(created),
						NbOfTxs: "1",
						// SttlmInf ...
					},
					CdtTrfTxInf: []pacs_008_001_08.CreditTransferTransaction39{
						// PmtId, IntrBkSttlmAmt, Dbtr, Cdtr ...
					},
				},
			},
		},
	},
}

bs, err := xml.MarshalIndent(msg, "", "  ")
if err != nil {
	log.Fatal(err)
}
```

Parse a pacs.002 payment status report from the FedNow Service:

```go
var msg fednow_outgoing_external.FedNowOutgoing
if err := xml.Unmarshal(bs, &msg); err != nil {
	log.Fatal(err)
}

if status := msg.FedNowOutgoingMessage.FedNowPaymentStatus; status != nil {
	for _, tx := range status.Document.FIToFIPmtStsRpt.TxInfAndSts {
		fmt.Println(*tx.OrgnlEndToEndId, *tx.TxSts) // Scenario01EtoEId001 ACSC
	}
}
if reject := msg.FedNowOutgoingMessage.FedNowMessageReject; reject != nil {
	// admi.002: the element is Document.Admi00200101
	fmt.Println(reject.Document.Admi00200101.Rsn.RjctgPtyRsn) // T505
}
```

The complete examples are in [example_test.go](example_test.go): `Example_buildCustomerCreditTransfer`, `Example_parsePaymentStatus` and `Example_parseMessageReject`. They are tested with `go test`.

Amounts in the generated types are strings, for example `ActiveCurrencyAndAmount{Ccy: "USD", Text: "51.74"}`. Write them with two decimal places.

### Helpers

The `pkg/fednow` package has helpers for FedNow messages:

- `fednow.MessageID(when, partyIdentification, reference)` makes a message identification from the date, your FedNow connection party identification and your reference.
- `fednow.Eastern()` returns the US Eastern time zone.
- `fednow.ISODate` and `fednow.ISODateTime` are the date and date time types in the generated code.

The `pkg/fednowtest` package has `fednowtest.FlipMessageDirection(xml)`. It changes a message from the outgoing envelope to the incoming envelope, or from incoming to outgoing, so that you can use a message that you sent as a message that you receive in tests. It supports only some messages.

### Validation

This library does not validate messages. It does not check the XSD rules or the FedNow business rules. For example, a pacs.008 without `CdtTrfTxInf` encodes with no error.

Validate the XML against the XSD files in `xsd/` before you send it. For example, with `xmllint`:

```
xmllint --noout --schema xsd/fednow-incoming_external.xsd message.xml
xmllint --noout --schema xsd/fednow-outgoing_external.xsd message.xml
```

`TestBuildCustomerCreditTransferValidates` in [example_test.go](example_test.go) validates the quick start message with `xmllint`.

### XML namespaces

Go has several [known limitations around XML namespace prefixes](https://github.com/golang/go/issues/13400). The generated types do not use namespace prefixes:

- Decoding accepts messages with or without namespace prefixes (for example `<p8:Document xmlns:p8="...">`).
- Encoding writes each namespace as a default namespace (`xmlns="..."`), with no prefixes. When you encode a message that you decoded, some elements can get a `xmlns` attribute that is not necessary. The XML is still valid against the XSD.

### Generate the Go code

The code in `gen/` is generated from the XSD files in `xsd/` with [`moov-io/xsd2go`](https://github.com/moov-io/xsd2go) and `scripts/generate.sh`. Do not change the files in `gen/` manually. To generate the code again you need `make`, and then run:

```
make install-generator generate
```

## Getting Help

 channel | info
 ------- | -------
Twitter [@moov](https://twitter.com/moov) | You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/fednow20022/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#fednow`) to have an interactive discussion about the development of the project.

## Contributing

Yes please! Please open an issue or pull request. Moov's [Contributing guide](https://github.com/moov-io/ach/blob/master/CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) also apply to this project.

Run the tests with `go test ./...`. `make check` also runs the Moov linters, and it needs `make` and `wget`.

The [docs](docs/) folder has FedNow documents from the Federal Reserve Banks. They are from 2023 and 2024, so check [FedNow Explorer](https://explore.fednow.org/resources) for the current versions.

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
