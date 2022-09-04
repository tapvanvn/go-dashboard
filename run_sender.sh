checkError() {
    if [ $? -ne 0 ]; then 
        exit 1
    fi 
}

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

mkdir -p "$DIR/.temp/config"

go build -o "$DIR/.temp/sender" $DIR/sender/main/main.go

rs=$?
if [ $rs -eq 0 ]; then 
    echo "SUCCESS"

    cp config/config_local.jsonc "$DIR/.temp/config/config.jsonc"

    app="$DIR/.temp/sender"
    kill $(lsof -t -i:8081)
    PORT=8081 $app
else
    echo "FAIL"
fi
