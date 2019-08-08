package rest

//TxByHashHandler ,which is a http handler, query tx by hash.
// func TxByHashHandler(w http.ResponseWriter, r *http.Request) {
//
// 	args := mux.Vars(r)
//
// 	var hashStr string
//
// 	if v, ok := args["hash"]; ok {
// 		hashStr = v
// 	} else {
// 		w.Write([]byte("please provide the hash like this, http:ip/tx/{hash}"))
// 		return
// 	}
//
// 	dbInstance, err := db.GetDatabaseInstance("mongodb://localhost:27017", "binance")
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
//
// 	collection, err := dao.NewTxCollection(dbInstance)
//
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
//
// 	txInfo, err := collection.QueryTxInfoByHash(hashStr)
//
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
//
// 	bz, err := json.Marshal(txInfo)
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
//
// 	w.Write(bz)
// }
//
// //TxsHandler ,which is a http handler, query tx list by page as parameters.
// func TxsHandler(w http.ResponseWriter, r *http.Request) {
//
// 	pageAsInt64, sizeAsInt64 := int64(0), int64(0)
// 	args := mux.Vars(r)
//
// 	if v, ok := args["page"]; ok {
//
// 		i, err := strconv.ParseInt(v, 10, 64)
// 		if err != nil {
// 			pageAsInt64 = 1
// 		} else {
// 			pageAsInt64 = i
// 		}
// 	}
//
// 	if v, ok := args["size"]; ok {
// 		i, err := strconv.ParseInt(v, 10, 64)
// 		if err != nil {
// 			sizeAsInt64 = 10
// 		} else {
// 			sizeAsInt64 = i
// 		}
// 	}
//
// 	dbInstance, err := db.GetDatabaseInstance("mongodb://localhost:27017", "binance")
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
//
// 	collection, err := dao.NewTxCollection(dbInstance)
//
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
//
// 	txInfoArr, err := collection.QueryTxInfoByPage(pageAsInt64, sizeAsInt64)
//
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	bz, err := json.Marshal(txInfoArr)
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
//
// 	w.Write(bz)
// }
