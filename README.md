# Parking Lot

## Problem Statement

I own a parking lot that can hold up to 'n' cars at any given point in time. Each slot is
given a number starting at 1 increasing with increasing distance from the entry point
in steps of one. I want to create an automated ticketing system that allows my
customers to use my parking lot without human intervention.

When a car enters my parking lot, I want to have a ticket issued to the driver. The
ticket issuing process includes us documenting the registration number (number
plate) and the colour of the car and allocating an available parking slot to the car
before actually handing over a ticket to the driver (we assume that our customers are
nice enough to always park in the slots allocated to them). The customer should be
allocated a parking slot which is nearest to the entry. At the exit the customer returns
the ticket which then marks the slot they were using as being available.

Due to government regulation, the system should provide me with the ability to findout:

1. Registration numbers of all cars of a particular colour.
2. Slot number in which a car with a given registration number is parked.
3. Slot numbers of all slots where a car of a particular colour is parked.

The system allows input in two ways:-

1. Interactive command prompt based shell where commands can be typed in.
2. Accept a filename as a parameter at the command prompt and read the commands from that file

> **Note:- Project should be placed in `$GOPATH/src/github.com/singhmeghna79`**

### To install all dependencies, compile and run tests:
> run `./bin/setup` from root directory of the project. 

### To run the code so it accepts input from a file:
> run `./bin/parking_lot file_inputs.txt` from root directory of the project.
