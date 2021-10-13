#!/usr/bin/env bats

@test "Add Different Elements To Nodes & Check For Successfull Sync" {
  ports="$(docker ps | awk '/set/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	node1="${ports_list[0]}"
	node2="${ports_list[1]}"

	response="$(curl -sS -i http://$node1/set/add --data '{"value": "1"}' | awk ' /HTTP/ {print $2}')" && [ "$response" == "200" ]
	response="$(curl -sS -i http://$node1/set/add --data '{"value": "2"}' | awk ' /HTTP/ {print $2}')" && [ "$response" == "200" ]
	response="$(curl -sS -i http://$node1/set/add --data '{"value": "3"}' | awk ' /HTTP/ {print $2}')" && [ "$response" == "200" ]

	response="$(curl -sS -i http://$node2/set/add --data '{"value": "4"}' | awk ' /HTTP/ {print $2}')" && [ "$response" == "200" ]
	response="$(curl -sS -i http://$node2/set/add --data '{"value": "5"}' | awk ' /HTTP/ {print $2}')" && [ "$response" == "200" ]

	response="$(curl -sS -X GET http://$node1/set/list)" && [ "$response" == "[1,2,3]" ]
	response="$(curl -sS -X GET http://$node2/set/list)" && [ "$response" == "[4,5]" ]

	response="$(curl -sS -i http://$node2/set/sync | awk ' /HTTP/ {print $2}')" && [ "$response" == "200" ]
	response="$(curl -sS -X GET http://$node1/set/list)" && [ "$response" == "[1,2,3,4,5]" ]
	response="$(curl -sS -X GET http://$node2/set/list)" && [ "$response" == "[1,2,3,4,5]" ]
}
