package routes

import (
	"database/sql"
	"e-identitet/tick-api/utils"
	"log"
)

func GetCompanies(db *sql.DB) []Company {
	rows, err := db.Query("SELECT * FROM company;")
	checkError := utils.CheckError
	checkError(err)
	companies := make([]Company, 0)
	for rows.Next() {
		var company Company
		if err := rows.Scan(&company.CompanyId, &company.CompanyName, &company.CompanyAddress, &company.CompanyMail, &company.CompanyPhone); err != nil {
			log.Fatal(err)
		}

		companies = append(companies, company)
	}
	return companies
}

func GetCompanyNames(db *sql.DB) []string {
	rows, err := db.Query("SELECT company_name FROM company;")
	checkError := utils.CheckError
	checkError(err)
	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}
	return names
}

func GetCompany(db *sql.DB) {
}

type Company struct {
	CompanyName    string `json:"companyName"`
	CompanyId      string `json:"companyId"`
	CompanyAddress string `json:"companyAddress"`
	CompanyPhone   string `json:"companyPhone"`
	CompanyMail    string `json:"companyMail"`
}
