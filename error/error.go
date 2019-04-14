package error

// type HttpResponseErr struct {
// 	Code int	''
// 	Message string
// }

func CheckErr(err error){
	if err!=nil{
		panic(err)
	}
}
