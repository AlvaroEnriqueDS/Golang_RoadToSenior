package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func InitSession() (*session.Session, error) {
	a := &aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(
			"ASIAYCKKXG5MGCOZIFBW",
			"bGTDFoDB/hBdxKXg8h+hRE8BMNl6EcGdpqnnjkS+",
			"FwoGZXIvYXdzEJX//////////wEaDGN9gwwFxaOJqINniSKaAoCIRWX4MMmEuYKPluWUraQHtcoxLWBbSvvgAtZhBOJ9+PyI3PVfLRglcGH66FL8aOabG0FfRSTYYbZldJpqZWe9zWq8ps7KbXFPsKNhdPbP9hs77acsNr0YG+S+TY8PjdLnwIUmaCMzLP0NGK0pluiOZ7UJl7f844TabFeruEImetsBsdb3k8MeK61Mjar/zg5vN2UCCxhvyjwQ7Ssj4/TMPfWEbRciqgF2tjinza4xtREG8jr5spqzYX4twxA9mAOWXcDAys042UVgyYU/2G4bSL1Os27OPhKbKfk/vVqh/f1FBYFfAfJ6iz18KO3QGfekn/U+NqRAefsis9qnJ9acjK7ns6NFmExZOAArhfF1/qpuFB6BgGL5Zij+v/z9BTIpkZ1opJkerTZtT7uawMHGmRai1Fc3OhEj1k0WK6c17B6/EgoPuUKyvrk=",
		)}

	sess, err := session.NewSession(a)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func main()  {
	sess, err := InitSession()
	if err != nil {
		panic("NO SE PUDO CONECATAR, LO SIENTO")
	}


	file, err := os.Open("./files/gopher.png")
	if err != nil {
			  panic(err)
	  }


	uploader := s3manager.NewUploader(sess)

	output, err := uploader.Upload(&s3manager.UploadInput{
		//ACL:                       aws.String("public-read"),
		Body:                      file,
		Bucket:                    aws.String("elembio.storage.test"),
		Key:                       aws.String(file.Name()),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}