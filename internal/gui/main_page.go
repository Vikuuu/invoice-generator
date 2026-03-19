package gui

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (c *Config) GreatingPage(a fyne.App, w fyne.Window) {
	sb := c.sidebar(a, w)
	greetingPage := widget.NewLabel("Yokosho watashino invoice generator")

	c.Cont = container.NewHSplit(sb, greetingPage)
	c.Cont.Offset = 0.2

	w.SetContent(c.Cont)
}

func (c *Config) CompanyPage(a fyne.App, w fyne.Window) fyne.CanvasObject {
	contentArea := container.NewStack(c.createCompanyList())

	title := widget.NewLabel("Company Page")
	title.TextStyle = fyne.TextStyle{Bold: true}

	btn := widget.NewButton("Add New Company", func() {
		slog.Info("Called: Add new company")
		form := c.addNewCompany(a, w)
		backBtn := widget.NewButton("<- Back", func() {
			contentArea.Objects = []fyne.CanvasObject{c.createCompanyList()}
			contentArea.Refresh()
		})

		formWithNav := container.NewBorder(
			container.NewHBox(backBtn), // Top
			nil, nil, nil,
			container.NewScroll(form), // Center - scrollable
		)

		contentArea.Objects = []fyne.CanvasObject{formWithNav}
		contentArea.Refresh()
	})

	// Header with title on left, button on right
	header := container.NewBorder(nil, nil, title, btn)

	return container.NewBorder(header, nil, nil, nil, contentArea)
}

func (c *Config) PaymentDetailPage(a fyne.App, w fyne.Window) fyne.CanvasObject {
	contentArea := container.NewStack(c.createPaymentDetailList())

	title := widget.NewLabel("Payment Detail Page")
	title.TextStyle = fyne.TextStyle{Bold: true}

	btn := widget.NewButton("Add Payment Detail", func() {
		slog.Info("Called: Add payment detail")
		form := c.addNewPaymentMethod(a, w)
		backBtn := widget.NewButton("<- Back", func() {
			contentArea.Objects = []fyne.CanvasObject{c.createPaymentDetailList()}
			contentArea.Refresh()
		})

		formWithNav := container.NewBorder(
			container.NewHBox(backBtn), // Top
			nil, nil, nil,
			container.NewScroll(form), // Center - scrollable
		)

		contentArea.Objects = []fyne.CanvasObject{formWithNav}
		contentArea.Refresh()
	})

	// Header with title on left, button on right
	header := container.NewBorder(nil, nil, title, btn)

	return container.NewBorder(header, nil, nil, nil, contentArea)
}

func (c *Config) ItemPage(a fyne.App, w fyne.Window) fyne.CanvasObject {
	contentArea := container.NewStack(c.createItemList())

	title := widget.NewLabel("Items Page")
	title.TextStyle = fyne.TextStyle{Bold: true}

	btn := widget.NewButton("Add Items", func() {
		slog.Info("Called: Add Items")
		form := c.addNewItems(a, w)
		backBtn := widget.NewButton("<- Back", func() {
			contentArea.Objects = []fyne.CanvasObject{c.createItemList()}
			contentArea.Refresh()
		})

		formWithNav := container.NewBorder(
			container.NewHBox(backBtn), // Top
			nil, nil, nil,
			container.NewScroll(form), // Center - scrollable
		)

		contentArea.Objects = []fyne.CanvasObject{formWithNav}
		contentArea.Refresh()
	})

	// Header with title on left, button on right
	header := container.NewBorder(nil, nil, title, btn)

	return container.NewBorder(header, nil, nil, nil, contentArea)
}

func (c *Config) ShippingAddressPage(a fyne.App, w fyne.Window) fyne.CanvasObject {
	contentArea := container.NewStack(c.createShippingAddressList())

	title := widget.NewLabel("Shipping Address Page")
	title.TextStyle = fyne.TextStyle{Bold: true}

	btn := widget.NewButton("Add Shipping Address", func() {
		slog.Info("Called: Add Shipping Address")
		form := c.addShippingAddress(a, w)
		backBtn := widget.NewButton("<- Back", func() {
			contentArea.Objects = []fyne.CanvasObject{c.createShippingAddressList()}
			contentArea.Refresh()
		})

		formWithNav := container.NewBorder(
			container.NewHBox(backBtn), // Top
			nil, nil, nil,
			container.NewScroll(form), // Center - scrollable
		)

		contentArea.Objects = []fyne.CanvasObject{formWithNav}
		contentArea.Refresh()
	})

	// Header with title on left, button on right
	header := container.NewBorder(nil, nil, title, btn)

	return container.NewBorder(header, nil, nil, nil, contentArea)
}

func (c *Config) InvoicePage(a fyne.App, w fyne.Window) fyne.CanvasObject {
	contentArea := container.NewStack(c.createInvoiceList())

	title := widget.NewLabel("Invoice Page")
	title.TextStyle = fyne.TextStyle{Bold: true}

	btn := widget.NewButton("New Invoice", func() {
		slog.Info("GUI:", "msg", "Clicked New Invoice button")
		form := c.generateInvoicePage(a, w)
		backBtn := widget.NewButton("<- Back", func() {
			contentArea.Objects = []fyne.CanvasObject{}
			contentArea.Refresh()
		})

		formWithNav := container.NewBorder(
			container.NewHBox(backBtn), // Top
			nil, nil, nil,
			container.NewScroll(form),
		)

		contentArea.Objects = []fyne.CanvasObject{formWithNav}
		contentArea.Refresh()
	})

	header := container.NewBorder(nil, nil, title, btn)

	return container.NewBorder(header, nil, nil, nil, contentArea)
}
