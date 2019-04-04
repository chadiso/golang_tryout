package models

import "fmt"

const PROHIBITED = "prohibited"
const COLLECTIONS = "collections"
const RECEIPTS = "receipts"

const FROM_CLIENT = "from_client"
const OBO_CLIENT = "obo_client"
const OBO_CLIENTS_CUSTOMER = "obo_clients_customer"

type Classification struct {
	fundingType string
	fundingMode string
}

func Classify(sender Sender, account AccountClassification, houseAccount AccountClassification) Classification {
	fundingType := identifyFundingType(sender, account, houseAccount)
	fundingMode := identifyFundingMode(fundingType, sender, account, houseAccount)

	return Classification{
		fundingType: fundingType,
		fundingMode: fundingMode,
	}
}

func identifyFundingType(sender Sender, account AccountClassification, houseAccount AccountClassification) string {
	switch {
	case ifNoComplianceRelationship(account, houseAccount) ||
		ifNestedPaymentsWithCollections(sender, account, houseAccount) ||
		ifRegulatedAffiliateReceipts(sender, account, houseAccount):
		return PROHIBITED

	case sender.IsNotAccountHolder() || sender.IsApprovedFundingPartner() || isCorporateCollections(account, houseAccount, sender):
		return COLLECTIONS
	case sender.IsAccountHolder():
		return RECEIPTS
	}

	return ""
}

func identifyFundingMode(fundingType string, sender Sender, account AccountClassification, houseAccount AccountClassification) string {
	if fundingType == PROHIBITED {
		return ""
	}

	switch fundingType {
	case RECEIPTS:
		if account.IsClient() {
			return fmt.Sprintf("%s_%s", fundingType, FROM_CLIENT)
		} else {
			return fmt.Sprintf("%s_%s", fundingType, OBO_CLIENT)
		}

	case COLLECTIONS:
		if isNestedCollections(account, houseAccount, sender) {
			return fmt.Sprintf("%s_%s", fundingType, OBO_CLIENTS_CUSTOMER)
		} else {
			return fmt.Sprintf("%s_%s", fundingType, OBO_CLIENT)
		}
	}

	return ""
}

//////////

func ifNoComplianceRelationship(account AccountClassification, houseAccount AccountClassification) bool {
	return account.IsNonClient() && houseAccount.IsNonClient()
}

func ifNestedPaymentsWithCollections(sender Sender, account AccountClassification, houseAccount AccountClassification) bool {
	var isRegulated bool
	if houseAccount.complianceRelationship != "" {
		isRegulated = houseAccount.IsRegulated()
	} else {
		isRegulated = account.IsRegulated()
	}

	return isRegulated && account.IsClient() &&
		(sender.IsNotAccountHolder() || sender.IsApprovedFundingPartner())
}

func ifRegulatedAffiliateReceipts(sender Sender, account AccountClassification, houseAccount AccountClassification) bool {
	return account.IsClient() && sender.IsAccountHolder() &&
		houseAccount.complianceRelationship != "" && houseAccount.IsNonClient() && houseAccount.IsRegulated()
}

func isNestedCollections(account, houseAccount AccountClassification, sender Sender) bool {
	return account.IsNonClient() && houseAccount.IsClient() && houseAccount.IsRegulated() && sender.IsNotAccountHolder()
}

func isCorporateCollections(account, houseAccount AccountClassification, sender Sender) bool {
	return account.IsNonClient() && sender.IsAccountHolder() &&
		houseAccount.complianceRelationship != "" && houseAccount.IsClient() && houseAccount.IsUnregulated()
}
