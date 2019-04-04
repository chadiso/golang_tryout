package models

type Sender struct {
	classification string
	status         string
}

func (s Sender) IsAccountHolder() bool {
	return s.classification == "account_holder"
}

func (s Sender) IsNotAccountHolder() bool {
	return s.classification == "not_account_holder"
}

func (s Sender) IsApprovedFundingPartner() bool {
	return s.classification == "approved_funding_partner"
}

func (s Sender) IsUnknown() bool {
	return s.classification == "unknown"
}

func (s Sender) IfScreeningRequired() bool {
	list := []string{"unchecked", "always_check", "compliance_review_required"}
	return isIncluded(list, s.status) && !s.IsApprovedFundingPartner()
}

func isIncluded(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
