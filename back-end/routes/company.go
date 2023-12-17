package routes

import (
	"database/sql"
	"e-identitet/tick-api/utils"
	"log"
)

func GetCompanyNames(db *sql.DB) []string {
	rows, companyTableErr := db.Query("SELECT company_name FROM company;")
	checkError := utils.CheckError
	checkError(companyTableErr)
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
	companyName    string
	companyId      string
	companyAddress string
	companyPhone   string
	companyMail    string
}
