# ECNL Module

A Go module to provide access to ECNL data.

![GitHub Actions](https://github.com/jedi-knights/ecnl/workflows/CI/badge.svg)

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


## References

* [ECNL](https://www.ecnlgirls.com/)
* [ECNL Schedule](https://www.ecnlgirls.com/schedule/)
* [Total Global Sports Swagger](https://public.totalglobalsports.com/swagger/index.html)