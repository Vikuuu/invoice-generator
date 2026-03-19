package gui

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/Vikuuu/invoice_generator/internal/utils"
)

func (c *Config) addNewCompany(a fyne.App, w fyne.Window) *widget.Form {
	company := widget.NewEntry()
	gst := widget.NewEntry()
	address := widget.NewMultiLineEntry()

	signatureEntry := widget.NewEntry()
	signatureEntry.SetPlaceHolder("No file selected...")
	signatureEntry.Disable()
	var sigPath string

	browseBtn := widget.NewButton("Browse", func() {
		sigImage := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if reader == nil {
				slog.Info("GUI: add new company", "msg", "got nil in reader")
				return
			}
			defer reader.Close()

			signatureEntry.SetText(reader.URI().Name())

			assetPath, _ := utils.GetAssetsAppPath()
			sigPath = filepath.Join(assetPath, reader.URI().Name())
			f, err := os.Create(sigPath)
			if err != nil {
				slog.Error("GUI:", "msg", err)
				return
			}
			defer f.Close()

			_, err = io.Copy(f, reader)
			if err != nil {
				slog.Error("GUI:", "msg", err)
				return
			}
		}, w)
		sigImage.Show()
	})

	sigFileRow := container.NewBorder(nil, nil, nil, browseBtn, signatureEntry)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Company Name", Widget: company},
			{Text: "GST", Widget: gst},
			{Text: "Address", Widget: address},
			{Text: "Signature Image", Widget: sigFileRow},
		},
	}

	form.OnSubmit = func() {
		slog.Info("Form Submission button clicked")
		name := strings.TrimSpace(company.Text)
		g := strings.ToUpper(strings.TrimSpace(gst.Text))
		addr := strings.TrimSpace(address.Text)
		sigPath = strings.TrimSpace(sigPath)

		var err error
		if err = validateCompanyDetail(name, g, addr); err != nil {
			slog.Error("Validation", "error", err)
			dialog.ShowError(err, w)
			return
		}
		if err = c.dbAddCompany(name, g, addr, sigPath); err != nil {
			slog.Error("DB Error", "error", err)
			dialog.ShowError(err, w)
			return
		}

		slog.Info("Company information added successfully")
		dialog.ShowInformation("Success", "Company information added successfully", w)

		for _, item := range form.Items {
			if entry, ok := item.Widget.(*widget.Entry); ok {
				entry.SetText("")
			}
			item.Widget.Refresh()
		}
	}

	return form
}

func (c *Config) addNewPaymentMethod(a fyne.App, w fyne.Window) *widget.Form {
	accHolder := widget.NewEntry()
	accNumber := widget.NewEntry()
	ifsc := widget.NewEntry()
	branch := widget.NewEntry()
	bankName := widget.NewEntry()
	virtualPaymentAddr := widget.NewEntry()
	companyId := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{},
	}

	form.Append("Account Holder Name", accHolder)
	form.Append("Account Number", accNumber)
	form.Append("IFSC Code", ifsc)
	form.Append("Branch Name", branch)
	form.Append("Bank Name", bankName)
	form.Append("Virtual Payment Address", virtualPaymentAddr)

	form.OnSubmit = func() {
		slog.Info("Form Submitted")

		accNum, _ := strconv.Atoi(accNumber.Text)
		comId, _ := strconv.Atoi(companyId.Text)
		if err := c.dbAddPaymentMethod(
			accHolder.Text,
			int64(accNum),
			ifsc.Text,
			branch.Text,
			bankName.Text,
			virtualPaymentAddr.Text,
			int64(comId),
		); err != nil {
			slog.Error("DB Error", "error", err)
			dialog.ShowError(err, w)
			return
		}

		slog.Info("Payment information added success")
		dialog.ShowInformation("Success", "Payment information add success", w)

		for _, item := range form.Items {
			if entry, ok := item.Widget.(*widget.Entry); ok {
				entry.SetText("")
			}
			item.Widget.Refresh()
		}
	}

	return form
}

func (c *Config) addNewItems(a fyne.App, w fyne.Window) *widget.Form {
	itemName := widget.NewEntry()
	hsn := widget.NewEntry()
	price := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{},
	}

	form.Append("Item Name", itemName)
	form.Append("HSN/SAC", hsn)
	form.Append("Price", price)

	form.OnSubmit = func() {
		slog.Info("Form Submitted")
		h, err := strconv.Atoi(hsn.Text)
		if err != nil {
			slog.Error("Strconv:", "msg", err)
		}
		p, err := strconv.ParseFloat(price.Text, 64)
		if err != nil {
			slog.Error("Strconv:", "msg", err)
		}
		if err := c.dbAddItemMethod(itemName.Text, int64(h), p); err != nil {
			slog.Error("DB Error", "error", err)
			dialog.ShowError(err, w)
			return
		}

		slog.Info("Item information added success")
		dialog.ShowInformation("Success", "Item information added success", w)

		for _, item := range form.Items {
			if entry, ok := item.Widget.(*widget.Entry); ok {
				entry.SetText("")
			}
			item.Widget.Refresh()
		}
	}

	return form
}

func (c *Config) addShippingAddress(a fyne.App, w fyne.Window) *widget.Form {
	name := widget.NewEntry()
	address := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{},
	}

	form.Append("Name", name)
	form.Append("Address", address)

	form.OnSubmit = func() {
		slog.Info("Form Submitted")

		if err := c.dbAddShippingAddress(name.Text, address.Text); err != nil {
			slog.Error("DB Error", "error", err)
			dialog.ShowError(err, w)
			return
		}

		slog.Info("Shipping address information added successfully")
		dialog.ShowInformation("Success", "Company information added successfully", w)

		for _, item := range form.Items {
			if entry, ok := item.Widget.(*widget.Entry); ok {
				entry.SetText("")
			}
			item.Widget.Refresh()
		}
	}

	return form
}
