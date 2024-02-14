# Email Microsservice
## Uber BackEnd Old Challange

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![AWS SES](https://img.shields.io/badge/Amazon_AWS-232F3E?style=for-the-badge&logo=amazon-aws&logoColor=white)

This project is an API using **Go and AWS Simple Email Service**

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)

## Installation
1. Clone the repository
```bash
git clone https://github.com/rcarvalho-pb/emailsender-aws-ses-golang
```
2. Install dependencies with:
```bash
go mod tidy
```
3. Update `.env` file putting yours AWS Credentials and port for the API
```env
PORT=
AWS_REGION=
AWS_ACESSKEYID=
AWS_SECRETKEY=
SEND_EMAIL=
CHARSET=
```
## Usage
1. Start the application with:
```bash
go run cmd/main.go
```
2. The API will be accessible at http:localhost:`Port`

## API Endpoints
The API provides the following endpoints:

**Get Email**
```markdown
POST /api/emails - Send an e-mail from your sender to the destination
```

**Body**
```json
{
    "to": "example@email.com",
    "subject": "Teste",
    "body": "Test from AWS SES"
}
```