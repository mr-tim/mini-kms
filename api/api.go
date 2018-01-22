package api

import "../kms"

type Api interface {
	Run()
}

func New(k kms.Kms) Api {
	return GinApi{k}
}
