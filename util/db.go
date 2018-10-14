package myDB

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
	"log"
  "os"
  "bufio"
	"strings"
)

var db *sql.DB
var dbHost string
var dbName string
var dbAccount string
var dbPassword string

func Query(sql string, args string) (o_rows *sql.Rows, o_err error) {
  rows, err := db.Query(sql, args)

  if err != nil {
    log.Println(err)
  }

  return rows, err
}

func GetStockName(stockNo string) (o_rows *sql.Rows, o_err error)  {
  stmt, err := db.Prepare("SELECT stock_name, stock_type FROM tbl_stock_list WHERE stock_no = ?")
  if err != nil {
    log.Fatalln(err)
  }
  defer stmt.Close()

  rows, err :=  stmt.Query(stockNo)
  if err != nil{
    log.Println(err)
  }

  return rows, err
}

func readini() {
  log.Println("-----------------Read INI Start-----------------")
  var err error
  file, err := os.Open("db.ini")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    tmpString := strings.SplitN(scanner.Text(),"=",2)
    if (tmpString[0] == "DB_HOST") {
      dbHost = tmpString[1]
    } else if (tmpString[0] == "DB_NAME") {
      dbName = tmpString[1]
    } else if (tmpString[0] == "DB_ACCOUNT") {
      dbAccount = tmpString[1]
    } else if (tmpString[0] == "DB_PASSWORD") {
      dbPassword = tmpString[1]
    }
  }

  log.Println("DB_HOST = ", dbHost)
  log.Println("DB_NAME = ", dbName)
  log.Println("DB_ACCOUNT = ", dbAccount)
  log.Println("DB_PASSWORD = ", dbPassword)
  log.Println("-----------------Read INI End-----------------")
}

func init() {
    var err error

    readini()

    log.Println("-----------------Open DB Start-----------------")

    dbConnectString := dbAccount + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName;
    log.Println("ConnectString = ", dbConnectString)

    db, err = sql.Open("mysql", dbConnectString)
    if err != nil {
      log.Fatalf("Open database error: %s\n", err)
    } else {
      log.Println("DB Connected")
    }
    err = db.Ping()
    if err != nil {
      log.Fatal(err)
    }
    log.Println("-----------------Open DB End-----------------")
}
