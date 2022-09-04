checkError() {
    if [ $? -ne 0 ]; then 
        exit 1
    fi 
}

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

mkdir -p "$DIR/.temp/config"

go build -o "$DIR/.temp/daskboard" $DIR/main.go

rs=$?
if [ $rs -eq 0 ]; then 
    echo "SUCCESS"

    cp config/config_local.jsonc "$DIR/.temp/config/config.jsonc"
    cp -r static "$DIR/.temp/"

    app="$DIR/.temp/daskboard"
    kill $(lsof -t -i:8080)
    PORT=8080 $app
else
    echo "FAIL"
fi
