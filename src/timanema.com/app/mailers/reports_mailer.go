package mailers

import (
  "github.com/robfig/revel"
  "net/smtp"
  "bytes"
  "io/ioutil"
  "os"
  "mime/multipart"
  "fmt"
  "path"
)

const CRLF = "\r\n"

func SendReport(reported_id string, recipients []string) error {
  return sendMail("mail", "a signature has been reported", recipients)
}

func sendMail(template, subject string, recipients []string) error {
  host := revel.Config.StringDefault("mail.host", "")
  full_url := fmt.Sprintf("%s:%d", host, revel.Config.IntDefault("mail.port", 25))
  c, err := smtp.Dial(full_url)
  if err != nil {
    return err
  }
  if ok, _ := c.Extension("STARTTLS"); ok {
    if err = c.StartTLS(nil); err != nil {
      return err
    }
  }
  if err = c.Auth(smtp.PlainAuth(
      revel.Config.StringDefault("mail.from", ""),
      revel.Config.StringDefault("mail.username", ""),
      getPassword(),
      host,
    )); err != nil {
       return err
  }
  if err = c.Mail(revel.Config.StringDefault("mail.username", "")); err != nil {
    return err
  }
  for _, addr := range recipients {
    if err = c.Rcpt(addr); err != nil {
      return err
    }
  }
  w, err := c.Data()
  if err != nil {
    return err
  }

  multiw := multipart.NewWriter(w)
  header := renderHeader("this is the subject", multiw)
  body := renderBody("mail", multiw)

  mail := []byte(header + body)
  fmt.Println(string(mail))
  _, err = w.Write(mail)
  if err != nil {
    return err
  }
  err = w.Close()
  if err != nil {
    return err
  }
  return c.Quit()
}

func renderHeader(subject string, multi *multipart.Writer) string{
  s := "Subject: " + subject + CRLF
  f := "From: " + revel.Config.StringDefault("mail.from", revel.Config.StringDefault("mail.username", "")) + CRLF
  content_type := "Content-Type: multipart/alternative; boundary=" + multi.Boundary() + CRLF + CRLF
  return s + f + content_type
}

func renderBody(template_name string, multi *multipart.Writer) string {
  body := ""
  contents := map[string]string{"plain": renderTemplate("mail", "txt"), "html": renderTemplate("mail", "html")}
  for k, v := range contents {
    body += "--" + multi.Boundary() + CRLF + "Content-Type: text/" + k + "; charset=UTF-8" + CRLF + CRLF + v + CRLF + CRLF
  }
  body +=  "--" + multi.Boundary() + "--" + CRLF + CRLF
  return body
}

func renderTemplate(template_name, mime string) string {
  var body bytes.Buffer
  template, _ := revel.MainTemplateLoader.Template("mailers/" + template_name + "." + mime)
  template.Render(&body, nil)
  return body.String()
}

func getPassword() string {
  password := ""
  wd, _ := os.Getwd()
  email_pwd_path := path.Clean(path.Join(wd, "./email.pwd"))
  password_byte, err := ioutil.ReadFile(email_pwd_path)
  if err != nil {
      fmt.Println(err)
      password = revel.Config.StringDefault("mail.password", "")
  }else{
    password = string(password_byte)
  }
  return password
}
