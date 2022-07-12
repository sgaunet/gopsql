package main

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var log = logrus.New()

// initTrace initialize log instance with the level in parameter
func initTrace(debugLevel string) {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	// log.SetFormatter(&log.TextFormatter{
	// 	DisableColors: true,
	// 	FullTimestamp: true,
	// })

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	switch debugLevel {
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	var err error
	var dbhost string
	var dbport int
	var dbuser string
	var rqt string
	var dbpassword string

	flag.StringVar(&dbhost, "h", "localhost", "host of database server")
	flag.StringVar(&dbuser, "U", "postgres", "User ")
	flag.StringVar(&dbpassword, "P", "postgres", "Password ")
	flag.IntVar(&dbport, "p", 5432, "Port of postgreql server")
	flag.StringVar(&rqt, "c", "", "Request to execute")
	flag.Parse()
	dbname := "mydb"

	initTrace("debug")

	if rqt == "" {
		log.Errorln("no request to execute")
		os.Exit(1)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpassword, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Debugln(psqlInfo)
		log.Fatalln("Failed to connect to database")
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln("Failed to connect to database")
		os.Exit(1)
	}
	defer db.Close()

	var myMap = make(map[string]interface{})

	rows, err := db.Query(rqt)
	defer rows.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
	colNames, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	cols := make([]interface{}, len(colNames))
	colPtrs := make([]interface{}, len(colNames))

	// Print column name
	for i := 0; i < len(colNames); i++ {
		colPtrs[i] = &cols[i]
		fmt.Printf("%s ", colNames[i])
	}
	fmt.Println()
	// end print columns

	for rows.Next() {
		err = rows.Scan(colPtrs...)
		if err != nil {
			log.Fatal(err)
		}
		for i, col := range cols {
			myMap[colNames[i]] = col
			fmt.Printf("%v ", col)
		}
		fmt.Println()
		// Do something with the map
		// for key, val := range myMap {
		// fmt.Printf("Key: %15s    Value Type: %10s   Value: %v\n", key, reflect.TypeOf(val), val)
		// }
	}

	fmt.Println("***************************")
	b := new(bytes.Buffer)
	enc := yaml.NewEncoder(b)
	enc.Encode(myMap)
	enc.Close()
	b.WriteTo(os.Stdout)
}
