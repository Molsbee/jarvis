package config

import "fmt"

type Environment interface {
	Name() string
	ElasticSearchMainURL() string
}

type environment struct {
	name                 string
	elasticSearchMainURL string
}

func (e environment) Name() string {
	return e.name
}

func (e environment) ElasticSearchMainURL() string {
	return e.elasticSearchMainURL
}

var (
	DEV = environment{
		name:                 "dev",
		elasticSearchMainURL: "http://search-lb1.t3dev.dom:1337/main/_search?pretty=true",
	}
	PPE = environment{
		name:                 "ppe",
		elasticSearchMainURL: "http://search-lb2.t3ppe.dom:1337/main/_search?pretty=true",
	}
	PROD = environment{
		name:                 "prod",
		elasticSearchMainURL: "http://search-uc1.t3n.dom:1337/main/_search?pretty=true",
	}
)

func GetEnvironment(env string) (e Environment, err error) {
	switch env {
	case DEV.name:
		e = DEV
		return
	case PPE.name:
		e = PPE
		return
	case PROD.name:
		e = PROD
		return
	}

	err = fmt.Errorf("%s is an unsupported environment - supported values(%s, %s, %s)", env, DEV, PPE, PROD)
	return
}
