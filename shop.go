package yzgo

//Shop the youzan shop
type Shop struct {
	SID         string `json:"sid"`
	URL         string `json:"url"`
	PhysicalURL string `json:"physical_url"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	CertType    string `json:"cert_type"`
}
