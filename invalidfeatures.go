package main

import (
   "fmt"
   "database/sql"
   "time"
   _ "github.com/lib/pq"
)

const (
   host     = "10.91.10.41"
   port     = 49187
   user     = "pdok_featured_owner"
   password = "pPk9en7CI5BTjP1jdSDmJ9PjNEGUSFRa"
   dbname   = "pdok_featured"
	//host     = "localhost"
	//port     = 5432
	//user     = "pdok_owner"
	//password = "pdok_owner"
	//dbname   = "pdok"
)

const (
   // sqlSelect = "SELECT feature_type, feature_id FROM extract_wijkenbuurten2013.delta_data"
   sqlSelect = `SELECT collection, feature_id, validity, action, attributes
				FROM featured_bgtv3.feature_stream
				ORDER BY collection, feature_id`
   sqlSequenceScan  = "set enable_seqscan = off"
)
func main() {

   db, err := setupDatabase()
   checkError(err)
   defer db.Close()


   db.Exec(sqlSequenceScan)
   rows, err := db.Query(sqlSelect)
   checkError(err)

   var firstTime = true
   var previousFeature feature
   var features []feature

   var collection string
   var featureId string
   var action string
   var attributes string
   var validity time.Time


   for rows.Next() {
      err := rows.Scan(&collection, &featureId, &validity, &action, &attributes)
      checkError(err)

      fmt.Printf("Value: %s => %s\n", collection, featureId)
      feature := NewFeature(collection, featureId, validity, action, attributes)

      if firstTime {
         previousFeature = *feature
         firstTime = false
      }

      if previousFeature.equals(*feature) {
         features = append(features, *feature)
      } else {
         processFeatures(features)
         //  Rest the array
         features = nil    //  or features = features[:0]
         features = append(features, *feature)
      }
      previousFeature = *feature

   }

}


func processFeatures(features []feature) {
   fmt.Printf("There are %d elements", len(features))

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

