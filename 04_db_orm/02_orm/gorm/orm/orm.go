package orm
//
//import (
//        "database/sql"
//        "gorm.io/gorm"
//        "gorm.io/driver/postgres"
//        "time"
//)
//
//func init() {
//        db, err := gorm.Open(postgres.New(postgres.Config{
//                //DriverName:           "",
//                DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
//                //PreferSimpleProtocol: false,
//                //Conn:                 nil,
//        }),  &gorm.Config{})
//
//
//}
//
//type User struct {
//        ID           uint `json:"id",gorm:"primaryKey"`
//        Name         string
//        Email        *string
//        Age          uint8
//        Birthday     *time.Time
//        MemberNumber sql.NullString
//        ActivedAt    sql.NullTime
//        CreatedAt    time.Time
//        UpdatedAt    time.Time
//}
