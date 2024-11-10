#!/bin/sh
set -e  # Exit on error

cd /app

if [ "$ENVIRONMENT" = "development" ]; then
    make tailwind-watch
    make templ-watch
else
    make tailwind-build
    make templ-generate
fi
