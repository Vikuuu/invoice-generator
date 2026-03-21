package generator

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/Vikuuu/invoice_generator/internal/database"
)

type CompanyDetail struct {
	name, gstin, address string
}

type Item struct {
	name  string
	hsn   int
	qty   int
	price int
	gst   int
	total int
}

type Payment struct {
	name, number, ifsc,
	branch, bankName, virtualAddress string
}

type Invoice struct {
	from, to        CompanyDetail
	date            *time.Time
	items           []Item
	shippingAddress string
	pay             Payment
	// subTotal float32
	// igst float32
}

func NewInvoice(
	fromCompany, toCompany database.Company,
	date *time.Time,
	item database.Item,
	qty int,
	shipAddr database.ShippingAddress,
	pay database.PaymentDetail,
) *Invoice {
	accNumber := strconv.Itoa(int(pay.AccNumber))
	return &Invoice{
		from: CompanyDetail{
			name:    fromCompany.Name,
			gstin:   fromCompany.Gst,
			address: fromCompany.Address,
		},
		to: CompanyDetail{
			name:    toCompany.Name,
			gstin:   toCompany.Gst,
			address: toCompany.Address,
		},
		items: []Item{
			{
				name:  item.Name,
				hsn:   int(item.Hsn),
				qty:   qty,
				price: int(item.Price),
			},
		},
		shippingAddress: shipAddr.Address,
		pay: Payment{
			name:           pay.AccHolder,
			number:         accNumber,
			ifsc:           pay.Ifsc,
			branch:         pay.Branch,
			bankName:       pay.BankName,
			virtualAddress: pay.VirtualPaymentAddr,
		},
	}
}

func InvoiceDataMap(
	invoiceNumber int64,
	fromCompany, toCompany database.Company,
	date *time.Time,
	item database.Item,
	qty int,
	shipAddr database.ShippingAddress,
	pay database.PaymentDetail,
) map[string]any {
	typstItems := []map[string]any{
		{
			"product-name": item.Name,
			"hsn-sac":      item.Hsn,
			"qty":          qty,
			"price":        item.Price,
			"gst":          0,
			"total":        0,
		},
	}
	data := map[string]any{
		"company":         fromCompany.Name,
		"company-gstin":   fromCompany.Gst,
		"company-address": fromCompany.Address,
		"invoice-date":    date,
		"invoice-number":  invoiceNumber,
		"bill-to-name":    toCompany.Name,
		"bill-to-gstin":   toCompany.Gst,
		"bill-to-address": toCompany.Address,
		"items":           typstItems,
		"ship-to-address": shipAddr.Address,
		"payment-data": map[string]any{
			"acc-name":        pay.AccHolder,
			"acc-number":      pay.AccNumber,
			"ifsc":            pay.Ifsc,
			"branch":          pay.Branch,
			"bank-name":       pay.BankName,
			"virtual-address": pay.VirtualPaymentAddr,
		},
		"sub-total":  float32(0.00),
		"igst":       float32(0.00),
		"image-path": fromCompany.SignatureImagePath,
	}
	slog.Info("Data to Typst", "data", data)
	return data
}
