package mailers

type UserMailer struct {
  Mailer
}

func (u UserMailer) SendReport(reported_id string) error {
  return u.Send(H{
            "subject": "a signature has been reported",
            "to": []string{"timanema@gmail.com"},
            "reported_id": reported_id,
          })
}
