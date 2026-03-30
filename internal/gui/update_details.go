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

	"github.com/Vikuuu/invoice_generator/internal/database"
	"github.com/Vikuuu/invoice_generator/internal/utils"
)

func (c *Config) updateCompanyForm(
	a fyne.App,
	w fyne.Window,
	companyData database.Company,
) *widget.Form {
	company := widget.NewEntry()
	company.SetText(companyData.Name)
	gst := widget.NewEntry()
	gst.SetText(companyData.Gst)
	address := widget.NewMultiLineEntry()
	address.SetText(companyData.Address)
	signatureEntry := widget.NewEntry()
	signatureEntry.SetText(companyData.SignatureImagePath)
	signatureEntry.Disable()
	sigPath := ""

	browseBtn := widget.NewButton("Browse", func() {
		sigImage := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if reader == nil {
				slog.Info("GUI: update company", "id", companyData.ID, "msg", "got nil in reader")
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
		slog.Info(
			"GUI: update company",
			"id",
			companyData.ID,
			"msg",
			"Form Submission button clicked",
		)
		name := strings.TrimSpace(company.Text)
		g := strings.ToUpper(strings.TrimSpace(gst.Text))
		addr := strings.TrimSpace(address.Text)
		sigPath = strings.TrimSpace(sigPath)

		var err error
		if err = validateCompanyDetail(name, g, addr); err != nil {
			slog.Error("GUI: update company", "id", companyData.ID, "msg", err)
			dialog.ShowError(err, w)
			return
		}
		var updated database.Company
		if updated, err = c.dbUpdateCompany(name, g, addr, sigPath, companyData); err != nil {
			slog.Error("DB:", "id", companyData.ID, "msg", err)
			dialog.ShowError(err, w)
			return
		}
		slog.Info("Company information updated successfully")
		dialog.ShowInformation("Success", "Company information updated successfully", w)

		// Update the form widgets with the new updated data
		company.SetText(updated.Name)
		gst.SetText(updated.Gst)
		address.SetText(updated.Address)
		signatureEntry.SetPlaceHolder(updated.SignatureImagePath)
	}

	return form
}

func (c *Config) updatePaymentDetailForm(
	a fyne.App,
	w fyne.Window,
	pd database.PaymentDetail,
) *widget.Form {
	accHolder := widget.NewEntry()
	accHolder.SetText(pd.AccHolder)
	accNumber := widget.NewEntry()
	accNumber.SetText(string(pd.AccNumber))
	ifsc := widget.NewEntry()
	ifsc.SetText(pd.Ifsc)
	branch := widget.NewEntry()
	branch.SetText(pd.Branch)
	bankName := widget.NewEntry()
	bankName.SetText(pd.BankName)
	virtualPaymentAddr := widget.NewEntry()
	virtualPaymentAddr.SetText(pd.VirtualPaymentAddr)
	// TODO: update this to get the options to select from
	companyId := widget.NewEntry()
	companyId.SetText(string(pd.FkCompanyID))

	form := &widget.Form{
		Items: []*widget.FormItem{},
	}

	form.Append("Account Holder Name", accHolder)
	form.Append("Account Number", accNumber)
	form.Append("IFSC Code", ifsc)
	form.Append("Branch Name", branch)
	form.Append("Bank Name", bankName)
	form.Append("Virtual Payment Address", virtualPaymentAddr)
	form.Append("Company", companyId)

	form.OnSubmit = func() {
		slog.Info(
			"GUI: update payment detail",
			"id",
			pd.ID,
			"msg",
			"Update form submission button clicked",
		)
		accNum, err := strconv.Atoi(accNumber.Text)
		if err != nil {
			slog.Error("strconv:", "msg", err)
		}
		comId, err := strconv.Atoi(companyId.Text)
		if err != nil {
			slog.Error("strconv:", "msg", err)
		}
		var updated database.PaymentDetail
		if updated, err = c.dbUpdatePaymentMethod(accHolder.Text, int64(accNum), ifsc.Text, branch.Text, bankName.Text, virtualPaymentAddr.Text, int64(comId), pd); err != nil {
			slog.Error("DB:", "msg", err)
			dialog.ShowError(err, w)
			return
		}
		slog.Info("Payment detail updated successfully")
		dialog.ShowInformation("Success", "Payment detail updated successfully", w)

		// Update the form widgets with the new updated data
		accHolder.SetText(updated.AccHolder)
		accNumber.SetText(string(updated.AccNumber))
		ifsc.SetText(updated.Ifsc)
		branch.SetText(updated.Branch)
		bankName.SetText(updated.BankName)
		virtualPaymentAddr.SetText(updated.VirtualPaymentAddr)
		// TODO: update this to get the name of the company
		companyId.SetText(string(updated.FkCompanyID))
	}
	return form
}

func (c *Config) updateItemForm(
	a fyne.App, w fyne.Window, item database.Item,
) *widget.Form {
	itemName := widget.NewEntry()
	itemName.SetText(item.Name)
	// update
	hsn := newNumericalEntry(0)
	hsn.SetNumber(item.Hsn)
	gst := newNumericalEntry(2)
	gst.SetNumber(item.Gst)
	price := widget.NewEntry()
	price.SetText(strconv.FormatFloat(item.Price, 'g', -1, 64))

	form := &widget.Form{
		Items: []*widget.FormItem{},
	}

	form.Append("Item Name", itemName)
	form.Append("HSN/SAC", hsn)
	form.Append("Price", price)

	form.OnSubmit = func() {
		slog.Info("GUI: Update item", "id", item.ID, "msg", "Update item button clicked")
		h, err := strconv.Atoi(hsn.Text)
		if err != nil {
			slog.Error("Strconv:", "msg", err)
		}
		p, err := strconv.ParseFloat(price.Text, 64)
		if err != nil {
			slog.Error("Strconv:", "msg", err)
		}
		g, err := strconv.Atoi(gst.Text)
		if err != nil {
			slog.Error("Strconv:", "msg", err)
		}
		var updated database.Item
		if updated, err = c.dbUpdateItem(itemName.Text, int64(h), int64(g), p, item); err != nil {
			slog.Error("DB:", "msg", err)
			dialog.ShowError(err, w)
			return
		}
		slog.Info("Item detail updated successfully")
		dialog.ShowInformation("Success", "Item detail updated successfully", w)

		// Update the form widgets with the new updated data
		itemName.SetText(updated.Name)
		hsn.SetText(string(updated.Hsn))
		gst.SetText(string(updated.Gst))
		price.SetText(strconv.FormatFloat(updated.Price, 'g', -1, 64))
	}
	return form
}

func (c *Config) updateShippingAddressForm(
	a fyne.App, w fyne.Window, sa database.ShippingAddress,
) *widget.Form {
	name := widget.NewEntry()
	name.SetText(sa.Name)
	address := widget.NewMultiLineEntry()
	address.SetText(sa.Address)

	form := &widget.Form{
		Items: []*widget.FormItem{},
	}

	form.Append("Name", name)
	form.Append("Address", address)

	form.OnSubmit = func() {
		slog.Info(
			"GUI: Update shipping address",
			"id",
			sa.ID,
			"msg",
			"Update shipping address button clicked",
		)
		var updated database.ShippingAddress
		var err error
		if updated, err = c.dbUpdateShippingAddress(name.Text, address.Text, sa); err != nil {
			slog.Error("DB:", "msg", err)
			dialog.ShowError(err, w)
			return
		}
		slog.Info("Shipping address detail updated successfully")
		dialog.ShowInformation("Success", "Item detail updated successfully", w)

		// Update the form widgets with the new updated data
		name.SetText(updated.Name)
		address.SetText(updated.Address)
	}
	return form
}
