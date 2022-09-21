[//]: # "Copyright (c) 2017-2021 Aarna Networks, Inc."
[//]: # "All rights reserved."
[//]: # "Licensed under the Apache License, Version 2.0 (the \"License\");"
[//]: # "you may not use this file except in compliance with the License."
[//]: # "You may obtain a copy of the License at"
[//]: # "          http://www.apache.org/licenses/LICENSE-2.0"
[//]: # "Unless required by applicable law or agreed to in writing, software"
[//]: # "distributed under the License is distributed on an \"AS IS\" BASIS,"
[//]: # "WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied."
[//]: # "See the License for the specific language governing permissions and"
[//]: # "limitations under the License."

## Local setup

for running the app in a local setup first install the dependencies by running `npm install`.
Then run `startup.sh`

## Production build

for creating a production build, run `npm run build`. A production ready build will be available at /build directory

## Available scripts

### `startup.sh`

This script basically calls npm start.
This script runs the app in the development mode.<br />
Before running the script update the backend address if backend is not running locally.
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />
You will also see any lint errors in the console.

### `npm run build`

Builds the app for production to the `build` folder.<br />
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br />

## Building docker image

To build docker images of the EMCO GUI, run the following commands:

    docker build -t emco-gui:latest .
    docker build -t emco-gui-dbhook:latest db_udpate
    docker build -t emco-gui-authgw:latest authgateway
    docker build -t emco-gui-middleend:latest guimiddleend
