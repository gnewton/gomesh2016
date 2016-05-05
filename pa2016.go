package gomesh2016


type PaRoot struct {
	PharmacologicalActionSet *PharmacologicalActionSet `xml:"PharmacologicalActionSet,omitempty" json:"PharmacologicalActionSet,omitempty"`
}

type PharmacologicalActionSet struct {
	PharmacologicalAction []*PharmacologicalAction `xml:"PharmacologicalAction,omitempty" json:"PharmacologicalAction,omitempty"`
}

type PharmacologicalAction struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
	PharmacologicalActionSubstanceList *PharmacologicalActionSubstanceList `xml:"PharmacologicalActionSubstanceList,omitempty" json:"PharmacologicalActionSubstanceList,omitempty"`
}

type PharmacologicalActionSubstanceList struct {
	Substance []*Substance `xml:"Substance,omitempty" json:"Substance,omitempty"`
}

type Substance struct {
	RecordName *RecordName `xml:"RecordName,omitempty" json:"RecordName,omitempty"`
	RecordUI *RecordUI `xml:"RecordUI,omitempty" json:"RecordUI,omitempty"`
}

type RecordUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RecordName struct {
	String *String `xml:"String,omitempty" json:"String,omitempty"`
}

