# we assume that 

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

pushd "$DIR/../../"

tag=$(<./version.txt)

server_url=tapvanvn

docker build -t $server_url/rpi_dashboard:$tag -f docker/rpi.dockerfile ./

docker push $server_url/rpi_dashboard:$tag

popd