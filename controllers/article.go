package controllers

import (
	"blog/models"
	"fmt"
	"net/http"
)

// ArticleItem 列表
func ArticleItem(w http.ResponseWriter, r *http.Request) {
	var goodsModel models.GoodsModel

	var goodsId int64
	var goodsName string
	selectField := models.SelectValues{
		"goods_id":   &goodsId,
		"goods_name": &goodsName,
	}

	where := models.WhereValues{}

	qr, err := goodsModel.Query(selectField, where, 0, 10)

	if err != nil {
		fmt.Println(err)
		//		http.NotFound(w, r)
		return
	}
	fmt.Println(qr)
}
