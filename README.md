
[![GitHub license](https://img.shields.io/github/license/kryptodev/sports-event-timer.svg)](https://github.com/kryptodev/sports-event-timer/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/kryptodev/sports-event-timer)](https://goreportcard.com/report/github.com/kryptodev/sports-event-timer)
[![DeepScan grade](https://deepscan.io/api/projects/2329/branches/13742/badge/grade.svg)](https://deepscan.io/dashboard#view=project&pid=2329&bid=13742)
[![Build Status](https://travis-ci.org/kryptodev/sports-event-timer.svg?branch=master)](https://travis-ci.org/kryptodev/sports-event-timer)

<p align="center">
  <img src="https://github.com/kryptodev/sports-event-timer/blob/master/source/frontend/src/assets/images/logo-alt.png?raw=true" alt="Sports Event Timing"/>
</p>

# Sports Event Timing


Sports Event Timing' demonstrates the simulation of a race among dummy participants, 
both server-sided and client-sided.

The simulation is a choice that can be run either solely on the client-side, concurrently 
using reactjs redux-sagas OR entirely on the server side where reactjs ensures to update the 
client-side accordingly.

<p align="center">
  <img src="https://github.com/kryptodev/sports-event-timer/blob/master/source/frontend/src/assets/images/demo.gif?raw=true" alt="App Demo"/>
</p>

### Race Preface

Every athlete carries a chip with a unique identifier. 
There are two set time points in the race, both of which also carrying unique chip identifiers,
individually trigger a signal to the server side, indicating the athlete has reached said time point.

###### Client Side Simulation

In case of the client side simulation, the client fetches dummy race participants, and dummy 
time points. Concurrently, the sagas then simulate the athletes running on the track, where
whenever any given athlete trips over a given time point, the following information is sent
to the server:
1. Athlete Chip Identifier
2. Time taken for the athlete to reach that time point
3. Time Point Chip Identifier

The server then acnkowledges the request using REST endpoints and then stores the information in the database.

###### Server Side Simulation

On the other hand, for the server side simulation, the server executes a goroutine-based race among
athletes, stores their data in the database and makes that data available for client side fetching
via REST endpoints. The server also uses similar mechanics to detect athlete chip identifiers
and time point chip identifiers.  


### Prerequisites

```
1. Go development environment and;
2. NPM/Yarn.
```

### Installing

On a Windows machine, project can be built using the build.bat file.
Subsequently, it will create a dist folder that will contain the production ReactJS files
as well as the backend binary. 

```
1. Clone the project in your GOPATH or download the repository as zip and extract it in your GOPATH.
2. Change directory into the project root.
3. Run ./build.bat.
```  

Alternatively, production files can also be generated separately(backend and frontend) as follows:

```
1. Clone the project in your GOPATH or download the repository as zip and extract it in your GOPATH.
2. Change directory into the %project root%/source/backend folder.
3. Run go build to get the built executable.
4. Change directory into the %project root%/source/frontend folder.
5. Run 'npm install' or just 'yarn'(project uses yarn) to install dependencies.
6. Run 'npm run build' or 'yarn build' to generate production files that will be placed inside %project root%/dist/ folder.
7. Make sure the server executable is in the same folder as all other distribution files i. e. same folder as index.html
```

**Backend REST API Overview**
----
  _The backend uses the package 
  [httprouter](https://github.com/julienschmidt/httprouter "httprouter") as its main mux to server REST Endpoints._
  
  By default, the server port is **8082**
  


### Server Simulated Race EndPoints
| endpoint      | method             | description                       |
|:--------------|:------------------|:----------------------------------|
| `/server/race-simulation/start`      |GET| *Starts of the Race Event*
| `/server/race-simulation/stop`    |GET| *Stops the Race Event*|
| `/server/race-simulation/fetch/live-standings` |GET| *Returns the list of athletes with their race-relevant attributes in a live race* |
| `/server/race-simulation/fetch/last-standings`      |GET| *Returns the list of athletes with their race-relevant attributes from the last race whereby, a race is currently not live* |

### Sample Calls
**Start Race**

Request:
  ```
  curl 'http://localhost:8082/server/race-simulation/start'
  ```
Response:
```json
{"responseMessage":"starting race","athletes":[{"startNumber":1,"fullName":"Manish Singh","identifier":"bbk15uu8auq2p50gu1i0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":2,"fullName":"Madhushree Singh","identifier":"bbk15uu8auq2p50gu1ig","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":3,"fullName":"Siim Kaspar Uustalu","identifier":"bbk15uu8auq2p50gu1j0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":4,"fullName":"Eliisabeth Käbin","identifier":"bbk15uu8auq2p50gu1jg","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":5,"fullName":"Ahti Liin","identifier":"bbk15uu8auq2p50gu1k0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1}]}
```
  
---

**Stop Race**

Request:
  ```
  curl 'http://localhost:8082/server/race-simulation/stop'
  ```
Response:
```json
{"responseMessage":"stopping race"}
```
  
---
**Get live Athlete values in a live Race**

Request:
  ```
  curl 'http://localhost:8082/server/race-simulation/fetch/live-standings'
  ```
Response:
```json
{"responseMessage":"race is in process, showing live data","athletes":[{"startNumber":1,"fullName":"Manish Singh","identifier":"bbk15uu8auq2p50gu1i0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":2,"fullName":"Madhushree Singh","identifier":"bbk15uu8auq2p50gu1ig","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":3,"fullName":"Siim Kaspar Uustalu","identifier":"bbk15uu8auq2p50gu1j0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":4,"fullName":"Eliisabeth Käbin","identifier":"bbk15uu8auq2p50gu1jg","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":5,"fullName":"Ahti Liin","identifier":"bbk15uu8auq2p50gu1k0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1}]}
```
  
---
**Get Athlete values in a from the last Race**

Request:
  ```
  curl 'http://localhost:8082/server/race-simulation/fetch/last-standings'
  ```
Response:
```json
{"responseMessage":"race is currently not in process, showing last race data","athletes":[{"startNumber":1,"fullName":"Manish Singh","identifier":"bbk15uu8auq2p50gu1i0","location":209,"timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":2,"fullName":"Madhushree Singh","identifier":"bbk15uu8auq2p50gu1ig","location":203,"timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":3,"fullName":"Siim Kaspar Uustalu","identifier":"bbk15uu8auq2p50gu1j0","location":226,"timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":4,"fullName":"Eliisabeth Käbin","identifier":"bbk15uu8auq2p50gu1jg","location":201,"timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":5,"fullName":"Ahti Liin","identifier":"bbk15uu8auq2p50gu1k0","location":201,"timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1}]}
```
  
---

### Client Simulated Race EndPoints
| endpoint      | method             | description                       |
|:--------------|:------------------|:----------------------------------|
| `/client/race-simulation/start`      |GET| *Starts of the Race Event* |
| `/client/race-simulation/stop`    |GET| *Stops the Race Event* |
| `/client/race-simulation/fetch/current-standings` |GET| *Returns the list of athletes with their race-relevant attributes in a live race* |
| `/client/race-simulation/push/current-standings`      |POST| *Accepts athlete attributes which it then stores in the database in a live race* |

### Sample Calls
**Start Race**

*Request:*
  ```
  curl 'http://localhost:8082/client/race-simulation/start'
  ```
*Response:*
```json
{"responseMessage":"starting race","athletes":[{"startNumber":1,"fullName":"Manish Singh","identifier":"bbk0q968auq2p50gu1b0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":2,"fullName":"Madhushree Singh","identifier":"bbk0q968auq2p50gu1bg","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":3,"fullName":"Siim Kaspar Uustalu","identifier":"bbk0q968auq2p50gu1c0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":4,"fullName":"Eliisabeth Käbin","identifier":"bbk0q968auq2p50gu1cg","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":5,"fullName":"Ahti Liin","identifier":"bbk0q968auq2p50gu1d0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1}],"timePoints":[{"name":"Corridor Timepoint","location":800,"identifier":"bbk0q968auq2p50gu1dg"},{"name":"Finish Line Timepoint","location":1000,"identifier":"bbk0q968auq2p50gu1e0"}]}
```
  
---
**Stop Race**

*Request:*
  ```
  curl 'http://localhost:8082/client/race-simulation/stop'
  ```
*Response:*
```json
{"responseMessage":"stopping race"}
```
  
---
**Fetch Current Athletes in the Race**

*Request:*
  ```
  curl 'http://localhost:8082/client/race-simulation/fetch/current-standings'
  ```
*Response:*
```json
{"responseMessage":"race is in process, showing live data","athletes":[{"startNumber":1,"fullName":"Manish Singh","identifier":"bbk15hu8auq2p50gu1eg","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":2,"fullName":"Madhushree Singh","identifier":"bbk15hu8auq2p50gu1f0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":3,"fullName":"Siim Kaspar Uustalu","identifier":"bbk15hu8auq2p50gu1fg","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":4,"fullName":"Eliisabeth Käbin","identifier":"bbk15hu8auq2p50gu1g0","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":5,"fullName":"Ahti Liin","identifier":"bbk15hu8auq2p50gu1gg","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1}]}
```
  
---
**Push updated Athletes values from Client to Server**

*Request:*
  ```
  curl 'http://localhost:8082/client/race-simulation/push/current-standings' --data-binary '{"requestCommand":"commit feed","athletes":[{"timeElapsed":33614,"identifier":"bbk0ntm8auq2p50gu17g","timePointIdentifier":"bbk0ntm8auq2p50gu1a0"}]}'
  ```
*Response:*
```json
{"responseMessage":"race is in process, data was committed, showing committed data","athletes":[{"startNumber":1,"fullName":"Manish Singh","identifier":"bbk0ntm8auq2p50gu17g","inFinishCorridor":true,"timeTakenToReachFinishCorridor":33614,"timeTakenToFinish":-1},{"startNumber":2,"fullName":"Madhushree Singh","identifier":"bbk0ntm8auq2p50gu180","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":3,"fullName":"Siim Kaspar Uustalu","identifier":"bbk0ntm8auq2p50gu18g","timeTakenToReachFinishCorridor":-1,"timeTakenToFinish":-1},{"startNumber":4,"fullName":"Eliisabeth Käbin","identifier":"bbk0ntm8auq2p50gu190","inFinishCorridor":true,"timeTakenToReachFinishCorridor":33485,"timeTakenToFinish":-1},{"startNumber":5,"fullName":"Ahti Liin","identifier":"bbk0ntm8auq2p50gu19g","inFinishCorridor":true,"timeTakenToReachFinishCorridor":33647,"timeTakenToFinish":-1}]}
```
  
---

* **Notes:**

_Ideally, the response message should be in the format of 'success' or 'fail' and should use status codes more effectively to better adhere to REST EndPoint standards. 
The backend uses defensive analysis to ascertain the condition of the race event before responding to the client. The client side,
might however, lack coping mechanisms to accurately reflect the backend's status responses. This is true at the time of writing._

_1, May, 2018_ 


___

## To-do
* Thorough Unit-Testing.
* Replace mutexes in backend with pure channel-based solution.
* Stories on frontend. 

## Built With

* [httprouter](https://github.com/julienschmidt/httprouter) - The backend router used
* [react](https://reactjs.org) - The frontend javascript library used
* [redux](https://redux.js.org/) - State Management
* [redux-saga](https://github.com/redux-saga/redux-saga) - To handle side-effects and other asynchronous tasks
* [sqlite3](https://www.sqlite.org/index.html) - Relational drop-in database for quick demo

## Author

* **Manish Singh** - *Initial work* - [kryptodev](https://github.com/kryptodev)

## License

This project is licensed under The Unlicence License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* Mooncascade for the opportunity given(https://www.mooncascade.com)
