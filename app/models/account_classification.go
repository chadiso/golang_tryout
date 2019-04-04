package models

type AccountClassification struct {
	regulatedService       string
	complianceRelationship string
}

func (a AccountClassification) IsClient() bool {
	return a.complianceRelationship == "client"
}

func (a AccountClassification) IsNonClient() bool {
	return a.complianceRelationship == "non-client" || a.complianceRelationship == ""
}

func (a AccountClassification) IsRegulated() bool {
	return a.regulatedService == "regulated"
}

func (a AccountClassification) IsUnregulated() bool {
	return a.regulatedService == "unregulated"
}
