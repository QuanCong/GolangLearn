package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"config"
	"myutil"
	"fmt"
)

var (
	name string
	kecheng string
	score int
)

func printRows(rows *sql.Rows) {
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name, &kecheng, &score)
		myutil.LogError(err)
		fmt.Printf("%s\t%s\t%d\n", name, kecheng, score)
	}
}

func fetchData(db *sql.DB) {
	rows, err := db.Query("select *from score order by fenshu desc limit 10")
	myutil.LogError(err)
	printRows(rows)
}
func prepareQuery(db *sql.DB) {
	stmOut, err := db.Prepare("select *from score where sname=? order by fenshu desc limit 10");
	myutil.LogError(err)
	defer stmOut.Close()
	rows, err := stmOut.Query("李四")
	printRows(rows)//打印结果
}

func singleQuery(db *sql.DB) {
	stmt, err := db.Prepare("select count(*) from score where sname=?")
	myutil.LogError(err)
	defer stmt.Close()
	var count int
	err = stmt.QueryRow("song").Scan(&count)
	myutil.LogError(err)
	fmt.Println("count:", count)
}
func insertData(db *sql.DB) {
	stmIn, err := db.Prepare("INSERT INTO score(sname,kecheng,fenshu) VALUES(?,?,?)")
	myutil.LogError(err)
	defer stmIn.Close()

	res, err := stmIn.Exec("song", "english", 120)
	myutil.LogError(err)
	lastId, err := res.LastInsertId();
	rowCnt, err := res.RowsAffected();

	fmt.Printf("ID=%d,affected=%d\n", lastId, rowCnt)
}
func updateData(db *sql.DB) {
	stmIn, err := db.Prepare("UPDATE score SET fenshu=? where sname=?")
	myutil.LogError(err)
	defer stmIn.Close()

	res, err := stmIn.Exec(80, "song")
	myutil.LogError(err)
	lastId, err := res.LastInsertId();
	rowCnt, err := res.RowsAffected();

	fmt.Printf("ID=%d,affected=%d\n", lastId, rowCnt)
}

func txUpdateData(db *sql.DB) {
	tx, err := db.Begin()
	myutil.LogError(err)
	stmIn, err := tx.Prepare("UPDATE score SET fenshu=? where sname=?")
	myutil.LogError(err)
	defer stmIn.Close()
	res, err := stmIn.Exec(60, "song")
	myutil.LogError(err)
	lastId, err := res.LastInsertId();
	rowCnt, err := res.RowsAffected();
	fmt.Printf("ID=%d,affected=%d\n", lastId, rowCnt)
	tx.Commit()//或 tx.Rollback()
}

func main() {
	//"user:password@host/database"
	connStr := config.MysqlUser + ":" + config.MysqlPass + "@tcp(" + config.MysqlHost + ")/exercise"
	db, err := sql.Open("mysql", connStr)
	myutil.LogError(err)
	defer db.Close();
	/*fetchData(db)
	fmt.Println(">>>>>>>>>>>>>>>>")
	prepareQuery(db)
	fmt.Println(">>>>>>>>>>>>>>>>")
	singleQuery(db)
	fmt.Println(">>>>>>>>>>>>>>>>")
	insertData(db)*/
	/*res, err := db.Exec("INSERT INTO score(sname,kecheng,fenshu) VALUES(?,?,?)", "song", "english", 120)
	myutil.LogError(err)
	lastId, err := res.LastInsertId();
	rowCnt, err := res.RowsAffected();

	fmt.Printf("ID=%d,affected=%d\n", lastId, rowCnt)

	updateData(db)*/

	txUpdateData(db)
}
