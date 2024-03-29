package handler

import (
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/donations", handleGETDonations)
	http.HandleFunc("/donate", handlePOSTMakeDonation)
	http.HandleFunc("/payment_intent", handleGetPaymentIntent)
	http.HandleFunc("/log_donation", HandleLogDonation)

}
