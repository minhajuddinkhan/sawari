package sawari

// Creds Creds
type Creds struct {
	Name           string
	EmployeeCode   string
	Designation    string
	MonthlyCeiling string
}

// Entry Entry
type Entry struct {
	Serial      int
	Date        string
	Description string
	Amount      string
}
