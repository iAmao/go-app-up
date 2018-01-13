#!/bin/bash

mkdir -p "${HOME}/go-app-up"
cp "./main" "${HOME}/go-app-up"
echo "alias go-app-up='${HOME}/go-app-up/main'" > "${HOME}/.bash_profile";
source "${HOME}/.bash_profile";