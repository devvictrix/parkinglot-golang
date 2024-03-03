# Parking Lot Application

This application simulates a parking lot management system, capable of handling car parking, leaving, and querying operations.

## Getting Started

Build and run the application using Docker. Two modes are supported: Interactive Mode and File Input Mode.

### Interactive Mode

1. Build the Docker image for interactive mode:
   ```sh
   docker build -f storage/private/dockerfiles/DockerfileInteractive -t crea-parkinglot-interactive .
   ```
2. Run the container
   ```sh
   docker run -it crea-parkinglot-interactive
   ```
3. Follow the prompts to enter commands.
   - create_parking_lot 6
   - park KA-01-HH-1234 White
   - park KA-01-HH-9999 White
   - park KA-01-BB-0001 Black
   - park KA-01-HH-7777 Red
   - park KA-01-HH-2701 Blue
   - park KA-01-HH-3141 Black
   - leave 4
   - status
   - park KA-01-P-333 White
   - park DL-12-AA-9999 White
   - registration_numbers_for_cars_with_color White
   - slot_numbers_for_cars_with_color White
   - slot_number_for_registration_number KA-01-HH-3141
   - slot_number_for_registration_number MH-04-AY-1111

### File Input Mode

1. Build the Docker image for file input mode:
   ```sh
   docker build -f storage/private/dockerfiles/DockerfileFileInput -t crea-parkinglot-fileinput .
   ```
2. Run the container, specifying the input file:
   ```sh
   docker run -it crea-parkinglot-fileinput
   ```
