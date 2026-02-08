package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (c *config) addNewCompany(a fyne.App, w fyne.Window) *widget.Form {
	company := widget.NewEntry()
	gst := widget.NewEntry()
	address := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Company Name", Widget: company},
			{Text: "GST", Widget: gst},
			{Text: "Address", Widget: address},
		},
	}

	form.OnSubmit = func() {
		log.Println("Form submitted: ", company.Text)
		log.Println("multiline: ", address.Text)
		log.Println("GST: ", gst.Text)
		a.SendNotification(fyne.NewNotification("Form submitted ", "Added "+company.Text))
		company.Text = ""
		gst.Text = ""
		address.Text = ""
		company.Refresh()
		gst.Refresh()
		address.Refresh()
	}

	return form
}

func (c *config) addNewPaymentMethod(a fyne.App, w fyne.Window) *widget.Form {
	acc_holder := widget.NewEntry()
	acc_number := widget.NewEntry()
	ifsc := widget.NewEntry()
	branch := widget.NewEntry()
	bank_name := widget.NewEntry()
	virtual_payment_addr := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{},
	}

	form.Append("Account Holder Name", acc_holder)
	form.Append("Account Number", acc_number)
	form.Append("IFSC Code", ifsc)
	form.Append("Branch Name", branch)
	form.Append("Bank Name", bank_name)
	form.Append("Virtual Payment Address", virtual_payment_addr)

	form.OnSubmit = func() {
		log.Println("Form Submitted")
		log.Println("Acc Holder: ", acc_holder)
		log.Println(acc_number, ifsc, branch, bank_name, virtual_payment_addr)

		acc_holder.Text = ""
		acc_number.Text = ""
		ifsc.Text = ""
		branch.Text = ""
		bank_name.Text = ""
		virtual_payment_addr.Text = ""
		acc_holder.Refresh()
		acc_number.Refresh()
		ifsc.Refresh()
		branch.Refresh()
		bank_name.Refresh()
		virtual_payment_addr.Refresh()
	}

	return form
}
