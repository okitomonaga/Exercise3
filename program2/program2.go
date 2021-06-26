package main

func getdigit(str string) []string {
	search_id := 0
	start_id := 0
	for i := 0; i < len(str); i++ { //1000の位を切り出すためのfor
		if str[search_id] == 'M' {
			search_id++
		} else {
			break
		}
	}
	str_1000 := str[start_id:search_id] //1000の位を切り出す
	start_id = search_id                //start_idを初期化
	for i := start_id; i < len(str); i++ {
		if str[search_id] == 'C' || str[search_id] == 'D' {
			search_id++
		} else {
			if str[search_id] == 'M' {
				search_id++
			}
			break
		}
	}
	str_100 := str[start_id:search_id] //100の位を切り出す
	start_id = search_id
	for i := start_id; i < len(str); i++ {
		if str[search_id] == 'X' || str[search_id] == 'L' {
			search_id++
		} else {
			if str[search_id] == 'C' {
				search_id++
			}
			break
		}
	}
	str_10 := str[start_id:search_id] //10の位を切り出す
	start_id = search_id
	str_1 := str[start_id:len(str)] //残りの1の位を切り出す

	return []string{str_1000, str_100, str_10, str_1}
}

func main() {

}
