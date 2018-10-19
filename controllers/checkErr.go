package controllers

func checkErr(errMasg error) {
	if errMasg != nil {
		panic(errMasg)
	}
}
