package setr

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document01400102 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:setr.014.001.02 Document"`
	Message *SwitchOrderCancellationInstructionV02 `xml:"setr.014.001.02"`
}

func (d *Document01400102) AddMessage() *SwitchOrderCancellationInstructionV02 {
	d.Message = new(SwitchOrderCancellationInstructionV02)
	return d.Message
}

// Scope
// The SwitchOrderCancellationInstruction message is sent by an instructing party, eg, an investment manager or its authorised representative, to an executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is sent to cancel a previously sent SwitchOrderInstruction message.
// Usage
// The SwitchOrderCancellationInstruction message is used to cancel the entire previously sent SwitchOrder message and all the individual legs that it contains. There is no amendment, but a cancellation and re-instruct policy.
// This message must contain the reference of the message to be cancelled. This message may also contain all the details of the message to be cancelled, but this is not recommended.
// The deadline and acceptance of a cancellation instruction is subject to a service level agreement (SLA). This cancellation message is a cancellation instruction. There is no automatic acceptance of the cancellation instruction.
// The rejection or acceptance of a cancellation message instruction is made using an OrderCancellationStatusReport message.
type SwitchOrderCancellationInstructionV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *iso20022.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef"`

	// Common information related to all the orders to be cancelled.
	OrderToBeCancelled *iso20022.SwitchOrderInstruction1 `xml:"OrdrToBeCanc,omitempty"`
}

func (s *SwitchOrderCancellationInstructionV02) AddMasterReference() *iso20022.AdditionalReference3 {
	s.MasterReference = new(iso20022.AdditionalReference3)
	return s.MasterReference
}

func (s *SwitchOrderCancellationInstructionV02) AddPoolReference() *iso20022.AdditionalReference3 {
	s.PoolReference = new(iso20022.AdditionalReference3)
	return s.PoolReference
}

func (s *SwitchOrderCancellationInstructionV02) AddPreviousReference() *iso20022.AdditionalReference3 {
	s.PreviousReference = new(iso20022.AdditionalReference3)
	return s.PreviousReference
}

func (s *SwitchOrderCancellationInstructionV02) AddOrderToBeCancelled() *iso20022.SwitchOrderInstruction1 {
	s.OrderToBeCancelled = new(iso20022.SwitchOrderInstruction1)
	return s.OrderToBeCancelled
}
func ( d *Document01400102 ) String() (result string, ok bool) { return }
