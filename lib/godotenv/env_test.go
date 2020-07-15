package godotenv

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")
	log.Println(s3Bucket, secretKey)
	// now do something with s3 or whatever
}
