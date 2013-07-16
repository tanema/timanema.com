package controllers

import (
  "github.com/robfig/revel"
  "timanema.com/app/models"
  "timanema.com/app/mimes"
  "labix.org/v2/mgo/bson"
  "timanema.com/app/mailers"
  "math"
)

type Signature struct {
	*revel.Controller
}

func (c Signature) Index() revel.Result {
  limit := 10
  page := 0
  c.Params.Bind(&page, "page")
  signatures := []models.Signature{}
  models.Signatures().All(&signatures, bson.M{"order": "-_id","limit": limit, "skip": page*limit})
  count, _ := models.Signatures().Count(nil)
  var next_page, prev_page int
  page_count := int(math.Ceil(float64(count / limit))) + 1
  if page * limit < page_count - 1 {
    next_page = page + 1
  }
  if page > 0 {
    prev_page = page - 1
  }else{
    prev_page = -1
  }
  page++
  return c.Render(signatures, count, page, next_page, prev_page, page_count)
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

func (c Signature) Report(id string) revel.Result {
  var s models.Signature
  models.Signatures().Find(&s, id)
  s.Reported = true
  s.Save()
  err := mailers.UserMailer{}.SendReport(id)
  if err != nil {
    c.Flash.Error("There was a problem reporting this image.")
  } else {
    c.Flash.Success("This image has been reported and will be reviewed shortly.")
  }
	return c.Redirect(App.Index)
}
