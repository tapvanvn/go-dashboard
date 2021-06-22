# we assume that 

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

pushd "$DIR/../../"

tag=$(<./version.txt)

server_url=gcr.io/mydefipet

docker build -t $server_url/dashboard:$tag -f docker/gcloud.dockerfile ./

docker push $server_url/dashboard:$tag

popd