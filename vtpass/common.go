package vtpass

type CommonService service

func (s *CommonService) GetVariation(service_id string) (interface{}, error) {
	path := "/api/service-variations?serviceID=" + service_id
	goTvVariationCodes, err := s.client.Call("GET", path, nil)
	return goTvVariationCodes, err
}
