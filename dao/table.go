package dao

// func (dao *Dao) GetData() {
//
// 	rows, err := dao.Mysql
//
// }

// func (_ *Dao) GetTable() {
//
// 	var (
// 		id   int
// 		name string
// 	)
// 	rows, err := db.Query("select id, name from users where id = ?", 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		err := rows.Scan(&id, &name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Println(id, name)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
