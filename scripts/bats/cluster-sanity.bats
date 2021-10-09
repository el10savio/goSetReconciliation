#!/usr/bin/env bats

@test "Check Replicas Count" {
  count="$(docker ps | grep -c set)"
  [ "$count" -eq 2 ]
}

@test "Check Replicas Are Avaialable" {
  ports="$(docker ps | awk '/set/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g')"
	IFS=', ' read -r -a ports_list <<< "$ports"

	for port in "${ports_list[@]}"; do
			response="$(curl -sS -X GET http://$port/)"
			[ "$response" == "Hello World Set Node" ]
	done 
}