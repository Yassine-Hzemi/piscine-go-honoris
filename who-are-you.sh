curl --max-time 10 "https://honoriscentraleit.01-edu.org/assets/superhero/all.json" |
    jq '.[] | select(.id == 70) | .name'

