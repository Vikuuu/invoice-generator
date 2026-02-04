package generator

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

func GenerateInvoice() {
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
		panic(err)
	}

	invoiceFilePath := filepath.Join(pwd, "invoices", invoiceFileName)

	paths, err := filepath.Glob("typst/*.typ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+q", paths)

	path, err := filepath.Abs("typst/main.typ")
	if err != nil {
		panic(err)
	}
	typstCmd := exec.Command("typst", "compile", path, invoiceFilePath)

	out, err := typstCmd.Output()
	if err != nil {
		var execErr *exec.Error
		var exitErr *exec.ExitError
		switch {
		case errors.As(err, &execErr):
			fmt.Println("Failed executing: ", err)
		case errors.As(err, &exitErr):
			exitCode := exitErr.ExitCode()
			fmt.Println("command rc = ", exitCode)
			fmt.Println(string(exitErr.Stderr))
		default:
			panic(err)
		}
	}
	fmt.Println(string(out))
}
