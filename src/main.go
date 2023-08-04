package main

import (
	"grafana_user_listerner/db_listener"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db_type := os.Getenv("db_type")
	db_username := os.Getenv("db_username")
	db_url := os.Getenv("db_url")
	db_password := os.Getenv("db_password")
	db_name := os.Getenv("db_name")
	graf_url := os.Getenv("graf_url")
	graf_admintoken := os.Getenv("graf_admintoken")
	org_id := os.Getenv("org_id")
	// 健康检查
	// http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "OK")
	// })

	// log.Fatal(http.ListenAndServe(":1300", nil))
	db_listener.NewUser_watch(db_type, db_username, db_password, db_url, db_name, graf_url, graf_admintoken, org_id)
}
