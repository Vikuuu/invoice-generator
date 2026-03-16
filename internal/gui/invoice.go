package gui

import (
	"log/slog"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/Vikuuu/invoice_generator/internal/database"
	"github.com/Vikuuu/invoice_generator/internal/generator"
)

// TO-DO: Add option for the invoice_number
func (c *Config) generateInvoicePage(a fyne.App, w fyne.Window) *widget.Form {
	slog.Info("GUI:", "msg", "Generate invoice page")

	invoiceNumber := widget.NewEntry()
	invNum, err := c.dbGetLatestInvoiceNumber()
	if err != nil {
		slog.Error("DB:", "msg", err)
		invNum = 1
	}
	invoiceNumber.Text = strconv.Itoa(int(invNum))

	companies, err := c.dbListCompany()
	if err != nil {
		slog.Error("DB:", "msg", err)
	}
	companiesMap := map[string]database.Company{}
	companySelectEntryData := make([]string, 0, len(companies))
	for _, company := range companies {
		companySelectEntryData = append(companySelectEntryData, company.Name)
		companiesMap[company.Name] = company
	}

	fromCompany := widget.NewSelectEntry(companySelectEntryData)
	toCompany := widget.NewSelectEntry(companySelectEntryData)

	dateEntry := widget.NewDateEntry()
	currentDate := time.Now()
	dateEntry.SetDate(&currentDate)

	items, err := c.dbListItem()
	if err != nil {
		slog.Error("DB:", "msg", err)
	}
	itemsMap := map[string]database.Item{}
	itemsSelectEntryData := make([]string, 0, len(items))
	for _, item := range items {
		itemsSelectEntryData = append(itemsSelectEntryData, item.Name)
		itemsMap[item.Name] = item
	}

	itemEntry := widget.NewSelectEntry(itemsSelectEntryData)
	qtyEntry := widget.NewEntry()

	shipAddrs, err := c.dbListShippingAddress()
	if err != nil {
		slog.Error("DB:", "msg", err)
	}
	shipAddrsMap := map[string]database.ShippingAddress{}
	shipAddrSelectEntryData := make([]string, 0, len(shipAddrs))
	for _, shipAddr := range shipAddrs {
		shipAddrSelectEntryData = append(shipAddrSelectEntryData, shipAddr.Name)
		shipAddrsMap[shipAddr.Name] = shipAddr
	}
	shipTo := widget.NewSelectEntry(shipAddrSelectEntryData)

	paymentDetails, err := c.dbListPaymentDetail()
	if err != nil {
		slog.Error("DB:", "msg", err)
	}
	paymentDetailsMap := map[string]database.PaymentDetail{}
	paymentDetailSelectEntryData := make([]string, 0, len(paymentDetails))
	for _, paymentDetail := range paymentDetails {
		paymentDetailSelectEntryData = append(
			paymentDetailSelectEntryData,
			paymentDetail.AccHolder,
		)
		paymentDetailsMap[paymentDetail.AccHolder] = paymentDetail
	}
	paymentTo := widget.NewSelectEntry(paymentDetailSelectEntryData)

	form := &widget.Form{}
	form.Append("Invoice From", fromCompany)
	form.Append("Invoice To", toCompany)
	form.Append("Date", dateEntry)
	form.Append("Item", itemEntry)
	form.Append("Qty", qtyEntry)
	form.Append("Ship to", shipTo)
	form.Append("Payment to", paymentTo)

	form.OnSubmit = func() {
		slog.Info("User Input:", "invoiceFrom", fromCompany.Text)

		invNumber, err := strconv.Atoi(invoiceNumber.Text)
		if err != nil {
			slog.Error("Strconv:", "msg", err)
		}
		fromCompany := companiesMap[fromCompany.Text]
		toCompany := companiesMap[toCompany.Text]
		date := dateEntry.Date
		item := itemsMap[itemEntry.Text]
		qty, err := strconv.Atoi(qtyEntry.Text)
		if err != nil {
			slog.Error("Strconv:", "msg", err)
		}
		shipAddr := shipAddrsMap[shipTo.Text]
		pay := paymentDetailsMap[paymentTo.Text]

		if err := generator.ValidateInvoiceData(
			fromCompany, toCompany, date, item, qty, shipAddr, pay,
		); err != nil {
			slog.Error("Validation:", "msg", err)
		}

		slog.Info(
			"Invoice data from invoice.go", "from company", fromCompany,
			"to company", toCompany, "date", date, "item", item, "qty", qty,
			"shipAddr", shipAddr, "pay", pay,
		)

		invoiceData := generator.InvoiceDataMap(
			int64(invNumber),
			fromCompany,
			toCompany,
			date,
			item,
			qty,
			shipAddr,
			pay,
		)

		// TO-DO: Save the invoice data into database
		err = c.dbAddInvoice(
			int64(invNumber), fromCompany.ID, toCompany.ID, pay.ID,
			shipAddr.ID, float64(0), // TO-DO: calculate the Igst and total
			float64(0), *date, item.ID, int64(qty),
		)
		if err != nil {
			slog.Error("DB:", "msg", err)
			dialog.ShowError(err, w)
		}

		// Generate the invoice PDF
		err = generator.GenerateInvoice(invoiceData, c.TypstBinPath)
		if err != nil {
			slog.Error("Typst:", "msg", err)
			dialog.ShowError(err, w)
		}

		dialog.ShowInformation("Success", "Invoice generated successfully", w)

		for _, item := range form.Items {
			switch item.Widget.(type) {
			case *widget.Entry:
				item.Widget.(*widget.Entry).SetText("")
			case *widget.SelectEntry:
				item.Widget.(*widget.SelectEntry).SetText("")
			case *widget.DateEntry:
				item.Widget.(*widget.DateEntry).SetDate(&currentDate)
			}
			item.Widget.Refresh()
		}
	}

	return form
}
