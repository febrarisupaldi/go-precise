package models

import (
	"net/http"

	"github.com/febrarisupaldi/go-precise/db"
)

type Country struct {
	Id      int    `json:"country_id"`
	Code	string `json:"country_code"`
	Name    string `json:"country_name"`
	CreatedOn string `json:"created_on"`
	CreatedBy string `json:"created_by"`
	UpdatedOn *string `json:"updated_on"`
	UpdatedBy *string `json:"updated_by"`
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

func AddCountry(country_code string, country_name string, created_by string) (Response, error) {
	var res Response
	con := db.Conn()
	sqlStatement := "insert into precise.country(country_code, country_name, created_by) values(?,?,?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(country_code, country_name, created_by)

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

func UpdateCountry(country_id int, country_code string, country_name string, updated_by string)(Response, error){
	var res Response
	con := db.Conn()
	sqlStatement := "update precise.country set country_code = ?, country_name = ?, updated_by = ? where country_id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	
	result, err := stmt.Exec(country_code, country_name, updated_by, country_id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = rowsAffected

	return res, nil

}

func DeleteCountry(id int)(Response, error){
	var res Response

	con := db.Conn()
	sqlStatement := "delete from precise.country where country_id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil{
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil{
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil{
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = rowsAffected

	return res, nil
}
