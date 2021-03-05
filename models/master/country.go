package master

import (
	"net/http"
	"github.com/febrarisupaldi/go-precise/db"
	r "github.com/febrarisupaldi/go-precise/models"
)

type Country struct {
	Code	string `json:"country_code"`
	Name    string `json:"country_name"`
}

type Countries struct{
	Id      int    `json:"country_id"`
	Code	string `json:"country_code"`
	Name    string `json:"country_name"`
	CreatedOn string `json:"created_on"`
	CreatedBy string `json:"created_by"`
	UpdatedOn *string `json:"updated_on"`
	UpdatedBy *string `json:"updated_by"`
}

func AllCountry() (r.Response, error) {
	var obj Countries
	var arrobj []Countries
	var res r.Response
	db.Init()
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

func ShowCountry(country_id int)(Country, error){
	var obj Country
	db.Init()
	con := db.Conn()
	sqlStatement := "select country_code, country_name from precise.country where country_id = ?"
	err := con.QueryRow(sqlStatement, country_id).Scan(
		&obj.Code, &obj.Name,
	)
	defer con.Close()

	if err != nil {
		return obj, err
	}
	
	return obj, nil
}

func AddCountry(country_code string, country_name string, created_by string) (r.Response, error) {
	var res r.Response
	
	db.Init()
	con := db.Conn()
	sqlStatement := "insert into precise.country(country_code, country_name, created_by) values(?,?,?)"
	stmt, err := con.Prepare(sqlStatement)
	defer con.Close()
	
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

func UpdateCountry(country_id int, country_code string, country_name string, updated_by string, reason string)(r.Response, error){
	var res r.Response
	
	db.Init()
	con := db.Conn()
	tx, err := con.Begin()
	if err != nil {
		tx.Rollback()
		return res, err
	}

	sqlStatement := "update precise.country set country_code = ?, country_name = ?, updated_by = ? where country_id = ?"
	stmt, err := con.Prepare(sqlStatement)
	defer con.Close()

	if err != nil {
		tx.Rollback()
		return res, err
	}
	
	result, err := stmt.Exec(country_code, country_name, updated_by, country_id)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return res, err
	}

	sqlStatement = "set @username = ?, @reason = ?"
	stmt, err = con.Prepare(sqlStatement)
	//defer con.Close()

	if err != nil {
		tx.Rollback()
		return res, err
	}
	
	result, err = stmt.Exec(updated_by, reason)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = rowsAffected

	return res, nil

}

func DeleteCountry(id int, deleted_by string, reason string)(r.Response, error){
	var res r.Response

	db.Init()
	con := db.Conn()
	tx, err := con.Begin()
	if err != nil {
		tx.Rollback()
		return res, err
	}
	sqlStatement := "delete from precise.country where country_id = ?"
	defer con.Close()

	stmt, err := con.Prepare(sqlStatement)
	if err != nil{
		tx.Rollback()
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil{
		tx.Rollback()
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil{
		tx.Rollback()
		return res, err
	}

	sqlStatement = "set @username = ?, @reason = ?"
	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		tx.Rollback()
		return res, err
	}
	
	result, err = stmt.Exec(deleted_by, reason)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = rowsAffected

	return res, nil
}

func CheckCountry(tipe string, value string)(int, error){
	var sqlStatement string
	var rowCount int
	db.Init()
	con := db.Conn()

	if tipe == "code"{
		sqlStatement = "select count(*) from precise.country where country_code = ?"
	}else if tipe == "name"{
		sqlStatement = "select count(*) from precise.country where country_name = ?"
	}
	defer con.Close()

	err := con.QueryRow(sqlStatement, value).Scan(
		&rowCount,
	)

	if err != nil {
		return 0, err
	}

	return rowCount, nil

}