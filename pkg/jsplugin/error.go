package jsplugin

import "errors"

var (
	ErrTypeConversion = errors.New("error converting golang type into js variable")
)
