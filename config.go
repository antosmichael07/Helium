package main

type config struct {
	color_config color_config
}

func (he *heightmap_editor) init_config() {
	he.config = config{}

	he.config.init_color_config()
}
