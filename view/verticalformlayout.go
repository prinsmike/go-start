package view

import (
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
)

/*
VerticalFormLayout.

CSS needed for VerticalFormLayout:

	form label:after {
		content: ":";
	}

	form input[type=checkbox] + label:after {
		content: "";
	}

Additional CSS for labels above input fields (except checkboxes):

	form label {
		display: block;
	}

	form input[type=checkbox] + label {
		display: inline;
	}

DIV classes for coloring:

	form .required {}
	form .error {}
	form .success {}

*/
type VerticalFormLayout struct {
}

func (self *VerticalFormLayout) BeginFormContent(form *Form, formFields Views) Views {
	return formFields
}

func (self *VerticalFormLayout) SubmitSuccess(message string, form *Form, formFields Views) Views {
	return append(formFields, form.GetFieldFactory().NewSuccessMessage(message, form))
}

func (self *VerticalFormLayout) SubmitError(message string, form *Form, formFields Views) Views {
	return append(formFields, form.GetFieldFactory().NewGeneralErrorMessage(message, form))
}

func (self *VerticalFormLayout) EndFormContent(fieldValidationErrs, generalValidationErrs []error, form *Form, formFields Views) Views {
	fieldFactory := form.GetFieldFactory()
	for _, err := range generalValidationErrs {
		debug.Dump(err)
		formFields = append(formFields, fieldFactory.NewGeneralErrorMessage(err.Error(), form))
	}
	formId := &HiddenInput{Name: FormIDName, Value: form.FormID}
	submitButton := fieldFactory.NewSubmitButton(form.GetSubmitButtonText(), form)
	return append(formFields, formId, submitButton)
}

func (self *VerticalFormLayout) BeginStruct(strct *model.MetaData, form *Form, formFields Views) Views {
	return formFields
}

func (self *VerticalFormLayout) StructField(field *model.MetaData, validationErr error, form *Form, formFields Views) Views {
	if field.Kind != model.ValueKind || form.IsFieldExcluded(field) {
		return formFields
	}
	fieldFactory := form.GetFieldFactory()
	if form.IsFieldHidden(field) {
		return append(formFields, fieldFactory.NewHiddenInput(field, form))
	}
	views := make(Views, 0, 2)
	input := fieldFactory.NewInput(field, form)
	if self.fieldNeedsLabel(field) {
		label := fieldFactory.NewLabel(input, field, form)
		views = append(views, label)
	}
	if validationErr != nil {
		views = append(views, fieldFactory.NewFieldErrorMessage(validationErr.Error(), field, form))
	}
	views = append(views, input)
	return append(formFields, P(views))
}

func (self *VerticalFormLayout) EndStruct(strct *model.MetaData, validationErr error, form *Form, formFields Views) Views {
	return formFields
}

func (self *VerticalFormLayout) BeginArray(array *model.MetaData, form *Form, formFields Views) Views {
	return formFields
}

func (self *VerticalFormLayout) ArrayField(field *model.MetaData, validationErr error, form *Form, formFields Views) Views {

	return self.StructField(field, validationErr, form, formFields) // todo replace

	return formFields
}

func (self *VerticalFormLayout) EndArray(array *model.MetaData, validationErr error, form *Form, formFields Views) Views {
	return formFields
}

func (self *VerticalFormLayout) BeginSlice(slice *model.MetaData, form *Form, formFields Views) Views {
	return formFields
}

func (self *VerticalFormLayout) SliceField(field *model.MetaData, validationErr error, form *Form, formFields Views) Views {

	return self.StructField(field, validationErr, form, formFields) // todo replace

	return formFields
}

func (self *VerticalFormLayout) EndSlice(slice *model.MetaData, validationErr error, form *Form, formFields Views) Views {
	return formFields
}

func (self *VerticalFormLayout) fieldNeedsLabel(field *model.MetaData) bool {
	switch field.Value.Addr().Interface().(type) {
	case *model.Bool:
		return false
	}
	return true
}
