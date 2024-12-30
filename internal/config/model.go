package config

type AppConfig struct {
	Global         *AppGlobal `json:"global,omitempty"`
	Infrastructure *AppInfra  `json:"infrastructure"`
}

type AppGlobal struct {
	LogLevel string `json:"logLevel,omitempty"`
}

type AppInfra struct {
	Notion   *NotionConfig   `json:"notion"`
	Calendar *CalendarConfig `json:"calendar"`
}

type NotionConfig struct {
	IntegrationKey string                `json:"integrationKey"`
	Database       *NotionDatabaseConfig `json:"database"`
}

type NotionDatabaseConfig struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Filter string `json:"filter"`
}

type CalendarConfig struct {
	Provider string               `json:"provider"`
	Google   GoogleCalendarConfig `json:"google"`
}

type GoogleCalendarConfig struct {
}
