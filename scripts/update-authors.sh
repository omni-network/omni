#!/bin/bash

# This script updates the AUTHORS file with the list of all the authors that
# have contributed in the code. It removes already existing file and re-creates
# it every time it is executed.
#
# PS: This script needs to be executed from the root of the omni repo
#     ex: ./scripts/update-authors.sh
#

AUTHORS_FILE="AUTHORS"

OUTPUT=`git log --format='%an <%ae>' | sort | uniq`
cat > $AUTHORS_FILE <<EOL
${OUTPUT}
EOL