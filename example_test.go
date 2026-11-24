package fednow20022_test

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/fednow20022/gen/fednow_incoming_external"
	"github.com/moov-io/fednow20022/gen/fednow_outgoing_external"
	"github.com/moov-io/fednow20022/gen/head_001_001_02"
	"github.com/moov-io/fednow20022/gen/pacs_008_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow"
)

func ptr[T any](v T) *T {
	return &v
}

// buildCustomerCreditTransfer returns a pacs.008 customer credit transfer in
// the FedNowIncoming envelope, which is the envelope for messages that a
// participant sends to the FedNow Service.
func buildCustomerCreditTransfer() fednow_incoming_external.FedNowIncoming {
	created := time.Date(2023, time.June, 4, 10, 55, 26, 0, fednow.Eastern())
	msgID := fednow.MessageID(created, "111111111", "Sc01Step1MsgId")

	// Agents in the pacs.008 use the ABA routing number with the USABA code.
	agent := func(routingNumber string) pacs_008_001_08.BranchAndFinancialInstitutionIdentification6 {
		return pacs_008_001_08.BranchAndFinancialInstitutionIdentification6{
			FinInstnId: pacs_008_001_08.FinancialInstitutionIdentification18{
				ClrSysMmbId: &pacs_008_001_08.ClearingSystemMemberIdentification2{
					ClrSysId: &pacs_008_001_08.ClearingSystemIdentification2Choice{
						Cd: ptr(pacs_008_001_08.ExternalClearingSystemIdentification1Code("USABA")),
					},
					MmbId: pacs_008_001_08.Max35Text(routingNumber),
				},
			},
		}
	}
	account := func(number string) *pacs_008_001_08.CashAccount38 {
		return &pacs_008_001_08.CashAccount38{
			Id: pacs_008_001_08.AccountIdentification4Choice{
				Othr: &pacs_008_001_08.GenericAccountIdentification1{
					Id: pacs_008_001_08.Max34Text(number),
				},
			},
		}
	}
	// The AppHdr uses the routing number with no clearing system code.
	party := func(routingNumber string) head_001_001_02.Party44Choice {
		return head_001_001_02.Party44Choice{
			FIId: &head_001_001_02.BranchAndFinancialInstitutionIdentification6{
				FinInstnId: head_001_001_02.FinancialInstitutionIdentification18{
					ClrSysMmbId: &head_001_001_02.ClearingSystemMemberIdentification2{
						MmbId: head_001_001_02.Max35Text(routingNumber),
					},
				},
			},
		}
	}

	return fednow_incoming_external.FedNowIncoming{
		FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
			FedNowCustomerCreditTransfer: &fednow_incoming_external.FedNowCustomerCreditTransfer{
				AppHdr: head_001_001_02.BusinessApplicationHeaderV02{
					Fr:        party("111111111"),
					To:        party("021150706"),
					BizMsgIdr: head_001_001_02.Max35Text(msgID),
					MsgDefIdr: "pacs.008.001.08",
					MktPrctc: &head_001_001_02.ImplementationSpecification1{
						Regy: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/FedNow_Service",
						Id:   "frb.fednow.01",
					},
					CreDt: fednow.ISODateTime(created),
				},
				Document: pacs_008_001_08.Document{
					FIToFICstmrCdtTrf: pacs_008_001_08.FIToFICustomerCreditTransferV08{
						GrpHdr: pacs_008_001_08.GroupHeader93{
							MsgId:   pacs_008_001_08.Max35Text(msgID),
							CreDtTm: fednow.ISODateTime(created),
							NbOfTxs: "1",
							SttlmInf: pacs_008_001_08.SettlementInstruction7{
								SttlmMtd: "CLRG",
								ClrSys: &pacs_008_001_08.ClearingSystemIdentification3Choice{
									Cd: ptr(pacs_008_001_08.ExternalCashClearingSystem1Code("FDN")),
								},
							},
						},
						CdtTrfTxInf: []pacs_008_001_08.CreditTransferTransaction39{
							{
								PmtId: pacs_008_001_08.PaymentIdentification7{
									InstrId:    ptr(pacs_008_001_08.Max35Text("Scenario01InstrId001")),
									EndToEndId: "Scenario01EtoEId001",
									UETR:       ptr(pacs_008_001_08.UUIDv4Identifier("8a562c67-ca16-48ba-b074-65581be6f011")),
								},
								PmtTpInf: &pacs_008_001_08.PaymentTypeInformation28{
									LclInstrm: &pacs_008_001_08.LocalInstrument2Choice{
										Prtry: ptr(pacs_008_001_08.Max35Text("FDNA")),
									},
									CtgyPurp: &pacs_008_001_08.CategoryPurpose1Choice{
										Prtry: ptr(pacs_008_001_08.Max35Text("CONS")),
									},
								},
								// Amounts are strings. Write them with two decimal places.
								IntrBkSttlmAmt: pacs_008_001_08.ActiveCurrencyAndAmount{
									Ccy:  "USD",
									Text: "51.74",
								},
								IntrBkSttlmDt: ptr(fednow.ISODate(created)),
								ChrgBr:        "SLEV",
								InstgAgt:      ptr(agent("111111111")),
								InstdAgt:      ptr(agent("333333333")),
								Dbtr: pacs_008_001_08.PartyIdentification135{
									Nm: ptr(pacs_008_001_08.Max140Text("Individual A")),
								},
								DbtrAcct: account("44444444444"),
								DbtrAgt:  agent("111111111"),
								CdtrAgt:  agent("333333333"),
								Cdtr: pacs_008_001_08.PartyIdentification135{
									Nm: ptr(pacs_008_001_08.Max140Text("Corporation B")),
								},
								CdtrAcct: account("55555555555"),
							},
						},
					},
				},
			},
		},
	}
}

// Build a pacs.008 customer credit transfer and encode it as XML.
func Example_buildCustomerCreditTransfer() {
	msg := buildCustomerCreditTransfer()

	bs, err := xml.MarshalIndent(msg, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// The library does not validate the message. Validate the XML against
	// xsd/fednow-incoming_external.xsd before you send it.
	var decoded fednow_incoming_external.FedNowIncoming
	if err := xml.Unmarshal(bs, &decoded); err != nil {
		log.Fatal(err)
	}
	transfer := decoded.FedNowIncomingMessage.FedNowCustomerCreditTransfer.Document.FIToFICstmrCdtTrf
	fmt.Println(transfer.GrpHdr.MsgId)
	fmt.Println(transfer.CdtTrfTxInf[0].IntrBkSttlmAmt.Ccy, transfer.CdtTrfTxInf[0].IntrBkSttlmAmt.Text)

	// Output:
	// 20230604111111111Sc01Step1MsgId
	// USD 51.74
}

// Parse a pacs.002 payment status report that the FedNow Service sends to a
// participant in the FedNowOutgoing envelope.
func Example_parsePaymentStatus() {
	bs, err := os.ReadFile(filepath.Join("testdata", "fednow-sample-message-outgoing-pacs002.xml"))
	if err != nil {
		log.Fatal(err)
	}

	var msg fednow_outgoing_external.FedNowOutgoing
	if err := xml.Unmarshal(bs, &msg); err != nil {
		log.Fatal(err)
	}

	status := msg.FedNowOutgoingMessage.FedNowPaymentStatus
	if status == nil {
		log.Fatal("not a pacs.002 message")
	}
	report := status.Document.FIToFIPmtStsRpt
	for _, tx := range report.TxInfAndSts {
		fmt.Println(*tx.OrgnlEndToEndId, *tx.TxSts)
	}

	// Output:
	// Scenario01EtoEId001 ACSC
}

// Parse an admi.002 message reject that the FedNow Service sends when it
// cannot process a message.
func Example_parseMessageReject() {
	bs, err := os.ReadFile(filepath.Join("testdata", "fednow-sample-message-admi002.xml"))
	if err != nil {
		log.Fatal(err)
	}

	var msg fednow_outgoing_external.FedNowOutgoing
	if err := xml.Unmarshal(bs, &msg); err != nil {
		log.Fatal(err)
	}

	reject := msg.FedNowOutgoingMessage.FedNowMessageReject
	if reject == nil {
		log.Fatal("not an admi.002 message")
	}
	rsn := reject.Document.Admi00200101.Rsn
	fmt.Println(reject.Document.Admi00200101.RltdRef.Ref, rsn.RjctgPtyRsn)

	// Output:
	// NOTAVAILABLE T505
}

// TestBuildCustomerCreditTransferValidates checks the XML of
// Example_buildCustomerCreditTransfer against the bundled XSD with xmllint.
func TestBuildCustomerCreditTransferValidates(t *testing.T) {
	xmllint, err := exec.LookPath("xmllint")
	if err != nil {
		t.Skip("xmllint is not installed")
	}

	bs, err := xml.MarshalIndent(buildCustomerCreditTransfer(), "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "pacs008.xml")
	if err := os.WriteFile(path, bs, 0600); err != nil {
		t.Fatal(err)
	}

	out, err := exec.Command(xmllint, "--noout", "--schema", filepath.Join("xsd", "fednow-incoming_external.xsd"), path).CombinedOutput()
	if err != nil {
		t.Fatalf("xmllint: %v\n%s", err, out)
	}
}
