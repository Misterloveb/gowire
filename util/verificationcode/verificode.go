package verificationcode

import "github.com/mojocn/base64Captcha"

var store = base64Captcha.DefaultMemStore

type VerifiCode struct {
	store  base64Captcha.Store
	driver base64Captcha.Driver
}

func NewVerifiCode(driver base64Captcha.Driver) *VerifiCode {
	res := &VerifiCode{
		driver: driver,
		store:  store,
	}
	return res
}
func (v *VerifiCode) Generate() (id, b64s string, err error) {
	c := base64Captcha.NewCaptcha(v.driver, v.store)
	id, b64s, err = c.Generate()
	return
}

func (v *VerifiCode) Verify(captid string, answer string) bool {
	return v.store.Verify(captid, answer, true)
}
