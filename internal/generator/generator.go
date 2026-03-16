package generator

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	typst "github.com/Dadido3/go-typst"
)

const typstMainTemplate = `
#import "./invoice_template.typ": *

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
`

func GenerateInvoice(input map[string]any, typstBinPath string) error {
	now := time.Now()
	invoiceFileName := fmt.Sprintf(
		"invoice-%d-%02d-%02d-%02d:%02d.pdf",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Minute(),
		now.Hour(),
	)
	pwd, err := filepath.Abs(".")
	if err != nil {
		slog.Error("Filepath:", "msg", err)
		return err
	}

	invoiceFilePath := filepath.Join(pwd, "invoices", invoiceFileName)

	typstMainFilePath, err := filepath.Abs("typst/main.typ")
	if err != nil {
		slog.Error("Filepath:", "msg", err)
		return err
	}

	var markup bytes.Buffer
	if err := typst.InjectValues(&markup, input); err != nil {
		slog.Error("Typst Dep:", "msg", err)
		return err
	}

	markup.WriteString(typstMainTemplate)

	// invoice file
	invoiceFile, err := os.Create(invoiceFilePath)
	if err != nil {
		slog.Error("File:", "msg", err)
		return err
	}
	defer invoiceFile.Close()

	typstCaller := typst.CLI{
		ExecutablePath:   typstBinPath,
		WorkingDirectory: filepath.Dir(typstMainFilePath),
	}

	slog.Info("Typst: ", "markup", markup.String())

	if err = typstCaller.Compile(&markup, invoiceFile, nil); err != nil {
		slog.Error("Typst:", "msg", err)
		return err
	}
	slog.Info("Typst: Success", "msg", "Created invoice successfully")
	return nil
}
