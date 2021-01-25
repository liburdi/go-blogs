package tools
/**
 * 公有ERROR处理
 * @params error errMsg
 */
func CheckErr(errMsg error) {
	if errMsg != nil {
		panic(errMsg)
		return
	}
}
