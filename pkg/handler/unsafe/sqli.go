package unsafe

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"vuln-go-app/pkg/conf"
)

func BadQueryUser(c *gin.Context) {
	db, err := sql.Open("mysql", conf.DataSourceName)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	{
		querySQL := "select username from test.user where id = " + c.Query("id")

		rows, err := db.QueryContext(c, querySQL)
		if err != nil {
			// c.String(200, "Error: %s", err)		// 避免报错注入
			log.Print(err)
			c.String(200, "execute sql error")
			return
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				panic(err)
			}
		}(rows)

		var name string
		for rows.Next() {
			err := rows.Scan(&name)
			if err != nil {
				panic(err)
			}
			break
		}
		c.JSON(200, gin.H{
			"username": name,
		})
	}
}

// /unsafe/query_user?id=%31%3b%73%65%6c%65%63%74%28%73%6c%65%65%70%28%33%29%29%3b ---> "1;select(3);"
// /unsafe/query_user?id=1
