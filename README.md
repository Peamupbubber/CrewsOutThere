# CrewsOutThere
Welcome to CrewsOutThere! This codebase is an instance of my Computer Science major senior project at Western Washington University that ran from January 2023 - December 2023.

CoT is a fully functioning text-messaging system created for the Bellingham Civil Air Patrol, an auxilary of the US Airforce, intended to get flight crews in contact with each other on short notice.

Functionality includes: inviting other users to the system (intended only for use within CAP), adding yourself to roles and airports, viewing roles and airports (both yours and all available options), requesting members of a flight crews (by specifying the role and airport the user needs), and help message features.

Documents created for this project: Vision & Scope, Software Requirements Specification, Schedule, ER Diagrams, Use Case Diagrams, Test Cases, Project Documentation, and Design Specs for future implementations.

Technologies used for this project: Golang, MySQL, Docker, Bash, Linux, NetBSD.

There are 4 distinct folders in this repo, each one serves a specific purpose:

 - crewCLI: This is a command line interface developed for modifying the database easily. The only person who has access to this program is our client from CAP. This interface allows you to update the available roles and airports within the system, as well as manage users, allowing for the removal of single users or all users invited by a specific user.

 - crewsControl: This is a mirror of the actual system intended for testing. There is an sql script located outside the folder that can be used to populate a local Docker database and test new features without interfering with the deployment build folder.

 - crewsOutThere: this is all of the code for the actual system, including connecting to our http web server for interacting with Twilio, our SMS provider, connecting to our MySQL database, logging information for debugging and system management, and the command handling with uses regexes to parse user input and backend logic writen in Golang to handle the required response, who should be notified, and how the DB should be affected. There is also an encryption folder that holds information about the system.

 - iscotDown: uses a bash script to notify a developer if the system has shut down unexpectedly.

My primary contribution was to write most of the code found in the crewsOutThere/command folder, primarily these go files: cleanupDB, commandFly, commandParser, commandUtilities, and parserCalls.

I also co-designed the architecture of the system as well as co-wrote all of the project documentation, which can be found on the project's wiki page.
