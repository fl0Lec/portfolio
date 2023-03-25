package data

type Data struct {
	Name   string            `yaml:"name"`
	Skills map[string]Skills `yaml:"skills"`
	Works  map[string]Work   `yaml:"work"`
}
type Skills map[string][]string

type Work struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	From        string `yaml:"from"`
	To          string `yaml:"to"`
	Where       string `yaml:"where"`
}
