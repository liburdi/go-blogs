package models

var tableName = "php41_goods"

// BlogModel 模型
type GoodsModel struct {
	ModelProperty
}

// GoodsProperty 属性
type GoodsProperty struct {
	Goods_id   string
	Goods_name string
}
