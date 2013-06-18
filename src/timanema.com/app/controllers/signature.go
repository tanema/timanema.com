package controllers

import (
  "github.com/robfig/revel"
  "timanema.com/app/models"
  "timanema.com/app/mimes"
)

type Signature struct {
	*revel.Controller
}

func (c Signature) Create(signature models.Signature) revel.Result {
  saved := models.Signatures().Create(&signature, c.Validation)
  if !saved || c.Validation.HasErrors() {
    return c.Render(signature)
  }
  return c.Redirect(App.Index)
}

func (c Signature) Show(id string) revel.Result {
  var s models.Signature
  models.Signatures().Find(&s, id)
	return mimes.Png(s.Png)
}
