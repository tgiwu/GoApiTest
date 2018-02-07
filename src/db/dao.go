package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"models"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB = nil

func connectToDB(){
	if nil == db {
		cdb, err := sql.Open("mysql", "first:123456@tcp(127.0.0.1:3306)/firstDB?charset=utf8")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("connect success ")
			db = cdb
			db.SetMaxOpenConns(2000)
			db.SetMaxIdleConns(1000)
			db.Ping()
		}
	}
}

func InsertStudent(student *models.Student)  {
	if nil == db {
		connectToDB()
	}

	tx, txerr := db.Begin()
	if nil != txerr {
		log.Fatal(txerr)
	}
	defer tx.Rollback()

	stmt, err:= db.Prepare(`INSERT students (name, sex, age, tel) VALUE (?, ?, ?, ?)`)
	if nil != err {
		fmt.Println("1", err)
	}
	defer stmt.Close()

	fmt.Println("name ", student.Name)

	res, err:= stmt.Exec(student.Name, student.Sex, student.Age, student.Tel)
	if nil != err {
		fmt.Println("2", err)
	}

	id, err := res.LastInsertId()
	if nil != err {
		fmt.Println("3", err)
	}
	err = tx.Commit()
	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(id)
}

func QueryStudent()  {

	if nil == db {
		connectToDB()
	}
	id := 0
	name := "a"
	rows := db.QueryRow("SELECT id, name FROM students WHERE id = ? ", 3)
	rows.Scan(&id, &name)
	fmt.Println(id, name)


}

func GetAllStudent()  {
	if nil == db {
		connectToDB()
	}

	stmt, err := db.Prepare("SELECT * from students")
	if nil != err {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if nil != err {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if nil != err {
		fmt.Println(err)
		return
	}
	count := len(columns)

	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			fmt.Println("vaule", i, &values[i])
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if nil != err {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}

func getToJson(sqlString string) (string, error) {
	if nil == db {
		connectToDB()
	}

	stmt, err := db.Prepare(sqlString)
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	count := len(columns)

	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuesPtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuesPtrs[i] = &values[i]
		}

		rows.Scan(valuesPtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)

			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err :=json.Marshal(tableData)
	if nil != err {
		fmt.Println(err)
		return "", err
	}

	return string(jsonData), nil
}


func QueryAllStudent(){
	if nil == db {
		connectToDB()
	}
	rows, _ := db.Query("SELECT * FROM students")
	defer rows.Close()
	for rows.Next()  {

		var id int64
		var name string
		var sex  string
		var age  int16
		var tel  string
		var student = new(models.Student)
		rows.Columns()
		err := rows.Scan(&id, &name, &sex, &age, &tel)
		if nil != err {
			fmt.Println(err)
		}

		fmt.Println(id, name, sex, age, tel)
		student.Id = id
		student.Name = name
		student.Sex = sex
		student.Age = age
		student.Tel = tel

		b, er := json.Marshal(student)
		if er != nil {
			fmt.Println("encode faild")
		} else {
			fmt.Println(string(b))
		}
	}

}
