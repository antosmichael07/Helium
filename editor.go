package main

type heightmap_editor struct {
	tool_window tool_window
	config      config
}

func init_heightmap_editor() heightmap_editor {
	return heightmap_editor{
		tool_window: init_tool_window(),
		config:      init_config(),
	}
}
