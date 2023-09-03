package vtung

type BalanceService service

func (s *BalanceService) GetWalletBalance() (interface{}, error) {
	path := "/wp-json/api/v1/balance?" +
		"username=" + s.client.Username +
		"&password=" + s.client.Password
	resp, err := s.client.Call("GET", path, nil)
	return resp, err
}
