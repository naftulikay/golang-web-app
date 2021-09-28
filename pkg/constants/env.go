package constants

const (
	DevEnvironment        = "dev"
	StagingEnvironment    = "stg"
	ProductionEnvironment = "prod"
)

func EnvironmentList() []string {
	return []string{DevEnvironment, StagingEnvironment, ProductionEnvironment}
}

func Environments() map[string]bool {
	return map[string]bool{
		DevEnvironment:        true,
		StagingEnvironment:    true,
		ProductionEnvironment: true,
	}
}
