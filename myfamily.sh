curl --compressed --silent --show-error --max-time 10 https://honoriscentraleit.01-edu.org/assets/superhero/all.json |
    jq ".[] | select(.id == $HERO_ID)" | grep relatives | cut -d'"' -f4