# Morning Box By Team Asdos Sisop
## Hackfest 2023

### Table of Contents
- [My Team](#my-team)
- [Description](#description)
- [Solution](#solution)
- [Tech Stack](#tech-stack)
- [Documentation Postman](#documentation-postman)
- [How to run](#how-to-run)

### My Team
- [Nur Muhammad Ainul Yaqin](https://www.linkedin.com/in/nurmuhammad22/)
- [Benedictus Wicaksono](https://www.linkedin.com/in/benewicaksono/)
- [Kurnia Cahya Febryanto](https://www.linkedin.com/in/kurnia-cahya-febryanto/)
- [Christhoper Marcelino Mamahit](https://www.linkedin.com/in/christhopermarcelino/)

### Description 
A breakfast habituation platform for overseas students as well as monitoring for parents. Morning Box goes hand in hand and as an operational supporter of distributed food catering.

### Solution
1. ✅ Early Food Delivery </br>
    Breakfast delivery when you're not busy or on the move in the morning (05.30-07.30)!
2. ✅ Reward System </br>
    The more you have your breakfast, the more bonuses you can get! Strike is the key!
3. ✅ Normalizing the Breakfast Ecosystem </br>
    Invite your friends, share your breakfast, and let your parents know about your breakfast habit progress!

### Tech Stack
- Golang
    - Gin
- Flutter
- Firebase
    - Firestore
- Google Compute Engine
- Tensorflow

# Documentation Postman
https://documenter.getpostman.com/view/19606163/2s93JtQ44o

# How to run
1. Clone this repository
    ```bash
    git clone https://github.com/kurniacf/morning-box-hackfest-backend.git
    ```
2. Install all dependencies
    ```bash
    go get ./...
    ```
3. Run the server
    ```bash
    go run . 
    ```
    or
    ```bash
    go run main.go
    ```
    description:
    - `go run .` to run the server
    - `go run . -migrate` to run the server and migrate the database
    - `go run . -seed` to run the server and seed the database

4. Run the client
