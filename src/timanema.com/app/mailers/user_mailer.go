package mailers

import (
  "github.com/tanema/revel_mailer"
)

type UserMailer struct {
  revel_mailer.Mailer
}

func (u UserMailer) SendReport(reported_id string) error {
  return u.Send(revel_mailer.H{
            "subject": "a signature has been reported",
            "to": []string{"timanema@gmail.com"},
            "reported_id": reported_id,
          })
}
