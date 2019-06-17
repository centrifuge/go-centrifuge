// +build testworld

package testworld

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func defaultDocumentPayload(documentType string, collaborators []string) map[string]interface{} {
	switch documentType {
	case typeInvoice:
		return defaultInvoicePayload(collaborators)
	case typePO:
		return defaultPOPayload(collaborators)
	default:
		return defaultInvoicePayload(collaborators)
	}
}

func defaultPOPayload(collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_created": "2018-09-26T23:12:37.902198664Z",
			"total_amount": "40",
			"currency":     "USD",
		},
		"write_access": collaborators,
		"attributes":   defaultAttributePayload(),
	}
}

func defaultEntityPayload(identity string, collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"identity":   identity,
			"legal_name": "test company",
			"contacts": []map[string]interface{}{
				{
					"name": "test name",
				},
			},
		},
		"write_access": collaborators,
	}
}

func defaultRelationshipPayload(identity, targetID string) map[string]interface{} {
	return map[string]interface{}{
		"identity":        identity,
		"target_identity": targetID,
	}
}

func defaultInvoicePayload(collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_due":     "2018-09-26T23:12:37.902198664Z",
			"gross_amount": "40",
			"currency":     "USD",
			"net_amount":   "40",
			"line_items": []map[string]interface{}{
				{
					"item_number":  "12345",
					"tax_amount":   "1.99",
					"total_amount": "2.99",
					"description":  "line item description",
				},
			},
		},
		"write_access": collaborators,
		"attributes":   defaultAttributePayload(),
	}
}

func defaultFundingPayload(borrowerId, funderId string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"amount":             "20000",
			"apr":                "0.33",
			"days":               "90",
			"currency":           "USD",
			"fee":                "30.30",
			"repayment_due_date": "2018-09-26T23:12:37.902198664Z",
			"borrower_id":        borrowerId,
			"funder_id":          funderId,
		},
	}
}

func updateFundingPayload(agreementId, borrowerId, funderId string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"agreement_id":       agreementId,
			"funder_id":          funderId,
			"borrower_id":        borrowerId,
			"amount":             "10000",
			"apr":                "0.55",
			"days":               "90",
			"currency":           "USD",
			"fee":                "30.30",
			"repayment_due_date": "2018-09-26T23:12:37.902198664Z",
		},
	}
}

func wrongInvoicePayload(collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_due":     "2018-09-26T23:12:37.902198664Z",
			"gross_amount": "40",
			"currency":     "USD",
			"net_amount":   "40",
			"line_items": []map[string]interface{}{
				{
					"item_number":  "12345",
					"tax_amount":   "1.99",
					"total_amount": "2.99",
					"description":  "line item description",
				},
			},
		},
		"write_access": collaborators,
		"attributes":   wrongAttributePayload(),
	}
}

func invoiceNFTPayload(collaborators []string, sender string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"number":        "12324",
			"date_due":      "2018-09-26T23:12:37.902198664Z",
			"gross_amount":  "40",
			"currency":      "USD",
			"net_amount":    "40",
			"document_type": "invoice",
			"sender":        sender,
			"status":        "unpaid",
		},
		"write_access": collaborators,
	}
}

func poNFTPayload(collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"number":        "123245",
			"date_created":  "2018-09-26T23:12:37.902198664Z",
			"currency":      "USD",
			"total_amount":  "40",
			"document_type": "po",
		},
		"write_access": collaborators,
	}
}

func defaultNFTPayload(documentType string, collaborators []string, sender string) map[string]interface{} {
	switch documentType {
	case typeInvoice:
		return invoiceNFTPayload(collaborators, sender)
	case typePO:
		return poNFTPayload(collaborators)
	default:
		return invoiceNFTPayload(collaborators, sender)
	}

}

func updatedDocumentPayload(documentType string, collaborators []string) map[string]interface{} {
	switch documentType {
	case typeInvoice:
		return updatedInvoicePayload(collaborators)
	case typePO:
		return updatedPOPayload(collaborators)
	default:
		return updatedInvoicePayload(collaborators)
	}
}

func updatedPOPayload(collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_created": "2018-09-26T23:12:37.902198664Z",
			"currency":     "EUR",
			"total_amount": "42",
		},
		"write_access": collaborators,
	}

}

func updatedInvoicePayload(collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_due":     "2018-09-26T23:12:37.902198664Z",
			"gross_amount": "40",
			"currency":     "EUR",
			"net_amount":   "42",
		},
		"write_access": collaborators,
	}

}

func updatedEntityPayload(identity string, collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"identity":   identity,
			"legal_name": "edited test company",
			"contacts": []map[string]interface{}{
				{
					"name": "test name",
				},
			},
		},
		"write_access": collaborators,
	}
}

func defaultProofPayload(documentType string) map[string]interface{} {
	if documentType == typeInvoice {

		return map[string]interface{}{
			"type":   "http://github.com/centrifuge/centrifuge-protobufs/invoice/#invoice.InvoiceData",
			"fields": []string{"invoice.net_amount", "invoice.currency"},
		}
	}
	return map[string]interface{}{
		"type":   "http://github.com/centrifuge/centrifuge-protobufs/purchaseorder/#purchaseorder.PurchaseOrderData",
		"fields": []string{"po.total_amount", "po.currency"},
	}
}

func wrongAttributePayload() map[string]map[string]string {
	payload := defaultAttributePayload()
	payload["test_invalid"] = map[string]string{
		"type":  "timestamp",
		"value": "some invalid time stamp",
	}

	return payload
}

func defaultAttributePayload() map[string]map[string]string {
	return map[string]map[string]string{
		"test_string": {
			"type":  "string",
			"value": "string value",
		},

		"test_decimal": {
			"type":  "decimal",
			"value": "100.000001",
		},

		"test_integer": {
			"type":  "integer",
			"value": "123456",
		},

		"test_bytes": {
			"type":  "bytes",
			"value": hexutil.Encode([]byte("byte value")),
		},

		"test_timestamp": {
			"type":  "timestamp",
			"value": "2019-05-03T12:49:05Z",
		},
	}
}
