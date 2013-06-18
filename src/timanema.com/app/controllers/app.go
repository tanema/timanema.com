package controllers

import (
  "github.com/robfig/revel"
  "timanema.com/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
  signatures := []models.Signature{}
  models.Signatures().Where(&signatures, nil, nil)
	return c.Render(signatures)
}
