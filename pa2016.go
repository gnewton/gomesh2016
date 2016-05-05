package gomesh2016

const PHARMACOLOGICAL_RECORD = "PharmacologicalAction"

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
	RecordName string `xml:"RecordName,omitempty" json:"RecordName,omitempty"`
	RecordUI string `xml:"RecordUI,omitempty" json:"RecordUI,omitempty"`
}


