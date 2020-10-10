package main

import (
	"context"
	"flag"
	"log"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/momotaro98/mixlunch-email-notification/emailsender"
)

func init() {
	emailsender.InitEmailSender()
	InitBodyTemplate()
}

func sendEmailToRecipients(ctx context.Context, recipients []*Recipient) error {
	emailSender := emailsender.NewSender()
	g, ctx := errgroup.WithContext(ctx)
	for _, r := range recipients {
		r := r // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Assemble email body
			body, err := r.AssembleEmailBody()
			if err != nil {
				return err
			}
			// Send to the recipient
			if err := r.Send(ctx, emailSender, body); err != nil {
				return err
			}
			log.Println(body)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func main() {
	var (
		targetDate = flag.String("target-date", "Year-Month-Day", "Target date to send email to users")
	)
	flag.Parse()
	log.Printf("Process start with target-date %s", *targetDate)

	// Step: Collect target recipients
	log.Println("Collecting target recipients process starts.")
	collectCtx, collectCancel := context.WithTimeout(context.Background(), time.Second*30)
	defer collectCancel()
	recipients, err := CollectRecipients(collectCtx, *targetDate)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Result of CollectRecipients: %v", recipients)
	for i, r := range recipients {
		log.Printf(" %d th recipient, %v", i, r)
	}

	// Step: Send Email to recipients
	log.Println("Sending emails process starts.")
	sendCtx, sendCancel := context.WithTimeout(context.Background(), time.Second*30)
	defer sendCancel()
	if err := sendEmailToRecipients(sendCtx, recipients); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Process done.")
}
