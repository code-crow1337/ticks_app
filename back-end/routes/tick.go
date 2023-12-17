package routes

import (
	"database/sql"
	"e-identitet/tick-api/utils"
	"log"
)

func GetTicks() {

}
func GetTickByType(db *sql.DB, tickType string) []Tick {

	rows, err := db.Query(formatQuery(tickType))
	checkError := utils.CheckError
	checkError(err)
	ticks := make([]Tick, 0)
	for rows.Next() {
		var tick Tick
		if err := rows.Scan(&tick.CompanyId, &tick.TickId, &tick.CoreService, &tick.TypeTick, &tick.Date); err != nil {
			log.Fatal(err)
		}

		ticks = append(ticks, tick)
	}
	return ticks
}

func formatQuery(query string) string {
	switch query {
	case BankidAuth:
		return "SELECT * FROM tick WHERE type_tick = 'bankid-auth'"
	case BankidSign:
		return "SELECT * FROM tick WHERE type_tick = 'bankid-sign'"
	case SparLookup:
		return "SELECT * FROM tick WHERE type_tick = 'spar'"
	default:
		return "SELECT * FROM tick WHERE type_tick = 'sms'"
	}

}

const (
	BankidAuth = "bankid-auth"
	BankidSign = "bankid-sign"
	SparLookup = "spar"
	SmsAuth    = "sms"
)

type Tick struct {
	CompanyId   string `json:"CompanyId"`
	TickId      string `json:"tickId"`
	CoreService string `json:"coreService"`
	TypeTick    string `json:"typeTick"`
	Date        string `json:"date"`
}
