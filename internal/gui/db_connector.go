package gui

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/Vikuuu/invoice_generator/internal/database"
	db "github.com/Vikuuu/invoice_generator/internal/database"
)

func (c *Config) dbAddCompany(name, gst, address, path string) error {
	arg := db.CreateCompanyParams{
		Name:               name,
		Gst:                gst,
		Address:            address,
		SignatureImagePath: path,
	}

	_, err := c.Queries.CreateCompany(c.Context, arg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) dbUpdateCompany(
	name, gst, address, path string,
	company database.Company,
) (database.Company, error) {
	arg := database.UpdateCompanyParams{
		ID:                 company.ID,
		Name:               company.Name,
		Gst:                company.Gst,
		Address:            company.Address,
		SignatureImagePath: company.SignatureImagePath,
	}

	if name != company.Name {
		arg.Name = name
	}
	if gst != company.Gst {
		arg.Gst = gst
	}
	if address != company.Address {
		arg.Address = address
	}
	if path != company.SignatureImagePath && strings.TrimSpace(path) != "" {
		arg.SignatureImagePath = path
	}

	updated, err := c.Queries.UpdateCompany(c.Context, arg)
	if err != nil {
		return updated, err
	}
	return updated, nil
}

func (c *Config) dbAddPaymentMethod(
	holder string, number int64, ifsc, branch, name,
	virtualPaymentAddress string, companyId int64,
) error {
	arg := db.CreatePaymentDetailParams{
		AccHolder:          holder,
		AccNumber:          number,
		Ifsc:               ifsc,
		Branch:             branch,
		BankName:           name,
		VirtualPaymentAddr: virtualPaymentAddress,
		FkCompanyID:        companyId,
	}

	_, err := c.Queries.CreatePaymentDetail(c.Context, arg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) dbUpdatePaymentMethod(
	accHolderName string, accNum int64, ifsc, branchName, bankName,
	virtualPaymentAddr string, companyId int64, pd database.PaymentDetail,
) (database.PaymentDetail, error) {
	arg := database.UpdatePaymentDetailParams{
		AccHolder:          pd.AccHolder,
		AccNumber:          pd.AccNumber,
		Ifsc:               pd.Ifsc,
		Branch:             pd.Branch,
		BankName:           pd.BankName,
		VirtualPaymentAddr: pd.VirtualPaymentAddr,
		FkCompanyID:        pd.FkCompanyID,
	}
	if accHolderName != pd.AccHolder && strings.TrimSpace(accHolderName) != "" {
		arg.AccHolder = accHolderName
	}
	if accNum != pd.AccNumber {
		arg.AccNumber = accNum
	}
	if ifsc != pd.Ifsc && strings.TrimSpace(ifsc) != "" {
		arg.Ifsc = ifsc
	}
	if branchName != pd.Branch && strings.TrimSpace(branchName) != "" {
		arg.Branch = branchName
	}
	if bankName != pd.BankName && strings.TrimSpace(bankName) != "" {
		arg.BankName = bankName
	}
	if virtualPaymentAddr != pd.VirtualPaymentAddr && strings.TrimSpace(virtualPaymentAddr) != "" {
		arg.VirtualPaymentAddr = virtualPaymentAddr
	}
	if companyId != pd.FkCompanyID {
		arg.FkCompanyID = companyId
	}

	updated, err := c.Queries.UpdatePaymentDetail(c.Context, arg)
	if err != nil {
		return updated, err
	}
	return updated, nil
}

func (c *Config) dbAddItemMethod(name string, hsn int64, price float64) error {
	arg := db.CreateItemParams{
		Name:  name,
		Hsn:   hsn,
		Price: price,
	}

	_, err := c.Queries.CreateItem(c.Context, arg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) dbUpdateItem(
	name string,
	hsn, gst int64,
	price float64,
	item database.Item,
) (database.Item, error) {
	arg := database.UpdateItemParams{
		Name:  item.Name,
		Hsn:   item.Hsn,
		Price: item.Price,
		ID:    item.ID,
		Gst:   item.Gst,
	}
	if name != item.Name && strings.TrimSpace(name) != "" {
		arg.Name = name
	}
	if hsn != item.Hsn {
		arg.Hsn = hsn
	}
	if gst != item.Gst {
		arg.Gst = gst
	}
	if price != item.Price {
		arg.Price = price
	}
	updated, err := c.Queries.UpdateItem(c.Context, arg)
	if err != nil {
		return updated, err
	}
	return updated, nil
}

func (c *Config) dbAddShippingAddress(name, address string) error {
	arg := db.CreateShippingAddressParams{
		Name:    name,
		Address: address,
	}

	_, err := c.Queries.CreateShippingAddress(c.Context, arg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) dbUpdateShippingAddress(
	name, address string,
	sa database.ShippingAddress,
) (database.ShippingAddress, error) {
	arg := database.UpdateShippingAddressParams{
		Name:    sa.Name,
		Address: sa.Address,
		ID:      sa.ID,
	}
	if name != sa.Name && strings.TrimSpace(name) != "" {
		arg.Name = name
	}
	if address != sa.Address && strings.TrimSpace(address) != "" {
		arg.Address = address
	}

	updated, err := c.Queries.UpdateShippingAddress(c.Context, arg)
	if err != nil {
		return updated, err
	}
	return updated, nil
}

func (c *Config) dbListCompany() ([]db.Company, error) {
	companies, err := c.Queries.ListCompany(c.Context)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (c *Config) dbListPaymentDetail() ([]db.PaymentDetail, error) {
	paymentDetails, err := c.Queries.ListPaymentDetail(c.Context)
	if err != nil {
		return nil, err
	}
	return paymentDetails, nil
}

func (c *Config) dbListItem() ([]db.Item, error) {
	items, err := c.Queries.ListItem(c.Context)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Config) dbListShippingAddress() ([]db.ShippingAddress, error) {
	address, err := c.Queries.ListShippingAddress(c.Context)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (c *Config) dbAddInvoice(
	invoiceNumber int64,
	fkFromCompany, fkToCompany, fkPaymentDetail, fkShippingAddress int64,
	igst, total float64,
	date time.Time,
	fkItem int64,
	qty int64,
) error {
	invoiceArg := db.CreateInvoiceParams{
		InvoiceNumber:     invoiceNumber,
		Date:              date,
		FkFromCompany:     fkFromCompany,
		FkToCompany:       fkToCompany,
		FkPaymentDetail:   fkPaymentDetail,
		Igst:              igst,
		Total:             total,
		FkShippingAddress: fkShippingAddress,
	}
	invoice, err := c.Queries.CreateInvoice(c.Context, invoiceArg)
	if err != nil {
		return err
	}

	itemArg := db.CreateInvoiceItemParams{
		FkInvoice: invoice.ID,
		FkItem:    fkItem,
		Qty:       qty,
	}
	_, err = c.Queries.CreateInvoiceItem(c.Context, itemArg)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) dbGetLatestInvoiceNumber() (int64, error) {
	num, err := c.Queries.GetLatestInvoiceNumber(c.Context)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return int64(1), nil
		}
		return int64(1), err
	}

	return num + 1, nil
}

func (c *Config) dbListInvoices() ([]db.Invoice, error) {
	invoice, err := c.Queries.ListInvoice(c.Context)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}
