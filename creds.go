package sawari

import (
	"os"
)

// NewCreds gives credentials
func NewCreds() Creds {

	employeeName := os.Getenv("NAME")
	employeeDesignation := os.Getenv("DESIGNATION")
	monthlyCeiling := os.Getenv("CEILING")
	employeeCode := os.Getenv("EMPLOYEE_CODE")

	c := Creds{}
	if len(employeeName) == 0 {
		c.Name = TakeInput("Name?")
	} else {
		c.Name = employeeName
	}

	if len(employeeDesignation) == 0 {
		c.Designation = TakeInput("Designation?")
	} else {
		c.Designation = employeeDesignation
	}

	if len(employeeCode) == 0 {
		c.EmployeeCode = TakeInput("Employee Code?")
	} else {
		c.EmployeeCode = employeeCode
	}
	if len(monthlyCeiling) == 0 {
		c.MonthlyCeiling = TakeInput("Monthly Ceiling?")
	} else {
		c.MonthlyCeiling = monthlyCeiling
	}

	return c
}
