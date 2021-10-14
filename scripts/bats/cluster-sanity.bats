#!/usr/bin/env bats

@test "Check Replicas Count" {
  count="$(docker ps | grep -c set)"
  [ "$count" -eq 2 ]
}

@test "Check Replicas Are Avaialable" {
  ports="$(docker ps | awk '/set/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	for port in "${ports_list[@]}"; do
			response="$(curl -sS -X GET http://$port/)"
			[ "$response" == "Hello World Set Node" ]
	done 
}

@test "Writes Are Succesfull" {
  ports="$(docker ps | awk '/set/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	for port in "${ports_list[@]}"; do
			response="$(curl -sS -i http://$port/set/add --data '{"values": [1,2]}' | awk ' /HTTP/ {print $2}')" && [ "$response" == "200" ]
	done 
}

@test "Reads Are Succesfull" {
  ports="$(docker ps | awk '/set/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	for port in "${ports_list[@]}"; do
			response="$(curl -sS -X GET http://$port/set/list)"
			[ "$response" == "[1,2]" ]
	done 
}

@test "Writes Are Idempotent" {
  ports="$(docker ps | awk '/set/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	for port in "${ports_list[@]}"; do
			response="$(curl -sS -i http://$port/set/add --data '{"values": [1,2]}' | awk ' /HTTP/ {print $2}')" && [ "$response" == "200" ]
	done

	for port in "${ports_list[@]}"; do
			response="$(curl -sS -X GET http://$port/set/list)"
			[ "$response" == "[1,2]" ]
	done
}

@test "Set Debug Clear" {
  ports="$(docker ps | awk '/set/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	node1="${ports_list[0]}"
	node2="${ports_list[1]}"

	response="$(curl -sS -i http://$node1/set/debug/clear | awk ' /HTTP/ {print $2}')" && [ "$response" == "200" ]
	response="$(curl -sS -i http://$node2/set/debug/clear | awk ' /HTTP/ {print $2}')" && [ "$response" == "200" ]

	response="$(curl -sS -X GET http://$node1/set/list)" && [ "$response" == "[]" ]
	response="$(curl -sS -X GET http://$node2/set/list)" && [ "$response" == "[]" ]
}
