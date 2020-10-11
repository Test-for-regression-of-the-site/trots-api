package configuration

import (
	"github.com/Test-for-regression-of-the-site/trots-api/lighthouse"
)

type Configuration struct {
	Address    string
	Lighthouse lighthouse.Config
}
