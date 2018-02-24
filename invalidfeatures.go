package main

import (
   "fmt"
   "database/sql"
   "time"
   _ "github.com/lib/pq"
)

const (
   host     = "localhost"
   port     = 5432
   user     = "pdok_owner"
   password = "pdok_owner"
   dbname   = "pdok"
)

const (
   sqlSelect = "SELECT feature_type, feature_id FROM extract_wijkenbuurten2013.delta_data"
   sqlSequenceScan  = "set enable_seqscan = off"
)
func main() {


   var f1 = feature{"Iets", "id", time.Now(), "", ""}
   var f2 = feature{"Iets", "id", time.Now(), "", ""}
   f1.equals(f2)

   db, error := setupDatabase()
   checkError(error)
   defer db.Close()


   db.Exec(sqlSequenceScan)
   rows, error := db.Query(sqlSelect)
   checkError(error)

   for rows.Next() {
      var collection string
      var featureId string

      error := rows.Scan(&collection, &featureId)
      checkError(error)
      fmt.Printf("Value: %s => %s\n", collection, featureId)
   }

}


func setupDatabase() (sql.DB,  error) {
   dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
                  host, port, user, password, dbname)
   db, err := sql.Open("postgres", dbinfo)
   return *db, err
}



func checkError(err error) {
   if err != nil {
      panic(fmt.Sprintf("%s", err))
   }
}

