#!/bin/bash

SUPERVISOR_CONFIG=${DISCOVERY_PROXY:-nginx}
SUPERVISOR_CONFIG="/etc/supervisor/conf.d/supervisord_$SUPERVISOR_CONFIG.conf"

/usr/bin/supervisord  -c $SUPERVISOR_CONFIG