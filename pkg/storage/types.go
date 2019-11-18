package storage

import (
	"strings"
	"time"
	"unicode"
)

type Revision struct {
	ID          string  `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	URL         string  `db:"url" json:"url"`
	FileHashSum *string `db:"file_hash_sum" json:"file_hash_sum"`
	Removed     int     `db:"removed" json:"removed"`
	Added       int     `db:"added" json:"added"`
	CreatedAt   string  `db:"created_at" json:"created_at"`
}

type Vehicle struct {
	ID            string    `db:"id" json:"id"`
	Brand         string    `db:"brand" json:"brand"`
	Color         *string   `db:"color" json:"color,omitempty"`
	Number        *string   `db:"number" json:"number,omitempty"`
	BodyNumber    *string   `db:"body_number" json:"body_number,omitempty"`
	ChassisNumber *string   `db:"chassis_number" json:"chassis_number,omitempty"`
	EngineNumber  *string   `db:"engine_number" json:"engine_number,omitempty"`
	OVD           string    `db:"ovd" json:"ovd"`
	Kind          string    `db:"kind" json:"kind"`
	Status        Status    `db:"status" json:"status"`
	RevisionID    string    `db:"revision_id" json:"revision_id"`
	TheftDate     time.Time `db:"theft_date" json:"theft_date"`
	InsertDate    time.Time `db:"insert_date" json:"insert_date"`
}

// Status represent status of the vehicle search.
// Can be either "stolen" or "removed".
type Status string

const (
	StatusStolen  Status = "stolen"
	StatusRemoved Status = "removed"
)

type WantedVehicle struct {
	ID            string  `db:"id" json:"ID"`
	OVD           string  `db:"ovd" json:"OVD"`
	Brand         string  `db:"brand" json:"BRAND"`
	Color         *string `db:"color" json:"COLOR"`
	Number        *string `db:"number" json:"VEHICLENUMBER"`
	BodyNumber    *string `db:"body_number" json:"BODYNUMBER"`
	ChassisNumber *string `db:"chassis_number" json:"CHASSISNUMBER"`
	EngineNumber  *string `db:"engine_number" json:"ENGINENUMBER"`
	Kind          string  `db:"kind" json:"-"`
	TheftDate     string  `db:"theft_date" json:"THEFT_DATA"`
	InsertDate    string  `db:"insert_date" json:"INSERT_DATE"`
	Status        Status  `db:"status" json:"-"`
	RevisionID    string  `db:"revision_id" json:"-"`
}

// ParseKind splits lexeme into kind and other lexems.
//
// Example:
// ParseKind("Tesla - Model X (Легковий автотранспорт)")
//  =>
// other: "Hyundai - i30", kind: "Легковий автотранспорт"
func ParseKind(lexeme string) (other string, kind string) {
	stack := make([]rune, 0)

	for i := len(lexeme) - 1; i >= 0; i-- {
		switch lexeme[i] {
		case ')':
			stack = append(stack, ')')
		case '(':
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			other, kind = lexeme[:i], lexeme[i:]
			break
		}
	}

	if len(stack) != 0 {
		var tmp []rune
		i := 0
		for ; i < len([]rune(lexeme)) && lexeme[i] != '('; i++ {
			tmp = append(tmp, []rune(lexeme)[i])
		}
		other = string(tmp)
		kind = lexeme[i:]
	}

	other = strings.TrimFunc(other, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	})

	kind = strings.TrimFunc(kind, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	return other, kind
}
