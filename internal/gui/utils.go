package gui

import (
	"log/slog"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (c *Config) createCompanyList(
	a fyne.App,
	w fyne.Window,
	contentArea *fyne.Container,
) fyne.CanvasObject {
	companyList := container.NewVBox()
	companies, _ := c.dbListCompany()

	for _, company := range companies {
		// company := company
		nameLabel := widget.NewLabel(company.Name)

		updateBtn := widget.NewButton("update", func() {
			slog.Info("Called: Update company", "id", company.ID)
			form := c.updateCompanyForm(a, w, company)

			backBtn := widget.NewButton("<- Back", func() {
				contentArea.Objects = []fyne.CanvasObject{c.createCompanyList(a, w, contentArea)}
				contentArea.Refresh()
			})

			formWithNav := container.NewBorder(
				container.NewHBox(backBtn),
				nil, nil, nil,
				container.NewScroll(form),
			)

			contentArea.Objects = []fyne.CanvasObject{formWithNav}
			contentArea.Refresh()
		})

		row := container.NewBorder(nil, nil, nil, updateBtn, nameLabel)
		companyList.Add(row)
		// companyList.Add(widget.NewLabel(company.Name))
	}

	if len(companyList.Objects) == 0 {
		companyList.Add(widget.NewLabel("Your companies will appear here..."))
	}

	return companyList
}

func (c *Config) createPaymentDetailList(
	a fyne.App,
	w fyne.Window,
	contentArea *fyne.Container,
) fyne.CanvasObject {
	paymentDetailList := container.NewVBox()
	paymentDetails, err := c.dbListPaymentDetail()
	if err != nil {
		slog.Error("DB: Payment detail list", "msg", err)
	}

	for _, paymentDetail := range paymentDetails {
		nameLabel := widget.NewLabel(paymentDetail.AccHolder)

		updateBtn := widget.NewButton("update", func() {
			slog.Info("Called: Update payment detail", "id", paymentDetail.ID)
			form := c.updatePaymentDetailForm(a, w, paymentDetail)

			backBtn := widget.NewButton("<- Back", func() {
				contentArea.Objects = []fyne.CanvasObject{
					c.createPaymentDetailList(a, w, contentArea),
				}
				contentArea.Refresh()
			})

			formWithNav := container.NewBorder(
				container.NewHBox(backBtn),
				nil, nil, nil,
				container.NewScroll(form),
			)
			contentArea.Objects = []fyne.CanvasObject{formWithNav}
			contentArea.Refresh()
		})
		row := container.NewBorder(nil, nil, nil, updateBtn, nameLabel)
		paymentDetailList.Add(row)
		// paymentDetailList.Add(widget.NewLabel(paymentDetail.AccHolder))
	}

	if len(paymentDetailList.Objects) == 0 {
		paymentDetailList.Add(widget.NewLabel("Your payment details list will appear here..."))
	}

	return paymentDetailList
}

func (c *Config) createItemList(
	a fyne.App,
	w fyne.Window,
	contentArea *fyne.Container,
) fyne.CanvasObject {
	itemList := container.NewVBox()
	items, err := c.dbListItem()
	if err != nil {
		slog.Error("DB: Item list", "msg", err)
	}
	for _, item := range items {
		nameLabel := widget.NewLabel(item.Name)

		updateBtn := widget.NewButton("update", func() {
			slog.Info("Called: Update Item detail", "id", item.ID)
			form := c.updateItemForm(a, w, item)
			backBtn := widget.NewButton("<- Back", func() {
				contentArea.Objects = []fyne.CanvasObject{c.createItemList(a, w, contentArea)}
				contentArea.Refresh()
			})

			formWithNav := container.NewBorder(
				container.NewHBox(backBtn),
				nil, nil, nil,
				container.NewScroll(form),
			)
			contentArea.Objects = []fyne.CanvasObject{formWithNav}
			contentArea.Refresh()
		})
		row := container.NewBorder(nil, nil, nil, updateBtn, nameLabel)
		itemList.Add(row)
		// itemList.Add(widget.NewLabel(item.Name))
	}

	if len(itemList.Objects) == 0 {
		itemList.Add(widget.NewLabel("Your Items list will appear here..."))
	}

	return itemList
}

func (c *Config) createShippingAddressList(
	a fyne.App, w fyne.Window, contentArea *fyne.Container,
) fyne.CanvasObject {
	addrList := container.NewVBox()
	addrs, err := c.dbListShippingAddress()
	if err != nil {
		slog.Error("DB: Shipping address list", "msg", err)
	}

	for _, addr := range addrs {
		nameLabel := widget.NewLabel(addr.Name)

		updateBtn := widget.NewButton("update", func() {
			slog.Info("Called: Update shipping address detail", "id", addr.ID)
			form := c.updateShippingAddressForm(a, w, addr)
			backBtn := widget.NewButton("<- Back", func() {
				contentArea.Objects = []fyne.CanvasObject{
					c.createShippingAddressList(a, w, contentArea),
				}
				contentArea.Refresh()
			})

			formWithNav := container.NewBorder(
				container.NewHBox(backBtn),
				nil, nil, nil,
				container.NewScroll(form),
			)
			contentArea.Objects = []fyne.CanvasObject{formWithNav}
			contentArea.Refresh()
		})
		row := container.NewBorder(nil, nil, nil, updateBtn, nameLabel)
		addrList.Add(row)
	}

	if len(addrList.Objects) == 0 {
		addrList.Add(widget.NewLabel("Your shipping addresses details list will appear here..."))
	}

	return addrList
}

func (c *Config) createInvoiceList() fyne.CanvasObject {
	invoiceList := container.NewVBox()
	invoices, err := c.dbListInvoices()
	if err != nil {
		slog.Error("DB: Invoice list", "msg", err)
	}
	for _, invoice := range invoices {
		invoiceList.Add(widget.NewLabel(strconv.Itoa(int(invoice.InvoiceNumber))))
	}

	if len(invoiceList.Objects) == 0 {
		invoiceList.Add(widget.NewLabel("Your invoices will appear here..."))
	}

	return invoiceList
}
