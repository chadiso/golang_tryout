package models

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

var testData = [][]string{
	// regulated_service account house_account sender funding_type funding_mode

	// receipts
	[]string{"unregulated", "client", "client", "account_holder", "receipts", "receipts_from_client"},
	[]string{"unregulated", "client", "non-client", "account_holder", "receipts", "receipts_from_client"},
	[]string{"regulated", "client", "client", "account_holder", "receipts", "receipts_from_client"},
	[]string{"regulated", "non-client", "client", "account_holder", "receipts", "receipts_obo_client"},

	// collections
	[]string{"unregulated", "non-client", "client", "account_holder", "collections", "collections_obo_client"},
	[]string{"unregulated", "client", "client", "not_account_holder", "collections", "collections_obo_client"},
	[]string{"unregulated", "client", "non-client", "not_account_holder", "collections", "collections_obo_client"},
	[]string{"unregulated", "non-client", "client", "not_account_holder", "collections", "collections_obo_client"},
	[]string{"regulated", "non-client", "client", "not_account_holder", "collections", "collections_obo_clients_customer"},

	// prohibited
	[]string{"regulated", "non-client", "non-client", "account_holder", "prohibited", ""},
	[]string{"unregulated", "non-client", "non-client", "account_holder", "prohibited", ""},
	[]string{"regulated", "client", "non-client", "account_holder", "prohibited", ""},
	[]string{"regulated", "non-client", "non-client", "not_account_holder", "prohibited", ""},
	[]string{"unregulated", "non-client", "non-client", "not_account_holder", "prohibited", ""},
	[]string{"regulated", "client", "client", "not_account_holder", "prohibited", ""},
	[]string{"regulated", "client", "non-client", "not_account_holder", "prohibited", ""},
}

func TestClassify(t *testing.T) {
	for _, row := range testData {
		account := AccountClassification{regulatedService: row[0], complianceRelationship: row[1]}
		houseAccount := AccountClassification{regulatedService: row[0], complianceRelationship: row[2]}
		sender := Sender{classification: row[3], status: "unchecked"}

		expectedClassification := Classification{
			fundingType: row[4],
			fundingMode: row[5],
		}

		classification := Classify(sender, account, houseAccount)

		Equal(t, classification, expectedClassification)
	}
}
