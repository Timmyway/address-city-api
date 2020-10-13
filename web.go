package main

import ("fmt"
	"net/http"
	// "io/ioutil"
	"encoding/json"
)

func index_handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>My title</h1>
		<p style="background: indigo">Hey there</p>
		`)
}

type Lead struct {
	Email string `json:email`
	Firstname string `json:fname`
	Lastname string `json:lname`
	Zipcode string `json:zcode`
	Phone string `json:phone`
	Civility string `json:civ`
	AffiliateID string `json:affiliate_id`	
}

func echo_handle(w http.ResponseWriter, r *http.Request) {
	// resp, _ := http.Get("http://testim.kontikimedia.fr/echo.php?email=fr2devkontiki@gmail.com&fname=tim&lname=Kontiki&zcode=20000&phone=0033123456789&civ=madame&a=2889224")
	// bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	lead := Lead{Email: "fr2devkontiki@gmail.com", Firstname: "tim", Lastname: "Kontiki", Zipcode: "20000", Phone: "0033123456789", Civility: "madame", AffiliateID: "2889224"}	
	// resp.Body.Close()	
	json.NewEncoder(w).Encode(lead)
}

func loop_handle(w http.ResponseWriter, r *http.Request) {
	i:=0
	for i<10 {
		fmt.Println(i)
		i++
	}	
}

func map_handle(w http.ResponseWriter, r *http.Request) {
	couples := make(map[string]string)

	couples["Laka"] = "Mika"
	couples["Tim"] = "Irina"
	couples["Kevin"] = "Medina"

	fmt.Println(couples)

	myCouple := couples["Tim"]
	fmt.Println(myCouple)
	delete(couples, "Kevin")
	for k, v := range(couples) {
		fmt.Println("%s => %s", k, v)
	}
	json.NewEncoder(w).Encode(couples)
}

func main() {	
	http.HandleFunc("/", index_handle)
	http.HandleFunc("/echo", echo_handle)
	http.HandleFunc("/loop", loop_handle)
	http.HandleFunc("/map", map_handle)
	http.ListenAndServe(":8081", nil)
}