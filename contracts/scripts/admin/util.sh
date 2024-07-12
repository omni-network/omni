# shared admin script utils

# check if required environment variables are set
# usage: require_env <var1> <var2> ...
require_env() {
  for var in $@; do
    if [ -z "${!var}" ]; then
      echo "Missing required environment variable: $var"
      return 1
    fi
  done
}
