package config

import "path"

var Runtime = &RuntimeConfig{
	&File{
		Location: path.Join(GetConfigRoot(), "keeshrc"),
	},
}
