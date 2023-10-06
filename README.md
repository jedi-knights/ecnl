# ECNL Module

A Go module to provide access to ECNL data.

![GitHub Actions](https://github.com/jedi-knights/ecnl/workflows/CI/badge.svg)

## Local Development

### Requirements 

In order to get setup to develop in this project there are a number of requirements that need to be met 
by your development environment.  Although it is possible to set yourself up in other ways I would 
recommend that you follow the instructions below.

* [Homebrew Package Manager](https://brew.sh/)
* [OpenSSL](https://www.openssl.org/)
* [Node Version Manager](https://github.com/nvm-sh/nvm)
* [The Go Programming Language](https://go.dev/)


#### Setting up Homebrew

The following command will utilize curl to get the Homebrew install script over https and immediately
execute via Bash on your machine.

```sh
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```


#### Setting up the Node Version Manager

Unfortunately the Node Version Manager does not work with Homebrew.  You will need to install it manually.

```sh
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash
```

Installing Go and OpenSSL using Homebrew:

```sh
brew install go
brew install openssl
```

## SSL Setup

To use HTTPS with the Echo server, we need to create and configure an HTTPS server using the 'net/http' package and then use it as a handler for the Echo router.

Steps to set up HTTPS:

1. Generate SSL/TLS Certificates
2. Import Required Packages
3. Configure the Echo Server with HTTPS
4. Register routes and handlers as needed within the main function
5. Start the Echo server

Remember to replace "cert.pem" and "key.pem" with the actual paths to your SSL certificate and private key files.
Additionally, make sure that your firewall allows traffic on port 443 (the default HTTPS port).

Once the server is configured with HTTPS, it will listen for secure HTTPS connections on port 443, and you 
can access it using the https protocol.


### Generating SSL/TLS Certificates

You need SSL/TLS certificates to enable HTTPS.  You can either purchase a certificate from a certificate authority
or create a self-signed certificate for testing purposes.

In order to create a self-signed certificate you will need OpenSSL.  If you don't have it and are 
using Homebrew, you can install it with the following command, otherwise see https://www.openssl.org/.

```sh
brew install openssl
```

Creating a self-signed certificate using OpenSSL:

```sh
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
```




This command generates a self-signed certificate ('cert.pem') and a private key ('key.pem') and saves them in the current directory.
These files can be used for HTTPS.


## Get Organization List

GET /api/Script/get-org-club-list-by-orgID/{orgID}

GET /api/Script/get-club-info/{clubID}

GET /api/Script/get-club-complex-list-by-club/{clubID}

GET /api/Script/get-individual-club-info-with-teams-and-staff/{orgID}/{clubID}/{eventID}

GET /api/Script/get-club-schedules-by-eventID-and-clubID/{eventID}/{clubID}

GET /api/Script/get-event-list-by-season-id/{orgSeasonID}/{eventID}

GET /api/Script/get-division-list-by-event-id/{orgID}/{eventID}/{flightID}

GET /api/Script/get-conference-schedules/{orgID}/{orgSeasonID}/{eventID}/{flightID}/{complexID}

POST /api/Script/get-filtered-conference-schedules/{orgID}/{orgSeasonID}/{eventID}/{flightID}/{complexID}

GET /api/Script/get-conference-standings/{eventID}/{orgID}/{orgSeasonID}/{divisionID}/{standingValue}

GET /api/Script/get-game-complex/{eventID}/{complexID}/{venueID}/{zip}/{matchID}

This endpoint requires the following parameters:

* eventID - The ID of the event.  This can be retrieved from the ECNL Schedule page.
* orgID - The ID of the organization.  This can be retrieved from the ECNL Schedule page.
* orgSeasonID - The ID of the organization season.  This can be retrieved from the ECNL Schedule page.
* divisionID - The ID of the division.  This can be retrieved from the ECNL Schedule page.
* standingValue - The value of the standing.  This can be retrieved from the ECNL Schedule page.


eventId:
* ECNL Girls Southwest 2023-2024 [2837]
* ECNL Girls Southeast 2023-2024 [2838]
* ECNL Girls Mid-Atlantic 2023-2024 [2839]
* ECNL Girls Midwest 2023-2024 [2849]
* ECNL Girls North Atlantic 2023-2024 [2850]

There appears to be 6 ECNL Organizations.
* BOYS PRE-ECNL [22]
* ECNL Boys [12]
* ECNL Boys Regional League [16]
* ECNL Girls [9]
* ECNL Girls Regional League [13]
* GIRLS PRE-ECNL [21]

An Organization contains:
* Id
* Name
* SeasonId
* SeasonGroupId

All the orgnizations have the same SeasonGroupId of 8 so I don't think that is important.

The SeasonId varies across the organizations so I'm unsure as to its significance.

In order to retrieve a list of clubs I needed to provide an Organization's Id, which
for the "ECNL Girls" organziation is 9.

Since the concept of an organization appears to be at the root of their structure, I have 
added the ability to retrieve an organazation by name and id.

After examining the list of clubs for the ECNL Girls organization I
have noticed that all clubs have the same OrgId and OrgSeasonId.  I
assume that to mean that the OrgSeasonId represents the current season
across the league. 

In the list of clubs there is some variability in the EventId so
I am assuming that has something to do with the conference the
club is associated with.

Ok so I think I have figured out the EventId.  It appears to be
an identifier that represents the conference the club is associated
with.  There are 10 ECNL Girls conferences and 10 event Ids so I 
think that is the case.

Each event appears to represent

## Example

Step 1: Determine an Organization

```text
There are a total of 9 organizations.
0: Name: "BOYS PRE-ECNL", Id: 22, SeasonId: 56, SeasonGroupId: 8
1: Name: "ECNL Boys", Id: 12, SeasonId: 50, SeasonGroupId: 8
2: Name: "ECNL Boys Regional League", Id: 16, SeasonId: 52, SeasonGroupId: 8
3: Name: "ECNL Girls", Id: 9, SeasonId: 49, SeasonGroupId: 8
4: Name: "ECNL Girls Regional League", Id: 13, SeasonId: 51, SeasonGroupId: 8
5: Name: "GIRLS PRE-ECNL", Id: 21, SeasonId: 55, SeasonGroupId: 8
6: Name: "Texas Club Soccer League", Id: 23, SeasonId: 59, SeasonGroupId: 0
7: Name: "US Club Soccer", Id: 4, SeasonId: 54, SeasonGroupId: 8
8: Name: "USC TX", Id: 18, SeasonId: 53, SeasonGroupId: 8
```

I Select Organization 9 or ECNL Girls

ECNL Girls Clubs

```text
There are a total of 128 clubs.
Name: "Alabama FC", OrgId: 9, OrgSeasonId: 49, ClubId: 18, City: "Birmingham", StateCode: AL, EventId: 2838, EventCounts: 0
Name: "Albion Hurricanes FC", OrgId: 9, OrgSeasonId: 49, ClubId: 10, City: "Houston", StateCode: TX, EventId: 2857, EventCounts: 0
Name: "Arlington Soccer", OrgId: 9, OrgSeasonId: 49, ClubId: 525, City: "Arlington", StateCode: VA, EventId: 2839, EventCounts: 0
Name: "Atlanta Fire United", OrgId: 9, OrgSeasonId: 49, ClubId: 14, City: "Duluth", StateCode: GA, EventId: 2838, EventCounts: 0
Name: "AZ Arsenal", OrgId: 9, OrgSeasonId: 49, ClubId: 592, City: "Phoenix", StateCode: AZ, EventId: 2837, EventCounts: 0
Name: "Bay Area Surf", OrgId: 9, OrgSeasonId: 49, ClubId: 3986, City: "Bay Area", StateCode: CA, EventId: 2854, EventCounts: 0
Name: "Beach FC  (CA)", OrgId: 9, OrgSeasonId: 49, ClubId: 16, City: "Torrance", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "Beach FC  (VA)", OrgId: 9, OrgSeasonId: 49, ClubId: 628, City: "Virginia Beach", StateCode: VA, EventId: 2839, EventCounts: 0
Name: "Bethesda SC", OrgId: 9, OrgSeasonId: 49, ClubId: 17, City: "Gaithersburg", StateCode: MD, EventId: 2850, EventCounts: 0
Name: "Boise Timbers | Thor...", OrgId: 9, OrgSeasonId: 49, ClubId: 51, City: "Meridian", StateCode: ID, EventId: 2855, EventCounts: 0
Name: "Carolina Elite Socce...", OrgId: 9, OrgSeasonId: 49, ClubId: 380, City: "Greenville", StateCode: SC, EventId: 2838, EventCounts: 0
Name: "Challenge SC", OrgId: 9, OrgSeasonId: 49, ClubId: 25, City: "Spring", StateCode: TX, EventId: 2857, EventCounts: 0
Name: "Charlotte Independen...", OrgId: 9, OrgSeasonId: 49, ClubId: 1010, City: "Cornelius", StateCode: NC, EventId: 2839, EventCounts: 0
Name: "Charlotte SA", OrgId: 9, OrgSeasonId: 49, ClubId: 26, City: "Pineville ", StateCode: NC, EventId: 2839, EventCounts: 0
Name: "Chicago Inter Soccer", OrgId: 9, OrgSeasonId: 49, ClubId: 2167, City: "New Lenox", StateCode: IL, EventId: 2849, EventCounts: 0
Name: "Classics Elite Socce...", OrgId: 9, OrgSeasonId: 49, ClubId: 1375, City: "San Antonio", StateCode: TX, EventId: 2857, EventCounts: 0
Name: "Cleveland Force SC", OrgId: 9, OrgSeasonId: 49, ClubId: 3478, City: "Bedford Heights", StateCode: OH, EventId: 2856, EventCounts: 0
Name: "Colorado Rapids", OrgId: 9, OrgSeasonId: 49, ClubId: 29, City: "Aurora", StateCode: CO, EventId: 2857, EventCounts: 0
Name: "Concorde Fire Platinum", OrgId: 9, OrgSeasonId: 49, ClubId: 2360, City: "Atlanta", StateCode: GA, EventId: 2838, EventCounts: 0
Name: "Concorde Fire Premier", OrgId: 9, OrgSeasonId: 49, ClubId: 30, City: "Atlanta", StateCode: GA, EventId: 2838, EventCounts: 0
Name: "Connecticut FC", OrgId: 9, OrgSeasonId: 49, ClubId: 31, City: "Bethany", StateCode: CT, EventId: 2853, EventCounts: 0
Name: "Crossfire Premier", OrgId: 9, OrgSeasonId: 49, ClubId: 33, City: "Redmond", StateCode: WA, EventId: 2855, EventCounts: 0
Name: "Dallas Texans", OrgId: 9, OrgSeasonId: 49, ClubId: 37, City: "Dallas", StateCode: TX, EventId: 2857, EventCounts: 0
Name: "Davis Legacy", OrgId: 9, OrgSeasonId: 49, ClubId: 1425, City: "Davis", StateCode: CA, EventId: 2854, EventCounts: 0
Name: "De Anza Force", OrgId: 9, OrgSeasonId: 49, ClubId: 2510, City: "Cupertino", StateCode: CA, EventId: 2854, EventCounts: 0
Name: "Del Mar Sharks", OrgId: 9, OrgSeasonId: 49, ClubId: 39, City: "San Diego", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "DKSC", OrgId: 9, OrgSeasonId: 49, ClubId: 35, City: "Farmers Branch", StateCode: TX, EventId: 2857, EventCounts: 0
Name: "Eagles Soccer Club", OrgId: 9, OrgSeasonId: 49, ClubId: 41, City: "Camarillo", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "East Meadow SC", OrgId: 9, OrgSeasonId: 49, ClubId: 42, City: "North Bellmore", StateCode: NY, EventId: 2853, EventCounts: 0
Name: "Eastside FC", OrgId: 9, OrgSeasonId: 49, ClubId: 398, City: "Issaquah", StateCode: WA, EventId: 2855, EventCounts: 0
Name: "Eclipse Select SC", OrgId: 9, OrgSeasonId: 49, ClubId: 43, City: "Arlington Heights", StateCode: IL, EventId: 2849, EventCounts: 0
Name: "Fairfax BRAVE Soccer...", OrgId: 9, OrgSeasonId: 49, ClubId: 4058, City: "Vienna", StateCode: VA, EventId: 2839, EventCounts: 0
Name: "FC Alliance", OrgId: 9, OrgSeasonId: 49, ClubId: 1693, City: "Knoxville", StateCode: TN, EventId: 2856, EventCounts: 0
Name: "FC Dallas", OrgId: 9, OrgSeasonId: 49, ClubId: 47, City: "Dallas", StateCode: TX, EventId: 2857, EventCounts: 0
Name: "FC DELCO", OrgId: 9, OrgSeasonId: 49, ClubId: 598, City: "Conshohocken", StateCode: PA, EventId: 2850, EventCounts: 0
Name: "FC Pride", OrgId: 9, OrgSeasonId: 49, ClubId: 1501, City: "Indianapolis", StateCode: IN, EventId: 2856, EventCounts: 0
Name: "FC Prime", OrgId: 9, OrgSeasonId: 49, ClubId: 2329, City: "Sunrise", StateCode: FL, EventId: 2838, EventCounts: 0
Name: "FC Stars Blue ", OrgId: 9, OrgSeasonId: 49, ClubId: 2358, City: "Acton", StateCode: MA, EventId: 2853, EventCounts: 0
Name: "FC Stars White", OrgId: 9, OrgSeasonId: 49, ClubId: 52, City: "Acton", StateCode: MA, EventId: 2853, EventCounts: 0
Name: "FC Wisconsin", OrgId: 9, OrgSeasonId: 49, ClubId: 54, City: "Germantown", StateCode: WI, EventId: 2849, EventCounts: 0
Name: "Florida Elite Soccer...", OrgId: 9, OrgSeasonId: 49, ClubId: 1353, City: "St Johns", StateCode: FL, EventId: 2838, EventCounts: 0
Name: "Florida Kraze Krush", OrgId: 9, OrgSeasonId: 49, ClubId: 379, City: "Winter Springs", StateCode: FL, EventId: 2838, EventCounts: 0
Name: "Florida Premier FC", OrgId: 9, OrgSeasonId: 49, ClubId: 2731, City: "New Port Richey", StateCode: FL, EventId: 2838, EventCounts: 0
Name: "Florida West FC", OrgId: 9, OrgSeasonId: 49, ClubId: 3384, City: "Fort Myers", StateCode: FL, EventId: 2838, EventCounts: 0
Name: "FSA FC", OrgId: 9, OrgSeasonId: 49, ClubId: 431, City: "Farmington", StateCode: CT, EventId: 2853, EventCounts: 0
Name: "Gretna Elite Academy", OrgId: 9, OrgSeasonId: 49, ClubId: 373, City: "Gretna", StateCode: NE, EventId: 2849, EventCounts: 0
Name: "GSA", OrgId: 9, OrgSeasonId: 49, ClubId: 56, City: "Liburn", StateCode: GA, EventId: 2838, EventCounts: 0
Name: "Heat FC", OrgId: 9, OrgSeasonId: 49, ClubId: 57, City: "Henderson", StateCode: NV, EventId: 2837, EventCounts: 0
Name: "HEX FC", OrgId: 9, OrgSeasonId: 49, ClubId: 46, City: "Perkasie", StateCode: PA, EventId: 2850, EventCounts: 0
Name: "Indiana Fire ", OrgId: 9, OrgSeasonId: 49, ClubId: 436, City: "Westfield", StateCode: IN, EventId: 2856, EventCounts: 0
Name: "Internationals SC", OrgId: 9, OrgSeasonId: 49, ClubId: 58, City: "Medina", StateCode: OH, EventId: 2856, EventCounts: 0
Name: "Jacksonville FC", OrgId: 9, OrgSeasonId: 49, ClubId: 990, City: "Jacksonville", StateCode: FL, EventId: 2838, EventCounts: 0
Name: "Kansas City Athletics", OrgId: 9, OrgSeasonId: 49, ClubId: 2330, City: "Merriam", StateCode: KS, EventId: 2849, EventCounts: 0
Name: "Kings Hammer SC", OrgId: 9, OrgSeasonId: 49, ClubId: 3491, City: "Covington", StateCode: KY, EventId: 2856, EventCounts: 0
Name: "LA Breakers FC", OrgId: 9, OrgSeasonId: 49, ClubId: 2289, City: "Los Angeles", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "La Roca", OrgId: 9, OrgSeasonId: 49, ClubId: 1760, City: "South Weber", StateCode: UT, EventId: 2855, EventCounts: 0
Name: "LAFC SO CAL", OrgId: 9, OrgSeasonId: 49, ClubId: 82, City: "Woodland Hills", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "Legends FC", OrgId: 9, OrgSeasonId: 49, ClubId: 61, City: "Eastvale ", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "Liverpool FC IA Mich...", OrgId: 9, OrgSeasonId: 49, ClubId: 3712, City: "Pontiac", StateCode: MI, EventId: 2849, EventCounts: 0
Name: "Marin FC", OrgId: 9, OrgSeasonId: 49, ClubId: 1585, City: "Greenbrae", StateCode: CA, EventId: 2854, EventCounts: 0
Name: "Maryland United FC", OrgId: 9, OrgSeasonId: 49, ClubId: 419, City: "Bowie", StateCode: MD, EventId: 2850, EventCounts: 0
Name: "Match Fit Academy", OrgId: 9, OrgSeasonId: 49, ClubId: 64, City: "Mercer County", StateCode: NJ, EventId: 2850, EventCounts: 0
Name: "McLean Youth Soccer", OrgId: 9, OrgSeasonId: 49, ClubId: 65, City: "McLean", StateCode: VA, EventId: 2839, EventCounts: 0
Name: "Michigan Hawks", OrgId: 9, OrgSeasonId: 49, ClubId: 2359, City: "Plymouth", StateCode: MI, EventId: 2849, EventCounts: 0
Name: "Michigan Hawks Magic", OrgId: 9, OrgSeasonId: 49, ClubId: 3457, City: "Plymouth", StateCode: MI, EventId: 2849, EventCounts: 0
Name: "Midwest United FC", OrgId: 9, OrgSeasonId: 49, ClubId: 3721, City: "Kentwood", StateCode: MI, EventId: 2849, EventCounts: 0
Name: "Minnesota Thunder Ac...", OrgId: 9, OrgSeasonId: 49, ClubId: 67, City: "Richfield", StateCode: MN, EventId: 2849, EventCounts: 0
Name: "Mustang SC", OrgId: 9, OrgSeasonId: 49, ClubId: 69, City: "Danville", StateCode: CA, EventId: 2854, EventCounts: 0
Name: "MVLA Soccer Club", OrgId: 9, OrgSeasonId: 49, ClubId: 294, City: "Mountain View", StateCode: CA, EventId: 2854, EventCounts: 0
Name: "NC Courage", OrgId: 9, OrgSeasonId: 49, ClubId: 3404, City: "RALEIGH", StateCode: NC, EventId: 2839, EventCounts: 0
Name: "NC Fusion", OrgId: 9, OrgSeasonId: 49, ClubId: 71, City: "Bermuda Run", StateCode: NC, EventId: 2839, EventCounts: 0
Name: "NCFC Youth", OrgId: 9, OrgSeasonId: 49, ClubId: 22, City: "Raleigh", StateCode: NC, EventId: 2839, EventCounts: 0
Name: "Northern Virginia Al...", OrgId: 9, OrgSeasonId: 49, ClubId: 4369, City: "Leesburg", StateCode: VA, EventId: 2839, EventCounts: 0
Name: "Northwest Elite FC", OrgId: 9, OrgSeasonId: 49, ClubId: 737, City: "Beaverton", StateCode: OR, EventId: 2855, EventCounts: 0
Name: "Ohio Elite Soccer Ac...", OrgId: 9, OrgSeasonId: 49, ClubId: 75, City: "Cincinnati", StateCode: OH, EventId: 2856, EventCounts: 0
Name: "Ohio Premier", OrgId: 9, OrgSeasonId: 49, ClubId: 437, City: "Dublin", StateCode: OH, EventId: 2856, EventCounts: 0
Name: "Oklahoma Energy FC", OrgId: 9, OrgSeasonId: 49, ClubId: 2133, City: "Edmond", StateCode: OK, EventId: 2857, EventCounts: 0
Name: "Orlando City Youth S...", OrgId: 9, OrgSeasonId: 49, ClubId: 77, City: "Sanford", StateCode: FL, EventId: 2838, EventCounts: 0
Name: "Pacific Northwest So...", OrgId: 9, OrgSeasonId: 49, ClubId: 2331, City: "Tukwila", StateCode: WA, EventId: 2855, EventCounts: 0
Name: "Pateadores", OrgId: 9, OrgSeasonId: 49, ClubId: 186, City: "Costa Mesa ", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "PDA", OrgId: 9, OrgSeasonId: 49, ClubId: 78, City: "Somerset", StateCode: NJ, EventId: 2850, EventCounts: 0
Name: "PDA Blue", OrgId: 9, OrgSeasonId: 49, ClubId: 2357, City: "Somerset", StateCode: NJ, EventId: 2850, EventCounts: 0
Name: "PDA South", OrgId: 9, OrgSeasonId: 49, ClubId: 2628, City: "Westampton", StateCode: NJ, EventId: 2850, EventCounts: 0
Name: "Penn Fusion SA", OrgId: 9, OrgSeasonId: 49, ClubId: 79, City: "West Chester", StateCode: PA, EventId: 2850, EventCounts: 0
Name: "Phoenix Rising FC", OrgId: 9, OrgSeasonId: 49, ClubId: 1141, City: "Scottsdale", StateCode: AZ, EventId: 2837, EventCounts: 0
Name: "Pipeline SC", OrgId: 9, OrgSeasonId: 49, ClubId: 3405, City: "Baldwin", StateCode: MD, EventId: 2850, EventCounts: 0
Name: "Pittsburgh Riverhounds", OrgId: 9, OrgSeasonId: 49, ClubId: 538, City: "Pittsburgh", StateCode: PA, EventId: 2856, EventCounts: 0
Name: "Placer United Soccer...", OrgId: 9, OrgSeasonId: 49, ClubId: 3054, City: "Rocklin", StateCode: CA, EventId: 2854, EventCounts: 0
Name: "Pleasanton Rage", OrgId: 9, OrgSeasonId: 49, ClubId: 80, City: "Pleasanton", StateCode: CA, EventId: 2854, EventCounts: 0
Name: "Portland Thorns Acad...", OrgId: 9, OrgSeasonId: 49, ClubId: 3406, City: "Portland", StateCode: OR, EventId: 2855, EventCounts: 0
Name: "Pride Soccer Club", OrgId: 9, OrgSeasonId: 49, ClubId: 2335, City: "Colorado Springs", StateCode: CO, EventId: 2857, EventCounts: 0
Name: "Racing Louisville Ac...", OrgId: 9, OrgSeasonId: 49, ClubId: 3408, City: "Louisville", StateCode: KY, EventId: 2856, EventCounts: 0
Name: "Real Colorado", OrgId: 9, OrgSeasonId: 49, ClubId: 81, City: "Centennial", StateCode: CO, EventId: 2857, EventCounts: 0
Name: "Real Colorado Athlet...", OrgId: 9, OrgSeasonId: 49, ClubId: 3470, City: "Centennial", StateCode: CO, EventId: 2857, EventCounts: 0
Name: "Rebels SC", OrgId: 9, OrgSeasonId: 49, ClubId: 300, City: "Chula Vista", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "Richmond United", OrgId: 9, OrgSeasonId: 49, ClubId: 417, City: "Glen Allen", StateCode: VA, EventId: 2839, EventCounts: 0
Name: "Rockford Raptors FC", OrgId: 9, OrgSeasonId: 49, ClubId: 2017, City: "Loves park", StateCode: IL, EventId: 2849, EventCounts: 0
Name: "San Diego Surf Socce...", OrgId: 9, OrgSeasonId: 49, ClubId: 87, City: "San Diego", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "San Juan SC", OrgId: 9, OrgSeasonId: 49, ClubId: 88, City: "Rancho Cordova", StateCode: CA, EventId: 2854, EventCounts: 0
Name: "Santa Rosa United", OrgId: 9, OrgSeasonId: 49, ClubId: 91, City: "Santa Rosa", StateCode: CA, EventId: 2854, EventCounts: 0
Name: "Scorpions SC", OrgId: 9, OrgSeasonId: 49, ClubId: 94, City: "Auburn", StateCode: MA, EventId: 2853, EventCounts: 0
Name: "Seattle United", OrgId: 9, OrgSeasonId: 49, ClubId: 1643, City: "Seattle", StateCode: WA, EventId: 2855, EventCounts: 0
Name: "SLAMMERS FC", OrgId: 9, OrgSeasonId: 49, ClubId: 96, City: "Newport Beach", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "Slammers FC HB Koge", OrgId: 9, OrgSeasonId: 49, ClubId: 1656, City: "Costa Mesa", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "SLSG Green", OrgId: 9, OrgSeasonId: 49, ClubId: 4064, City: "Fenton", StateCode: MO, EventId: 2849, EventCounts: 0
Name: "SLSG MO", OrgId: 9, OrgSeasonId: 49, ClubId: 97, City: "Fenton", StateCode: MO, EventId: 2849, EventCounts: 0
Name: "So Cal Blues SC", OrgId: 9, OrgSeasonId: 49, ClubId: 98, City: "Laguna Hills", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "Solar Soccer Club", OrgId: 9, OrgSeasonId: 49, ClubId: 101, City: "Allen", StateCode: TX, EventId: 2857, EventCounts: 0
Name: "South Carolina United", OrgId: 9, OrgSeasonId: 49, ClubId: 1689, City: "Columbia", StateCode: SC, EventId: 2838, EventCounts: 0
Name: "Sporting Blue Valley", OrgId: 9, OrgSeasonId: 49, ClubId: 102, City: "Overland Park", StateCode: KS, EventId: 2849, EventCounts: 0
Name: "Sporting California ...", OrgId: 9, OrgSeasonId: 49, ClubId: 13, City: "Ontario", StateCode: CA, EventId: 2837, EventCounts: 0
Name: "Sporting Iowa", OrgId: 9, OrgSeasonId: 49, ClubId: 3483, City: "Des Moines", StateCode: IA, EventId: 2849, EventCounts: 0
Name: "Sting Austin", OrgId: 9, OrgSeasonId: 49, ClubId: 1690, City: "Pflugerville", StateCode: TX, EventId: 2857, EventCounts: 0
Name: "Sting Dallas Black", OrgId: 9, OrgSeasonId: 49, ClubId: 36, City: "Dallas", StateCode: TX, EventId: 2857, EventCounts: 0
Name: "Sting Dallas Royal", OrgId: 9, OrgSeasonId: 49, ClubId: 3437, City: "Dallas", StateCode: TX, EventId: 2857, EventCounts: 0
Name: "Strikers FC ECNL", OrgId: 9, OrgSeasonId: 49, ClubId: 2, City: "Irvine", StateCode: CA, EventId: 0, EventCounts: 0
Name: "SUSA FC", OrgId: 9, OrgSeasonId: 49, ClubId: 1404, City: "Central Islip", StateCode: NY, EventId: 2853, EventCounts: 0
Name: "Tampa Bay United", OrgId: 9, OrgSeasonId: 49, ClubId: 874, City: "Tampa", StateCode: FL, EventId: 2838, EventCounts: 0
Name: "Tennessee Soccer Club", OrgId: 9, OrgSeasonId: 49, ClubId: 1354, City: "Franklin", StateCode: TN, EventId: 2856, EventCounts: 0
Name: "Tulsa SC", OrgId: 9, OrgSeasonId: 49, ClubId: 106, City: "Tulsa", StateCode: OK, EventId: 2857, EventCounts: 0
Name: "United Futbol Academy", OrgId: 9, OrgSeasonId: 49, ClubId: 3407, City: "Cumming", StateCode: GA, EventId: 2838, EventCounts: 0
Name: "Utah Avalanche", OrgId: 9, OrgSeasonId: 49, ClubId: 107, City: "Salt Lake City", StateCode: UT, EventId: 2855, EventCounts: 0
Name: "Utah Royals FC", OrgId: 9, OrgSeasonId: 49, ClubId: 3716, City: "Mesa", StateCode: AZ, EventId: 2837, EventCounts: 0
Name: "Virginia Development...", OrgId: 9, OrgSeasonId: 49, ClubId: 1785, City: "Woodbridge", StateCode: VA, EventId: 2839, EventCounts: 0
Name: "Washington Premier", OrgId: 9, OrgSeasonId: 49, ClubId: 112, City: "Puyallup", StateCode: WA, EventId: 2855, EventCounts: 0
Name: "Wilmington Hammerhea...", OrgId: 9, OrgSeasonId: 49, ClubId: 2181, City: "Wilmington", StateCode: NC, EventId: 2839, EventCounts: 0
Name: "WNY Flash", OrgId: 9, OrgSeasonId: 49, ClubId: 343, City: "Elma ", StateCode: NY, EventId: 2856, EventCounts: 0
Name: "World Class FC", OrgId: 9, OrgSeasonId: 49, ClubId: 116, City: "Orangeburg", StateCode: NY, EventId: 2853, EventCounts: 0
```

Events for organization "ECNL Girls" (9)

```text
There are a total of 10 events.
Name: "ECNL Girls Mid-Atlantic 2023-24", Id: 2839, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
Name: "ECNL Girls Midwest 2023-24", Id: 2849, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
Name: "ECNL Girls New England 2023-24", Id: 2853, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
Name: "ECNL Girls North Atlantic 2023-24", Id: 2850, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
Name: "ECNL Girls Northern Cal 2023-24", Id: 2854, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
Name: "ECNL Girls Northwest 2023-24", Id: 2855, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
Name: "ECNL Girls Ohio Valley 2023-24", Id: 2856, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
Name: "ECNL Girls Southeast 2023-24", Id: 2838, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
Name: "ECNL Girls Southwest 2023-24", Id: 2837, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
Name: "ECNL Girls Texas 2023-24", Id: 2857, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
```

Now I need to find all the match results for "ECNL Girls".

How do I get match results?

What do I know?

- OrgId = 9
- OrgSeasonId = 49
- Each club seems to have different events.
- Events seem to be conferences for a particular season (for example ECNL Girls Southeast 2023-2024 [2838])

I get the events for the "ECNL Girls" organization by specifying the organizaiton Id = 9.

Now for each of these events I need to get the match results using

orgId and event.ID somehow

I'm using the Southeast conference as an example.

```text
Name: "ECNL Girls Southeast 2023-24", Id: 2838, OrgId: 9, OrgName: ECNL Girls, OrgSeasonId: 49, OrgSeasonName: "ECNL Girls 2023-24 Season"
```

So

- OrgId = 9
- EventId = 2838
- OrgSeasonId = 49
- ClubId = 30 (Concorde Fire Premier)


/api/Club/get-club-schedule-list/{clubID}/{eventID}
This appears to give us the club schedule for all age groups

/api/Club/get-score-reporting-schedule-list/{clubID}/{eventID}
This gives us the scores for all the matches for a given club across the conference

/api/Club/get-event-divisions-by-event-and-gender/{eventID}/{gender}
This gives us genders by event and gender (where gender is M or F)

It returns a response like:

```json
{
  "result": "success",
  "data": [
    {
      "divisionID": 11600,
      "divisionName": "G2006/2005"
    },
    {
      "divisionID": 11599,
      "divisionName": "G2007"
    },
    {
      "divisionID": 11598,
      "divisionName": "G2008"
    },
    {
      "divisionID": 11597,
      "divisionName": "G2009"
    },
    {
      "divisionID": 11596,
      "divisionName": "G2010"
    },
    {
      "divisionID": 11595,
      "divisionName": "G2011"
    }
  ]
}
```

Since I'm interested in G2009 I know my divisionId is 11597.

So

- OrgId = 9
- EventId = 2838
- OrgSeasonId = 49
- ClubId = 30 (Concorde Fire Premier)
- DivisionId = 11597 (G2009)

/api/Event/get-flight-list-by-division/{divisionID}
This gives us the flights for a given division
It returns something like this

```json
{
  "result": "success",
  "data": [
    {
      "flightID": 19789,
      "name": "ECNL"
    }
  ]
}
```

We can also get the teams in the conference for a given age group once we know the eventID and divisionID

/api/Event/get-event-division-teams/{eventID}/{divisionID}

This returns something like this

```json
{
  "result": "success",
  "data": {
    "teamList": [
      {
        "teamID": 58734,
        "teamName": "Alabama FC ECNL G09",
        "status": 2,
        "clubID": 18,
        "initialSeed": 1,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/_da463c0400e545cbbfb1f23ee81d0cde_afc.png",
        "firstName": "Thomas",
        "lastName": "Brower",
        "wdl": "32W 10D 16L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 60033,
        "teamName": "Atlanta Fire United ECNL G09",
        "status": 2,
        "clubID": 14,
        "initialSeed": 2,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Lee_Wilder_db1845a2eea8444082a1d8c62031bb81_ATLFire_Un.png",
        "firstName": "Domenic",
        "lastName": "Martelli",
        "wdl": "15W 7D 30L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 46817,
        "teamName": "CESA Liberty ECNL G09",
        "status": 2,
        "clubID": 380,
        "initialSeed": 3,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/ClubImages/Logos/CarolinaESA.jpg",
        "firstName": "James",
        "lastName": "Atkins",
        "wdl": "40W 17D 23L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 59255,
        "teamName": "Concorde Fire Platinum ECNL G09",
        "status": 2,
        "clubID": 2360,
        "initialSeed": 4,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Kevin_Sanchez_b3827dbbef3340c6a827bb5601cfbde9_ConcordeFi.jpg",
        "firstName": "Jonn",
        "lastName": "Warde",
        "wdl": "39W 7D 12L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 59236,
        "teamName": "Concorde Fire Premier ECNL G09",
        "status": 2,
        "clubID": 30,
        "initialSeed": 5,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/ClubImages/Logos/Concorde-Fire.jpg",
        "firstName": "Brian",
        "lastName": "Kelly",
        "wdl": "13W 11D 25L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 46728,
        "teamName": "FC Prime ECNL G09",
        "status": 2,
        "clubID": 2329,
        "initialSeed": 6,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Adrian_Mosquera_06c95d2a2f604090a8e440d26da037db_2021_fcp_l.png",
        "firstName": "Tyrone",
        "lastName": "Mears",
        "wdl": "40W 14D 10L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 69292,
        "teamName": "FL Premier FC ECNL G09",
        "status": 2,
        "clubID": 2731,
        "initialSeed": 7,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Dylan_Hughes_b9b0e98257d84392860e8301596f98fd_Florida_Premier_TGS.jpg",
        "firstName": "Stuart",
        "lastName": "Campbell",
        "wdl": "11W 4D 12L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 45064,
        "teamName": "Florida Elite ECNL G09",
        "status": 2,
        "clubID": 1353,
        "initialSeed": 8,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/ClubImage/Logo For ECNL Page48DD2017034806.pn",
        "firstName": "Michael",
        "lastName": "Dowling",
        "wdl": "29W 13D 24L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 34958,
        "teamName": "Florida Krush ECNL G09",
        "status": 2,
        "clubID": 379,
        "initialSeed": 9,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/ClubImages/Logos/FKK.png",
        "firstName": "Mark",
        "lastName": "Hansen",
        "wdl": "11W 8D 37L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 45629,
        "teamName": "Florida West FC ECNL G09",
        "status": 2,
        "clubID": 3384,
        "initialSeed": 10,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Daniel_Padilla_6c61c8269be74ea0988e03471528def5_Florida-West-Football-Club.png",
        "firstName": "Daniel",
        "lastName": "Fahey",
        "wdl": "27W 7D 25L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 59311,
        "teamName": "GSA ECNL G09",
        "status": 2,
        "clubID": 56,
        "initialSeed": 11,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/kk_kk_93be4e524dee4d938f98de1b01e720ab_GSALogo18D.jpg",
        "firstName": "Russell",
        "lastName": "Staggs",
        "wdl": "28W 11D 19L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 45667,
        "teamName": "Jacksonville FC ECNL G09",
        "status": 2,
        "clubID": 990,
        "initialSeed": 12,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/JOEL_HERNANDEZ_310dd57f9c604f8a9237bbafc95329d7_JFC_Logo.jpg",
        "firstName": "Ayomide",
        "lastName": "Leyimu",
        "wdl": "8W 2D 49L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 59913,
        "teamName": "Orlando City ECNL G09",
        "status": 2,
        "clubID": 77,
        "initialSeed": 13,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/_0ede1643f06145889d4f593cc8d4b0a1_LeagueTeam.png",
        "firstName": "Craig",
        "lastName": "Melton",
        "wdl": "9W 15D 28L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 50959,
        "teamName": "South Carolina United ECNL G09",
        "status": 2,
        "clubID": 1689,
        "initialSeed": 14,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Yoshikl_Sako_4009bd6b36b14e96ae900fd4b47b8b6f_SCUFC_TM_s.jpg",
        "firstName": "Adrian",
        "lastName": "Pinasco",
        "wdl": "21W 4D 30L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 49152,
        "teamName": "Tampa Bay United ECNL G09",
        "status": 2,
        "clubID": 874,
        "initialSeed": 15,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Hugo_Acosta_15ad970dfb1e4ea39726de93b4ff5a6e_tbu_2_star.png",
        "firstName": "Kayla",
        "lastName": "Corona Grimsley",
        "wdl": "13W 8D 36L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 59536,
        "teamName": "United Futbol Academy ECNL G09",
        "status": 2,
        "clubID": 3407,
        "initialSeed": 16,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Jonathon_Lake_ce2883d65e694e248155158ad7164fa4_2020 UFA.png",
        "firstName": "Gilbert",
        "lastName": "Jean-Baptiste",
        "wdl": "44W 6D 9L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      }
    ],
    "teamsCountForFlightList": [
      {
        "teamsCount": 16,
        "flightID": 19789,
        "flightName": "ECNL"
      }
    ],
    "flightGroups": []
  }
}
```

So

- OrgId = 9
- EventId = 2838
- OrgSeasonId = 49
- ClubId = 30 (Concorde Fire Premier)
- DivisionId = 11597 (G2009)
- FlightId = 19789 (ECNL)

But there could be more than one flight for a given division.

Now that we have the FlightId

/api/Event/get-event-division-teams/{eventID}/{divisionID}/{flightID}
This gives us all of the teams in the age group

It gives us something like this.

```json
{
  "result": "success",
  "data": {
    "teamList": [
      {
        "teamID": 58734,
        "teamName": "Alabama FC ECNL G09",
        "status": 2,
        "clubID": 18,
        "initialSeed": 1,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/_da463c0400e545cbbfb1f23ee81d0cde_afc.png",
        "firstName": "Thomas",
        "lastName": "Brower",
        "wdl": "32W 10D 16L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 60033,
        "teamName": "Atlanta Fire United ECNL G09",
        "status": 2,
        "clubID": 14,
        "initialSeed": 2,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Lee_Wilder_db1845a2eea8444082a1d8c62031bb81_ATLFire_Un.png",
        "firstName": "Domenic",
        "lastName": "Martelli",
        "wdl": "15W 7D 30L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 46817,
        "teamName": "CESA Liberty ECNL G09",
        "status": 2,
        "clubID": 380,
        "initialSeed": 3,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/ClubImages/Logos/CarolinaESA.jpg",
        "firstName": "James",
        "lastName": "Atkins",
        "wdl": "40W 17D 23L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 59255,
        "teamName": "Concorde Fire Platinum ECNL G09",
        "status": 2,
        "clubID": 2360,
        "initialSeed": 4,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Kevin_Sanchez_b3827dbbef3340c6a827bb5601cfbde9_ConcordeFi.jpg",
        "firstName": "Jonn",
        "lastName": "Warde",
        "wdl": "39W 7D 12L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 59236,
        "teamName": "Concorde Fire Premier ECNL G09",
        "status": 2,
        "clubID": 30,
        "initialSeed": 5,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/ClubImages/Logos/Concorde-Fire.jpg",
        "firstName": "Brian",
        "lastName": "Kelly",
        "wdl": "13W 11D 25L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 46728,
        "teamName": "FC Prime ECNL G09",
        "status": 2,
        "clubID": 2329,
        "initialSeed": 6,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Adrian_Mosquera_06c95d2a2f604090a8e440d26da037db_2021_fcp_l.png",
        "firstName": "Tyrone",
        "lastName": "Mears",
        "wdl": "40W 14D 10L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 69292,
        "teamName": "FL Premier FC ECNL G09",
        "status": 2,
        "clubID": 2731,
        "initialSeed": 7,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Dylan_Hughes_b9b0e98257d84392860e8301596f98fd_Florida_Premier_TGS.jpg",
        "firstName": "Stuart",
        "lastName": "Campbell",
        "wdl": "11W 4D 12L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 45064,
        "teamName": "Florida Elite ECNL G09",
        "status": 2,
        "clubID": 1353,
        "initialSeed": 8,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/ClubImage/Logo For ECNL Page48DD2017034806.pn",
        "firstName": "Michael",
        "lastName": "Dowling",
        "wdl": "29W 13D 24L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 34958,
        "teamName": "Florida Krush ECNL G09",
        "status": 2,
        "clubID": 379,
        "initialSeed": 9,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/ClubImages/Logos/FKK.png",
        "firstName": "Mark",
        "lastName": "Hansen",
        "wdl": "11W 8D 37L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 45629,
        "teamName": "Florida West FC ECNL G09",
        "status": 2,
        "clubID": 3384,
        "initialSeed": 10,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Daniel_Padilla_6c61c8269be74ea0988e03471528def5_Florida-West-Football-Club.png",
        "firstName": "Daniel",
        "lastName": "Fahey",
        "wdl": "27W 7D 25L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 59311,
        "teamName": "GSA ECNL G09",
        "status": 2,
        "clubID": 56,
        "initialSeed": 11,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/kk_kk_93be4e524dee4d938f98de1b01e720ab_GSALogo18D.jpg",
        "firstName": "Russell",
        "lastName": "Staggs",
        "wdl": "28W 11D 19L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 45667,
        "teamName": "Jacksonville FC ECNL G09",
        "status": 2,
        "clubID": 990,
        "initialSeed": 12,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/JOEL_HERNANDEZ_310dd57f9c604f8a9237bbafc95329d7_JFC_Logo.jpg",
        "firstName": "Ayomide",
        "lastName": "Leyimu",
        "wdl": "8W 2D 49L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 59913,
        "teamName": "Orlando City ECNL G09",
        "status": 2,
        "clubID": 77,
        "initialSeed": 13,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/_0ede1643f06145889d4f593cc8d4b0a1_LeagueTeam.png",
        "firstName": "Craig",
        "lastName": "Melton",
        "wdl": "9W 15D 28L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 50959,
        "teamName": "South Carolina United ECNL G09",
        "status": 2,
        "clubID": 1689,
        "initialSeed": 14,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Yoshikl_Sako_4009bd6b36b14e96ae900fd4b47b8b6f_SCUFC_TM_s.jpg",
        "firstName": "Adrian",
        "lastName": "Pinasco",
        "wdl": "21W 4D 30L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 49152,
        "teamName": "Tampa Bay United ECNL G09",
        "status": 2,
        "clubID": 874,
        "initialSeed": 15,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Hugo_Acosta_15ad970dfb1e4ea39726de93b4ff5a6e_tbu_2_star.png",
        "firstName": "Kayla",
        "lastName": "Corona Grimsley",
        "wdl": "13W 8D 36L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
      {
        "teamID": 59536,
        "teamName": "United Futbol Academy ECNL G09",
        "status": 2,
        "clubID": 3407,
        "initialSeed": 16,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/Jonathon_Lake_ce2883d65e694e248155158ad7164fa4_2020 UFA.png",
        "firstName": "Gilbert",
        "lastName": "Jean-Baptiste",
        "wdl": "44W 6D 9L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      }
    ],
    "teamsCountForFlightList": [
      {
        "teamsCount": 16,
        "flightID": 19789,
        "flightName": "ECNL"
      }
    ],
    "flightGroups": [
      {
        "flightgroupID": 26793,
        "flightID": 19789,
        "name": "Group A",
        "colorcode": null
      }
    ]
  }
}
```



Not Pertinent but interesting

/api/Player/get-commitment-status-data-by-playerID/{playerID}
This gives us the commitment status for a given player

I don't know how to find the playerID yet


The most important endpoint of all so far is this one

/api/Club/get-score-reporting-schedule-list/{clubID}/{eventID}
This returns the match results for a given club and event

So I would need to make this call for each club in the ECNL for the "ECNL Girls events"

Step 1. Determine the list of clubs in the event

Step 2. Call this for each club in the event caching the results

Each result looks like this

```json
{
  "result": "success",
  "data": {
    "reportedScore": 34,
    "unReportedScore": 56,
    "eventPastScheduleList": [
      {
        "matchID": 496876,
        "gameDate": "2023-09-09T10:00:00",
        "hometeamID": 59913,
        "homeTeamClubID": 77,
        "awayteamID": 59236,
        "awayTeamClubID": 30,
        "gameTime": "10:00:00",
        "flight": "ECNL",
        "division": "G2009",
        "homeclublogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/_0ede1643f06145889d4f593cc8d4b0a1_LeagueTeam.png",
        "awayclublogo": "https://s3.amazonaws.com/images.totalglobalsports.com/ClubImages/Logos/Concorde-Fire.jpg",
        "homeTeam": "Orlando City ECNL G09",
        "awayTeam": "Concorde Fire Premier ECNL G09",
        "complex": "Seminole",
        "venue": "Field 1 - Wayne Densch Trust",
        "scheduleID": 7965,
        "homeTeamScore": 1,
        "awayTeamScore": 1,
        "eventName": "ECNL Girls Southeast 2023-24",
        "eventLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/EventImages/Logos/8fec3d2f970c4ee2b9d76404cb824716_ECNL_South.png",
        "startDate": "2023-08-01T00:00:00",
        "endDate": "2024-07-01T00:00:00",
        "eventTypeID": 2
      }
   ]
  }
}
```

## References

* [ECNL](https://www.ecnlgirls.com/)
* [ECNL Schedule](https://www.ecnlgirls.com/schedule/)
* [Total Global Sports Swagger](https://public.totalglobalsports.com/swagger/index.html)
