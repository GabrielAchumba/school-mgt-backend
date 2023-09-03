package vtung

type CustomerVerificationService service

func (s *CustomerVerificationService) MeterNumberVerification(customer_id string, service_id string,
	variation_id string) (interface{}, error) {
	//https://vtu.ng/wp-json/api/v1/verify-customer?username=Frank&password=123456&customer_id=62418234034&service_id=ikeja-electric&variation_id=prepaid
	path := "/wp-json/api/v1/verify-customer?" +
		"username=" + s.client.Username +
		"&password=" + s.client.Password +
		"&customer_id=" + customer_id +
		"&service_id=" + service_id +
		"&variation_id=" + variation_id
	resp, err := s.client.Call("GET", path, nil)

	/* Sample Response on Success
	{"code":"success","message":"Customer details successfully retrieved","data":{"customer_id":"62418234034","customer_name":"FIRSTNAME LASTNAME","customer_address":"10 Example Street, Town, State","customer_arrears":"0.00","decoder_status":null,"decoder_due_date":null,"decoder_balance":null}}
	Sample Response on Failure
	{"code":"failure","message":"Invalid Meter Number"}
	*/
	return resp, err
}

func (s *CustomerVerificationService) SmartCardNumberVerification(customer_id string, service_id string) (interface{}, error) {
	//https://vtu.ng/wp-json/api/v1/verify-customer?username=Frank&password=123456&customer_id=62418234034&service_id=ikeja-electric
	path := "/wp-json/api/v1/verify-customer?" +
		"username=" + s.client.Username +
		"&password=" + s.client.Password +
		"&customer_id=" + customer_id +
		"&service_id=" + service_id
	resp, err := s.client.Call("GET", path, nil)

	/* Sample Response on Success
	{"code":"success","message":"Customer details successfully retrieved","data":{"customer_id":"62418234034","customer_name":"FIRSTNAME LASTNAME","customer_address":"10 Example Street, Town, State","customer_arrears":"0.00","decoder_status":null,"decoder_due_date":null,"decoder_balance":null}}
	Sample Response on Failure
	{"code":"failure","message":"Invalid Meter Number"}
	*/
	return resp, err
}
