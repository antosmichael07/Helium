package main

type config struct {
	color_config color_config
}

func init_config() config {
	return config{
		color_config: init_color_config(),
	}
}
