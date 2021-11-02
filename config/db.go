package config
import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)
func GetDB() (*sql.DB,error){
 DBDriver := os.Getenv("DB_DRIVER")  
 DBName := os.Getenv("DB_NAME")
 DBUser := os.Getenv("DB_USER")
 DBPassword := os.Getenv("DB_PASSWORD")
 DBURL := DBUser + ":" + DBPassword + "@/" + DBName
 db,err := sql.Open(DBDriver,DBURL)
 if err != nil{
  return db,err
 }
 
 return db,err
}