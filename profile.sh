#!/bin/bash

# Get the directory where the script is located
COUNTER_SOURCE_CODE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Generate mock files for all services, putting the results in the proper file.
# If you update any service, run this function to update the service for the tests.
counterRegenerateMockServices () {
  packageName="mockgen"

  # Check if mockgen is installed
  which "${packageName}"

  # If mockgen is installed (exit code 0)
  if [[ "$?" == "0" ]]; then
    MOCK_FOLDER="Services"
    SERVICE_DIR="${COUNTER_SOURCE_CODE_DIR}/src/${MOCK_FOLDER}"

    # Ensure the mocks directory exists
    mkdir -p "${SERVICE_DIR}/mocks/${MOCK_FOLDER}_mocks"

    # Find all _service.go files in the SERVICE_DIR
    SERVICE_FILES=$(find "${SERVICE_DIR}" -maxdepth 1 -name "*_service.go")

    # Loop over each _service.go file
    for SERVICE_FILE in ${SERVICE_FILES}; do
      if [[ -f ${SERVICE_FILE} ]]; then
        # Extract the file's basename (without the directory and _service.go suffix)
        FOLDER_NAME="${SERVICE_FILE##*/}"
        FOLDER_NAME="${FOLDER_NAME%_service.go}"

        # Generate mock file with mockgen
        echo "Running mockgen for ${SERVICE_FILE}, output: mocks/${MOCK_FOLDER}_mocks/${FOLDER_NAME}_mock.go"

        mockgen \
          -source=${SERVICE_FILE} \
          -destination="${SERVICE_DIR}/mocks/${MOCK_FOLDER}_mocks/${FOLDER_NAME}_mock.go" \
          -package="${MOCK_FOLDER}_mocks" \
          -mock_names Service=Mock_${FOLDER_NAME}

        # Check if mockgen succeeded
        if [[ $? -ne 0 ]]; then
          echo "mockgen failed for ${SERVICE_FILE}. Exiting."
          exit 1
        fi
      fi
    done

    echo "All mock files generated successfully."

  else
    # If mockgen is not found, provide installation instructions
    echo "Error: missing required package 'mockgen'. Please run the following commands and try again:"
    echo "1. Install protobuf, and then run:"
    echo "$ go get -u github.com/golang/protobuf/protoc-gen-go"
  fi
}
