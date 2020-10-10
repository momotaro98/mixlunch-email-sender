package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/momotaro98/mixlunch-email-notification/pb"
)

func CollectRecipients(ctx context.Context, targetDate string) (recipients []*Recipient, err error) {
	dt, err := ParseDatetimeString(targetDate)
	if err != nil {
		return nil, err
	}
	// [Business Logic] [2020年7月4日時点] 翌日のマッチング結果を返す
	lunchDT := dt.AddDate(0, 0, 1)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	var partyServiceAPIAddr = os.Getenv("PARTYSERVICE_API_ADDRESS")
	var partyServiceAPIPort = os.Getenv("PARTYSERVICE_API_PORT")
	conn, err := grpc.Dial(partyServiceAPIAddr+":"+partyServiceAPIPort, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewMixLunchClient(conn)

	g, ctx := errgroup.WithContext(ctx)
	partyChan := make(chan *pb.Party)

	g.Go(func() error {
		if err := receiveParties(ctx,
			client,
			&pb.TargetDate{
				Date: fmt.Sprintf(lunchDT.Format("2006-01-02")),
			},
			partyChan,
		); err != nil {
			return err
		}
		return nil
	})

	for party := range partyChan {
		mappedRecipients, _ := mapPartyToRecipients(party)
		for _, r := range mappedRecipients {
			recipients = append(recipients, r)
		}
	}
	// Check whether any of the goroutines failed. Since g is accumulating the
	// errors, we don't need to send them (or check for them) in the individual
	// results sent on the channel.
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return recipients, nil
}

// receiveParties is called by CollectRecipients to get Parties from service-api via gRPC
func receiveParties(ctx context.Context, client pb.MixLunchClient, targetDate *pb.TargetDate, partyChan chan<- *pb.Party) error {
	defer close(partyChan)
	stream, err := client.GetParties(ctx, targetDate)
	if err != nil {
		log.Printf("%v.GetParties(_) = _, %v", client, err)
		return err
	}
	for {
		party, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("%v.GetParties(_) = _, %v", client, err)
			return err
		}
		partyChan <- party
	}
	return nil
}

// mapPartyToRecipients is called by CollectRecipients to translate from "one party protoc model" to Recipient models
func mapPartyToRecipients(protocParty *pb.Party) (membersInAParty []*Recipient, err error) {
	var preMapMembers []*Recipient
	// Generate members from a protoc party
	for _, preMapMember := range protocParty.Members {
		var r Recipient
		r.UserId = preMapMember.UserId
		r.UserName = preMapMember.UserName
		r.Email = preMapMember.Email
		preMapMembers = append(preMapMembers, &r)
	}
	// Generate a party from a protoc party
	beginDT, endDT, err := ParseBeginEndTime(protocParty.StartFrom, protocParty.EndTo)
	if err != nil {
		return nil, err
	}
	var preMapParty Party
	preMapParty.Date = fmt.Sprintf("%04d-%02d-%02d", beginDT.Year(), int(beginDT.Month()), beginDT.Day())
	preMapParty.TimeFrom = fmt.Sprintf("%02d:%02d", beginDT.Hour(), beginDT.Minute())
	preMapParty.TimeTo = fmt.Sprintf("%02d:%02d", endDT.Hour(), endDT.Minute())
	preMapParty.ChatRoomId = protocParty.ChatRoomId
	for _, pmm := range preMapMembers {
		preMapParty.MemberUserNames = append(preMapParty.MemberUserNames, pmm.UserName)
	}

	// Populate to return object
	for _, preMapUser := range preMapMembers {
		var r Recipient
		r.UserId = preMapUser.UserId
		r.Email = preMapUser.Email
		r.UserName = preMapUser.UserName
		r.Party = preMapParty
		// Append to object value
		membersInAParty = append(membersInAParty, &r)
	}
	return membersInAParty, nil
}
