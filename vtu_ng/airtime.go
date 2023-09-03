package vtung

type AirtimeService service

func (s *AirtimeService) PurchaseAirtime(phone string, network_id string, amount string) (interface{}, error) {
	//https://vtu.ng/wp-json/api/v1/airtime?username=Frank&password=123456&phone=07045461790&network_id=mtn&amount=2000
	path := "/wp-json/api/v1/airtime?" +
		"username=" + s.client.Username +
		"&password=" + s.client.Password +
		"&phone=" + phone +
		"&network_id=" + network_id +
		"&amount=" + amount
	resp, err := s.client.Call("GET", path, nil)

	/* Sample Response on Success
	{"code":"success","message":"Airtime successfully delivered","data":{"network":"MTN","phone":"07045461790","amount":"NGN2000","order_id":"3100"}}
	Sample Response on Failure
	{"code":"failure","message":"Your wallet balance (NGN1067.65) is insufficient to make this airtime purchase of NGN2000","order_id":"3289"} */
	return resp, err
}
