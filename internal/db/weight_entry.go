package db

type WeightEntry struct {
	Id     int64
	UserId int64
	Weight float32
	Date   string
}

func (w *WeightEntry) SqlDate() string {
	return w.Date[:10]
}
