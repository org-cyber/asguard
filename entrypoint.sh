#!/bin/sh
# Start both services in the background
/root/general_engine &
/root/face_service &

# Start the reverse proxy (foreground, keeps container alive)
/root/proxy