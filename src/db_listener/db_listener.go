package db_listener

import (
	"database/sql"
	"fmt"
	"grafana_user_listerner/grafana_org_invite"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
}

func NewUser_watch(db_type string, db_username string, db_password string, db_url string, db_name string, graf_url string, graf_admintoken string, org_id string) {
	grafana_url := graf_url
	graf_admtoken := graf_admintoken
	driverName := db_type
	dataSourceName := db_username + ":" + db_password + "@tcp" + "(" + db_url + ")" + "/" + db_name
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 定义用于存储用户名的数组
	var usernames []string

	for {
		// 清空新用户数组
		usernames = usernames[:0]
		// 执行查询
		rows, err := db.Query("SELECT login FROM user WHERE created > NOW() - INTERVAL 5 SECOND")
		if err != nil {
			log.Fatal(err)
		}

		// 遍历查询结果
		for rows.Next() {
			var username string

			// 通过Scan函数将查询结果赋值给username变量
			if err := rows.Scan(&username); err != nil {
				log.Fatal(err)
			}

			// 输出新增用户的用户名
			// fmt.Println("New User:", username)

			// 将用户名添加至usernames数组
			usernames = append(usernames, username)
			for _, user := range usernames {
				// 输出新增用户的用户名
				fmt.Println("新用户: ", user)
				grafana_org_invite.PuPuTech_Invite(user, grafana_url, graf_admtoken, org_id)
			}

		}
		rows.Close()

		// 每次查询完成后，等待5秒
		time.Sleep(5 * time.Second)
	}
}
