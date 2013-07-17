package tests

import (
  "github.com/robfig/revel"
  "github.com/tanema/revel_mock"
  "timanema.com/app/models"
)

type SignatureModelTest struct {
	revel.TestSuite
  c *revel.Controller
  s models.Signature
}

func (t *SignatureModelTest) Before() {
  t.c = revel_mock.MockController("","")
  t.s = models.Signature{
          Name: "Test Name",
          Word: "word",
          Png: "picture",
        }
}
func (t *SignatureModelTest) After() {}

func (t SignatureModelTest) TestValidateWord() {
  t.s.Validate(t.c.Validation)
  t.Assert(!t.c.Validation.HasErrors())
  t.s.Word = ""
  t.s.Validate(t.c.Validation)
  t.Assert(t.c.Validation.HasErrors())
  t.Assert(t.c.Validation.Errors[0].Message == "How did you misplace your word?")
}

func (t SignatureModelTest) TestValidatePng() {
  t.s.Validate(t.c.Validation)
  t.Assert(!t.c.Validation.HasErrors())
  t.s.Png = ""
  t.s.Validate(t.c.Validation)
  t.Assert(t.c.Validation.HasErrors())
  t.Assert(t.c.Validation.Errors[0].Message == "You need to draw something!")
}
