package models

import (
  "github.com/robfig/revel"
  "labix.org/v2/mgo/bson"
  "time"
)

type Signature struct {
  Document    "-"
  Id          bson.ObjectId "_id,omitempty"
  Name        string
  Email       string
  Png         string
  Word        string
  Comment     string
  Created_at  time.Time
}

func Signatures() *Collection{
  return GetCollection(Signature{})
}

func (signature *Signature) Validate(v *revel.Validation) {
  v.Required(signature.Email).Message("Your Email is required for validation purposes and so I know who drew that lewd picture!")
  v.Required(signature.Word).Message("The word you are drawing is required")
  v.Required(signature.Png).Message("You need to draw something!")
}
