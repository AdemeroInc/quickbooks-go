// Copyright (c) 2018, Randy Westlund. All rights reserved.
// This code is under the BSD-2-Clause license.

package quickbooks

import (
	"strconv"
	"strings"
)

// UserInfo describes an employee.
type Employee struct {
	Id               string
	SyncToken        string
	PrimaryAddr      PhysicalAddress
	PrimaryEmailAddr EmailAddress
	DisplayName      string
	Title            string
	BillableTime     bool
	GivenName        string
	BirthDate        Date
	MiddleName       string
	SSN              string
	PrimaryPhone     TelephoneNumber
	Active           bool
	ReleasedDate     Date
	MetaData         MetaData
	CostRate         float64
	Mobile           TelephoneNumber
	Gender           string
	HiredDate        Date
	BillRate         float64
	Organization     bool
	Suffix           string
	FamilyName       string
	PrintOnCheckName string
	EmployeeNumber   string
	V4IDPseudonym    string
}

// GetEmployees gets the employees
func (c *Client) GetEmployees(startpos int, pageSizeOverride int) ([]Employee, error) {

	pageSize := queryPageSize
	if pageSizeOverride > 0 {
		pageSize = pageSizeOverride
	}

	var r struct {
		QueryResponse struct {
			Employee      []Employee
			StartPosition int
			MaxResults    int
		}
	}
	q := "SELECT * FROM Employee ORDERBY FamilyName,GivenName STARTPOSITION " +
		strconv.Itoa(startpos) + " MAXRESULTS " + strconv.Itoa(pageSize)
	err := c.query(q, &r)
	if err != nil {
		return nil, err
	}

	if r.QueryResponse.Employee == nil {
		r.QueryResponse.Employee = make([]Employee, 0)
	}
	return r.QueryResponse.Employee, nil
}

// QueryEmployeeById gets an employee with a given Id.
func (c *Client) QueryEmployeeById(id string) (*Employee, error) {

	var r struct {
		QueryResponse struct {
			Employee []Employee
		}
	}
	err := c.query("SELECT * FROM Employee WHERE Id = '"+
		strings.Replace(id, "'", "''", -1)+"'", &r)
	if err != nil {
		return nil, err
	}

	if len(r.QueryResponse.Employee) > 0 {
		return &r.QueryResponse.Employee[0], nil
	}

	return nil, nil
}
