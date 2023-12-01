(
    cd ../contracts
    mkdir -p abis
    mkdir -p bins
    for FILE in *.sol
    do
        nm=${FILE%.*}
        lcnm=$(echo "$nm" | awk '{print tolower($0)}')
        solc --abi $FILE -o abis --overwrite >/dev/null 2>&1
        solc --bin $FILE -o bins --overwrite >/dev/null 2>&1
        mkdir -p ../src/pkg/abis/${nm}
        abigen --bin=bins/${nm}.bin --abi=abis/${nm}.abi --pkg=${lcnm} --out=../src/pkg/abis/${nm}/${nm}.go
    done
    rm -r abis bins
)