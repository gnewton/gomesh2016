type Root struct {
	SupplementalRecordSet *SupplementalRecordSet `xml:"SupplementalRecordSet,omitempty" json:"SupplementalRecordSet,omitempty"`
}

type SupplementalRecordSet struct {
	Attr_LanguageCode string `xml:" LanguageCode,attr"  json:",omitempty"`
	SupplementalRecord []*SupplementalRecord `xml:"SupplementalRecord,omitempty" json:"SupplementalRecord,omitempty"`
}

type SupplementalRecord struct {
	Attr_SCRClass string `xml:" SCRClass,attr"  json:",omitempty"`
	ConceptList *ConceptList `xml:"ConceptList,omitempty" json:"ConceptList,omitempty"`
	DateCreated *DateCreated `xml:"DateCreated,omitempty" json:"DateCreated,omitempty"`
	DateRevised *DateRevised `xml:"DateRevised,omitempty" json:"DateRevised,omitempty"`
	Frequency *Frequency `xml:"Frequency,omitempty" json:"Frequency,omitempty"`
	HeadingMappedToList *HeadingMappedToList `xml:"HeadingMappedToList,omitempty" json:"HeadingMappedToList,omitempty"`
	IndexingInformationList *IndexingInformationList `xml:"IndexingInformationList,omitempty" json:"IndexingInformationList,omitempty"`
	Note *Note `xml:"Note,omitempty" json:"Note,omitempty"`
	PharmacologicalActionList *PharmacologicalActionList `xml:"PharmacologicalActionList,omitempty" json:"PharmacologicalActionList,omitempty"`
	PreviousIndexingList *PreviousIndexingList `xml:"PreviousIndexingList,omitempty" json:"PreviousIndexingList,omitempty"`
	SourceList *SourceList `xml:"SourceList,omitempty" json:"SourceList,omitempty"`
	SupplementalRecordName *SupplementalRecordName `xml:"SupplementalRecordName,omitempty" json:"SupplementalRecordName,omitempty"`
	SupplementalRecordUI *SupplementalRecordUI `xml:"SupplementalRecordUI,omitempty" json:"SupplementalRecordUI,omitempty"`
}

type Note struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Frequency struct {
	Text int16 `xml:",chardata" json:",omitempty"`
}

type SourceList struct {
	Source []*Source `xml:"Source,omitempty" json:"Source,omitempty"`
}

type Source struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IndexingInformationList struct {
	IndexingInformation []*IndexingInformation `xml:"IndexingInformation,omitempty" json:"IndexingInformation,omitempty"`
}

type IndexingInformation struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
	QualifierReferredTo *QualifierReferredTo `xml:"QualifierReferredTo,omitempty" json:"QualifierReferredTo,omitempty"`
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

type PreviousIndexingList struct {
	PreviousIndexing []*PreviousIndexing `xml:"PreviousIndexing,omitempty" json:"PreviousIndexing,omitempty"`
}

type PreviousIndexing struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PharmacologicalActionList struct {
	PharmacologicalAction []*PharmacologicalAction `xml:"PharmacologicalAction,omitempty" json:"PharmacologicalAction,omitempty"`
}

type PharmacologicalAction struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
}

type SupplementalRecordName struct {
	String *String `xml:"String,omitempty" json:"String,omitempty"`
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

type HeadingMappedToList struct {
	HeadingMappedTo []*HeadingMappedTo `xml:"HeadingMappedTo,omitempty" json:"HeadingMappedTo,omitempty"`
}

type HeadingMappedTo struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
	QualifierReferredTo *QualifierReferredTo `xml:"QualifierReferredTo,omitempty" json:"QualifierReferredTo,omitempty"`
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
	TermList *TermList `xml:"TermList,omitempty" json:"TermList,omitempty"`
}

type ConceptName struct {
	String *String `xml:"String,omitempty" json:"String,omitempty"`
}

type RegistryNumber struct {
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
	String *String `xml:"String,omitempty" json:"String,omitempty"`
	TermUI *TermUI `xml:"TermUI,omitempty" json:"TermUI,omitempty"`
	ThesaurusIDlist *ThesaurusIDlist `xml:"ThesaurusIDlist,omitempty" json:"ThesaurusIDlist,omitempty"`
}

type ThesaurusIDlist struct {
	ThesaurusID []*ThesaurusID `xml:"ThesaurusID,omitempty" json:"ThesaurusID,omitempty"`
}

type ThesaurusID struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TermUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CASN1Name struct {
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

type Concept2UI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Concept1UI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RelatedRegistryNumberList struct {
	RelatedRegistryNumber []*RelatedRegistryNumber `xml:"RelatedRegistryNumber,omitempty" json:"RelatedRegistryNumber,omitempty"`
}

type RelatedRegistryNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ConceptUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SupplementalRecordUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DateRevised struct {
	Day *Day `xml:"Day,omitempty" json:"Day,omitempty"`
	Month *Month `xml:"Month,omitempty" json:"Month,omitempty"`
	Year *Year `xml:"Year,omitempty" json:"Year,omitempty"`
}

