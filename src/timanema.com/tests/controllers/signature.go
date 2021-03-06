package tests

import (
  "github.com/robfig/revel"
  "timanema.com/app/controllers"
  "timanema.com/app/mimes"
  "github.com/tanema/revel_mock"
  "timanema.com/app/models"
  "labix.org/v2/mgo/bson"
  "net/url"
  "reflect"
)

type SignatureControllerTest struct {
	revel.TestSuite
}

func (t SignatureControllerTest) Before() { }
func (t SignatureControllerTest) After() { }

func (t SignatureControllerTest) TestIndexFunctional() {
	t.Get("/signatures")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t SignatureControllerTest) TestIndexResult() {
  result, ok := (controllers.Signature{revel_mock.MockController("Signature","Index")}.Index()).(*revel.RenderTemplateResult)
  t.Assert(ok) //succeeded rendering

  signatures := []models.Signature{}
  models.Signatures().All(&signatures, bson.M{"order": "-_id", "limit": 10})
  count, _ := models.Signatures().Count(nil)

  t.Assert(result.RenderArgs["count"] == count)
  t.Assert(reflect.ValueOf(result.RenderArgs["signatures"]).Len() == len(signatures))
}

func (t SignatureControllerTest) TestShowFunctional() {
  signatures := []models.Signature{}
  models.Signatures().All(&signatures, nil)
	t.Get("/signatures/"+signatures[0].Id.Hex())
	t.AssertOk()
	t.AssertContentType("image/png")
}

func (t SignatureControllerTest) TestShowResult() {
  signatures := []models.Signature{}
  models.Signatures().All(&signatures, nil)

  result, ok := (controllers.Signature{revel_mock.MockController("Signature","Show")}.Show(signatures[0].Id.Hex())).(mimes.Png)
  t.Assert(ok) //succeeded rendering

  t.Assert(string(result) == signatures[0].Png)
}

func (t SignatureControllerTest) TestReportFunctional() {
  signatures := []models.Signature{}
  models.Signatures().All(&signatures, nil)
	t.Get("/signatures/"+signatures[0].Id.Hex())
	t.AssertOk()
	t.AssertContentType("image/png")
}

func (t SignatureControllerTest) TestReportResult() {
  signatures := []models.Signature{}
  models.Signatures().All(&signatures, nil)
  signatures[0].Reported = false
  signatures[0].Save(nil)

  c := controllers.Signature{revel_mock.MockController("Signature","Report")}
  _, ok := (c.Report(signatures[0].Id.Hex())).(*revel.RedirectToActionResult)
  t.Assert(ok) //succeeded redirecting

  t.Assert(c.Flash.Out["success"] == "This image has been reported and will be reviewed shortly.")

  var s models.Signature
  models.Signatures().Find(&s, signatures[0].Id.Hex())
  t.Assert(s.Reported == true)
}

func (t SignatureControllerTest) TestCreateFuntional() {
  count, _ := models.Signatures().Count(nil)

  v := url.Values{}
  v.Set("signature.Name", "testers name")
  v.Set("signature.Comment", "This is the comment")
  v.Set("signature.Word", "test")
  //stephs snail
  v.Set("signature.Png", "data:image/png;base64,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")
	t.PostForm("/signatures", v)
	t.AssertOk()
	t.AssertContentType("text/html")

  newcount, _ := models.Signatures().Count(nil)
  t.Assert(newcount >= (count+1))
}

func (t SignatureControllerTest) TestCreateResult() {
  count, _ := models.Signatures().Count(nil)

  valid_sig := models.Signature{
    Name: "Testers Name",
    Word: "Test",
    Png: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAUUAAAGQCAYAAAA5h3LBAAAa70lEQVR4Xu3dv48tyVUHcHaNsYVARJDaJgfZCQkhEiKDFEHARvAn2ZEdGGLIcEJEgoTQAjm2EyQkAgIs+ffS99mzzBt331NdXdV1uuvzpJHsnbpdpz6n6vv63jt33ge/5A8BAgQIfCrwAQsCBAgQ+H8BoWg3ECBA4JWAULQdCBAgIBTtAQIECKwLuFO0MwgQIOBO0R4gQICAO0V7gAABAqGAp88hkQEECMwkIBRn6ra1EiAQCgjFkMgAAgRmEhCKM3XbWgkQCAWEYkhkAAECMwkIxZm6ba0ECIQCQjEkMoAAgZkEhOJM3bZWAgRCAaEYEhlAgMBMAkJxpm5bKwECoYBQDIkMIEBgJgGhOFO3rZUAgVBAKIZEBhAgMJOAUJyp29ZKgEAoIBRDIgMIEJhJQCjO1G1rJUAgFBCKIZEBBAjMJCAUZ+q2tRIgEAoIxZDIAAIEZhIQijN121oJEAgFhGJIZAABAjMJCMWZum2tBAiEAkIxJDKAAIGZBITiTN22VgIEQgGhGBIZQIDATAJCcaZuWysBAqGAUAyJDCBAYCYBoThTt62VAIFQQCiGRAYQIDCTgFCcqdvWSoBAKCAUQyIDCBCYSUAoztRtayVAIBQQiiGRAQQIzCQgFGfqtrUSIBAKCMWQyAACBGYSEIozddtaCRAIBYRiSGQAAQIzCQjFmbptrQQIhAJCMSQygACBmQSE4kzdtlYCBEIBoRgSGUCAwEwCQnGmblsrAQKhgFAMiQwgQGAmAaE4U7etlQCBUEAohkQGECAwk4BQnKnb1kqAQCggFEMiAwgQmElAKM7UbWslQCAUEIohkQEECMwkIBRn6ra1EiAQCgjFkMgAAgRmEhCKM3XbWgkQCAWEYkhkAAECMwkIxZm6ba0ECIQCQjEkMoAAgZkEhOJM3bZWAgRCAaEYEhlAgMBMAkJxpm5bKwECoYBQDIkMIEBgJgGhOFO3rZUAgVBAKIZEBhAgMJOAUJyp29ZKgEAoIBRDIgMIEJhJQCjO1G1rJUAgFBCKIZEBBAjMJCAUZ+q2tRIgEAoIxZDIAAIEZhIQijN121oJEAgFhGJIZAABAjMJCMWZum2tBAiEAkIxJDKAAIGZBITiTN22VgIEQgGhGBIZQIDATAJCcaZuWysBAqGAUAyJDCBAYCYBoThTt62VAIFQQCiGRAYQIDCTgFCcqdvWSoBAKCAUQyIDCBCYSUAoztRtayVAIBQQiiGRAQQIzCQgFGfqtrUSIBAKCMWQyAACBGYSEIozddtaCRAIBYRiSGQAAQIzCQjFmbptrQQIhAJCMSQygACBmQSE4kzdtlYCBEIBoRgSGUCAwEwCQnGmblsrAQKhgFAMiQwgQGAmAaE4U7etlQCBUEAohkQGECAwk4BQnKnb1kqAQCggFEMiAwgQmElAKM7UbWslQCAUEIohkQEECMwkIBRn6ra1EiAQCgjFkMgAAgRmEhCKM3XbWgkQCAWEYkhkAAECMwkIxZm6ba0ECIQCQjEkMoAAgZkEhOJM3bZWAgRCAaEYEhlAgMBMAkJxpm5bKwECoYBQDIkMIEBgJgGhOFO3rZUAgVBAKIZEBhAgMJOAUJyp29ZKgEAoIBRDIgMIEJhJQCjO1G1rJUAgFBCKIZEBBAjMJCAUZ+q2tRIgEAoIxZDIAAIEZhIQijN121oJEAgFhGJIZAABAjMJCMWZum2tBAiEAkIxJDKAAIGZBITiTN22VgIEQgGhGBIZQIDATAJCcaZuWysBAqGAUAyJDCBAYCYBoThTt62VAIFQQCiGRAYQIDCTgFCcqdvWSoBAKCAUQyIDCBCYSUAoztRtayVAIBQQiiGRAQQIzCQgFGfqtrUSIBAKCMWQyAACBGYSEIozddtaCRAIBYRiSGQAAQIzCQjFmbptrQQIhAJCMSQygACBmQSE4kzdtlYCBEIBoRgSGUCAwEwCQnGmbltrBoGfrhTxYYbC1PAzAaFoJxA4R+ARhs/Om7N4Th/CWTQiJDJgcoHXd3a1d3RRID6IP1m+aq8/eYvaLl8otvV0tXsJrIVZzZl5BF7Jn68sgz4uGWhMP4GaBverxpUJjBV4+3rf2vnYe0f3T8uSfm/Hsj5axn5jx3hDGwsIxcagLndZgZKnuC+Le3vn973lG7++sfLSu8SXh//38j9+87KKNyhcKN6giZZwWGBPIG5NtnUHuRWKj6fK/7J8rZ1B5/JwS+svAL/eziPvI7D3bm5r5VtPt9+O//byH770hM+5HLi34A/EN3UagbND8TGfH89J0/73CxGKSRujrFMFzg7FaHHOZSTU8fvwO+K69GUEeoZizeuVzuXArQN/IL6p0wjUBNda8VvnaW/oOpcDtwb8gfimTiWwN7j2hOLe0HUuB24N+APxTZ1OYG94vV5A9EPde67tXA7cGvAH4ps6rcCeAHssIgrEl4WWXte5HLg14A/EN3V6gdYhVvIUvTRg0+NdtUCheNXOqftMge8vk31uZcI9AVYSsP+7zLH1ccEz1zv1XEJx6vZbfKHAHy7jvrUytvT8PAvEf1+u+7uFdRh2gkBpU08oxRQEUgusPfUtPT9bT5sfYfmZ1KuesLjSpk5IY8kE3hPoEYrOX8JNpikJm6KklAK1oeguMWU7t4sSihdrmHKHCdSE4rN3m529Ya18PrHGJG2MstIJtAzFPe9ap4O4e0FC8e4dtr5WAntD8d+WiX9nY3LnrlVXOlxHczqguuQtBfaG4g8Xhc+uSDhzybeHBiVvkPLSCOwNxa3XE525NC1dL0SDkjdIeWkEhGKaVvQtRCj29XX1+wi0CMUfLRy/ch+Se65EKN6zr1bVXmBPKH51mf4vV0rwkb72fWl+RaHYnNQFbyqwJxS9nnjhTSAUL9w8pZ8qIBRP5R43mVAcZ2/mawmUhuI3l2X92crSvrb8t7+61pLnrFYoztl3q94vUBqKnjrvt031CKGYqh2KSSwgFBM3p2VpQrGlpmvdWeBIKP71AvPnd8a509qE4p26aS09BY6EonPWszONr61ZjUFd7rYCQvG2rX1/YUJxkkZb5mEBoXiY8BoXEIrX6JMqxwusheLa70UsDc/xK1LBqoBQtDEIPBd4/ONSjz9bZ+VtCK6N+4fl8X8A+hoCQvEafVLlGIFn/5zA3oqctb1ig8Zr1CB406YXKPnH6/cswlnbozVwrEYNxDd1aoGWd4nPnn6nRpixOKE4Y9etuUSgZSj6h6pKxJOMEYpJGqGMdAKtQlEgpmvt84KE4sUaptzTBGp+tGYrSAXjaW07PpFQPG7oCvcUqAnFZ2/OOGsX2ScadZFGKfN0gZpQfBTpbvH0VrWdUCi29XS1+wjUhuKzu0VPoy+wP4TiBZqkxCECtaH47G7x8T1nbkg7yyfVoHIrI+cSOBKK31+oPrfB5cwl30calLxByhsmcCQUH0VvPY3+wfK9zw9blYlDAaEYEhkwqcDRUHz2NNq5S7ypNCdxcwpLe/ktLoXDnw77sMVFbnINoXiTRu5dhlDcK5ZjfOtfVvB6Va/DYOaQbBGKW68tOnc5ztFqFZqTuDkrpT17Ab/XSp593O3OodkiFLeeQvvRnF67tcF1hWIDxJMu8ZNlnowh9DY8MtZY06Keofiox9mr6coJj9GYE5AbTNHz6XKD8n7hEnd4Ci4Ue+yMC1xTKOZu0tXC8Jnm1Z4ytgrFrR46e0nPnsbka0z0b4K8fVOkxdPVM8P3KuHYKhS3Xld09vKdvXcVaUyuxuwJp8fYz3Qqf08dtSVkD0ehWNvZiz9OKOZqYOkvNh39qYiWoZn1jRqhmOtsnFaNUDyNOpyoJGiy312VrCGEWAaMXmfrny9sGbAlfsYcEBCKB/AaP/TZXeLokDiy1CNBOWrdW72oPS9C8cgOOvmxtU0+ucwppmt9ELOhHQnHs/dp614IxWy78Uk9Z2+2C9GcXurawRl1p9Rz8bXheOZeXevFkddxhWLPHdX42mdutMal3+5ysx2cveF45l8QrXvR+nq32/yZFiQU83TjzIMT/WadFj/7WCMbBeVZwbi3F3+/LPb3l6/HRzG/t3z95/L1z8vX15avj3/+xtFbD2evZoec8BiNOQG5YIrW73Y+mzIKnpfHjvxRmehHk3rv2z2h+F8L2G89Af9o+d7XV77few0F286QNQGNybEvWr+wv7Wq0kDcevxZd2qP+aNg7FnLnlCM6vzOspYvCsUcB62kCqFYotR/zBmheDQQXyu81NvzaXZJvb2CsWUofneB+4JQ7H+IWs0gFFtJHrtO63c716qJ7mhqV9AzIEcF455Q/J8F7jc8fa7dPvkeJxRz9GTPIaytuFcovtTT666t5Kl06328tx8/Wor85VeNebzZ8s3l66vLlzdaanfsoMe13kyDlnH5afcewr0L3rrjWguykruzrfl7vjnzLNRb7+OafnxrQXm8A/2Py9cfvQGqud7eHhvfSKD1ZmpU1nSX6X1ojrxmeSQkt+7yal+LPON3E/b4SYDe/Z3uwPRcsFDsqVt+7d6HZu36NU93jwbka5HavTfC6lF3bb1bfzEcuV75zjJyt4DG7Cbr8oARB/1I71uEY00onxEwR+6qS19WOBqyXTahi/5M4MjBYNhO4Gqh+LLyo+FYs/9GWB353PMZT/nb7URXEopJ9sCIg14TSGtcR4KxpoZWLwXsWUtNnS/X73HnmWTb3rOMI82+p8iYVV05FN+K7QnJmv3XM2R6XLtniI/ZrTeftWZT3pxkyPLuFIq9X1fr+XS0R4D17u2QDXvnSYViju72Pji9r1+i2LKGHuHVI2x7XLPE2pgDAkLxAF7Dh7YMjLWyel+/hKJlDWc9zT3yZuSzlxGcu5IdM2iM5gyCfzNty8CYIRRb34Ft/cB27Y8NPXqwFdxHgjbHbr15FUIxR4OF4v4+tDQ7887zSNDuV/KI3QJCcTdZlwe0POCXvlP85JNPvrws4I+Xr7/74IMPHr9MYetPS7OzQlEgdjk+bS8qFNt61l6t5QG/bCgugfgXS/Fff7WAj5Zg/MYGakuztWsd+YHtHp+frt1bHrdTQCjuBOs0vOUBbx2Ka/+eS80vdAjXuITi299N+J0lFL80KBSPnI0ed56dtp7LvhU40nia7QTCwDg4Ve31wzcLliD7NDSXAHsWlmENy7V+Ycxyza09Gl6v0Kz1mzaPaYViIX7GYUIxR1daHfCWr79Fn0x5ybD39tCREBsUij0CrPXT8Ry7dJIqhGKORmcMxWd3ie/UVm7s3v3njTvGcI2JQvHoGyLhWnNsO1WsCQjFHPui9yGquX5tKP7Sxt1iWEOiUDx6LsK15th2qhCKefdA70NUc/0oFB8ZthoeFwrFHq8nbr2meDRo8+7em1WmUTkauhZAR5/CvV5Zq1B8XOexZ97VtnZnd7Gnzz1eT+wVtDl26gRVCMUcTe5xOHuE4nv7ZSMUP17uFL+ywhoG88lPn/f8Y157dknvXu6pxdgKAaFYgdbhIb3vLsJAahVij+tc5Olzr/DqfdffYfu55GsBoZhnP9QEV2n1NdcOH7Nxp/jeHeqrd6KrrrcRsEc/MbL1+HeZXoq6Mu4nj5cVVv77kWseKMdDawQ0q0atz2PC0Dgwbc21t+54Pi1jCcV/Xf7P47PKT/98+OG7lx/X9tqn/+3nPwT+C2N23HXuCbStu8Sjr+P2vuOPqH2/gYBQbIDY6BI1wVU6dc21o3efS+feHPfqTnNzH54cikfPg6fOh3fF+Asc3QTjV3CfCmqCq3T1Ndc+IxTD+ldCcetubM8vcFhb257Hb9W9dt1HvZ8JF2pAGgGhmKYVq5+XbdWfdKH4+EjzT3+69rsm3mvIj5dxn33ToqNvkPR8ilvjnGcHquSdQKtDh/O4QM8DVXPt6LPP1SsuDMQ9Hxfc81rg0VB9tu4a52pHD+wjIBT7uNZcteeBqr1202AsDMOH3VYgvvveCu6efdzzdb+jtdXsG49pLLBnMzWe2uUKnha26k+Pw1ocmDvC8EHyg2X855/sjqNrOfp4d4o3P7qtDt3NmU5ZXs8X6XsGwSk4ryY5spaeT51b3MWebWm+g087APYV8LQu9j3yQ9vP3k1vdXNwJLDj1RtxikCrzXBKsTefxLuicYNr7/SiHy9qdQ6EYtzD9CNabYb0C71AgT0/InaHw1r70bwoEPe8cx1tozs4R2u8/feFYq4W9zpUva57pt5WuEU/dP0sFFsGotcUz9wNHecSih1xKy7dK7x6XbdiiVUPqb1L3Lr7fhTRY+9f3bmqOXd7UI+NcTejM9fT61D1uu5ZNrV3iT1fp3279jPnOst9ynmEYq629wqvXtc9Q+/Iu8Y939F/vfZnP7PpjJ2xSxrOoWENMRtcqld49bpugyU/vcSzQIxeS9x6yt3jFzQcCe7ehq6/U0Ao7gTrPLxXePW6bk+O6F3jaO/W/vhOzZq25mr9Rk5NbR6zUyDaWDsvZ/hBgV7hddbTyIPL//ThRwOx9o2ZmvoFYo1a4scIxVzN6XXAoh9Lea2w9uv0z1Iq+Tx1yZ6tfWNm7zq9lrhX7ALjSzbYBZYxtMTwlwIWVPe4s/nV5Su6Q4peR9uaKrru2uNqHlOw1NUhpfuwZNyZr+/1+kus1tHjGgiUbLIG09zyEiV3NT0WviesXu76RtXacv3RXo3WWPsXSs1fNFGtLV1cq7GA5tWBRgew7qp9HvU6RK/a72d1P3v98LVoy7WfeTfaZ1e46qZAy40yC/OVAvF1T17eCb1S/dG7t6VrabXPo/miemc5I5deZ6vNcmmEHcXveeq647KnDb3KXWMULlE49bhDLJnTeTptK/ebSBPLba8eiCWvjT1egyw5/OVq5SOjIHxcaU9tJdcrra6k985SqWbycRpZ1qD/WIZ96cnQIwfwe8t1H+88r/15hMCPl68fLl+/Vlbq4VGvA2Dkj+e8XUhJMD0ec6QXWz2Izkn0/cNNcYHzBDSzzDo6kGc5/s1S7p++Kvmj5X9/Y2UJe+6oIoGXtZ8ZkK9/zGmPbet3mF9CdsuodQBHvfD9EwT2bLgTykk5xXeWqr6wUdmIQ/HlpZY/Wb7+dvn6+IlYy2B8mSb6y+FIcL4EYc2e7NWHZ4a95kx5CGYqqmYDzuTz7E7h28s3fzs5Ro9gjJYcBefa42v3Yc9gEohRp2/6/drNeFOO1WVtHfKr240IzJb7pqe/QGzZqYtdq+fGuhjFZrlrofjdZfQXb7LAq4Vjr7vD0qfvzsxNNv7WMjQ4bvBaKN7dLVtQ9grCR/f3rLVnHfFONOIUgbsf7haIM4bia7c9odHC++V13CNv2pTUUbMu56VE9uJjNDlu4OyhGAvtu9vaul6Lu7CS31hUs+db1FbiaEwCgZoNkqDsU0sQiqdyV01Wc9dXOpFALJW6yTihGDdSKMZGI0f0CERBOLKjg+cWinEDsoRiyVPDt6vp/bpcrNd3ROmvDSutQhiWSt14nFCMm7sWir0PT+3H3NZWM+JjerFq/YiWNi9V9O5n/Wo98nQBoRiTb/3wdu1Biu74evZkay0j7ygjj9cdamVT27t4txhxeYFWm+zyEE8W8Oxja5Hf2wMfjR/pWPPxvKP19vIQekc7M/Hje23KO5FGH/naWivb83eBMDzf/HYzOrhlLe3xDmfZzEaVCgjEUinjngoIxfINMuLpZe0bATOFuDAs38NGFggIxQKknw85OxRbHvY7hWRLl/LuGzmNgFAsb3XLYMl0sFuuq1zz/ZGZPGrX4HE3ERCK+xpZEyAO/D5jowkMFRCKQ/lNToBANgGhmK0j6iFAYKiAUBzKb3ICBLIJCMVsHVEPAQJDBYTiUH6TEyCQTUAoZuuIeggQGCogFIfym5wAgWwCQjFbR9RDgMBQAaE4lN/kBAhkExCK2TqiHgIEhgoIxaH8JidAIJuAUMzWEfUQIDBUQCgO5Tc5AQLZBIRito6ohwCBoQJCcSi/yQkQyCYgFLN1RD0ECAwVEIpD+U1OgEA2AaGYrSPqIUBgqIBQHMpvcgIEsgkIxWwdUQ8BAkMFhOJQfpMTIJBNQChm64h6CBAYKiAUh/KbnACBbAJCMVtH1EOAwFABoTiU3+QECGQTEIrZOqIeAgSGCgjFofwmJ0Agm4BQzNYR9RAgMFRAKA7lNzkBAtkEhGK2jqiHAIGhAkJxKL/JCRDIJiAUs3VEPQQIDBUQikP5TU6AQDYBoZitI+ohQGCogFAcym9yAgSyCQjFbB1RDwECQwWE4lB+kxMgkE1AKGbriHoIEBgqIBSH8pucAIFsAkIxW0fUQ4DAUAGhOJTf5AQIZBMQitk6oh4CBIYKCMWh/CYnQCCbgFDM1hH1ECAwVEAoDuU3OQEC2QSEYraOqIcAgaECQnEov8kJEMgmIBSzdUQ9BAgMFRCKQ/lNToBANgGhmK0j6iFAYKiAUBzKb3ICBLIJCMVsHVEPAQJDBYTiUH6TEyCQTUAoZuuIeggQGCogFIfym5wAgWwCQjFbR9RDgMBQAaE4lN/kBAhkExCK2TqiHgIEhgoIxaH8JidAIJuAUMzWEfUQIDBUQCgO5Tc5AQLZBIRito6ohwCBoQJCcSi/yQkQyCYgFLN1RD0ECAwVEIpD+U1OgEA2AaGYrSPqIUBgqIBQHMpvcgIEsgkIxWwdUQ8BAkMFhOJQfpMTIJBNQChm64h6CBAYKiAUh/KbnACBbAJCMVtH1EOAwFABoTiU3+QECGQTEIrZOqIeAgSGCgjFofwmJ0Agm4BQzNYR9RAgMFRAKA7lNzkBAtkEhGK2jqiHAIGhAkJxKL/JCRDIJiAUs3VEPQQIDBUQikP5TU6AQDYBoZitI+ohQGCogFAcym9yAgSyCQjFbB1RDwECQwWE4lB+kxMgkE1AKGbriHoIEBgqIBSH8pucAIFsAkIxW0fUQ4DAUAGhOJTf5AQIZBMQitk6oh4CBIYKCMWh/CYnQCCbgFDM1hH1ECAwVEAoDuU3OQEC2QSEYraOqIcAgaECQnEov8kJEMgmIBSzdUQ9BAgMFRCKQ/lNToBANgGhmK0j6iFAYKiAUBzKb3ICBLIJCMVsHVEPAQJDBYTiUH6TEyCQTUAoZuuIeggQGCogFIfym5wAgWwCQjFbR9RDgMBQAaE4lN/kBAhkExCK2TqiHgIEhgoIxaH8JidAIJuAUMzWEfUQIDBUQCgO5Tc5AQLZBIRito6ohwCBoQJCcSi/yQkQyCYgFLN1RD0ECAwV+D/8fTe+SeECWgAAAABJRU5ErkJggg==",
  }
  invalid_sig := models.Signature{}

  c := controllers.Signature{revel_mock.MockController("Signature","Create")}
  _, ok := (c.Create(invalid_sig)).(*revel.RenderTemplateResult)
  t.Assert(ok) //re renders form to correct issues
  c.Validation = &revel.Validation{}
  _, ok = (c.Create(valid_sig)).(*revel.RedirectToActionResult)
  t.Assert(ok) //redirects to index

  newcount, _ := models.Signatures().Count(nil)
  t.Assert(newcount >= (count+1))
}
