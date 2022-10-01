package dtos

type EmailDto struct {
	Origin    string
	EmailFrom string
	SMTPHost  string
	SMTPPass  string
	SMTPPort  int
	SMTPUser  string
}
