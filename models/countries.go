package models

import (
	"net/http"
	"time"

	"github.com/febrarisupaldi/go-learning-api/db"
)

type Country struct {
	Id      int    `json:"country_id"`
	Code	string `json:"country_code"`
	Name    string `json:"country_name"`
	CreatedOn string `json:"created_on"`
	CreatedBy string `json:"created_by"`
	UpdatedOn string `json:"updated_on"`
	UpdatedBy string `json:"updated_by"`
}

func GetAllCountry() (Response, error) {
	var obj Country
	var arrobj []Country
	var res Response

	con := db.Conn()

	sqlStatement := `select country_id, country_code, country_name, created_on, created_by, updated_on, updated_by
	from precise.country`

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Code, &obj.Name, &obj.CreatedOn, &obj.CreatedBy, &obj.UpdatedOn, &obj.UpdatedBy)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}

func AddCustomers(customer_name string, customer_address string, customer_contact string) (Response, error) {
	var res Response
	con := db.Conn()
	time := time.Now()
	sqlStatement := "insert into customers(customer_name, customer_address, customer_contact, created_at, sales_id) values(?,?,?,?,120)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(customer_name, customer_address, customer_contact, time)

	if err != nil {
		return res, err
	}

	getId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{"id": getId}

	return res, nil
}
