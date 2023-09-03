package vtung

type ElectricityService service

func (s *ElectricityService) PurchaseElectricity(phone string, meter_number string, service_id string,
	variation_id string, amount string) (interface{}, error) {
	//https://vtu.ng/wp-json/api/v1/electricity?username=Frank&password=123456&phone=07045461790&meter_number=62418234034&service_id=ikeja-electric&variation_id=prepaid&amount=3000
	path := "/wp-json/api/v1/electricity?" +
		"username=" + s.client.Username +
		"&password=" + s.client.Password +
		"&phone=" + phone +
		"&meter_number=" + meter_number +
		"&service_id=" + service_id +
		"&variation_id=" + variation_id +
		"&amount=" + amount
	resp, err := s.client.Call("GET", path, nil)

	/* Sample Response on Success
	{"code":"success","message":"Electricity bill successfully paid","data":{"electricity":"Ikeja (IKEDC)","meter_number":"62418234034","token":"Token: 5345 8765 3456 3456 1232","units":"47.79kwH","phone":"07045461790","amount":"NGN3000","amount_charged":"NGN2970","order_id":"4324"}}
	Sample Response on Failure
	{"code":"failure","message":"Invalid Meter Number","order_id":"3907"}
	*/
	return resp, err
}
