package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Palind(input int) bool {
	// for check if input 1 digit, its return true
	if input > 0 && input < 10 {
		return true
	}
	// convert integer to string
	new := strconv.Itoa(input)
	// looping to check palindrome or not
	for j := 0; j < len(new)/2; j++ {
		if new[j] != new[len(new)-1-j] {
			return false
		}
	}
	return true
}

func CountPalindrome(n, m int) int {
	var count int
	// count how much palindrome from n to m
	for i := n; i <= m; i++ {
		if Palind(i) {
			count++
		}

	}
	return count
}
func palindrome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		param1, _ := strconv.Atoi(r.FormValue("param1"))
		param2, _ := strconv.Atoi(r.FormValue("param2"))
		var err error

		result := CountPalindrome(param1, param2)
		response, err := json.Marshal(result)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
		return

	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/palindrome", palindrome)

	fmt.Println("Program Starting")
	http.ListenAndServe(":8080", nil)
}
