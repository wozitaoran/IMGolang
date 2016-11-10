package main

//mysql test
import (
	"database/sql"
	"fmt"
	//"time"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {

	//获取连接池 sql 包的 Close 方法只有 3个，除了 *sql.Db 是连接池对象，使用中是不会关闭的 其他的两个 Rows.Close 和 Stmt.Close 是需要关的
	dbtemp, err := sql.Open("mysql", "root:yxkj@tcp(192.168.19.37:3307)/lexiangccb?charset=utf8")
	checkErr(err)
	db = dbtemp
}

func main() {

	//插入数据   INSERT im_user SET token='hdfh89wfh923',insert_time=NOW(),beizhu='beizhu'

	stmt, err := db.Prepare("INSERT im_user SET token=?,insert_time=NOW(),beizhu=?")
	checkErr(err)

	res, err := stmt.Exec("hdfh89wfh923", "beizhu")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update im_user set token=? where user_id=?")
	checkErr(err)

	res, err = stmt.Exec("xuruiupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM im_user")
	checkErr(err)

	for rows.Next() {
		var user_id int
		var token string
		var insert_time string
		var beizhu string
		err = rows.Scan(&user_id, &token, &insert_time, &beizhu)
		checkErr(err)
		fmt.Println(user_id)
		fmt.Println(token)
		fmt.Println(insert_time)
		fmt.Println(beizhu)
	}

	//删除数据
	//	stmt, err = db.Prepare("delete from im_user where user_id=?")
	//	checkErr(err)

	//	res, err = stmt.Exec(id)
	//	checkErr(err)

	//	affect, err = res.RowsAffected()
	//	checkErr(err)

	//	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
