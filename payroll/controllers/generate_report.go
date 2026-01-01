package controllers

import (
	"fmt"
	"os"

	"payroll/models"
	"payroll/utils"

	"github.com/jung-kurt/gofpdf"
)

func GenerateReport(date string, employeePayroll *[]models.Payroll) {
	// extract year and month
	y, m, err := utils.ParseMonthYear(date)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	// making the date format consistent for saving data
	date = fmt.Sprintf("%d/%d", m, y)

	// -----------------------------------
	pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    
	// Set font for title
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Employee Report")
	pdf.Ln(15)

	// Set font for content
	pdf.SetFont("Arial", "", 12)

	// display the date
	pdf.Cell(0, 10, date)
	pdf.Ln(15)

	for _, e := range *employeePayroll {
		if err := generateEmployeeReport(e, pdf); err != nil {
			fmt.Fprintln(os.Stderr, "Error: generating report", err)
		}
	}

	fileName := fmt.Sprintf("%d-%d", m, y)
	path := fmt.Sprintf("reports/%s.pdf", fileName)
	if err := pdf.OutputFileAndClose(path); err != nil {
		fmt.Fprintln(os.Stderr, "Error: creating pdf file for report error", err)
		return
	}

	fmt.Println("Employee report successfully created at", path)
}

func generateEmployeeReport(e models.Payroll, pdf *gofpdf.Fpdf) error {
	// Add employee details
	pdf.Cell(40, 10, "ID:")
	pdf.Cell(0, 10, e.Employee.EmployeeId)
	pdf.Ln(6)

	pdf.Cell(40, 10, "Name:")
	pdf.Cell(0, 10, e.Employee.Name)
	pdf.Ln(6)

	pdf.Cell(40, 10, "Position:")
	pdf.Cell(0, 10, e.Employee.Position)
	pdf.Ln(6)

	pdf.Cell(40, 10, "Hours:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f", e.TotalHoursWorked))
	pdf.Ln(6)

	pdf.Cell(40, 10, "OT:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f", e.TotalOverTime))
	pdf.Ln(6)

	pdf.Cell(40, 10, "Gross Income:")
	pdf.Cell(0, 10, fmt.Sprintf("$ %.2f", e.GrossIncome))
	pdf.Ln(30)

	return nil
}