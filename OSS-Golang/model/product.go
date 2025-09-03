import "gorm"

package model

type Product struct {
    ID          uint    ` + "`gorm:\"primaryKey\" json:\"id\"`" + `
    Name        string  ` + "`json:\"name\" binding:\"required\"`" + `
    Description string  ` + "`json:\"description\"`" + `
    Price       float64 ` + "`json:\"price\" binding:\"required\"`" + `
}