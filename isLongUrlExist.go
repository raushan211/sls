package main

import (
	"fmt"
	"log"
	"time"
)

func isLongUrlExist(long_link string) (bool, int, int) {
	count := 0
	id := 0
	//a := ""
	currentTime := time.Now().UTC()
	var expire_at string
	fmt.Println("currenttime", currentTime, "test")
	SQL := `SELECT COUNT(long_link) FROM link WHERE long_link=$1`

	rows := DB.QueryRow(SQL, long_link)

	err := rows.Scan(&count)
	if err != nil {
		log.Println(err)
		return false, count, id
	}

	if count > 0 {

		NeWSQL := `SELECT id,expire_at FROM link WHERE long_link=$1 ORDER by id DESC limit 1`

		Newrows := DB.QueryRow(NeWSQL, long_link)

		err2 := Newrows.Scan(&id, &expire_at)
		//layout := "2006-01-02 15:04:05.000Z"
		expire_time, e := time.Parse(time.RFC3339, expire_at)
		fmt.Println(expire_time)

		if e != nil {
			fmt.Println(e)
		}

		if currentTime.After(expire_time) {
			count := -1
			return true, count, id

		}
		if err2 != nil {
			//log.Fatal(err2)

		}
		return true, count, id
	}
	return true, count, id
}
