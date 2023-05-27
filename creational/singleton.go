package creational

type Singletone struct {
	field string
}

var instance *Singletone

func GetSingletoneInstance() *Singletone {
	if instance == nil {
		instance = &Singletone{field: "I am Singleton"}
	}

	return instance
}
