package generator

import (
	"bytes"
	"fmt"
	"io"
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

func GenerateInvoice(
	input map[string]any,
	typstBinPath string,
	typstTemplateEmbedData []byte,
) error {
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

	typstTempDir, err := os.MkdirTemp("", "typst")
	if err != nil {
		slog.Error("Generator:", "msg", "error creating temp folder"+err.Error())
		return err
	}

	input["image-path"], err = resolveImagePath(typstTempDir, input["image-path"].(string))
	if err != nil {
		slog.Error("Generator:", "msg", err)
	}

	var markup bytes.Buffer
	if err := typst.InjectValues(&markup, input); err != nil {
		slog.Error("Typst Dep:", "msg", err)
		return err
	}

	markup.WriteString(typstMainTemplate)

	// Copy the embed typst template in the temp directory
	tempTemplateFile, err := os.Create(filepath.Join(typstTempDir, "invoice_template.typ"))
	if err != nil {
		slog.Error("Generator:", "msg", "error creating template file in temp dir: "+err.Error())
	}

	_, err = tempTemplateFile.Write(typstTemplateEmbedData)
	if err != nil {
		slog.Error("Generator:", "msg", "error writing template data: "+err.Error())
	}

	// invoice file
	invoiceFile, err := os.Create(invoiceFilePath)
	if err != nil {
		slog.Error("File:", "msg", err)
		return err
	}
	defer invoiceFile.Close()

	typstCaller := typst.CLI{
		ExecutablePath:   typstBinPath,
		WorkingDirectory: typstTempDir,
	}

	slog.Info("Typst: ", "markup", markup.String())

	if err = typstCaller.Compile(&markup, invoiceFile, nil); err != nil {
		slog.Error("Typst:", "msg", err)
		return err
	}
	slog.Info("Typst: Success", "msg", "Created invoice successfully")
	return nil
}

func resolveImagePath(typstDir, absImagePath string) (string, error) {
	// Copy image into typst working dir and return relative path
	filename := filepath.Base(absImagePath)
	destPath := filepath.Join(typstDir, filename)

	src, err := os.Open(absImagePath)
	if err != nil {
		return "", fmt.Errorf("error opening input image file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(destPath)
	if err != nil {
		return "", fmt.Errorf("error creating output image file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("error copying image data: %w", err)
	}

	return filename, nil // relative path — typst can find it
}
