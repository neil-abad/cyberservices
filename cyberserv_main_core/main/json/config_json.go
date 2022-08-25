package json

//go:generate go run cyberservices.com/core/common/errors/errorgen

import (
	"io"

	"cyberservices.com/core"
	"cyberservices.com/core/common"
	"cyberservices.com/core/common/cmdarg"
	"cyberservices.com/core/infra/conf/serial"
	"cyberservices.com/core/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				r, err := confloader.LoadExtConfig(v)
				if err != nil {
					return nil, newError("failed to execute v2ctl to convert config file.").Base(err).AtWarning()
				}
				return core.LoadConfig("protobuf", "", r)
			case io.Reader:
				return serial.LoadJSONConfig(v)
			default:
				return nil, newError("unknow type")
			}
		},
	}))
}
