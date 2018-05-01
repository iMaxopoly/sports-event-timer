**FrontEnd Overview**
----
_The frontend is based on ReactJS, using Redux for state-management
and Redux Saga as a side-effects middleware._

Bootstrap is used as the CSS framework of choice and SCSS is used to
modify the theme.

[Click here](https://github.com/kryptodev/sports-event-timer/blob/master/source/frontend/src/assets/images/componentDiagram.png?raw=true) to see the component diagram of the application structure. 

* **Notes:**

_CSS modules in combination with SCSS files that allow for global CSS/SCSS
dependencies(or from node_modules) to stay intact would have been a much proficient set up, not found in this project._
Usage of FlipMove external dependency seems to be bad with performance, however, ideal for this project,
compared to my slide animation attempt using react-transition-group, but definitely something to retry.

_1, May, 2018_