package controllers

import (
  "github.com/robfig/revel"
  "timanema.com/app/models"
  "labix.org/v2/mgo/bson"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
  signatures := []models.Signature{}
  models.Signatures().All(&signatures, bson.M{"order": "-_id", "limit": 5})
  count, _ := models.Signatures().Count(nil)
	return c.Render(signatures, count)
}
