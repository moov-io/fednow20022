package fednow

import (
	"fmt"
	"strings"
	"time"
)

// MessageID formats a FedNow Message Identification which is a reference assigned by the sender of the message.
// A MessageID is composed of
//   - the Calendar Date (8 numeric characters, CCYYMMDD),
//   - sender's FedNow Connection Party Identification (9 alphanumeric characters)
//   - a reference assigned by the sender (up to 18 alphanumeric characters)
func MessageID(when time.Time, partyIdentification string, reference string) string {
	return strings.TrimSpace(fmt.Sprintf("%8.8s%9.9s%-18.18s", when.Format("20060102"), partyIdentification, reference))
}
