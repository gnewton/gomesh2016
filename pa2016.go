package gomesh2016


type Root struct {
	PharmacologicalActionSet *PharmacologicalActionSet `xml:"PharmacologicalActionSet,omitempty" json:"PharmacologicalActionSet,omitempty"`
}

type PharmacologicalActionSet struct {
	PharmacologicalAction []*PharmacologicalAction `xml:"PharmacologicalAction,omitempty" json:"PharmacologicalAction,omitempty"`
}

type PharmacologicalAction struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
	PharmacologicalActionSubstanceList *PharmacologicalActionSubstanceList `xml:"PharmacologicalActionSubstanceList,omitempty" json:"PharmacologicalActionSubstanceList,omitempty"`
}

type DescriptorReferredTo struct {
	DescriptorName *DescriptorName `xml:"DescriptorName,omitempty" json:"DescriptorName,omitempty"`
	DescriptorUI *DescriptorUI `xml:"DescriptorUI,omitempty" json:"DescriptorUI,omitempty"`
}

type DescriptorUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DescriptorName struct {
	String *String `xml:"String,omitempty" json:"String,omitempty"`
}

type String struct {
	Text string `xml:",chardata" json:",omitempty"`
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

