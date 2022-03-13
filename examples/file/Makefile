
output_name:=file-example

build: clean
	@go build -o "${output_name}" *.go && echo "built ${output_name}"
	sudo setcap CAP_DAC_OVERRIDE,CAP_SETPCAP,CAP_SETFCAP+ep "${output_name}"

clean:
	@rm "${output_name}" 2> /dev/null && echo "${output_name} cleaned" || echo "nothing to clean"

