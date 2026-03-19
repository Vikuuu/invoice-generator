package gui

import (
	"log/slog"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (c *Config) createCompanyList() fyne.CanvasObject {
	companyList := container.NewVBox()
	companies, _ := c.dbListCompany()
	for _, company := range companies {
		companyList.Add(widget.NewLabel(company.Name))
	}

	if len(companyList.Objects) == 0 {
		companyList.Add(widget.NewLabel("Your companies will appear here..."))
	}

	return companyList
}

func (c *Config) createPaymentDetailList() fyne.CanvasObject {
	paymentDetailList := container.NewVBox()
	paymentDetails, _ := c.dbListPaymentDetail()
	for _, paymentDetail := range paymentDetails {
		paymentDetailList.Add(widget.NewLabel(paymentDetail.AccHolder))
	}

	if len(paymentDetailList.Objects) == 0 {
		paymentDetailList.Add(widget.NewLabel("Your payment details list will appear here..."))
	}

	return paymentDetailList
}

func (c *Config) createItemList() fyne.CanvasObject {
	itemList := container.NewVBox()
	items, err := c.dbListItem()
	if err != nil {
		slog.Error("Item List", "Error", err)
	}
	for _, item := range items {
		itemList.Add(widget.NewLabel(item.Name))
	}

	if len(itemList.Objects) == 0 {
		itemList.Add(widget.NewLabel("Your Items list will appear here..."))
	}

	return itemList
}

func (c *Config) createShippingAddressList() fyne.CanvasObject {
	addrList := container.NewVBox()
	addrs, _ := c.dbListShippingAddress()
	for _, addr := range addrs {
		addrList.Add(widget.NewLabel(addr.Address))
	}

	if len(addrList.Objects) == 0 {
		addrList.Add(widget.NewLabel("Your shipping addresses details list will appear here..."))
	}

	return addrList
}

func (c *Config) createInvoiceList() fyne.CanvasObject {
	invoiceList := container.NewVBox()
	invoices, _ := c.dbListInvoices()
	for _, invoice := range invoices {
		invoiceList.Add(widget.NewLabel(strconv.Itoa(int(invoice.InvoiceNumber))))
	}

	if len(invoiceList.Objects) == 0 {
		invoiceList.Add(widget.NewLabel("Your invoices will appear here..."))
	}

	return invoiceList
}
