#!/bin/sh

rsync_file="10.29.8.25::download_rsync_0/20141218/77026291_41344886_0.zip"

rsync -avz --password-file=.download_password  rsync_auth@$rsync_file .
