package lrjson

import (
	"fmt"
	"testing"
)

func benchmarkDynamicUnmarshal(b *testing.B, data string) {
	for n := 0; n < b.N; n++ {
		unmarshalled, error := DynamicUnmarshal(data)
		if error != nil {
			fmt.Println(error)
		} else {
			_ = unmarshalled["Uid"]
		}
	}
}

func BenchmarkDynamicUnmarshal(b *testing.B) {
	benchmarkDynamicUnmarshal(b, profile)
}

func BenchmarkUnmarshalGetManageAccountProfilesByEmail(b *testing.B) {
	benchmarkUnmarshalGetManageAccountProfilesByEmail(b, profile)
}

func benchmarkUnmarshalGetManageAccountProfilesByEmail(b *testing.B, data string) {
	for n := 0; n < b.N; n++ {
		unmarshalled, error := UnmarshalGetManageAccountProfilesByEmail(data)
		if error != nil {
			fmt.Println(error)
		} else {
			_ = unmarshalled.UID
		}
	}
}

var profile = `
{
  "Identities": [
  ],
  "UserName": "JohnSmith",
  "Uid": "f2005bd20****************3278343",
  "PasswordExpirationDate": "2028-02-07T19:02:55.803Z",
  "CustomFields": {
    "id": "https://graph.facebook.com/*****/picture?type=small",
    "id2": "https://graph.Profile.com/*****/picture?type=square"
  },
  "LastPasswordChangeToken": "1542daac-4df0-4270-b5b5-5d7922492d7e",
  "phoneId": "824525-add135",
  "ExternalUserLoginId": "asybkyabskdhbkad",
  "RegistrationProvider": "Email",
  "ID": "dfb7****fc947618****f0ae46ff",
  "Provider": "Email",
  "Prefix": "Mr.",
  "FirstName": "Test1",
  "MiddleName": "M.",
  "LastName": "Account2",
  "Suffix": "Junior",
  "FullName": "Test1 M. Account2",
  "NickName": "TMC",
  "ProfileName": "TMA",
  "BirthDate": "1990-02-07",
  "Gender": "F",
  "Website": "www.whatwebsite.com",
  "Email": [
    {
      "Type": "Primary",
      "Value": "NorthOfKing@Westeros0.com"
    }
  ],
  "Country": {
    "Code": "HS",
    "Name": "House Smith"
  },
  "ThumbnailImageUrl": "www.image.com",
  "ImageUrl": "www.image.com",
  "Favicon": "www.image.com",
  "ProfileUrl": "www.image.com",
  "HomeTown": "Vancouver",
  "State": "BC",
  "City": "Vancouver",
  "Industry": "IT",
  "About": "a junior engineer",
  "TimeZone": "+8",
  "LocalLanguage": "English",
  "CoverPhoto": "www.beatiufulimage.com",
  "TagLine": "#engineer",
  "Language": "English",
  "Verified": "false",
  "UpdatedTime": "2018-02-06T19:02:55.803Z",
  "Positions": [
    {
      "Position": "astronaut",
      "Summary": "An astronaut.",
      "StartDate": "2008-02-07T19:02:55.803Z",
      "EndDate": "2015-02-07T19:02:55.803Z",
      "IsCurrent": "false",
      "Location": "Houston",
      "Comapny": {
        "Name": "NASA",
        "Type": "astronaut",
        "Industry": "National secure"
      },
      "Company": {
        "Name": "NASA",
        "Type": "astronaut",
        "Industry": "National secure"
      }
    }
  ],
  "Educations": [
    {
      "School": "UCL",
      "year": "2010",
      "type": "student",
      "notes": "great",
      "activities": "sports",
      "degree": "master",
      "fieldofstudy": "astronaut",
      "StartDate": "2004-02-07T19:02:55.803Z",
      "EndDate": "2005-02-07T19:02:55.803Z"
    }
  ],
  "PhoneNumbers": [
    {
      "PhoneType": "Mobile",
      "PhoneNumber": "84684*543"
    }
  ],
  "IMAccounts": [
    {
      "AccountType": "primary",
      "AccountName": "Guessnot"
    }
  ],
  "Addresses": [
    {
      "Type": "Primary",
      "Address1": "1281 West Mall",
      "Address2": "Room 121",
      "City": "Edmonton",
      "State": "Alberta",
      "PostalCode": "546841",
      "Region": "Myself",
      "Country": "Canada"
    }
  ],
  "MainAddress": "1281 West Mall rm 121",
  "Created": "true",
  "CreatedDate": "2017-08-24T12:21:41.808Z",
  "ModifiedDate": "2017-08-24T13:04:12.1810017Z",
  "LocalCity": "Vancouver",
  "ProfileCity": "Edmonton",
  "LocalCountry": "House Smith",
  "ProfileCountry": "Canada",
  "IsProtected": false,
  "RelationshipStatus": "single",
  "Quota": "Winter is coming",
  "Quote": "North remembers",
  "InterestedIn": [
    "blah",
    "blah"
  ],
  "Interests": [
    {
      "InterestedType": "Most",
      "InterestedName": "Shoot"
    }
  ],
  "Religion": "None",
  "Political": "None",
  "Sports": [
    {
      "Id": "545",
      "Name": "Swimming"
    }
  ],
  "InspirationalPeople": [
    {
      "Id": "5846",
      "Name": "Sansa"
    }
  ],
  "HttpsImageUrl": "www.image.com",
  "FollowersCount": 0,
  "FriendsCount": 0,
  "IsGeoEnabled": "false",
  "TotalStatusesCount": 0,
  "Associations": "North Army",
  "NumRecommenders": 0,
  "Honors": "the engineer of the year",
  "Awards": [
    {
      "Id": "001",
      "Name": "One of Smith",
      "Issuer": "Lord Smith"
    }
  ],
  "Skills": [
    {
      "Id": "133",
      "Name": "debug"
    }
  ],
  "CurrentStatus": [
    {
      "Id": "328",
      "Text": "Go for help",
      "Source": "testsource",
      "CreateDate": "2006-02-07T19:02:55.803Z"
    }
  ],
  "Certifications": [
    {
      "Id": "512",
      "Name": "Engineer certification",
      "Authority": "The Great Engineer Association",
      "Number": "1",
      "StartDate": "2006-02-07T19:02:55.803Z",
      "EndDate": "2007-02-07T19:02:55.803Z"
    }
  ],
  "Courses": [
    {
      "Id": "888",
      "Name": "Advanced Algorithm",
      "Number": "1"
    }
  ],
  "Volunteer": [
    {
      "Id": "456",
      "Role": "Teacher",
      "Organization": "The Great Wall",
      "Cause": "Teach people to defend the Wall"
    }
  ],
  "RecommendationsReceived": [
    {
      "Id": "846",
      "RecommendationType": "Primary",
      "RecommendationText": "The Great Wall",
      "Recommender": "King"
    }
  ],
  "Languages": [
    {
      "Id": "484",
      "Name": "Latin",
      "Proficiency": "Okay"
    }
  ],
  "Projects": [
    {
      "Id": "789",
      "Name": "Night Watcher",
      "Summary": "Defend the wall",
      "With": [
        {
          "Id": "555",
          "Name": "Meng"
        }
      ],
      "StartDate": "2010-02-07T19:02:55.803Z",
      "EndDate": "2011-02-07T19:02:55.803Z",
      "isCurrent": "false"
    }
  ],
  "Games": [
    {
      "Id": "777",
      "Name": "Working out",
      "Category": "Sports",
      "CreateDate": "2010-10-07T19:02:55.803Z"
    }
  ],
  "Family": [
    {
      "Id": "055",
      "Name": "Robb Smith",
      "Relationship": "Brother"
    },
    {
      "Id": "056",
      "Name": "Eddard Smith",
      "Relationship": "Father"
    }
  ],
  "TeleVisionShow": [
    {
      "Id": "777",
      "Name": "Getting out",
      "Category": "Comedy",
      "CreateDate": "2011-02-07T19:02:55.803Z"
    }
  ],
  "MutualFriends": [
    {
      "Id": "060",
      "Name": "Robb Smith",
      "FirstName": "Robb",
      "LastName": "Smith",
      "Birthday": "2000-08-24T12:21:41.808Z",
      "Hometown": "Vancouver",
      "Link": "classmate",
      "Gender": "M"
    },
    {
      "Id": "061",
      "Name": "Sansa Smith",
      "FirstName": "Sansa",
      "LastName": "Smith",
      "Birthday": "2000-08-24T12:21:41.808Z",
      "Hometown": "Vancouver",
      "Link": "classmate",
      "Gender": "F"
    }
  ],
  "Movies": [
    {
      "Id": "788",
      "Name": "Getting in",
      "Category": "Comedy",
      "CreateDate": "2017-02-07T19:02:55.803Z"
    }
  ],
  "Books": [
    {
      "Id": "788",
      "Name": "Getting in",
      "Category": "Interest",
      "CreateDate": "2000-01-15T19:02:55.803Z"
    }
  ],
  "AgeRange": {
    "Min": 15,
    "Max": 20
  },
  "PublicRepository": "Master",
  "Hireable": false,
  "Age": "18",
  "RepositoryUrl": "www.master.com",
  "Patents": [
    {
      "Id": "516",
      "Title": "POST",
      "Date": "2009-02-07T19:02:55.803Z"
    }
  ],
  "FavoriteThings": [
    {
      "Id": "781",
      "Name": "BasketBall",
      "Type": "Sports"
    }
  ],
  "ProfessionalHeadline": "dragon",
  "ProviderAccessCredential": {
    "AccessToken": "anlkeruna;sa",
    "TokenSecret": "adga!rgagads"
  },
  "RelatedProfileViews": [
    {
      "Id": "884",
      "FirstName": "Jemery",
      "LastName": "Fir"
    }
  ],
  "KloutScore": {
    "KloutId": "51656",
    "Score": "52.0"
  },
  "LRUserID": "526526",
  "PlacesLived": [
    {
      "Name": "Allision road",
      "IsPrimary": "True"
    }
  ],
  "Publications": [
    {
      "Id": "51656",
      "Title": "ASA",
      "Publisher": "Fisher",
      "Authors": [
        {
          "Id": "51",
          "Name": "Darker"
        }
      ],
      "Date": "1997-02-07T19:02:55.803Z",
      "Url": "whatwebsite.com",
      "Summary": "What would you do as soon as you can?"
    }
  ],
  "JobBookmarks": [
    {
      "IsApplied": true,
      "ApplyTimestamp": "541",
      "IsSaved": false,
      "SavedTimestamp": "588",
      "Job": {
        "Active": true,
        "Id": "51",
        "DescriptionSnippet": "Darker",
        "Compony": {
          "Id": "489",
          "Name": "Night Watch"
        },
        "Position": {
          "Title": "Commander"
        },
        "PostingTimestamp": "988"
      }
    }
  ],
  "Suggestions": {
    "CompaniestoFollow": [
      {
        "Id": "568",
        "Name": "Amazon"
      }
    ],
    "IndustriestoFollow": [
      {
        "Id": "695",
        "Name": "Software"
      }
    ],
    "NewssourcetoFollow": [
      {
        "Id": "352",
        "Name": "NewYorkFans"
      }
    ],
    "PeopletoFollow": [
      {
        "Id": "845",
        "Name": "Anna Smith"
      }
    ]
  },
  "Badges": [
    {
      "BadgeId": "531",
      "Name": "MVP",
      "BadgeMessage": "Most Valuable Person",
      "Description": "An award to triple double",
      "ImageUrl": "www.imageMVP.com"
    }
  ],
  "MemberUrlResources": [
    {
      "Url": "www.memberurl.com",
      "UrlName": "MemberURL"
    }
  ],
  "TotalPrivateRepository": 0,
  "Currency": "CAD",
  "StarredUrl": "www.image.com",
  "GistsUrl": "www.gists.com",
  "PublicGists": 0,
  "PrivateGists": 0,
  "Subscription": {
    "Name": "Song of Ice and Fire",
    "Space": "none",
    "PrivateRepos": "testRepo1",
    "Collaborators": "No idea"
  },
  "Company": "Night Watch",
  "GravatarImageUrl": "www.image.com",
  "ProfileImageUrls": {
    "test1": "https://graph.facebook.com/*****/picture?type=small",
    "test2": "https://graph.Profile.com/*****/picture?type=square"
  },
  "WebProfiles": {
    "Small": "https://graph.facebook.com/*****/picture?type=small",
    "Square": "https://graph.Profile.com/*****/picture?type=square",
    "Large": "https://graph.facebook.com/*****/picture?type=large",
    "Profile": "https://graph.facebook.com/******/picture?type=normal"
  },
  "IsEmailSubscribed": false
}`
