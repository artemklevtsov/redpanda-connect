package stat_table

import (
	"github.com/Jeffail/shutdown"
	"github.com/redpanda-data/benthos/v4/public/service"
)

func inputFromParsed(conf *service.ParsedConfig, mgr *service.Resources) (service.BatchInput, error) {
	input := &benthosInput{
		logger:  mgr.Logger(),
		shutSig: shutdown.NewSignaller(),
		query:   &apiQuery{},
	}

	var err error

	input.query.IDs, err = conf.FieldIntList("ids")
	if err != nil {
		return nil, err
	}

	input.query.Metrics, err = conf.FieldStringList("metrics")
	if err != nil {
		return nil, err
	}
	// if len(input.query.Metrics) == 0 {
	// 	return nil, errors.New("metrics not defined")
	// }

	if conf.Contains("token") {
		input.token, err = conf.FieldString("token")
		if err != nil {
			return nil, err
		}
	}

	if conf.Contains("dimensions") {
		input.query.Dimensions, err = conf.FieldStringList("dimensions")
		if err != nil {
			return nil, err
		}
	}

	if conf.Contains("filters") {
		input.query.Filters, err = conf.FieldString("filters")
		if err != nil {
			return nil, err
		}
	}

	if conf.Contains("sort") {
		input.query.Sort, err = conf.FieldStringList("sort")
		if err != nil {
			return nil, err
		}
	}

	if conf.Contains("date1") {
		input.query.Date1, err = conf.FieldString("date1")
		if err != nil {
			return nil, err
		}
	}

	if conf.Contains("date2") {
		input.query.Date2, err = conf.FieldString("date2")
		if err != nil {
			return nil, err
		}
	}

	if conf.Contains("lang") {
		input.query.Lang, err = conf.FieldString("lang")
		if err != nil {
			return nil, err
		}
	}

	if conf.Contains("preset") {
		input.query.Preset, err = conf.FieldString("preset")
		if err != nil {
			return nil, err
		}
	}

	if conf.Contains("timezone") {
		input.query.Timezone, err = conf.FieldString("timezone")
		if err != nil {
			return nil, err
		}
	}

	if conf.Contains("direct_client_logins") {
		input.query.DirectLogins, err = conf.FieldStringList("direct_client_logins")
		if err != nil {
			return nil, err
		}
	}

	if conf.Contains("format_keys") {
		input.formatKeys, err = conf.FieldBool("format_keys")
		if err != nil {
			return nil, err
		}
	}

	return input, nil
}
