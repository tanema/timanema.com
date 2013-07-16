package models

import (
  "github.com/robfig/revel"
  "labix.org/v2/mgo/bson"
  "github.com/tanema/mgorx"
  "time"
)

type Signature struct {
  mgorx.Document    "-"
  Id          bson.ObjectId "_id,omitempty"
  Name        string
  Email       string
  Png         string
  Word        string
  Comment     string
  Reported    bool
  Created_at  time.Time
}

func Signatures() *mgorx.Collection{
  return mgorx.GetCollection(Signature{})
}

func (signature *Signature) Validate(v *revel.Validation) {
  //v.Required(signature.Email).Message("Your Email is required for validation purposes and so I know who drew that lewd picture!")
  v.Required(signature.Png).Message("You need to draw something!")
  v.Required(signature.Word).Message("How did you misplace your word?")
}
