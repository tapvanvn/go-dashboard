# we assume that 

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

pushd "$DIR/../../"

tag=$(<./version.txt)

server_url=tapvanvn

docker build -t $server_url/rpi64_dashboard:$tag -f docker/rpi64.dockerfile ./

docker push $server_url/rpi64_dashboard:$tag

popd