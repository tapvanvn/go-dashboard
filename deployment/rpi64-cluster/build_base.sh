# we assume that 

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

pushd "$DIR/../../"

tag=latest

server_url=tapvanvn

docker build -t $server_url/rpi64_dashboard_base:$tag -f docker/rpi64_base.dockerfile ./

docker push $server_url/rpi64_dashboard_base:$tag

popd