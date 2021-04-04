#!/bin/bash
PORT=$1

echo ">Basic ping test: "
wget -qO- "http://localhost:${PORT}/ping"
echo ""

echo ">Adding candidate Carl"
wget -qO- "http://localhost:${PORT}/api/user/" \
    --header='Content-Type:application/json' \
    --post-data='{"name": "Carl", "mail": "carl@gmail.com", "type": "candidate"}'
echo ""

echo ">Adding candidate Ines"
wget -qO- "http://localhost:${PORT}/api/user/" \
    --header='Content-Type:application/json' \
    --post-data='{"name": "Ines", "mail": "ines@gmail.com", "type": "interviewer"}'
echo ""

echo ">Adding candidate Ingrid"
wget -qO- "http://localhost:${PORT}/api/user/" \
    --header='Content-Type:application/json' \
    --post-data='{"name": "Ingrid", "mail": "ingrid@gmail.com", "type": "interviewer"}'
echo ""

echo ">Adding Carl's slots"
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Carl", "week":33, "day": "monday", "hour_start":9, "hour_end":10}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Carl", "week":33, "day": "tuesday", "hour_start":9, "hour_end":10}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Carl", "week":33, "day": "wednesday", "hour_start":9, "hour_end":10}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Carl", "week":33, "day": "thursday", "hour_start":9, "hour_end":10}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Carl", "week":33, "day": "friday", "hour_start":9, "hour_end":10}'
echo ""

echo ">Adding Ines's slots"
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Ines", "week":33, "day": "monday", "hour_start":9, "hour_end":16}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Ines", "week":33, "day": "tuesday", "hour_start":9, "hour_end":16}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Ines", "week":33, "day": "wednesday", "hour_start":9, "hour_end":16}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Ines", "week":33, "day": "thursday", "hour_start":9, "hour_end":16}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Ines", "week":33, "day": "friday", "hour_start":9, "hour_end":16}'
echo ""

echo ">Adding Ingrids's slots"
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Ingrid", "week":33, "day": "monday", "hour_start":12, "hour_end":18}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Ingrid", "week":33, "day": "tuesday", "hour_start":9, "hour_end":12}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Ingrid", "week":33, "day": "wednesday", "hour_start":12, "hour_end":18}'
echo ""
wget -qO- "http://localhost:${PORT}/api/slots/" \
    --header='Content-Type:application/json' \
    --post-data='{"name":"Ingrid", "week":33, "day": "thursday", "hour_start":9, "hour_end":12}'
echo ""

echo ">See possible slots for meeting with Carl, Ines and Ingrid"
wget -qO- "http://localhost:${PORT}/api/schedule/" \
    --header='Content-Type:application/json' \
    --method=GET \
    --body-data='{"candidate": "Carl", "interviewers": "Ines,Ingrid"}'
