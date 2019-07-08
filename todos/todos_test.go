package todos

// func ExampleResponseRecorder() {
// 	req := httptest.NewRequest("GET", "/hello", nil)
// 	w := httptest.NewRecorder()
// 	TodosHandler(w, req)
// 	resp := w.Result()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(resp.StatusCode)
// 	fmt.Println(resp.Header.Get("Content-Type"))
// 	fmt.Println(string(body))

// 	// Output:
// 	// 200
// 	// text/plain; charset=utf-8
// 	// Hello, world!
// }
