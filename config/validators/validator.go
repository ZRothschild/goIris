package validators

import (
	"fmt"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	viper2 "github.com/spf13/viper"
	"reflect"
	"unicode/utf8"
)

type ValidatorConf struct {
	Locale string
	Label  string
}

func NewValidators(viperKey string, viper *viper2.Viper) (validate *validator.Validate, trans ut.Translator, err error) {
	var (
		validatorCon ValidatorConf
		translator   locales.Translator
		uni          *ut.UniversalTranslator
	)

	if err = viper.UnmarshalKey(viperKey, &validatorCon); err != nil {
		return
	}

	// NOTE: ommitting allot of error checking for brevity
	if validatorCon.Locale == "zh" {
		translator = zh.New()
	} else if validatorCon.Locale == "en" {
		translator = en.New()
	}

	uni = ut.New(translator, translator)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ = uni.GetTranslator(validatorCon.Locale) // zh
	validate = validator.New()
	// 注册方法
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get(validatorCon.Label) // label
		if label == "" {
			return field.Name
		}
		return label
	})

	// 自定义函数checkName与 struct tag 关联起来
	if err = validate.RegisterValidation("checkName", checkName); err != nil {
		return validate, trans, err
	}

	if err := zhTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		return validate, trans, err
	}
	return validate, trans, err
}

// validate, trans, _ := validators.NewValidators()
// validators.TranslateAll(validate, trans)
// validators.TranslateIndividual(validate, trans)
// validators.TranslateOverride(validate, trans)

// 自定义校验函数
func checkName(f validator.FieldLevel) bool {
	// FieldLevel contains all the information and helper functions to validate a field
	count := utf8.RuneCountInString(f.Field().String()) // 通过utf8编码，获取字符串长度
	if count >= 2 && count <= 12 {
		return true
	}
	return false
}

func TranslateAll(validate *validator.Validate, trans ut.Translator) {
	type User struct {
		Username string `validate:"checkName"`
		// Tagline  string `validate:"required,lt=10"`
		// Tagline2 string `validate:"required,gt=1"`
	}

	user := User{
		Username: "Joeyb55555555555555loggs",
		// Tagline:  "This tagline is way too long.",
		// Tagline2: "1",
	}

	if err := validate.Struct(user); err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)
		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		fmt.Println(errs.Translate(trans))
	}
}

func TranslateIndividual(validate *validator.Validate, trans ut.Translator) {
	type User struct {
		Username string `validate:"checkName"`
	}
	var user User
	if err := validate.Struct(user); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
	}
}

func TranslateOverride(validate *validator.Validate, trans ut.Translator) {
	if err := validate.RegisterTranslation("checkName", trans, func(ut ut.Translator) error {
		return ut.Add("checkName", "{0} must have a value 赵桥旺!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("checkName", fe.Field())

		return t
	}); err != nil {
		fmt.Println(err)
	}

	type User struct {
		Username string `validate:"checkName"`
	}

	var user User
	if err := validate.Struct(user); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
	}
}
