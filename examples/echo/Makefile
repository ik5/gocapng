
output_name:=echo-server

build:
	@rm "${output_name}" 2> /dev/null && echo "${output_name} cleaned" || true
	@go build -o "${output_name}" *.go && echo "built ${output_name}"
	sudo setcap CAP_NET_RAW,CAP_NET_BIND_SERVICE+ep "${output_name}"

clean:
	@rm "${output_name}" 2> /dev/null && echo "${output_name} cleaned" || echo "nothing to clean"
