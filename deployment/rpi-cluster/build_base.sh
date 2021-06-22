# we assume that 

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

pushd "$DIR/../../"

tag=latest

server_url=tapvanvn

docker build -t $server_url/rpi_dashboard_base:$tag -f docker/rpi_base.dockerfile ./

docker push $server_url/rpi_dashboard_base:$tag

popd