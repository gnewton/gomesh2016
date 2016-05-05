
package gomesh2016

type Root struct {
	DescriptorRecordSet *DescriptorRecordSet `xml:"DescriptorRecordSet,omitempty" json:"DescriptorRecordSet,omitempty"`
}

type DescriptorRecordSet struct {
	Attr_LanguageCode string `xml:" LanguageCode,attr"  json:",omitempty"`
	DescriptorRecord []*DescriptorRecord `xml:"DescriptorRecord,omitempty" json:"DescriptorRecord,omitempty"`
}

type DescriptorRecord struct {
	Attr_DescriptorClass string `xml:" DescriptorClass,attr"  json:",omitempty"`
	AllowableQualifiersList *AllowableQualifiersList `xml:"AllowableQualifiersList,omitempty" json:"AllowableQualifiersList,omitempty"`
	Annotation *Annotation `xml:"Annotation,omitempty" json:"Annotation,omitempty"`
	ConceptList *ConceptList `xml:"ConceptList,omitempty" json:"ConceptList,omitempty"`
	ConsiderAlso *ConsiderAlso `xml:"ConsiderAlso,omitempty" json:"ConsiderAlso,omitempty"`
	DateCreated *DateCreated `xml:"DateCreated,omitempty" json:"DateCreated,omitempty"`
	DateEstablished *DateEstablished `xml:"DateEstablished,omitempty" json:"DateEstablished,omitempty"`
	DateRevised *DateRevised `xml:"DateRevised,omitempty" json:"DateRevised,omitempty"`
	DescriptorName *DescriptorName `xml:"DescriptorName,omitempty" json:"DescriptorName,omitempty"`
	DescriptorUI *DescriptorUI `xml:"DescriptorUI,omitempty" json:"DescriptorUI,omitempty"`
	EntryCombinationList *EntryCombinationList `xml:"EntryCombinationList,omitempty" json:"EntryCombinationList,omitempty"`
	HistoryNote *HistoryNote `xml:"HistoryNote,omitempty" json:"HistoryNote,omitempty"`
	NLMClassificationNumber *NLMClassificationNumber `xml:"NLMClassificationNumber,omitempty" json:"NLMClassificationNumber,omitempty"`
	OnlineNote *OnlineNote `xml:"OnlineNote,omitempty" json:"OnlineNote,omitempty"`
	PharmacologicalActionList *PharmacologicalActionList `xml:"PharmacologicalActionList,omitempty" json:"PharmacologicalActionList,omitempty"`
	PreviousIndexingList *PreviousIndexingList `xml:"PreviousIndexingList,omitempty" json:"PreviousIndexingList,omitempty"`
	PublicMeSHNote *PublicMeSHNote `xml:"PublicMeSHNote,omitempty" json:"PublicMeSHNote,omitempty"`
	SeeRelatedList *SeeRelatedList `xml:"SeeRelatedList,omitempty" json:"SeeRelatedList,omitempty"`
	TreeNumberList *TreeNumberList `xml:"TreeNumberList,omitempty" json:"TreeNumberList,omitempty"`
}

type ConceptList struct {
	Concept []*Concept `xml:"Concept,omitempty" json:"Concept,omitempty"`
}

type Concept struct {
	Attr_PreferredConceptYN string `xml:" PreferredConceptYN,attr"  json:",omitempty"`
	CASN1Name *CASN1Name `xml:"CASN1Name,omitempty" json:"CASN1Name,omitempty"`
	ConceptName *ConceptName `xml:"ConceptName,omitempty" json:"ConceptName,omitempty"`
	ConceptRelationList *ConceptRelationList `xml:"ConceptRelationList,omitempty" json:"ConceptRelationList,omitempty"`
	ConceptUI *ConceptUI `xml:"ConceptUI,omitempty" json:"ConceptUI,omitempty"`
	RegistryNumber *RegistryNumber `xml:"RegistryNumber,omitempty" json:"RegistryNumber,omitempty"`
	RelatedRegistryNumberList *RelatedRegistryNumberList `xml:"RelatedRegistryNumberList,omitempty" json:"RelatedRegistryNumberList,omitempty"`
	ScopeNote *ScopeNote `xml:"ScopeNote,omitempty" json:"ScopeNote,omitempty"`
	TermList *TermList `xml:"TermList,omitempty" json:"TermList,omitempty"`
}

type RelatedRegistryNumberList struct {
	RelatedRegistryNumber []*RelatedRegistryNumber `xml:"RelatedRegistryNumber,omitempty" json:"RelatedRegistryNumber,omitempty"`
}

type RelatedRegistryNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ConceptRelationList struct {
	ConceptRelation []*ConceptRelation `xml:"ConceptRelation,omitempty" json:"ConceptRelation,omitempty"`
}

type ConceptRelation struct {
	Attr_RelationName string `xml:" RelationName,attr"  json:",omitempty"`
	Concept1UI *Concept1UI `xml:"Concept1UI,omitempty" json:"Concept1UI,omitempty"`
	Concept2UI *Concept2UI `xml:"Concept2UI,omitempty" json:"Concept2UI,omitempty"`
}

type Concept1UI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Concept2UI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TermList struct {
	Term []*Term `xml:"Term,omitempty" json:"Term,omitempty"`
}

type Term struct {
	Attr_ConceptPreferredTermYN string `xml:" ConceptPreferredTermYN,attr"  json:",omitempty"`
	Attr_IsPermutedTermYN string `xml:" IsPermutedTermYN,attr"  json:",omitempty"`
	Attr_LexicalTag string `xml:" LexicalTag,attr"  json:",omitempty"`
	Attr_RecordPreferredTermYN string `xml:" RecordPreferredTermYN,attr"  json:",omitempty"`
	DateCreated *DateCreated `xml:"DateCreated,omitempty" json:"DateCreated,omitempty"`
	EntryVersion *EntryVersion `xml:"EntryVersion,omitempty" json:"EntryVersion,omitempty"`
	SortVersion *SortVersion `xml:"SortVersion,omitempty" json:"SortVersion,omitempty"`
	String *String `xml:"String,omitempty" json:"String,omitempty"`
	TermUI *TermUI `xml:"TermUI,omitempty" json:"TermUI,omitempty"`
	ThesaurusIDlist *ThesaurusIDlist `xml:"ThesaurusIDlist,omitempty" json:"ThesaurusIDlist,omitempty"`
}

type SortVersion struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TermUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type String struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DateCreated struct {
	Day *Day `xml:"Day,omitempty" json:"Day,omitempty"`
	Month *Month `xml:"Month,omitempty" json:"Month,omitempty"`
	Year *Year `xml:"Year,omitempty" json:"Year,omitempty"`
}

type Year struct {
	Text int16 `xml:",chardata" json:",omitempty"`
}

type Month struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

type Day struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

type ThesaurusIDlist struct {
	ThesaurusID []*ThesaurusID `xml:"ThesaurusID,omitempty" json:"ThesaurusID,omitempty"`
}

type ThesaurusID struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EntryVersion struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ConceptUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ConceptName struct {
	String *String `xml:"String,omitempty" json:"String,omitempty"`
}

type CASN1Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RegistryNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ScopeNote struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HistoryNote struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PharmacologicalActionList struct {
	PharmacologicalAction []*PharmacologicalAction `xml:"PharmacologicalAction,omitempty" json:"PharmacologicalAction,omitempty"`
}

type PharmacologicalAction struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
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

type NLMClassificationNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AllowableQualifiersList struct {
	AllowableQualifier []*AllowableQualifier `xml:"AllowableQualifier,omitempty" json:"AllowableQualifier,omitempty"`
}

type AllowableQualifier struct {
	Abbreviation *Abbreviation `xml:"Abbreviation,omitempty" json:"Abbreviation,omitempty"`
	QualifierReferredTo *QualifierReferredTo `xml:"QualifierReferredTo,omitempty" json:"QualifierReferredTo,omitempty"`
}

type QualifierReferredTo struct {
	QualifierName *QualifierName `xml:"QualifierName,omitempty" json:"QualifierName,omitempty"`
	QualifierUI *QualifierUI `xml:"QualifierUI,omitempty" json:"QualifierUI,omitempty"`
}

type QualifierUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type QualifierName struct {
	String *String `xml:"String,omitempty" json:"String,omitempty"`
}

type Abbreviation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type OnlineNote struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PreviousIndexingList struct {
	PreviousIndexing []*PreviousIndexing `xml:"PreviousIndexing,omitempty" json:"PreviousIndexing,omitempty"`
}

type PreviousIndexing struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TreeNumberList struct {
	TreeNumber []*TreeNumber `xml:"TreeNumber,omitempty" json:"TreeNumber,omitempty"`
}

type TreeNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Annotation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EntryCombinationList struct {
	EntryCombination []*EntryCombination `xml:"EntryCombination,omitempty" json:"EntryCombination,omitempty"`
}

type EntryCombination struct {
	ECIN *ECIN `xml:"ECIN,omitempty" json:"ECIN,omitempty"`
	ECOUT *ECOUT `xml:"ECOUT,omitempty" json:"ECOUT,omitempty"`
}

type ECOUT struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
	QualifierReferredTo *QualifierReferredTo `xml:"QualifierReferredTo,omitempty" json:"QualifierReferredTo,omitempty"`
}

type ECIN struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
	QualifierReferredTo *QualifierReferredTo `xml:"QualifierReferredTo,omitempty" json:"QualifierReferredTo,omitempty"`
}

type DateRevised struct {
	Day *Day `xml:"Day,omitempty" json:"Day,omitempty"`
	Month *Month `xml:"Month,omitempty" json:"Month,omitempty"`
	Year *Year `xml:"Year,omitempty" json:"Year,omitempty"`
}

type DateEstablished struct {
	Day *Day `xml:"Day,omitempty" json:"Day,omitempty"`
	Month *Month `xml:"Month,omitempty" json:"Month,omitempty"`
	Year *Year `xml:"Year,omitempty" json:"Year,omitempty"`
}

type ConsiderAlso struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PublicMeSHNote struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SeeRelatedList struct {
	SeeRelatedDescriptor []*SeeRelatedDescriptor `xml:"SeeRelatedDescriptor,omitempty" json:"SeeRelatedDescriptor,omitempty"`
}

type SeeRelatedDescriptor struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
}

