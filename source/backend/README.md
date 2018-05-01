**REST API Overview**
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