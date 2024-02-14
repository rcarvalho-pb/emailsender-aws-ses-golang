package ses

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type AmazonSES struct {
	Sender    string
	Charset   string
	Region    string
	AKAID     string
	SecretKey string
}

func NewAmazonSES() *AmazonSES {
	return &AmazonSES{
		Sender:    os.Getenv("SEND_EMAIL"),
		Charset:   os.Getenv("CHARSET"),
		Region:    os.Getenv("AWS_REGION"),
		AKAID:     os.Getenv("AWS_ACESSKEYID"),
		SecretKey: os.Getenv("AWS_SECRETKEY"),
	}
}

func (amazonSES *AmazonSES) SendEmailGateway(toEmail, subject, body string) {
	ases := NewAmazonSES()
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(ases.Region),
		Credentials: credentials.NewStaticCredentials(ases.AKAID, ases.SecretKey, ""),
	})
	if err != nil {
		panic(err)
	}

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(toEmail),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(ases.Charset),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(ases.Charset),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(ases.Sender),
	}

	result, err := svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Println("Email Sent to address: " + toEmail)
	fmt.Println(result)

}
