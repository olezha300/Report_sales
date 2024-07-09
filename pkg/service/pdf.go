package service

import (
	"bytes"
	"fmt"
	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/jung-kurt/gofpdf"
	"strconv"
)

var widths = []float64{60, 60, 20, 20, 30}

func PdfMaker(info [][]rgm.SalesPdf, biggestSale []rgm.SalesPdf, lowerSale []rgm.SalesPdf) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	createHeader(pdf, "Saller", "Name Product", "Price", "Count", "Total")
	createTable(pdf, info)
	pdf.Ln(-1)
	pdf.Cell(20, 10, "Biggest Sale:")
	pdf.Ln(-1)
	createHeader(pdf, "id", "Name Product", "Price", "Count", "Total")
	createParamTable(pdf, biggestSale)
	pdf.Ln(-1)
	pdf.Cell(20, 10, "Lowest Sale:")
	pdf.Ln(-1)
	createHeader(pdf, "id", "Name Product", "Price", "Count", "Total")
	createParamTable(pdf, lowerSale)
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func createTable(pdf *gofpdf.Fpdf, info [][]rgm.SalesPdf) {
	var sum float64
	data := [][]string{}
	for _, sales := range info {
		for _, product := range sales {
			sum += (product.Price * float64(product.Amount))
			res := []string{product.Saller, product.Title, strconv.FormatFloat(product.Price, 'f', 2, 64),
				strconv.Itoa(product.Amount), strconv.FormatFloat((product.Price * float64(product.Amount)), 'f', 2, 64)}
			data = append(data, res)
		}

	}

	for _, row := range data {
		for i, str := range row {
			pdf.CellFormat(widths[i], 10, str, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}
	str := fmt.Sprintf("Total amount sales: %.2f", sum)
	pdf.Cell(20, 10, str)
	pdf.Ln(-1)
}

func createParamTable(pdf *gofpdf.Fpdf, info []rgm.SalesPdf) {
	var sum float64
	data := [][]string{}
	for _, sales := range info {
		sum += (sales.Price * float64(sales.Amount))
		res := []string{sales.Saller, sales.Title, strconv.FormatFloat(sales.Price, 'f', 2, 64),
			strconv.Itoa(sales.Amount), strconv.FormatFloat((sales.Price * float64(sales.Amount)), 'f', 2, 64)}
		data = append(data, res)

	}

	for _, row := range data {
		for i, str := range row {
			pdf.CellFormat(widths[i], 10, str, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}

	str := fmt.Sprintf("Total amount sales: %.2f", sum)
	pdf.Cell(20, 10, str)
	pdf.Ln(-1)
}

func createHeader(pdf *gofpdf.Fpdf, header ...string) {
	for i, str := range header {
		pdf.CellFormat(widths[i], 10, str, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)
}
