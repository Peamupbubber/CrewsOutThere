package command

import (
	"crewFinder/db"
	"log"
)

// Given a timestamp, clean up the database of the requests that are more than five minutes old, then send out any deferred messages that haven't timed out
// Returns whether or not the user texting the system was orginally deferred. Mitigates the case where they said yes to a request they haven't gotten yet
func cleanupDB(phone string, timeStamp int64) bool {
	fiveMinutesAgo := timeStamp - int64(300000)

	notifyRequestersOfTimedOutRequest(fiveMinutesAgo)
	clearOldInvites(fiveMinutesAgo)
	clearOldRequests(fiveMinutesAgo)
	clearOldDeferred(fiveMinutesAgo)

	query := "SELECT * FROM Contacts WHERE timestamp < ?"
	selectResult, err := db.DB.Query(query, fiveMinutesAgo)

	if err != nil {
		log.Fatalf("Impossible select from Contacts: %s", err)
	}

	defer selectResult.Close()

	_, _, message, _ := getTopItemFromDef(phone)
	originalSenderWasRemovedFromContacts := false

	// For each timestamp that has timed out, remove them from all tables then handle any outgoing contacts
	for selectResult.Next() {
		var rPhone string
		var cPhone string
		var timedOutTimeStamp int64

		err = selectResult.Scan(&rPhone, &cPhone, &timedOutTimeStamp)

		if err != nil {
			log.Fatalf("Impossible to get row from selected results: %s", err)
		}

		if cPhone == phone {
			originalSenderWasRemovedFromContacts = true
		}

		handleOutgoingContacts(cPhone, timedOutTimeStamp)
	}

	return message != "" && originalSenderWasRemovedFromContacts
}

/* This set of functions clears out old entries from stated tables */

// Only clears invalid members
func clearOldInvites(cutoff int64) {
	query := "DELETE FROM Members WHERE timestamp < ? AND isValid = 0"
	_, err := db.DB.Exec(query, cutoff)
	if err != nil {
		log.Fatalf("Impossible delete from Members: %s", err)
	}
}

func clearOldDeferred(cutoff int64) {
	query := "DELETE FROM Deferred WHERE timestamp < ?"
	_, err := db.DB.Exec(query, cutoff)
	if err != nil {
		log.Fatalf("Impossible delete from Deferred: %s", err)
	}
}

func clearOldRequests(cutoff int64) {
	query := "DELETE FROM Requester WHERE timestamp < ?"
	_, err := db.DB.Exec(query, cutoff)
	if err != nil {
		log.Fatalf("Impossible delete from Contacts: %s", err)
	}
}

// When a user's request is timed out, notify them
func notifyRequestersOfTimedOutRequest(cutoff int64) {
	query := "SELECT * FROM Requester WHERE timestamp < ?"
	toNotify, err := db.DB.Query(query, cutoff)

	if err != nil {
		log.Fatalf("Impossible select from Requester: %s", err)
	}

	defer toNotify.Close()

	for toNotify.Next() {
		var timeStamp int64
		var rPhone string
		var message string

		err = toNotify.Scan(&timeStamp, &rPhone, &message)

		if err != nil {
			log.Fatalf("Impossible to get row from selected results: %s", err)
		}

		MessageUser(rPhone, "Your request has timed out")
	}
}
