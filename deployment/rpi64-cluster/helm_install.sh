DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
tag=$(<../../version.txt)

namespace=default
if [ -f "$DIR/../../namespace.txt" ]; then
    namespace=$(<$DIR/../../namespace.txt)
fi

helm install go-dashboard $DIR/../../helm/rpi64-cluster --set image.tag=$tag,namespace=$namespace --namespace=$namespace