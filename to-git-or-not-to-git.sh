curl -s https://honoriscentraleit.01-edu.org/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"taha.ba\"}}){id}}"}' | tr -dc '0-9' 
echo -e
