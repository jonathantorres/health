package db

type BloodReading struct {
	Id        int64
	UserId    int64
	Systolic  int32
	Diastolic int32
	Pulse     int32
	Date      string
}

type BloodSeverity struct {
	Text  string
	Class string
}

func (b *BloodReading) Severity() *BloodSeverity {
	text := "N/A"
	class := "normal"

	if b.Systolic <= 120 && b.Diastolic <= 80 {
		text = "Normal"
		class = "primary"
	} else if (b.Systolic > 120 && b.Systolic <= 139) || (b.Diastolic > 80 && b.Diastolic <= 89) {
		text = "Pre Hypertension"
		class = "warning"
	} else if (b.Systolic >= 140 && b.Systolic <= 159) || (b.Diastolic >= 90 && b.Diastolic <= 99) {
		text = "Stage 1 Hypertension"
		class = "danger"
	} else if b.Systolic >= 160 && b.Diastolic >= 100 {
		text = "Stage 2 Hypertension"
		class = "danger"
	}

	return &BloodSeverity{
		Text:  text,
		Class: class,
	}
}

func (b *BloodReading) SqlDate() string {
	return b.Date[:10]
}
