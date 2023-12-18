package routes

import (
	"database/sql"
	"e-identitet/tick-api/utils"
	"fmt"
	"log"
)

func GetCompanies(db *sql.DB) []CompanyWithTick {
	rows, err := db.Query("SELECT company.company_id, company.company_name, company.company_address, company.company_phone, company.company_mail, tick.tick_id, tick.core_service, tick.date, tick.type_tick FROM company INNER JOIN tick ON company.company_id = tick.company_id")
	checkError := utils.CheckError
	checkError(err)
	companies := make([]CompanyWithTick, 0)
	for rows.Next() {
		var company CompanyWithTick
		if err := rows.Scan(&company.CompanyId, &company.CompanyName, &company.CompanyAddress, &company.CompanyPhone, &company.CompanyMail, &company.TickId, &company.CoreService, &company.Date, &company.TypeTick); err != nil {
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

func GetCompany(db *sql.DB, companyName string) []CompanyWithTick {
	singleCompanyQuery := "SELECT company.company_id, company.company_name, company.company_address, company.company_phone, company.company_mail, tick.tick_id, tick.core_service, tick.date, tick.type_tick FROM company INNER JOIN tick ON company.company_id = tick.company_id WHERE company.company_name ILIKE " + fmt.Sprint("'"+companyName+"'")
	log.Print("🧁", singleCompanyQuery)
	rows, err := db.Query(singleCompanyQuery)
	checkError := utils.CheckError
	checkError(err)
	companies := make([]CompanyWithTick, 0)
	for rows.Next() {
		var company CompanyWithTick
		if err := rows.Scan(&company.Company.CompanyId, &company.CompanyName, &company.CompanyAddress, &company.CompanyMail, &company.CompanyPhone, &company.TickId, &company.TypeTick, &company.CoreService, &company.Date); err != nil {
			log.Fatal(err)
		}

		companies = append(companies, company)
	}
	return companies
}

type Company struct {
	CompanyName    string `json:"companyName"`
	CompanyId      string `json:"companyId"`
	CompanyAddress string `json:"companyAddress"`
	CompanyPhone   string `json:"companyPhone"`
	CompanyMail    string `json:"companyMail"`
}
type CompanyWithTick struct {
	Company
	Tick
}
