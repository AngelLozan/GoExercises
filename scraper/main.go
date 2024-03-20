package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

type Malware struct {
	link  string
	title string
}

func sendEmail(items []Malware) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	emailAppPassword := os.Getenv("APP_PASS")
	yourMail := os.Getenv("APP_PASS")
	recipient := os.Getenv("APP_PASS")
	hostAddress := "smtp.mail.yahoo.com"
	hostPort := "465"
	mailSubject := "Malicious packages found on Snap"
	mailBody := fmt.Sprintf("Please review the following packages: %v \n", items)
	fullServerAddress := hostAddress + ":" + hostPort

	headerMap := make(map[string]string)
	headerMap["From"] = yourMail
	headerMap["To"] = recipient
	headerMap["Subject"] = mailSubject
	mailMessage := ""
	for k, v := range headerMap {
		mailMessage += fmt.Sprintf("%s: %s\\r", k, v)
	}
	mailMessage += "\\r" + mailBody

	authenticate := smtp.PlainAuth("", yourMail, emailAppPassword, hostAddress)
	tlsConfigurations := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         hostAddress,
	}

	conn, err := tls.Dial("tcp", fullServerAddress, tlsConfigurations)

	if err != nil {
		log.Panic(err)
	}

	newClient, err := smtp.NewClient(conn, hostAddress)

	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = newClient.Auth(authenticate); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = newClient.Mail(yourMail); err != nil {
		log.Panic(err)
	}

	if err = newClient.Rcpt(headerMap["To"]); err != nil {
		log.Panic(err)
	}

	// Data
	writer, err := newClient.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = writer.Write([]byte(mailMessage))

	if err != nil {
		log.Panic(err)
	}

	err = writer.Close()

	if err != nil {
		log.Panic(err)
	}

	err = newClient.Quit()

	if err != nil {
		fmt.Println("THERE WAS AN ERROR")
	}

	fmt.Println("Successful, the mail was sent!")

}

func main() {
	c := colly.NewCollector()

	url := "https://snapcraft.io/search?q=exodus"

	element := ".p-media-object"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Got a response from %v\n\n", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("An error occurred!:", e)
	})

	var items []Malware

	c.OnHTML(element, func(e *colly.HTMLElement) {

		maliciousItem := Malware{}

		link := e.Attr("href")
		title := e.Attr("title")

		cleanLink := strings.TrimSpace(link)
		cleanTitle := strings.TrimSpace(title)

		maliciousItem.link = fmt.Sprintf("https://snapcraft.io%s", cleanLink)
		maliciousItem.title = cleanTitle

		if strings.Contains(strings.ToLower(maliciousItem.title), "wallet") {
			items = append(items, maliciousItem)
		}

	})

	c.OnScraped(func(r *colly.Response) {
		if len(items) > 0 {
			fmt.Println("Found some malicious items:", items)
			sendEmail(items)
		} else {
			fmt.Println("Nothing found today")
		}
	})

	err := c.Visit(url)

	if err != nil {
		log.Fatal(err)
	}
}
