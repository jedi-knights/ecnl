# ecnl
A Go module to provide access to ECNL data.

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

## References

* [ECNL](https://www.ecnlgirls.com/)
* [ECNL Schedule](https://www.ecnlgirls.com/schedule/)
* [Total Global Sports Swagger](https://public.totalglobalsports.com/swagger/index.html)