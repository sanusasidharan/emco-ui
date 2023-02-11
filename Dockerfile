# =======================================================================
# Copyright (c) 2017-2021 Aarna Networks, Inc.
# All rights reserved.
# ======================================================================
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#           http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# ========================================================================

# => Build container
FROM node:18-alpine as builder
WORKDIR /app
COPY package.json .
COPY package-lock.json .
RUN npm install -g npm@9.4.2
RUN npm install
COPY src ./src
COPY public ./public
# => Pass the reuired version
RUN REACT_APP_VERSION=v2.4.1 REACT_APP_PRODUCT=EMCO PUBLIC_URL=/ npm run build
#RUN REACT_APP_VERSION=v2.3 REACT_APP_PRODUCT=AMCOP npm run build

# => Run container
FROM  nginxinc/nginx-unprivileged

# Nginx config
COPY default.conf /etc/nginx/conf.d/

# Static build
COPY --from=builder /app/build /usr/share/nginx/html/

# Default port exposure
EXPOSE 8080
