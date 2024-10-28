######################################
# Prepare npm_builder
######################################
FROM node:18 as npm_builder
WORKDIR /go/src/github.com/Allen-Career-Institute/flagr

# setting default arg as dev currently to test the flow in dev/sandbox
ARG ENVIRONMENT=dev
ENV ENVIRONMENT=${ENVIRONMENT}
COPY . .
ARG FLAGR_UI_POSSIBLE_ENTITY_TYPES=null
ENV VUE_APP_FLAGR_UI_POSSIBLE_ENTITY_TYPES ${FLAGR_UI_POSSIBLE_ENTITY_TYPES}
RUN make build_ui

######################################
# Prepare go_builder
######################################
FROM golang:1.21-alpine as go_builder
WORKDIR /go/src/github.com/Allen-Career-Institute/flagr

RUN apk add --no-cache build-base git make
COPY . .
RUN make build

FROM alpine

COPY --from=go_builder /go/src/github.com/Allen-Career-Institute/flagr/flagr .

ENV HOST=0.0.0.0
# ENV PORT=3000 (for local testing)
ENV PORT=18000

ENV FLAGR_DB_DBDRIVER=sqlite3
ENV FLAGR_DB_DBCONNECTIONSTR=/data/demo_sqlite3.db
# for local testing set FLAGR_RECORDER_ENABLED to false
ENV FLAGR_RECORDER_ENABLED=false

# JWT Environment Variables
ENV FLAGR_JWT_AUTH_ENABLED=true
ENV FLAGR_JWT_AUTH_DEBUG=true
ENV FLAGR_JWT_AUTH_WHITELIST_PATHS="/api/v1/health,/api/v1/evaluation,/login,/callback,/static,/favicon.ico"
ENV FLAGR_JWT_AUTH_EXACT_WHITELIST_PATHS=",/,/login,/callback"
ENV FLAGR_JWT_AUTH_COOKIE_TOKEN_NAME="access_token"
ENV FLAGR_JWT_AUTH_SECRET="4fg88hsdf04ea6f26e3d1d3c5f17aba9770fab35de1400d7de89212cb2d32e240323f33a261181bd60b779cc97d15affbf6c51128b07c26964911c0b16b155f6c0c4e6f96e55649d03e9cbc7ca9681102e267be067ccb611c6cb35e07612c0449358a1e0cc7f638ac7c228f25e9650d8c6ea72e2619ef5474a11d9733afec91a"
ENV FLAGR_JWT_AUTH_NO_TOKEN_STATUS_CODE=307
# temp changes to test flow in dev end to end
ENV FLAGR_JWT_AUTH_NO_TOKEN_REDIRECT_URL="http://127.0.0.1:18000/login"
# ENV FLAGR_JWT_AUTH_NO_TOKEN_REDIRECT_URL="http://localhost:18000/login" (to be overriden in k8 yaml file as per env)
ENV FLAGR_JWT_AUTH_USER_CLAIM=uid
ENV FLAGR_JWT_AUTH_SIGNING_METHOD=HS256

# CORS Environment Variables
ENV FLAGR_CORS_ALLOWED_METHODS="GET,POST,PUT,DELETE,PATCH,OPTIONS"
ENV FLAGR_CORS_ALLOWED_HEADERS="*"

COPY --from=npm_builder /go/src/github.com/Allen-Career-Institute/flagr/browser/flagr-ui/dist ./browser/flagr-ui/dist

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

COPY --chown=appuser:appgroup ./buildscripts/demo_sqlite3.db /data/demo_sqlite3.db

# EXPOSE 3000 (for local testing)
EXPOSE 18000

CMD "./flagr"
