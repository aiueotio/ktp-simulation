package go_say_hello

import "fmt"
// DEFFER
func DefferKtp(){
	messeage := recover()
	if messeage != nil{
		fmt.Println("Error Messeage :",messeage)
	}
	fmt.Println("Aplication open")
}

func PanicsKtp(error bool){
	defer DefferKtp()
	if error {
		panic("ERROR 404")
	}
	fmt.Println("Aplication loading...")
}

type Filter func(string)string

func ToxicFilter(name string,filter Filter){
	fmt.Println("Hallo",filter(name))
}

func NameFilter(name string) string{
	if name == "anjing"{
		return "Gunakan Bahasa yang Baik"
	} else {
		return name
	}
}

func KtpRun(request ModelKTP)(string, string,string){
	return request.Name, request.Address, request.Age
}

func Hallo(name string)string{
	return "Hello " + name

}