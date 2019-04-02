package validators

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	entrans "gopkg.in/go-playground/validator.v9/translations/en"
	zhtrans "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	uni       *ut.UniversalTranslator
	Validator UniversalValidator
)

type UniversalValidator struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

func InitValidate() UniversalValidator {
	Validator.Validate = validator.New()
	return Validator
}

func InitUniversalValidator(lang string, v UniversalValidator) UniversalValidator {
	switch lang {
	case "en-US":
		return registerEnTranslation(v)
	case "zh-CN":
		return registerZhTranslation(v)
	default:
		return registerEnTranslation(v)
	}
}

func registerEnTranslation(v UniversalValidator) UniversalValidator {
	enTr := en.New()
	uni = ut.New(enTr, enTr)
	trans, _ := uni.GetTranslator("en")
	v.Trans = trans
	_ = entrans.RegisterDefaultTranslations(v.Validate, v.Trans)
	return v
}

func registerZhTranslation(v UniversalValidator) UniversalValidator {
	zhTr := zh.New()
	uni = ut.New(zhTr, zhTr)
	trans, _ := uni.GetTranslator("zh")
	v.Trans = trans
	_ = zhtrans.RegisterDefaultTranslations(v.Validate, v.Trans)
	return v
}
