package requests_to_server

import (
	"OISA_2x_sistem/settings"
	"log"
)

// TODO: Implement this and delete blocking functions everywhere

// also what is this https://pkg.go.dev/net/http/httptest

// TODO: if request with retries fails cut that bookie out from main, turn it off
// https://brandur.org/fragments/go-http-retry

//package main
//
//import (
//    "fmt"
//    "io/ioutil"
//    "log"
//    "net/http"
//)
//
//func main() {
//    var (
//        err      error
//        response *http.Response
//        retries  int = 3
//    )
//    for retries > 0 {
//        response, err = http.Get("https://non-existent")
//        // response, err = http.Get("https://google.com/robots.txt")
//        if err != nil {
//            log.Println(err)
//            retries -= 1
//        } else {
//            break
//        }
//    }
//    if response != nil {
//        defer response.Body.Close()
//        data, err := ioutil.ReadAll(response.Body)
//        if err != nil {
//            log.Fatal(err)
//        }
//        fmt.Printf("data = %s\n", data)
//    }
//}

func GetSidebarSportsBlocking() map[string]interface{} {

	response := GetSidebarSports()

	for i := 0; i < settings.NumOfTries; i++ {

		if response != nil {
			break
		}

		log.Println("Merkurxtip: Stuck on GetSidebarSports()...")
		response = GetSidebarSports()
	}

	return response
}

func GetAllSubgamesBlocking() map[string]interface{} {

	response := GetAllSubgames()

	for i := 0; i < settings.NumOfTries; i++ {

		if response != nil {
			break
		}

		log.Println("Merkurxtip: Stuck on GetAllSubgames()...")
		response = GetAllSubgames()

	}

	return response
}

func GetSidebarSportGroupsBlocking(sportID string) map[string]interface{} {

	response := GetSidebarSportGroups(sportID)

	for i := 0; i < settings.NumOfTries; i++ {

		if response != nil {
			break
		}

		log.Println("Merkurxtip: Stuck on GetSidebarSportGroups(sportID=" + sportID + ")...")
		response = GetSidebarSportGroups(sportID)
	}

	return response
}

func GetSidebarSportGroupLeaguesBlocking(sportID string, groupID string) map[string]interface{} {

	response := GetSidebarSportGroupLeagues(sportID, groupID)

	for i := 0; i < settings.NumOfTries; i++ {

		if response != nil {
			break
		}

		log.Println("Merkurxtip: Stuck on GetSidebarSportGroupLeagues(sportID=" + sportID + ", groupID=" + groupID + ")...")
		response = GetSidebarSportGroupLeagues(sportID, groupID)
	}

	return response
}

func GetMatchIDsBlocking(sportID string, leagueID string) map[string]interface{} {

	response := GetMatchIDs(sportID, leagueID)

	for i := 0; i < settings.NumOfTries; i++ {

		if response != nil {
			break
		}

		log.Println("Merkurxtip: Stuck on GetMatchIDs(sportID=" + sportID + ", leagueID=" + leagueID + ")...")
		response = GetMatchIDs(sportID, leagueID)
	}

	return response
}
