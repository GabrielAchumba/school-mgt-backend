package vtung

type CableTVService service

func (s *CableTVService) PurchaseCableTV(phone string, service_id string, smartcard_number string,
	variation_id string) (interface{}, error) {
	//https://vtu.ng/wp-json/api/v1/tv?username=Frank&password=123456&phone=07045461790&service_id=gotv&smartcard_number=7032400086&variation_id=gotv-max
	path := "/wp-json/api/v1/tv?" +
		"username=" + s.client.Username +
		"&password=" + s.client.Password +
		"&phone=" + phone +
		"&service_id=" + service_id +
		"&smartcard_number=" + smartcard_number +
		"&variation_id=" + variation_id
	resp, err := s.client.Call("GET", path, nil)

	/* Sample Response on Success
	{"code":"success","message":"Cable TV subscription successfully delivered","data":{"cable_tv":"GOtv","subscription_plan":"GOtv Max","smartcard_number":"7032400086","phone":"07045461790","amount":"NGN3280","amount_charged":"NGN3247.2","service_fee":"NGN0.00","order_id":"2876"}}
	Sample Response on Failure
	{"code":"failure","message":"Invalid Smartcard Number","order_id":"3652"}
	*/
	return resp, err
}
