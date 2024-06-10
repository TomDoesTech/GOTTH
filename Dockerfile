FROM golang:1.22-alpine AS base

WORKDIR /app

# Install dependencies required for the build process
RUN apk add --no-cache make curl

RUN curl -L https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.4/tailwindcss-linux-arm64 -o tailwindcss \
    && chmod +x tailwindcss

RUN go install github.com/a-h/templ/cmd/templ@latest

FROM base AS development

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN make build

FROM base AS production

COPY --from=development --chown=golang:golang /app /app

USER golang
