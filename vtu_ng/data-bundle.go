package vtung

type DataBundleService service

func (s *DataBundleService) PurchaseDataBundle(phone string, network_id string, amount string) (interface{}, error) {
	//https://vtu.ng/wp-json/api/v1/data?username=Frank&password=123456&phone=07045461790&network_id=mtn&variation_id=M1024
	path := "/wp-json/api/v1/data?" +
		"username=" + s.client.Username +
		"&password=" + s.client.Password +
		"&phone=" + phone +
		"&network_id=" + network_id +
		"&variation_id=" + amount
	resp, err := s.client.Call("GET", path, nil)

	/* Sample Response on Success
	{"code":"success","message":"Airtime successfully delivered","data":{"network":"MTN","phone":"07045461790","amount":"NGN2000","order_id":"3100"}}
	Sample Response on Failure
	{"code":"failure","message":"Your wallet balance (NGN1067.65) is insufficient to make this airtime purchase of NGN2000","order_id":"3289"} */
	return resp, err
}
