package go_say_hello

func KtpRun(request ModelKTP)(string, string,string){
	return request.Name, request.Address, request.Age
}