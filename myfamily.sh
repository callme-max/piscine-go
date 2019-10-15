curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --arg ID_Num "$HERO_ID" '.[] | select( .id == ($ID_Num|tonumber)) .connections.relatives' | sed 's/"//g'
