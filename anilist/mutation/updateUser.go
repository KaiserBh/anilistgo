package anilist

import (
	anilist "AniListGoWrapper/anilist/query"
	"encoding/json"
	"fmt"
)

//UpdateUserAbout User's about/bio text.
func UpdateUserAbout(about string, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(about:"%v"){
				about,
			},
		}`, about),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserTitleLanguage User's title language/ The language the user wants to see media titles in such as ROMAJI, ENGLISH, NATIVE, ROMAJI_STYLISED, ENGLISH_STYLISED, NATIVE_STYLISED
func UpdateUserTitleLanguage(titleLanguage string, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(titleLanguage:%v){
				options {
      				titleLanguage,
    			},
			},
		}`, titleLanguage),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserAdultContent If the user should see media marked as adult-only boolean
func UpdateUserAdultContent(displayAdultContent bool, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(displayAdultContent:%v){
				options {
      				displayAdultContent,
    			},
			},
		}`, displayAdultContent),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserAiringNotification If the user should get notifications when a show they are watching aires
func UpdateUserAiringNotification(airingNotifications bool, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(airingNotifications:%v){
				options {
					  displayAdultContent,
					  airingNotifications,
    			},
			},
		}`, airingNotifications),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserScoreFormat The user's list scoring system/Media list scoring type such as POINT_100, POINT_10_DECIMAL, POINT_10, POINT_5, POINT_3
func UpdateUserScoreFormat(scoreFormat string, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(scoreFormat:%v){
				mediaListOptions {
					scoreFormat,
    			},
			},
		}`, scoreFormat),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserRowOrder The user's default list order such as Score, title, last updated, last added
func UpdateUserRowOrder(rowOrder string, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(rowOrder:"%v"){
				mediaListOptions {
					rowOrder,
    			},
			},
		}`, rowOrder),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserProfileColor The user's Profile highlight color: blue, purple, green, orange, red, pink
func UpdateUserProfileColor(profileColor string, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(profileColor:"%v"){
				options {
					profileColor,
    			},
			},
		}`, profileColor),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

// UpdateUserNotificationOption The user's Notification options
/*func

ACTIVITY_MESSAGE
A user has sent you message

ACTIVITY_REPLY
A user has replied to your activity

FOLLOWING
A user has followed you

ACTIVITY_MENTION
A user has mentioned you in their activity

THREAD_COMMENT_MENTION
A user has mentioned you in a forum comment

THREAD_SUBSCRIBED
A user has commented in one of your subscribed forum threads

THREAD_COMMENT_REPLY
A user has replied to your forum comment

AIRING
An anime you are currently watching has aired

ACTIVITY_LIKE
A user has liked your activity

ACTIVITY_REPLY_LIKE
A user has liked your activity reply

THREAD_LIKE
A user has liked your forum thread

THREAD_COMMENT_LIKE
A user has liked your forum comment

ACTIVITY_REPLY_SUBSCRIBED
A user has replied to activity you have also replied to

RELATED_MEDIA_ADDITION
A new anime or manga has been added to the site where its related media is on the user's list
*/
func UpdateUserNotificationOption(notificationType string, enabled bool, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(notificationOptions:{type:%v,enabled:%v}) {
				options {
					notificationOptions {
						type,
						enabled,
					},
				},
			},
		}`, notificationType, enabled),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

// UpdateUserAnimeListOptionSectionOrder A user's list options for anime or manga lists
func UpdateUserAnimeListOptionSectionOrder(sectionOrder string, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(animeListOptions:{sectionOrder:"%v"}) {
				mediaListOptions {
					animeList {
						sectionOrder
						splitCompletedSectionByFormat
						customLists
						advancedScoring
						advancedScoringEnabled
					}
				},
			},
		}`, sectionOrder),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}
