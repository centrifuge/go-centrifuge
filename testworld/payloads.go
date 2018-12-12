// +build testworld

package testworld

func defaultDocumentPayload( documentType string, collaborators []string) map[string]interface{} {

 switch documentType{
 case TypeInvoice: return defaultInvoicePayload(collaborators)
 case TypePO: return defaultPOPayload(collaborators)
 default: return defaultInvoicePayload(collaborators)
 }

}



func defaultPOPayload(collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"po_number": "12324",
			"due_date":       "2018-09-26T23:12:37.902198664Z",
			"gross_amount":   "40",
			"currency":       "USD",
			"net_amount":     "40",
		},
		"collaborators": collaborators,
	}

}
func defaultInvoicePayload(collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"invoice_number": "12324",
			"due_date":       "2018-09-26T23:12:37.902198664Z",
			"gross_amount":   "40",
			"currency":       "USD",
			"net_amount":     "40",
		},
		"collaborators": collaborators,
	}

}

func defaultNFTPayload(collaborators []string) map[string]interface{} {

	return map[string]interface{}{
		"data": map[string]interface{}{
			"invoice_number": "12324",
			"due_date":       "2018-09-26T23:12:37.902198664Z",
			"gross_amount":   "40",
			"currency":       "USD",
			"net_amount":     "40",
			"document_type":  "invoice",
		},
		"collaborators": collaborators,
	}

}

func updatedDocumentPayload( documentType string, collaborators []string) map[string]interface{} {
	switch documentType{
	case TypeInvoice: return updatedInvoicePayload(collaborators)
	case TypePO: return updatedPOPayload(collaborators)
	default: return updatedInvoicePayload(collaborators)
	}
}


func updatedPOPayload(collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"po_number": "12324",
			"due_date":       "2018-09-26T23:12:37.902198664Z",
			"gross_amount":   "40",
			"currency":       "EUR",
			"net_amount":     "42",
		},
		"collaborators": collaborators,
	}

}

func updatedInvoicePayload(collaborators []string) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"invoice_number": "12324",
			"due_date":       "2018-09-26T23:12:37.902198664Z",
			"gross_amount":   "40",
			"currency":       "EUR",
			"net_amount":     "42",
		},
		"collaborators": collaborators,
	}

}

func defaultProofPayload(documentType string) map[string]interface{} {
	if documentType == TypeInvoice {

		return map[string]interface{}{
			"type":   "http://github.com/centrifuge/centrifuge-protobufs/invoice/#invoice.InvoiceData",
			"fields": []string{"invoice.net_amount", "invoice.currency"},
		}
	}
	return map[string]interface{}{
		"type":   "http://github.com/centrifuge/centrifuge-protobufs/purchaseorder/#purchaseorder.PurchaseOrderData",
		"fields": []string{"purchaseorder.net_amount", "purchaseorder.currency"},
	}
}
