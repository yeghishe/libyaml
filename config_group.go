package libyaml

type ConfigGroup struct {
	Name        string        `yaml:"name" json:"name"`
	Title       string        `yaml:"title" json:"title"`
	Description string        `yaml:"description" json:"description"`
	TestProc    *TestProc     `yaml:"test_proc" json:"test_proc"`
	Items       []*ConfigItem `yaml:"items" json:"items"`
	When        string        `yaml:"when" json:"when"`
}
