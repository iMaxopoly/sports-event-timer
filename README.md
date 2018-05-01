
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
1. Clone or download the repository as zip
2. Place the folder inside your GOPATH
3. Change directory into the project root
4. Run ./build.bat
```  

Alternatively, production files can also be generated separately(backend and frontend) as follows:

```
1. Clone or download the repository as zip
2. Place the folder inside your GOPATH
3. Change directory into the %project root%/source/backend folder
4. Run go build to get the built executable
5. Change directory into the %project root%/source/frontend folder
6. Run 'npm install' or just 'yarn'(project uses yarn) to install dependencies
7. Run 'npm build' or 'yarn build' to generate production files that will be placed 
inside %project root%/dist/ folder
```

## Built With

* [httprouter](https://github.com/julienschmidt/httprouter) - The backend router used
* [react](https://reactjs.org) - The frontend javascript library used
* [redux](https://redux.js.org/) - State Management
* [redux-saga](https://github.com/redux-saga/redux-saga) - To handle side-effects and other asynchronous tasks

## Author

* **Manish Singh** - *Initial work* - [kryptodev](https://github.com/kryptodev)

## License

This project is licensed under The Unlicence License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* React & Go creators
* Mooncascade for the opportunity given(https://www.mooncascade.com)
