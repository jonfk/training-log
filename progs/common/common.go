package common

type Exercise struct {
	Name     string `name`
	Weight   string `weight,omitempty`
	Sets     string `sets`
	Reps     string `reps`
	Exertion string `exertion`
}

type TrainingLog struct {
	Date       string     `date`
	Time       string     `time,omitempty`
	Length     string     `length,omitempty`
	Bodyweight string     `bodyweight,omitempty`
	Event      string     `event,omitempty`
	Wilks      string     `wilks,omitempty`
	Total      string     `total,omitempty`
	Workout    []Exercise `workout`
	Notes      []string   `notes,omitempty`
}
