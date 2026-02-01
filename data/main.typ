#import "./invoice_template.typ": *

#let company = ""
#let company-gstin = ""
#let company-address = ""
#let invoice-date = datetime.today()
#let invoice-number = 1
#let bill-to-name = ""
#let bill-to-gstin = ""
#let bill-to-address = ""

#let items = (
  (
    product-name: "",
    hsn-sac: "",
    qty: "",
    price: "",
    gst: "",
    total: "",
  ),
)
#let ship-to-address = ""
#let payment-data = (
  acc-name:  "",
  acc-number: "",
  ifsc: "",
  branch: "",
  bank-name: "",
  virtual-address: "",
)
#let sub-total = decimal("0.00")
#let igst = decimal("0.00")
#let image-path = "fake-sign.jpg"

#show: invoice.with(
  company: company,
  company-gstin: company-gstin,
  company-address: company-address,
  invoice-date: invoice-date,
  invoice-number: invoice-number,
  bill-to-name: bill-to-name,
  bill-to-gstin: bill-to-gstin,
  bill-to-address: bill-to-address,
  items: items,
  ship-to-address: ship-to-address,
  payment-data: payment-data,
  sub-total: sub-total,
  igst: igst,
  image-path: image-path,
)
