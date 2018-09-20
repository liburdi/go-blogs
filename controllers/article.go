package controllers

import (
	"blog/config"
	"blog/models"
	"fmt"
	"html/template"
	"net/http"
)

func ArticleItem(w http.ResponseWriter, r *http.Request) {
	var goodsModel models.GoodsModel
	goodsModel.ModelProperty.TableName = "php41_goods"
	var typeModel models.TypeModel
	typeModel.ModelProperty.TableName = "php41_type"
	var goodsId, typeId int64
	var goodsName, typeName string
	selectField := models.SelectValues{
		"goods_id":   &goodsId,
		"goods_name": &goodsName,
	}
	typeField := models.SelectValues{
		"type_id":   &typeId,
		"type_name": &typeName,
	}
	where := models.WhereValues{}

	qr, err := goodsModel.Query(selectField, where, 0, 10)

	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}
	fmt.Println(qr)
	typeqr, err := typeModel.Query(typeField, where, 0, 10)

	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}
	fmt.Println(typeqr)

	t, err := template.ParseFiles(config.TemplateDir + "/Index/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, struct {
		TypeData []models.SelectResult
	}{
		TypeData: typeqr,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
