package main

func main() {
	arr := [5]string{"I","am","stupid","and","weak"}

	for i,v := range arr {
		if v == "stupid" {
			arr[i] = "smart"
		}else if v == "weak" {
			arr[i] = "strong"
		}
	}

	for i, v := range arr {
		println(i,v)
	}



}
