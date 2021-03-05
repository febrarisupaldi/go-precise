package models

import (
	"database/sql"
	"strings"
	"github.com/febrarisupaldi/go-precise/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	user_id string
}

func Login(user_id, password string)(bool, int64, error){

	var obj User
	var pwd string

	db.Init()
	con := db.Conn()

	sqlStatement := "select user_id, password from users where user_id = ?"
	defer con.Close()
	err := con.QueryRow(sqlStatement, user_id).Scan(
		&obj.user_id, &pwd,
	)

	if err == sql.ErrNoRows{
		return false, 0, err
	}

	if err != nil{
		return false, 0, err
	}
	pwd = strings.ReplaceAll(pwd, "$2y$", "$2a$")
	checked, err := CheckPasswordHash(password, pwd)
	

	sqlStatement = "insert into precise.log_user_login (user_id,login_on) values(?,sysdate())"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return false, 0, err
	}

	result, err := stmt.Exec(user_id)

	if err != nil {
		return false, 0, err
	}

	getId, err := result.LastInsertId()
	if err != nil {
		return false, 0, err
	}

	if !checked{
		return false, 0, err
	}

	return true, getId, nil
}

func HashPassword(password string)(string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string)(bool, error){
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil{
		return false, err
	}

	return true, nil
}