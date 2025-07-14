package vacancy

import "time"

type VacancyCreateForm struct {
	Role     string
	Company  string
	Type     string
	Salary   string
	Location string
	Email    string
}

type Vacancy struct {
	ID        int `db:"id"`
	Email     string `db:"email"`
	Role      string `db:"role"`
	Company   string `db:"company"`
	Type      string `db:"type"`
	Salary   string `db:"salary"`
	Location  string `db:"location"`
	CreatedAt time.Time `db:"createdat"`
}