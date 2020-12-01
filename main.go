package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var csvFileName string

// User alkal user entity
type User struct {
	FirstName string
	LastName  string
	UserName  string
	Password  string
	Gender    string
	Role      string
	JobID     int
}

func init() {
	flag.StringVar(&csvFileName, "f", "", "\tUSE: '-f filename.csv' where filename is the absolute path of the csv file that contains your data for conversion [MANDATORY]")
}

func main() {
	flag.Parse()
	fmt.Println(csvFileName)
	file, err := os.Open(csvFileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(-3)
	}

	reader := csv.NewReader(file)
	var userName strings.Builder

	for lineCount := 0; ; lineCount++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if lineCount > 0 {
			name := strings.Split(record[1], " ")
			firstName := name[0]
			lastName := strings.Join(name[1:], " ")
			jobID := GetJobID(record[3])
			userName.WriteString(strings.ToLower(strings.ReplaceAll(record[1], " ", "")))
			userName.WriteString(fmt.Sprintf("%d%d", jobID, lineCount))
			password := GenPassword(userName.String())
			fmt.Printf("(%d,'%s','%s','%s','%s','%s',%d),", lineCount, firstName, lastName, userName.String(), password, "L", jobID)
			userName.Reset()
		} else {
			fmt.Printf("INSERT INTO `alkal_users` (id,first_name,last_name,username,password,gender,jobId) VALUES ")
		}
	}
}

// GetJobID get the job id from predefined job set
func GetJobID(job string) int {
	switch job {
	case "Pengemudi Alat Berat":
		return 2
	case "Pengemudi Kendaraan Operasional Lapangan":
		return 3
	case "Petugas Mekanikal Elektrikal":
		return 4
	case "Petugas Pemeliharaan Jalan dan Jembatan":
		return 5
	}
	return 0
}

//GenPassword generate password
func GenPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
