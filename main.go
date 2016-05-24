package main

import (
    "net/http"
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
    "github.com/qor/qor"
    "github.com/qor/admin"
)

// Create a GORM-backend model
type User struct {
  gorm.Model
  Name string
}

// Create another GORM-backend model
type Career struct {
  gorm.Model
  Title        string
  Description string
}

func main() {
  DB, _ := gorm.Open("sqlite3", "demo.db")
  DB.AutoMigrate(&User{}, &Career{})

  // Initalize
  Admin := admin.New(&qor.Config{DB: DB})
  Admin.SetSiteName("MergEye")

  // Create resources from GORM-backend model
  Admin.AddResource(&User{})
  Admin.AddResource(&Career{})

  // Register route
  mux := http.NewServeMux()
  // amount to /admin, so visit `/admin` to view the admin interface
  Admin.MountTo("/admin", mux)

  fmt.Println("Listening on: 9000")
  http.ListenAndServe(":9000", mux)
}
